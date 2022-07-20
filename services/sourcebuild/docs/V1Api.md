# \V1Api

All URIs are relative to *https://sourcebuild.apigw.ntruss.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainerRegistry**](V1Api.md#GetContainerRegistry) | **Get** /containerregistry/registry | 
[**GetComputeEnv**](V1Api.md#GetComputeEnv) | **Get** /env/compute | 
[**GetDockerEnv**](V1Api.md#GetDockerEnv) | **Get** /env/docker | 
[**GetOsEnv**](V1Api.md#GetOsEnv) | **Get** /env/os | 
[**GetRuntimeEnv**](V1Api.md#GetRuntimeEnv) | **Get** /env/os/{osId}/runtime | 
[**GetRuntimeVersionEnv**](V1Api.md#GetRuntimeVersionEnv) | **Get** /env/os/{osId}/runtime/{runtimeId}/version | 
[**GetObjectstorageBucket**](V1Api.md#GetObjectstorageBucket) | **Get** /objectstorage/bucket | 
[**GetProjects**](V1Api.md#GetProjects) | **Get** /project | 
[**CreateProject**](V1Api.md#CreateProject) | **Post** /project | 
[**CancelBuild**](V1Api.md#CancelBuild) | **Delete** /project/{projectId}/build | 
[**StartBuild**](V1Api.md#StartBuild) | **Post** /project/{projectId}/build | 
[**DeleteProject**](V1Api.md#DeleteProject) | **Delete** /project/{projectId} | 
[**GetProject**](V1Api.md#GetProject) | **Get** /project/{projectId} | 
[**GetBuildHistory**](V1Api.md#GetBuildHistory) | **Get** /project/{projectId}/history | 
[**ChangeProject**](V1Api.md#ChangeProject) | **Patch** /project/{projectId} | 
[**GetSourcecommitRepositories**](V1Api.md#GetSourcecommitRepositories) | **Get** /sourcecommit/repository | 
[**GetSourcecommitRepositoryBranches**](V1Api.md#GetSourcecommitRepositoryBranches) | **Get** /sourcecommit/repository/{repositoryName}/branch | 


# **GetContainerRegistry**
> GetContainerRegistryResponse GetContainerRegistry()


### Required Parameters

This endpoint does not need any parameter.

### Return type

*[**GetContainerRegistryResponse**](GetContainerRegistryResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComputeEnv**
> GetComputeEnvResponse GetComputeEnv()


### Required Parameters

This endpoint does not need any parameter.

### Return type

*[**GetComputeEnvResponse**](GetComputeEnvResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDockerEnv**
> GetDockerEnvResponse GetDockerEnv()


### Required Parameters

This endpoint does not need any parameter.

### Return type

*[**GetDockerEnvResponse**](GetDockerEnvResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOsEnv**
> GetOsEnvResponse GetOsEnv()


### Required Parameters

This endpoint does not need any parameter.

### Return type

*[**GetOsEnvResponse**](GetOsEnvResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuntimeEnv**
> GetRuntimeEnvResponse GetRuntimeEnv(osId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**osId** | **string** | osId |

### Return type

*[**GetRuntimeEnvResponse**](GetRuntimeEnvResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuntimeVersionEnv**
> GetRuntimeVersionEnvResponse GetRuntimeVersionEnv(osId, runtimeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**osId** | **string** | osId | **runtimeId** | **string** | runtimeId |

### Return type

*[**GetRuntimeVersionEnvResponse**](GetRuntimeVersionEnvResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectstorageBucket**
> GetObjsBucketResponse GetObjectstorageBucket()


### Required Parameters

This endpoint does not need any parameter.

### Return type

*[**GetObjsBucketResponse**](GetObjsBucketResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjects**
> GetProjectListResponse GetProjects(optional)


### Optional Parameters
Optional parameters are passed through a map[string]interface{}.
The keys are pageNo, pageSize and projectName.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pageNo** | **string** | Page No |
**pageSize** | **string** | Page Size |
**projectName** | **string** | Project Name |

### Return type

*[**GetProjectListResponse**](GetProjectListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**
> CreateProjectResponse CreateProject(body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*CreateProject](CreateProject.md)** |  |

### Return type

*[**CreateProjectResponse**](CreateProjectResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelBuild**
> CancelBuild(body, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*CancelBuild](CancelBuild.md)** |  | **projectId** | **string** | projectId |

### Return type

 (empty response body)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartBuild**
> StartBuildResponse StartBuild(projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId |

### Return type

*[**StartBuildResponse**](StartBuildResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject(projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId |

### Return type

 (empty response body)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> GetProjectDetailResponse GetProject(projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId |

### Return type

*[**GetProjectDetailResponse**](GetProjectDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildHistory**
> GetBuildHistoryResponse GetBuildHistory(projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId |

### Return type

*[**GetBuildHistoryResponse**](GetBuildHistoryResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeProject**
> CreateProjectResponse ChangeProject(body, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*ChangeProject](ChangeProject.md)** |  | **projectId** | **string** | projectId |

### Return type

*[**CreateProjectResponse**](CreateProjectResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourcecommitRepositories**
> GetScRepositoryResponse GetSourcecommitRepositories()


### Required Parameters

This endpoint does not need any parameter.

### Return type

*[**GetScRepositoryResponse**](GetScRepositoryResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourcecommitRepositoryBranches**
> GetScBranchResponse GetSourcecommitRepositoryBranches(repositoryName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**repositoryName** | **string** | repositoryName | 

### Return type

*[**GetScBranchResponse**](GetScBranchResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

