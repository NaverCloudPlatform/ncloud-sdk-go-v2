# GlobalCdnRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtocolTypeCode** | **string** | 프로토콜구분코드 | [optional] [default to null]
**ServiceDomainTypeCode** | **string** | 서비스도메인구분코드 | [optional] [default to null]
**OriginUrl** | **string** | 원본URL | [optional] [default to null]
**OriginPath** | **string** | 원본경로 | [optional] [default to null]
**OriginHttpPort** | **int32** | 원본HTTP포트 | [optional] [default to null]
**OriginHttpsPort** | **int32** | 원본HTTPS포트 | [optional] [default to null]
**ForwardHostHeaderTypeCode** | **string** | 포워드호스트헤더구분코드 | [optional] [default to null]
**ForwardHostHeader** | **string** | 포워드호스트헤더 | [optional] [default to null]
**CacheKeyHostNameTypeCode** | **string** | 캐쉬키호스트명구분코드 | [optional] [default to null]
**IsGzipCompressionUse** | **bool** | GZIP압축사용여부 | [optional] [default to null]
**CachingOptionTypeCode** | **string** | 캐싱옵션구분코드 | [optional] [default to null]
**IsErrorContentsResponseUse** | **bool** | 오류내용응답사용여부 | [optional] [default to null]
**CachingTtlTime** | **string** | TTL캐싱 | [optional] [default to null]
**IsQueryStringIgnoreUse** | **bool** | 쿼리스트링무시여부 | [optional] [default to null]
**IsRemoveVaryHeaderUse** | **bool** | 헤더제거사용여부 | [optional] [default to null]
**IsLargeFileOptimizationUse** | **bool** | 대용량파일최적화사용여부 | [optional] [default to null]
**GzipResponseTypeCode** | **string** | GZIP응답구분코드 | [optional] [default to null]
**IsReferrerDomainUse** | **bool** | 레퍼러도메인사용여부 | [optional] [default to null]
**ReferrerDomainList** | **[]string** | 레퍼러도메인리스트 | [optional] [default to null]
**IsReferrerDomainRestrictUse** | **bool** | 레퍼러도메인제한사용여부 | [optional] [default to null]
**IsSecureTokenUse** | **bool** | 보안토큰사용여부 | [optional] [default to null]
**SecureTokenPassword** | **string** | 보안토큰비밀번호 | [optional] [default to null]
**IsReissueSecureTokenPassword** | **bool** | 보안토큰재발급여부 | [optional] [default to null]
**CertificateName** | **string** | 인증서이름 | [optional] [default to null]
**IsAccessLogUse** | **bool** | 엑세스로그사용여부 | [optional] [default to null]
**AccessLogFileStorageContainerName** | **string** | 엑세스로그파일스토리지인스턴스이름 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


