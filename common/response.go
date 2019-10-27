package common

type RestResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetOKResponse(obj interface{}) RestResponse {
	r := RestResponse{
		Code: OkCode,
		Msg:  "OK",
		Data: obj,
	}
	return r
}

func GetErrorResponse() RestResponse {
	r := RestResponse{
		Code: ErrorCode,
		Msg:  "Please have a retry",
		Data: nil,
	}
	return r
}

const (
	OkCode    = 0
	ErrorCode = 9999
)
