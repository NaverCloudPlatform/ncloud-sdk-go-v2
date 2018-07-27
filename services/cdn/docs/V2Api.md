# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/cdn/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCdnPlusInstanceList**](V2Api.md#GetCdnPlusInstanceList) | **Post** /getCdnPlusInstanceList | 
[**GetCdnPlusPurgeHistoryList**](V2Api.md#GetCdnPlusPurgeHistoryList) | **Post** /getCdnPlusPurgeHistoryList | 
[**GetGlobalCdnInstanceList**](V2Api.md#GetGlobalCdnInstanceList) | **Post** /getGlobalCdnInstanceList | 
[**GetGlobalCdnPurgeHistoryList**](V2Api.md#GetGlobalCdnPurgeHistoryList) | **Post** /getGlobalCdnPurgeHistoryList | 
[**RequestCdnPlusPurge**](V2Api.md#RequestCdnPlusPurge) | **Post** /requestCdnPlusPurge | 
[**RequestGlobalCdnPurge**](V2Api.md#RequestGlobalCdnPurge) | **Post** /requestGlobalCdnPurge | 


# **GetCdnPlusInstanceList**
> GetCdnPlusInstanceListResponse GetCdnPlusInstanceList(getCdnPlusInstanceListRequest)


CDN+인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCdnPlusInstanceListRequest** | **[\*GetCdnPlusInstanceListRequest](GetCdnPlusInstanceListRequest.md)** | getCdnPlusInstanceListRequest | 

### Return type

*[**GetCdnPlusInstanceListResponse**](getCdnPlusInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCdnPlusPurgeHistoryList**
> GetCdnPlusPurgeHistoryListResponse GetCdnPlusPurgeHistoryList(getCdnPlusPurgeHistoryListRequest)


CDN+퍼지기록조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCdnPlusPurgeHistoryListRequest** | **[\*GetCdnPlusPurgeHistoryListRequest](GetCdnPlusPurgeHistoryListRequest.md)** | getCdnPlusPurgeHistoryListRequest | 

### Return type

*[**GetCdnPlusPurgeHistoryListResponse**](getCdnPlusPurgeHistoryListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalCdnInstanceList**
> GetGlobalCdnInstanceListResponse GetGlobalCdnInstanceList(getGlobalCdnInstanceListRequest)


Global CDN 인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getGlobalCdnInstanceListRequest** | **[\*GetGlobalCdnInstanceListRequest](GetGlobalCdnInstanceListRequest.md)** | getGlobalCdnInstanceListRequest | 

### Return type

*[**GetGlobalCdnInstanceListResponse**](getGlobalCdnInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalCdnPurgeHistoryList**
> GetGlobalCdnPurgeHistoryListResponse GetGlobalCdnPurgeHistoryList(getGlobalCdnPurgeHistoryListRequest)


Global CDN퍼지기록조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getGlobalCdnPurgeHistoryListRequest** | **[\*GetGlobalCdnPurgeHistoryListRequest](GetGlobalCdnPurgeHistoryListRequest.md)** | getGlobalCdnPurgeHistoryListRequest | 

### Return type

*[**GetGlobalCdnPurgeHistoryListResponse**](getGlobalCdnPurgeHistoryListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestCdnPlusPurge**
> RequestCdnPlusPurgeResponse RequestCdnPlusPurge(requestCdnPlusPurgeRequest)


CDN+퍼지요청

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**requestCdnPlusPurgeRequest** | **[\*RequestCdnPlusPurgeRequest](RequestCdnPlusPurgeRequest.md)** | requestCdnPlusPurgeRequest | 

### Return type

*[**RequestCdnPlusPurgeResponse**](requestCdnPlusPurgeResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestGlobalCdnPurge**
> RequestGlobalCdnPurgeResponse RequestGlobalCdnPurge(requestGlobalCdnPurgeRequest)


Global CDN퍼지요청

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**requestGlobalCdnPurgeRequest** | **[\*RequestGlobalCdnPurgeRequest](RequestGlobalCdnPurgeRequest.md)** | requestGlobalCdnPurgeRequest | 

### Return type

*[**RequestGlobalCdnPurgeResponse**](requestGlobalCdnPurgeResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

