package main

import (
	"database/sql"
	"github.com/hunter1271/todos/internal/database"
	"github.com/hunter1271/todos/internal/graphql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8888"

func main() {
	db, err := sql.Open("pgx", "postgres://todos:123456@localhost:25432/todos?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	queries := database.New(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: graphql.NewResolver(queries)})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
