/*
 * vautoscaling
 *
 * VPC Auto Scaling 관련 API<br/>https://ncloud.apigw.ntruss.com/vautoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vautoscaling

type AutoScalingGroup struct {

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`

	// 서브넷번호
SubnetNo *string `json:"subnetNo,omitempty"`

	// 서버이름Prefix
ServerNamePrefix *string `json:"serverNamePrefix,omitempty"`

	// 오토스케일링그룹번호
AutoScalingGroupNo *string `json:"autoScalingGroupNo,omitempty"`

	// 오토스케일링그룹이름
AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	// 론치설정번호
LaunchConfigurationNo *string `json:"launchConfigurationNo,omitempty"`

	// 최소용량
MinSize *int32 `json:"minSize,omitempty"`

	// 최대용량
MaxSize *int32 `json:"maxSize,omitempty"`

	// 기대용량
DesiredCapacity *int32 `json:"desiredCapacity,omitempty"`

	// 쿨다운기본값
DefaultCoolDown *int32 `json:"defaultCoolDown,omitempty"`

	// 헬스체크보류기간
HealthCheckGracePeriod *int32 `json:"healthCheckGracePeriod,omitempty"`

	// 헬스체크유형
HealthCheckType *CommonCode `json:"healthCheckType,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// 오토스케일링그룹상태
AutoScalingGroupStatus *CommonCode `json:"autoScalingGroupStatus,omitempty"`

	// 타겟그룹번호리스트
TargetGroupNoList []*string `json:"targetGroupNoList,omitempty"`

	// 오토스케일링그룹에속한서버인스턴스리스트
InAutoScalingGroupServerInstanceList []*InAutoScalingGroupServerInstance `json:"inAutoScalingGroupServerInstanceList,omitempty"`

	// ACG번호리스트
AccessControlGroupNoList []*string `json:"accessControlGroupNoList,omitempty"`

	// 일시정지된프로세스리스트
SuspendedProcessList []*SuspendedProcess `json:"suspendedProcessList,omitempty"`
}