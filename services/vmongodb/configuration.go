/*
 * vmongodb
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmongodb/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmongodb

import (
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"os"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

func NewConfiguration(apiKeys ...*ncloud.APIKey) *ncloud.Configuration {
	cfg := &ncloud.Configuration{
		BasePath:      "https://ncloud.apigw.ntruss.com/vmongodb/v2",
		DefaultHeader: make(map[string]string),
		UserAgent:     "vmongodb/1.0.1/go",
	}
	if len(apiKeys) > 0 {
		cfg.APIKey = apiKeys[0]
	}
	cfg.InitCredentials()
	if os.Getenv("NCLOUD_API_GW") != "" {
		cfg.BasePath = os.Getenv("NCLOUD_API_GW") + "/vmongodb/v2"
	}
	return cfg
}
