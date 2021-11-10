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
	go func() {
		fmt.Printf("Server is running on : %v\n", os.Getenv("API_PORT"))

		if err := s.ListenAndServe(); err != nil {
			fmt.Errorf("Server failed to start : %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
