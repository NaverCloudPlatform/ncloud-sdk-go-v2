/*
 * vnks
 *
 * <br/>https://nks.apigw.ntruss.com/vnks/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnks

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
		BasePath:      "https://nks.apigw.ntruss.com/vnks/v2",
		DefaultHeader: make(map[string]string),
		UserAgent:     "vnks/1.0.0/go",
	}
	if len(apiKeys) > 0 {
		cfg.APIKey = apiKeys[0]
	}
	cfg.InitCredentials()

	var ncloudApiGw string
	if os.Getenv("NCLOUD_API_GW") == "" {
		ncloudApiGw = "https://nks.apigw.ntruss.com"
	} else {
		ncloudApiGw = os.Getenv("NCLOUD_API_GW")
	}

	ncloudApiGw = strings.Replace(ncloudApiGw, "https://ncloud.", "https://nks.", 1)
	ncloudApiGw = strings.Replace(ncloudApiGw, "https://fin-ncloud.", "https://nks.", 1)

	switch {
	case strings.HasPrefix(region, "F"):
		cfg.BasePath = fmt.Sprintf("%s/nks/v2", ncloudApiGw)
	case region == "KR", strings.Contains(region, "CS"):
		cfg.BasePath = fmt.Sprintf("%s/vnks/v2", ncloudApiGw)
	default:
		cfg.BasePath = fmt.Sprintf("%s/vnks/%s-v2", ncloudApiGw, strings.ToLower(region))
	}

	cfg.DefaultHeader["x-ncp-region_code"] = region

	return cfg
}

func NewConfigurationWithUserAgent(region string, userAgent string, apiKeys ...*ncloud.APIKey) *ncloud.Configuration {
	cfg := NewConfiguration(region, apiKeys...)
	cfg.UserAgent = userAgent

	return cfg
}
