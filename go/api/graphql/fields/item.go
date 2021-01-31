package fields

import (
	"github.com/francesc2509/glided-rose/api/graphql/types"
	"github.com/francesc2509/glided-rose/api/services"
	"github.com/graphql-go/graphql"
)

func HandleItemQuery(fields graphql.Fields) {
	fields["get_item"] = get()
}

func HandleItemMutation(fields graphql.Fields) {
	fields["update_item_inventory"] = updateInventory()
}

func get() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.ItemType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return services.Item.Get()
		},
	}
}

func updateInventory() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.ItemType),
		Args: graphql.FieldConfigArgument{
			"days": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			days := params.Args["days"].(int)

			return services.Item.UpdateInventory(days)
		},
	}
}
