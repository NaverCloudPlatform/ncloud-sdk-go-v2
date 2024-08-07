/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetServerProductListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// 서버이미지상품코드
ServerImageProductCode *string `json:"serverImageProductCode,omitempty"`

	// 제외할상품코드
ExclusionProductCode *string `json:"exclusionProductCode,omitempty"`

	// 상품코드
ProductCode *string `json:"productCode,omitempty"`

	// 세대코드
GenerationCode *string `json:"generationCode,omitempty"`

	// 회원서버이미지인스턴스번호
MemberServerImageInstanceNo *string `json:"memberServerImageInstanceNo,omitempty"`
}
