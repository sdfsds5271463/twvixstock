package main

import(
	"os"
	"testing"
	"bytes"
	"strings"
	"net/http"
	"net/http/httptest"
	"io"

	"twvixstock/internal/config"
	"twvixstock/internal/repository"

	"log"
	//"fmt"
)

//測試伺服器初始
func TestServerInit(t *testing.T){

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
}

//測試API stockDB
func TestServerApiStockDB(t *testing.T){
	//檢查路由 /api/v1/stockDB
	//模擬發送
	r := SetupRouter()
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/stockDB?start=2026-01-01&type=taipe",
		nil,
	)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	//檢查狀態
	if w.Code != http.StatusOK {
		t.Errorf("expected 200 but got %d", w.Code)
	}
	body, _ := io.ReadAll(w.Body)
	//檢查內容
    if !strings.Contains(string(body), `"code":1000`) {
        t.Errorf("路由 /api/v1/stockDB 沒有出現預期字串 code 1000")
    }
}

//測試API geminiData
func TestServerApiGeminiData(t *testing.T){
	//檢查路由 /api/v1/geminiData
	//模擬發送
	r := SetupRouter()
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/geminiData",
		strings.NewReader(`{"date":""}`),
	)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	//檢查狀態
	if w.Code != http.StatusOK {
		t.Errorf("expected 200 but got %d", w.Code)
	}
	body, _ := io.ReadAll(w.Body)
	//檢查內容
    if !strings.Contains(string(body), `"code":1000`) {
        t.Errorf("路由 /api/v1/stockDB 沒有出現預期字串 code 1000")
    }
}

//測試靜態檔
func TestServerApiStatic(t *testing.T){
	//檢查路由 /sitemap.xml
	//模擬發送
	r := SetupRouter()
	req := httptest.NewRequest(
		http.MethodGet,
		"/sitemap.xml",
		nil,
	)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	//檢查狀態
	if w.Code != http.StatusOK {
		t.Errorf("expected 200 but got %d", w.Code)
	}
}

//測試SPA Fallback (隨機路徑都要返回 200 code)
func TestSpaFallback(t *testing.T){
	//檢查路由 /sitemap.xml
	//模擬發送
	r := SetupRouter()
	req := httptest.NewRequest(
		http.MethodGet,
		"/test/888/666",
		nil,
	)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	//檢查狀態
	if w.Code != http.StatusOK {
		t.Errorf("expected 200 but got %d", w.Code)
	}
}