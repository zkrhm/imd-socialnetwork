package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetRequest(r *http.Request, reqObj interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	json.Unmarshal(b, reqObj)

	return nil
}

func WriteReponse(w http.ResponseWriter, resObj interface{}) error {
	b, err := json.Marshal(resObj)
	if err != nil {
		return err
	}
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintln(w, string(b))
	return nil
}
