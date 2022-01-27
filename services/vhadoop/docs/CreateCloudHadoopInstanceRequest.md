# CreateCloudHadoopInstanceRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionCode** | ***string** | REGION코드 | [optional] [default to null]
**VpcNo** | ***string** | VPC번호 | [default to null]
**CloudHadoopClusterName** | ***string** | Cloud Hadoop 클러스터 이름 | [default to null]
**CloudHadoopImageProductCode** | ***string** | Cloud Hadoop이미지상품코드 | [optional] [default to null]
**CloudHadoopClusterTypeCode** | ***string** | Cloud Hadoop 클러스터 유형 코드 | [default to null]
**CloudHadoopAddOnCodeList** | **[]\*string** | Cloud Hadoop Add-On 코드 리스트 | [optional] [default to null]
**CloudHadoopAdminUserName** | ***string** | 클러스터 관리자 계정 | [default to null]
**CloudHadoopAdminUserPassword** | ***string** | 클러스터 관리자 계정 패스워드 | [default to null]
**LoginKeyName** | ***string** | 인증키 이름 | [optional] [default to null]
**BucketName** | ***string** | 버킷 이름 | [default to null]
**EdgeNodeProductCode** | ***string** | Cloud Hadoop 엣지노드 상품 코드 | [optional] [default to null]
**EdgeNodeSubnetNo** | ***string** | 엣지노드의 Subnet 번호 | [default to null]
**MasterNodeProductCode** | ***string** | Cloud Hadoop 마스터노드 상품 코드 | [optional] [default to null]
**MasterNodeSubnetNo** | ***string** | 마스터노드의 Subnet 번호 | [default to null]
**MasterNodeDataStorageTypeCode** | ***string** | 마스터노드의 데이터 스토리지 타입 코드 | [optional] [default to null]
**MasterNodeDataStorageSize** | ***int32** | 마스터노드의 데이터 스토리지 크기 | [default to null]
**WorkerNodeProductCode** | ***string** | Cloud Hadoop 작업자노드 상품 코드 | [optional] [default to null]
**WorkerNodeCount** | ***int32** | 작업자노드 개수 | [optional] [default to null]
**WorkerNodeSubnetNo** | ***string** | 작업자노드의 Subnet 번호 | [default to null]
**WorkerNodeDataStorageTypeCode** | ***string** | 작업자노드의 데이터 스토리지 타입 코드 | [optional] [default to null]
**WorkerNodeDataStorageSize** | ***int32** | 작업자노드의 데이터 스토리지 크기 | [default to null]
**UseKdc** | ***string** | 커버로스 인증 구성 여부 | [optional] [default to null]
**KdcRealm** | ***string** | KDC의 Realm 정보 | [optional] [default to null]
**KdcPassword** | ***string** | KDC admin 계정의 패스워드 | [optional] [default to null]
**UseBootstrapScript** | ***string** | Cloud Hadoop 부트스트랩 스크립트 사용 여부 | [optional] [default to null]
**BootstrapScript** | ***string** | Cloud Hadoop 부트스트랩 스크립트 사용 여부 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


