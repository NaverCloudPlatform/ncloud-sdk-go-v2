package ncloud

import (
	"bufio"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud/credentials"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type APIKey struct {
	AccessKey string
	SecretKey string
}

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
	APIKey        *APIKey
	Credentials   *credentials.Credentials
}

// Keys This is for backward compatibility. Will be deprecated soon.
func Keys() *APIKey {
	apiKey := &APIKey{
		AccessKey: "",
		SecretKey: "",
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if usr.HomeDir == "" {
		log.Fatal("use.HomeDir is nil")
		return nil
	}

	configureFile := filepath.Join(usr.HomeDir, ".ncloud", "configure")
	file, err := os.Open(configureFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "=")
		switch strings.TrimSpace(s[0]) {
		case "ncloud_access_key_id":
			apiKey.AccessKey = strings.TrimSpace(s[1])
		case "ncloud_secret_access_key":
			apiKey.SecretKey = strings.TrimSpace(s[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	return apiKey
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func (c *Configuration) GetCredentials() *credentials.Credentials {
	if c.ValidCredentials() {
		return c.Credentials.Retrieve()
	}
	return nil
}

func (c *Configuration) ValidCredentials() bool {
	return c.Credentials != nil && c.Credentials.Valid()
}

func (c *Configuration) InitCredentials() {
	if c.APIKey != nil {
		c.Credentials = credentials.NewValueProviderCreds(c.APIKey.AccessKey, c.APIKey.SecretKey)
	} else if !c.ValidCredentials() {
		c.Credentials = credentials.LoadCredentials(credentials.DefaultCredentialsChain())
	}
}
