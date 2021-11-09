package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/imrushi/Go-REST-API/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/health", handler.Health).Methods("GET")
	r.HandleFunc("/api/getProducts", handler.GetProducts).Methods("GET")
	r.HandleFunc("/api/addProducts", handler.AddProduct).Methods("POST")
	r.HandleFunc("/api/updateProduct/{id:[0-9]+}", handler.UpdateProducts).Methods("PUT")
	r.HandleFunc("/api/deleteProduct/{id:[0-9]+}", handler.DeleteProducts).Methods(http.MethodDelete)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	s := &http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv("API_PORT")),
		Handler: loggedRouter,
	}
	fmt.Printf("Server is running on : %v", os.Getenv("API_PORT"))
	if err := s.ListenAndServe(); err != nil {
		fmt.Errorf("Server failed to start : %s", err)
	}
}
