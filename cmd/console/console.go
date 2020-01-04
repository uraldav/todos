package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/hunter1271/todos/internal/database"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func main() {
	db, err := sql.Open("pgx", "postgres://todos:123456@localhost:25432/todos?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	queries := database.New(db)

	todo, err := queries.CreateTodo(context.Background(), "New task")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Create new todo #%d.\n", todo.ID)
}
