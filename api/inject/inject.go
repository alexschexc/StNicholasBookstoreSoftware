package inject

import (
	"github.com/alexschexc/bookstore/api/models"
	"github.com/alexschexc/bookstore/api/restapi/implementations"
	"github.com/alexschexc/bookstore/api/restapi/operations"
	"github.com/upper/db/v4"
)

// InjectApi wires up concrete implementations of the operations
func InjectApi(api *operations.BookstoreAPI, sess db.Session) error {
	// Hook up BasicAuth here
	api.BasicAuth = func(user, password string) (*models.Principal, error) {
		// This is a very simple implementataion, we can get fancier
		// by using BearerToken later.
		return implementations.Authenticate(user, password)
	}

	api.GetInventoryItemsHandler = &implementations.InventoryItemsGetHandler{Session: sess}
	// api.GetInventoryItemsIDHandler = nil
	// api.DeleteInventoryItemsIDHandler = nil
	// api.PatchInventoryItemsIDHandler = nil
	// api.PostInventoryItemsHandler = nil

	return nil
}
