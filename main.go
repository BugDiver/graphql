package main

import (
	"encoding/json"
	"net/http"

	"github.com/bugdiver/graphql/schema"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/mnmtanish/go-graphiql"
)

func main() {
	http.HandleFunc("/", graphiql.ServeGraphiQL)
	http.HandleFunc("/graphql", serverGraphQL)
	http.ListenAndServe(":9001", nil)
}

func serverGraphQL(w http.ResponseWriter, r *http.Request) {
	opts := handler.NewRequestOptions(r)

	params := graphql.Params{
		Schema:         schema.Root,
		RequestString:  opts.Query,
		VariableValues: opts.Variables,
		OperationName:  opts.OperationName,
	}
	result := graphql.Do(params)
	// render result

	if err := json.NewEncoder(w).Encode(result); err != nil {
		sendError(w, err)
	}

}

var sendError = func(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}
