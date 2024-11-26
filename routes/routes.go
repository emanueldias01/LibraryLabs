package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)


var r = mux.NewRouter()

func HandleRequest(){
	http.ListenAndServe(":8000", r)
}