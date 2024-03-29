/*
 * vhadoop
 *
 * <br/>https://ncloud.apigw.ntruss.com/vhadoop/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vhadoop

type GetCloudHadoopProductListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// CloudHadoop이미지상품코드
CloudHadoopImageProductCode *string `json:"cloudHadoopImageProductCode,omitempty"`

	// 조회할상품코드
ProductCode *string `json:"productCode,omitempty"`

	// Cloud Hadoop 기타 서버 infra 상세 상품 코드
InfraResourceDetailTypeCode *string `json:"infraResourceDetailTypeCode,omitempty"`

	// 제외할상품코드
ExclusionProductCode *string `json:"exclusionProductCode,omitempty"`
}
