package seed

import (
	"boilerplate/pkg/db"
	"fmt"

	"syreclabs.com/go/faker"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"
)

func Run(conn *sqlx.DB) error {
	var err error
	fmt.Println("Seeding")
	err = SeedUsers(conn)
	if err != nil {
		return err
	}
	err = SeedTodos(conn)
	if err != nil {
		return err
	}
	fmt.Println("Seed complete")
	return nil
}
func SeedUsers(conn *sqlx.DB) error {
	names := []string{"Corey", "Yong", "Mac", "Ben", "Jett", "Reece", "Natalie", "Alex", "John", "Jamie", "Maricel"}
	for _, name := range names {
		p := &db.User{
			Name: name,
		}
		err := p.Insert(conn, boil.Infer())
		if err != nil {
			return err
		}
	}
	return nil
}

func SeedTodos(conn *sqlx.DB) error {
	users, err := db.Users().All(conn)
	if err != nil {
		return err
	}
	for _, user := range users {
		for i := 0; i < 5; i++ {
			p := &db.Todo{
				ID:        uuid.Must(uuid.NewV4()).String(),
				UserID:    user.ID,
				Body:      faker.Company().Bs(),
				Completed: false,
			}
			err := p.Insert(conn, boil.Infer())
			if err != nil {
				return err
			}

		}
	}
	return nil
}
