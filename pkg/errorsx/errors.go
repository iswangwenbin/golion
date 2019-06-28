package errorsx

import (
	"encoding/json"
)

type AppError struct {
	Id            string `json:"id"`
	Message       string `json:"message"`              // Message to be display to the end user without debugging information
	DetailedError string `json:"detailed_error"`       // Internal error string to help the developer
	RequestId     string `json:"request_id,omitempty"` // The RequestId that's also set in the header
	Code          int    `json:"code,omitempty"`       // The error code
	Where         string `json:"-"`                    // The function where it happened in the form of Struct.Func
	params        map[string]interface{}
}

func (er *AppError) Error() string {
	return er.Where + ": " + er.Message + ", " + er.DetailedError
}

func (er *AppError) ToJson() string {
	b, _ := json.Marshal(er)
	return string(b)
}

func New(where string, id string, params map[string]interface{}, details string, code int) *AppError {
	ap := &AppError{}
	ap.Id = id
	ap.params = params
	ap.Message = id
	ap.Where = where
	ap.DetailedError = details
	ap.Code = code
	return ap
}
