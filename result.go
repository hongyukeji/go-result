package result

import (
	"encoding/json"
)

type Result struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResult(code string, message string, data interface{}) *Result {
	return &Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func Success() *Result {
	return NewResult(SuccessCode, MessageCodeMap[SuccessCode], nil)
}

func Fail() *Result {
	return NewResult(FailCode, MessageCodeMap[FailCode], nil)
}

func Error() *Result {
	return NewResult(ErrorCode, MessageCodeMap[ErrorCode], nil)
}

func (r *Result) GetMessage() string {
	return r.Message
}

func (r *Result) SetMessage(message string) *Result {
	r.Message = message
	return r
}

func (r *Result) GetCode() interface{} {
	return r.Code
}

func (r *Result) SetCode(code interface{}) *Result {
	r.Code = code.(string)
	//r.Code = fmt.Sprintf("%d", code)
	return r
}

func (r *Result) GetData() interface{} {
	return r.Data
}

func (r *Result) SetData(data interface{}) *Result {
	r.Data = data
	return r
}

func (r *Result) PutData(key string, value interface{}) *Result {
	data := map[string]interface{}{}
	if r.Data != nil {
		data = r.Data.(map[string]interface{})
	}
	data[key] = value
	r.Data = data
	return r
}

func (r *Result) toJson() map[string]any {
	// JSON 字节切片数据
	jsonData, _ := json.Marshal(r)
	// JSON 对象
	var jsonObject map[string]interface{}
	// JSON 字符串
	//jsonStr := string(jsonData)
	// 方式1 (推荐)
	_ = json.Unmarshal(jsonData, &jsonObject)
	// 方式2
	//_ = json.Unmarshal([]byte(jsonStr), &jsonObject)
	return jsonObject
}

func (r *Result) toJsonString() string {
	jsonData, _ := json.Marshal(r)
	jsonStr := string(jsonData)
	return jsonStr
}

func (r *Result) Error() string {
	return r.toJsonString()
}
