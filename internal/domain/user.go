package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellido        string             `bson:"apellido" json:"apellido,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          UserImage          `bson:"avatar" json:"avatar,omitempty"`
	Banner          UserImage          `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitioweb" json:"sitioweb,omitempty"`
	Token           string             `json:"token,omitempty"`
}

type UserImage struct {
	PublicID string `bson:"public_id" json:"public_id,omitempty"`
	ImgUrl   string `bson:"imgurl" json:"imgurl,omitempty"`
}
