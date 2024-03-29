/*
 * vautoscaling
 *
 * VPC Auto Scaling 관련 API<br/>https://ncloud.apigw.ntruss.com/vautoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vautoscaling

type GetAutoScalingActivityLogListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 오토스케일링그룹번호
AutoScalingGroupNo *string `json:"autoScalingGroupNo"`

	// 액티비티번호리스트
ActivityNoList []*string `json:"activityNoList,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`
}
