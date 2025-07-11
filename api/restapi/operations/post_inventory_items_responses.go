// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/alexschexc/bookstore/api/models"
)

// PostInventoryItemsCreatedCode is the HTTP code returned for type PostInventoryItemsCreated
const PostInventoryItemsCreatedCode int = 201

/*
PostInventoryItemsCreated Successful operation

swagger:response postInventoryItemsCreated
*/
type PostInventoryItemsCreated struct {

	/*
	  In: Body
	*/
	Payload *models.InventoryItem `json:"body,omitempty"`
}

// NewPostInventoryItemsCreated creates PostInventoryItemsCreated with default headers values
func NewPostInventoryItemsCreated() *PostInventoryItemsCreated {

	return &PostInventoryItemsCreated{}
}

// WithPayload adds the payload to the post inventory items created response
func (o *PostInventoryItemsCreated) WithPayload(payload *models.InventoryItem) *PostInventoryItemsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post inventory items created response
func (o *PostInventoryItemsCreated) SetPayload(payload *models.InventoryItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostInventoryItemsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostInventoryItemsUnauthorizedCode is the HTTP code returned for type PostInventoryItemsUnauthorized
const PostInventoryItemsUnauthorizedCode int = 401

/*
PostInventoryItemsUnauthorized Autentication information is missing or invalid

swagger:response postInventoryItemsUnauthorized
*/
type PostInventoryItemsUnauthorized struct {
	/*

	 */
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewPostInventoryItemsUnauthorized creates PostInventoryItemsUnauthorized with default headers values
func NewPostInventoryItemsUnauthorized() *PostInventoryItemsUnauthorized {

	return &PostInventoryItemsUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the post inventory items unauthorized response
func (o *PostInventoryItemsUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *PostInventoryItemsUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the post inventory items unauthorized response
func (o *PostInventoryItemsUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *PostInventoryItemsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
