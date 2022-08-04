package api

import (
	"net/http"
)

// Respond writes a JSON or JSONP response
func RespondWithStatus(w http.ResponseWriter, r *http.Request, resp Response, status int) {
	w.WriteHeader(status)
	Respond(w, r, resp)
}
