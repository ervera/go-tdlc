package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ervera/tdlc-gin/internal/localGoogle"
	"github.com/ervera/tdlc-gin/internal/login"
	"github.com/ervera/tdlc-gin/internal/media"
	"github.com/ervera/tdlc-gin/internal/team"
	"github.com/ervera/tdlc-gin/internal/user"
	"github.com/ervera/tdlc-gin/pkg/db"
	"github.com/ervera/tdlc-gin/pkg/sendgrid"
	"github.com/ervera/tdlc-gin/src/api/handlers"
	"github.com/ervera/tdlc-gin/src/api/routers"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("error log")
	// }
	// //tkoen := os.Getenv("TOKEN")
	// db := db.NewConnection()
	// r := gin.Default()
	// handler := cors.AllowAll().Handler(r)
	// router := routes.NewRouter(r, db)
	// router.MapRoutes()
	// PORT := os.Getenv("PORT")
	// if PORT == "" {
	// 	PORT = "8080"
	// }
	// log.Fatal(http.ListenAndServe(":"+PORT, handler))

	err := godotenv.Load()
	if err != nil {
		fmt.Println("error log")
	}
	//tkoen := os.Getenv("TOKEN")
	db := db.NewConnection()
	repoUser := user.NewRepository(db)
	serviceMedia := media.NewService()
	serviceSendgrid := sendgrid.NewService()
	serviceUsers := user.NewService(repoUser, serviceSendgrid)

	serviceLogin := login.NewService(repoUser)

	serviceUser := user.NewService(repoUser, serviceSendgrid)
	localGoogleService := localGoogle.NewService(repoUser, serviceUser, serviceMedia)

	repoTeams := team.NewRepository(db)
	serviceTeam := team.NewService(repoTeams)

	userHandler := handlers.NewHandlerUser(serviceUsers, serviceMedia)
	loginHandler := handlers.NewLoginHandler(serviceLogin)
	googleHandler := handlers.NewGoogleHandler(localGoogleService)
	teamHandler := handlers.NewTeamHandler(serviceTeam)

	mapper := routers.NewMapper(
		userHandler,
		loginHandler,
		googleHandler,
		teamHandler,
	)

	router := routers.CreateRouter(mapper)
	routerHandler := cors.AllowAll().Handler(router)
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+PORT, routerHandler))
	//_ = router.Run(PORT)
}
