/*
 * server
 *
 * <br/>https://ncloud.beta-apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-28T05:08:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetMemberServerImageListRequest struct {

	// 회원서버이미지번호리스트
MemberServerImageNoList []*string `json:"memberServerImageNoList,omitempty"`

	// 플랫폼타입코드리스트
PlatformTypeCodeList []*string `json:"platformTypeCodeList,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`

	// 소팅대상
SortedBy *string `json:"sortedBy,omitempty"`

	// 소팅순서
SortingOrder *string `json:"sortingOrder,omitempty"`
}
