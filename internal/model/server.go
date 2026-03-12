package model

// 伺服器 api 返回的股票格式
type GetStocksRet struct{
	Type string
	Date string
	Close float64
}

// 伺服器 api 返回的 Gemini text 格式
type GeminiJson struct{
	Date  string
	Point float64
	Price float64
	Reason string
}