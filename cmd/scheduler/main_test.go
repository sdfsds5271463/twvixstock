package main

import (
	"testing"
	"os"
	"bytes"
	"strings"

	"twvixstock/internal/crawler"
	"twvixstock/internal/config"
	"twvixstock/internal/repository"
	"twvixstock/internal/service"

	"log"
)

//測試排程初始化
func TestSchedulerInit(t *testing.T){

	//建立輸出log接收
	var buf bytes.Buffer
	originalOutput := log.Writer()
    log.SetOutput(&buf)  
		//後續檢查 buf.String() 字串即可
		//log.SetOutput(os.Stderr)   //復原
    defer log.SetOutput(originalOutput)

	//初始化 config test
	os.Chdir("../..")
	config.LoadConfig("")
	if config.AppConfig.Database.Host == "" {
		t.Error("未能讀取到資料庫 Host 設定")
	} else {
		t.Logf("讀取成功，Host 為: %s", config.AppConfig.Database.Host)
	}

	//初始化 DB test
	repository.InitDB()
    if !strings.Contains(buf.String(), "Database connection established") {
        t.Errorf("初始化 DB test 沒出現預期字串 connection established")
    }
	buf.Reset()
	if repository.DB == nil {
		t.Error("執行完 InitDB 後，DB 實例不應為 nil")
	}

	//檢查 CheckForNeedCrawler
	_ = crawler.CheckForNeedCrawler()
    if strings.Contains(buf.String(), "ERROR") {
        t.Errorf("CheckForNeedCrawler 出現 log ERROR")
    }
	buf.Reset()
//log.SetOutput(os.Stderr)
}

//測試排程API
func TestSchedulerApi(t *testing.T){
	
	//建立輸出log接收
	var buf bytes.Buffer
	originalOutput := log.Writer()
    log.SetOutput(&buf)  
    defer log.SetOutput(originalOutput)

	//檢查 FetchDataTAIEX
	crawler.FetchDataTAIEX(false, config.AppConfig.Crawler)
    if strings.Contains(buf.String(), "失敗") {
        t.Errorf("FetchDataTAIEX 出現 log 失敗")
    }
	buf.Reset()

	//檢查 FetchDataVIXTWN
	crawler.FetchDataVIXTWN(false, config.AppConfig.Crawler)
    if strings.Contains(buf.String(), "失敗") {
        t.Errorf("FetchDataVIXTWN 出現 log 失敗")
    }
	buf.Reset()

	//檢查 FetchDataTAIPE
	crawler.FetchDataTAIPE(false, config.AppConfig.Crawler)
    if strings.Contains(buf.String(), "失敗") {
        t.Errorf("FetchDataTAIPE 出現 log 失敗")
    }
	buf.Reset()

	//檢查 FetchGeminiApi
	_, err := service.GeminiApiFetch(config.AppConfig.Crawler)
    if err!=nil {
        t.Errorf("FetchGeminiApi 出現 err: %v", err)
    }
}