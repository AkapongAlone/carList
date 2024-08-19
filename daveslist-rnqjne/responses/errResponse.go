package responses

type AppErr struct {
	Status     bool
	StatusCode int
	Message    string
}

func NewAppErr(statusCode int, messageErr string) *AppErr {
	return &AppErr{Status: false, StatusCode: statusCode, Message: messageErr}
}

func (r *AppErr) Error() string {
	return r.Message
}

func SuccessResponse(data interface{}, code int) Success {
	return Success{
		Status: true,
		Code:   code,
		Data:   data,
	}
}

type Success struct {
	Status bool
	Code   int
	Data   interface{}
}
