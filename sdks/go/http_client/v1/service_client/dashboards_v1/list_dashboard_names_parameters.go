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

package dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListDashboardNamesParams creates a new ListDashboardNamesParams object
// with the default values initialized.
func NewListDashboardNamesParams() *ListDashboardNamesParams {
	var ()
	return &ListDashboardNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListDashboardNamesParamsWithTimeout creates a new ListDashboardNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListDashboardNamesParamsWithTimeout(timeout time.Duration) *ListDashboardNamesParams {
	var ()
	return &ListDashboardNamesParams{

		timeout: timeout,
	}
}

// NewListDashboardNamesParamsWithContext creates a new ListDashboardNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListDashboardNamesParamsWithContext(ctx context.Context) *ListDashboardNamesParams {
	var ()
	return &ListDashboardNamesParams{

		Context: ctx,
	}
}

// NewListDashboardNamesParamsWithHTTPClient creates a new ListDashboardNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListDashboardNamesParamsWithHTTPClient(client *http.Client) *ListDashboardNamesParams {
	var ()
	return &ListDashboardNamesParams{
		HTTPClient: client,
	}
}

/*ListDashboardNamesParams contains all the parameters to send to the API endpoint
for the list dashboard names operation typically these are written to a http.Request
*/
type ListDashboardNamesParams struct {

	/*Limit
	  Limit size.

	*/
	Limit *int32
	/*Offset
	  Pagination offset.

	*/
	Offset *int32
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Query
	  Query filter the search search.

	*/
	Query *string
	/*Sort
	  Sort to order the search.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list dashboard names params
func (o *ListDashboardNamesParams) WithTimeout(timeout time.Duration) *ListDashboardNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list dashboard names params
func (o *ListDashboardNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list dashboard names params
func (o *ListDashboardNamesParams) WithContext(ctx context.Context) *ListDashboardNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list dashboard names params
func (o *ListDashboardNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list dashboard names params
func (o *ListDashboardNamesParams) WithHTTPClient(client *http.Client) *ListDashboardNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list dashboard names params
func (o *ListDashboardNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list dashboard names params
func (o *ListDashboardNamesParams) WithLimit(limit *int32) *ListDashboardNamesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list dashboard names params
func (o *ListDashboardNamesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the list dashboard names params
func (o *ListDashboardNamesParams) WithOffset(offset *int32) *ListDashboardNamesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list dashboard names params
func (o *ListDashboardNamesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOwner adds the owner to the list dashboard names params
func (o *ListDashboardNamesParams) WithOwner(owner string) *ListDashboardNamesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the list dashboard names params
func (o *ListDashboardNamesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithQuery adds the query to the list dashboard names params
func (o *ListDashboardNamesParams) WithQuery(query *string) *ListDashboardNamesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the list dashboard names params
func (o *ListDashboardNamesParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the list dashboard names params
func (o *ListDashboardNamesParams) WithSort(sort *string) *ListDashboardNamesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list dashboard names params
func (o *ListDashboardNamesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListDashboardNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}