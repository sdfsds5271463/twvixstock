package model

import (
	"time"
)

// GeminiText 對應到資料表
type GeminiText struct {
	ID     uint       `gorm:"primaryKey;autoIncrement" json:"-"`
	Text   string     `gorm:"type:text"`
	Date   time.Time  `gorm:"type:date;index"`
	Remark string     `gorm:"type:varchar(50)"`
}

// TableName 指定資料表名稱 (若不指定，GORM 預設 gemini_text_records )
func (GeminiText) TableName() string {
	return "gemini_text"
}


