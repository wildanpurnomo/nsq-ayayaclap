package gqlresolvers

import (
	"github.com/graphql-go/graphql"
	appcontrollers "github.com/wildanpurnomo/nsq-ayayaclap/main-service/controllers"
)

var (
	RegisterResolver = func(params graphql.ResolveParams) (interface{}, error) {
		username := params.Args["username"].(string)
		email := params.Args["email"].(string)
		password := params.Args["password"].(string)

		newUser, err := appcontrollers.RegisterNewUser(username, email, password)
		if err != nil {
			return nil, err
		}

		return newUser, nil
	}
)
