# LoadBalancerInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerInstanceNo** | ***string** | 로드밸런서인스턴스번호 | [optional] [default to null]
**VirtualIp** | ***string** | virtualIp | [optional] [default to null]
**LoadBalancerName** | ***string** | 로드밸런서명 | [optional] [default to null]
**LoadBalancerAlgorithmType** | **[*CommonCode](CommonCode.md)** | 로드밸런서알고리즘구분코 | [optional] [default to null]
**LoadBalancerDescription** | ***string** | 로드밸런서설명 | [optional] [default to null]
**CreateDate** | ***string** | 생성일자 | [optional] [default to null]
**DomainName** | ***string** | 도메인명 | [optional] [default to null]
**InternetLineType** | **[*CommonCode](CommonCode.md)** | 인터넷회선구분 | [optional] [default to null]
**LoadBalancerInstanceStatusName** | ***string** | 로드밸런서인스턴스상태명 | [optional] [default to null]
**LoadBalancerInstanceStatus** | **[*CommonCode](CommonCode.md)** | 로드밸런서인스턴스상태 | [optional] [default to null]
**LoadBalancerInstanceOperation** | **[*CommonCode](CommonCode.md)** | 로드밸런서인스턴스OP | [optional] [default to null]
**NetworkUsageType** | **[*CommonCode](CommonCode.md)** | 네트워크사용구분 | [optional] [default to null]
**IsHttpKeepAlive** | ***bool** | httpKeepAlive사용여부 | [optional] [default to null]
**ConnectionTimeout** | ***int32** | 커넥션타임아웃 | [optional] [default to null]
**CertificateName** | ***string** | SSL인증명 | [optional] [default to null]
**LoadBalancerRuleList** | **[[]\*LoadBalancerRule](LoadBalancerRule.md)** |  | [optional] [default to null]
**LoadBalancedServerInstanceList** | **[[]\*LoadBalancedServerInstance](LoadBalancedServerInstance.md)** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


