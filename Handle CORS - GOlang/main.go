package main

import (
	"log"
	"net/http"
	B "smartdq/middleware/src/azure"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func handleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/test", B.Test).Methods("GET")
	r.HandleFunc("/getdata", B.GetData).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})

	handler := c.Handler(r)
	log.Fatal((http.ListenAndServe(":8089", handler)))
	http.Handle("/", r)
}

func main() {
	handleRequests()
}
