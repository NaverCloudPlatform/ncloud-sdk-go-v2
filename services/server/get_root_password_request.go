/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-05T07:55:47Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetRootPasswordRequest struct {

	// 개인키
PrivateKey *string `json:"privateKey,omitempty"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo"`
}