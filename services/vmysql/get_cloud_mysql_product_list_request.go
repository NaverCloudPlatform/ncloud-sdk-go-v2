/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type GetCloudMysqlProductListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// CloudMysql이미지상품코드
CloudMysqlImageProductCode *string `json:"cloudMysqlImageProductCode"`

	// 조회할상품코드
ProductCode *string `json:"productCode,omitempty"`

	// 제외할상품코드
ExclusionProductCode *string `json:"exclusionProductCode,omitempty"`
}
