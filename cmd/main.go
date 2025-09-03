package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet{
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message" : "pong"})
		}
	})

	http.ListenAndServe(":8000", nil)
}