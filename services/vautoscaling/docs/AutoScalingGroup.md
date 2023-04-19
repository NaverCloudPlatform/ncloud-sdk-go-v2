# AutoScalingGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcNo** | ***string** | VPC번호 | [optional] [default to null]
**SubnetNo** | ***string** | 서브넷번호 | [optional] [default to null]
**ServerNamePrefix** | ***string** | 서버이름Prefix | [optional] [default to null]
**AutoScalingGroupNo** | ***string** | 오토스케일링그룹번호 | [optional] [default to null]
**AutoScalingGroupName** | ***string** | 오토스케일링그룹이름 | [optional] [default to null]
**LaunchConfigurationNo** | ***string** | 론치설정번호 | [optional] [default to null]
**MinSize** | ***int32** | 최소용량 | [optional] [default to null]
**MaxSize** | ***int32** | 최대용량 | [optional] [default to null]
**DesiredCapacity** | ***int32** | 기대용량 | [optional] [default to null]
**DefaultCoolDown** | ***int32** | 쿨다운기본값 | [optional] [default to null]
**HealthCheckGracePeriod** | ***int32** | 헬스체크보류기간 | [optional] [default to null]
**HealthCheckType** | **[*CommonCode](CommonCode.md)** | 헬스체크유형 | [optional] [default to null]
**CreateDate** | ***string** | 생성일시 | [optional] [default to null]
**AutoScalingGroupStatus** | **[*CommonCode](CommonCode.md)** | 오토스케일링그룹상태 | [optional] [default to null]
**TargetGroupNoList** | **[]\*string** | 타겟그룹번호리스트 | [optional] [default to null]
**InAutoScalingGroupServerInstanceList** | **[[]\*InAutoScalingGroupServerInstance](InAutoScalingGroupServerInstance.md)** | 오토스케일링그룹에속한서버인스턴스리스트 | [optional] [default to null]
**AccessControlGroupNoList** | **[]\*string** | ACG번호리스트 | [optional] [default to null]
**SuspendedProcessList** | **[[]\*SuspendedProcess](SuspendedProcess.md)** | 일시정지된프로세스리스트 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


