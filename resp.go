package resp

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
		return m.Code == SUCCESS_CODE
	}
}

type ResponseCode struct {
	Code string
	Info string
}

type Error struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

const (
	SUCCESS_CODE       = "RC00000"
	GENERAL_ERROR_CODE = "RC60000"
	PARAM_ERROR_CODE   = "RC80000"
	SYS_ERROR_CODE     = "RC90000"
	SAVE_ERROR_CODE    = "RC90001"
	DELETE_ERROR_CODE  = "RC90002"
)

var SUCCESS = &ResponseCode{Code: "RC00000", Info: "SUCCESS"}
var GENERAL_ERROR = &ResponseCode{Code: "RC60000", Info: "FAILURE"}
var PARAM_ERROR = &ResponseCode{Code: "RC70000", Info: "ParamError"}
var DB_ERROR = &ResponseCode{Code: "R890000", Info: "DBError"}
var SYS_ERROR = &ResponseCode{Code: "RC90000", Info: "SysError"}

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
