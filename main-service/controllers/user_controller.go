package appcontrollers

import (
	"errors"

	"github.com/wildanpurnomo/nsq-ayayaclap/main-service/db/models"
	"github.com/wildanpurnomo/nsq-ayayaclap/main-service/db/repositories"
	"github.com/wildanpurnomo/nsq-ayayaclap/main-service/libs"
	nsqutil "github.com/wildanpurnomo/nsq-ayayaclap/main-service/nsq"
	"golang.org/x/crypto/bcrypt"
)

func ConfirmNewUser(email string) error {
	var user models.User
	if err := repositories.Repo.GetUnverifiedUser(email, &user); err != nil {
		return err
	}

	if err := repositories.Repo.ConfirmUserRegistration(user.Email); err != nil {
		return err
	}

	return nil
}

func RegisterNewUser(username string, email string, password string) (models.User, error) {
	if !libs.ValidateUsername(username) {
		return models.User{}, errors.New("Invalid username")
	}

	if !libs.ValidateEmail(email) {
		return models.User{}, errors.New("Invalid email")
	}

	if !libs.ValidatePassword(password) {
		return models.User{}, errors.New("Invalid password")
	}

	var input models.User
	input.Username = username
	input.Email = email

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return models.User{}, errors.New("monkaW")
	}
	input.Password = string(hashed)

	if err := repositories.Repo.CreateNewUser(&input); err != nil {
		return models.User{}, err
	}

	event := nsqutil.Event{
		EventName: "register_new_user",
		Data:      input.Email,
	}
	nsqutil.NsqPublisher.Publish("test", event)

	return input, nil
}
