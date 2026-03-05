package crawler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"twvixstock/internal/service"
)

// 台灣加權爬蟲
func FetchDataTAIEX(all bool, token string) {

	// 全資料all時，抓取366天；否則抓取31天
	var getDays int64 = 31
	if(all){
		getDays = 366
	}

	// 時間準備
	timestamp := time.Now().Unix() - 86400*getDays  //抓取天數
	taipei_zone, _ := time.LoadLocation("Asia/Taipei")
	t := time.Unix( timestamp, 0 ).In(taipei_zone)  //台北時區
	start_date := t.Format("2006-01-02")

	// 爬蟲
	url := "https://api.finmindtrade.com/api/v4/data?dataset=TaiwanStockPrice&data_id=TAIEX&start_date=" + start_date
	res, err := http.Get(url)
	if(err != nil){ 
		log.Printf("TAIEX 爬蟲請求失敗: %v", err)
		return
	}
	defer res.Body.Close()

	if(res.StatusCode != 200){ // 狀態
		log.Printf("TAIEX 爬蟲狀態失敗: %v", res.StatusCode)
		return
	}

	ioBytes, err := io.ReadAll(res.Body) // 讀取
	if(err != nil){ 
		log.Printf("TAIEX 爬蟲 io 讀取失敗: %v", err)
		return
	}

	// 交給服務進行資料解析
	count, err := service.TaiexAnalysis(ioBytes)
	if(err != nil){ 
		log.Printf("TAIEX 資料解析失敗: %v", err)
		return
	}

	fmt.Println("TAIEX 解析筆數:", count)
}

// 台灣恐慌爬蟲
func FetchDataVIXTWN(all bool, token string) {

	// 全資料all時，抓取4個月檔案；否則抓取1個月檔案
	var getMonth = 1
	if(all){
		getMonth = 4
	}

	// 時間準備
	taipei_zone, _ := time.LoadLocation("Asia/Taipei")
	t := time.Now().In(taipei_zone)  //台北時區
	t = t.Add( time.Duration( time.Hour * -14 ) )  //扣14小時 (因為下午兩點以後才可能封盤更新)

	// 爬蟲抓取
	var strArr []string
	for i:=0; i<getMonth; i++ {
		// 時間檔字串
		year := t.Year()
		mon := int(t.Month()) - i //每次往回一個月
		if(mon <= 0){ year -= 1; mon += 12 }
		fileStr := fmt.Sprintf("%04d%02dnew.txt", year, mon)

		// 爬蟲
		url := "https://www.taifex.com.tw/file/taifex/Dailydownload/vix/log2data/" + fileStr
		res, err := http.Get(url)
		if(err != nil){ 
			log.Printf("VIXTWN 爬蟲請求失敗: %v", err)
			continue
		}
		defer res.Body.Close()

		if(res.StatusCode != 200){ // 狀態
			log.Printf("VIXTWN 爬蟲狀態失敗: %v", res.StatusCode)
			return
		}

		ioBytes, err := io.ReadAll(res.Body) // 讀取
		if(err != nil){ 
			log.Printf("VIXTWN 爬蟲 io 讀取失敗: %v", err)
			continue
		}

		// 存入
		strArr = append(strArr, string(ioBytes))
		time.Sleep(time.Second * 3) //睡3秒
	}

	// 交給服務進行資料解析
	count, err := service.VixtwnAnalysis(strArr)
	if(err != nil){ 
		log.Printf("VIXTWN 資料解析失敗: %v", err)
		return
	}

	fmt.Println("VIXTWN 解析筆數:", count)
}