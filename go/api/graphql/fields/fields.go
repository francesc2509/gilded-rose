package fields

import "github.com/graphql-go/graphql"

func HandleQuery() graphql.Fields {
	fieldMap := make(graphql.Fields)

	HandleItemQuery(fieldMap)

	return fieldMap
}

func HandleMutation() graphql.Fields {
	fieldMap := make(graphql.Fields)

	HandleItemMutation(fieldMap)

	return fieldMap
}
