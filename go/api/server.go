package api

import (
	"net/http"

	"github.com/francesc2509/glided-rose/api/graphql"
	servergo "github.com/francesc2509/http-wrapper"
)

func Start() {
	router := servergo.New()

	router.HandleFunc("/graphql", graphql.Init()).Methods("OPTIONS", "POST", "GET", "PUT", "DELETE")

	http.ListenAndServe(":8080", router)
}
