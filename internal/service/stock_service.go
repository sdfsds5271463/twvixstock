package service

import (
	"errors"
	//"fmt"
	"strings"
	"twvixstock/internal/model"
	"twvixstock/internal/repository"
	"encoding/json"
)

// 根據日期過濾資料
func GetStocksByDate(startDate string, endDate string, stocktype string) ([]model.GetStocksRet, error) {
	var stocks []model.StockRecord

	//空白消除
	startDate = strings.TrimSpace(startDate)
	endDate = strings.TrimSpace(endDate)
	stocktype = strings.TrimSpace(stocktype)

	//類別準備
	stocktype = strings.ToUpper(stocktype)
	if(stocktype == "ALL"){  //全部的話也就不限制
		stocktype = ""
	}

	// 呼叫 Repository 層的 DB 物件進行查詢
	db := repository.DB

	//if startDate != "" && endDate != "" {
	//	db = db.Where("date BETWEEN ? AND ?", startDate, endDate)
	//}

	if startDate != ""{
		db = db.Where("date >= ?", startDate)
	}
	if endDate != ""{
		db = db.Where("date <= ?", endDate)
	}

	if stocktype != "" {
		db = db.Where("type = ? ", stocktype )
	}
	
	// 查詢
	err := db.Order("date ASC").Find(&stocks).Error
	if(err != nil){
		return nil, err
	}

	// 格式化日期
	var ret []model.GetStocksRet
	for _,v := range(stocks){
		ret = append(ret, model.GetStocksRet{
			Type: v.Type,
			Close: v.Close,
			Date: v.Date.Format("2006-01-02"),
		})
	}

	return ret, err
}

// 取得 Gemini Api 的分析
func GetGeminiTextByDate(date string)(data model.GeminiJson, err error){

	if(date == ""){
		err = errors.New("date can't null")
		return
	}

	var geminiText []model.GeminiText

	// 呼叫 Repository 層的 DB 物件進行查詢
	db := repository.DB
	db = db.Where("date = ?", date)

	// 查詢
	err = db.Find(&geminiText).Error
	if(err != nil){
		return
	}

	// 解析
	if len(geminiText) > 0 { //有資料
		err =  json.Unmarshal([]byte(geminiText[0].Text) , &data )
		if(err != nil){
			return
		}
		data.Date = date
	}

	return
}