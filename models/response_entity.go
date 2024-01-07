package models

import "net/http"

type ResponseEntity struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (r *ResponseEntity) Ok(data any) ResponseEntity {
	r.Code = http.StatusOK
	r.Message = "success"
	r.Data = data
	return *r
}

func (r *ResponseEntity) Fail(data any) ResponseEntity {
	r.Code = http.StatusInternalServerError
	r.Message = "fail"
	r.Data = data
	return *r
}

func (r *ResponseEntity) FailMessage(msg string) ResponseEntity {
	r.Code = http.StatusInternalServerError
	r.Message = msg
	return *r
}

func (r *ResponseEntity) Msg(code int, msg string, data any) ResponseEntity {
	r.Code = code
	r.Message = msg
	r.Data = data
	return *r
}
