package migrate

import (
	"log"

	"infinite-craft/database"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

func Migrate(args []string) {
	db, err := database.DBConn()
	if err != nil {
		log.Fatal(err)
	}

	err = db.RunInTransaction(func(tx *pg.Tx) error {
		oldVersion, newVersion, err := migrations.Run(tx, args...)
		if err != nil {
			return err
		}
		if newVersion != oldVersion {
			log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
		} else {
			log.Printf("version is %d\n", oldVersion)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

}

func Reset() {
	db, err := database.DBConn()
	if err != nil {
		log.Fatal(err)
	}

	version, err := migrations.Version(db)
	if err != nil {
		log.Fatal(err)
	}

	err = db.RunInTransaction(func(tx *pg.Tx) error {
		for version != 0 {
			oldVersion, newVersion, err := migrations.Run(tx, "down")
			if err != nil {
				return err
			}
			log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
			version = newVersion
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
