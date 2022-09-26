package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ConcesionarioBack/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
	}
	return client, ctx, cancel
}

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Connect successfully")
	return nil
}

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()
}

func GetAutos() model.Autos {
	uri := os.Getenv("URI")
	db := os.Getenv("DATABASE")
	col := os.Getenv("COLECCION_AUTOS")

	client, ctx, cancel := connect(uri)

	collection := client.Database(db).Collection(col)

	defer close(client, ctx, cancel)

	ping(client, ctx)

	data, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Panic(err)
	}

	var results []model.Auto
	if err = data.All(ctx, &results); err != nil {
		log.Panic(err)
	}

	return results
}

func GetAuto(id int) model.Auto {
	uri := os.Getenv("URI")
	db := os.Getenv("DATABASE")
	col := os.Getenv("COLECCION_AUTOS")

	client, ctx, cancel := connect(uri)

	collection := client.Database(db).Collection(col)

	defer close(client, ctx, cancel)

	ping(client, ctx)

	var result = model.Auto{}
	err := collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&result)
	if err != nil {
		log.Panic(err)
	}

	return result
}

func AddUsuario(usuario model.Usuario) error {
	uri := os.Getenv("URI")
	db := os.Getenv("DATABASE")
	col := os.Getenv("COLECCION_USUARIOS")

	client, ctx, cancel := connect(uri)

	collection := client.Database(db).Collection(col)

	usuarioDB := bson.D{
		{Key: "nombre", Value: usuario.Nombre}, 
		{Key: "apellido", Value: usuario.Apellido}, 
		{Key: "email", Value: usuario.Email}, 
		{Key: "detalles", Value: usuario.Detalles},
	}
	_, err := collection.InsertOne(context.TODO(), usuarioDB)
	if err != nil {
		log.Panic(err)
		return err
	}

	ping(client, ctx)

	defer close(client, ctx, cancel)

	return err
}
