/*
 * vsourcedeploy
 *
 * <br/>https://vpcsourcedeploy.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vsourcedeploy

type GetDeployHistoryDetailResponseTargets struct {

Server *GetDeployHistoryDetailResponseServer `json:"server,omitempty"`

Status *string `json:"status,omitempty"`

Time *GetDeployHistoryDetailResponseTime `json:"time,omitempty"`

Step *GetDeployHistoryDetailResponseStep `json:"step,omitempty"`
}
