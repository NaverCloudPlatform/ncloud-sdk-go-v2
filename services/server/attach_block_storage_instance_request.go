/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-08-08T01:43:18Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type AttachBlockStorageInstanceRequest struct {

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo"`

	// 블록스토리지인스턴스번호
BlockStorageInstanceNo *string `json:"blockStorageInstanceNo"`
}
