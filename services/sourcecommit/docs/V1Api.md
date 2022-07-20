# \V1Api

All URIs are relative to *https://sourcecommit.apigw.ntruss.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRepositories**](V1Api.md#GetRepositories) | **Get** /repository | 
[**GetRepository**](V1Api.md#GetRepository) | **Get** /repository/{repositoryName} | 
[**GetRepositoryById**](V1Api.md#GetRepositoryById) | **Get** /repository/{repositoryId} | 
[**CreateRepository**](V1Api.md#CreateRepository) | **Post** /repository | 
[**DeleteRepository**](V1Api.md#DeleteRepository) | **Delete** /repository/id/{repositoryId} | 
[**ChangeRepository**](V1Api.md#ChangeRepository) | **Patch** /repository/id/{repositoryId} | 

# **GetRepositories**
> GetRepositoryListResponse GetRepositories()


### Required Parameters

This endpoint does not need any parameter.



### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pageNo** | **int32** | page No | 
**pageSize** | **int32** | page Size|
 **projectName** | **string** | project Name | 

### Return type

*[**GetRepositoryListResponse**](GetRepositoryListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# **GetRepository**
> GetRepositoryDetailResponse GetRepository(repositoryName)


### Required Parameters
This endpoint does not need any parameter.


### Return type

*[**GetRepositoryDetailResponse**](GetRepositoryDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryById**
> GetRepositoryDetailResponse GetRepositoryById(repositoryId)


### Required Parameters
This endpoint does not need any parameter.


### Return type

*[**GetRepositoryDetailResponse**](GetRepositoryDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository**
> bool RepositoryPost(body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*CreateRepository](CreateRepository.md)** |  |

### Return type

**bool**

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# **DeleteRepository**
> bool DeleteRepository(repositoryId)


### Required Parameters

This endpoint does not need any parameter.


### Return type

**bool**

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# **ChangeRepository**
> bool ChangeRepository(body, repositoryId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*ChangeRepository](ChangeRepository.md)** |  |

### Return type

**bool**

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
