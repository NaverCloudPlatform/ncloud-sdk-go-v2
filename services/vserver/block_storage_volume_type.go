/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type BlockStorageVolumeType struct {

	// ZONE코드리스트
ZoneCodeList []*string `json:"zoneCodeList,omitempty"`

	// 블록스토리지볼륨타입
BlockStorageVolumeType *CommonCode `json:"blockStorageVolumeType,omitempty"`

	// 하이퍼바이저타입
HypervisorType *CommonCode `json:"hypervisorType,omitempty"`

	// 블록스토리지 볼륨타입 최소 throughput
MinThroughput *int32 `json:"minThroughput,omitempty"`

	// 블록스토리지 볼륨타입 최대 throughput
MaxThroughput *int32 `json:"maxThroughput,omitempty"`

	// 블록스토리지 볼륨타입 최소 iops
MinIops *int32 `json:"minIops,omitempty"`

	// 블록스토리지 볼륨타입 최대 iops
MaxIops *int32 `json:"maxIops,omitempty"`

	// 블록스토리지 볼륨타입 최소 volumeSize
MinVolumeSize *int32 `json:"minVolumeSize,omitempty"`

	// 블록스토리지 볼륨타입 최대 volumeSize
MaxVolumeSize *int32 `json:"maxVolumeSize,omitempty"`

	// 블록스토리지 볼륨타입 최소 rootVolumeSize
MinBaseVolumeSize *int32 `json:"minBaseVolumeSize,omitempty"`

	// 블록스토리지 볼륨타입 최대 rootVolumeSize
MaxBaseVolumeSize *int32 `json:"maxBaseVolumeSize,omitempty"`

	// 기본스토리지 가능 여부
IsAvailableBase *bool `json:"isAvailableBase,omitempty"`
}