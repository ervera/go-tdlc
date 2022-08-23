package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// var (
// 	StorageDB *sql.DB
// )

func NewConnection() *sql.DB {
	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	// SIN EL var err error, no funciona.
	dataSource := "postgres://pwvhgatecfevnn:2fd1a4fd5a04c4a8fe659816c1a75d3e2d74d43da53e0efa8157916abaa0953e@ec2-34-199-68-114.compute-1.amazonaws.com:5432/d9p1i6md4jhb75"
	StorageDB, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	return StorageDB
}

// package db

// import (
// 	"context"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// ///"mongodb+srv://evcvera:ernestit0@tdlc.jdjtp.mongodb.net/?retryWrites=true&w=majority"

// func NewConnection(dataSource string) *mongo.Client {
// 	clientOptions := options.Client().ApplyURI(dataSource)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 		return client
// 	}
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 		return client
// 	}
// 	log.Printf("Success")
// 	return client
// }

// func CheckConnection(connection *mongo.Client) bool {
// 	err := connection.Ping(context.TODO(), nil)
// 	return !(err != nil)
// }
