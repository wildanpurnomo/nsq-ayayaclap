package nsqutil

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/nsqio/go-nsq"
	smtputil "github.com/wildanpurnomo/nsq-ayayaclap/smtp-service/smtp"
)

type ConsumerWrapper struct {
	nsqConsumer *nsq.Consumer
}

type MessageHandler struct {
}

var (
	SMTPConsumer *ConsumerWrapper
)

func (m *MessageHandler) HandleMessage(message *nsq.Message) error {
	if len(message.Body) == 0 {
		return errors.New("body is blank, reenqueue message")
	}

	var data map[string]interface{}
	json.Unmarshal(message.Body, &data)

	if err := smtputil.SendRegistrationConfirmationMail(data["data"].(string)); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func InjectConsumer(consumer *nsq.Consumer) {
	SMTPConsumer = &ConsumerWrapper{nsqConsumer: consumer}
}

func CreateNewConsumer() (*nsq.Consumer, error) {
	config := nsq.NewConfig()

	config.MaxAttempts = 10
	config.MaxInFlight = 10
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0

	consumer, err := nsq.NewConsumer("register_new_user", "smtp", config)
	if err != nil {
		return nil, err
	}

	consumer.AddHandler(&MessageHandler{})

	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func StopConsumer() {
	SMTPConsumer.nsqConsumer.Stop()
}
