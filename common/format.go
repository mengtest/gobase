package common

// RPFormat 用于描述一个包
type RPFormat struct {
	ErrCode int    `json:"errCode"`
	ErrDesc string `json:"errDesc"`
	Data    string `json:"data"`
}

func newRPFormat() *RPFormat {
	return &RPFormat{}
}
