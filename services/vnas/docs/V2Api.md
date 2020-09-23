# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/vnas/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNasVolumeAccessControl**](V2Api.md#AddNasVolumeAccessControl) | **Post** /addNasVolumeAccessControl | 
[**ChangeNasVolumeSize**](V2Api.md#ChangeNasVolumeSize) | **Post** /changeNasVolumeSize | 
[**CreateNasVolumeInstance**](V2Api.md#CreateNasVolumeInstance) | **Post** /createNasVolumeInstance | 
[**DeleteNasVolumeInstances**](V2Api.md#DeleteNasVolumeInstances) | **Post** /deleteNasVolumeInstances | 
[**GetNasVolumeInstanceDetail**](V2Api.md#GetNasVolumeInstanceDetail) | **Post** /getNasVolumeInstanceDetail | 
[**GetNasVolumeInstanceList**](V2Api.md#GetNasVolumeInstanceList) | **Post** /getNasVolumeInstanceList | 
[**RemoveNasVolumeAccessControl**](V2Api.md#RemoveNasVolumeAccessControl) | **Post** /removeNasVolumeAccessControl | 
[**SetNasVolumeAccessControl**](V2Api.md#SetNasVolumeAccessControl) | **Post** /setNasVolumeAccessControl | 


# **AddNasVolumeAccessControl**
> AddNasVolumeAccessControlResponse AddNasVolumeAccessControl(addNasVolumeAccessControlRequest)


NAS볼륨접근제어추가

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


NAS볼륨사이즈변경

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

# **CreateNasVolumeInstance**
> CreateNasVolumeInstanceResponse CreateNasVolumeInstance(createNasVolumeInstanceRequest)


NAS볼륨인스턴스생성

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

# **DeleteNasVolumeInstances**
> DeleteNasVolumeInstancesResponse DeleteNasVolumeInstances(deleteNasVolumeInstancesRequest)


NAS볼륨인스턴스제거

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

# **GetNasVolumeInstanceDetail**
> GetNasVolumeInstanceDetailResponse GetNasVolumeInstanceDetail(getNasVolumeInstanceDetailRequest)


NAS볼륨인스턴스상세조회

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


NAS볼륨인스턴스리스트조회

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

# **RemoveNasVolumeAccessControl**
> RemoveNasVolumeAccessControlResponse RemoveNasVolumeAccessControl(removeNasVolumeAccessControlRequest)


NAS볼륨접근제어삭제

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

# **SetNasVolumeAccessControl**
> SetNasVolumeAccessControlResponse SetNasVolumeAccessControl(setNasVolumeAccessControlRequest)


NAS볼륨접근제어설정

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

