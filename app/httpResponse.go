package app

import "log"

type HttpResponseType[T interface{}] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data    T      `json:"data,omitempty"`
}

func HttpSuccessResponse(message string, code int, data interface{}) HttpResponseType[interface{}] {
	return HttpResponseType[interface{}]{
		Code:    code,
		Message: message,
		Status:  true,
		Data:    data,
	}
}

func HttpErrorResponse(message string, code int, data interface{}, err error) HttpResponseType[interface{}] {
	if err != nil {
		log.Printf("Error: %s: %v", message, err)
	}
	return HttpResponseType[interface{}]{
		Code:    code,
		Message: message,
		Status:  false,
		Data:    data,
	}
}
