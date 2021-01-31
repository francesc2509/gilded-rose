package types

import "github.com/graphql-go/graphql"

var ItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Item",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Float,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"sellIn": &graphql.Field{
			Type: graphql.Int,
		},
		"quality": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
