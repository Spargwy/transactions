package main

import (
	"log"
	"net/http"
	"os"
	environment "transactions/env"
	"transactions/graph"
	"transactions/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	godotenv.Load("../.env")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	env := environment.New()

	err := env.DBPing()
	if err != nil {
		log.Fatalf("failed DBPing: %v", err)
	}

	log.Print("Connected to db")
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Env: env}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
