/*
 * sourcebuild
 *
 * <br/>https://sourcebuild.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcebuild

type GetProjectDetailResponseEnvPlatform struct {
	Type_ *string `json:"type,omitempty"`

	Config *EnvPlatformConfigResponse `json:"config,omitempty"`
}