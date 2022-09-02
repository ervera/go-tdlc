package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ervera/tdlc-gin/cmd/server/routes"
	"github.com/ervera/tdlc-gin/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error log")
	}
	//tkoen := os.Getenv("TOKEN")
	db := db.NewConnection()
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
