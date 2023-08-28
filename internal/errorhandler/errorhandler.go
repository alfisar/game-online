package errorhandler

type ErrorData struct {
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

func ErrorMapping(rc int, message string, err error) ErrorData {
	return ErrorData{
		ResponseCode: rc,
		Message:      message,
		Data:         err.Error(),
	}
}
