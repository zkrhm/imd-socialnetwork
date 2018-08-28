package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	resp := make(map[string]string)

	resp["message"] = error
	b, err := json.Marshal(resp)
	if err != nil {

	}
	fmt.Fprintln(w, string(b))
}
