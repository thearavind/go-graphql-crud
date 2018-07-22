package main

import (
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/thearavind/goGraphql/resolvers"
	"github.com/thearavind/goGraphql/db"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", graphqlPlaygroundHandler).Methods("GET")
	schemaString, err := parseSchema("./schema.graphql")
	if err != nil {
		panic(err)
	}
	schema := graphql.MustParseSchema(schemaString, &resolvers.Resolver{DB: db.ConnectToDb()})
	router.Handle("/graphql", &relay.Handler{Schema: schema})
	http.ListenAndServe(":8080", router)
}

func parseSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
