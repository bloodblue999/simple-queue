package rmq

import (
	"log"
	"time"

	"github.com/adjust/rmq/v5"
)

const RedisDatabase = 1
const PrefetchLimit = 2
const PollDuration = time.Second

func Run(resultChannel chan string) (conn rmq.Connection, queue rmq.Queue, err error) {
	error_channel := make(chan error)

	go func() {
		for err := range error_channel {
			log.Println(err.Error())
		}
	}()

	conn, err = rmq.OpenConnection("simple_queue_service", "tcp", "localhost:6379", RedisDatabase, error_channel)
	if err != nil {
		return nil, nil, err
	}

	queue, err = conn.OpenQueue("default_queue")
	if err != nil {
		return nil, nil, err
	}

	if err := queue.StartConsuming(PrefetchLimit, PollDuration); err != nil {
		return nil, nil, err
	}

	if _, err := queue.AddConsumer("default_consumer", Consumer{resultChannel: resultChannel}); err != nil {
		return nil, nil, err
	}

	log.Println("starting queue...")

	return conn, queue, nil
}
