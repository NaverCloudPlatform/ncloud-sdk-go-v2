/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetServerImageDetailRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 서버이미지번호
ServerImageNo *string `json:"serverImageNo,omitempty"`
}