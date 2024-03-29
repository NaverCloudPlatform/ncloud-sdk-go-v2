/*
 * vmssql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmssql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmssql

type CreateCloudMssqlSlaveInstanceRequest struct {

	// Cloud MSSQL 인스턴스 번호
CloudMssqlInstanceNo *string `json:"cloudMssqlInstanceNo"`

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// private domain postfix
PrivateDomainPostfix *string `json:"privateDomainPostfix"`
}
