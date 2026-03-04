package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Crawler  CrawlerConfig
}

type ServerConfig struct {
	Port int
	Mode string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type CrawlerConfig struct {
	BaseURL      string
	IntervalHour int
	Token        string
}

var AppConfig *Config

func LoadConfig(path string) {
	if path == "" { path = "./config" }

	// 1. 先讀取 .env (如果檔案存在)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from system environment")
	}

	// 2. 設定 Viper 讀取 YAML
	viper.SetConfigName("config") // 檔案名稱 (不含副檔名)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path) // 尋找路徑

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 3. 將資料對應到 Struct
	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// 4. 從環境變數覆蓋讀取
	AppConfig.Database.Host = os.Getenv("DB_HOST")
	AppConfig.Database.User = os.Getenv("DB_USER")
	AppConfig.Database.Password = os.Getenv("DB_PASSWORD")
	AppConfig.Database.DBName = os.Getenv("DB_NAME")
	AppConfig.Crawler.Token = os.Getenv("CRAWLER_TOKEN")

	log.Println("Config loaded successfully")
}