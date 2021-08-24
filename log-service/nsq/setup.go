package nsqutil

import (
	"time"

	"github.com/nsqio/go-nsq"
)

type ConsumerWrapper struct {
	nsqConsumers []*nsq.Consumer
}

var (
	LoggerConsumerWrapper *ConsumerWrapper
)

func InitLoggerConsumers() {
	LoggerConsumerWrapper = &ConsumerWrapper{nsqConsumers: make([]*nsq.Consumer, 0)}
}

func AppendConsumer(consumer *nsq.Consumer) {
	LoggerConsumerWrapper.nsqConsumers = append(LoggerConsumerWrapper.nsqConsumers, consumer)
}

func CreateNewConsumer(topic string, channel string, handler nsq.Handler) (*nsq.Consumer, error) {
	config := nsq.NewConfig()

	config.MaxAttempts = 10
	config.MaxInFlight = 10
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0

	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, err
	}

	consumer.AddHandler(handler)

	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func StopConsumers() {
	for _, consumer := range LoggerConsumerWrapper.nsqConsumers {
		consumer.Stop()
	}
}
