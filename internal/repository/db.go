package repository

import (
	"fmt"
	"log"
	"twvixstock/internal/config"
	//"twvixstock/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化資料庫連線
func InitDB() {
	// 執行 InitDB() 之前必須先有 config.LoadConfig("")
	c := config.AppConfig.Database
	
	// 組裝 DSN (Data Source Name)
	// 格式: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FTaipei",  //時區 +8
		c.User, c.Password, c.Host, c.Port, c.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 測試環境開啟 Log，可以看到 SQL 語法
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自動遷移 (Auto Migration)
	// 這會自動幫你在 MySQL 建立表，如果表已存在則會檢查欄位是否需要更新
	// err = db.AutoMigrate(&model.StockRecord{})
	// if err != nil {
	// 	log.Fatalf("Failed to migrate database: %v", err)
	// }

	DB = db
	log.Println("Database connection established")
}