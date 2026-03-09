package main

import (
	"fmt"
	"twvixstock/internal/config"
	"twvixstock/internal/handler"
	"twvixstock/internal/repository"

	"github.com/didip/tollbooth/v7"
    "github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化
	config.LoadConfig("")
	repository.InitDB()

	// 設定 Gin
	r := gin.Default()

	// 設定 tollbooth 限流 ( 每2分鐘最多請求20次 => 1/6秒充能速率+20充能上限 )
	limiter := tollbooth.NewLimiter((10.0/60.0), nil)  //每秒恢復請求充能 (10.0/60.0) (每六秒才能充一次)
	limiter.SetBurst(20)  //請求充能上限 (每次請求扣1，沒充能不給訪問)
    limiter.SetMessage("{'code':3001, 'error':'too many requests'}")
	limiter.SetMessageContentType("application/json")
	limiter.SetStatusCode(429)

	// 允許跨域 (CORS) - 為了讓你的 Vue 前端能順利存取
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// 設定路由群組
	api := r.Group("/api")
	{
		api.Use(tollbooth_gin.LimitHandler(limiter))  // 進行 tollbooth_gin 限流

		api.GET("/stockDB", handler.StockDataShow)
		api.POST("/stockDB", handler.StockDataShow)

		/*
			回應 {'code':<code>, 'error':<msg>, 'data':<data>}
			error只在錯誤時回應， <msg> = 錯誤細節
			data 只在正確時回應， <data> = 資料
			code 必填，<code> =
				1000 正常        (200)
				2001 請求參數有誤 (400)
				2002 資料傳輸異常 (500)
				3001 請求過於頻繁 (429)
		*/
	}

	// 根目錄靜態檔 (讓 vue 掛載用)
	r.Static("/assets", "./static/assets")
	r.Static("/images", "./static/images")

	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// 啟動
	port := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	fmt.Printf("Server 啟動於 %s\n", port)
	r.Run(port)
}