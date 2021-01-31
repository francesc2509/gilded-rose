package fields

import (
	"github.com/francesc2509/glided-rose/api/graphql/types"
	"github.com/francesc2509/glided-rose/api/services"
	"github.com/graphql-go/graphql"
)

func HandleItemQuery(fields graphql.Fields) {
	fields["get_item"] = get()
}

func get() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.ItemType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return services.Item.Get()
		},
	}
}
