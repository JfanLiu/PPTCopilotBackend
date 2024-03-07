package controllers

// Response 返回响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// MakeResponse 生成一个响应
func MakeResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
