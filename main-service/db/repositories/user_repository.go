package repositories

import "github.com/wildanpurnomo/nsq-ayayaclap/main-service/db/models"

func (r *Repository) CreateNewUser(user *models.User) error {
	stmt, err := r.db.Prepare("INSERT INTO users(username, email, password) VALUES($1,$2,$3)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
