# \V1Api

All URIs are relative to *https://clouddatastreamingservice.apigw.ntruss.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterChangeCountOfBrokerNodeServiceGroupInstanceNoPost**](V1Api.md#ClusterChangeCountOfBrokerNodeServiceGroupInstanceNoPost) | **Post** /cluster/changeCountOfBrokerNode/{serviceGroupInstanceNo} | 
[**ClusterCreateCDSSClusterPost**](V1Api.md#ClusterCreateCDSSClusterPost) | **Post** /cluster/createCDSSCluster | 
[**ClusterDeleteCDSSClusterServiceGroupInstanceNoDelete**](V1Api.md#ClusterDeleteCDSSClusterServiceGroupInstanceNoDelete) | **Delete** /cluster/deleteCDSSCluster/{serviceGroupInstanceNo} | 
[**ClusterDisableBrokerNodePublicEndpointServiceGroupInstanceNoGet**](V1Api.md#ClusterDisableBrokerNodePublicEndpointServiceGroupInstanceNoGet) | **Get** /cluster/disableBrokerNodePublicEndpoint/{serviceGroupInstanceNo} | 
[**ClusterDisablePublicDomainServiceGroupInstanceNoGet**](V1Api.md#ClusterDisablePublicDomainServiceGroupInstanceNoGet) | **Get** /cluster/disablePublicDomain/{serviceGroupInstanceNo} | 
[**ClusterDownloadCertificateServiceGroupInstanceNoGet**](V1Api.md#ClusterDownloadCertificateServiceGroupInstanceNoGet) | **Get** /cluster/downloadCertificate/{serviceGroupInstanceNo} | 
[**ClusterEnableBrokerNodePublicEndpointServiceGroupInstanceNoPost**](V1Api.md#ClusterEnableBrokerNodePublicEndpointServiceGroupInstanceNoPost) | **Post** /cluster/enableBrokerNodePublicEndpoint/{serviceGroupInstanceNo} | 
[**ClusterEnablePublicDomainServiceGroupInstanceNoGet**](V1Api.md#ClusterEnablePublicDomainServiceGroupInstanceNoGet) | **Get** /cluster/enablePublicDomain/{serviceGroupInstanceNo} | 
[**ClusterGetAcgInfoListServiceGroupInstanceNoGet**](V1Api.md#ClusterGetAcgInfoListServiceGroupInstanceNoGet) | **Get** /cluster/getAcgInfoList/{serviceGroupInstanceNo} | 
[**ClusterGetBrokerInfoServiceGroupInstanceNoGet**](V1Api.md#ClusterGetBrokerInfoServiceGroupInstanceNoGet) | **Get** /cluster/getBrokerInfo/{serviceGroupInstanceNo} | 
[**ClusterGetCDSSVersionListGet**](V1Api.md#ClusterGetCDSSVersionListGet) | **Get** /cluster/getCDSSVersionList | 
[**ClusterGetClusterInfoListPost**](V1Api.md#ClusterGetClusterInfoListPost) | **Post** /cluster/getClusterInfoList | 
[**ClusterGetClusterNodeListServiceGroupInstanceNoGet**](V1Api.md#ClusterGetClusterNodeListServiceGroupInstanceNoGet) | **Get** /cluster/getClusterNodeList/{serviceGroupInstanceNo} | 
[**ClusterGetClusterStatusServiceGroupInstanceNoGet**](V1Api.md#ClusterGetClusterStatusServiceGroupInstanceNoGet) | **Get** /cluster/getClusterStatus/{serviceGroupInstanceNo} | 
[**ClusterGetLoadBalancerInstanceListServiceGroupInstanceNoGet**](V1Api.md#ClusterGetLoadBalancerInstanceListServiceGroupInstanceNoGet) | **Get** /cluster/getLoadBalancerInstanceList/{serviceGroupInstanceNo} | 
[**ClusterGetNodeProductListPost**](V1Api.md#ClusterGetNodeProductListPost) | **Post** /cluster/getNodeProductList | 
[**ClusterGetOsProductListGet**](V1Api.md#ClusterGetOsProductListGet) | **Get** /cluster/getOsProductList | 
[**ClusterGetSubnetListPost**](V1Api.md#ClusterGetSubnetListPost) | **Post** /cluster/getSubnetList | 
[**ClusterGetVpcListGet**](V1Api.md#ClusterGetVpcListGet) | **Get** /cluster/getVpcList | 
[**ClusterResetCMAKPasswordServiceGroupInstanceNoPost**](V1Api.md#ClusterResetCMAKPasswordServiceGroupInstanceNoPost) | **Post** /cluster/resetCMAKPassword/{serviceGroupInstanceNo} | 
[**ClusterRestartAllServicesServiceGroupInstanceNoGet**](V1Api.md#ClusterRestartAllServicesServiceGroupInstanceNoGet) | **Get** /cluster/restartAllServices/{serviceGroupInstanceNo} | 
[**ClusterRestartCMAKServiceServiceGroupInstanceNoGet**](V1Api.md#ClusterRestartCMAKServiceServiceGroupInstanceNoGet) | **Get** /cluster/restartCMAKService/{serviceGroupInstanceNo} | 
[**ClusterRestartKafkaServicePerNodeServiceGroupInstanceNoPost**](V1Api.md#ClusterRestartKafkaServicePerNodeServiceGroupInstanceNoPost) | **Post** /cluster/restartKafkaServicePerNode/{serviceGroupInstanceNo} | 
[**ClusterRestartKafkaServiceServiceGroupInstanceNoGet**](V1Api.md#ClusterRestartKafkaServiceServiceGroupInstanceNoGet) | **Get** /cluster/restartKafkaService/{serviceGroupInstanceNo} | 
[**MonitoringGetCdssMonitoringDataServiceGroupInstanceNoPost**](V1Api.md#MonitoringGetCdssMonitoringDataServiceGroupInstanceNoPost) | **Post** /monitoring/getCdssMonitoringData/{serviceGroupInstanceNo} | 
[**MonitoringGetOsMonitoringDataServiceGroupInstanceNoPost**](V1Api.md#MonitoringGetOsMonitoringDataServiceGroupInstanceNoPost) | **Post** /monitoring/getOsMonitoringData/{serviceGroupInstanceNo} | 


# **ClusterChangeCountOfBrokerNodeServiceGroupInstanceNoPost**
> ResponseVoBoolean ClusterChangeCountOfBrokerNodeServiceGroupInstanceNoPost(ctx, changeCountOfBrokerNode, serviceGroupInstanceNo, xNcpRegionNo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeCountOfBrokerNode** | [**AddNodesInCluster**](AddNodesInCluster.md)| 브로커 노드 추가 파라미터 | 
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNcpRegionNo** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterCreateCDSSClusterPost**
> ResponseVoBoolean ClusterCreateCDSSClusterPost(ctx, createCluster, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createCluster** | [**CreateCluster**](CreateCluster.md)| 클러스터 생성 파라미터 | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterDeleteCDSSClusterServiceGroupInstanceNoDelete**
> ResponseVoBoolean ClusterDeleteCDSSClusterServiceGroupInstanceNoDelete(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


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

# **ClusterDisableBrokerNodePublicEndpointServiceGroupInstanceNoGet**
> ResponseVoBoolean ClusterDisableBrokerNodePublicEndpointServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


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

# **ClusterDisablePublicDomainServiceGroupInstanceNoGet**
> ResponseVoBoolean ClusterDisablePublicDomainServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| X-NCP-REGION_NO | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterDownloadCertificateServiceGroupInstanceNoGet**
> ResponseVoGetCertFileResponseVo ClusterDownloadCertificateServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| X-NCP-REGION_NO | 

### Return type

[**ResponseVoGetCertFileResponseVo**](ResponseVo_GetCertFileResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterEnableBrokerNodePublicEndpointServiceGroupInstanceNoPost**
> ResponseVoBoolean ClusterEnableBrokerNodePublicEndpointServiceGroupInstanceNoPost(ctx, enableBrokerNodePublicEndpoint, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enableBrokerNodePublicEndpoint** | [**EnableBrokerNodePublicEndpoint**](EnableBrokerNodePublicEndpoint.md)| Broker Node Public Endpoint 활성화 | 
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

# **ClusterEnablePublicDomainServiceGroupInstanceNoGet**
> ResponseVoBoolean ClusterEnablePublicDomainServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| X-NCP-REGION_NO | 

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

# **ClusterGetBrokerInfoServiceGroupInstanceNoGet**
> ResponseVoGetBrokerNodeListsResponseVo ClusterGetBrokerInfoServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| X-NCP-REGION_NO | 

### Return type

[**ResponseVoGetBrokerNodeListsResponseVo**](ResponseVo_GetBrokerNodeListsResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetCDSSVersionListGet**
> ResponseVoGetKafkaVersionListResponseVo ClusterGetCDSSVersionListGet(ctx, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetKafkaVersionListResponseVo**](ResponseVo_GetKafkaVersionListResponseVo.md)

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

# **ClusterGetClusterNodeListServiceGroupInstanceNoGet**
> ResponseVoGetOpenApiServiceGroupResponseVo ClusterGetClusterNodeListServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


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

# **ClusterGetClusterStatusServiceGroupInstanceNoGet**
> ResponseVoGetClusterStatusPerNodeResponseVo ClusterGetClusterStatusServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| X-NCP-REGION_NO | 

### Return type

[**ResponseVoGetClusterStatusPerNodeResponseVo**](ResponseVo_GetClusterStatusPerNodeResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetLoadBalancerInstanceListServiceGroupInstanceNoGet**
> ResponseVoGetLoadBalancerListsResponseVo ClusterGetLoadBalancerInstanceListServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetLoadBalancerListsResponseVo**](ResponseVo_GetLoadBalancerListsResponseVo.md)

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
  **getNodeProductList** | [**NodeProduct**](NodeProduct.md)| 매니저, 브로커 노드에서 사용할 수 있는 상품 스펙 조회 파라미터 | 
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
> ResponseVoGetOpenApiVpcSubnetListResponseVo ClusterGetSubnetListPost(ctx, getSubnetList, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getSubnetList** | [**GetSubnetList**](GetSubnetList.md)| Subnet 목록 조회 | 
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetOpenApiVpcSubnetListResponseVo**](ResponseVo_GetOpenApiVpcSubnetListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGetVpcListGet**
> ResponseVoGetOpenApiVpcConfigListResponseVo ClusterGetVpcListGet(ctx, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xNCPREGIONNO** | **int32**| Region No | 

### Return type

[**ResponseVoGetOpenApiVpcConfigListResponseVo**](ResponseVo_GetOpenApiVpcConfigListResponseVo.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterResetCMAKPasswordServiceGroupInstanceNoPost**
> ResponseVoBoolean ClusterResetCMAKPasswordServiceGroupInstanceNoPost(ctx, cmakPassword, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cmakPassword** | [**ResetCmakPassword**](ResetCmakPassword.md)| cmakPassword | 
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

# **ClusterRestartAllServicesServiceGroupInstanceNoGet**
> ResponseVoBoolean ClusterRestartAllServicesServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


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

# **ClusterRestartCMAKServiceServiceGroupInstanceNoGet**
> ResponseVoBoolean ClusterRestartCMAKServiceServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| X-NCP-REGION_NO | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterRestartKafkaServicePerNodeServiceGroupInstanceNoPost**
> ResponseVoBoolean ClusterRestartKafkaServicePerNodeServiceGroupInstanceNoPost(ctx, restartKafkaServicePerNode, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **restartKafkaServicePerNode** | [**RestartKafkaServicePerNode**](RestartKafkaServicePerNode.md)| 노드별 Kafka 서비스 재시작 파라미터 | 
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

# **ClusterRestartKafkaServiceServiceGroupInstanceNoGet**
> ResponseVoBoolean ClusterRestartKafkaServiceServiceGroupInstanceNoGet(ctx, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceGroupInstanceNo** | **string**| serviceGroupInstanceNo | 
  **xNCPREGIONNO** | **int32**| X-NCP-REGION_NO | 

### Return type

[**ResponseVoBoolean**](ResponseVo_boolean.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitoringGetCdssMonitoringDataServiceGroupInstanceNoPost**
> ResponseVoQueryMultipleDataResponseVo MonitoringGetCdssMonitoringDataServiceGroupInstanceNoPost(ctx, getCdssMonitoringData, serviceGroupInstanceNo, xNCPREGIONNO)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getCdssMonitoringData** | [**GetCdssMonitoringData**](GetCdssMonitoringData.md)| CDSS 모니터링 데이터 조회 파라미터 | 
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

