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

// UpdateOneReader is a Reader for the UpdateOne structure.
type UpdateOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateOneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateOneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /{id} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateOneOK creates a UpdateOneOK with default headers values
func NewUpdateOneOK() *UpdateOneOK {
	return &UpdateOneOK{}
}

/*UpdateOneOK handles this case with default header values.

OK
*/
type UpdateOneOK struct {
	Payload *todoclientmodels.Item
}

func (o *UpdateOneOK) Error() string {
	return fmt.Sprintf("[PUT /{id}][%d] updateOneOK  %+v", 200, o.Payload)
}

func (o *UpdateOneOK) GetPayload() *todoclientmodels.Item {
	return o.Payload
}

func (o *UpdateOneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(todoclientmodels.Item)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOneNotFound creates a UpdateOneNotFound with default headers values
func NewUpdateOneNotFound() *UpdateOneNotFound {
	return &UpdateOneNotFound{}
}

/*UpdateOneNotFound handles this case with default header values.

Not Found
*/
type UpdateOneNotFound struct {
	Payload *todoclientmodels.Error
}

func (o *UpdateOneNotFound) Plain() (code string, message string, attributes map[string]string) {
	return o.Payload.Code, o.Payload.Message, o.Payload.Attributes
}

func (o *UpdateOneNotFound) Error() string {
	return fmt.Sprintf("[PUT /{id}][%d] updateOneNotFound  %+v", 404, o.Payload)
}

func (o *UpdateOneNotFound) GetPayload() *todoclientmodels.Error {
	return o.Payload
}

func (o *UpdateOneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(todoclientmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOneInternalServerError creates a UpdateOneInternalServerError with default headers values
func NewUpdateOneInternalServerError() *UpdateOneInternalServerError {
	return &UpdateOneInternalServerError{}
}

/*UpdateOneInternalServerError handles this case with default header values.

Internal Error
*/
type UpdateOneInternalServerError struct {
	Payload *todoclientmodels.Error
}

func (o *UpdateOneInternalServerError) Plain() (code string, message string, attributes map[string]string) {
	return o.Payload.Code, o.Payload.Message, o.Payload.Attributes
}

func (o *UpdateOneInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /{id}][%d] updateOneInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateOneInternalServerError) GetPayload() *todoclientmodels.Error {
	return o.Payload
}

func (o *UpdateOneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(todoclientmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}