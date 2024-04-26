/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type ExportDbServerLogToObjectStorageRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 데이터베이스서버로그파일타입
LogType *string `json:"logType"`

	// 데이터베이스서버로그파일이름
FileName *string `json:"fileName"`

	// ObjectStorage버킷이름
BucketName *string `json:"bucketName"`

	// CloudMysql서버인스턴스번호
CloudMysqlServerInstanceNo *string `json:"cloudMysqlServerInstanceNo"`

	// 폴더경로
FolderPath *string `json:"folderPath,omitempty"`
}
