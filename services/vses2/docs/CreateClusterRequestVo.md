# CreateClusterRequestVo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** | 클러스터 이름 | [default to null]
**DataNodeCount** | **int32** | 데이터 노드 갯수 | min:3 | only 3 to up | [default to null]
**DataNodeProductCode** | **string** | 데이터 노드 상품(Spec) 코드 | [default to null]
**DataNodeStorageSize** | **int64** | 데이터 노드 스토리지 사이즈 | [default to null]
**DataNodeStorageType2Code** | **string** | 데이터 노드 스토리지 타입 | Default: SSD | [default to null]
**DataNodeSubnetName** | **string** | 데이터 서브넷 이름 | [default to null]
**DataNodeSubnetNo** | **int32** | 데이터 서브넷 번호 | [default to null]
**ElasticSearchHttpPort** | **string** | elasticSearch 접근 포트 | [optional] [default to null]
**ElasticSearchVersionCode** | **string** | elasticSearch 버전 코드 | [optional] [default to null]
**IsDualManager** | **bool** | 이중화 여부 | [default to null]
**IsMasterOnlyNodeActivated** | **bool** | 마스터 전용 노드 활성화 여부 | [default to null]
**KibanaHttpPort** | **string** | Kibana 접근 포트 | [optional] [default to null]
**KibanaUserName** | **string** | 키바나 접근 유저 네임 | [optional] [default to null]
**KibanaUserPassword** | **string** | 키바나 접근 유저 비밀번호 | [optional] [default to null]
**LoginKeyName** | **string** | 서버 접속 인증키 네임 | [default to null]
**ManagerNodeCount** | **int32** | 매니저 노드 갯수 | Default:1 | only 1 | [default to null]
**ManagerNodeProductCode** | **string** | 매니저 노드 상품(Spec) 코드 | [default to null]
**ManagerNodeSubnetName** | **string** | 매니저 서브넷 이름 | [optional] [default to null]
**ManagerNodeSubnetNo** | **int32** | 매니저 서브넷 번호 | [optional] [default to null]
**MasterNodeCount** | **int32** | 마스터 노드 갯수 | only 3 or 5 | [default to null]
**MasterNodeProductCode** | **string** | 마스터 노드 상품(Spec) 코드 | [default to null]
**MasterNodeSubnetName** | **string** | 마스터 노드 서브넷 이름 | [default to null]
**MasterNodeSubnetNo** | **int32** | 마스터 노드 서브넷 번호 | [default to null]
**RegionNo** | **string** | regionNo | [default to null]
**RequestTypeCode** | **string** | 요청 구분 MC|API | [default to null]
**SearchEngineDashboardPort** | **string** | Search Engine Dashboard 접근 포트 | [optional] [default to null]
**SearchEnginePort** | **string** | Search Engine 접근 포트 | [optional] [default to null]
**SearchEngineUserName** | **string** | Search Engine 유저 ID | [optional] [default to null]
**SearchEngineUserPassword** | **string** | Search Engine 유저 Password | [optional] [default to null]
**SearchEngineVersionCode** | **string** | Search Engine 버전 코드 | [optional] [default to null]
**SoftwareProductCode** | **string** | 소프트웨어 상품(OS) 코드 | [default to null]
**ValidDataNodeStorageSize** | **bool** |  | [optional] [default to null]
**VpcName** | **string** | VPC 이름 | [optional] [default to null]
**VpcNo** | **int32** | VPC 번호 | [optional] [default to null]
**ZoneNo** | **string** | zone no | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


