package metadata

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var endpoint = "http://169.254.169.254"
var basePath = "/latest/meta-data/"

type ApiClient struct {
	client *http.Client
}

func NewApiClient() *ApiClient {
	if os.Getenv("NCLOUD_METADATA_API_ENDPOINT") != "" {
		endpoint = os.Getenv("NCLOUD_METADATA_API_ENDPOINT")
	}
	c := &http.Client{
		Timeout: 3 * time.Second,
	}
	return &ApiClient{client: c}
}

func (c ApiClient) GetMetadata(url string) (string, error) {
	resp, err := c.client.Get(endpoint + basePath + url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}
