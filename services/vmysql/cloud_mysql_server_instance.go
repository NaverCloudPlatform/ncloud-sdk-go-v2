/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type CloudMysqlServerInstance struct {

	// CloudMysql서버인스턴스번호
CloudMysqlServerInstanceNo *string `json:"cloudMysqlServerInstanceNo,omitempty"`

	// CloudMysql서버이름
CloudMysqlServerName *string `json:"cloudMysqlServerName,omitempty"`

	// CloudMysql서버역할
CloudMysqlServerRole *CommonCode `json:"cloudMysqlServerRole,omitempty"`

	// CloudMysql인스턴스상태이름
CloudMysqlServerInstanceStatusName *string `json:"cloudMysqlServerInstanceStatusName,omitempty"`

	// CloudMysql서버인스턴스상태
CloudMysqlServerInstanceStatus *CommonCode `json:"cloudMysqlServerInstanceStatus,omitempty"`

	// CloudMysql서버인스턴스OP
CloudMysqlServerInstanceOperation *CommonCode `json:"cloudMysqlServerInstanceOperation,omitempty"`

	// CloudMysql상품코드
CloudMysqlProductCode *string `json:"cloudMysqlProductCode,omitempty"`

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`

	// Subnet번호
SubnetNo *string `json:"subnetNo,omitempty"`

	// PublicSubnet여부
IsPublicSubnet *bool `json:"isPublicSubnet,omitempty"`

	// 공인도메인명
PublicDomain *string `json:"publicDomain,omitempty"`

	// 사설도메인명
PrivateDomain *string `json:"privateDomain,omitempty"`

	// 내부IP
PrivateIp *string `json:"privateIp,omitempty"`

	// 데이터스토리지타입
DataStorageType *CommonCode `json:"dataStorageType,omitempty"`

	// 데이터스토리지암호화여부
IsStorageEncryption *bool `json:"isStorageEncryption,omitempty"`

	// 데이터스토리지사이즈
DataStorageSize *int64 `json:"dataStorageSize,omitempty"`

	// 사용중인데이터스토리지사이즈
UsedDataStorageSize *int64 `json:"usedDataStorageSize,omitempty"`

	// virtualCPU개수
CpuCount *int32 `json:"cpuCount,omitempty"`

	// 메모리사이즈
MemorySize *int64 `json:"memorySize,omitempty"`

	// 업시간
Uptime *string `json:"uptime,omitempty"`

	// 생성일자
CreateDate *string `json:"createDate,omitempty"`
}
