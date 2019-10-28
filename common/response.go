package common

type RestResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *RestResponse) SetData(d interface{}) {
	r.Data = d
}

func GetOKResponse(obj interface{}) RestResponse {
	r := RestResponse{
		Code: OkCode,
		Msg:  "OK",
		Data: obj,
	}
	return r
}

func GetErrorResponse(msg string) RestResponse {
	if msg == "" {
		msg = "Please have a retry"
	}

	r := RestResponse{
		Code: ErrorCode,
		Msg:  msg,
		Data: nil,
	}
	return r
}

const (
	OkCode    = 0
	ErrorCode = 9999
)
