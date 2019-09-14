package platform

import (
	"context"
	"net/http"
	"boilerplate/pkg/api"

	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
)

type API struct {
	Conn *sqlx.DB
	Log  *zap.SugaredLogger
}

func (s *API) Run(ctx context.Context) error {
	s.Log.Info("starting server on :8080")
	return http.ListenAndServe(":8080", api.Routes(s.Conn, s.Log))
}
