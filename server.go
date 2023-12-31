package main

import (
	"log"
	"net/http"
	"os"

	"review-pull-request-be/graph"
	"review-pull-request-be/graph/services"
	database "review-pull-request-be/internal/pkg/db/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := database.InitDB()
	defer database.CloseDB(db)
	database.Migrate(db)

	service := services.New(db)

	mux := http.NewServeMux()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Srv: service}}))
	handler := cors.Default().Handler(mux)

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
