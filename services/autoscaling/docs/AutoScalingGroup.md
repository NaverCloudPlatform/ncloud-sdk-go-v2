# AutoScalingGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScalingGroupName** | ***string** | 오토스케일링그룹명 | [optional] [default to null]
**AutoScalingGroupNo** | ***string** | 오토스케일링그룹번호 | [optional] [default to null]
**LaunchConfigurationName** | ***string** | 론치설정명 | [optional] [default to null]
**LaunchConfigurationNo** | ***string** | 론치설정번호 | [optional] [default to null]
**DesiredCapacity** | ***int32** | 기대능력치 | [optional] [default to null]
**MinSize** | ***int32** | 최소사이즈 | [optional] [default to null]
**MaxSize** | ***int32** | 최대사이즈 | [optional] [default to null]
**DefaultCooldown** | ***int32** |  | [optional] [default to null]
**LoadBalancerInstanceSummaryList** | **[[]\*LoadBalancerInstanceSummary](LoadBalancerInstanceSummary.md)** | 로드밸런서인스턴스Summary리스트 | [optional] [default to null]
**HealthCheckGracePeriod** | ***int32** | 헬스체크보류기간 | [optional] [default to null]
**HealthCheckType** | **[*CommonCode](CommonCode.md)** | 헬스체크유형 | [optional] [default to null]
**CreateDate** | ***string** | 생성일시 | [optional] [default to null]
**InAutoScalingGroupServerInstanceList** | **[[]\*InAutoScalingGroupServerInstance](InAutoScalingGroupServerInstance.md)** | 오토스케일링그룹에속한서버인스턴스리스트 | [optional] [default to null]
**SuspendedProcessList** | **[[]\*SuspendedProcess](SuspendedProcess.md)** | 보류된프로세스리스트 | [optional] [default to null]
**ZoneList** | **[[]\*Zone](Zone.md)** | ZONE리스트 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


