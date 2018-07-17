# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/loadbalancer/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLoadBalancerSslCertificate**](V2Api.md#AddLoadBalancerSslCertificate) | **Post** /addLoadBalancerSslCertificate | 
[**ChangeLoadBalancedServerInstances**](V2Api.md#ChangeLoadBalancedServerInstances) | **Post** /changeLoadBalancedServerInstances | 
[**ChangeLoadBalancerInstanceConfiguration**](V2Api.md#ChangeLoadBalancerInstanceConfiguration) | **Post** /changeLoadBalancerInstanceConfiguration | 
[**CreateLoadBalancerInstance**](V2Api.md#CreateLoadBalancerInstance) | **Post** /createLoadBalancerInstance | 
[**DeleteLoadBalancerInstances**](V2Api.md#DeleteLoadBalancerInstances) | **Post** /deleteLoadBalancerInstances | 
[**DeleteLoadBalancerSslCertificate**](V2Api.md#DeleteLoadBalancerSslCertificate) | **Post** /deleteLoadBalancerSslCertificate | 
[**GetLoadBalancedServerInstanceList**](V2Api.md#GetLoadBalancedServerInstanceList) | **Post** /getLoadBalancedServerInstanceList | 
[**GetLoadBalancerInstanceList**](V2Api.md#GetLoadBalancerInstanceList) | **Post** /getLoadBalancerInstanceList | 
[**GetLoadBalancerSslCertificateList**](V2Api.md#GetLoadBalancerSslCertificateList) | **Post** /getLoadBalancerSslCertificateList | 
[**GetLoadBalancerTargetServerInstanceList**](V2Api.md#GetLoadBalancerTargetServerInstanceList) | **Post** /getLoadBalancerTargetServerInstanceList | 


# **AddLoadBalancerSslCertificate**
> AddLoadBalancerSslCertificateResponse AddLoadBalancerSslCertificate(ctx, addLoadBalancerSslCertificateRequest)


로드밸런서SSL인증서추가

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **addLoadBalancerSslCertificateRequest** | [**AddLoadBalancerSslCertificateRequest**](AddLoadBalancerSslCertificateRequest.md)| addLoadBalancerSslCertificateRequest | 

### Return type

[**AddLoadBalancerSslCertificateResponse**](addLoadBalancerSslCertificateResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeLoadBalancedServerInstances**
> ChangeLoadBalancedServerInstancesResponse ChangeLoadBalancedServerInstances(ctx, changeLoadBalancedServerInstancesRequest)


로드밸런서에Bind된서버인스턴스변경

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **changeLoadBalancedServerInstancesRequest** | [**ChangeLoadBalancedServerInstancesRequest**](ChangeLoadBalancedServerInstancesRequest.md)| changeLoadBalancedServerInstancesRequest | 

### Return type

[**ChangeLoadBalancedServerInstancesResponse**](changeLoadBalancedServerInstancesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeLoadBalancerInstanceConfiguration**
> ChangeLoadBalancerInstanceConfigurationResponse ChangeLoadBalancerInstanceConfiguration(ctx, changeLoadBalancerInstanceConfigurationRequest)


로드밸런서인스턴스설정변경

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **changeLoadBalancerInstanceConfigurationRequest** | [**ChangeLoadBalancerInstanceConfigurationRequest**](ChangeLoadBalancerInstanceConfigurationRequest.md)| changeLoadBalancerInstanceConfigurationRequest | 

### Return type

[**ChangeLoadBalancerInstanceConfigurationResponse**](changeLoadBalancerInstanceConfigurationResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerInstance**
> CreateLoadBalancerInstanceResponse CreateLoadBalancerInstance(ctx, createLoadBalancerInstanceRequest)


로드밸런서인스턴스생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **createLoadBalancerInstanceRequest** | [**CreateLoadBalancerInstanceRequest**](CreateLoadBalancerInstanceRequest.md)| createLoadBalancerInstanceRequest | 

### Return type

[**CreateLoadBalancerInstanceResponse**](createLoadBalancerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerInstances**
> DeleteLoadBalancerInstancesResponse DeleteLoadBalancerInstances(ctx, deleteLoadBalancerInstancesRequest)


로드밸런서인스턴스삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deleteLoadBalancerInstancesRequest** | [**DeleteLoadBalancerInstancesRequest**](DeleteLoadBalancerInstancesRequest.md)| deleteLoadBalancerInstancesRequest | 

### Return type

[**DeleteLoadBalancerInstancesResponse**](deleteLoadBalancerInstancesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerSslCertificate**
> DeleteLoadBalancerSslCertificateResponse DeleteLoadBalancerSslCertificate(ctx, deleteLoadBalancerSslCertificateRequest)


로드밸런서SSL인증서삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deleteLoadBalancerSslCertificateRequest** | [**DeleteLoadBalancerSslCertificateRequest**](DeleteLoadBalancerSslCertificateRequest.md)| deleteLoadBalancerSslCertificateRequest | 

### Return type

[**DeleteLoadBalancerSslCertificateResponse**](deleteLoadBalancerSslCertificateResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancedServerInstanceList**
> GetLoadBalancedServerInstanceListResponse GetLoadBalancedServerInstanceList(ctx, getLoadBalancedServerInstanceListRequest)


로드밸런서Bind된서버인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getLoadBalancedServerInstanceListRequest** | [**GetLoadBalancedServerInstanceListRequest**](GetLoadBalancedServerInstanceListRequest.md)| getLoadBalancedServerInstanceListRequest | 

### Return type

[**GetLoadBalancedServerInstanceListResponse**](getLoadBalancedServerInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerInstanceList**
> GetLoadBalancerInstanceListResponse GetLoadBalancerInstanceList(ctx, getLoadBalancerInstanceListRequest)


로드밸런서인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getLoadBalancerInstanceListRequest** | [**GetLoadBalancerInstanceListRequest**](GetLoadBalancerInstanceListRequest.md)| getLoadBalancerInstanceListRequest | 

### Return type

[**GetLoadBalancerInstanceListResponse**](getLoadBalancerInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerSslCertificateList**
> GetLoadBalancerSslCertificateListResponse GetLoadBalancerSslCertificateList(ctx, getLoadBalancerSslCertificateListRequest)


로드밸런서SSL인증서조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getLoadBalancerSslCertificateListRequest** | [**GetLoadBalancerSslCertificateListRequest**](GetLoadBalancerSslCertificateListRequest.md)| getLoadBalancerSslCertificateListRequest | 

### Return type

[**GetLoadBalancerSslCertificateListResponse**](getLoadBalancerSslCertificateListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerTargetServerInstanceList**
> GetLoadBalancerTargetServerInstanceListResponse GetLoadBalancerTargetServerInstanceList(ctx, getLoadBalancerTargetServerInstanceListRequest)


로드밸런서Target서버인스턴스리스트

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getLoadBalancerTargetServerInstanceListRequest** | [**GetLoadBalancerTargetServerInstanceListRequest**](GetLoadBalancerTargetServerInstanceListRequest.md)| getLoadBalancerTargetServerInstanceListRequest | 

### Return type

[**GetLoadBalancerTargetServerInstanceListResponse**](getLoadBalancerTargetServerInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

