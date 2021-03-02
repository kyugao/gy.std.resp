package stdresp

import "fmt"

type Response struct {
	Code   string      `json:"code"`
	Info   string      `json:"info"`
	Result interface{} `json:"result,omitempty"`
}

func (m *Response) isSuccess() bool {
	if m == nil {
		return false
	} else {
		return m.Code == SUCCESS.Code
	}
}

type Error struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

func GenSuccess(result interface{}) (resp *Response) {
	resp = Gen(SUCCESS, "", result)
	return
}

func GenError(code *ResponseCode, info string) (resp *Response) {
	resp = Gen(code, info, nil)
	return
}

func Gen(code *ResponseCode, hint string, result interface{}) (resp *Response) {
	resp = &Response{Code: code.Code}
	if len(hint) > 0 {
		resp.Info = fmt.Sprintf("%s: %s", code.Info, hint)
	} else {
		resp.Info = code.Info
	}
	resp.Result = result
	return
}
