package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/hunter1271/todos/database"
	"github.com/hunter1271/todos/graphql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/cors"
	"log"
	"net/http"
)

const defaultPort = "8888"

func main() {
	db, err := sql.Open("pgx", "postgres://todos:123456@localhost:25432/todos?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	queries := database.New(db)

	router := chi.NewRouter()
	opts := cors.Options{
		AllowedOrigins: []string{"*"},
		Debug:          true,
	}
	router.Use(cors.New(opts).Handler)
	router.Handle("/", handler.Playground("Todo list API", "/query"))
	router.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: graphql.NewResolver(queries)})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	err = http.ListenAndServe(":"+defaultPort, router)
	if err != nil {
		log.Panic(err)
	}
}
