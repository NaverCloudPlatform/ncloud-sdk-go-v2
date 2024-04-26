/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type GetCloudMysqlImageProductListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 조회할상품코드
ProductCode *string `json:"productCode,omitempty"`

	// 제외할상품코드
ExclusionProductCode *string `json:"exclusionProductCode,omitempty"`

	// 세대코드
GenerationCode *string `json:"generationCode,omitempty"`
}
