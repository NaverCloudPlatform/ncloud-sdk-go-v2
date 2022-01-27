# CloudMysqlInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudMysqlInstanceNo** | ***string** | CloudMysql인스턴스번호 | [optional] [default to null]
**CloudMysqlServiceName** | ***string** | CloudMysql서비스이름 | [optional] [default to null]
**CloudMysqlInstanceStatusName** | ***string** | CloudMysql인스턴스상태이름 | [optional] [default to null]
**CloudMysqlInstanceStatus** | **[*CommonCode](CommonCode.md)** | CloudMysql인스턴스상태 | [optional] [default to null]
**CloudMysqlInstanceOperation** | **[*CommonCode](CommonCode.md)** | CloudMysql인스턴스OP | [optional] [default to null]
**CloudMysqlImageProductCode** | ***string** | CloudMysql이미지상품코드 | [optional] [default to null]
**EngineVersion** | ***string** | CloudMysql엔진버전 | [optional] [default to null]
**License** | **[*CommonCode](CommonCode.md)** | CloudMysql라이선스 | [optional] [default to null]
**CloudMysqlPort** | ***int32** | CloudMysql포트 | [optional] [default to null]
**IsHa** | ***bool** | 고가용성여부 | [optional] [default to null]
**IsMultiZone** | ***bool** | 멀티존여부 | [optional] [default to null]
**IsBackup** | ***bool** | 백업여부 | [optional] [default to null]
**BackupFileRetentionPeriod** | ***int32** | 백업파일보관기간 | [optional] [default to null]
**BackupTime** | ***string** | 백업시간 | [optional] [default to null]
**CreateDate** | ***string** | 생성일자 | [optional] [default to null]
**AccessControlGroupNoList** | **[*AccessControlGroupNoList](AccessControlGroupNoList.md)** | ACG번호리스트 | [optional] [default to null]
**CloudMysqlConfigList** | **[*CloudMysqlConfigList](CloudMysqlConfigList.md)** | CloudMysqlConfig리스트 | [optional] [default to null]
**CloudMysqlServerInstanceList** | **[[]\*CloudMysqlServerInstance](CloudMysqlServerInstance.md)** | CloudMysql서버인스턴스리스트 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


