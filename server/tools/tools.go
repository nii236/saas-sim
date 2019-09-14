// +build tools

package tools

//go:generate go build -o ../../bin/caddy github.com/caddyserver/caddy
//go:generate go build -o ../../bin/migrate github.com/golang-migrate/migrate
//go:generate go build -o ../../bin/realize github.com/oxequa/realize
//go:generate go build -o ../../bin/sqlboiler github.com/volatiletech/sqlboiler
//go:generate go build -o ../../bin/sqlboiler-psql github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

import (
	_ "github.com/caddyserver/caddy"
	_ "github.com/golang-migrate/migrate"
	_ "github.com/oxequa/realize"
	_ "github.com/volatiletech/sqlboiler"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql"
)
