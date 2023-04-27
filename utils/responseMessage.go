package utils

type errorObject struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DataObject struct {
	Item       interface{} `json:"item,omitempty"`
	Items      interface{} `json:"items,omitempty"`
	Total      *int64      `json:"total,omitempty"`
	Id         uint        `json:"id,omitempty"`
	StatusCode *int        `json:"statusCode,omitempty"`
}

type ErrorMessagePrototype struct {
	APIVersion string      `json:"apiVersion"`
	Error      errorObject `json:"error"`
}

type SuccessMessagePrototype struct {
	APIVersion string     `json:"apiVersion"`
	Data       DataObject `json:"data"`
}

func ErrorMessage(message string, code int) ErrorMessagePrototype {
	err := errorObject{
		Code:    code,
		Message: message,
	}
	return ErrorMessagePrototype{APIVersion: "1.0.0", Error: err}
}

func SuccessMessage(data DataObject) SuccessMessagePrototype {
	return SuccessMessagePrototype{APIVersion: "1.0.0", Data: data}
}
