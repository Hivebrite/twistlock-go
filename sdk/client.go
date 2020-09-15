package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	Token      string
	httpClient *http.Client
}

type ServerError struct {
	StatusCode int
	Err        string `json:"err"`
}

func NewServerError(statusCode int, body []byte) error {
	apiErr := ServerError{StatusCode: statusCode}

	decodeErr := json.Unmarshal(body, &apiErr)
	if decodeErr != nil {
		apiErr.Err = string(body)
	}

	return &apiErr
}

func (err ServerError) Error() string {
	return fmt.Sprintf("code: %d - reason: %s", err.StatusCode, err.Err)
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	u := c.BaseURL.ResolveReference(&url.URL{Path: path})

	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("user-agent", c.UserAgent)

	if body != nil {
		req.Header.Set("content-type", "application/json")
	}

	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, NewServerError(resp.StatusCode, body)
	}

	if v == nil {
		return resp, nil
	}

	return resp, json.Unmarshal(body, v)
}

func NewClient(apiUrl string) (*Client, error) {
	u, err := url.Parse(apiUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURL:    u,
		httpClient: http.DefaultClient,
	}, nil
}
