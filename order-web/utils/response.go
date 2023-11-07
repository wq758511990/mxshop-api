package utils

type ResponseStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(code int, message string) ResponseStruct {
	return ResponseStruct{Code: code, Message: message, Data: nil}
}

func (r *ResponseStruct) WithData(data interface{}) *ResponseStruct {
	return &ResponseStruct{
		Code:    r.Code,
		Message: r.Message,
		Data:    data,
	}
}

func (r *ResponseStruct) WithMessage(message string) *ResponseStruct {
	return &ResponseStruct{
		Code:    r.Code,
		Message: message,
		Data:    r.Data,
	}
}
