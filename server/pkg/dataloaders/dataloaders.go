//go:generate go run github.com/vektah/dataloaden UserLoader string *boilerplate/pkg/db.User
//go:generate go run github.com/vektah/dataloaden TodoLoader string *boilerplate/pkg/db.Todo
//go:generate go run github.com/vektah/dataloaden UserTodoLoader string []*boilerplate/pkg/db.Todo

package dataloaders

import (
	"boilerplate/pkg/db"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/volatiletech/sqlboiler/queries/qm"
)

type contextKey string

const loaderKey contextKey = "dataloaders"

// FromContext gets the loaders loaded early by Middleware
func FromContext(ctx context.Context) (*Loaders, bool) {
	loaders, ok := ctx.Value(loaderKey).(*Loaders)
	return loaders, ok
}

// Middleware to load new dataloader instance per request
func Middleware(conn *sqlx.DB) func(http.Handler) http.Handler {
	fn := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			loaders := newLoaders(conn)
			ctx := context.WithValue(r.Context(), loaderKey, loaders)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	return fn
}

// newLoaders will initialise the loaders needed for the resolvers
func newLoaders(conn *sqlx.DB) *Loaders {
	todoLoader := NewTodoLoader(TodoLoaderConfig{
		Wait:     1 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(keys []string) ([]*db.Todo, []error) {
			op := "todoLoader"
			fmt.Println("RUN:", op)
			args := []interface{}{}
			for _, key := range keys {
				args = append(args, key)
			}
			records, err := db.Todos(qm.WhereIn("id in ?", args...)).All(conn)
			if err != nil {
				return nil, []error{err}
			}

			result := []*db.Todo{}
			for _, key := range keys {
				for _, record := range records {
					if record.ID == key {
						result = append(result, record)
					}
				}
			}
			return result, nil
		},
	})
	userLoader := NewUserLoader(UserLoaderConfig{
		Wait:     1 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(keys []string) ([]*db.User, []error) {
			op := "userLoader"
			fmt.Println("RUN:", op)
			args := []interface{}{}
			for _, key := range keys {

				args = append(args, key)
			}
			records, err := db.Users(qm.WhereIn("id in ?", args...)).All(conn)
			if err != nil {
				return nil, []error{err}
			}

			result := []*db.User{}
			for _, key := range keys {
				for _, record := range records {
					if record.ID == key {
						result = append(result, record)
					}
				}
			}
			return result, nil
		},
	})
	userTodoLoader := NewUserTodoLoader(UserTodoLoaderConfig{
		Wait:     1 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(keys []string) ([][]*db.Todo, []error) {
			op := "userTodoLoader"
			fmt.Println("RUN:", op)
			var errs []error
			sliceResult := [][]*db.Todo{}
			for _, key := range keys {
				records, err := db.Todos(db.TodoWhere.UserID.EQ(key)).All(conn)
				if err != nil {
					errs = append(errs, err)
					continue
				}
				sliceResult = append(sliceResult, records)

			}
			if len(errs) > 0 {
				return nil, errs
			}
			return sliceResult, nil
		},
	})

	return &Loaders{
		todoLoader,
		userLoader,
		userTodoLoader,
	}
}

// Loaders holds all the loaders
type Loaders struct {
	*TodoLoader
	*UserLoader
	*UserTodoLoader
}
