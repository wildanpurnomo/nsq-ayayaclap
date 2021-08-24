package appcontrollers

import (
	"errors"
	"log"
	"strings"

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

	event := nsqutil.Event{
		EventName: "confirm_new_user",
		Data:      email,
	}
	nsqutil.NsqPublisher.Publish("confirm_new_user", event)

	return nil
}

func Login(identifier string, password string) (models.User, error) {
	var user models.User
	if err := repositories.Repo.GetVerifiedUser(identifier, &user); err != nil {
		log.Println(err.Error())
		return models.User{}, errors.New("Invalid credentials or unverified")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.User{}, errors.New("Invalid credentials or unverified")
	}

	event := nsqutil.Event{
		EventName: "user_login",
		Data:      user.Email,
	}
	nsqutil.NsqPublisher.Publish("user_login", event)

	user.Password = ""
	return user, nil
}

func Register(username string, email string, password string) (models.User, error) {
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
	input.Username = strings.TrimSpace(username)
	input.Email = strings.TrimSpace(email)

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
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
	nsqutil.NsqPublisher.Publish("register_new_user", event)

	return input, nil
}
