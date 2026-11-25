# CreateLoadBalancerInstanceRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionCode** | ***string** | REGION코드 | [optional] [default to null]
**IdleTimeout** | ***int32** | 연결타임아웃 | [optional] [default to null]
**LoadBalancerDescription** | ***string** | 로드밸런서설명 | [optional] [default to null]
**LoadBalancerNetworkTypeCode** | ***string** | 로드밸런서네트워크유형코드 | [optional] [default to null]
**LoadBalancerTypeCode** | ***string** | 로드밸런서유형코드 | [default to null]
**LoadBalancerListenerList** | **[[]\*LoadBalancerListenerParameter](LoadBalancerListenerParameter.md)** | 로드밸런서리스너리스트 | [optional] [default to null]
**LoadBalancerName** | ***string** | 로드밸런서이름 | [optional] [default to null]
**ThroughputTypeCode** | ***string** | 부하처리성능유형코드 | [optional] [default to null]
**VpcNo** | ***string** | VPC번호 | [default to null]
**SubnetNoList** | **[]\*string** | 서브넷번호리스트 | [default to null]
**LoadBalancerSubnetList** | **[[]\*LoadBalancerSubnetParameter](LoadBalancerSubnetParameter.md)** | 로드밸런서서브넷리스트 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


