/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetFlowLogConfigurationListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 네트워크인터페이스번호리스트
NetworkInterfaceNoList []*string `json:"networkInterfaceNoList"`
}
