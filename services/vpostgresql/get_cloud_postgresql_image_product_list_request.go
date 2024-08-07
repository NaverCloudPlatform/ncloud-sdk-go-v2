/*
 * vpostgresql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vpostgresql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpostgresql

type GetCloudPostgresqlImageProductListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 조회할상품코드
ProductCode *string `json:"productCode,omitempty"`

	// 제외할상품코드
ExclusionProductCode *string `json:"exclusionProductCode,omitempty"`
}
