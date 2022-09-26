package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auto struct {
	IDMongo string `bson:"_id" json:"_id"`
	ID      int    `bson:"id" json:"id"`
	Marca   string `bson:"marca" json:"marca"`
	Modelo  string `bson:"modelo" json:"modelo"`
	Anio    int    `bson:"anio" json:"anio"`
	Km      int    `bson:"km" json:"km"`
	Precio  int    `bson:"precio" json:"precio"`
	Img     string `bson:"img" json:"img"`
}

type Autos []Auto

type Usuario struct {
	ID       primitive.ObjectID   `bson:"_id" json:"_id"`
	Nombre   string `bson:"nombre" json:"nombre"`
	Apellido string `bson:"apellido" json:"apellido"`
	Email    string `bson:"email" json:"email"`
	Detalles string `bson:"detalles" json:"detalles"`
}

type Usuarios []Usuario
