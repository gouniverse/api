package api

// Error returns an error message
func Error(message string) Response {
	return Response{Status: "error", Message: message, Data: map[string]interface{}{}}
}

// ErrorWithData returns an error message with data
func ErrorWithData(message string, data map[string]interface{}) Response {
	return Response{Status: "error", Message: message, Data: data}
}

// Forbidden returns an forbidden message, user is authenticated but insufficient permissions
func Forbidden(message string) Response {
	return Response{Status: "forbidden", Message: message, Data: map[string]interface{}{}}
}

// Success returns a success message
func Success(message string) Response {
	return Response{Status: "success", Message: message, Data: map[string]interface{}{}}
}

// SuccessWithData returns a success message with data
func SuccessWithData(message string, data map[string]interface{}) Response {
	return Response{Status: "success", Message: message, Data: data}
}

// Unauthenticated returns an unauthenticated message, user is not known (authenticated)
func Unauthenticated(message string) Response {
	return Response{Status: "unauthenticated", Message: message, Data: map[string]interface{}{}}
}

// Unauthorized returns an unauthorized message, user is known but not unauthorized to do the action
func Unauthorized(message string) Response {
	return Response{Status: "unauthorized", Message: message, Data: map[string]interface{}{}}
}
