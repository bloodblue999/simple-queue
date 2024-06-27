package server

import (
	"log"

	"github.com/adjust/rmq/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/savi2w/simple-queue/router"
)

func Run(queue rmq.Queue, resultChannel chan string) error {
	server := echo.New()
	server.Use(middleware.Recover())

	router := &router.Router{
		Queue:         queue,
		Server:        server,
		ResultChannel: resultChannel,
	}

	server.GET("/", router.Handler)

	server.HideBanner = true
	server.HidePort = true

	log.Println("starting server...")

	return server.Start(":3001")
}
