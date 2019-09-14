//go:generate ../../../bin/sqlboiler ../../../bin/sqlboiler-psql --wipe --tag db --config ../../sqlboiler.toml --output ../db
//go:generate go run github.com/99designs/gqlgen

package graphql
