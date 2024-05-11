package result

import "github.com/moker/common/merr"

type ResponseSuccessBean struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{200, "OK", data}
}

type ResponseErrorBean struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func ErrorResp(errCode uint32) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, merr.StringMapErrMsg(errCode)}
}
