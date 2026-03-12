package crawler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"twvixstock/internal/service"
	"twvixstock/internal/model"
	"twvixstock/internal/config"
)

// 檢查是不是要爬蟲
func CheckForNeedCrawler()(needCrawler model.NeedCrawler){

	//旗標(預設全抓)
	needCrawler = model.NeedCrawler{
		TAIEX:   true,  //預設抓，當DB有時不抓
		VIXTWN:  true,  //預設抓，當DB有時不抓
		TAIPE:   true,  //預設抓，當DB有時不抓
		GEMINI:  false, //預設不抓，當 TAIEX VIXTWN 有時才抓
	}

	//檢查 TAIEX VIXTWN
	timeStr := time.Now().Format("2006-01-02") //今日日期
	stocks, err := service.GetStocksByDate(timeStr, timeStr, "all")
	if err != nil {
		fmt.Println("CheckForNeedCrawler DB ERROR!!")
		return
	}

	for _ ,v := range(stocks){
		if(v.Type == "TAIEX"){ //TAIEX 有紀錄，不用抓了
			needCrawler.TAIEX = false
			fmt.Println("TAIEX check already has datas")
		}
		if(v.Type == "VIXTWN"){ //VIXTWN 有紀錄，不用抓了
			needCrawler.VIXTWN = false
			fmt.Println("VIXTWN check already has datas")
		}
	}

	//檢查 TAIPE
	//上個月日期
	ly := int(time.Now().Year())  //今年
	lm := int(time.Now().Month())  //今月
	lm -= 1  //上個月計算
	if(lm <= 0){
		lm += 12
		ly -= 1
	}
	lastMonthtimeStr := fmt.Sprintf("%04d-%02d-%02d", ly, lm, 1)
	stocks2, err := service.GetStocksByDate(lastMonthtimeStr, lastMonthtimeStr, "all")
	if err != nil {
		fmt.Println("CheckForNeedCrawler DB ERROR!!")
		return
	}
	if(len(stocks2) == 1){ //TAIPE 有紀錄，不用抓了
		needCrawler.TAIPE = false
		fmt.Println("TAIPE check already has datas")
	}

	//檢查 GEMINI (當 TAIEX VIXTWN 都有時才抓)
	if (!needCrawler.TAIEX && !needCrawler.VIXTWN) {
		needCrawler.GEMINI = true
	}else{
		fmt.Println("GEMINI skip because TAIEX|VIXTWN not ready")
	}
	if ( needCrawler.GEMINI == true ){  //如果要抓
		geminiChk, err := service.GetGeminiTextByDate(timeStr)  //檢查有沒有資料囉
		if err != nil {
			fmt.Println("CheckForNeedGemini DB ERROR!!")
			return
		}
		if (geminiChk.Reason != ""){  //有資料
			needCrawler.GEMINI = false  //別抓
			fmt.Println("GEMINI check already has datas")
		}
	}
	return
}


// 台灣加權爬蟲
func FetchDataTAIEX(all bool, configCrawler config.CrawlerConfig) {

	// 全資料all時，抓取366天；否則抓取31天
	var getDays int64 = 31
	if(all){
		getDays = 366
	}

	// 時間準備
	timestamp := time.Now().Unix() - 86400*getDays  //抓取天數
	taipei_zone := time.FixedZone("UTC+8", 8 * 3600)
	t := time.Unix( timestamp, 0 ).In(taipei_zone)  //台北時區
	start_date := t.Format("2006-01-02")

	// 爬蟲
	url := "https://api.finmindtrade.com/api/v4/data?dataset=TaiwanStockPrice&data_id=TAIEX&start_date=" + start_date
	res, err := http.Get(url)
	if(err != nil){ 
		log.Printf("TAIEX 爬蟲請求失敗: %v", err)
		return
	}
	defer res.Body.Close()

	if(res.StatusCode != 200){ // 狀態
		log.Printf("TAIEX 爬蟲狀態失敗: %v", res.StatusCode)
		return
	}

	ioBytes, err := io.ReadAll(res.Body) // 讀取
	if(err != nil){ 
		log.Printf("TAIEX 爬蟲 io 讀取失敗: %v", err)
		return
	}

	// 交給服務進行資料解析
	count, err := service.TaiexAnalysis(ioBytes)
	if(err != nil){ 
		log.Printf("TAIEX 資料解析失敗: %v", err)
		return
	}

	fmt.Println("TAIEX 解析筆數:", count)
}

// 台灣恐慌爬蟲
func FetchDataVIXTWN(all bool, configCrawler config.CrawlerConfig) {

	// 全資料all時，抓取4個月檔案；否則抓取1個月檔案
	var getMonth = 1
	if(all){
		getMonth = 4
	}

	// 時間準備
	taipei_zone := time.FixedZone("UTC+8", 8 * 3600)
	t := time.Now().In(taipei_zone)  //台北時區
	t = t.Add( time.Duration( time.Hour * -14 ) )  //扣14小時 (因為下午兩點以後才可能封盤更新)

	// 爬蟲抓取
	var strArr []string
	for i:=0; i<getMonth; i++ {
		// 時間檔字串
		year := t.Year()
		mon := int(t.Month()) - i //每次往回一個月
		if(mon <= 0){ year -= 1; mon += 12 }
		fileStr := fmt.Sprintf("%04d%02dnew.txt", year, mon)

		// 爬蟲
		url := "https://www.taifex.com.tw/file/taifex/Dailydownload/vix/log2data/" + fileStr
		res, err := http.Get(url)
		if(err != nil){ 
			log.Printf("VIXTWN 爬蟲請求失敗: %v", err)
			continue
		}
		defer res.Body.Close()

		if(res.StatusCode != 200){ // 狀態
			log.Printf("VIXTWN 爬蟲狀態失敗: %v", res.StatusCode)
			return
		}

		ioBytes, err := io.ReadAll(res.Body) // 讀取
		if(err != nil){ 
			log.Printf("VIXTWN 爬蟲 io 讀取失敗: %v", err)
			continue
		}

		// 存入
		strArr = append(strArr, string(ioBytes))
		time.Sleep(time.Second * 3) //睡3秒
	}

	// 交給服務進行資料解析
	count, err := service.VixtwnAnalysis(strArr)
	if(err != nil){ 
		log.Printf("VIXTWN 資料解析失敗: %v", err)
		return
	}

	fmt.Println("VIXTWN 解析筆數:", count)
}

// 台灣PE爬蟲
func FetchDataTAIPE(all bool, configCrawler config.CrawlerConfig) {

	// 全資料
	sortInsert := !all //短插入
	
	//目標網站(元大)沒有時間過濾，永遠抓全部
	// sortInsert 用來決定是否短插入，就不用處理那麼多DB
	//https://www.yuanta-etfadvisor.com/map/chart?area=80008d7a-64b7-450c-9b96-49d7bf712b34&staticCategoryName=MacroEconomicIndex&product1=507f7d79-b6e2-44b1-902d-f304253322fc
	//https://api.yuantafunds.com/ECTranslationAI/api/bridge?APIType=ConsultingAPI&FuncId=Index%2FMarcoEconomicIndex&code=Y00114&CompanyName=YUANTAFUNDS&Platform=ETFXAI&AppName=ETFxAI

	// 時間準備

	// 爬蟲
	url := "https://api.yuantafunds.com/ECTranslationAI/api/bridge?APIType=ConsultingAPI&FuncId=Index%2FMarcoEconomicIndex&code=Y00114&CompanyName=YUANTAFUNDS&Platform=ETFXAI&AppName=ETFxAI"
	res, err := http.Get(url)
	if(err != nil){ 
		log.Printf("TAIPE 爬蟲請求失敗: %v", err)
		return
	}
	defer res.Body.Close()

	if(res.StatusCode != 200){ // 狀態
		log.Printf("TAIPE 爬蟲狀態失敗: %v", res.StatusCode)
		return
	}

	ioBytes, err := io.ReadAll(res.Body) // 讀取
	if(err != nil){ 
		log.Printf("TAIPE 爬蟲 io 讀取失敗: %v", err)
		return
	}

	// 交給服務進行資料解析
	count, err := service.TaipeAnalysis(ioBytes, sortInsert)
	if(err != nil){ 
		log.Printf("TAIPE 資料解析失敗: %v", err)
		return
	}

	fmt.Println("TAIPE 解析筆數:", count)
}

// GEMINI API
func FetchGeminiApi(all bool, configCrawler config.CrawlerConfig) {
	_ = all //不用這個參數，固定給 6 個月 + 15天 分析

	// 交給服務進行資料解析
	count, err := service.GeminiApiFetch(configCrawler)
	fmt.Println("GEMINI:", count, err)
}