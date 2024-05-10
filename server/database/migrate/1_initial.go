package migrate

import (
	"fmt"

	"github.com/go-pg/migrations"
)

const recipeTable = `
CREATE TABLE recipes (
    id serial NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp with time zone DEFAULT current_timestamp,
    first_ingredient text NOT NULL,
    second_ingredient text NOT NULL,
    result text NOT NULL,
	emoji text NOT NULL,
    PRIMARY KEY (id)
)`

func init() {
	up := []string{
		recipeTable,
	}

	down := []string{
		`DROP TABLE recipes`,
	}

	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating initial tables")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("dropping initial tables")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
