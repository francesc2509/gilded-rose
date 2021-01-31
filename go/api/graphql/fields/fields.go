package fields

import "github.com/graphql-go/graphql"

func HandleQuery() graphql.Fields {
	fieldMap := make(graphql.Fields)

	HandleItemQuery(fieldMap)

	return fieldMap
}
