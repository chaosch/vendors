// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/go-swagger/go-swagger/examples/oauth2/models"
)

// GetAuthCallbackOKCode is the HTTP code returned for type GetAuthCallbackOK
const GetAuthCallbackOKCode int = 200

/*GetAuthCallbackOK login

swagger:response getAuthCallbackOK
*/
type GetAuthCallbackOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetAuthCallbackOKBody `json:"body,omitempty"`
}

// NewGetAuthCallbackOK creates GetAuthCallbackOK with default headers values
func NewGetAuthCallbackOK() *GetAuthCallbackOK {

	return &GetAuthCallbackOK{}
}

// WithPayload adds the payload to the get auth callback o k response
func (o *GetAuthCallbackOK) WithPayload(payload *models.GetAuthCallbackOKBody) *GetAuthCallbackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get auth callback o k response
func (o *GetAuthCallbackOK) SetPayload(payload *models.GetAuthCallbackOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthCallbackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAuthCallbackDefault error

swagger:response getAuthCallbackDefault
*/
type GetAuthCallbackDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAuthCallbackDefault creates GetAuthCallbackDefault with default headers values
func NewGetAuthCallbackDefault(code int) *GetAuthCallbackDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAuthCallbackDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get auth callback default response
func (o *GetAuthCallbackDefault) WithStatusCode(code int) *GetAuthCallbackDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get auth callback default response
func (o *GetAuthCallbackDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get auth callback default response
func (o *GetAuthCallbackDefault) WithPayload(payload *models.Error) *GetAuthCallbackDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get auth callback default response
func (o *GetAuthCallbackDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthCallbackDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
