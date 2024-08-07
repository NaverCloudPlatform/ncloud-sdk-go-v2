/*
 * vpostgresql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vpostgresql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpostgresql

type CloudPostgresqlBackup struct {

	// CloudPostgresql인스턴스번호
CloudPostgresqlInstanceNo *string `json:"cloudPostgresqlInstanceNo,omitempty"`

	// CloudPostgresql서비스이름
CloudPostgresqlServiceName *string `json:"cloudPostgresqlServiceName,omitempty"`

	// 백업유형
BackupType *string `json:"backupType,omitempty"`

	// 백업파일보관기간
BackupRetention *int32 `json:"backupRetention,omitempty"`

	// 백업시간
BackupTime *string `json:"backupTime,omitempty"`

	// 백업사이즈
BackupSize *int64 `json:"backupSize,omitempty"`

	// 마지막백업일시
LastBackupDate *string `json:"lastBackupDate,omitempty"`

	// 백업파일보관개수
BackupFileStorageCount *int32 `json:"backupFileStorageCount,omitempty"`
}
