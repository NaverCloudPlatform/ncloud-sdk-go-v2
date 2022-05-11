/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type GetVpcConfigListResponseDetailVo struct {
	CreatedDate *DateTimeVo `json:"createdDate,omitempty"`
	Ipv4Cidr string `json:"ipv4Cidr,omitempty"`
	Permission string `json:"permission,omitempty"`
	RegionNo string `json:"regionNo,omitempty"`
	StatusCode string `json:"statusCode,omitempty"`
	VpcName string `json:"vpcName,omitempty"`
	VpcNo int32 `json:"vpcNo,omitempty"`
}