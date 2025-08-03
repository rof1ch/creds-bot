package main


import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string
	var isDown bool

	flag.StringVar(&storagePath, "storage-path", "", "storage path")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migration table")
	flag.BoolVar(&isDown, "down", false, "is down migration")
	flag.Parse()

	if storagePath == "" {
		panic("storage-path is required")
	}

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New("file://"+migrationsPath, fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable))

	if err != nil {
		panic(err)
	}

	if isDown {
		m.Down()
	} else {
		err := m.Up()
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		} else if err != nil {
			panic(err)
		}
		return
	}
}
