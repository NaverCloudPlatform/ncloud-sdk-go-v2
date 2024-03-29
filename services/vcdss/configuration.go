/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

import (
	"fmt"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"os"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

func NewConfiguration(region string, apiKeys ...*ncloud.APIKey) *ncloud.Configuration {
	cfg := &ncloud.Configuration{
		BasePath:      "https://clouddatastreamingservice.apigw.ntruss.com/api/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "vcdss//go",
	}
	if len(apiKeys) > 0 {
		cfg.APIKey = apiKeys[0]
	}
	cfg.InitCredentials()

	var ncloudApiGw string
	if os.Getenv("NCLOUD_API_GW") == "" {
		ncloudApiGw = "https://clouddatastreamingservice.apigw.ntruss.com"
	} else {
		ncloudApiGw = os.Getenv("NCLOUD_API_GW")
	}

	ncloudApiGw = strings.Replace(ncloudApiGw, "https://ncloud.", "https://clouddatastreamingservice.", 1)
	ncloudApiGw = strings.Replace(ncloudApiGw, "https://fin-ncloud.", "https://fin-clouddatastreamingservice.", 1)

	switch region {
	case "KR": // 민간, 공공-수도권
		cfg.BasePath = fmt.Sprintf("%s/api/v1", ncloudApiGw)
	case "FKR": // 금융
		cfg.BasePath = fmt.Sprintf("%s/api/v1", ncloudApiGw)
	default: // 민간-싱가폴, 공공-남부권
		cfg.BasePath = fmt.Sprintf("%s/api/%s-v1", ncloudApiGw, strings.ToLower(region))
	}
	return cfg
}
