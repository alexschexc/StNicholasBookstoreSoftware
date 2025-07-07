// Package migrations is responsible for DB schema.
// It relies on https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md
// to manage schema migrations as embedded SQL scripts.
//
// Migration scripts have naming comvention: <migration number>_<migration name>.(up|down).sql,
// where migrations are applied sequentially ordered by migration number. Up/Down migrations
// must be symmetric to bring DB schema state to the expected state.
//
// To add new migrations, install CLI tool:
// go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
//
// and then run:
// migrate create -ext sql -dir api/db/migrations -seq "<migration name>"
// from the root of the repo to create migration Up/Down scripts.
package migrations

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"

	// must be explicitly imported to support SQLite3 driver.
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed *.sql
var migrationsFS embed.FS

// Up performs DB migrations, will result in no-op if no migrations pending.
func Up(dbURI string) error {
	migrations, _, err := migrations(dbURI)
	if err != nil {
		return fmt.Errorf("failed to setup migrations: %w", err)
	}
	defer migrations.Close()

	err = migrations.Up()
	if err != nil && err.Error() != "no change" {
		return fmt.Errorf("failed to execute migrations: %w", err)
	}

	return nil
}

// TODO: implement Down()
func Down(dbURI string, ver int) error {
	panic("not implemented")
}

func migrations(
	dbURI string,
) (*migrate.Migrate, bool, error) {
	dbDriver, err := iofs.New(migrationsFS, ".")
	if err != nil {
		return nil, false, fmt.Errorf("failed to load migrations: %w", err)
	}
	defer dbDriver.Close()

	migrations, err := migrate.NewWithSourceInstance("iofs", dbDriver, dbURI)
	if err != nil {
		return nil, false, fmt.Errorf("failed to initiate migrations: %w", err)
	}

	curr, dirty, err := migrations.Version()
	if dirty {
		return nil, false, errors.New("previous migration failed. manual resulution is required.")
	}

	if err != nil {
		if !isNoMigrations(err) {
			return nil, false, fmt.Errorf("failed to initiate migrations: %w", err)
		}

		return migrations, true, nil
	}

	ver, err := dbDriver.First()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return migrations, false, nil
		}

		return nil, false, fmt.Errorf("failed to initiate migrations: %w", err)
	}

	return migrations, ver > curr, nil
}

// Pending returns whether there are pending migrations
// that have not been applied to the DB.
func Pending(dbURI string) (bool, error) {
	migrations, hasMigrations, err := migrations(dbURI)
	if err != nil {
		return false, fmt.Errorf("failed to setup migrations: %w", err)
	}
	defer migrations.Close()

	return hasMigrations, nil
}

func isNoMigrations(err error) bool {
	return err != nil && strings.Contains(err.Error(), "no migration")
}
