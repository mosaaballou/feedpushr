// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "feedpushr": feed Resource Client
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
	"strconv"
)

// CreateFeedPath computes a request path to the create action of feed.
func CreateFeedPath() string {

	return fmt.Sprintf("/v1/feeds")
}

// Create a new feed
func (c *Client) CreateFeed(ctx context.Context, path string, url_ string, tags *string, title *string) (*http.Response, error) {
	req, err := c.NewCreateFeedRequest(ctx, path, url_, tags, title)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateFeedRequest create the request corresponding to the create action endpoint of the feed resource.
func (c *Client) NewCreateFeedRequest(ctx context.Context, path string, url_ string, tags *string, title *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("url", url_)
	if tags != nil {
		values.Set("tags", *tags)
	}
	if title != nil {
		values.Set("title", *title)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// DeleteFeedPath computes a request path to the delete action of feed.
func DeleteFeedPath(id string) string {
	param0 := id

	return fmt.Sprintf("/v1/feeds/%s", param0)
}

// Delete a feed
func (c *Client) DeleteFeed(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteFeedRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteFeedRequest create the request corresponding to the delete action endpoint of the feed resource.
func (c *Client) NewDeleteFeedRequest(ctx context.Context, path string) (*http.Request, error) {
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

// GetFeedPath computes a request path to the get action of feed.
func GetFeedPath(id string) string {
	param0 := id

	return fmt.Sprintf("/v1/feeds/%s", param0)
}

// Retrieve feed with given ID
func (c *Client) GetFeed(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetFeedRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetFeedRequest create the request corresponding to the get action endpoint of the feed resource.
func (c *Client) NewGetFeedRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListFeedPath computes a request path to the list action of feed.
func ListFeedPath() string {

	return fmt.Sprintf("/v1/feeds")
}

// Retrieve all feeds
func (c *Client) ListFeed(ctx context.Context, path string, limit *int, page *int) (*http.Response, error) {
	req, err := c.NewListFeedRequest(ctx, path, limit, page)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListFeedRequest create the request corresponding to the list action endpoint of the feed resource.
func (c *Client) NewListFeedRequest(ctx context.Context, path string, limit *int, page *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if limit != nil {
		tmp27 := strconv.Itoa(*limit)
		values.Set("limit", tmp27)
	}
	if page != nil {
		tmp28 := strconv.Itoa(*page)
		values.Set("page", tmp28)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// StartFeedPath computes a request path to the start action of feed.
func StartFeedPath(id string) string {
	param0 := id

	return fmt.Sprintf("/v1/feeds/%s/start", param0)
}

// Start feed aggregation
func (c *Client) StartFeed(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewStartFeedRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewStartFeedRequest create the request corresponding to the start action endpoint of the feed resource.
func (c *Client) NewStartFeedRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// StopFeedPath computes a request path to the stop action of feed.
func StopFeedPath(id string) string {
	param0 := id

	return fmt.Sprintf("/v1/feeds/%s/stop", param0)
}

// Stop feed aggregation
func (c *Client) StopFeed(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewStopFeedRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewStopFeedRequest create the request corresponding to the stop action endpoint of the feed resource.
func (c *Client) NewStopFeedRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateFeedPath computes a request path to the update action of feed.
func UpdateFeedPath(id string) string {
	param0 := id

	return fmt.Sprintf("/v1/feeds/%s", param0)
}

// Update a feed
func (c *Client) UpdateFeed(ctx context.Context, path string, tags *string, title *string) (*http.Response, error) {
	req, err := c.NewUpdateFeedRequest(ctx, path, tags, title)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateFeedRequest create the request corresponding to the update action endpoint of the feed resource.
func (c *Client) NewUpdateFeedRequest(ctx context.Context, path string, tags *string, title *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if tags != nil {
		values.Set("tags", *tags)
	}
	if title != nil {
		values.Set("title", *title)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
