package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func Respond(w http.ResponseWriter, data interface{}, err error) {
	response := Response{Data: data, Success: true}
	if err != nil {
		response.Data = nil
		response.Success = false
		response.Error = err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
