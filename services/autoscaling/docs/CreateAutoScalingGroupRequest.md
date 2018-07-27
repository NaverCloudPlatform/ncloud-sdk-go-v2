# CreateAutoScalingGroupRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScalingGroupName** | ***string** | 오토스케일링그룹명 | [optional] [default to null]
**LaunchConfigurationName** | ***string** | 론치설정명 | [default to null]
**DesiredCapacity** | ***int32** | 기대용량치 | [optional] [default to null]
**MinSize** | ***int32** | 최소사이즈 | [default to null]
**MaxSize** | ***int32** | 최대사이즈 | [default to null]
**DefaultCooldown** | ***int32** | 디폴트쿨다운타임 | [optional] [default to null]
**LoadBalancerNameList** | **[]\*string** | 로드밸런서명리스트 | [optional] [default to null]
**HealthCheckGracePeriod** | ***int32** | 헬스체크보류기간 | [optional] [default to null]
**HealthCheckTypeCode** | ***string** | 헬스체크유형코드 | [optional] [default to null]
**ZoneNoList** | **[]\*string** | ZONE번호리스트 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


