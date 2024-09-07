package api

// Forbidden returns an forbidden message, user is authenticated but insufficient permissions
func Forbidden(message string) Response {
	return Response{Status: "forbidden", Message: message, Data: map[string]any{}}
}

// ForbiddenWithData returns a forbidden message with data
func ForbiddenWithData(message string, data map[string]any) Response {
	return Response{Status: "success", Message: message, Data: data}
}
