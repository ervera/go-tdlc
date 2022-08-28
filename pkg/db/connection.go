package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	StorageDB *sql.DB
// )

func NewConnection() *sql.DB {
	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	// SIN EL var err error, no funciona.
	// dataSource := "postgres://aofrimvubeextn:b8c9703a6ed32b362eefdba80459839a431c33903793ed4c58d2caa297fc5f31@ec2-3-214-2-141.compute-1.amazonaws.com:5432/d2cs8ep2ntrekh"
	// StorageDB, err := sql.Open("postgres", dataSource)
	DSN := "aklzy92giy81gwchwkbg:pscale_pw_jjFJOZZ6dck7ZNxxWEP8zMegDbvAQ8rjIWt6lIddPvv@tcp(aws-sa-east-1.connect.psdb.cloud)/evcvera?tls=true&parseTime=true"
	StorageDB, err := sql.Open("mysql", DSN)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully connected to PlanetScale!")
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
