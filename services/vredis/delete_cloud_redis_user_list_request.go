/*
 * vredis
 *
 * <br/>https://ncloud.apigw.ntruss.com/vredis/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vredis

type DeleteCloudRedisUserListRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// CloudRedis인스턴스번호
	CloudRedisInstanceNo *string `json:"cloudRedisInstanceNo"`

	// CloudRedisDBUser리스트
	CloudRedisUserList []*CloudRedisUserKeyParameter `json:"cloudRedisUserList"`
}
