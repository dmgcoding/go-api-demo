package myUtils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("5xx error, code: %v, msg: %v", code, msg)
	}

	type errorResponse struct {
		Msg string `json:"msg"`
	}

	RespondWithJson(w, code, errorResponse{
		Msg: msg,
	})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	//convert payload to binary
	dat, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	w.WriteHeader(code)
	//write header
	w.Header().Set("Content-Type", "application/json")
	//return value
	w.Write(dat)

}
