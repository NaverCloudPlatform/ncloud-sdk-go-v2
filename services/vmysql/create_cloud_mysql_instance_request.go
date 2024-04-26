/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type CreateCloudMysqlInstanceRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo"`

	// CloudMysql이미지상품코드
CloudMysqlImageProductCode *string `json:"cloudMysqlImageProductCode,omitempty"`

	// CloudMysql상품코드
CloudMysqlProductCode *string `json:"cloudMysqlProductCode,omitempty"`

	// 데이터스토리지유형코드
DataStorageTypeCode *string `json:"dataStorageTypeCode,omitempty"`

	// 고가용성여부
IsHa *bool `json:"isHa,omitempty"`

	// 멀티존여부
IsMultiZone *bool `json:"isMultiZone,omitempty"`

	// 데이터스토리지암호화여부
IsStorageEncryption *bool `json:"isStorageEncryption,omitempty"`

	// 백업여부
IsBackup *bool `json:"isBackup,omitempty"`

	// 백업파일보관기간
BackupFileRetentionPeriod *int32 `json:"backupFileRetentionPeriod,omitempty"`

	// 백업시간
BackupTime *string `json:"backupTime,omitempty"`

	// 백업시간자동여부
IsAutomaticBackup *bool `json:"isAutomaticBackup,omitempty"`

	// CloudMysql서비스이름
CloudMysqlServiceName *string `json:"cloudMysqlServiceName"`

	// CloudMysql서버이름
CloudMysqlServerNamePrefix *string `json:"cloudMysqlServerNamePrefix"`

	// CloudMysql유저명
CloudMysqlUserName *string `json:"cloudMysqlUserName"`

	// CloudMysql유저패스워드
CloudMysqlUserPassword *string `json:"cloudMysqlUserPassword"`

	// 호스트IP
HostIp *string `json:"hostIp"`

	// CloudMysql포트
CloudMysqlPort *int32 `json:"cloudMysqlPort,omitempty"`

	// CloudMysqlDB명
CloudMysqlDatabaseName *string `json:"cloudMysqlDatabaseName"`

	// Subnet번호
SubnetNo *string `json:"subnetNo"`

	// StandbyMasterSubnet번호
StandbyMasterSubnetNo *string `json:"standbyMasterSubnetNo,omitempty"`

	// privateSubDomain사용여부
IsPrivateSubDomain *bool `json:"isPrivateSubDomain,omitempty"`

	// privateSubDomain접두어
PrivateSubDomainPrefix *string `json:"privateSubDomainPrefix,omitempty"`

	// engineVersionCode
EngineVersionCode *string `json:"engineVersionCode,omitempty"`
}
