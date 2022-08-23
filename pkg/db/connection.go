package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// var (
// 	StorageDB *sql.DB
// )

func NewConnection() *sql.DB {
	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	// SIN EL var err error, no funciona.
	dataSource := os.Getenv("DATA_SOURCE")
	StorageDB, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	// asd, _ := StorageDB.Query("select * from users")
	// fmt.Println(asd.Columns())
	// fmt.Println("dasdsa")
	// log.Println("database Configured")
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
