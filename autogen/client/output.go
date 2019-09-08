// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "feedpushr": output Resource Client
//
// Command:
// $ goagen
// --design=github.com/ncarlier/feedpushr/design
// --out=$(GOPATH)/src/github.com/ncarlier/feedpushr/autogen
// --version=v1.4.0

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateOutputPayload is the output create action payload.
type CreateOutputPayload struct {
	// Alias of the output
	Alias string `form:"alias" json:"alias" yaml:"alias" xml:"alias"`
	// Name of the output
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Output properties
	Props map[string]interface{} `form:"props,omitempty" json:"props,omitempty" yaml:"props,omitempty" xml:"props,omitempty"`
	// Comma separated list of tags
	Tags *string `form:"tags,omitempty" json:"tags,omitempty" yaml:"tags,omitempty" xml:"tags,omitempty"`
}

// CreateOutputPath computes a request path to the create action of output.
func CreateOutputPath() string {

	return fmt.Sprintf("/v1/outputs")
}

// Create a new output
func (c *Client) CreateOutput(ctx context.Context, path string, payload *CreateOutputPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateOutputRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateOutputRequest create the request corresponding to the create action endpoint of the output resource.
func (c *Client) NewCreateOutputRequest(ctx context.Context, path string, payload *CreateOutputPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteOutputPath computes a request path to the delete action of output.
func DeleteOutputPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/v1/outputs/%s", param0)
}

// Delete an output
func (c *Client) DeleteOutput(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteOutputRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteOutputRequest create the request corresponding to the delete action endpoint of the output resource.
func (c *Client) NewDeleteOutputRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetOutputPath computes a request path to the get action of output.
func GetOutputPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/v1/outputs/%s", param0)
}

// Retrieve output with given ID
func (c *Client) GetOutput(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetOutputRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetOutputRequest create the request corresponding to the get action endpoint of the output resource.
func (c *Client) NewGetOutputRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListOutputPath computes a request path to the list action of output.
func ListOutputPath() string {

	return fmt.Sprintf("/v1/outputs")
}

// Retrieve all outputs definitions
func (c *Client) ListOutput(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListOutputRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListOutputRequest create the request corresponding to the list action endpoint of the output resource.
func (c *Client) NewListOutputRequest(ctx context.Context, path string) (*http.Request, error) {
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

// SpecsOutputPath computes a request path to the specs action of output.
func SpecsOutputPath() string {

	return fmt.Sprintf("/v1/outputs/_specs")
}

// Retrieve all output types available
func (c *Client) SpecsOutput(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSpecsOutputRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSpecsOutputRequest create the request corresponding to the specs action endpoint of the output resource.
func (c *Client) NewSpecsOutputRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateOutputPayload is the output update action payload.
type UpdateOutputPayload struct {
	// Alias of the output
	Alias *string `form:"alias,omitempty" json:"alias,omitempty" yaml:"alias,omitempty" xml:"alias,omitempty"`
	// Output status
	Enabled bool `form:"enabled" json:"enabled" yaml:"enabled" xml:"enabled"`
	// Output properties
	Props map[string]interface{} `form:"props,omitempty" json:"props,omitempty" yaml:"props,omitempty" xml:"props,omitempty"`
	// Comma separated list of tags
	Tags *string `form:"tags,omitempty" json:"tags,omitempty" yaml:"tags,omitempty" xml:"tags,omitempty"`
}

// UpdateOutputPath computes a request path to the update action of output.
func UpdateOutputPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/v1/outputs/%s", param0)
}

// Update an output
func (c *Client) UpdateOutput(ctx context.Context, path string, payload *UpdateOutputPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateOutputRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateOutputRequest create the request corresponding to the update action endpoint of the output resource.
func (c *Client) NewUpdateOutputRequest(ctx context.Context, path string, payload *UpdateOutputPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
