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
	r.Message = MessageCodeMap[SuccessCode]
	return r
}

func (r *ResultCode) Fail() *ResultCode {
	r.Code = FailCode
	r.Message = MessageCodeMap[FailCode]
	return r
}

const (
	SuccessCode             = "200"
	FailCode                = "400"
	ErrorCode               = "-1"
	UnauthorizedCode        = "401"
	NotFoundCode            = "404"
	MethodNotAllowedCode    = "405"
	InternalServerErrorCode = "500"
)

var MessageCodeMap = map[interface{}]string{
	SuccessCode:             "Success",
	FailCode:                "Fail",
	ErrorCode:               "Error",
	UnauthorizedCode:        "Unauthorized",
	NotFoundCode:            "Not Found",
	MethodNotAllowedCode:    "Method Not Allowed",
	InternalServerErrorCode: "Internal Server Error",
}
