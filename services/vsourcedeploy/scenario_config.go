/*
 * vsourcedeploy
 *
 * <br/>https://vpcsourcedeploy.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vsourcedeploy

type ScenarioConfig struct {

Strategy *string `json:"strategy,omitempty"`

File *ScenarioConfigFile `json:"file,omitempty"`

Rollback *bool `json:"rollback,omitempty"`

Cmd *ScenarioConfigCmd `json:"cmd,omitempty"`

LoadBalancer *ScenarioConfigLoadBalancer `json:"loadBalancer,omitempty"`

Manifest *ScenarioConfigManifest `json:"manifest,omitempty"`

CanaryConfig *ScenarioConfigCanaryConfig `json:"canaryConfig,omitempty"`

Path []*ScenarioConfigCmdDeploy `json:"path,omitempty"`
}
