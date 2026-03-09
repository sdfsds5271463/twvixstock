package main

import (
	"fmt"
	"twvixstock/internal/config"
	"twvixstock/internal/crawler"
	"twvixstock/internal/repository"

	"twvixstock/internal/model"
	"flag"
	"strings"
)

func main() {
	fmt.Println("排程爬蟲啟動中...")

	// 初始化
	config.LoadConfig("")
	repository.InitDB()

	// 爬蟲 flag (EX: go run cmd/scheduler/main.go -k=taiex -all)
	flag_k := flag.String("k", "all", "指數種類?(all|taiex|vixtwn|taipe)")
	flag_all := flag.Bool("all", false, "是否抓取全紀錄?")
	flag.Parse()
	crawlerKind := model.KindName( strings.ToLower(*flag_k) )

	// 台灣加權指數
	if(crawlerKind == model.CrawlerKind.All || crawlerKind == model.CrawlerKind.TAIEX){
		crawler.FetchDataTAIEX(*flag_all, config.AppConfig.Crawler.Token)
	}
	// 台灣恐慌指數
	if(crawlerKind == model.CrawlerKind.All || crawlerKind == model.CrawlerKind.VIXTWN){
		crawler.FetchDataVIXTWN(*flag_all, config.AppConfig.Crawler.Token)
	}
	// 台灣PE指數
	if(crawlerKind == model.CrawlerKind.All || crawlerKind == model.CrawlerKind.TAIPE){
		crawler.FetchDataTAIPE(*flag_all, config.AppConfig.Crawler.Token)
	}

}

