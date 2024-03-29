/*
 * vloadbalancer
 *
 * VPC Load Balancer 관련 API<br/>https://ncloud.apigw.ntruss.com/vloadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancer

type GetTargetGroupListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 타겟그룹번호리스트
TargetGroupNoList []*string `json:"targetGroupNoList,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 정렬리스트
SortList *string `json:"sortList,omitempty"`

	// 타겟유형코드
TargetTypeCode *string `json:"targetTypeCode,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`
}
