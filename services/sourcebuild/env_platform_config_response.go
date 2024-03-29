/*
 * sourcebuild
 *
 * <br/>https://sourcebuild.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcebuild

type EnvPlatformConfigResponse struct {
	Os *EnvPlatformConfigResponseOs `json:"os,omitempty"`

	Runtime *EnvPlatformConfigResponseRuntime `json:"runtime,omitempty"`

	Registry *string `json:"registry,omitempty"`

	Image *string `json:"image,omitempty"`

	Tag *string `json:"tag,omitempty"`
}
