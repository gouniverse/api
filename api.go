package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Response defines an response for the API
type Response struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// Error returns an error message
func Error(message string) Response {
	return Response{Status: "error", Message: message, Data: map[string]interface{}{}}
}

// ErrorWithData returns an error message with data
func ErrorWithData(message string, data map[string]interface{}) Response {
	return Response{Status: "error", Message: message, Data: data}
}

// Success returns a success message
func Success(message string) Response {
	return Response{Status: "success", Message: message, Data: map[string]interface{}{}}
}

// SuccessWithData returns a success message with data
func SuccessWithData(message string, data map[string]interface{}) Response {
	return Response{Status: "success", Message: message, Data: data}
}

// Respond writes a JSON or JSONP response
func Respond(w http.ResponseWriter, r *http.Request, resp Response) {
	callback := strings.Trim(r.URL.Query().Get("callback"), "")

	response, _ := json.Marshal(resp)

	if callback != "" {
		response := callback + "(" + string(response) + ");"
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte(response))
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
