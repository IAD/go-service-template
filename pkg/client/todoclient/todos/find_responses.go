// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	todoclientmodels "github.com/IAD/go-swagger-template-example/pkg/client/todoclientmodels"
)

// FindReader is a Reader for the Find structure.
type FindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewFindNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewFindInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET / returns an error %d: %s", response.Code(), string(data))
	}
}

// NewFindOK creates a FindOK with default headers values
func NewFindOK() *FindOK {
	return &FindOK{}
}

/*FindOK handles this case with default header values.

OK
*/
type FindOK struct {
	Payload []*todoclientmodels.Item
}

func (o *FindOK) Error() string {
	return fmt.Sprintf("[GET /][%d] findOK  %+v", 200, o.Payload)
}

func (o *FindOK) GetPayload() []*todoclientmodels.Item {
	return o.Payload
}

func (o *FindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindNotFound creates a FindNotFound with default headers values
func NewFindNotFound() *FindNotFound {
	return &FindNotFound{}
}

/*FindNotFound handles this case with default header values.

Not Found
*/
type FindNotFound struct {
	Payload *todoclientmodels.Error
}

func (o *FindNotFound) Plain() (code string, message string, attributes map[string]string) {
	return o.Payload.Code, o.Payload.Message, o.Payload.Attributes
}

func (o *FindNotFound) Error() string {
	return fmt.Sprintf("[GET /][%d] findNotFound  %+v", 404, o.Payload)
}

func (o *FindNotFound) GetPayload() *todoclientmodels.Error {
	return o.Payload
}

func (o *FindNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(todoclientmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindInternalServerError creates a FindInternalServerError with default headers values
func NewFindInternalServerError() *FindInternalServerError {
	return &FindInternalServerError{}
}

/*FindInternalServerError handles this case with default header values.

Internal Error
*/
type FindInternalServerError struct {
	Payload *todoclientmodels.Error
}

func (o *FindInternalServerError) Plain() (code string, message string, attributes map[string]string) {
	return o.Payload.Code, o.Payload.Message, o.Payload.Attributes
}

func (o *FindInternalServerError) Error() string {
	return fmt.Sprintf("[GET /][%d] findInternalServerError  %+v", 500, o.Payload)
}

func (o *FindInternalServerError) GetPayload() *todoclientmodels.Error {
	return o.Payload
}

func (o *FindInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(todoclientmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
