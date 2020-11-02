package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/squok/twittor/middleware"
	"github.com/squok/twittor/routers"
)

/*Manejadores set handlers y serv arrriba */
func Manejadores(){
	router := mux.NewRouter()

	router.HandleFunc("/registro",middleware.ChequeoBD(routers.Registro)).Methods("POST")


	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT,handler))
}