/*
 * vpostgresql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vpostgresql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpostgresql

type CloudPostgresqlDatabase struct {

	// Database이름
DatabaseName *string `json:"databaseName,omitempty"`

	// Database소유자
Owner *string `json:"owner,omitempty"`
}