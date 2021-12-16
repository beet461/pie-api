package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Respond(w http.ResponseWriter, v interface{}) {
	enc, err := json.Marshal(v)
	errorCheck(err, "RESPONSE ENCODE")
	i, err := w.Write([]byte(enc))
	errorCheck(err, "Returned int:", i, "\n")
}

func Unpack(w http.ResponseWriter, r *http.Request, s interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	errorCheck(err, "BODY READING")

	jerr := json.Unmarshal([]byte(body), s)
	errorCheck(jerr, "UNMARSHAL")
}
