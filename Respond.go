package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

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
