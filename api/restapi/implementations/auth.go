package implementations

import (
	"github.com/alexschexc/bookstore/api/models"
)

func Authenticate(user, password string) (*models.Principal, error) {
	// return nil, errors.New("error authenticate")

	// TODO: read scopes from Users DB once implemented
	return &models.Principal{Scopes: "inventory:read"}, nil
}
