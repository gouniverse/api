package api

// Success returns a success message
func Success(message string) Response {
	return Response{Status: TYPE_SUCCESS, Message: message, Data: map[string]any{}}
}

// SuccessWithData returns a success message with data
func SuccessWithData(message string, data map[string]any) Response {
	return Response{Status: TYPE_SUCCESS, Message: message, Data: data}
}
