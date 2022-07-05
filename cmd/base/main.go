package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	rApi := r.PathPrefix("/api").Subrouter()

	rApi.HandleFunc("/hello",func(w http.ResponseWriter, r* http.Request){
		w.Write([]byte("Hello!"))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}