package common

type response struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
	StatusText string      `json:"statusText"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
}

func NewStandardResponse(data interface{}, statusCode int, statusText, err string, message string) response {
	return response{Data: data, StatusCode: statusCode, StatusText: statusText, Error: err, Message: message}
}
