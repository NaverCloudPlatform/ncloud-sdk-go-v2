# Cluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | ***string** | 클러스터 uuid | [default to null]
**AcgName** | ***string** | 클러스터 acg 이름 | [default to null]
**AcgNo** | ***int32** | 클러스터 acg no | [default to null]
**Name** | ***string** | 클러스터 이름 | [default to null]
**Capacity** | ***string** | 클러스터 용량 | [default to null]
**ClusterType** | ***string** | 클러스터 타입 | [default to null]
**HypervisorCode** | ***string** | Hypervisor Code | [default to null]
**NodeCount** | ***int32** | 등록된 노드 총 개수 | [default to null]
**NodeMaxCount** | ***int32** | 사용할 수 있는 노드의 최대 개수 | [default to null]
**CpuCount** | ***int32** | cpu 개수 | [default to null]
**MemorySize** | ***int64** | 메모리 용량 | [default to null]
**CreatedAt** | ***string** | 생성 일자 | [default to null]
**Endpoint** | ***string** | Control Plane API 주소 | [default to null]
**K8sVersion** | ***string** | 쿠버네티스 버전 | [default to null]
**RegionCode** | ***string** | region의 코드 | [default to null]
**Status** | ***string** | 클러스터 상태 | [default to null]
**KubeNetworkPlugin** | ***string** | CNI Plugin Code | [default to null]
**SubnetLbName** | ***string** | 로드밸런서 전용 서브넷 이름 | [optional] [default to null]
**SubnetLbNo** | ***int32** | 로드밸런서 전용 Private Subnet No | [optional] [default to null]
**LbPrivateSubnetNo** | ***int32** | 로드밸런서 전용 Private Subnet No | [default to null]
**LbPublicSubnetNo** | ***int32** | 로드밸런서 전용 Public Subnet No | [optional] [default to null]
**SubnetName** | ***string** | 서브넷 이름 | [optional] [default to null]
**SubnetNoList** | **[]\*int32** | 서브넷 No 목록 | [default to null]
**UpdatedAt** | ***string** | 최근 업데이트 일자 | [default to null]
**VpcName** | ***string** | vpc 이름 | [default to null]
**VpcNo** | ***int32** | vpc 번호 | [default to null]
**ZoneCode** | ***string** | zone 코드 | [default to null]
**ZoneNo** | ***int32** | zone 번호 | [default to null]
**InitScriptNo** | ***int32** | InitScript 번호 | [optional] [default to null]
**InitScriptName** | ***string** | InitScript 이름 | [optional] [default to null]
**PodSecurityPolicyEnabled** | ***bool** | Pod Security Policy 설정 여부 | [optional] [default to null]
**LoginKeyName** | ***string** | 로그인 키 이름 | [default to null]
**NodePool** | **[[]\*NodePool](NodePool.md)** | 노드풀 | [default to null]
**Log** | **[*ClusterLogInput](ClusterLogInput.md)** |  | [default to null]
**PublicNetwork** | ***bool** | Public Network | [default to null]
**ReturnProtection** | ***bool** |  | [default to null]
**KmsKeyTag** | ***string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


