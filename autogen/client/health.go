// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "feedpushr": health Resource Client
//
// Command:
// $ goagen
// --design=github.com/ncarlier/feedpushr/design
// --out=$(GOPATH)/src/github.com/ncarlier/feedpushr/autogen
// --version=v1.4.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// GetHealthPath computes a request path to the get action of health.
func GetHealthPath() string {

	return fmt.Sprintf("/v1/healthz")
}

// Perform health check.
func (c *Client) GetHealth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetHealthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetHealthRequest create the request corresponding to the get action endpoint of the health resource.
func (c *Client) NewGetHealthRequest(ctx context.Context, path string) (*http.Request, error) {
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
