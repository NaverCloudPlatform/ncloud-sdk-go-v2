# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/monitoring/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListMetrics**](V2Api.md#GetListMetrics) | **Post** /getListMetrics | 
[**GetMetricStatistics**](V2Api.md#GetMetricStatistics) | **Post** /getMetricStatistics | 


# **GetListMetrics**
> GetListMetricsResponse GetListMetrics(getListMetricsRequest)


B.메트릭 리스트 조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getListMetricsRequest** | **[\*GetListMetricsRequest](GetListMetricsRequest.md)** | getListMetricsRequest | 

### Return type

*[**GetListMetricsResponse**](getListMetricsResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetricStatistics**
> GetMetricStatisticsResponse GetMetricStatistics(getMetricStatisticsRequest)


A.메트릭 통계 조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getMetricStatisticsRequest** | **[\*GetMetricStatisticsRequest](GetMetricStatisticsRequest.md)** | getMetricStatisticsRequest | 

### Return type

*[**GetMetricStatisticsResponse**](getMetricStatisticsResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

