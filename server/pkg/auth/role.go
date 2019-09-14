package auth

import (
	"context"
	"errors"

	"github.com/go-chi/jwtauth"
	"github.com/ninja-software/e"
)

func RoleFromContext(ctx context.Context) (string, error) {
	op := "roleFromContext"
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return "", &e.Error{
			Err:     err,
			Message: "Could not validate the provided token",
			Op:      op,
			Code:    e.EINVALID,
		}
	}

	roleI, ok := claims.Get("role")
	if !ok {
		return "", &e.Error{
			Err:     errors.New("could not get claim in token"),
			Message: "Could not validate the provided token",
			Op:      op,
			Code:    e.EINVALID,
		}
	}

	jwtRoleStr := roleI.(string)

	return jwtRoleStr, nil
}
