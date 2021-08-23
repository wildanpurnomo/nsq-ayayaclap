package gqlschema

import (
	"github.com/graphql-go/graphql"
	gqltypes "github.com/wildanpurnomo/nsq-ayayaclap/main-service/gql/types"
)

func InitGQLSchema() (graphql.Schema, error) {
	schemaConfig := graphql.SchemaConfig{
		Query:    gqltypes.QueryType,
		Mutation: gqltypes.MutationType,
	}
	return graphql.NewSchema(schemaConfig)
}
