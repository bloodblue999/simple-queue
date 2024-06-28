package rmq

import (
	"log"
	"sync"
	"time"

	"github.com/adjust/rmq/v5"
)

const RedisDatabase = 1
const PrefetchLimit = 4
const PollDuration = time.Second

func Run() (conn rmq.Connection, responder *Responder, queue rmq.Queue, err error) {
	error_channel := make(chan error)

	go func() {
		for err := range error_channel {
			log.Println(err.Error())
		}
	}()

	conn, err = rmq.OpenConnection("simple_queue_service", "tcp", "localhost:6379", RedisDatabase, error_channel)
	if err != nil {
		return nil, nil, nil, err
	}

	queue, err = conn.OpenQueue("default_queue")
	if err != nil {
		return nil, nil, nil, err
	}

	if err := queue.StartConsuming(PrefetchLimit, PollDuration); err != nil {
		return nil, nil, nil, err
	}

	responder = &Responder{
		channels: sync.Map{},
	}

	if _, err := queue.AddConsumer("default_consumer", Consumer{responder: responder}); err != nil {
		return nil, nil, nil, err
	}
	if _, err := queue.AddConsumer("default_consumer", Consumer{responder: responder}); err != nil {
		return nil, nil, nil, err
	}
	if _, err := queue.AddConsumer("default_consumer", Consumer{responder: responder}); err != nil {
		return nil, nil, nil, err
	}
	if _, err := queue.AddConsumer("default_consumer", Consumer{responder: responder}); err != nil {
		return nil, nil, nil, err
	}

	log.Println("starting queue...")

	return conn, responder, queue, nil
}
