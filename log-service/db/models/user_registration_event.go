package dbmodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRegistrationEvent struct {
	EventName string              `bson:"event_name, omitempty"`
	Email     string              `bson:"email, omitempty"`
	Timestamp primitive.Timestamp `bson:"timestamp, omitempty"`
}
