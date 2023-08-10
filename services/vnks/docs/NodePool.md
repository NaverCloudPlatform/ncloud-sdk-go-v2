# NodePool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceNo** | ***int32** | 인스턴스 번호 | [default to null]
**IsDefault** | ***bool** | default pool 여부 | [default to null]
**Name** | ***string** | 노드풀 이름 | [default to null]
**NodeCount** | ***int32** | 노드 개수 | [default to null]
**SubnetNoList** | **[]\*int32** | Subnet 번호 목록 | [default to null]
**SubnetNameList** | **[]\*string** | Subnet 이름 목록 | [default to null]
**Status** | ***string** | 노드풀 상태 | [default to null]
**Autoscale** | **[*AutoscaleOption](AutoscaleOption.md)** |  | [default to null]
**SoftwareCode** | ***string** | Software Code(OS) | [default to null]
**ProductCode** | ***string** | 상품 코드 | [optional] [default to null]
**K8sVersion** | ***string** | 쿠버네티스 버전 | [default to null]
**ServerSpecCode** | ***string** | 서버 스펙 코드 | [optional] [default to null]
**StorageSize** | ***int32** | 스토리지 크기 | [optional] [default to null]
**Labels** | **[[]\*NodePoolLabel](NodePoolLabel.md)** |  | [default to null]
**Taints** | **[[]\*NodePoolTaint](NodePoolTaint.md)** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


