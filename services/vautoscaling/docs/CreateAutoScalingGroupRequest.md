# CreateAutoScalingGroupRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionCode** | ***string** | REGION코드 | [optional] [default to null]
**LaunchConfigurationNo** | ***string** | 론치설정번호 | [default to null]
**AutoScalingGroupName** | ***string** | 오토스케일링그룹이름 | [optional] [default to null]
**VpcNo** | ***string** | VPC번호 | [default to null]
**SubnetNo** | ***string** | 서브넷번호 | [default to null]
**AccessControlGroupNoList** | **[]\*string** | ACG번호리스트 | [default to null]
**ServerNamePrefix** | ***string** | 서버이름Prefix | [optional] [default to null]
**MinSize** | ***int32** | 최소용량 | [default to null]
**MaxSize** | ***int32** | 최대용량 | [default to null]
**DesiredCapacity** | ***int32** | 기대용량 | [optional] [default to null]
**DefaultCoolDown** | ***int32** | 쿨다운기본값 | [optional] [default to null]
**HealthCheckGracePeriod** | ***int32** | 헬스체크보류기간 | [optional] [default to null]
**HealthCheckTypeCode** | ***string** | 헬스체크유형코드 | [optional] [default to null]
**TargetGroupNoList** | **[]\*string** | 타겟그룹번호리스트 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


