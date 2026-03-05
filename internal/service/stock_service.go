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
)

// 根據日期過濾資料
func GetStocksByDate(startDate string, endDate string) ([]model.StockRecord, error) {
	var stocks []model.StockRecord

	// 呼叫 Repository 層的 DB 物件進行查詢
	db := repository.DB

	if startDate != "" && endDate != "" {
		db = db.Where("date BETWEEN ? AND ?", startDate, endDate)
	}

	err := db.Find(&stocks).Error
	return stocks, err
}

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

fmt.Println( dateStr, date.Format("2006/01/02 15:04:05 -07"), date.Unix() )

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

fmt.Println( dateStr, date.Format("2006/01/02 15:04:05 -07"), date.Unix() )

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
