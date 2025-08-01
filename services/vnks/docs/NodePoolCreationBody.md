# NodePoolCreationBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | ***string** | 노드풀 이름 | [default to null]
**NodeCount** | ***int32** | 등록 될 노드 개수 | [optional] [default to null]
**SubnetNo** | ***int32** | Subnet 번호 | [optional] [default to null]
**SubnetNoList** | **[]\*int32** | Subnet 번호 | [optional] [default to null]
**SoftwareCode** | ***string** | Server image code | [optional] [default to null]
**ProductCode** | ***string** | 상품 코드 [서버 스펙 목록](/docs/compute-vserver-server-common-getserverproductlist) | [optional] [default to null]
**ServerSpecCode** | ***string** | Server spec code | [optional] [default to null]
**StorageSize** | ***int32** | Storage size | [optional] [default to null]
**Autoscale** | **[*AutoscalerUpdate](AutoscalerUpdate.md)** |  | [optional] [default to null]
**Labels** | **[[]\*NodePoolLabel](NodePoolLabel.md)** |  | [optional] [default to null]
**Taints** | **[[]\*NodePoolTaint](NodePoolTaint.md)** |  | [optional] [default to null]
**ServerRoleId** | ***string** |  | [optional] [default to null]
**FabricCluster** | **[*FabricClusterPool](FabricClusterPool.md)** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


