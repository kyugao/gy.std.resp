package stdresp

import "fmt"

type PageResponse struct {
	Code     string        `json:"code"`
	Info     string        `json:"info"`
	PageInfo *PageInfo     `json:"page_info,omitempty"`
	Results  []interface{} `json:"results,omitempty"`
}

func (m *PageResponse) isSuccess() bool {
	if m == nil {
		return false
	} else {
		return m.Code == SUCCESS.Code
	}
}

type PageInfo struct {
	Total    int `json:"total"`
	PageNo   int `json:"page_no"`
	PageSize int `json:"page_size"`
}

func GenSuccessPageResponse(total, pageNo, pageSize int, results []interface{}) (resp *PageResponse) {
	resp = GenPageResponse(SUCCESS, "")
	resp.PageInfo = &PageInfo{
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
	}
	resp.Results = results
	return
}

func GenPageResponse(code *ResponseCode, hint string) (resp *PageResponse) {
	resp = &PageResponse{Code: code.Code}
	if len(hint) > 0 {
		resp.Info = fmt.Sprintf("%s: %s", code.Info, hint)
	} else {
		resp.Info = code.Info
	}
	return
}
