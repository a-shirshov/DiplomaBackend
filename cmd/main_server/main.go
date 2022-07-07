package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {

	viper.AddConfigPath("../../config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("Config isn't found")
		os.Exit(1)
	}

	r := mux.NewRouter()
	rApi := r.PathPrefix("/api").Subrouter()

	rApi.HandleFunc("/hello",func(w http.ResponseWriter, r* http.Request){
		w.Write([]byte("Hello!"))
	})

	port := viper.GetString("base.port")
	if err := http.ListenAndServe(":" + port, r); err != nil {
		log.Fatal(err)
	}
}