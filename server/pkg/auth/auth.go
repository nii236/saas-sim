package auth

import (
	"boilerplate/pkg/db"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ninja-software/e"

	"github.com/jmoiron/sqlx"

	"github.com/go-chi/jwtauth"
)

type Auther struct {
	tokenAuth *jwtauth.JWTAuth
	conn      *sqlx.DB
}

func New(jwtsecret string, conn *sqlx.DB) *Auther {
	result := &Auther{
		tokenAuth: jwtauth.New("HS256", []byte(jwtsecret), []byte(jwtsecret)),
		conn:      conn,
	}
	return result
}
func (a *Auther) FromContext(ctx context.Context) (*db.User, error) {
	op := "auth.FromContext"
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return nil, &e.Error{
			Op:      op,
			Err:     err,
			Message: "Could not find current user",
			Code:    e.EINTERNAL,
		}
	}
	idI, ok := claims.Get("id")
	if !ok {
		err := errors.New("could not get ID from claims")
		return nil, &e.Error{
			Op:      op,
			Err:     err,
			Message: "Could not find current user",
			Code:    e.EINTERNAL,
		}
	}
	id := idI.(string)
	fmt.Println("GOT ID:", id)
	u, err := db.FindUser(a.conn, id)
	if err != nil {
		return nil, &e.Error{
			Op:      op,
			Err:     err,
			Message: "Could not find current user",
			Code:    e.EINTERNAL,
		}
	}

	return u, nil

}
func (a *Auther) GenerateJWT(user, id, role string) (string, error) {
	_, tokenString, err := a.tokenAuth.Encode(jwtauth.Claims{"user": user, "id": id, "role": role})
	fmt.Println("SET ID:", id)
	return tokenString, err
}
func (a *Auther) VerifyMiddleware() func(http.Handler) http.Handler {
	return jwtauth.Verifier(a.tokenAuth)
}
