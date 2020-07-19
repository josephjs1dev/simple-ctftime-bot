package ctftime

import (
	"io/ioutil"
	"net/http"
)

var defaultClient = &http.Client{}

// GetWithClient requests and returns body
func GetWithClient(url string, client *http.Client) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
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

// Get calls GetWithClient with default client
func Get(url string) (string, error) {
	return GetWithClient(url, defaultClient)
}
