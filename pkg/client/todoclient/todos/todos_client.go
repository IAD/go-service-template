// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new todos API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for todos API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddOne add one API
*/
func (a *Client) AddOne(params *AddOneParams) (*AddOneCreated, *AddOneInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addOne",
		Method:             "POST",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *AddOneCreated:
		return v, nil, nil
	case *AddOneInternalServerError:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
DestroyOne destroy one API
*/
func (a *Client) DestroyOne(params *DestroyOneParams) (*DestroyOneNoContent, *DestroyOneNotFound, *DestroyOneInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDestroyOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "destroyOne",
		Method:             "DELETE",
		PathPattern:        "/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DestroyOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DestroyOneNoContent:
		return v, nil, nil, nil
	case *DestroyOneNotFound:
		return nil, v, nil, nil
	case *DestroyOneInternalServerError:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Find find API
*/
func (a *Client) Find(params *FindParams) (*FindOK, *FindNotFound, *FindInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "find",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *FindOK:
		return v, nil, nil, nil
	case *FindNotFound:
		return nil, v, nil, nil
	case *FindInternalServerError:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
UpdateOne update one API
*/
func (a *Client) UpdateOne(params *UpdateOneParams) (*UpdateOneOK, *UpdateOneNotFound, *UpdateOneInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOne",
		Method:             "PUT",
		PathPattern:        "/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateOneOK:
		return v, nil, nil, nil
	case *UpdateOneNotFound:
		return nil, v, nil, nil
	case *UpdateOneInternalServerError:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}