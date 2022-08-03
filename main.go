package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ervera/tdlc-gin/cmd/server/routes"
	"github.com/ervera/tdlc-gin/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	dataSource := "mongodb+srv://evcvera:ernestit0@tdlc.jdjtp.mongodb.net/?retryWrites=true&w=majority"
	db := db.NewConnection(dataSource)
	r := gin.Default()
	handler := cors.AllowAll().Handler(r)
	router := routes.NewRouter(r, db)
	router.MapRoutes()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
