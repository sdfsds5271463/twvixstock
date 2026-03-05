package model

import (
	"time"
)

// StockRecord 對應到資料表
type StockRecord struct {
	ID    uint      `gorm:"primaryKey;autoIncrement"`
	Type  string    `gorm:"type:enum('TAIEX','VIXTWN');index"`
	Date  time.Time `gorm:"type:date;index"`
	Close float64   `gorm:"type:decimal(10,2);"`
}

// TableName 指定資料表名稱 (若不指定，GORM 預設 stock_records )
func (StockRecord) TableName() string {
	return "stock"
}