# Go API client for sourcebuild

    <br/>https://sourcebuild.apigw.ntruss.com/api/v1

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2022-04-22T08:24:50Z
- Package version: 
- Build package: io.swagger.codegen.languages.NcpGoForVnksClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
"./sourcebuild"
```

## Documentation for API Endpoints

All URIs are relative to *https://sourcebuild.apigw.ntruss.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V1Api* | [**GetContainerRegistry**](docs/V1Api.md#GetContainerRegistry) | **Get** /containerregistry/registry | 
*V1Api* | [**GetEnvCompute**](docs/V1Api.md#GetEnvCompute) | **Get** /env/compute | 
*V1Api* | [**GetEnvDocker**](docs/V1Api.md#GetEnvDocker) | **Get** /env/docker | 
*V1Api* | [**GetEnvOs**](docs/V1Api.md#GetEnvOs) | **Get** /env/os | 
*V1Api* | [**GetEnvRuntime**](docs/V1Api.md#GetEnvRuntime) | **Get** /env/os/{osId}/runtime | 
*V1Api* | [**GetEnvRuntimeVersion**](docs/V1Api.md#GetEnvRuntimeVersion) | **Get** /env/os/{osId}/runtime/{runtimeId}/version | 
*V1Api* | [**GetObjectstorageBucket**](docs/V1Api.md#GetObjectstorageBucket) | **Get** /objectstorage/bucket | 
*V1Api* | [**GetProjects**](docs/V1Api.md#GetProjects) | **Get** /project | 
*V1Api* | [**CreateProject**](docs/V1Api.md#CreateProject) | **Post** /project | 
*V1Api* | [**CancelBuild**](docs/V1Api.md#CancelBuild) | **Delete** /project/{projectId}/build | 
*V1Api* | [**StartBuild**](docs/V1Api.md#StartBuild) | **Post** /project/{projectId}/build | 
*V1Api* | [**DeleteProject**](docs/V1Api.md#DeleteProject) | **Delete** /project/{projectId} | 
*V1Api* | [**GetProject**](docs/V1Api.md#GetProject) | **Get** /project/{projectId} | 
*V1Api* | [**GetBuildHistory**](docs/V1Api.md#GetBuildHistory) | **Get** /project/{projectId}/history | 
*V1Api* | [**ChangeProject**](docs/V1Api.md#ChangeProject) | **Patch** /project/{projectId} | 
*V1Api* | [**GetSourcecommitRepositories**](docs/V1Api.md#GetSourcecommitRepositories) | **Get** /sourcecommit/repository | 
*V1Api* | [**GetSourcecommitRepositoryBranches**](docs/V1Api.md#GetSourcecommitRepositoryBranches) | **Get** /sourcecommit/repository/{repositoryName}/branch | 


## Documentation For Models

 - [CancelBuild](docs/CancelBuild.md)
 - [ChangeProject](docs/ChangeProject.md)
 - [ChangeProjectEnv](docs/ChangeProjectEnv.md)
 - [ChangeProjectEnvPlatform](docs/ChangeProjectEnvPlatform.md)
 - [CreateProject](docs/CreateProject.md)
 - [CreateProjectArtifact](docs/CreateProjectArtifact.md)
 - [CreateProjectArtifactStorage](docs/CreateProjectArtifactStorage.md)
 - [CreateProjectCache](docs/CreateProjectCache.md)
 - [CreateProjectCmd](docs/CreateProjectCmd.md)
 - [CreateProjectCmdDockerbuild](docs/CreateProjectCmdDockerbuild.md)
 - [CreateProjectEnv](docs/CreateProjectEnv.md)
 - [CreateProjectEnvCompute](docs/CreateProjectEnvCompute.md)
 - [CreateProjectEnvDocker](docs/CreateProjectEnvDocker.md)
 - [CreateProjectEnvEnvVars](docs/CreateProjectEnvEnvVars.md)
 - [CreateProjectEnvPlatform](docs/CreateProjectEnvPlatform.md)
 - [CreateProjectLinked](docs/CreateProjectLinked.md)
 - [CreateProjectResponse](docs/CreateProjectResponse.md)
 - [CreateProjectSource](docs/CreateProjectSource.md)
 - [CreateProjectSourceConfig](docs/CreateProjectSourceConfig.md)
 - [EnvPlatformConfigRequest](docs/EnvPlatformConfigRequest.md)
 - [EnvPlatformConfigRequestOs](docs/EnvPlatformConfigRequestOs.md)
 - [EnvPlatformConfigRequestRuntime](docs/EnvPlatformConfigRequestRuntime.md)
 - [EnvPlatformConfigRequestRuntimeVersion](docs/EnvPlatformConfigRequestRuntimeVersion.md)
 - [EnvPlatformConfigResponse](docs/EnvPlatformConfigResponse.md)
 - [EnvPlatformConfigResponseOs](docs/EnvPlatformConfigResponseOs.md)
 - [EnvPlatformConfigResponseRuntime](docs/EnvPlatformConfigResponseRuntime.md)
 - [EnvPlatformConfigResponseRuntimeVersion](docs/EnvPlatformConfigResponseRuntimeVersion.md)
 - [GetBuildHistoryResponse](docs/GetBuildHistoryResponse.md)
 - [GetBuildHistoryResponseHistory](docs/GetBuildHistoryResponseHistory.md)
 - [GetComputeEnvResponse](docs/GetComputeEnvResponse.md)
 - [GetComputeEnvResponseCompute](docs/GetComputeEnvResponseCompute.md)
 - [GetContainerRegistryResponse](docs/GetContainerRegistryResponse.md)
 - [GetContainerRegistryResponseRegistry](docs/GetContainerRegistryResponseRegistry.md)
 - [GetDockerEnvResponse](docs/GetDockerEnvResponse.md)
 - [GetDockerEnvResponseDocker](docs/GetDockerEnvResponseDocker.md)
 - [GetObjsBucketResponse](docs/GetObjsBucketResponse.md)
 - [GetObjsBucketResponseBucket](docs/GetObjsBucketResponseBucket.md)
 - [GetOsEnvResponse](docs/GetOsEnvResponse.md)
 - [GetOsEnvResponseOs](docs/GetOsEnvResponseOs.md)
 - [GetProjectDetailResponse](docs/GetProjectDetailResponse.md)
 - [GetProjectDetailResponseArtifact](docs/GetProjectDetailResponseArtifact.md)
 - [GetProjectDetailResponseArtifactStorage](docs/GetProjectDetailResponseArtifactStorage.md)
 - [GetProjectDetailResponseCreated](docs/GetProjectDetailResponseCreated.md)
 - [GetProjectDetailResponseEnv](docs/GetProjectDetailResponseEnv.md)
 - [GetProjectDetailResponseEnvCompute](docs/GetProjectDetailResponseEnvCompute.md)
 - [GetProjectDetailResponseEnvPlatform](docs/GetProjectDetailResponseEnvPlatform.md)
 - [GetProjectDetailResponseLastBuild](docs/GetProjectDetailResponseLastBuild.md)
 - [GetProjectDetailResponseSource](docs/GetProjectDetailResponseSource.md)
 - [GetProjectDetailResponseSourceConfig](docs/GetProjectDetailResponseSourceConfig.md)
 - [GetProjectListResponse](docs/GetProjectListResponse.md)
 - [GetProjectListResponseProject](docs/GetProjectListResponseProject.md)
 - [GetRuntimeEnvResponse](docs/GetRuntimeEnvResponse.md)
 - [GetRuntimeVersionEnvResponse](docs/GetRuntimeVersionEnvResponse.md)
 - [GetScBranchResponse](docs/GetScBranchResponse.md)
 - [GetScRepositoryResponse](docs/GetScRepositoryResponse.md)
 - [GetScRepositoryResponseRepository](docs/GetScRepositoryResponseRepository.md)
 - [StartBuildResponse](docs/StartBuildResponse.md)
