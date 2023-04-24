package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/microservice-ec-site/graph/internal/graph"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		log.Println(e)

		// Todo: error handling
		ex := map[string]interface{}{"code": "INTERNAL_SERVER_ERROR", "status": http.StatusInternalServerError}

		return &gqlerror.Error{
			Message:    e.Error(),
			Path:       graphql.GetPath(ctx),
			Extensions: ex,
		}
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
