# \v2Api

All URIs are relative to *https://vpcsearchengine.apigw.ntruss.com/api/v2*

Method | HTTP request                                            | Description
------------- |---------------------------------------------------------| -------------
[**AddNodesInClusterUsingPOST**](v2Api.md#AddNodesInClusterUsingPOST) | **Post** /cluster/addNodesInCluster/{service-group-no}  | Add Nodes In Cluster
[**CreateClusterUsingPOST**](v2Api.md#CreateClusterUsingPOST) | **Post** /cluster/createCluster                         | Create Search Engine Cluster
[**DeleteClusterUsingDELETE**](v2Api.md#DeleteClusterUsingDELETE) | **Delete** /cluster/deleteCluster/{service-group-no}    | Delete Search Engine Cluster
[**GetAcgInfoListUsingGET**](v2Api.md#GetAcgInfoListUsingGET) | **Get** /cluster/getAcgInfoList/{service-group-no}      | Get Search Engine Cluster ACG Information
[**GetClusterInfoUsingGET**](v2Api.md#GetClusterInfoUsingGET) | **Get** /cluster/getClusterInfo/{service-group-no}      | Get Search Engine Cluster
[**GetClusterNodeListUsingGET**](v2Api.md#GetClusterNodeListUsingGET) | **Get** /cluster/getClusterNodeList/{service-group-no}  | Get Search Engine Cluster Node List
[**GetNodeProductListWithGetMethodUsingGET**](v2Api.md#GetNodeProductListWithGetMethodUsingGET) | **Get** /cluster/getNodeProductList                     | Get Node Product List
[**GetOsProductListUsingGET**](v2Api.md#GetOsProductListUsingGET) | **Get** /cluster/getOsProductList                       | Get Software Product List
[**GetSearchEngineVersionListUsingGET**](v2Api.md#GetSearchEngineVersionListUsingGET) | **Get** /cluster/getSearchEngineVersionList             | Get Search Engine Version List
[**ResetSearchEngineUserPasswordUsingPOST**](v2Api.md#resetadminpassworduUsingPost) | **Post** /cluster/resetSearchEngineUserPassword/{service-group-no} | Reset Search Engine User Password
[**RestartClusterUsingGET**](v2Api.md#RestartClusterUsingGET) | **Get** /cluster/restartCluster/{service-group-no}      | Restart Search Engine Cluster


# **AddNodesInClusterUsingPOST**
> ResponseVoboolean AddNodesInClusterUsingPOST(ctx, serviceGroupNo, changeCountOfDataNode)
Add Nodes In Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupNo** | **string**| service-group-no | 
  **changeCountOfDataNode** | [**AddNodesInClusterRequestVo**](AddNodesInClusterRequestVo.md)| changeCountOfDataNode | 

### Return type

[**ResponseVoboolean**](ResponseVo«boolean».md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterUsingPOST**
> ResponseVoCreateClusterResponseVo CreateClusterUsingPOST(ctx, createCluster)
Create Search Engine Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createCluster** | [**CreateClusterRequestVo**](CreateClusterRequestVo.md)| createCluster | 

### Return type

[**ResponseVoCreateClusterResponseVo**](ResponseVo«CreateClusterResponseVo».md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterUsingDELETE**
> ResponseVoboolean DeleteClusterUsingDELETE(ctx, serviceGroupNo)
Delete Search Engine Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupNo** | **string**| service-group-no | 

### Return type

[**ResponseVoboolean**](ResponseVo«boolean».md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAcgInfoListUsingGET**
> ResponseVoGetOpenApiAcgInfoList GetAcgInfoListUsingGET(ctx, serviceGroupNo)
Get Search Engine Cluster ACG Information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupNo** | **string**| service-group-no | 

### Return type

[**ResponseVoGetOpenApiAcgInfoList**](ResponseVo«GetOpenApiAcgInfoList».md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterInfoUsingGET**
> ResponseVoOpenApiGetClusterInfoResponseVo GetClusterInfoUsingGET(ctx, serviceGroupNo)
Get Search Engine Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupNo** | **string**| service-group-no | 

### Return type

[**ResponseVoOpenApiGetClusterInfoResponseVo**](ResponseVo«OpenApiGetClusterInfoResponseVo».md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterNodeListUsingGET**
> ResponseVoGetOpenApiServiceGroupResponseVo GetClusterNodeListUsingGET(ctx, serviceGroupNo)
Get Search Engine Cluster Node List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupNo** | **string**| service-group-no | 

### Return type

[**ResponseVoGetOpenApiServiceGroupResponseVo**](ResponseVo«GetOpenApiServiceGroupResponseVo».md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeProductListWithGetMethodUsingGET**
> ResponseVoGetOpenApiHwProductListResponseVo GetNodeProductListWithGetMethodUsingGET(ctx, optional)
Get Node Product List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***v2ApiGetNodeProductListWithGetMethodUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a v2ApiGetNodeProductListWithGetMethodUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **infraResourceTypeCode** | **optional.String**|  | 
 **softwareTypeCode** | **optional.String**|  | 
 **softwareProductCode** | **optional.String**|  | 
 **memberNo** | **optional.String**|  | 
 **regionNo** | **optional.String**|  | 
 **zoneNo** | **optional.Int32**|  | 
 **subnetNo** | **optional.Int32**|  | 

### Return type

[**ResponseVoGetOpenApiHwProductListResponseVo**](ResponseVo«GetOpenApiHwProductListResponseVo».md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOsProductListUsingGET**
> ResponseVoGetOpenApiOsProductListResponseVo GetOsProductListUsingGET(ctx, )
Get Software Product List

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseVoGetOpenApiOsProductListResponseVo**](ResponseVo«GetOpenApiOsProductListResponseVo».md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSearchEngineVersionListUsingGET**
> ResponseVoGetSearchEngineVersionListResponseVo GetSearchEngineVersionListUsingGET(ctx, )
Get Search Engine Version List

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseVoGetSearchEngineVersionListResponseVo**](ResponseVo«GetSearchEngineVersionListResponseVo».md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartClusterUsingGET**
> ResponseVoboolean RestartClusterUsingGET(ctx, serviceGroupNo)
Restart Search Engine Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupNo** | **string**| service-group-no | 

### Return type

[**ResponseVoboolean**](ResponseVo«boolean».md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

