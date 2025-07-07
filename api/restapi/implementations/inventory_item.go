package implementations

import (
	"github.com/alexschexc/bookstore/api/dbmodels"
	"github.com/alexschexc/bookstore/api/models"
	"github.com/alexschexc/bookstore/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/upper/db/v4"
)

type InventoryItemsGetHandler struct {
	Session db.Session
}

func (i *InventoryItemsGetHandler) Handle(
	params operations.GetInventoryItemsParams,
	principal *models.Principal,
) middleware.Responder {
	// first, ensure user is authorized to access this scope
	// TODO: switch to OAuth Scopes once switched
	scope := "inventory:read"
	if err := principal.Scoped(scope); err != nil {
		return operations.NewGetInventoryItemsIDUnauthorized().WithWWWAuthenticate(scope)
	}

	filter := db.Cond{}
	if params.Category != nil {
		filter["category"] = *params.Category
	}
	var items []*dbmodels.InventoryItem

	if err := i.Session.Collection("inventory_items").Find(filter).All(&items); err != nil {
		return operations.NewGetInventoryItemsInternalServerError().WithWWWError(err.Error())
	}

	var r []*models.InventoryItem
	for _, itm := range items {
		r = append(r, &models.InventoryItem{
			ID: strfmt.UUID(itm.ID.String()),
			InventoryFields: models.InventoryFields{
				Name:     itm.Name,
				Category: itm.Category,
			},
		})
	}

	return operations.NewGetInventoryItemsOK().WithPayload(r)
}
