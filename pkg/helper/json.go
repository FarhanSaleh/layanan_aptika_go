package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteResponseBody(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Error encode: ", err)
		panic(err)
	}
}

func ParseBody(r *http.Request, data any){
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		log.Println("Error decode: ", err)
		panic(err)
	}
}