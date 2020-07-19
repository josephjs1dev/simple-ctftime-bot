package ctftime

import (
	"io/ioutil"
	"net/http"
)

const baseURL = "https://ctftime.org"

// Client is our CTFTime Client struct
type Client struct {
	baseURL    string
	httpClient *http.Client
}

var defaultClient *Client

// GetClient returns DefaultClient
func GetClient() *Client {
	if defaultClient == nil {
		defaultClient = &Client{
			baseURL:    baseURL,
			httpClient: &http.Client{},
		}
	}

	return defaultClient
}

// Get requests and returns body string
func (c *Client) Get(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Need this
	req.Header.Set("User-Agent", "PostmanRuntime/7.26.1")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
