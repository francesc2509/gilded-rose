package graphql

import (
	"net/http"

	"github.com/francesc2509/glided-rose/api/graphql/fields"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func Init() http.HandlerFunc {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: fields.HandleQuery(),
	})

	// rootMutation := graphql.NewObject(graphql.ObjectConfig{
	// 	Name:   "Mutation",
	// 	Fields: nil,
	// })

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
		// Mutation: rootMutation,
	})

	if err != nil {
		panic(err)
	}

	handlerFn := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	return func(rw http.ResponseWriter, r *http.Request) {
		handlerFn.ContextHandler(r.Context(), rw, r)
	}
}
