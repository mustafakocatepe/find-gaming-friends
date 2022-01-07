package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/mustafakocatepe/find-gaming-friends/model"
)

func RespondWithError(w http.ResponseWriter, status int, err model.Error) {
	RespondWithJSON(w, &model.Error{Code: status, Message: err.Message, Success: false}, status)
}

func RespondWithJSON(w http.ResponseWriter, v interface{}, status int) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err := w.Write(buf.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
