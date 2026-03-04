package model

import (
	"time"
)

// StockRecord 對應你的資料表
type StockRecord struct {
	ID    uint      `gorm:"primaryKey;autoIncrement"`
	Type  string    `gorm:"type:enum('TAIEX','VIXTWN');"`
	Date  time.Time `gorm:"type:date;index"` // 建立索引提升查詢效率
	Close float64   `gorm:"type:decimal(10,2);"`
}

// TableName 指定資料表名稱 (若不指定，GORM 預設會變複數 stock_records)
func (StockRecord) TableName() string {
	return "stock"
}