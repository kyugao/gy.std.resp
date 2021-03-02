package stdresp

type ResponseCode struct {
	Code string
	Info string
}

var SUCCESS = &ResponseCode{Code: "RC00000", Info: "SUCCESS"}
var GENERAL_ERROR = &ResponseCode{Code: "RC60000", Info: "FAILURE"}
var PARAM_ERROR = &ResponseCode{Code: "RC70000", Info: "ParamError"}
var DB_ERROR = &ResponseCode{Code: "RC80000", Info: "DBError"}
var SYS_ERROR = &ResponseCode{Code: "RC90000", Info: "SysError"}
