package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el model del tipo usuario de mongodb*/
type Usuario struct {
	ID 	primitive.ObjectID 	   `bson:"_id, omitempty" json:"id"`
	Nombre	    string     	   `bson:"nombre" json:"nombre,omitempty"`
	Apellidos 	string   	   `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time  `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email       string         `bson:"email" json:"email"`
	Password    string		   `bson:"password" json:"password,omitempty"`
	Banner		string		   `bson:"banner" json:"banner"`
	Biografia	string		   `bson:"biografia" json:"biografia"`
	Ubicacion   string         `bson:"ubicacion" json:"ubicacion"`
	SitioWeb    string         `bson:"sitioWeb" json:"sitioWeb"`

}