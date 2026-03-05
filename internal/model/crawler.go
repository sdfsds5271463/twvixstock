package model

// 旗標 flag 參數型態
type KindName string
var CrawlerKind = struct {
    Both    KindName
    TAIEX   KindName
    VIXTWN  KindName
}{
    Both:   "both",
    TAIEX:  "taiex",
    VIXTWN: "vixtwn",
}

//TaiexApi 返回格式
type TaiexApi struct{
	Msg string           `json:"msg"`
	Status int64         `json:"status"`
	Data []map[string]interface{}   `json:"data"`
}