package dbmodels

import (
	"github.com/google/uuid"
)

type InventoryItem struct {
	ID       uuid.UUID `db:"id"`
	Category string    `db:"category"`
	Name     string    `db:"name"`
}
