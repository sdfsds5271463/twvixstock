package model

// 旗標 flag 參數型態
type KindName string
var CrawlerKind = struct {
    All     KindName
    TAIEX   KindName
    VIXTWN  KindName
    TAIPE  KindName
}{
    All:   "all",
    TAIEX:  "taiex",
    VIXTWN: "vixtwn",
    TAIPE:  "taipe",
}

//TaiexApi 返回格式
type TaiexApi struct{
	Msg string           `json:"msg"`
	Status int64         `json:"status"`
	Data []map[string]interface{}   `json:"data"`
}

//TaipeApi 返回格式
type TaipeApi struct{
    ResultCode int64     `json:"ResultCode"`
	ResultMsg string     `json:"ResultMsg"`
	Data []map[string]interface{}         `json:"Data"`
	ExecuteTime string   `json:"ExecuteTime"`
}