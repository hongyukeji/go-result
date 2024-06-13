package result

type ResultCode struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
}

func NewResultCode() *ResultCode {
	return &ResultCode{}
}

func (r *ResultCode) Success() *ResultCode {
	r.Code = SuccessCode
	r.Message = CodeMessage[SuccessCode]
	return r
}

func (r *ResultCode) Fail() *ResultCode {
	r.Code = FailCode
	r.Message = CodeMessage[FailCode]
	return r
}

const (
	SuccessCode             = "200"
	FailCode                = "400"
	ErrorCode               = "-1"
	CodeUnauthorized        = "401"
	CodeNotFound            = "404"
	CodeMethodNotAllowed    = "405"
	CodeInternalServerError = "500"
)

var CodeMessage = map[interface{}]string{
	SuccessCode:             "成功",
	FailCode:                "失败",
	ErrorCode:               "错误",
	CodeUnauthorized:        "未授权",
	CodeNotFound:            "未找到",
	CodeMethodNotAllowed:    "方法不被允许",
	CodeInternalServerError: "服务器内部错误",
}
