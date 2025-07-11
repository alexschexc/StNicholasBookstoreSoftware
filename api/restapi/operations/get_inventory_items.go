// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/alexschexc/bookstore/api/models"
)

// GetInventoryItemsHandlerFunc turns a function with the right signature into a get inventory items handler
type GetInventoryItemsHandlerFunc func(GetInventoryItemsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetInventoryItemsHandlerFunc) Handle(params GetInventoryItemsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetInventoryItemsHandler interface for that can handle valid get inventory items params
type GetInventoryItemsHandler interface {
	Handle(GetInventoryItemsParams, *models.Principal) middleware.Responder
}

// NewGetInventoryItems creates a new http.Handler for the get inventory items operation
func NewGetInventoryItems(ctx *middleware.Context, handler GetInventoryItemsHandler) *GetInventoryItems {
	return &GetInventoryItems{Context: ctx, Handler: handler}
}

/*
	GetInventoryItems swagger:route GET /inventory-items getInventoryItems

Get all available inventory
*/
type GetInventoryItems struct {
	Context *middleware.Context
	Handler GetInventoryItemsHandler
}

func (o *GetInventoryItems) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetInventoryItemsParams()
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
