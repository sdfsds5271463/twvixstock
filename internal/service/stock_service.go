package service

import (
	"twvixstock/internal/model"
	"twvixstock/internal/repository"
	"time"
)

// GetStocksByDate 根據日期過濾資料
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

// 爬蟲分析存入資料
func DataAnalysis(strdata []byte) (int, error){
	_ = strdata
	// 模擬假資料
	val := model.StockRecord{
		Type: "TAIEX",
		Date: time.Now(),
		Close: 123456,
	}
	data := []model.StockRecord{}
	data = append(data, val)
	data = append(data, val)

	//插入DB
	result := repository.DB.Create(&data)
	if result.Error != nil {
		return 0, result.Error
	}

	return len(data), nil
}