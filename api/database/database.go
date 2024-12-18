package database

import (
	"database/sql"
	"fmt"
	"walletTask/envs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var DB *sql.DB // variable for the database, through it we will access it

func InitDatabase() error {
	var err error
	env := envs.ServerEnvs

	// set the connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_HOST, env.POSTGRES_PORT, env.POSTGRES_NAME, env.POSTGRES_USE_SSL)

	// connect to the database
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// connection test
	if err = DB.Ping(); err != nil {
		return err
	}

	// create a driver for migration
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		return err
	}

	// download the migration file
	migrateDB, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return err
	}

	// perform migration (the table is created there)
	err = migrateDB.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
