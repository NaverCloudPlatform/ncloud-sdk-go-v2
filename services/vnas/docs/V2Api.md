# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/vnas/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNasVolumeAccessControl**](V2Api.md#AddNasVolumeAccessControl) | **Post** /addNasVolumeAccessControl | 
[**ChangeNasVolumeSize**](V2Api.md#ChangeNasVolumeSize) | **Post** /changeNasVolumeSize | 
[**ChangeNasVolumeSnapshotConfiguration**](V2Api.md#ChangeNasVolumeSnapshotConfiguration) | **Post** /changeNasVolumeSnapshotConfiguration | 
[**CreateNasVolumeInstance**](V2Api.md#CreateNasVolumeInstance) | **Post** /createNasVolumeInstance | 
[**CreateNasVolumeSnapshot**](V2Api.md#CreateNasVolumeSnapshot) | **Post** /createNasVolumeSnapshot | 
[**DeleteNasVolumeInstances**](V2Api.md#DeleteNasVolumeInstances) | **Post** /deleteNasVolumeInstances | 
[**DeleteNasVolumeSnapshot**](V2Api.md#DeleteNasVolumeSnapshot) | **Post** /deleteNasVolumeSnapshot | 
[**GetNasVolumeInstanceDetail**](V2Api.md#GetNasVolumeInstanceDetail) | **Post** /getNasVolumeInstanceDetail | 
[**GetNasVolumeInstanceList**](V2Api.md#GetNasVolumeInstanceList) | **Post** /getNasVolumeInstanceList | 
[**GetNasVolumeInstanceRatingList**](V2Api.md#GetNasVolumeInstanceRatingList) | **Post** /getNasVolumeInstanceRatingList | 
[**GetNasVolumeSnapshotConfigurationHistoryList**](V2Api.md#GetNasVolumeSnapshotConfigurationHistoryList) | **Post** /getNasVolumeSnapshotConfigurationHistoryList | 
[**GetNasVolumeSnapshotList**](V2Api.md#GetNasVolumeSnapshotList) | **Post** /getNasVolumeSnapshotList | 
[**RemoveNasVolumeAccessControl**](V2Api.md#RemoveNasVolumeAccessControl) | **Post** /removeNasVolumeAccessControl | 
[**RestoreNasVolumeWithSnapshot**](V2Api.md#RestoreNasVolumeWithSnapshot) | **Post** /restoreNasVolumeWithSnapshot | 
[**SetNasVolumeAccessControl**](V2Api.md#SetNasVolumeAccessControl) | **Post** /setNasVolumeAccessControl | 
[**SetNasVolumeReturnProtection**](V2Api.md#SetNasVolumeReturnProtection) | **Post** /setNasVolumeReturnProtection | 


# **AddNasVolumeAccessControl**
> AddNasVolumeAccessControlResponse AddNasVolumeAccessControl(addNasVolumeAccessControlRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addNasVolumeAccessControlRequest** | **[\*AddNasVolumeAccessControlRequest](AddNasVolumeAccessControlRequest.md)** | addNasVolumeAccessControlRequest | 

### Return type

*[**AddNasVolumeAccessControlResponse**](AddNasVolumeAccessControlResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeNasVolumeSize**
> ChangeNasVolumeSizeResponse ChangeNasVolumeSize(changeNasVolumeSizeRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**changeNasVolumeSizeRequest** | **[\*ChangeNasVolumeSizeRequest](ChangeNasVolumeSizeRequest.md)** | changeNasVolumeSizeRequest | 

### Return type

*[**ChangeNasVolumeSizeResponse**](ChangeNasVolumeSizeResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeNasVolumeSnapshotConfiguration**
> ChangeNasVolumeSnapshotConfigurationResponse ChangeNasVolumeSnapshotConfiguration(changeNasVolumeSnapshotConfigurationRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**changeNasVolumeSnapshotConfigurationRequest** | **[\*ChangeNasVolumeSnapshotConfigurationRequest](ChangeNasVolumeSnapshotConfigurationRequest.md)** | changeNasVolumeSnapshotConfigurationRequest | 

### Return type

*[**ChangeNasVolumeSnapshotConfigurationResponse**](ChangeNasVolumeSnapshotConfigurationResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNasVolumeInstance**
> CreateNasVolumeInstanceResponse CreateNasVolumeInstance(createNasVolumeInstanceRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createNasVolumeInstanceRequest** | **[\*CreateNasVolumeInstanceRequest](CreateNasVolumeInstanceRequest.md)** | createNasVolumeInstanceRequest | 

### Return type

*[**CreateNasVolumeInstanceResponse**](CreateNasVolumeInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNasVolumeSnapshot**
> CreateNasVolumeSnapshotResponse CreateNasVolumeSnapshot(createNasVolumeSnapshotRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createNasVolumeSnapshotRequest** | **[\*CreateNasVolumeSnapshotRequest](CreateNasVolumeSnapshotRequest.md)** | createNasVolumeSnapshotRequest | 

### Return type

*[**CreateNasVolumeSnapshotResponse**](CreateNasVolumeSnapshotResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNasVolumeInstances**
> DeleteNasVolumeInstancesResponse DeleteNasVolumeInstances(deleteNasVolumeInstancesRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteNasVolumeInstancesRequest** | **[\*DeleteNasVolumeInstancesRequest](DeleteNasVolumeInstancesRequest.md)** | deleteNasVolumeInstancesRequest | 

### Return type

*[**DeleteNasVolumeInstancesResponse**](DeleteNasVolumeInstancesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNasVolumeSnapshot**
> DeleteNasVolumeSnapshotResponse DeleteNasVolumeSnapshot(deleteNasVolumeSnapshotRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteNasVolumeSnapshotRequest** | **[\*DeleteNasVolumeSnapshotRequest](DeleteNasVolumeSnapshotRequest.md)** | deleteNasVolumeSnapshotRequest | 

### Return type

*[**DeleteNasVolumeSnapshotResponse**](DeleteNasVolumeSnapshotResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNasVolumeInstanceDetail**
> GetNasVolumeInstanceDetailResponse GetNasVolumeInstanceDetail(getNasVolumeInstanceDetailRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getNasVolumeInstanceDetailRequest** | **[\*GetNasVolumeInstanceDetailRequest](GetNasVolumeInstanceDetailRequest.md)** | getNasVolumeInstanceDetailRequest | 

### Return type

*[**GetNasVolumeInstanceDetailResponse**](GetNasVolumeInstanceDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNasVolumeInstanceList**
> GetNasVolumeInstanceListResponse GetNasVolumeInstanceList(getNasVolumeInstanceListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getNasVolumeInstanceListRequest** | **[\*GetNasVolumeInstanceListRequest](GetNasVolumeInstanceListRequest.md)** | getNasVolumeInstanceListRequest | 

### Return type

*[**GetNasVolumeInstanceListResponse**](GetNasVolumeInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNasVolumeInstanceRatingList**
> GetNasVolumeInstanceRatingListResponse GetNasVolumeInstanceRatingList(getNasVolumeInstanceRatingListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getNasVolumeInstanceRatingListRequest** | **[\*GetNasVolumeInstanceRatingListRequest](GetNasVolumeInstanceRatingListRequest.md)** | getNasVolumeInstanceRatingListRequest | 

### Return type

*[**GetNasVolumeInstanceRatingListResponse**](GetNasVolumeInstanceRatingListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNasVolumeSnapshotConfigurationHistoryList**
> GetNasVolumeSnapshotConfigurationHistoryListResponse GetNasVolumeSnapshotConfigurationHistoryList(getNasVolumeSnapshotConfigurationHistoryListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getNasVolumeSnapshotConfigurationHistoryListRequest** | **[\*GetNasVolumeSnapshotConfigurationHistoryListRequest](GetNasVolumeSnapshotConfigurationHistoryListRequest.md)** | getNasVolumeSnapshotConfigurationHistoryListRequest | 

### Return type

*[**GetNasVolumeSnapshotConfigurationHistoryListResponse**](GetNasVolumeSnapshotConfigurationHistoryListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNasVolumeSnapshotList**
> GetNasVolumeSnapshotListResponse GetNasVolumeSnapshotList(getNasVolumeSnapshotListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getNasVolumeSnapshotListRequest** | **[\*GetNasVolumeSnapshotListRequest](GetNasVolumeSnapshotListRequest.md)** | getNasVolumeSnapshotListRequest | 

### Return type

*[**GetNasVolumeSnapshotListResponse**](GetNasVolumeSnapshotListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveNasVolumeAccessControl**
> RemoveNasVolumeAccessControlResponse RemoveNasVolumeAccessControl(removeNasVolumeAccessControlRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**removeNasVolumeAccessControlRequest** | **[\*RemoveNasVolumeAccessControlRequest](RemoveNasVolumeAccessControlRequest.md)** | removeNasVolumeAccessControlRequest | 

### Return type

*[**RemoveNasVolumeAccessControlResponse**](RemoveNasVolumeAccessControlResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreNasVolumeWithSnapshot**
> RestoreNasVolumeWithSnapshotResponse RestoreNasVolumeWithSnapshot(restoreNasVolumeWithSnapshotRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**restoreNasVolumeWithSnapshotRequest** | **[\*RestoreNasVolumeWithSnapshotRequest](RestoreNasVolumeWithSnapshotRequest.md)** | restoreNasVolumeWithSnapshotRequest | 

### Return type

*[**RestoreNasVolumeWithSnapshotResponse**](RestoreNasVolumeWithSnapshotResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetNasVolumeAccessControl**
> SetNasVolumeAccessControlResponse SetNasVolumeAccessControl(setNasVolumeAccessControlRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setNasVolumeAccessControlRequest** | **[\*SetNasVolumeAccessControlRequest](SetNasVolumeAccessControlRequest.md)** | setNasVolumeAccessControlRequest | 

### Return type

*[**SetNasVolumeAccessControlResponse**](SetNasVolumeAccessControlResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetNasVolumeReturnProtection**
> SetNasVolumeReturnProtectionResponse SetNasVolumeReturnProtection(setNasVolumeReturnProtectionRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setNasVolumeReturnProtectionRequest** | **[\*SetNasVolumeReturnProtectionRequest](SetNasVolumeReturnProtectionRequest.md)** | setNasVolumeReturnProtectionRequest | 

### Return type

*[**SetNasVolumeReturnProtectionResponse**](SetNasVolumeReturnProtectionResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

