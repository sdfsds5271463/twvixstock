package main

import (
	"fmt"
	"twvixstock/internal/config"
	"twvixstock/internal/crawler"
	"twvixstock/internal/repository"
)

func main() {
	fmt.Println("排程爬蟲啟動中...")

	// 1. 初始化
	config.LoadConfig("")
	repository.InitDB()

	// 2. 模擬爬蟲
	crawler.FetchData(config.AppConfig.Crawler.Token)
}

