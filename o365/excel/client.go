package excel

import (
	"SaaS-Squash/common"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ExcelClient struct {
	// BaseURL      *url.URL
	BaseURL     string
	UserAgent   string
	DriveId     string
	SheetId     string
	SheetName   string
	TokenId     string
	HttpClient  common.HTTPClient
	APIKey      string
	Ticker      int
	TickerCell  string
	Commands    []Command
	AuthExpire  time.Time
	Debug       bool
	Credentials common.Config
}

func (c *ExcelClient) newRequest(method, path string, body *bytes.Buffer) (*http.Request, error) {
	if c.BaseURL == "" {
		return nil, errors.New("BaseURL is undefined")
	}

	u, _ := url.JoinPath(c.BaseURL, path)

	if body == nil {
		body = new(bytes.Buffer)
	}

	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}
	// Default request is json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	return req, nil
}

func (c *ExcelClient) do(req *http.Request,
	v interface{}) (*http.Response, error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)

	return resp, err
}

func (c *ExcelClient) do_noparse(req *http.Request) ([]byte, error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
