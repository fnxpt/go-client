// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the products client
type API interface {
	/*
	   DeleteChartrepoRepoChartsNameVersionLabelsID removes label from chart

	   Remove label from the specified chart version.*/
	DeleteChartrepoRepoChartsNameVersionLabelsID(ctx context.Context, params *DeleteChartrepoRepoChartsNameVersionLabelsIDParams) (*DeleteChartrepoRepoChartsNameVersionLabelsIDOK, error)
	/*
	   GetChartrepoRepoChartsNameVersionLabels returns the attahced labels of chart

	   Return the attahced labels of the specified chart version.*/
	GetChartrepoRepoChartsNameVersionLabels(ctx context.Context, params *GetChartrepoRepoChartsNameVersionLabelsParams) (*GetChartrepoRepoChartsNameVersionLabelsOK, error)
	/*
	   PostChartrepoRepoChartsNameVersionLabels marks label to chart

	   Mark label to the specified chart version.*/
	PostChartrepoRepoChartsNameVersionLabels(ctx context.Context, params *PostChartrepoRepoChartsNameVersionLabelsParams) (*PostChartrepoRepoChartsNameVersionLabelsOK, error)
	/*
	   PostEmailPing tests connection and authentication with email server

	   Test connection and authentication with email server.
	*/
	PostEmailPing(ctx context.Context, params *PostEmailPingParams) (*PostEmailPingOK, error)
}

// New creates a new products API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for products API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteChartrepoRepoChartsNameVersionLabelsID removes label from chart

Remove label from the specified chart version.
*/
func (a *Client) DeleteChartrepoRepoChartsNameVersionLabelsID(ctx context.Context, params *DeleteChartrepoRepoChartsNameVersionLabelsIDParams) (*DeleteChartrepoRepoChartsNameVersionLabelsIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteChartrepoRepoChartsNameVersionLabelsID",
		Method:             "DELETE",
		PathPattern:        "/chartrepo/{repo}/charts/{name}/{version}/labels/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteChartrepoRepoChartsNameVersionLabelsIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteChartrepoRepoChartsNameVersionLabelsIDOK), nil

}

/*
GetChartrepoRepoChartsNameVersionLabels returns the attahced labels of chart

Return the attahced labels of the specified chart version.
*/
func (a *Client) GetChartrepoRepoChartsNameVersionLabels(ctx context.Context, params *GetChartrepoRepoChartsNameVersionLabelsParams) (*GetChartrepoRepoChartsNameVersionLabelsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetChartrepoRepoChartsNameVersionLabels",
		Method:             "GET",
		PathPattern:        "/chartrepo/{repo}/charts/{name}/{version}/labels",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetChartrepoRepoChartsNameVersionLabelsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChartrepoRepoChartsNameVersionLabelsOK), nil

}

/*
PostChartrepoRepoChartsNameVersionLabels marks label to chart

Mark label to the specified chart version.
*/
func (a *Client) PostChartrepoRepoChartsNameVersionLabels(ctx context.Context, params *PostChartrepoRepoChartsNameVersionLabelsParams) (*PostChartrepoRepoChartsNameVersionLabelsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostChartrepoRepoChartsNameVersionLabels",
		Method:             "POST",
		PathPattern:        "/chartrepo/{repo}/charts/{name}/{version}/labels",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostChartrepoRepoChartsNameVersionLabelsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostChartrepoRepoChartsNameVersionLabelsOK), nil

}

/*
PostEmailPing tests connection and authentication with email server

Test connection and authentication with email server.

*/
func (a *Client) PostEmailPing(ctx context.Context, params *PostEmailPingParams) (*PostEmailPingOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostEmailPing",
		Method:             "POST",
		PathPattern:        "/email/ping",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostEmailPingReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostEmailPingOK), nil

}
