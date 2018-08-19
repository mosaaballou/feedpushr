// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "feedpushr": swagger Resource Client
//
// Command:
// $ goagen
// --design=github.com/ncarlier/feedpushr/design
// --out=$(GOPATH)/src/github.com/ncarlier/feedpushr/autogen
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// GetSwaggerPath computes a request path to the get action of swagger.
func GetSwaggerPath() string {

	return fmt.Sprintf("/v1/swagger.json")
}

// Get OpenAPI specifications
func (c *Client) GetSwagger(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetSwaggerRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetSwaggerRequest create the request corresponding to the get action endpoint of the swagger resource.
func (c *Client) NewGetSwaggerRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
