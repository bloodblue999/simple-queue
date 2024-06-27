package main

import (
	"log"

	"github.com/savi2w/simple-queue/rmq"
	"github.com/savi2w/simple-queue/server"
)

func main() {
	resultChannel := make(chan string)

	conn, queue, err := rmq.Run(resultChannel)
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer func() {
		channel := conn.StopAllConsuming()
		<-channel
	}()

	if err := server.Run(queue, resultChannel); err != nil {
		log.Println(err.Error())
		return
	}
}
