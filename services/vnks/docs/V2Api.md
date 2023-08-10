# \V2Api

All URIs are relative to *https://nks.apigw.ntruss.com/vnks/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersGet**](V2Api.md#ClustersGet) | **Get** /clusters | 
[**ClustersPost**](V2Api.md#ClustersPost) | **Post** /clusters | 
[**ClustersUuidAddSubnetPatch**](V2Api.md#ClustersUuidAddSubnetPatch) | **Patch** /clusters/{uuid}/add-subnet | 
[**ClustersUuidDelete**](V2Api.md#ClustersUuidDelete) | **Delete** /clusters/{uuid} | 
[**ClustersUuidGet**](V2Api.md#ClustersUuidGet) | **Get** /clusters/{uuid} | 
[**ClustersUuidIpAclGet**](V2Api.md#ClustersUuidIpAclGet) | **Get** /clusters/{uuid}/ip-acl | 
[**ClustersUuidIpAclPatch**](V2Api.md#ClustersUuidIpAclPatch) | **Patch** /clusters/{uuid}/ip-acl | 
[**ClustersUuidKubeconfigGet**](V2Api.md#ClustersUuidKubeconfigGet) | **Get** /clusters/{uuid}/kubeconfig | 
[**ClustersUuidKubeconfigResetPatch**](V2Api.md#ClustersUuidKubeconfigResetPatch) | **Patch** /clusters/{uuid}/kubeconfig/reset | 
[**ClustersUuidLbSubnetPatch**](V2Api.md#ClustersUuidLbSubnetPatch) | **Patch** /clusters/{uuid}/lb-subnet | 
[**ClustersUuidLogPatch**](V2Api.md#ClustersUuidLogPatch) | **Patch** /clusters/{uuid}/log | 
[**ClustersUuidNodePoolGet**](V2Api.md#ClustersUuidNodePoolGet) | **Get** /clusters/{uuid}/node-pool | 
[**ClustersUuidNodePoolInstanceNoDelete**](V2Api.md#ClustersUuidNodePoolInstanceNoDelete) | **Delete** /clusters/{uuid}/node-pool/{instanceNo} | 
[**ClustersUuidNodePoolInstanceNoLabelsPut**](V2Api.md#ClustersUuidNodePoolInstanceNoLabelsPut) | **Put** /clusters/{uuid}/node-pool/{instanceNo}/labels | 
[**ClustersUuidNodePoolInstanceNoPatch**](V2Api.md#ClustersUuidNodePoolInstanceNoPatch) | **Patch** /clusters/{uuid}/node-pool/{instanceNo} | 
[**ClustersUuidNodePoolInstanceNoSubnetsPatch**](V2Api.md#ClustersUuidNodePoolInstanceNoSubnetsPatch) | **Patch** /clusters/{uuid}/node-pool/{instanceNo}/subnets | 
[**ClustersUuidNodePoolInstanceNoTaintsPut**](V2Api.md#ClustersUuidNodePoolInstanceNoTaintsPut) | **Put** /clusters/{uuid}/node-pool/{instanceNo}/taints | 
[**ClustersUuidNodePoolInstanceNoUpgradePatch**](V2Api.md#ClustersUuidNodePoolInstanceNoUpgradePatch) | **Patch** /clusters/{uuid}/node-pool/{instanceNo}/upgrade | 
[**ClustersUuidNodePoolPost**](V2Api.md#ClustersUuidNodePoolPost) | **Post** /clusters/{uuid}/node-pool | 
[**ClustersUuidNodesGet**](V2Api.md#ClustersUuidNodesGet) | **Get** /clusters/{uuid}/nodes | 
[**ClustersUuidNodesInstanceNoDelete**](V2Api.md#ClustersUuidNodesInstanceNoDelete) | **Delete** /clusters/{uuid}/nodes/{instanceNo} | 
[**ClustersUuidOidcGet**](V2Api.md#ClustersUuidOidcGet) | **Get** /clusters/{uuid}/oidc | 
[**ClustersUuidOidcPatch**](V2Api.md#ClustersUuidOidcPatch) | **Patch** /clusters/{uuid}/oidc | 
[**ClustersUuidUpgradePatch**](V2Api.md#ClustersUuidUpgradePatch) | **Patch** /clusters/{uuid}/upgrade | 
[**OptionServerImageGet**](V2Api.md#OptionServerImageGet) | **Get** /option/server-image | 
[**OptionServerProductCodeGet**](V2Api.md#OptionServerProductCodeGet) | **Get** /option/server-product-code | 
[**OptionVersionGet**](V2Api.md#OptionVersionGet) | **Get** /option/version | 
[**RootGet**](V2Api.md#RootGet) | **Get** / | 


# **ClustersGet**
> ClustersRes ClustersGet()


### Required Parameters
This endpoint does not need any parameter.

### Return type

*[**ClustersRes**](ClustersRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersPost**
> CreateClusterRes ClustersPost(body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*ClusterInputBody](ClusterInputBody.md)** |  | 

### Return type

*[**CreateClusterRes**](CreateClusterRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidAddSubnetPatch**
> UpdateClusterRes ClustersUuidAddSubnetPatch(body, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*AddSubnetDto](AddSubnetDto.md)** |  | **uuid** | **string** | uuid | 

### Return type

*[**UpdateClusterRes**](UpdateClusterRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidDelete**
> ClustersUuidDelete(uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | 

### Return type

 (empty response body)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidGet**
> ClusterRes ClustersUuidGet(uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | 

### Return type

*[**ClusterRes**](ClusterRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidIpAclGet**
> IpAclsRes ClustersUuidIpAclGet(uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | 

### Return type

*[**IpAclsRes**](IpAclsRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidIpAclPatch**
> UpdateClusterRes ClustersUuidIpAclPatch(body, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*IpAclsDto](IpAclsDto.md)** |  | **uuid** | **string** | uuid | 

### Return type

*[**UpdateClusterRes**](UpdateClusterRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidKubeconfigGet**
> KubeconfigRes ClustersUuidKubeconfigGet(uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | 

### Return type

*[**KubeconfigRes**](KubeconfigRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidKubeconfigResetPatch**
> UpdateClusterRes ClustersUuidKubeconfigResetPatch(uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | 

### Return type

*[**UpdateClusterRes**](UpdateClusterRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidLbSubnetPatch**
> UpdateClusterLbSubnetRes ClustersUuidLbSubnetPatch(uuid, lbSubnetNo, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | **lbSubnetNo** | **int32** |  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | **lbSubnetNo** | **int32** |  | **igwYn** | **string** |  | 

### Return type

*[**UpdateClusterLbSubnetRes**](UpdateClusterLbSubnetRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidLogPatch**
> UpdateClusterRes ClustersUuidLogPatch(body, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*AuditLogDto](AuditLogDto.md)** |  | **uuid** | **string** | uuid | 

### Return type

*[**UpdateClusterRes**](UpdateClusterRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidNodePoolGet**
> NodePoolRes ClustersUuidNodePoolGet(uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | 

### Return type

*[**NodePoolRes**](NodePoolRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidNodePoolInstanceNoDelete**
> ClustersUuidNodePoolInstanceNoDelete(uuid, instanceNo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | **instanceNo** | **string** | instanceNo | 

### Return type

 (empty response body)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidNodePoolInstanceNoLabelsPut**
> UpdateNodePoolRes ClustersUuidNodePoolInstanceNoLabelsPut(body, uuid, instanceNo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*UpdateNodepoolLabelDto](UpdateNodepoolLabelDto.md)** |  | **uuid** | **string** | uuid | **instanceNo** | **string** | instanceNo | 

### Return type

*[**UpdateNodePoolRes**](UpdateNodePoolRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidNodePoolInstanceNoPatch**
> ClustersUuidNodePoolInstanceNoPatch(body, uuid, instanceNo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*NodePoolUpdateBody](NodePoolUpdateBody.md)** |  | **uuid** | **string** | uuid | **instanceNo** | **string** | instanceNo | 

### Return type

 (empty response body)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidNodePoolInstanceNoSubnetsPatch**
> UpdateNodePoolRes ClustersUuidNodePoolInstanceNoSubnetsPatch(body, uuid, instanceNo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*UpdateNodepoolSubnetDto](UpdateNodepoolSubnetDto.md)** |  | **uuid** | **string** | uuid | **instanceNo** | **string** | instanceNo | 

### Return type

*[**UpdateNodePoolRes**](UpdateNodePoolRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidNodePoolInstanceNoTaintsPut**
> UpdateNodePoolRes ClustersUuidNodePoolInstanceNoTaintsPut(body, uuid, instanceNo)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*UpdateNodepoolTaintDto](UpdateNodepoolTaintDto.md)** |  | **uuid** | **string** | uuid | **instanceNo** | **string** | instanceNo | 

### Return type

*[**UpdateNodePoolRes**](UpdateNodePoolRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidNodePoolInstanceNoUpgradePatch**
> UpdateNodePoolRes ClustersUuidNodePoolInstanceNoUpgradePatch(uuid, instanceNo, k8sVersion, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | **instanceNo** | **string** | instanceNo | **k8sVersion** | **string** |  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | **instanceNo** | **string** | instanceNo | **k8sVersion** | **string** |  | **maxSurge** | **int32** |  | **maxUnavailable** | **int32** |  | 

### Return type

*[**UpdateNodePoolRes**](UpdateNodePoolRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidNodePoolPost**
> UpdateClusterRes ClustersUuidNodePoolPost(body, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*NodePoolCreationBody](NodePoolCreationBody.md)** |  | **uuid** | **string** | uuid | 

### Return type

*[**UpdateClusterRes**](UpdateClusterRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidNodesGet**
> WorkerNodeRes ClustersUuidNodesGet(uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | 

### Return type

*[**WorkerNodeRes**](WorkerNodeRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidNodesInstanceNoDelete**
> ClustersUuidNodesInstanceNoDelete(uuid, instanceNo, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | **instanceNo** | **string** | instanceNo | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | **instanceNo** | **string** | instanceNo | **nodePoolId** | **string** | 노드풀 인스턴스 번호 | 

### Return type

 (empty response body)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidOidcGet**
> OidcRes ClustersUuidOidcGet(uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | 

### Return type

*[**OidcRes**](OidcRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidOidcPatch**
> UpdateClusterRes ClustersUuidOidcPatch(body, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*UpdateOidcDto](UpdateOidcDto.md)** |  | **uuid** | **string** | uuid | 

### Return type

*[**UpdateClusterRes**](UpdateClusterRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUuidUpgradePatch**
> UpdateClusterRes ClustersUuidUpgradePatch(uuid, k8sVersion, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | **k8sVersion** | **string** |  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uuid** | **string** | uuid | **k8sVersion** | **string** |  | **maxSurge** | **int32** |  | **maxUnavailable** | **int32** |  | 

### Return type

*[**UpdateClusterRes**](UpdateClusterRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptionServerImageGet**
> OptionsRes OptionServerImageGet(optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**hypervisorCode** | **string** |  | 

### Return type

*[**OptionsRes**](OptionsRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptionServerProductCodeGet**
> OptionsResForServerProduct OptionServerProductCodeGet(softwareCode, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**softwareCode** | **string** |  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**softwareCode** | **string** |  | **hypervisorCode** | **string** |  | **zoneCode** | **string** |  | **zoneNo** | **string** |  | 

### Return type

*[**OptionsResForServerProduct**](OptionsResForServerProduct.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptionVersionGet**
> OptionsRes OptionVersionGet(optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**from** | **string** |  | **to** | **string** |  | 

### Return type

*[**OptionsRes**](OptionsRes.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RootGet**
> RootGet()


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

