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
	flag_k := flag.String("k", "all", "指數種類或gemini三月分析?(all|taiex|vixtwn|taipe|gemini)")
	flag_all := flag.Bool("all", false, "是否抓取全紀錄?")
	flag.Parse()
	crawlerKind := model.KindName( strings.ToLower(*flag_k) )

	//檢查要不要抓
	needCrawler := crawler.CheckForNeedCrawler()

	// 台灣加權指數
	if(needCrawler.TAIEX || *flag_all){ //DB沒資料 或 all時抓
		if(crawlerKind == model.CrawlerKind.All || crawlerKind == model.CrawlerKind.TAIEX){
			crawler.FetchDataTAIEX(*flag_all, config.AppConfig.Crawler)
		}
	}
	// 台灣恐慌指數
	if(needCrawler.VIXTWN || *flag_all){ //DB沒資料 或 all時抓
		if(crawlerKind == model.CrawlerKind.All || crawlerKind == model.CrawlerKind.VIXTWN){
			crawler.FetchDataVIXTWN(*flag_all, config.AppConfig.Crawler)
		}
	}
	// 台灣PE指數
	if(needCrawler.TAIPE || *flag_all){ //DB沒資料 或 all時抓
		if(crawlerKind == model.CrawlerKind.All || crawlerKind == model.CrawlerKind.TAIPE){
			crawler.FetchDataTAIPE(*flag_all, config.AppConfig.Crawler)
		}
	}
	//檢查要不要抓 again (如果前面抓到了)
	needCrawler = crawler.CheckForNeedCrawler()
	// gemini三月分析
	if(needCrawler.GEMINI || *flag_all){ //DB沒資料 或 all時抓
		if(crawlerKind == model.CrawlerKind.All || crawlerKind == model.CrawlerKind.GEMINI){
			crawler.FetchGeminiApi(*flag_all, config.AppConfig.Crawler)
		}
	}
}

