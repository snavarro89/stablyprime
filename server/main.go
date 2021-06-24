package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	primeget "github.com/snavarro89/stablyprime/functions/prime/get"
)

//var application App

func init() {
	if err := godotenv.Load("../environment/.env"); err != nil {
		log.Print("No .env file found")
	}

}

func main() {

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Authorization", "X-Auth-Token", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"})

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/prime", primeget.Mux).Methods("POST")

	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}
