/*
 * server
 *
 * <br/>https://ncloud.beta-apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-28T05:08:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type PublicIpInstance struct {

	// 공인IP인스턴스번호
PublicIpInstanceNo *string `json:"publicIpInstanceNo,omitempty"`

	// 공인IP
PublicIp *string `json:"publicIp,omitempty"`

	// 공인IP설명
PublicIpDescription *string `json:"publicIpDescription,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// 인터넷라인구분
InternetLineType *CommonCode `json:"internetLineType,omitempty"`

	// 공인IP인스턴스상태명
PublicIpInstanceStatusName *string `json:"publicIpInstanceStatusName,omitempty"`

	// 공인IP인스턴스상태
PublicIpInstanceStatus *CommonCode `json:"publicIpInstanceStatus,omitempty"`

	// 공인IP인스턴스OP
PublicIpInstanceOperation *CommonCode `json:"publicIpInstanceOperation,omitempty"`

	// 공인IP종류구분
PublicIpKindType *CommonCode `json:"publicIpKindType,omitempty"`

	// 공인IP할당된서버인스턴스
ServerInstanceAssociatedWithPublicIp *ServerInstance `json:"serverInstanceAssociatedWithPublicIp,omitempty"`

	// 리전
Region *Region `json:"region,omitempty"`

	// ZONE
Zone *Zone `json:"zone,omitempty"`
}
