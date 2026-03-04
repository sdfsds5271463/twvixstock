package handler

import (
	"net/http"
	"twvixstock/internal/service"

	"github.com/gin-gonic/gin"
)

// StockDataShow 處理股票顯示請求
func StockDataShow(c *gin.Context) {
	// 1. 取得 Query Parameters (例如: /api/stockDB?start=2023-01-01&end=2023-01-31)
	startDate := c.Query("start")
	endDate := c.Query("end")

	// 2. 呼叫 Service 取得資料
	stocks, err := service.GetStocksByDate(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法取得資料: " + err.Error()})
		return
	}

	// 3. 回傳 JSON 給前端
	c.JSON(http.StatusOK, stocks)
}