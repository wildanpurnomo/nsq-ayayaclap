package restcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildanpurnomo/nsq-ayayaclap/main-service/db/models"
	"github.com/wildanpurnomo/nsq-ayayaclap/main-service/db/repositories"
	"github.com/wildanpurnomo/nsq-ayayaclap/main-service/libs"
)

func ConfirmUserEmail(c *gin.Context) {
	email, err := libs.ExtractEmailFromRedirToken(c.Query("redir_token"))
	if err != nil {
		c.String(http.StatusUnauthorized, "monkaW")
	}

	var user models.User
	if err := repositories.Repo.GetUnverifiedUser(email, &user); err != nil {
		c.String(http.StatusUnauthorized, "monkaW")
	}

	if err := repositories.Repo.ConfirmUserRegistration(user.Email); err != nil {
		c.String(http.StatusBadRequest, "monkaW")
	}

	c.Redirect(http.StatusFound, "http://localhost:8080/gql")
}
