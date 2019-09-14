package main

import (
	"boilerplate/pkg/log"
	"boilerplate/pkg/platform"
	"boilerplate/pkg/seed"
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"github.com/oklog/run"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const version = "v0.0.1"

var app *kingpin.Application
var databaseUser *string
var databasePass *string
var databaseHost *string
var databasePort *string
var databaseName *string

func init() {
	app = kingpin.New("boilerplate", "Wholesome hosting boilerplate")
	app.Version(version)
	databaseUser = app.Flag("database-user", "set the database user").Envar("DATABASE_USER").Default("boilerplate").String()
	databasePass = app.Flag("database-pass", "set the database pass").Envar("DATABASE_PASS").Default("dev").String()
	databaseHost = app.Flag("database-host", "set the database host").Envar("DATABASE_HOST").Default("localhost").String()
	databasePort = app.Flag("database-port", "set the database port").Envar("DATABASE_PORT").Default("5438").String()
	databaseName = app.Flag("database-name", "set the database name").Envar("DATABASE_NAME").Default("boilerplate").String()
	app.Command("serve", "begin webserver").Default()
	app.Command("seed", "seed the database")
}
func main() {
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	conn, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		*databaseUser,
		*databasePass,
		*databaseHost,
		*databasePort,
		*databaseName,
	))
	if err != nil {
		fmt.Println("could not initialise database:", err)
		return
	}

	switch cmd {
	case "seed":
		err = seed.Run(conn)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	default:
		g := &run.Group{}
		ctx, cancel := context.WithCancel(context.Background())
		g.Add(func() error {
			l := log.NewToStdOut("api", version, false)
			server := &platform.API{
				Conn: conn,
				Log:  l,
			}
			return server.Run(ctx)
		}, func(error) {
			cancel()
		})

		err = g.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
