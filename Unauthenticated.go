package api

// Unauthenticated returns an unauthenticated message, user is not known (not authenticated)
func Unauthenticated(message string) Response {
	return Response{Status: TYPE_UNAUTHENTICATED, Message: message, Data: map[string]any{}}
}

// Unauthenticated returns an unauthenticated message, user is not known (not authenticated)
func UnauthenticatedWithData(message string, data map[string]any) Response {
	return Response{Status: TYPE_UNAUTHENTICATED, Message: message, Data: data}
}
