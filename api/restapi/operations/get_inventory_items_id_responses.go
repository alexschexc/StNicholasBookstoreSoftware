// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/alexschexc/bookstore/api/models"
)

// GetInventoryItemsIDOKCode is the HTTP code returned for type GetInventoryItemsIDOK
const GetInventoryItemsIDOKCode int = 200

/*
GetInventoryItemsIDOK Successful operation

swagger:response getInventoryItemsIdOK
*/
type GetInventoryItemsIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.InventoryItem `json:"body,omitempty"`
}

// NewGetInventoryItemsIDOK creates GetInventoryItemsIDOK with default headers values
func NewGetInventoryItemsIDOK() *GetInventoryItemsIDOK {

	return &GetInventoryItemsIDOK{}
}

// WithPayload adds the payload to the get inventory items Id o k response
func (o *GetInventoryItemsIDOK) WithPayload(payload *models.InventoryItem) *GetInventoryItemsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get inventory items Id o k response
func (o *GetInventoryItemsIDOK) SetPayload(payload *models.InventoryItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInventoryItemsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInventoryItemsIDUnauthorizedCode is the HTTP code returned for type GetInventoryItemsIDUnauthorized
const GetInventoryItemsIDUnauthorizedCode int = 401

/*
GetInventoryItemsIDUnauthorized Autentication information is missing or invalid

swagger:response getInventoryItemsIdUnauthorized
*/
type GetInventoryItemsIDUnauthorized struct {
	/*

	 */
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewGetInventoryItemsIDUnauthorized creates GetInventoryItemsIDUnauthorized with default headers values
func NewGetInventoryItemsIDUnauthorized() *GetInventoryItemsIDUnauthorized {

	return &GetInventoryItemsIDUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the get inventory items Id unauthorized response
func (o *GetInventoryItemsIDUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *GetInventoryItemsIDUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the get inventory items Id unauthorized response
func (o *GetInventoryItemsIDUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *GetInventoryItemsIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
