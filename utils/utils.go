package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

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

func GetIDPlaceholders(ids []int) string {
	var str_ids strings.Builder
	for index := range ids {
		if str_ids.Len() == 0 {
			fmt.Fprintf(&str_ids, "$%d", index+1)
		} else {
			fmt.Fprintf(&str_ids, ",$%d", index+1)
		}
	}
	return str_ids.String()
}
