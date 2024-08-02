/*
 * vpostgresql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vpostgresql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpostgresql

type ExportBackupToObjectStorageRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 백업파일이름
FileName *string `json:"fileName"`

	// ObjectStorage버킷이름
BucketName *string `json:"bucketName"`

	// CloudPostgresql인스턴스번호
CloudPostgresqlInstanceNo *string `json:"cloudPostgresqlInstanceNo"`
}