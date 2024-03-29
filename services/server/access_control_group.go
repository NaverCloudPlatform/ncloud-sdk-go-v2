/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type AccessControlGroup struct {

	// 접근제어그룹설정번호
AccessControlGroupConfigurationNo *string `json:"accessControlGroupConfigurationNo,omitempty"`

	// 접근제어그룹명
AccessControlGroupName *string `json:"accessControlGroupName,omitempty"`

	// 접근제어그룹설명
AccessControlGroupDescription *string `json:"accessControlGroupDescription,omitempty"`

	// 디폴트그룹여부
IsDefaultGroup *bool `json:"isDefaultGroup,omitempty"`

	// 생성일자
CreateDate *string `json:"createDate,omitempty"`
}
