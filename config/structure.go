package config

const (
	RSP_SUCCESS   = 0
	RSP_ERRPR     = 1
	RSP_ERROR_MSG = "系统错误"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type NewPlayGame struct {
	WeddingId int64 `json:"weddingId"`
	Seconds   int64 `json:"seconds"`
}

type HLBUserContent struct {
	Type    int64  `json:"type"`
	Content string `json:"content"`
}

type HLBUser struct {
	Type    string  `json:"type"`
	IdList  []int64 `json:"idList"`
	Content string  `json:"content"`
}
