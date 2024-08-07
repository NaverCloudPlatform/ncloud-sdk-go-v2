/*
 * vpostgresql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vpostgresql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpostgresql

type AddCloudPostgresqlUserListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// CloudPostgresql인스턴스번호
CloudPostgresqlInstanceNo *string `json:"cloudPostgresqlInstanceNo"`

	// CloudPostgresqlDBUser리스트
CloudPostgresqlUserList []*CloudPostgresqlUserParameter `json:"cloudPostgresqlUserList"`
}
