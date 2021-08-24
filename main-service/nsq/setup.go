package nsqutil

import (
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

type Publisher struct {
	Producer *nsq.Producer
}

type Event struct {
	EventName string      `json:"event_name"`
	Data      interface{} `json:"data"`
}

var NsqPublisher *Publisher

func (p Publisher) Publish(topic string, event Event) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if err = p.Producer.Publish(topic, payload); err != nil {
		return err
	} else {
		return nil
	}
}

func InjectProducer(producer *nsq.Producer) {
	NsqPublisher = &Publisher{Producer: producer}
}

func CreateNewProducer() (*nsq.Producer, error) {
	config := nsq.NewConfig()

	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func StopProducer() {
	NsqPublisher.Producer.Stop()
	log.Println("Main service producer is stopping")
}
