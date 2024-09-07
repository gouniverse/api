package api

// Error returns an error message
func Error(message string) Response {
	return Response{Status: "error", Message: message, Data: map[string]any{}}
}

// ErrorWithData returns an error message with data
func ErrorWithData(message string, data map[string]any) Response {
	return Response{Status: "error", Message: message, Data: data}
}
