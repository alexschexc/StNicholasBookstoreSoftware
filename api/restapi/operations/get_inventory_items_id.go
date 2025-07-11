// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/alexschexc/bookstore/api/models"
)

// GetInventoryItemsIDHandlerFunc turns a function with the right signature into a get inventory items ID handler
type GetInventoryItemsIDHandlerFunc func(GetInventoryItemsIDParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetInventoryItemsIDHandlerFunc) Handle(params GetInventoryItemsIDParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetInventoryItemsIDHandler interface for that can handle valid get inventory items ID params
type GetInventoryItemsIDHandler interface {
	Handle(GetInventoryItemsIDParams, *models.Principal) middleware.Responder
}

// NewGetInventoryItemsID creates a new http.Handler for the get inventory items ID operation
func NewGetInventoryItemsID(ctx *middleware.Context, handler GetInventoryItemsIDHandler) *GetInventoryItemsID {
	return &GetInventoryItemsID{Context: ctx, Handler: handler}
}

/*
	GetInventoryItemsID swagger:route GET /inventory-items/{id} getInventoryItemsId

Get an inventory by ID
*/
type GetInventoryItemsID struct {
	Context *middleware.Context
	Handler GetInventoryItemsIDHandler
}

func (o *GetInventoryItemsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetInventoryItemsIDParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
