// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"

	models "github.com/IAD/go-swagger-template-example/internal/server/models"
)

// AddOneCreatedCode is the HTTP code returned for type AddOneCreated
const AddOneCreatedCode int = 201

/*AddOneCreated Created

swagger:response addOneCreated
*/
type AddOneCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewAddOneCreatedFunc is declaration for func that create response
type NewAddOneCreatedFunc func() *AddOneCreated

// NewAddOneCreated creates AddOneCreated with default headers values
func NewAddOneCreated() *AddOneCreated {

	return &AddOneCreated{}
}

// WithPayload adds the payload to the add one created response
func (o *AddOneCreated) WithPayload(payload *models.Item) *AddOneCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add one created response
func (o *AddOneCreated) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddOneCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			logrus.Panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddOneInternalServerErrorCode is the HTTP code returned for type AddOneInternalServerError
const AddOneInternalServerErrorCode int = 500

/*AddOneInternalServerError Internal Error

swagger:response addOneInternalServerError
*/
type AddOneInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddOneInternalServerErrorFunc is declaration for func that create response
type NewAddOneInternalServerErrorFunc func() *AddOneInternalServerError

// NewAddOneInternalServerError creates AddOneInternalServerError with default headers values
func NewAddOneInternalServerError() *AddOneInternalServerError {

	return &AddOneInternalServerError{}
}

// WithPayload adds the payload to the add one internal server error response
func (o *AddOneInternalServerError) WithPayload(payload *models.Error) *AddOneInternalServerError {
	o.Payload = payload
	return o
}

// WithErr adds the Error payload with a default code to the add one internal server error response
func (o *AddOneInternalServerError) WithErr(err error) *AddOneInternalServerError {
	type swaggerErr interface {
		Plain() (code string, message string, attributes map[string]string)
	}

	if swaggerErr, ok := err.(swaggerErr); ok {
		code, message, attributes := swaggerErr.Plain()

		o.Payload = &models.Error{
			Code:       code,
			Message:    message,
			Attributes: attributes,
		}
		return o
	}

	return o.WithError("500", err.Error())
}

// WithError  adds the Error payload to the add one internal server error response
func (o *AddOneInternalServerError) WithError(code string, message string) *AddOneInternalServerError {
	o.Payload = &models.Error{
		Attributes: nil,
		Code:       code,
		Message:    message,
	}
	return o
}

// SetPayload sets the payload to the add one internal server error response
func (o *AddOneInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddOneInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			logrus.Panic(err) // let the recovery middleware deal with this
		}
	}
}
