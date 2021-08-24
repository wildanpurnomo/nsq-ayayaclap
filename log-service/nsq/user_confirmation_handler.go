package nsqutil

import (
	"encoding/json"
	"time"

	"github.com/nsqio/go-nsq"
	dbmodels "github.com/wildanpurnomo/nsq-ayayaclap/log-service/db/models"
	dbrepos "github.com/wildanpurnomo/nsq-ayayaclap/log-service/db/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserConfirmationHandler struct {
}

func (h *UserConfirmationHandler) HandleMessage(message *nsq.Message) error {
	if len(message.Body) == 0 {
		return nil
	}

	var data map[string]interface{}
	json.Unmarshal(message.Body, &data)

	event := dbmodels.UserRegistrationEvent{
		EventName: "user_confirm",
		Email:     data["data"].(string),
		Timestamp: primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	if err := dbrepos.Repo.LogUserRegistrationEvent(event); err != nil {
		return err
	}

	return nil
}
