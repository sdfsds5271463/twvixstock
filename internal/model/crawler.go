package model

// 旗標 flag 參數型態
type KindName string
var CrawlerKind = struct {
    All     KindName
    TAIEX   KindName
    VIXTWN  KindName
    TAIPE   KindName
    GEMINI  KindName
}{
    All:    "all",
    TAIEX:  "taiex",
    VIXTWN: "vixtwn",
    TAIPE:  "taipe",
    GEMINI: "gemini",  //抓取近三個月的股市分析
}

//檢查是否要爬蟲
type NeedCrawler struct{
    TAIEX   bool
    VIXTWN  bool
    TAIPE   bool
    GEMINI  bool
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

//Gemini Api 返回格式
type GeminiApi struct{
    Candidates    []G_candidates    `json:"candidates"`
    UsageMetadata G_usageMetadata   `json:"usageMetadata"`
    ModelVersion  string            `json:"modelVersion"`
    ResponseId    string            `json:"responseId"`
}
type G_candidates struct{
    Content struct{
        Parts []map[string]string
        Role string
    }
    FinishReason string
    Index int
}
type G_usageMetadata struct{
    PromptTokenCount int
    CandidatesTokenCount int
    TotalTokenCount int
    PromptTokensDetails []struct{
        Modality string
        TokenCount int
    }
}
