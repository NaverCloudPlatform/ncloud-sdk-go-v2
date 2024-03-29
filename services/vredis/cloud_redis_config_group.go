/*
 * vredis
 *
 * <br/>https://ncloud.apigw.ntruss.com/vredis/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vredis

type CloudRedisConfigGroup struct {

	// configGroup번호
	ConfigGroupNo *string `json:"configGroupNo,omitempty"`

	// configGroup이름
	ConfigGroupName *string `json:"configGroupName,omitempty"`

	// configGroup설명
	ConfigGroupDescription *string `json:"configGroupDescription,omitempty"`

	// configGroup상태명
	ConfigGroupStatusName *string `json:"configGroupStatusName,omitempty"`

	// configGroup상태
	ConfigGroupStatus *CommonCode `json:"configGroupStatus,omitempty"`

	// cloudredis버전
	CloudRedisVersion *string `json:"cloudRedisVersion,omitempty"`

	// 생성일자
	CreateDate *string `json:"createDate,omitempty"`

	// 수정일자
	ModifyDate *string `json:"modifyDate,omitempty"`
}
