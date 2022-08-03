package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

///"mongodb+srv://evcvera:ernestit0@tdlc.jdjtp.mongodb.net/?retryWrites=true&w=majority"

func NewConnection(dataSource string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(dataSource)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Printf("Success")
	return client
}

func CheckConnection(connection *mongo.Client) bool {
	err := connection.Ping(context.TODO(), nil)
	return !(err != nil)
}
