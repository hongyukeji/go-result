package result

import (
	"fmt"
	"testing"
)

func ResultTest(t *testing.T) {
	// 自定义数据结构
	data := map[string]interface{}{
		"name":    "result",
		"version": "1.0.0",
	}

	// 返回1 预置状态码和消息 [规范]
	response := Success().SetCode(SuccessCode).SetMessage(CodeMessage[SuccessCode]).SetData(data)
	fmt.Println(response)

	// 返回2 自定义状态码和消息 [灵活]
	res := Success().SetCode("200").SetMessage("OK").SetData(data)
	fmt.Println(res)

	// 返回错误1 预置状态码和消息 [规范]
	resultError := Error()
	resultException := Exception(resultError)
	fmt.Println(resultException)

	// 返回错误2 自定义状态码和消息 [灵活]
	responseError := Fail().SetCode("500").SetMessage("Fail")
	fmt.Println(responseError)

	// panic
	//panic(Exception(Error()))
}
