/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type ChangeServerInstanceSpecRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// 서버인스턴스번호
	ServerInstanceNo *string `json:"serverInstanceNo"`

	// 서버상품코드
	ServerProductCode *string `json:"serverProductCode"`

	// 서버스펙코드
	ServerSpecCode *string `json:"serverSpecCode"`
}
