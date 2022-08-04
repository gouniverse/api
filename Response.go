package api

import "encoding/json"

// Response defines an response for the API
type Response struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func (r Response) ToString() string {
	str, _ := json.Marshal(r)
	return string(str)
}
