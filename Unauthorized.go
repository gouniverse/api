package api

// Unauthorized returns an unauthorized message, user is known but not unauthorized to do the action
func Unauthorized(message string) Response {
	return Response{Status: TYPE_UNAUTHORIZED, Message: message, Data: map[string]any{}}
}

// Unauthorized returns an unauthorized message, user is known but not unauthorized to do the action
func UnauthorizedWithData(message string, data map[string]any) Response {
	return Response{Status: TYPE_UNAUTHORIZED, Message: message, Data: data}
}
