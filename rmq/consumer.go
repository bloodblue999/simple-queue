package rmq

import (
	"github.com/google/uuid"
	"log"
	"time"

	"github.com/adjust/rmq/v5"
)

type Consumer struct {
	resultChannel chan string
}

func (c Consumer) Consume(delivery rmq.Delivery) {
	time.Sleep(2 * time.Second)

	generatedUUID := uuid.NewString()
	log.Println(generatedUUID)

	c.resultChannel <- generatedUUID

	if err := delivery.Ack(); err != nil {
		log.Println(err.Error())
	}
}
