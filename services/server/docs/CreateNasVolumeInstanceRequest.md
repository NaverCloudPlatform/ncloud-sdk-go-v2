# CreateNasVolumeInstanceRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeName** | ***string** | 볼륨이름 | [default to null]
**VolumeSize** | ***int32** | NAS볼륨사이즈 | [default to null]
**VolumeAllotmentProtocolTypeCode** | ***string** | 볼륨할당프로토콜유형코드 | [default to null]
**AccessControlRuleList** | **[[]\*AccessControlRuleParameter](AccessControlRuleParameter.md)** | 접근제어Rule리스트 | [optional] [default to null]
**CifsUserName** | ***string** | CIFS유저이름 | [optional] [default to null]
**CifsUserPassword** | ***string** | CIFS유저비밀번호 | [optional] [default to null]
**NasVolumeDescription** | ***string** | NAS볼륨설명 | [optional] [default to null]
**RegionNo** | ***string** | 리전번호 | [optional] [default to null]
**ZoneNo** | ***string** | ZONE번호 | [optional] [default to null]
**IsReturnProtection** | ***bool** | 반납보호여부 | [optional] [default to null]
**IsAsync** | ***bool** | Async여부 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


