// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRunArtifactReader is a Reader for the GetRunArtifact structure.
type GetRunArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetRunArtifactNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunArtifactOK creates a GetRunArtifactOK with default headers values
func NewGetRunArtifactOK() *GetRunArtifactOK {
	return &GetRunArtifactOK{}
}

/*GetRunArtifactOK handles this case with default header values.

A successful response.
*/
type GetRunArtifactOK struct {
	Payload string
}

func (o *GetRunArtifactOK) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactOK  %+v", 200, o.Payload)
}

func (o *GetRunArtifactOK) GetPayload() string {
	return o.Payload
}

func (o *GetRunArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactNoContent creates a GetRunArtifactNoContent with default headers values
func NewGetRunArtifactNoContent() *GetRunArtifactNoContent {
	return &GetRunArtifactNoContent{}
}

/*GetRunArtifactNoContent handles this case with default header values.

No content.
*/
type GetRunArtifactNoContent struct {
	Payload interface{}
}

func (o *GetRunArtifactNoContent) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactNoContent  %+v", 204, o.Payload)
}

func (o *GetRunArtifactNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactForbidden creates a GetRunArtifactForbidden with default headers values
func NewGetRunArtifactForbidden() *GetRunArtifactForbidden {
	return &GetRunArtifactForbidden{}
}

/*GetRunArtifactForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetRunArtifactForbidden struct {
	Payload interface{}
}

func (o *GetRunArtifactForbidden) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactForbidden  %+v", 403, o.Payload)
}

func (o *GetRunArtifactForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactNotFound creates a GetRunArtifactNotFound with default headers values
func NewGetRunArtifactNotFound() *GetRunArtifactNotFound {
	return &GetRunArtifactNotFound{}
}

/*GetRunArtifactNotFound handles this case with default header values.

Resource does not exist.
*/
type GetRunArtifactNotFound struct {
	Payload interface{}
}

func (o *GetRunArtifactNotFound) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactNotFound  %+v", 404, o.Payload)
}

func (o *GetRunArtifactNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
