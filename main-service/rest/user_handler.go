package resthandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appcontrollers "github.com/wildanpurnomo/nsq-ayayaclap/main-service/controllers"
	"github.com/wildanpurnomo/nsq-ayayaclap/main-service/libs"
)

func HandleEmailConfirmation(c *gin.Context) {
	email, err := libs.ExtractEmailFromRedirToken(c.Query("redir_token"))
	if err != nil {
		c.String(http.StatusUnauthorized, "monkaW")
	}

	if err = appcontrollers.ConfirmNewUser(email); err != nil {
		c.String(http.StatusBadRequest, "monkaW")
	}

	c.Redirect(http.StatusFound, "http://localhost:8080/gql")
}
