# \V1Api

All URIs are relative to *https://vpcsourcedeploy.apigw.ntruss.com/api/v1*

| Method                                                                        | HTTP request                                                                       | Description |
| ----------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ----------- |
| [**GetAutoscalingGroupTargetGroup**](V1Api.md#GetAutoscalingGroupTargetGroup) | **Get** /autoscaling/{autoscalingName}                                             |
| [**GetAutoscalingGroups**](V1Api.md#GetAutoscalingGroups)                     | **Get** /autoscaling                                                               |
| [**GetKubernetesClusters**](V1Api.md#GetKubernetesClusters)                   | **Get** /kubernetes/cluster                                                        |
| [**GetObjectstorageObjects**](V1Api.md#GetObjectstorageObjects)               | **Get** /objectstorage/bucket/{bucketName}                                         |
| [**GetObjectstorageBuckets**](V1Api.md#GetObjectstorageBuckets)               | **Get** /objectstorage/bucket                                                      |
| [**GetProjects**](V1Api.md#GetProjects)                                       | **Get** /project                                                                   |
| [**CreateProject**](V1Api.md#CreateProject)                                   | **Post** /project                                                                  |
| [**DeleteProject**](V1Api.md#DeleteProject)                                   | **Delete** /project/{projectId}                                                    |
| [**GetHistories**](V1Api.md#GetHistories)                                     | **Get** /project/{projectId}/history                                               |
| [**AcceptDeployApproval**](V1Api.md#AcceptDeployApproval)                     | **Post** /project/{projectId}/history/{historyId}/approval/accept                  |
| [**RejectDeployApproval**](V1Api.md#RejectDeployApproval)                     | **Post** /project/{projectId}/history/{historyId}/approval/reject                  |
| [**AcceptDeployCanary**](V1Api.md#AcceptDeployCanary)                         | **Post** /project/{projectId}/history/{historyId}/canary/accept                    |
| [**RejectDeployCanary**](V1Api.md#RejectDeployCanary)                         | **Post** /project/{projectId}/history/{historyId}/canary/reject                    |
| [**CancelDeploy**](V1Api.md#CancelDeploy)                                     | **Post** /project/{projectId}/history/{historyId}/cancel                           |
| [**GetHistory**](V1Api.md#GetHistory)                                         | **Get** /project/{projectId}/history/{historyId}                                   |
| [**GetCanaryReport**](V1Api.md#GetCanaryReport)                               | **Get** /project/{projectId}/history/{historyId}/report/{endtime}                  |
| [**GetCanaryReportEndtime**](V1Api.md#GetCanaryReportEndtime)                 | **Get** /project/{projectId}/history/{historyId}/report                            |
| [**GetStages**](V1Api.md#GetStages)                                           | **Get** /project/{projectId}/stage                                                 |
| [**CreateStage**](V1Api.md#CreateStage)                                       | **Post** /project/{projectId}/stage                                                |
| [**DeleteStage**](V1Api.md#DeleteStage)                                       | **Delete** /project/{projectId}/stage/{stageId}                                    |
| [**GetStage**](V1Api.md#GetStage)                                             | **Get** /project/{projectId}/stage/{stageId}                                       |
| [**ChangeStage**](V1Api.md#ChangeStage)                                       | **Patch** /project/{projectId}/stage/{stageId}                                     |
| [**GetScenarioes**](V1Api.md#GetScenarioes)                                   | **Get** /project/{projectId}/stage/{stageId}/scenario                              |
| [**CreateScenario**](V1Api.md#CreateScenario)                                 | **Post** /project/{projectId}/stage/{stageId}/scenario                             |
| [**DeleteScenario**](V1Api.md#DeleteScenario)                                 | **Delete** /project/{projectId}/stage/{stageId}/scenario/{scenarioId}              |
| [**Deploy**](V1Api.md#Deploy)                                                 | **Post** /project/{projectId}/stage/{stageId}/scenario/{scenarioId}/deploy         |
| [**DeployRequest**](V1Api.md#DeployRequest)                                   | **Post** /project/{projectId}/stage/{stageId}/scenario/{scenarioId}/deploy/request |
| [**GetScenario**](V1Api.md#GetScenario)                                       | **Get** /project/{projectId}/stage/{stageId}/scenario/{scenarioId}                 |
| [**ChangeScenario**](V1Api.md#ChangeScenario)                                 | **Patch** /project/{projectId}/stage/{stageId}/scenario/{scenarioId}               |
| [**GetServers**](V1Api.md#GetServers)                                         | **Get** /server                                                                    |
| [**GetSourcebuildProjects**](V1Api.md#GetSourcebuildProjects)                 | **Get** /sourcebuild/project                                                       |
| [**GetSourcecommitRepositories**](V1Api.md#GetSourcecommitRepositories)       | **Get** /sourcecommit/repository                                                   |
| [**GetSourceCommitBranches**](V1Api.md#GetSourceCommitBranches)               | **Get** /sourcecommit/repository/{repositoryName}/branch                           |

# **GetAutoscalingGroupTargetGroup**

> GetTargetGroupListResponse GetAutoscalingGroupTargetGroup(autoscalingName)

### Required Parameters

| Name                | Type       | Description     | Notes |
| ------------------- | ---------- | --------------- | ----- |
| **autoscalingName** | **string** | autoscalingName |

### Return type

\*[**GetTargetGroupListResponse**](GetTargetGroupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoscalingGroups**

> GetAutoScalingGroupListResponse GetAutoscalingGroups()

### Required Parameters

This endpoint does not need any parameter.

### Return type

\*[**GetAutoScalingGroupListResponse**](GetAutoScalingGroupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKubernetesClusters**

> GetKubernetesServiceClusterListResponse GetKubernetesClusters()

### Required Parameters

This endpoint does not need any parameter.

### Return type

\*[**GetKubernetesServiceClusterListResponse**](GetKubernetesServiceClusterListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectstorageObjects**

> GetObjectStorageObjectListResponse GetObjectstorageObjects(bucketName)

### Required Parameters

| Name           | Type       | Description | Notes |
| -------------- | ---------- | ----------- | ----- |
| **bucketName** | **string** | bucketName  |

### Return type

\*[**GetObjectStorageObjectListResponse**](GetObjectStorageObjectListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectstorageBuckets**

> GetObjectStorageBucketListResponse GetObjectstorageBuckets()

### Required Parameters

This endpoint does not need any parameter.

### Return type

\*[**GetObjectStorageBucketListResponse**](GetObjectStorageBucketListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjects**

> GetProjectListResponse GetProjects(optional)

### Required Parameters

| Name         | Type                       | Description         | Notes                |
| ------------ | -------------------------- | ------------------- | -------------------- |
| **optional** | **map[string]interface{}** | optional parameters | nil if no parameters |

### Optional Parameters

Optional parameters are passed through a map[string]interface{}.

| Name            | Type       | Description | Notes |
| --------------- | ---------- | ----------- | ----- |
| **pageNo**      | **string** |
| **pageSize**    | **string** |
| **projectName** | **string** |

### Return type

\*[**GetProjectListResponse**](GetProjectListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**

> CreateUpdateResponse CreateProject(projectCreate)

### Required Parameters

| Name              | Type                                    | Description   | Notes |
| ----------------- | --------------------------------------- | ------------- | ----- |
| **projectCreate** | **[\*CreateProject](CreateProject.md)** | 프로젝트 생성 |

### Return type

\*[**CreateUpdateResponse**](CreateUpdateResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**

> OkResponse DeleteProject(projectId)

### Required Parameters

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |

### Return type

\*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistories**

> GetDeployHistoryListResponse GetHistories(projectId, optional)

### Required Parameters

| Name          | Type                       | Description         | Notes                |
| ------------- | -------------------------- | ------------------- | -------------------- |
| **projectId** | **string**                 | projectId           |
| **optional**  | **map[string]interface{}** | optional parameters | nil if no parameters |

### Optional Parameters

Optional parameters are passed through a map[string]interface{}.

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |
| **pageNo**    | **string** | 페이지 NO   |
| **pageSize**  | **string** | 페이지 Size |

### Return type

\*[**GetDeployHistoryListResponse**](GetDeployHistoryListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AcceptDeployApproval**

> OkResponse AcceptDeployApproval(projectId, historyId)

### Required Parameters

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |
| **historyId** | **string** | historyId   |

### Return type

\*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RejectDeployApproval**

> OkResponse RejectDeployApproval(projectId, historyId)

### Required Parameters

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |
| **historyId** | **string** | historyId   |

### Return type

\*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AcceptDeployCanary**

> OkResponse AcceptDeployCanary(projectId, historyId)

### Required Parameters

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |
| **historyId** | **string** | historyId   |

### Return type

\*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RejectDeployCanary**

> OkResponse RejectDeployCanary(projectId, historyId)

### Required Parameters

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |
| **historyId** | **string** | historyId   |

### Return type

\*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelDeploy**

> OkResponse CancelDeploy(projectId, historyId)

### Required Parameters

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |
| **historyId** | **string** | historyId   |

### Return type

\*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistory**

> GetDeployHistoryDetailResponse GetHistory(projectId, historyId)

### Required Parameters

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |
| **historyId** | **string** | historyId   |

### Return type

\*[**GetDeployHistoryDetailResponse**](GetDeployHistoryDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCanaryReport**

> GetCanaryAnalysisReportResponse GetCanaryReport(projectId, historyId, endtime)

### Required Parameters

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |
| **historyId** | **string** | historyId   |
| **endtime**   | **string** | endtime     |

### Return type

\*[**GetCanaryAnalysisReportResponse**](GetCanaryAnalysisReportResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCanaryReportEndtime**

> GetCanaryAnalysisStageListResponse GetCanaryReportEndtime(projectId, historyId)

### Required Parameters

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |
| **historyId** | **string** | historyId   |

### Return type

\*[**GetCanaryAnalysisStageListResponse**](GetCanaryAnalysisStageListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStages**

> GetStageListResponse GetStages(projectId, optional)

### Required Parameters

| Name          | Type                       | Description         | Notes                |
| ------------- | -------------------------- | ------------------- | -------------------- |
| **projectId** | **string**                 | projectId           |
| **optional**  | **map[string]interface{}** | optional parameters | nil if no parameters |

### Optional Parameters

Optional parameters are passed through a map[string]interface{}.

| Name          | Type       | Description         | Notes |
| ------------- | ---------- | ------------------- | ----- |
| **projectId** | **string** | projectId           |
| **pageNo**    | **string** | 페이지 NO           |
| **pageSize**  | **string** | 페이지 Size         |
| **stageName** | **string** | Stage 이름으로 검색 |

### Return type

\*[**GetStageListResponse**](GetStageListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStage**

> CreateUpdateResponse CreateStage(stageCreate, projectId)

### Required Parameters

| Name            | Type                                | Description | Notes |
| --------------- | ----------------------------------- | ----------- | ----- |
| **stageCreate** | **[\*CreateStage](CreateStage.md)** | stage 생성  |
| **projectId**   | **string**                          | projectId   |

### Return type

\*[**CreateUpdateResponse**](CreateUpdateResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStage**

> OkResponse DeleteStage(projectId, stageId)

### Required Parameters

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |
| **stageId**   | **string** | stageId     |

### Return type

\*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStage**

> GetStageDetailResponse GetStage(projectId, stageId)

### Required Parameters

| Name          | Type       | Description | Notes |
| ------------- | ---------- | ----------- | ----- |
| **projectId** | **string** | projectId   |
| **stageId**   | **string** | stageId     |

### Return type

\*[**GetStageDetailResponse**](GetStageDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeStage**

> CreateUpdateResponse ChangeStage(stageUpdate, projectId, stageId)

### Required Parameters

| Name            | Type                                | Description | Notes |
| --------------- | ----------------------------------- | ----------- | ----- |
| **stageUpdate** | **[\*ChangeStage](ChangeStage.md)** | stage 수정  |
| **projectId**   | **string**                          | projectId   |
| **stageId**     | **string**                          | stageId     |

### Return type

\*[**CreateUpdateResponse**](CreateUpdateResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScenarioes**

> GetScenarioListResponse GetScenarioes(projectId, stageId, optional)

### Required Parameters

| Name          | Type                       | Description         | Notes                |
| ------------- | -------------------------- | ------------------- | -------------------- |
| **projectId** | **string**                 | projectId           |
| **stageId**   | **string**                 | stageId             |
| **optional**  | **map[string]interface{}** | optional parameters | nil if no parameters |

### Optional Parameters

Optional parameters are passed through a map[string]interface{}.

| Name             | Type       | Description            | Notes |
| ---------------- | ---------- | ---------------------- | ----- |
| **projectId**    | **string** | projectId              |
| **stageId**      | **string** | stageId                |
| **pageNo**       | **string** | 페이지 NO              |
| **pageSize**     | **string** | 페이지 Size            |
| **scenarioName** | **string** | 시나리오 이름으로 검색 |

### Return type

\*[**GetScenarioListResponse**](GetScenarioListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateScenario**

> CreateUpdateResponse CreateScenario(scenarioCreate, projectId, stageId)

### Required Parameters

| Name               | Type                                      | Description   | Notes |
| ------------------ | ----------------------------------------- | ------------- | ----- |
| **scenarioCreate** | **[\*CreateScenario](CreateScenario.md)** | 시나리오 생성 |
| **projectId**      | **string**                                | projectId     |
| **stageId**        | **string**                                | stageId       |

### Return type

\*[**CreateUpdateResponse**](CreateUpdateResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScenario**

> OkResponse DeleteScenario(projectId, stageId, scenarioId)

### Required Parameters

| Name           | Type       | Description | Notes |
| -------------- | ---------- | ----------- | ----- |
| **projectId**  | **string** | projectId   |
| **stageId**    | **string** | stageId     |
| **scenarioId** | **string** | scenarioId  |

### Return type

\*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Deploy**

> DeployRequestResponse Deploy(projectId, stageId, scenarioId)

### Required Parameters

| Name           | Type       | Description | Notes |
| -------------- | ---------- | ----------- | ----- |
| **projectId**  | **string** | projectId   |
| **stageId**    | **string** | stageId     |
| **scenarioId** | **string** | scenarioId  |

### Return type

\*[**DeployRequestResponse**](DeployRequestResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployRequest**

> DeployRequestResponse DeployRequest(projectId, stageId, scenarioId)

### Required Parameters

| Name           | Type       | Description | Notes |
| -------------- | ---------- | ----------- | ----- |
| **projectId**  | **string** | projectId   |
| **stageId**    | **string** | stageId     |
| **scenarioId** | **string** | scenarioId  |

### Return type

\*[**DeployRequestResponse**](DeployRequestResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScenario**

> GetScenarioDetailResponse GetScenario(projectId, stageId, scenarioId)

### Required Parameters

| Name           | Type       | Description | Notes           |
| -------------- | ---------- | ----------- | --------------- |
| **projectId**  | **string** | projectId   |
| **stageId**    | **string** | stageId     |
| **scenarioId** | **string** | scenarioId  | ### Return type |

\*[**GetScenarioDetailResponse**](GetScenarioDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeScenario**

> CreateUpdateResponse ChangeScenario(scenarioUpdate, projectId, stageId, scenarioId)

### Required Parameters

| Name               | Type                                      | Description   | Notes |
| ------------------ | ----------------------------------------- | ------------- | ----- |
| **scenarioUpdate** | **[\*ChangeScenario](ChangeScenario.md)** | 시나리오 수정 |
| **projectId**      | **string**                                | projectId     |
| **stageId**        | **string**                                | stageId       |
| **scenarioId**     | **string**                                | scenarioId    |

### Return type

\*[**CreateUpdateResponse**](CreateUpdateResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServers**

> GetServerListResponse GetServers()

### Required Parameters

This endpoint does not need any parameter.

### Return type

\*[**GetServerListResponse**](GetServerListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourcebuildProjects**

> GetSourceBuildProjectListResponse GetSourcebuildProjects()

### Required Parameters

This endpoint does not need any parameter.

### Return type

\*[**GetSourceBuildProjectListResponse**](GetSourceBuildProjectListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourcecommitRepositories**

> GetSourceCommitRepositoryListResponse GetSourcecommitRepositories()

### Required Parameters

This endpoint does not need any parameter.

### Return type

\*[**GetSourceCommitRepositoryListResponse**](GetSourceCommitRepositoryListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceCommitBranches**

> GetSourceCommitBranchListResponse GetSourceCommitBranches(repositoryName)

### Required Parameters

This endpoint does not need any parameter.

### Return type

\*[**GetSourceCommitBranchListResponse**](GetSourceCommitBranchListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

-   **Content-Type**: Not defined
-   **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
