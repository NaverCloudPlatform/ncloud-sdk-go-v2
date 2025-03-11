/*
 * vnks
 *
 * <br/>https://nks.apigw.ntruss.com/vnks/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnks

type Cluster struct {

	// 클러스터 uuid
	Uuid *string `json:"uuid"`

	// 클러스터 acg 이름
	AcgName *string `json:"acgName"`

	// 클러스터 acg no
	AcgNo *int32 `json:"acgNo"`

	// 클러스터 이름
	Name *string `json:"name"`

	// 클러스터 용량
	Capacity *string `json:"capacity"`

	// 클러스터 타입
	ClusterType *string `json:"clusterType"`

	// Hypervisor Code
	HypervisorCode *string `json:"hypervisorCode"`

	// 등록된 노드 총 개수
	NodeCount *int32 `json:"nodeCount"`

	// 사용할 수 있는 노드의 최대 개수
	NodeMaxCount *int32 `json:"nodeMaxCount"`

	// cpu 개수
	CpuCount *int32 `json:"cpuCount"`

	// 메모리 용량
	MemorySize *int64 `json:"memorySize"`

	// 생성 일자
	CreatedAt *string `json:"createdAt"`

	// Control Plane API 주소
	Endpoint *string `json:"endpoint"`

	// 쿠버네티스 버전
	K8sVersion *string `json:"k8sVersion"`

	// region의 코드
	RegionCode *string `json:"regionCode"`

	// 클러스터 상태
	Status *string `json:"status"`

	// CNI Plugin Code
	KubeNetworkPlugin *string `json:"kubeNetworkPlugin"`

	// 로드밸런서 전용 서브넷 이름
	SubnetLbName *string `json:"subnetLbName,omitempty"`

	// 로드밸런서 전용 Private Subnet No
	SubnetLbNo *int32 `json:"subnetLbNo,omitempty"`

	// 로드밸런서 전용 Private Subnet No
	LbPrivateSubnetNo *int32 `json:"lbPrivateSubnetNo"`

	// 로드밸런서 전용 Public Subnet No
	LbPublicSubnetNo *int32 `json:"lbPublicSubnetNo,omitempty"`

	// 서브넷 이름
	SubnetName *string `json:"subnetName,omitempty"`

	// 서브넷 No 목록
	SubnetNoList []*int32 `json:"subnetNoList"`

	// 최근 업데이트 일자
	UpdatedAt *string `json:"updatedAt"`

	// vpc 이름
	VpcName *string `json:"vpcName"`

	// vpc 번호
	VpcNo *int32 `json:"vpcNo"`

	// zone 코드
	ZoneCode *string `json:"zoneCode"`

	// zone 번호
	ZoneNo *int32 `json:"zoneNo"`

	// InitScript 번호
	InitScriptNo *int32 `json:"initScriptNo,omitempty"`

	// InitScript 이름
	InitScriptName *string `json:"initScriptName,omitempty"`

	// Pod Security Policy 설정 여부
	PodSecurityPolicyEnabled *bool `json:"podSecurityPolicyEnabled,omitempty"`

	// 로그인 키 이름
	LoginKeyName *string `json:"loginKeyName"`

	// 노드풀
	NodePool []*NodePool `json:"nodePool"`

	Log *ClusterLogInput `json:"log"`

	// Public Network
	PublicNetwork *bool `json:"publicNetwork"`

	ReturnProtection *bool `json:"returnProtection"`

	KmsKeyTag *string `json:"kmsKeyTag,omitempty"`
}
