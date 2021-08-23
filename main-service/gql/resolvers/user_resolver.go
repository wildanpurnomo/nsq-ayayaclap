package gqlresolvers

import (
	"errors"

	"github.com/graphql-go/graphql"
	models "github.com/wildanpurnomo/nsq-ayayaclap/main-service/db/models"
	"github.com/wildanpurnomo/nsq-ayayaclap/main-service/db/repositories"
	libs "github.com/wildanpurnomo/nsq-ayayaclap/main-service/libs"
	nsqutil "github.com/wildanpurnomo/nsq-ayayaclap/main-service/nsq"
	"golang.org/x/crypto/bcrypt"
)

var (
	RegisterResolver = func(params graphql.ResolveParams) (interface{}, error) {
		username := params.Args["username"].(string)
		email := params.Args["email"].(string)
		password := params.Args["password"].(string)

		if !libs.ValidateUsername(username) {
			return nil, errors.New("Invalid username")
		}

		if !libs.ValidateEmail(email) {
			return nil, errors.New("Invalid email")
		}

		if !libs.ValidatePassword(password) {
			return nil, errors.New("Invalid password")
		}

		var input models.User
		input.Username = username
		input.Email = email

		hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return nil, errors.New("monkaW")
		}
		input.Password = string(hashed)

		if err := repositories.Repo.CreateNewUser(&input); err != nil {
			return nil, err
		}

		event := nsqutil.Event{
			EventName: "register_new_user",
			Data:      input.Email,
		}
		nsqutil.NsqPublisher.Publish("test", event)

		return input, nil
	}
)
