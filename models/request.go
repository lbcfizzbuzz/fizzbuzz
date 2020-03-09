package models

// Request represents a request received by the server
type Request struct {
	Int1Param  uint64 `json:"int1_param"`
	Int2Param  uint64 `json:"int2_param"`
	LimitParam uint64 `json:"limit_param"`
	Str1Param  string `json:"str1_param"`
	Str2Param  string `json:"str2_param"`
	Count      int    `json:"count"`
}
