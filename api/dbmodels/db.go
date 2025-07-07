package dbmodels

import (
	"context"
	"fmt"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

func NewDBSession(ctx context.Context, connectionString string) (db.Session, error) {
	settings := sqlite.ConnectionURL{
		Database: connectionString,
	}

	sess, err := sqlite.Open(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %q : %w", connectionString, err)
	}

	return sess, nil
}
