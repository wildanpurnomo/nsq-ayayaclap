package gqltypes

import "github.com/graphql-go/graphql"

var (
	UserType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UserType",
			Fields: graphql.Fields{
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"is_email_verified": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		},
	)
)
