package api

import (
	"boilerplate/pkg/dataloaders"
	"boilerplate/pkg/graphql"
	"context"
	"net/http"

	gql "github.com/99designs/gqlgen/graphql"

	"go.uber.org/zap"

	"github.com/go-chi/chi/middleware"

	"github.com/jmoiron/sqlx"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
)

func Routes(conn *sqlx.DB, log *zap.SugaredLogger) http.Handler {
	// auther := auth.New("secret", conn)
	dataloaderMiddleware := dataloaders.Middleware(conn)
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(dataloaderMiddleware)
	r.Route("/api/gql", func(r chi.Router) {
		// r.Use(auther.VerifyMiddleware())
		r.Handle("/", handler.Playground("GraphQL playground", "/api/gql/query"))
		r.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{
			Resolvers: &graphql.Resolver{
				Conn: conn,
			},
			Directives: graphql.DirectiveRoot{
				HasRole: func(ctx context.Context, obj interface{}, next gql.Resolver, role graphql.Role) (res interface{}, err error) {
					return next(ctx)
				},
			},
		})))
	})

	return r
}
