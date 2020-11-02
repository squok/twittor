package main

import (
	"log"

	"github.com/squok/twittor/bd"
	"github.com/squok/twittor/handlers"
)
func main() {
	if bd.ChequeoConnection()==0 {
		log.Fatal("Sin coneccion a la base de datos")
		return
	}
	handlers.Manejadores()
}