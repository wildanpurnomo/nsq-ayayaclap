package dbrepos

import (
	"context"

	dbmodels "github.com/wildanpurnomo/nsq-ayayaclap/log-service/db/models"
)

func (r *Repository) LogUserRegistrationEvent(event dbmodels.UserRegistrationEvent) error {
	collection := r.dbClient.Database("nsq_test").Collection("user_registration_events")
	_, err := collection.InsertOne(context.Background(), event)
	if err != nil {
		return err
	}

	return nil
}
