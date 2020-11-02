package bd

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://pato:pato@cluster0.gdzcw.gcp.mongodb.net/twittor")

/* ConnectarBD conecta la base de datos*/
func ConnectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(),nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("connection exitosa de la bd")
	return client
}
/* ChequeoConnection cheque la conexion*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(),nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}