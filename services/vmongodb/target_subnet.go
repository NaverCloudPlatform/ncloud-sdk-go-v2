/*
 * vmongodb
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmongodb/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmongodb

type TargetSubnet struct {

	// Subnet번호
SubnetNo *string `json:"subnetNo,omitempty"`

	// Subnet이름
SubnetName *string `json:"subnetName,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`

	// VPC이름
VpcName *string `json:"vpcName,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// Subnet
Subnet *string `json:"subnet,omitempty"`

	// Public여부
IsPublic *string `json:"isPublic,omitempty"`

	// 생성시간
CreatedDate *string `json:"createdDate,omitempty"`
}
