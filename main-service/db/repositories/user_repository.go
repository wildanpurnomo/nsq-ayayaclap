package repositories

import "github.com/wildanpurnomo/nsq-ayayaclap/main-service/db/models"

func (r *Repository) CreateNewUser(user *models.User) error {
	stmt, err := r.db.Prepare("INSERT INTO users(username, email, password) VALUES($1, $2, $3)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ConfirmUserRegistration(email string) error {
	stmt, err := r.db.Prepare("UPDATE users SET is_email_verified = TRUE WHERE email = $1 AND is_email_verified = FALSE")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(email)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUnverifiedUser(emailInput string, model *models.User) error {
	stmt, err := r.db.Prepare("SELECT username, email, is_email_verified FROM users WHERE email = $1 AND is_email_verified = FALSE LIMIT 1")
	if err != nil {
		return err
	}

	err = stmt.QueryRow(emailInput).Scan(&model.Username, &model.Email, &model.IsEmailVerified)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetVerifiedUser(input string, model *models.User) error {
	stmt, err := r.db.Prepare("SELECT username, email, password, is_email_verified FROM users WHERE ((username = $1 OR email = $1) AND is_email_verified = TRUE) LIMIT 1")
	if err != nil {
		return err
	}

	err = stmt.QueryRow(input).Scan(&model.Username, &model.Email, &model.Password, &model.IsEmailVerified)
	if err != nil {
		return err
	}

	return nil
}
