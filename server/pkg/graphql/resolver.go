package graphql

import (
	"boilerplate/pkg/dataloaders"
	"boilerplate/pkg/db"
	"context"
	"errors"

	"syreclabs.com/go/faker"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	Conn *sqlx.DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) Owner(ctx context.Context, obj *db.Todo) (*db.User, error) {
	loaders, ok := dataloaders.FromContext(ctx)
	if !ok {
		return nil, errors.New("no dataloaders in context")
	}
	user, errs := loaders.UserLoader.Load(obj.UserID)
	if errs != nil {
		return nil, errs
	}
	return user, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) TodosConnection(ctx context.Context, obj *db.User, first int, after *string) (*TodoConnection, error) {
	todos, err := obj.Todos(qm.Limit(first)).All(r.Conn)
	if err != nil {
		return nil, err
	}

	edges := []*TodoEdge{}
	for _, todo := range todos {
		edges = append(edges, &TodoEdge{todo.ID, todo})
	}

	result := &TodoConnection{
		Edges:    edges,
		PageInfo: &PageInfo{},
	}
	return result, nil
}
func (r *userResolver) Todos(ctx context.Context, obj *db.User) ([]*db.Todo, error) {
	loaders, ok := dataloaders.FromContext(ctx)
	if !ok {
		return nil, errors.New("no dataloaders in context")
	}
	todos, errs := loaders.UserTodoLoader.Load(obj.ID)
	if errs != nil {
		return nil, errs
	}
	return todos, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) ToggleTodo(ctx context.Context, todoID string) (*db.Todo, error) {
	loaders, ok := dataloaders.FromContext(ctx)
	if !ok {
		return nil, errors.New("no dataloaders in context")
	}
	todo, err := loaders.TodoLoader.Load(todoID)

	if err != nil {
		return nil, err
	}

	todo.Completed = !todo.Completed
	_, err = todo.Update(r.Conn, boil.Whitelist(db.TodoColumns.Completed))
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (r *mutationResolver) AddTodo(ctx context.Context, userID string) (*db.Todo, error) {
	loaders, ok := dataloaders.FromContext(ctx)
	if !ok {
		return nil, errors.New("no dataloaders in context")
	}
	todo := &db.Todo{
		Body:   faker.Company().Bs(),
		UserID: userID,
	}
	err := todo.Insert(r.Conn, boil.Infer())
	if err != nil {
		return nil, err
	}
	loaders.UserTodoLoader.Clear(userID)
	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Viewer(ctx context.Context) (*Viewer, error) {
	panic("not implemented")
}
func (r *queryResolver) Todos(ctx context.Context) ([]*db.Todo, error) {
	return db.Todos().All(r.Conn)
}
func (r *queryResolver) Users(ctx context.Context) ([]*db.User, error) {
	return db.Users().All(r.Conn)
}
