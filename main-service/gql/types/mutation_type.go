package gqltypes

import (
	"github.com/graphql-go/graphql"
	gqlresolvers "github.com/wildanpurnomo/nsq-ayayaclap/main-service/gql/resolvers"
)

var (
	MutationType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "MutationType",
			Fields: graphql.Fields{
				"register": &graphql.Field{
					Type:        UserType,
					Description: "Register new user",
					Args: graphql.FieldConfigArgument{
						"username": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"email": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: gqlresolvers.RegisterResolver,
				},
				"login": &graphql.Field{
					Type:        UserType,
					Description: "Log a user in",
					Args: graphql.FieldConfigArgument{
						"identifier": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: gqlresolvers.LoginResolver,
				},
			},
		},
	)
)
