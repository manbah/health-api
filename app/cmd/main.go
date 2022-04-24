package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"main-1/handler"
	"net/http"
)

func main() {
	server := mux.NewRouter()
	server.HandleFunc("/health", handler.Health())
	srv := http.Server{
		Addr:    ":8000",
		Handler: server,
	}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
