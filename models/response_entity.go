package models

type ResponseEntity struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var ResponseEntityInstance = ResponseEntity{}

func (r *ResponseEntity) Ok(data interface{}) ResponseEntity {
	r.Code = 200
	r.Message = "success"
	r.Data = data
	return *r
}

func (r *ResponseEntity) Fail(data interface{}) ResponseEntity {
	r.Code = 500
	r.Message = "fail"
	r.Data = data
	return *r
}

func (r *ResponseEntity) FailMessage(msg string) ResponseEntity {
	r.Code = 500
	r.Message = msg
	return *r
}

func (r *ResponseEntity) Msg(code int, msg string, data interface{}) ResponseEntity {
	r.Code = code
	r.Message = msg
	r.Data = data
	return *r
}
