package graphql

import (
	"context"
	"fmt"
	"github.com/hunter1271/todos/internal/database"
	"strconv"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	queries *database.Queries
}

func NewResolver(queries *database.Queries) *Resolver {
	return &Resolver{queries: queries}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*database.Todo, error) {
	todo, err := r.queries.CreateTodo(ctx, input.Text)

	return &todo, err
}

func (r *mutationResolver) SetTodoDone(ctx context.Context, id *string, done *bool) (*database.Todo, error) {
	intId, err := parseID(id)
	if err != nil {
		return nil, err
	}
	todo, err := r.queries.UpdateTodoDone(ctx, database.UpdateTodoDoneParams{ID: intId, Done: *done})

	return &todo, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todo(ctx context.Context, id *string) (*database.Todo, error) {
	intId, err := parseID(id)
	if err != nil {
		return nil, err
	}
	todo, err := r.queries.GetTodo(ctx, intId)

	return &todo, err
}

func (r *queryResolver) Todos(ctx context.Context) ([]*database.Todo, error) {
	todos, err := r.queries.ListTodos(ctx)
	var list []*database.Todo
	for _, todo := range todos {
		list = append(list, &todo)
	}

	return list, err
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) ID(ctx context.Context, obj *database.Todo) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

func parseID(id *string) (int32, error) {
	int64Id, err := strconv.ParseInt(*id, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(int64Id), nil
}