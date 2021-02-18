package nauditor_sdk

import (
	"errors"
	"io/ioutil"
	"net/http"

	json "github.com/goccy/go-json"
)

//
var (
	baseURL = "https://nameauditor-whois-check.p.rapidapi.com"
	//
	ErrDataInvalid = errors.New("bad data")
)

// Client
type Client struct {
	key string
}

// NewClient
func NewClient(key string) *Client {
	return &Client{key}
}

// Ping
func (c *Client) Ping() error {
	req, err := http.NewRequest("GET", baseURL+"/ping", nil)
	if err != nil {
		return err
	}

	req.Header.Add("x-rapidapi-key", c.key)
	req.Header.Add("x-rapidapi-host", "nameauditor-whois-check.p.rapidapi.com")

	_, err = http.DefaultClient.Do(req)

	return err
}

// GetTLDs
func (c *Client) GetTLDs() (*TLDsList, error) {
	req, err := http.NewRequest("GET", baseURL+"/tlds", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", c.key)
	req.Header.Add("x-rapidapi-host", "nameauditor-whois-check.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response = struct {
		Code    int
		Message string
		Data    interface{}
	}{}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if response.Code != http.StatusOK {
		return nil, errors.New(response.Message)
	}

	b, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}
	var tlds = &TLDsList{}
	if err := json.Unmarshal(b, &tlds); err != nil {
		return nil, err
	}

	return tlds, nil
}

// GetWhois
func (c *Client) GetWhois(domain string) (*Domain, error) {
	req, _ := http.NewRequest("GET", baseURL+"/whois/"+domain, nil)

	req.Header.Add("x-rapidapi-key", c.key)
	req.Header.Add("x-rapidapi-host", "nameauditor-whois-check.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response = struct {
		Code    int
		Message string
		Data    interface{}
	}{}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if response.Code != http.StatusOK {
		return nil, errors.New(response.Message)
	}

	b, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var dom = &Domain{}
	if err := json.Unmarshal(b, &dom); err != nil {
		return nil, err
	}

	return dom, nil
}
