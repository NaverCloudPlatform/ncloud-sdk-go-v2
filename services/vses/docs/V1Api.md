# \V1Api

All URIs are relative to *https://vpcsearchengine.apigw.ntruss.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterChangeCountOfDataNodeServiceGroupInstanceNoPost**](V1Api.md#ClusterChangeCountOfDataNodeServiceGroupInstanceNoPost) | **Post** /cluster/changeCountOfDataNode/{serviceGroupInstanceNo} | 
[**ClusterCreateElasticsearchClusterPost**](V1Api.md#ClusterCreateElasticsearchClusterPost) | **Post** /cluster/createElasticsearchCluster | 
[**ClusterDeleteElasticsearchClusterServiceGroupInstanceNoDelete**](V1Api.md#ClusterDeleteElasticsearchClusterServiceGroupInstanceNoDelete) | **Delete** /cluster/deleteElasticsearchCluster/{serviceGroupInstanceNo} | 
[**ClusterGetAcgInfoListServiceGroupInstanceNoGet**](V1Api.md#ClusterGetAcgInfoListServiceGroupInstanceNoGet) | **Get** /cluster/getAcgInfoList/{serviceGroupInstanceNo} | 
[**ClusterGetClusterDataNodeListServiceGroupInstanceNoGet**](V1Api.md#ClusterGetClusterDataNodeListServiceGroupInstanceNoGet) | **Get** /cluster/getClusterDataNodeList/{serviceGroupInstanceNo} | 
[**ClusterGetClusterInfoListPost**](V1Api.md#ClusterGetClusterInfoListPost) | **Post** /cluster/getClusterInfoList | 
[**ClusterGetElasticsearchVersionListGet**](V1Api.md#ClusterGetElasticsearchVersionListGet) | **Get** /cluster/getElasticsearchVersionList | 
[**ClusterGetLoginKeyListGet**](V1Api.md#ClusterGetLoginKeyListGet) | **Get** /cluster/getLoginKeyList | 
[**ClusterGetNodeProductListPost**](V1Api.md#ClusterGetNodeProductListPost) | **Post** /cluster/getNodeProductList | 
[**ClusterGetOsProductListGet**](V1Api.md#ClusterGetOsProductListGet) | **Get** /cluster/getOsProductList | 
[**ClusterGetSubnetListPost**](V1Api.md#ClusterGetSubnetListPost) | **Post** /cluster/getSubnetList | 
[**ClusterGetVpcListGet**](V1Api.md#ClusterGetVpcListGet) | **Get** /cluster/getVpcList | 
[**ClusterResetKibanaPasswordServiceGroupInstanceNoPost**](V1Api.md#ClusterResetKibanaPasswordServiceGroupInstanceNoPost) | **Post** /cluster/resetKibanaPassword/{serviceGroupInstanceNo} | 
[**ClusterRestartElasticsearchClusterServiceGroupInstanceNoGet**](V1Api.md#ClusterRestartElasticsearchClusterServiceGroupInstanceNoGet) | **Get** /cluster/restartElasticsearchCluster/{serviceGroupInstanceNo} | 
[**DashboardGetDashboardInformationServiceGroupInstanceNoGet**](V1Api.md#DashboardGetDashboardInformationServiceGroupInstanceNoGet) | **Get** /dashboard/getDashboardInformation/{serviceGroupInstanceNo} | 
[**ImportCreateDataImportJobServiceGroupInstanceNoPost**](V1Api.md#ImportCreateDataImportJobServiceGroupInstanceNoPost) | **Post** /import/createDataImportJob/{serviceGroupInstanceNo} | 
[**ImportGetBucketListServiceGroupInstanceNoGet**](V1Api.md#ImportGetBucketListServiceGroupInstanceNoGet) | **Get** /import/getBucketList/{serviceGroupInstanceNo} | 
[**ImportGetDataImportHistoryServiceGroupInstanceNoGet**](V1Api.md#ImportGetDataImportHistoryServiceGroupInstanceNoGet) | **Get** /import/getDataImportHistory/{serviceGroupInstanceNo} | 
[**ImportStopDataImportJobServiceGroupInstanceNoGet**](V1Api.md#ImportStopDataImportJobServiceGroupInstanceNoGet) | **Get** /import/stopDataImportJob/{serviceGroupInstanceNo} | 
[**MonitoringGetElasticsearchMonitoringDataServiceGroupInstanceNoPost**](V1Api.md#MonitoringGetElasticsearchMonitoringDataServiceGroupInstanceNoPost) | **Post** /monitoring/getElasticsearchMonitoringData/{serviceGroupInstanceNo} | 
[**MonitoringGetOsMonitoringDataServiceGroupInstanceNoPost**](V1Api.md#MonitoringGetOsMonitoringDataServiceGroupInstanceNoPost) | **Post** /monitoring/getOsMonitoringData/{serviceGroupInstanceNo} | 
[**SnapshotCreateSnapshotServiceGroupInstanceNoPost**](V1Api.md#SnapshotCreateSnapshotServiceGroupInstanceNoPost) | **Post** /snapshot/createSnapshot/{serviceGroupInstanceNo} | 
[**SnapshotGetBucketListServiceGroupInstanceNoGet**](V1Api.md#SnapshotGetBucketListServiceGroupInstanceNoGet) | **Get** /snapshot/getBucketList/{serviceGroupInstanceNo} | 
[**SnapshotGetSnapshotHistoryServiceGroupInstanceNoGet**](V1Api.md#SnapshotGetSnapshotHistoryServiceGroupInstanceNoGet) | **Get** /snapshot/getSnapshotHistory/{serviceGroupInstanceNo} | 
[**SnapshotGetSnapshotSchedulingHistoryServiceGroupInstanceNoPost**](V1Api.md#SnapshotGetSnapshotSchedulingHistoryServiceGroupInstanceNoPost) | **Post** /snapshot/getSnapshotSchedulingHistory/{serviceGroupInstanceNo} | 
[**SnapshotReleaseSnapshotSchedulingServiceGroupInstanceNoGet**](V1Api.md#SnapshotReleaseSnapshotSchedulingServiceGroupInstanceNoGet) | **Get** /snapshot/releaseSnapshotScheduling/{serviceGroupInstanceNo} | 
[**SnapshotSetSnapshotSchedulingServiceGroupInstanceNoPost**](V1Api.md#SnapshotSetSnapshotSchedulingServiceGroupInstanceNoPost) | **Post** /snapshot/setSnapshotScheduling/{serviceGroupInstanceNo} | 
[**SnapshotUpdateAPIAuthenticationKeyServiceGroupInstanceNoPost**](V1Api.md#SnapshotUpdateAPIAuthenticationKeyServiceGroupInstanceNoPost) | **Post** /snapshot/updateAPIAuthenticationKey/{serviceGroupInstanceNo} | 


# **ClusterChangeCountOfDataNodeServiceGroupInstanceNoPost**
> ResponseVoBoolean ClusterChangeCountOfDataNodeServiceGroupInstanceNoPost(ctx, changeCountOfDataNode, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeCountOfDataNode** | [**AddNodesInCluster**](AddNodesInCluster.md)| 데이터 노드 추가 파라미터 | 
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterCreateElasticsearchClusterPost**
> ResponseVoCreateClusterResponseVo ClusterCreateElasticsearchClusterPost(ctx, createCluster, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createCluster** | [**CreateCluster**](CreateCluster.md)| 클러스터 생성 파라미터 | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoCreateClusterResponseVo**](ResponseVo_CreateClusterResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterDeleteElasticsearchClusterServiceGroupInstanceNoDelete**
> ResponseVoBoolean ClusterDeleteElasticsearchClusterServiceGroupInstanceNoDelete(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetAcgInfoListServiceGroupInstanceNoGet**
> ResponseVoGetOpenApiAcgInfoList ClusterGetAcgInfoListServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetOpenApiAcgInfoList**](ResponseVo_GetOpenApiAcgInfoList.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetClusterDataNodeListServiceGroupInstanceNoGet**
> ResponseVoGetOpenApiServiceGroupResponseVo ClusterGetClusterDataNodeListServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetOpenApiServiceGroupResponseVo**](ResponseVo_GetOpenApiServiceGroupResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetClusterInfoListPost**
> ResponseVoOpenApiGetClusterInfoListResponseVo ClusterGetClusterInfoListPost(ctx, clusterInfoRequest, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterInfoRequest** | [**GetClusterRequest**](GetClusterRequest.md)| 클러스터 조회 시 사용되는 파라미터 | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoOpenApiGetClusterInfoListResponseVo**](ResponseVo_OpenApiGetClusterInfoListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetElasticsearchVersionListGet**
> ResponseVoGetElasticSearchVersionListResponseVo ClusterGetElasticsearchVersionListGet(ctx, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetElasticSearchVersionListResponseVo**](ResponseVo_GetElasticSearchVersionListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetLoginKeyListGet**
> ResponseVoGetOpenApiLoginKeyListResponseVo ClusterGetLoginKeyListGet(ctx, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetOpenApiLoginKeyListResponseVo**](ResponseVo_GetOpenApiLoginKeyListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetNodeProductListPost**
> ResponseVoGetOpenApiHwProductListResponseVo ClusterGetNodeProductListPost(ctx, getNodeProductList, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getNodeProductList** | [**NodeProduct**](NodeProduct.md)| HW 상품 스펙 조회 파라미터 | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetOpenApiHwProductListResponseVo**](ResponseVo_GetOpenApiHwProductListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetOsProductListGet**
> ResponseVoGetOpenApiOsProductListResponseVo ClusterGetOsProductListGet(ctx, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetOpenApiOsProductListResponseVo**](ResponseVo_GetOpenApiOsProductListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetSubnetListPost**
> ResponseVoGetVpcSubnetListResponseVo ClusterGetSubnetListPost(ctx, getSubnetList, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getSubnetList** | [**GetSubnetList**](GetSubnetList.md)| Subnet 목록 조회 | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetVpcSubnetListResponseVo**](ResponseVo_GetVpcSubnetListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetVpcListGet**
> ResponseVoGetVpcConfigListResponseVo ClusterGetVpcListGet(ctx, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetVpcConfigListResponseVo**](ResponseVo_GetVpcConfigListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterResetKibanaPasswordServiceGroupInstanceNoPost**
> ResponseVoBoolean ClusterResetKibanaPasswordServiceGroupInstanceNoPost(ctx, kibanaPassword, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kibanaPassword** | [**ResetKibanaPassword**](ResetKibanaPassword.md)| 키바나 패스워드 변경 | 
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterRestartElasticsearchClusterServiceGroupInstanceNoGet**
> ResponseVoBoolean ClusterRestartElasticsearchClusterServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardGetDashboardInformationServiceGroupInstanceNoGet**
> ResponseVoGetOpenApiDashboardInformationResponseVo DashboardGetDashboardInformationServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 
 **optional** | ***V1ApiDashboardGetDashboardInformationServiceGroupInstanceNoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ApiDashboardGetDashboardInformationServiceGroupInstanceNoGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNo** | **optional.Int32**| pageNo | 
 **pageSize** | **optional.Int32**| pageSize | 

### Return type

[**ResponseVoGetOpenApiDashboardInformationResponseVo**](ResponseVo_GetOpenApiDashboardInformationResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportCreateDataImportJobServiceGroupInstanceNoPost**
> ResponseVoBoolean ImportCreateDataImportJobServiceGroupInstanceNoPost(ctx, createImportJob, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createImportJob** | [**CreateImportJob**](CreateImportJob.md)| Create Import Job 파라미터 | 
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportGetBucketListServiceGroupInstanceNoGet**
> ResponseVoBoolean ImportGetBucketListServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportGetDataImportHistoryServiceGroupInstanceNoGet**
> ResponseVoGetImportHistoryListResponseVo ImportGetDataImportHistoryServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 
 **optional** | ***V1ApiImportGetDataImportHistoryServiceGroupInstanceNoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ApiImportGetDataImportHistoryServiceGroupInstanceNoGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNo** | **optional.Int32**| Page No | 
 **pageSize** | **optional.Int32**| Page Size | 

### Return type

[**ResponseVoGetImportHistoryListResponseVo**](ResponseVo_GetImportHistoryListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportStopDataImportJobServiceGroupInstanceNoGet**
> ResponseVoBoolean ImportStopDataImportJobServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitoringGetElasticsearchMonitoringDataServiceGroupInstanceNoPost**
> ResponseVoQueryMultipleDataResponseVo MonitoringGetElasticsearchMonitoringDataServiceGroupInstanceNoPost(ctx, getElasticsearchMonitoringData, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getElasticsearchMonitoringData** | [**GetElasticsearchMonitoringData**](GetElasticsearchMonitoringData.md)| ES 모니터링 데이터 조회 파라미터 | 
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoQueryMultipleDataResponseVo**](ResponseVo_QueryMultipleDataResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitoringGetOsMonitoringDataServiceGroupInstanceNoPost**
> ResponseVoQueryMultipleDataResponseVo MonitoringGetOsMonitoringDataServiceGroupInstanceNoPost(ctx, getOsMonitoringData, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getOsMonitoringData** | [**GetOsMonitoringData**](GetOsMonitoringData.md)| OS 모니터링 데이터 조회 파라미터 | 
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoQueryMultipleDataResponseVo**](ResponseVo_QueryMultipleDataResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotCreateSnapshotServiceGroupInstanceNoPost**
> ResponseVoBoolean SnapshotCreateSnapshotServiceGroupInstanceNoPost(ctx, createSnapshot, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSnapshot** | [**CreateSnapshot**](CreateSnapshot.md)| Snapshot 생성 파라미터 | 
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotGetBucketListServiceGroupInstanceNoGet**
> ResponseVoGetObjectStorageBucketListResponseVo SnapshotGetBucketListServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetObjectStorageBucketListResponseVo**](ResponseVo_GetObjectStorageBucketListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotGetSnapshotHistoryServiceGroupInstanceNoGet**
> ResponseVoSnapshotHistoryListResponseVo SnapshotGetSnapshotHistoryServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 
 **optional** | ***V1ApiSnapshotGetSnapshotHistoryServiceGroupInstanceNoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ApiSnapshotGetSnapshotHistoryServiceGroupInstanceNoGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNo** | **optional.Int32**| pageNo | 
 **pageSize** | **optional.Int32**| pageSize | 

### Return type

[**ResponseVoSnapshotHistoryListResponseVo**](ResponseVo_SnapshotHistoryListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotGetSnapshotSchedulingHistoryServiceGroupInstanceNoPost**
> ResponseVoGetSnapshotSchedulingHistoryResponseVo SnapshotGetSnapshotSchedulingHistoryServiceGroupInstanceNoPost(ctx, getSnapshotSchedulingHistory, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getSnapshotSchedulingHistory** | [**GetSnapshotSchedulingHistory**](GetSnapshotSchedulingHistory.md)| Snapshot Scheduling 이력 조회 | 
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetSnapshotSchedulingHistoryResponseVo**](ResponseVo_GetSnapshotSchedulingHistoryResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotReleaseSnapshotSchedulingServiceGroupInstanceNoGet**
> ResponseVoBoolean SnapshotReleaseSnapshotSchedulingServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotSetSnapshotSchedulingServiceGroupInstanceNoPost**
> ResponseVoBoolean SnapshotSetSnapshotSchedulingServiceGroupInstanceNoPost(ctx, setSnapshotScheduling, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **setSnapshotScheduling** | [**SetSnapshotScheduling**](SetSnapshotScheduling.md)| Snapshot Scheduling 설정 | 
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotUpdateAPIAuthenticationKeyServiceGroupInstanceNoPost**
> ResponseVoBoolean SnapshotUpdateAPIAuthenticationKeyServiceGroupInstanceNoPost(ctx, updateAPIAuthenticationKey, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateAPIAuthenticationKey** | [**UpdateApiAuthenticationKey**](UpdateApiAuthenticationKey.md)| API 인증키 설정 파라미터 | 
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

