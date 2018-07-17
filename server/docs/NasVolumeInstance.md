# NasVolumeInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NasVolumeInstanceNo** | **string** | NAS볼륨인스턴스번호 | [optional] [default to null]
**NasVolumeInstanceStatus** | [***CommonCode**](CommonCode.md) | NAS볼륨인스턴스상태 | [optional] [default to null]
**NasVolumeInstanceOperation** | [***CommonCode**](CommonCode.md) | NAS볼륨인스턴스OP | [optional] [default to null]
**NasVolumeInstanceStatusName** | **string** | 볼륨인스턴스상태명 | [optional] [default to null]
**CreateDate** | **string** | 생성일시 | [optional] [default to null]
**NasVolumeInstanceDescription** | **string** | NAS볼륨인스턴스설명 | [optional] [default to null]
**MountInformation** | **string** | 마운트정보 | [optional] [default to null]
**VolumeAllotmentProtocolType** | [***CommonCode**](CommonCode.md) | 볼륨할당프로토콜구분 | [optional] [default to null]
**VolumeName** | **string** | 볼륨명 | [optional] [default to null]
**VolumeTotalSize** | **int64** | 볼륨총사이즈 | [optional] [default to null]
**VolumeSize** | **int64** | 볼륨사이즈 | [optional] [default to null]
**VolumeUseSize** | **int64** | 볼륨사용사이즈 | [optional] [default to null]
**VolumeUseRatio** | **float32** | 볼륨사용비율 | [optional] [default to null]
**SnapshotVolumeConfigurationRatio** | **float32** | 스냅샷볼륨설정비율 | [optional] [default to null]
**SnapshotVolumeConfigPeriodType** | [***CommonCode**](CommonCode.md) | 스냅샷볼륨설정기간구분 | [optional] [default to null]
**SnapshotVolumeConfigTime** | **int32** | 스냅샷볼륨설정시간 | [optional] [default to null]
**SnapshotVolumeSize** | **int64** | 스냅샷사이즈 | [optional] [default to null]
**SnapshotVolumeUseSize** | **int64** | 스냅사용사이즈 | [optional] [default to null]
**SnapshotVolumeUseRatio** | **float32** | 스냅샷사용비율 | [optional] [default to null]
**IsSnapshotConfiguration** | **bool** | 스냅샷설정여부 | [optional] [default to null]
**IsEventConfiguration** | **bool** | 이벤트설정여부 | [optional] [default to null]
**Region** | [***Region**](Region.md) | 리전 | [optional] [default to null]
**Zone** | [***Zone**](Zone.md) | ZONE | [optional] [default to null]
**NasVolumeInstanceCustomIpList** | [**[]NasVolumeInstanceCustomIp**](NasVolumeInstanceCustomIp.md) | NAS볼륨커스텀IP리스트 | [optional] [default to null]
**NasVolumeServerInstanceList** | [**[]ServerInstance**](ServerInstance.md) | NAS볼륨서버인스턴스리스트 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


