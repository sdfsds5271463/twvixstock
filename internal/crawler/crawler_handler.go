package crawler

import (
	"fmt"
	"log"
	"net/http"
	"twvixstock/internal/service"
	"io"
)

func FetchData(token string) {
	// 實作爬蟲
	res, err := http.Get("http://allenweb.ddns.net:9527/api1/res.php?token="+token)
	if(err != nil){ 
		log.Printf("爬蟲失敗: %v", err)
	}
	defer res.Body.Close()
	fmt.Println("爬蟲測試StatusCode:", res.StatusCode)  //200

	body, _ := io.ReadAll(res.Body)

	count, err := service.DataAnalysis(body)
	if(err != nil){ 
		log.Printf("資料解析失敗: %v", err)
	}
	fmt.Println("成功解析筆數:", count)  //200
}