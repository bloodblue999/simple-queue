package rmq

import (
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"time"
)

type Broker struct {
	Queue     rmq.Queue
	Responder *Responder
}

func (b *Broker) MakeRequest() ConsumerResponse {
	requestId := b.Responder.CreateResponseChannel()

	request := ConsumerRequest{
		RequestId: requestId,
	}

	requestJsonBytes, err := json.Marshal(request)
	if err != nil {
		b.Responder.channels.Delete(requestId)
		return ConsumerResponse{
			Error: err,
		}
	}

	err = b.Queue.PublishBytes(requestJsonBytes)
	if err != nil {
		b.Responder.channels.Delete(requestId)
		return ConsumerResponse{
			Error: err,
		}
	}

	return b.Responder.WaitForResponse(requestId, time.Second*30)
}
