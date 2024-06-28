package rmq

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"time"

	"github.com/adjust/rmq/v5"
)

type Consumer struct {
	responder *Responder
}

type ConsumerRequest struct {
	RequestId string
}

type ConsumerResponse struct {
	Uuid  string
	Error error
}

func (c Consumer) Consume(delivery rmq.Delivery) {
	var request ConsumerRequest

	err := json.Unmarshal([]byte(delivery.Payload()), &request)
	if err != nil {
		log.Println(err.Error())
		delivery.Reject()
		return
	}

	time.Sleep(2 * time.Second)

	response := ConsumerResponse{
		Uuid: uuid.NewString(),
	}
	log.Println(response.Uuid)

	c.responder.SendResponse(request.RequestId, response)

	if err := delivery.Ack(); err != nil {
		log.Println(err.Error())
	}
}
