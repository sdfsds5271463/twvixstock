package main

import (
	"fmt"
	"twvixstock/internal/config"
	"twvixstock/internal/handler"
	"twvixstock/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 初始化
	config.LoadConfig("")
	repository.InitDB()

	// 2. 設定 Gin
	r := gin.Default()

	// 允許跨域 (CORS) - 為了讓你的 Vue 前端能順利存取
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// 根目錄靜態檔 (讓 vue 掛載用)
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/img", "./static/img")

	// 3. 設定路由群組
	api := r.Group("/api")
	{
		api.GET("/stockDB", handler.StockDataShow)
	}

	// 4. 啟動
	port := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	fmt.Printf("Server 啟動於0 %s\n", port)
	r.Run(port)
}