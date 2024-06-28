package rmq

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

type Responder struct {
	channels sync.Map
}

func (r *Responder) CreateResponseChannel() string {
	requestID := uuid.NewString()

	channel := make(chan ConsumerResponse)

	r.channels.Store(requestID, channel)

	return requestID
}

func (r *Responder) WaitForResponse(requestID string, timeoutTime time.Duration) ConsumerResponse {
	channelInterface, ok := r.channels.Load(requestID)
	if !ok {
		return ConsumerResponse{
			Error: fmt.Errorf("no one is waiting for response of requestId %s", requestID),
		}
	}

	defer r.channels.Delete(requestID)

	channel, ok := channelInterface.(chan ConsumerResponse)
	if !ok {
		return ConsumerResponse{
			Error: fmt.Errorf("request with id %s has a invalid type", requestID),
		}
	}

	defer close(channel)

	var response ConsumerResponse
	select {

	case response = <-channel:

	case <-time.After(timeoutTime):
		response = ConsumerResponse{
			Error: errors.New("request reached timeout time"),
		}

	}

	return response
}

func (r *Responder) SendResponse(requestID string, response ConsumerResponse) {
	channelInterface, ok := r.channels.Load(requestID)
	if !ok {
		log.Println("no one is waiting for response of requestId ", requestID)
		return
	}

	channel, ok := channelInterface.(chan ConsumerResponse)
	if !ok {
		log.Printf("request with id %s has a invalid type\n", requestID)
		return
	}

	channel <- response
}
