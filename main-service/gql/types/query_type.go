package gqltypes

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var (
	QueryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "QueryType",
			Fields: graphql.Fields{
				"echo": &graphql.Field{
					Type:        graphql.String,
					Description: "Dummy cuz this won't work",
					Args: graphql.FieldConfigArgument{
						"message": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						message := fmt.Sprintf("Your message: %v", p.Args["message"].(string))
						return message, nil
					},
				},
			},
		},
	)
)
