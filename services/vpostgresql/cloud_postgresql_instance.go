/*
 * vpostgresql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vpostgresql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpostgresql

type CloudPostgresqlInstance struct {

	// CloudPostgresql인스턴스번호
CloudPostgresqlInstanceNo *string `json:"cloudPostgresqlInstanceNo,omitempty"`

	// CloudPostgresql서비스이름
CloudPostgresqlServiceName *string `json:"cloudPostgresqlServiceName,omitempty"`

	// CloudPostgresql인스턴스상태이름
CloudPostgresqlInstanceStatusName *string `json:"cloudPostgresqlInstanceStatusName,omitempty"`

	// CloudPostgresql인스턴스상태
CloudPostgresqlInstanceStatus *CommonCode `json:"cloudPostgresqlInstanceStatus,omitempty"`

	// CloudPostgresql인스턴스OP
CloudPostgresqlInstanceOperation *CommonCode `json:"cloudPostgresqlInstanceOperation,omitempty"`

	// CloudPostgresql이미지상품코드
CloudPostgresqlImageProductCode *string `json:"cloudPostgresqlImageProductCode,omitempty"`

	// CloudPostgresql엔진버전
EngineVersion *string `json:"engineVersion,omitempty"`

	// 세대코드
GenerationCode *string `json:"generationCode,omitempty"`

	// CloudPostgresql라이선스
License *string `json:"license,omitempty"`

	// CloudPostgresql포트
CloudPostgresqlPort *int32 `json:"cloudPostgresqlPort,omitempty"`

	// 고가용성여부
IsHa *bool `json:"isHa,omitempty"`

	// 멀티존여부
IsMultiZone *bool `json:"isMultiZone,omitempty"`

	// 백업여부
IsBackup *bool `json:"isBackup,omitempty"`

	// 백업파일보관기간
BackupFileRetentionPeriod *int32 `json:"backupFileRetentionPeriod,omitempty"`

	// 백업시간
BackupTime *string `json:"backupTime,omitempty"`

	// 생성일자
CreateDate *string `json:"createDate,omitempty"`

	// ACG번호리스트
AccessControlGroupNoList *AccessControlGroupNoList `json:"accessControlGroupNoList,omitempty"`

	// CloudPostgresqlConfig리스트
CloudPostgresqlConfigList *CloudPostgresqlConfigList `json:"cloudPostgresqlConfigList,omitempty"`

	// CloudPostgresql서버인스턴스리스트
CloudPostgresqlServerInstanceList []*CloudPostgresqlServerInstance `json:"cloudPostgresqlServerInstanceList,omitempty"`
}