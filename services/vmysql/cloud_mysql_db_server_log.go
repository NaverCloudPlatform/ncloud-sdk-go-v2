/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type CloudMysqlDbServerLog struct {

	// 데이터베이스서버로그파일이름
FileName *string `json:"fileName,omitempty"`

	// 데이터베이스서버로그파일사이즈
FileSize *int64 `json:"fileSize,omitempty"`

	// 데이터베이스서버로그파일수정일시
FileDate *string `json:"fileDate,omitempty"`
}
