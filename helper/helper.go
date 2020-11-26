package helper

import "github.com/go-playground/validator/v10"

//Response struct
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

//Meta struct
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

//APIResponse return response api
func APIResponse(code int, message string, status string, data interface{}) Response {
	meta := Meta{
		Code:    code,
		Message: message,
		Status:  status,
	}

	JSONResponse := Response{
		Meta: meta,
		Data: data,
	}

	return JSONResponse
}

//FormatError error response
func FormatError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
