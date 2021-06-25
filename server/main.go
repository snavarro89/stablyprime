package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	api "github.com/snavarro89/stablyprime/api/numbers"
	app "github.com/snavarro89/stablyprime/app"
	primeget "github.com/snavarro89/stablyprime/functions/prime/get"
)

var application app.App

func init() {
	if err := godotenv.Load("../environment/.env"); err != nil {
		log.Print("No .env file found")
	}

	application = app.App{
		Data: app.Data{
			Numbers: api.NumbersModel{DB: nil},
		},
	}

}

func main() {

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Authorization", "X-Auth-Token", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"})

	router := mux.NewRouter().StrictSlash(true)
	primeget.Data(application)
	router.HandleFunc("/prime/{number}", primeget.Mux).Methods("GET")

	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}
