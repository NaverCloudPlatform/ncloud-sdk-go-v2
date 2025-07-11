# ClusterInputBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | ***string** | 클러스터 이름 | [default to null]
**ClusterType** | ***string** | 클러스터 타입 | [default to null]
**K8sVersion** | ***string** | 쿠버네티스 버전 [Version 조회 API](/docs/compute-vpckubernetesservice-nksv2#k8sSupportedVersion) | [optional] [default to null]
**LoginKeyName** | ***string** | 로그인 키 이름 | [default to null]
**ZoneCode** | ***string** | Zone 코드 | [optional] [default to null]
**ZoneNo** | ***int32** | Zone 번호 | [optional] [default to null]
**HypervisorCode** | ***string** | Hypervisor Code | [optional] [default to null]
**RegionCode** | ***string** | Region의 코드 | [default to null]
**PublicNetwork** | ***bool** | Public network | [optional] [default to null]
**KubeNetworkPlugin** | ***string** | CNI Plugin Code (ncloud-vpc-cni or cilium) | [optional] [default to null]
**VpcNo** | ***int32** | [VPC 번호](/docs/networking-vpc-vpcmanagement-getvpclist) | [default to null]
**SubnetNoList** | **[]\*int32** | [서브넷 번호 목록](/docs/networking-vpc-subnetmanagement-getsubnetlist) | [default to null]
**SubnetLbNo** | ***int32** | [로드밸런서 전용 Private Subnet 번호](/docs/networking-vpc-subnetmanagement-getsubnetlist) | [optional] [default to null]
**LbPrivateSubnetNo** | ***int32** | [로드밸런서 전용 Private Subnet 번호](/docs/networking-vpc-subnetmanagement-getsubnetlist) | [optional] [default to null]
**LbPublicSubnetNo** | ***int32** | [로드밸런서 전용 Public Subnet 번호](/docs/networking-vpc-subnetmanagement-getsubnetlist) | [optional] [default to null]
**Log** | **[*ClusterLogInput](ClusterLogInput.md)** |  | [optional] [default to null]
**KmsKeyTag** | ***string** |  | [optional] [default to null]
**AuthType** | ***string** | 클러스터 인증 타입 | [optional] [default to null]
**BootstrapAccessEntry** | ***bool** | 클러스터 생성 시 bootstrap access entry 자동 생성 여부 | [optional] [default to null]
**DefaultNodePool** | **[*DefaultNodePoolParam](DefaultNodePoolParam.md)** |  | [optional] [default to null]
**NodePool** | **[[]\*NodePoolDto](NodePoolDto.md)** | 추가 노드풀 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


