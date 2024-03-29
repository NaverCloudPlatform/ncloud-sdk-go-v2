/*
 * api
 *
 * Cloud Data Streaming Service - Open API<br/>https://clouddatastreamingservice.beta-apigw.ntruss.com/api/v1
 *
 * API version: 2022-12-06T13:50:42Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type ResponseVoGetNodeProductListForSpecChangeResponseVo struct {
	Code      int32                                      `json:"code,omitempty"`
	Message   string                                     `json:"message,omitempty"`
	RequestId string                                     `json:"requestId,omitempty"`
	Result    *GetNodeProductListForSpecChangeResponseVo `json:"result,omitempty"`
}
