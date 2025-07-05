package inventory

import (
	"net/http"

	"github.com/alexschexc/bookstore/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// TODO: follow tutorial to implement: https://goswagger.io/go-swagger/generate/server/

func NewGetInventory(ctx *middleware.Context, handler operations.GetInventoryHandler) *GetInventory {
	return &GetInventory{Context: ctx, Handler: handler}
}

type GetInventory struct {
	Context *middleware.Context
	Params  operations.GetInventoryParams
	Handler operations.GetInventoryHandler
}

func (o *GetInventory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// TODO: sort out all the details in here, this is a dirty first pass only
	route, _, _ := o.Context.RouteInfo(r)
	_, _, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)
}
