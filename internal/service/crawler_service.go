package service

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"twvixstock/internal/model"
	"twvixstock/internal/repository"
	"strings"
	"strconv"
	"gorm.io/gorm/clause"
	"net/http"
	"bytes"
	"io"
	"twvixstock/internal/config"
	"sort"
	"errors"
)

// 爬蟲 TAIEX 分析存入資料
// 格式:
// {"msg":"success","status":200,"data":[{"date":"2026-03-02","stock_id":"TAIEX","Trading_Volume":14965963786,"Trading_money":1014040137978,"open":35277.48,"max":35345.72,"min":34605.36,"close":35095.09,"spread":-319.4,"Trading_turnover":6762693}]}
func TaiexAnalysis(ioBytes []byte) (count int, err error){

	// 解析 JSON
	var data model.TaiexApi
	err = json.Unmarshal(ioBytes, &data)
	if err != nil {
		log.Printf("TAIEX 分析 JSON 解析失敗: %v", err)
		return
	}

	// 準備插入
	insert := []model.StockRecord{}
	for _, v := range(data.Data){
		//資料
		close := v["close"].(float64)  //34700.82
		dateStr := v["date"].(string)  //2026-02-09
		date, _ := time.Parse(time.RFC3339, dateStr+"T00:00:00+08:00")  //用 +8時區字串轉

		val := model.StockRecord{
			Type: "TAIEX",
			Date: date,
			Close: close,
		}
		insert = append(insert, val)
	}

	// 執行插入或更新 (Upsert)
	result := repository.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "type"}, {Name: "date"}},  //唯獨索引
		DoUpdates: clause.AssignmentColumns([]string{"close"}),  //衝突時更新
	}).Create(&insert)

	if result.Error != nil {
		err = result.Error
		return
	}

	// 回報
	count = len(insert)
	return
}

// 爬蟲 VIXTWN 分析存入資料
/* 格式:
交易日期	時間(時/分/秒/毫秒)	臺指選擇權波動率指數	收盤前1分鐘平均指數
--------	-------------------	--------------------	-------------------
20260302	13450000			30.38		30.38
20260303	13450000			32.42		32.39
*/
func VixtwnAnalysis(strArr []string) (count int, err error){

	//插入用參數
	insert := []model.StockRecord{}

	//檔案組迴圈
	for _, v := range strArr{
		//檢查檔案字串格式
		if(! strings.Contains(v, "--------") ){
			fmt.Println("VIXTWN 發現一個未準備完成的檔案")  //未必是錯誤，有可能月初還沒準備好檔案
			continue
		}

		//切割字串
		exp1 := strings.Split(v, "\n")  //拆行
		for i:=2; i<len(exp1); i++{  //第二行開始
			exp2 := strings.Split(exp1[i], "\t")  //拆tab
			if(len(exp2) == 7){  //正確格式
				//參數
				dateStr := exp2[0]  //date 20260302
				close, _ := strconv.ParseFloat(exp2[4], 64)  //close 30.38

				//時間參數準備
				dateStr = fmt.Sprintf("%s-%s-%s", dateStr[0:4],dateStr[4:6],dateStr[6:8]) //2026-03-02
				date, _ := time.Parse(time.RFC3339, dateStr+"T00:00:00+08:00")  //用 +8時區字串轉

				val := model.StockRecord{
					Type: "VIXTWN",
					Date: date,
					Close: close,
				}
				insert = append(insert, val)
			}
		}
	}

	// 執行插入或更新 (Upsert)
	result := repository.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "type"}, {Name: "date"}},  //唯獨索引
		DoUpdates: clause.AssignmentColumns([]string{"close"}),  //衝突時更新
	}).Create(&insert)

	if result.Error != nil {
		err = result.Error
		return
	}

	// 回報
	count = len(insert)
	return
}

// 爬蟲 TAIPE 分析存入資料
// 格式:
// {"ResultCode":0,"ResultMsg":"","Data":[{"STK_CD":"","codeSerialNo":114,"code":"Y00114","name":"台股本益比","List":[{"value":18.47,"dataDate":"2007-05-01T00:00:00","UTC_dataDate":1177948800000},{"value":20.26,"dataDate":"2007-06-01T00:00:00","UTC_dataDate":1180627200000},{"value":18.72,"dataDate":"2007-08-01T00:00:00","UTC_dataDate":1185897600000}]}],"ExecuteTime": "2026-03-09 12:31:02.657"}
func TaipeAnalysis(ioBytes []byte, sortInsert bool) (count int, err error){

	// 解析 JSON
	var data model.TaipeApi
	err = json.Unmarshal(ioBytes, &data)
	if err != nil {
		log.Printf("TAIPE 分析 JSON 解析失敗: %v", err)
		return
	}

	// 將 List 拉出來斷言(陣列空介面)
	list := data.Data[0]["List"].([]interface{})

	// 準備插入
	insert := []model.StockRecord{}
	for _, v := range(list) {
		//第二層斷言
		data_sub := v.(map[string]interface{})
		
		//資料
		close := data_sub["value"].(float64)
		dataDate := data_sub["dataDate"].(string)  //2025-09-01T00:00:00
		date, _ := time.Parse(time.RFC3339, dataDate+"+08:00")  //用 +8時區字串轉

		//簡短插入 (避免處理過多sql)
		if(sortInsert){
			if(date.Before( time.Now().AddDate(0, -6, 0) )){  //6個月以前的不要
				continue
			}
		}

		val := model.StockRecord{
			Type: "TAIPE",
			Date: date,
			Close: close,
		}
		insert = append(insert, val)
	}

	// 執行插入或更新 (Upsert)
	result := repository.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "type"}, {Name: "date"}},  //唯獨索引
		DoUpdates: clause.AssignmentColumns([]string{"close"}),  //衝突時更新
	}).Create(&insert)

	if result.Error != nil {
		err = result.Error
		return
	}

	// 回報
	count = len(insert)
	return
}

//問 Gemini 最近三個月的股市狀況
func GeminiApiFetch(configCrawler config.CrawlerConfig)(count int, err error){

	//=============================階段(1/5) 取得數據=============================

	//定義統計資料
	type statistics = struct{
		TAIEX   map[string]float64
		VIXTWN  map[string]float64
		TAIPE   map[string]float64
	}

	//最近六個月(每月一條收盤價)
	getMonths := int(6)
	var monthData = statistics{
		TAIEX:  make(map[string]float64),
		VIXTWN: make(map[string]float64),
		TAIPE:  make(map[string]float64),
	}

	//最近15天資料
	recentDays := int(15)
	recentDaysTime := time.Now().AddDate(0, 0, -(recentDays+1))
	var recentData = statistics{
		TAIEX:  make(map[string]float64),
		VIXTWN: make(map[string]float64),
		TAIPE:  make(map[string]float64),
	}

	//開始抓
	//從6個月前的 1 號開始抓
	ly := int(time.Now().Year())  //今年
	lm := int(time.Now().Month())  //今月
	lm -= getMonths
	if(lm <= 0){
		lm += 12
		ly -= 1
	}
	last6MonthtimeStr := fmt.Sprintf("%04d-%02d-%02d", ly, lm, 1)

	//DB
	stocks, err := GetStocksByDate(last6MonthtimeStr, "", "all")
	if err != nil {
		fmt.Println("CheckForNeedCrawler DB ERROR!!")
		return
	}

	//全紀錄依序找
	nowM := fmt.Sprintf("%04d-%02d", int(time.Now().Year()), int(time.Now().Month()))  //當月
	for i := len(stocks) - 1; i >= 0; i-- {
		v := stocks[i]
		//每個月抓倒序第一條資料
		if(nowM != v.Date[0:7]){  //當月不抓(跟最近15天衝突)
			if(v.Type == "TAIEX" && monthData.TAIEX[v.Date[0:7]]==0){ 
				monthData.TAIEX[v.Date[0:7]] = v.Close  //當月最後一天
			}
			if(v.Type == "VIXTWN" && monthData.VIXTWN[v.Date[0:7]]==0){ 
				monthData.VIXTWN[v.Date[0:7]] = v.Close  //當月最後一天
			}
			if(v.Type == "TAIPE" && monthData.TAIPE[v.Date[0:7]]==0){ 
				monthData.TAIPE[v.Date[0:7]] = v.Close  //當月最後一天
			}
		}

		//最近的 15 天資料全部記錄下來
		if(v.Type == "TAIEX" || v.Type == "VIXTWN"){
			vTime, _ := time.Parse("2006-01-02", v.Date)
			if( vTime.After( recentDaysTime ) ){  //最近15天
				if(v.Type == "TAIEX"){
					recentData.TAIEX[v.Date] = v.Close
				}
				if(v.Type == "VIXTWN"){
					recentData.VIXTWN[v.Date] = v.Close
				}
			}
		}
	}

	//=============================階段(2/5) 模擬PE換算=============================

	// 模擬推演 每天的PE
	var stocks_sPE = make(map[string]model.GetStocksRet)  //模擬PE

	//換算用參數
	var savetmp_EX = make(map[string]model.GetStocksRet)  //暫存給順序換算
	var basePE float64  //基準PE
	var baseEX float64  //基準EX

	//  倒序找到最後一天 TAIEX 收盤
	for i := len(stocks) - 1; i >= 0; i-- {
		v := stocks[i]
		/*
			0. 假設現在是 3/12號，而PE有資料為 2/1號 (其實他是2月收盤價，應該是2/26才對)
			1. 倒序找基準，會發現 2/26時候的PE 以及 基準EX，記錄下
			2. 倒序換算，當日模擬PE 直接用基準PE 以及 基準EX 按比例縮放即可
			3. 順序換算，即完成
		*/

		//指判斷加權就可以換算
		if(v.Type == "TAIEX"){ 
			// 1. 找基準先
			if(basePE == 0 && monthData.TAIPE[v.Date[0:7]] != 0){
				basePE = monthData.TAIPE[v.Date[0:7]]
				baseEX = v.Close
			}else{
				//沒找到基準，儲存等等順序換算
				savetmp_EX[v.Date] = model.GetStocksRet{
					Type:  "TAIEX",
					Date:  v.Date,
					Close: v.Close,
				}
			}

			// 2. 倒序換算
			if(basePE != 0){
				//找到基準，倒序推演
				stocks_sPE[v.Date] = model.GetStocksRet{
					Type:  "TAIPE",
					Date:  v.Date,
					Close: basePE * (v.Close / baseEX),  //等比例換算
				}
			}
		}

		//最多紀錄 15 筆就好 (因為後面就拿15天最多)
		if ((len(stocks_sPE)+len(stocks_sPE)) >= recentDays  &&  basePE != 0){
			break
		}
	}
	//3. 順序換算
	for _, v := range(savetmp_EX){
		stocks_sPE[v.Date] = model.GetStocksRet{
			Type:  "TAIPE",
			Date:  v.Date,
			Close: basePE * (v.Close / baseEX),  //等比例換算
		}
	}
	
	//=============================階段(3/5) 串成字串=============================

	// 先串 6 個月
	var TAIEX_monthStr strings.Builder
	var VIXTWN_monthStr strings.Builder
	var TAIPE_monthStr strings.Builder
	//for k,v := range(monthData.TAIEX){
	//	TAIEX_monthStr.WriteString( fmt.Sprintf("%s=%.2f ", k, v) )
	//}
	for _,k := range(GetMapSortKeys(monthData.TAIEX)){
		TAIEX_monthStr.WriteString( fmt.Sprintf("%s=%.2f ", k, monthData.TAIEX[k] ) )
	}
	for _,k := range(GetMapSortKeys(monthData.VIXTWN)){
		VIXTWN_monthStr.WriteString( fmt.Sprintf("%s=%.2f ", k, monthData.VIXTWN[k] ) )
	}
	for _,k := range(GetMapSortKeys(monthData.TAIPE)){
		TAIPE_monthStr.WriteString( fmt.Sprintf("%s=%.2f ", k, monthData.TAIPE[k] ) )
	}

	// 再串 15 天
	var TAIEX_recentStr strings.Builder
	var VIXTWN_recentStr strings.Builder
	var TAIPE_recentStr strings.Builder
	for _,k := range(GetMapSortKeys(recentData.TAIEX)){
		TAIEX_recentStr.WriteString( fmt.Sprintf("%s=%.2f ", k, recentData.TAIEX[k] ) )
	}
	for _,k := range(GetMapSortKeys(recentData.VIXTWN)){
		VIXTWN_recentStr.WriteString( fmt.Sprintf("%s=%.2f ", k, recentData.VIXTWN[k] ) )
	}
	for _,k := range(GetMapSortKeys(stocks_sPE)){
		t, _ := time.Parse("2006-01-02", k)
		if( t.After( recentDaysTime ) ){  //抓15天內的資料
			TAIPE_recentStr.WriteString( fmt.Sprintf("%s=%.2f ", k, stocks_sPE[k].Close ) )
		}
	}

	// 準備
	var question strings.Builder
	question.WriteString("台股近六個月每月收盤價: \n")
	question.WriteString(" 台灣加權指數 TAIEX: " + TAIEX_monthStr.String() + "\n")
	question.WriteString(" 台灣恐慌指數 VIXTWN: " + VIXTWN_monthStr.String() + "\n")
	question.WriteString(" 台灣加權本益比 TAIPE: " + TAIPE_monthStr.String() + "\n")
	question.WriteString("台股近15天收盤價: \n")
	question.WriteString(" 台灣加權指數 TAIEX: " + TAIEX_recentStr.String() + "\n")
	question.WriteString(" 台灣恐慌指數 VIXTWN: " + VIXTWN_recentStr.String() + "\n")
	question.WriteString(" 台灣加權本益比(估值) TAIPE: " + TAIPE_recentStr.String() + "\n")
	question.WriteString(`
請評估台股現況及近兩周國際局勢:
 point: 0分為必離場、100分為必進場，請給分數
 price: 評估台灣加權股價要多少合理
 reason: 請給出約300字理由
回應格式為json: {'point': int, 'price': int, 'reason': string}
不需要追問我任何問題，感謝你。
	`)
	//fmt.Println(question.String())
	

	//=============================階段(4/5) 發送Gemini API=============================

	//模組  GeminiModel  |  GeminiModellite
	useModel := configCrawler.GeminiModel
	postData := []byte(`{'contents':[{'parts':[{'text':"`+question.String()+`"}]}]}`)

	//發送請求
	retry := 5
	var apiBody []byte
	for i := range(retry){
		//準備
		req, _ := http.NewRequest(
			"POST",
			fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent", useModel),
			bytes.NewBuffer( postData ), 
		)
		req.Header.Set("Content-Type", "application/json")  //header Json
		req.Header.Set("x-goog-api-key", configCrawler.Token)  //key

		// 請求
		fmt.Printf(" API post(%d/%d):\n",(i+1) ,retry)
		res, err := http.DefaultClient.Do(req)

		// 請求錯誤檢查
		if( err != nil  ||  res.StatusCode != 200 ) { 
			fmt.Printf(" - findError(%d): %v\n", res.StatusCode, err)

//apiBody, _ = io.ReadAll(res.Body)
//fmt.Println("errorBody:", string(apiBody))

			// 關閉連線
			res.Body.Close() 

			// 嘗試次數檢查
			if(i == 4 ){
				err = fmt.Errorf("%w (httpCode: %d)", err, res.StatusCode)  //紀錄 err 與 httpCode
				return 0, err
			}else{
				time.Sleep( time.Second * 30)  //等半分鐘
				continue  //還可以嘗試
			}
		}

		// 紀錄
		apiBody, err = io.ReadAll(res.Body)
		res.Body.Close() 
		if( err != nil ){  //格式怪怪的
			return 0, errors.New("api Body can't io.ReadAll")
		}else{
			break
		}
	}

	//=============================階段(5/5) 插入DB=============================

	// 解析 JSON
	var data model.GeminiApi
	err = json.Unmarshal(apiBody, &data)
	if err != nil {
		log.Printf("GEMINI 分析 JSON 解析失敗: %v", err)
		return
	}

	// 拆解結果
	rawText := string(data.Candidates[0].Content.Parts[0]["text"])
	cleanJSON := strings.TrimPrefix(rawText, "```json\n")
	cleanJSON = strings.TrimPrefix(cleanJSON, "```")
	cleanJSON = strings.TrimPrefix(cleanJSON, "json")
	cleanJSON = strings.TrimSuffix(cleanJSON, "\n```")
	cleanJSON = strings.TrimSuffix(cleanJSON, "```")
	cleanJSON = strings.TrimSpace(cleanJSON)
	
	// 取得最後一天 EX 日期
	allExKey := GetMapSortKeys(recentData.TAIEX)
	timeStr := allExKey[len(allExKey)-1] // 2026-03-11
	date, _ := time.Parse(time.RFC3339, timeStr+"T00:00:00+08:00")  //用 +8時區字串轉

	// 插入資料
	insert := []model.GeminiText{}
	insert = append(insert, model.GeminiText{
		Text: cleanJSON,
		Date: date,
		Remark: useModel + "(6m+15d)",
	})

	// 執行插入或更新 (Upsert)
	result := repository.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "date"}},  //唯獨索引
		DoUpdates: clause.AssignmentColumns([]string{"text"}),  //衝突時更新
	}).Create(&insert)

	if result.Error != nil {
		err = result.Error
		return
	}

	// 回報
	count = len(insert)
	return
}

//取得排序後的 map[string]type 鍵，返回 []string
// 使用 [K comparable, V any] 讓函數支援任何類型的 map
func GetMapSortKeys[V any](m map[string]V) []string {
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    return keys
}