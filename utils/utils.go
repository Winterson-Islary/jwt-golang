package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(req *http.Request, payload any) error {
	if req.Body == nil {
		return fmt.Errorf("Missing Request body")
	}
	return json.NewDecoder(req.Body).Decode(payload)
}

func WriteJSON(res http.ResponseWriter, status int, result any) error {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(status)

	return json.NewEncoder(res).Encode(result)
}

func WriteError(res http.ResponseWriter, status int, err error) {
	WriteJSON(res, status, map[string]string{"error": err.Error()})
}
