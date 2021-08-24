package gqlresolvers

import (
	"fmt"

	"github.com/graphql-go/graphql"
	appcontrollers "github.com/wildanpurnomo/nsq-ayayaclap/main-service/controllers"
	"github.com/wildanpurnomo/nsq-ayayaclap/main-service/libs"
)

var (
	RegisterResolver = func(params graphql.ResolveParams) (interface{}, error) {
		username := params.Args["username"].(string)
		email := params.Args["email"].(string)
		password := params.Args["password"].(string)

		newUser, err := appcontrollers.Register(username, email, password)
		if err != nil {
			return nil, err
		}

		return newUser, nil
	}
	LoginResolver = func(params graphql.ResolveParams) (interface{}, error) {
		identifier := params.Args["identifier"].(string)
		password := params.Args["password"].(string)

		user, err := appcontrollers.Login(identifier, password)
		if err != nil {
			return nil, err
		}

		token := libs.GenerateClientAuthToken(fmt.Sprint(user.UserID))

		contextValueWrapper := libs.ExtractContextValueWrapper(params.Context)
		contextValueWrapper.SetJwtCookie(token)

		return user, nil
	}
)
