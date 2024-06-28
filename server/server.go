package server

import (
	rmq2 "github.com/savi2w/simple-queue/rmq"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/savi2w/simple-queue/router"
)

func Run(broker *rmq2.Broker) error {
	server := echo.New()
	server.Use(middleware.Recover())

	router := &router.Router{
		Server: server,
		Broker: broker,
	}

	server.GET("/", router.Handler)

	server.HideBanner = true
	server.HidePort = true

	log.Println("starting server...")

	return server.Start(":3001")
}
