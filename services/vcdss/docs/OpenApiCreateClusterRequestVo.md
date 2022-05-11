# OpenApiCreateClusterRequestVo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokerNodeCount** | **int32** | 브로커 노드 갯수 | min:3 | only 3 to up | [default to null]
**BrokerNodeProductCode** | **string** | 브로커 노드 상품(Spec) 코드 | [default to null]
**BrokerNodeStorageSize** | **int64** | 데이터 노드 스토리지 사이즈 | [default to null]
**BrokerNodeStorageType2Code** | **string** | 데이터 노드 스토리지 타입 | Default: SSD | [default to null]
**BrokerNodeSubnetName** | **string** | 브로커 서브넷 이름 | [default to null]
**BrokerNodeSubnetNo** | **int32** | 브로커 서브넷 번호 | [default to null]
**ClusterName** | **string** | 클러스터 이름 | [default to null]
**KafkaManagerUserName** | **string** | 카프카 매니저 접근 유저 네임 | [default to null]
**KafkaManagerUserPassword** | **string** | 카프카 매니저 접근 유저 비밀번호 | [default to null]
**KafkaVersionCode** | **string** | Application 버전 코드 | [default to null]
**ManagerNodeCount** | **int32** | 매니저 노드 갯수 | Default:1 | only 1 | [default to null]
**ManagerNodeProductCode** | **string** | 매니저 노드 상품(Spec) 코드 | [default to null]
**ManagerNodeSubnetName** | **string** | 매니저 서브넷 이름 | [optional] [default to null]
**ManagerNodeSubnetNo** | **int32** | 매니저 서브넷 번호 | [optional] [default to null]
**SoftwareProductCode** | **string** | 소프트웨어 상품(OS) 코드 | [default to null]
**VpcName** | **string** | VPC 이름 | [optional] [default to null]
**VpcNo** | **int32** | VPC 번호 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


