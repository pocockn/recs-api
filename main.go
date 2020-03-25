package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pocockn/recs-api/config"
	"github.com/pocockn/recs-api/persistance"
	"github.com/pocockn/recs-api/recs/delivery"
	"github.com/pocockn/recs-api/recs/store"
	"log"
	"net/http"
)

func main() {
	config := config.NewConfig()

	db, err := persistance.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	connection, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	//TODO: this should come from config.
	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:8081", "http://localhost:3000"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
		middleware.Logger(),
	)

	recRepo := store.NewRecsStore(connection)
	echoHandler := delivery.NewHandler(config, recRepo)
	echoHandler.Register(e)

	e.Logger.Fatal(e.Start(":1323"))
}
