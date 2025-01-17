// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "feedpushr": opml Resource Client
//
// Command:
// $ goagen
// --design=github.com/ncarlier/feedpushr/v3/design
// --out=/home/fr23972/workspace/fe/feedpushr/autogen
// --version=v1.4.3

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GetOpmlPath computes a request path to the get action of opml.
func GetOpmlPath() string {

	return fmt.Sprintf("/v2/opml")
}

// Get all feeds as an OPML format
func (c *Client) GetOpml(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetOpmlRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetOpmlRequest create the request corresponding to the get action endpoint of the opml resource.
func (c *Client) NewGetOpmlRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// StatusOpmlPath computes a request path to the status action of opml.
func StatusOpmlPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/v2/opml/status/%s", param0)
}

// Get OPML import status
func (c *Client) StatusOpml(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewStatusOpmlRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewStatusOpmlRequest create the request corresponding to the status action endpoint of the opml resource.
func (c *Client) NewStatusOpmlRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UploadOpmlPath computes a request path to the upload action of opml.
func UploadOpmlPath() string {

	return fmt.Sprintf("/v2/opml")
}

// Upload an OPML file to create feeds
func (c *Client) UploadOpml(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewUploadOpmlRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUploadOpmlRequest create the request corresponding to the upload action endpoint of the opml resource.
func (c *Client) NewUploadOpmlRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
