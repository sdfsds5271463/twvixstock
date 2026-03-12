package handler

import (
	"net/http"
	"twvixstock/internal/service"

	"github.com/gin-gonic/gin"
	"fmt"
)

// StockDataShow 處理股票顯示請求
func StockDataShow(c *gin.Context) {

	fmt.Print(" ( ↓↓↓ Visitor IP: ", c.ClientIP()," ↓↓↓ ) ")

	// 過濾 binding 條件
	type ReqStruct struct {
		Start string `json:"start" form:"start" binding:"required,datetime=2006-01-02"`
		End   string `json:"end"   form:"end"   binding:"omitempty,datetime=2006-01-02"`
		Type  string `json:"type"  form:"type"`
	}

	// 解析參數 (傳入的資料有問題)
	var req ReqStruct
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code":2001, "error": "Request is fail: " + err.Error()})
		return
	}

	// 呼叫 Service 取得資料 (搜DB出錯)
	stocks, err := service.GetStocksByDate(req.Start, req.End, req.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code":2002, "error": "Can't get data: " + err.Error()})
		return
	}

	// 回傳 JSON 給前端
	c.JSON(http.StatusOK, gin.H{"code":1000, "data":stocks})
}


// GeminiDataShow 處理GeminiData顯示請求
func GeminiDataShow(c *gin.Context) {

	fmt.Print(" ( ↓↓↓ Visitor IP: ", c.ClientIP()," ↓↓↓ ) ")

	// 過濾 binding 條件
	type ReqStruct struct {
		Date string `json:"date" form:"date"`
	}

	// 解析參數 (傳入的資料有問題)
	var req ReqStruct
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code":2001, "error": "Request is fail: " + err.Error()})
		return
	}

	// 呼叫 Service 取得資料 (搜DB出錯)
	geminiText, err := service.GetGeminiTextByDate(req.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code":2002, "error": "Can't get data: " + err.Error()})
		return
	}

	// 回傳 JSON 給前端
	c.JSON(http.StatusOK, gin.H{"code":1000, "data":geminiText})
}