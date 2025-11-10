package api

import (
	"api_demo/internal/utils"
	"net/http"
	"strings"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	
	type successRes struct {
		Message string `json:"message"`
	}

	type errorRes struct {
		Error string `json:"error"`
	}
	
	if len(name) > 0 {
		first := strings.ToUpper(name)[0]
		if first >= 'A' && first <= 'M' {
			myUtils.RespondWithJson(w, http.StatusOK, successRes{Message: "Hello " + name})
		} else {
			myUtils.RespondWithJson(w, http.StatusBadRequest, errorRes{Error: "Invalid Input"})
		}
	}else{
		myUtils.RespondWithJson(w, http.StatusBadRequest, errorRes{Error: "Invalid Input"})
	}
}