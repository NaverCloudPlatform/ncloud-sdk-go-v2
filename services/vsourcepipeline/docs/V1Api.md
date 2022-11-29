# \V1Api

All URIs are relative to *https://vpcsourcepipeline.apigw.ntruss.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjects**](V1Api.md#GetProjects) | **Get** /project | 
[**CreateProject**](V1Api.md#CreateProject) | **Post** /project | 
[**DeleteProject**](V1Api.md#DeleteProject) | **Delete** /project/{projectId} | 
[**StartProject**](V1Api.md#StartProject) | **Post** /project/{projectId}/do | 
[**GetProject**](V1Api.md#GetProject) | **Get** /project/{projectId} | 
[**GetProjectHistories**](V1Api.md#GetProjectHistories) | **Get** /project/{projectId}/history | 
[**CancelProject**](V1Api.md#CancelProject) | **Post** /project/{projectId}/history/{historyId}/cancel | 
[**GetProjectHistory**](V1Api.md#GetProjectHistory) | **Get** /project/{projectId}/history/{historyId} | 
[**ChangeProject**](V1Api.md#ChangeProject) | **Patch** /project/{projectId} | 
[**GetSourcebuildProjects**](V1Api.md#GetSourcebuildProjects) | **Get** /sourcebuild/project | 
[**GetSourcecommitRepositories**](V1Api.md#GetSourcecommitRepositories) | **Get** /sourcecommit/repository | 
[**GetSourcecommitRepositoryBranches**](V1Api.md#GetSourcecommitRepositoryBranches) | **Get** /sourcecommit/repository/{repositoryName}/branch | 
[**GetSourcedeployProjects**](V1Api.md#GetSourcedeployProjects) | **Get** /sourcedeploy/project | 
[**GetSourcedeployProjectStages**](V1Api.md#GetSourcedeployProjectStages) | **Get** /sourcedeploy/project/{projectId}/stage | 
[**GetSourcedeployProjectScenarios**](V1Api.md#GetSourcedeployProjectScenarios) | **Get** /sourcedeploy/project/{projectId}/stage/{stageId}/scenario | 
[**GetTimeZone**](V1Api.md#GetTimeZone) | **Get** /trigger/timezone | 


# **GetProjects**
> GetProjectListResponse GetProjects

### Return type

*[**GetProjectListResponse**](GetProjectListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**
> CreateProjectResponse CreateProject(createProject)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createProject** | **[\*CreateProject](CreateProject.md)** |  

### Return type

*[**CreateProjectResponse**](CreateProjectResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProjectResponse DeleteProject(projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId 

### Return type

*[**DeleteProjectResponse**](DeleteProjectResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartProject**
> StartPipelineResponse StartProject(projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId 

### Return type

*[**StartPipelineResponse**](StartPipelineResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> GetProjectDetailResponse GetProject(projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId 

### Return type

*[**GetProjectDetailResponse**](GetProjectDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectHistories**
> GetHistoryListResponse GetProjectHistories(projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId 

### Return type

*[**GetHistoryListResponse**](GetHistoryListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelProject**
> CancelPipelineResponse CancelProject(projectId, historyId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId | **historyId** | **string** | historyId 

### Return type

*[**CancelPipelineResponse**](CancelPipelineResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectHistory**
> GetHistoryDetailResponse GetProjectHistory(projectId, historyId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId | **historyId** | **string** | historyId 

### Return type

*[**GetHistoryDetailResponse**](GetHistoryDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeProject**
> ChangeProjectReponse ChangeProject(changeProject, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**changeProject** | **[\*ChangeProject](ChangeProject.md)** |  | **projectId** | **string** | projectId 

### Return type

*[**ChangeProjectReponse**](ChangeProjectReponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourcebuildProjects**
> GetSbProjectResponse GetSourcebuildProjects

### Return type

*[**GetSbProjectResponse**](GetSbProjectResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourcecommitRepositories**
> GetScRepositoryResposne GetSourcecommitRepositories

### Return type

*[**GetScRepositoryResposne**](GetScRepositoryResposne.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourcecommitRepositoryBranches**
> GetScBranchResponse GetSourcecommitRepositoryBranches(repositoryName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**repositoryName** | **string** | repositoryName 

### Return type

*[**GetScBranchResponse**](GetScBranchResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourcedeployProjects**
> GetSdProjectResponse GetSourcedeployProjects

### Return type

*[**GetSdProjectResponse**](GetSdProjectResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourcedeployProjectStages**
> GetSdStageResponse GetSourcedeployProjectStages(projectId)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId 

### Return type

*[**GetSdStageResponse**](GetSdStageResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourcedeployProjectScenarios**
> GetSdScenarioResponse GetSourcedeployProjectScenarios(projectId, stageId)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**projectId** | **string** | projectId | **stageId** | **string** | stageId 

### Return type

*[**GetSdScenarioResponse**](GetSdScenarioResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeZone**
> GetTimeZone GetTimeZone

### Return type

*[**GetTimeZone**](GetTimeZone.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
