# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/server/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNasVolumeAccessControl**](V2Api.md#AddNasVolumeAccessControl) | **Post** /addNasVolumeAccessControl | 
[**AddPortForwardingRules**](V2Api.md#AddPortForwardingRules) | **Post** /addPortForwardingRules | 
[**AssociatePublicIpWithServerInstance**](V2Api.md#AssociatePublicIpWithServerInstance) | **Post** /associatePublicIpWithServerInstance | 
[**ChangeNasVolumeSize**](V2Api.md#ChangeNasVolumeSize) | **Post** /changeNasVolumeSize | 
[**ChangeServerInstanceSpec**](V2Api.md#ChangeServerInstanceSpec) | **Post** /changeServerInstanceSpec | 
[**CreateBlockStorageInstance**](V2Api.md#CreateBlockStorageInstance) | **Post** /createBlockStorageInstance | 
[**CreateLoginKey**](V2Api.md#CreateLoginKey) | **Post** /createLoginKey | 
[**CreateMemberServerImage**](V2Api.md#CreateMemberServerImage) | **Post** /createMemberServerImage | 
[**CreateNasVolumeInstance**](V2Api.md#CreateNasVolumeInstance) | **Post** /createNasVolumeInstance | 
[**CreatePublicIpInstance**](V2Api.md#CreatePublicIpInstance) | **Post** /createPublicIpInstance | 
[**CreateServerInstances**](V2Api.md#CreateServerInstances) | **Post** /createServerInstances | 
[**DeleteBlockStorageInstances**](V2Api.md#DeleteBlockStorageInstances) | **Post** /deleteBlockStorageInstances | 
[**DeleteLoginKey**](V2Api.md#DeleteLoginKey) | **Post** /deleteLoginKey | 
[**DeleteMemberServerImages**](V2Api.md#DeleteMemberServerImages) | **Post** /deleteMemberServerImages | 
[**DeleteNasVolumeInstance**](V2Api.md#DeleteNasVolumeInstance) | **Post** /deleteNasVolumeInstance | 
[**DeletePortForwardingRules**](V2Api.md#DeletePortForwardingRules) | **Post** /deletePortForwardingRules | 
[**DeletePublicIpInstances**](V2Api.md#DeletePublicIpInstances) | **Post** /deletePublicIpInstances | 
[**DisassociatePublicIpFromServerInstance**](V2Api.md#DisassociatePublicIpFromServerInstance) | **Post** /disassociatePublicIpFromServerInstance | 
[**GetAccessControlGroupList**](V2Api.md#GetAccessControlGroupList) | **Post** /getAccessControlGroupList | 
[**GetAccessControlGroupServerInstanceList**](V2Api.md#GetAccessControlGroupServerInstanceList) | **Post** /getAccessControlGroupServerInstanceList | 
[**GetAccessControlRuleList**](V2Api.md#GetAccessControlRuleList) | **Post** /getAccessControlRuleList | 
[**GetBlockStorageInstanceList**](V2Api.md#GetBlockStorageInstanceList) | **Post** /getBlockStorageInstanceList | 
[**GetBlockStorageSnapshotInstanceList**](V2Api.md#GetBlockStorageSnapshotInstanceList) | **Post** /getBlockStorageSnapshotInstanceList | 
[**GetLoginKeyList**](V2Api.md#GetLoginKeyList) | **Post** /getLoginKeyList | 
[**GetMemberServerImageList**](V2Api.md#GetMemberServerImageList) | **Post** /getMemberServerImageList | 
[**GetNasVolumeInstanceList**](V2Api.md#GetNasVolumeInstanceList) | **Post** /getNasVolumeInstanceList | 
[**GetNasVolumeInstanceRatingList**](V2Api.md#GetNasVolumeInstanceRatingList) | **Post** /getNasVolumeInstanceRatingList | 
[**GetPortForwardingRuleList**](V2Api.md#GetPortForwardingRuleList) | **Post** /getPortForwardingRuleList | 
[**GetPublicIpInstanceList**](V2Api.md#GetPublicIpInstanceList) | **Post** /getPublicIpInstanceList | 
[**GetPublicIpTargetServerInstanceList**](V2Api.md#GetPublicIpTargetServerInstanceList) | **Post** /getPublicIpTargetServerInstanceList | 
[**GetRaidList**](V2Api.md#GetRaidList) | **Post** /getRaidList | 
[**GetRegionList**](V2Api.md#GetRegionList) | **Post** /getRegionList | 
[**GetRootPassword**](V2Api.md#GetRootPassword) | **Post** /getRootPassword | 
[**GetServerImageProductList**](V2Api.md#GetServerImageProductList) | **Post** /getServerImageProductList | 
[**GetServerInstanceList**](V2Api.md#GetServerInstanceList) | **Post** /getServerInstanceList | 
[**GetServerProductList**](V2Api.md#GetServerProductList) | **Post** /getServerProductList | 
[**GetZoneList**](V2Api.md#GetZoneList) | **Post** /getZoneList | 
[**RebootServerInstances**](V2Api.md#RebootServerInstances) | **Post** /rebootServerInstances | 
[**RecreateServerInstance**](V2Api.md#RecreateServerInstance) | **Post** /recreateServerInstance | 
[**RemoveNasVolumeAccessControl**](V2Api.md#RemoveNasVolumeAccessControl) | **Post** /removeNasVolumeAccessControl | 
[**SetNasVolumeAccessControl**](V2Api.md#SetNasVolumeAccessControl) | **Post** /setNasVolumeAccessControl | 
[**StartServerInstances**](V2Api.md#StartServerInstances) | **Post** /startServerInstances | 
[**StopServerInstances**](V2Api.md#StopServerInstances) | **Post** /stopServerInstances | 
[**TerminateServerInstances**](V2Api.md#TerminateServerInstances) | **Post** /terminateServerInstances | 


# **AddNasVolumeAccessControl**
> AddNasVolumeAccessControlResponse AddNasVolumeAccessControl(ctx, addNasVolumeAccessControlRequest)


NAS볼륨인스턴스접근제어추가

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **addNasVolumeAccessControlRequest** | [**AddNasVolumeAccessControlRequest**](AddNasVolumeAccessControlRequest.md)| addNasVolumeAccessControlRequest | 

### Return type

[**AddNasVolumeAccessControlResponse**](addNasVolumeAccessControlResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPortForwardingRules**
> AddPortForwardingRulesResponse AddPortForwardingRules(ctx, addPortForwardingRulesRequest)


포트포워딩Rule추가

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **addPortForwardingRulesRequest** | [**AddPortForwardingRulesRequest**](AddPortForwardingRulesRequest.md)| addPortForwardingRulesRequest | 

### Return type

[**AddPortForwardingRulesResponse**](addPortForwardingRulesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssociatePublicIpWithServerInstance**
> AssociatePublicIpWithServerInstanceResponse AssociatePublicIpWithServerInstance(ctx, associatePublicIpWithServerInstanceRequest)


공인IP를서버인스턴스에할당

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **associatePublicIpWithServerInstanceRequest** | [**AssociatePublicIpWithServerInstanceRequest**](AssociatePublicIpWithServerInstanceRequest.md)| associatePublicIpWithServerInstanceRequest | 

### Return type

[**AssociatePublicIpWithServerInstanceResponse**](associatePublicIpWithServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeNasVolumeSize**
> ChangeNasVolumeSizeResponse ChangeNasVolumeSize(ctx, changeNasVolumeSizeRequest)


NAS볼륨사이즈변경

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **changeNasVolumeSizeRequest** | [**ChangeNasVolumeSizeRequest**](ChangeNasVolumeSizeRequest.md)| changeNasVolumeSizeRequest | 

### Return type

[**ChangeNasVolumeSizeResponse**](changeNasVolumeSizeResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeServerInstanceSpec**
> ChangeServerInstanceSpecResponse ChangeServerInstanceSpec(ctx, changeServerInstanceSpecRequest)


서버인스턴스스팩변경

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **changeServerInstanceSpecRequest** | [**ChangeServerInstanceSpecRequest**](ChangeServerInstanceSpecRequest.md)| changeServerInstanceSpecRequest | 

### Return type

[**ChangeServerInstanceSpecResponse**](changeServerInstanceSpecResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBlockStorageInstance**
> CreateBlockStorageInstanceResponse CreateBlockStorageInstance(ctx, createBlockStorageInstanceRequest)


블록스토리지인스턴스생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **createBlockStorageInstanceRequest** | [**CreateBlockStorageInstanceRequest**](CreateBlockStorageInstanceRequest.md)| createBlockStorageInstanceRequest | 

### Return type

[**CreateBlockStorageInstanceResponse**](createBlockStorageInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoginKey**
> CreateLoginKeyResponse CreateLoginKey(ctx, createLoginKeyRequest)


로그인키생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **createLoginKeyRequest** | [**CreateLoginKeyRequest**](CreateLoginKeyRequest.md)| createLoginKeyRequest | 

### Return type

[**CreateLoginKeyResponse**](createLoginKeyResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMemberServerImage**
> CreateMemberServerImageResponse CreateMemberServerImage(ctx, createMemberServerImageRequest)


회원서버이미지생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **createMemberServerImageRequest** | [**CreateMemberServerImageRequest**](CreateMemberServerImageRequest.md)| createMemberServerImageRequest | 

### Return type

[**CreateMemberServerImageResponse**](createMemberServerImageResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNasVolumeInstance**
> CreateNasVolumeInstanceResponse CreateNasVolumeInstance(ctx, createNasVolumeInstanceRequest)


NAS볼륨인스턴스생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **createNasVolumeInstanceRequest** | [**CreateNasVolumeInstanceRequest**](CreateNasVolumeInstanceRequest.md)| createNasVolumeInstanceRequest | 

### Return type

[**CreateNasVolumeInstanceResponse**](createNasVolumeInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePublicIpInstance**
> CreatePublicIpInstanceResponse CreatePublicIpInstance(ctx, createPublicIpInstanceRequest)


공인IP인스턴스생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **createPublicIpInstanceRequest** | [**CreatePublicIpInstanceRequest**](CreatePublicIpInstanceRequest.md)| createPublicIpInstanceRequest | 

### Return type

[**CreatePublicIpInstanceResponse**](createPublicIpInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServerInstances**
> CreateServerInstancesResponse CreateServerInstances(ctx, createServerInstancesRequest)


서버인스턴스생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **createServerInstancesRequest** | [**CreateServerInstancesRequest**](CreateServerInstancesRequest.md)| createServerInstancesRequest | 

### Return type

[**CreateServerInstancesResponse**](createServerInstancesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBlockStorageInstances**
> DeleteBlockStorageInstancesResponse DeleteBlockStorageInstances(ctx, deleteBlockStorageInstancesRequest)


블록스토리지인스턴스삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deleteBlockStorageInstancesRequest** | [**DeleteBlockStorageInstancesRequest**](DeleteBlockStorageInstancesRequest.md)| deleteBlockStorageInstancesRequest | 

### Return type

[**DeleteBlockStorageInstancesResponse**](deleteBlockStorageInstancesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoginKey**
> DeleteLoginKeyResponse DeleteLoginKey(ctx, deleteLoginKeyRequest)


로그인키삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deleteLoginKeyRequest** | [**DeleteLoginKeyRequest**](DeleteLoginKeyRequest.md)| deleteLoginKeyRequest | 

### Return type

[**DeleteLoginKeyResponse**](deleteLoginKeyResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMemberServerImages**
> DeleteMemberServerImagesResponse DeleteMemberServerImages(ctx, deleteMemberServerImagesRequest)


회원서버이미지삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deleteMemberServerImagesRequest** | [**DeleteMemberServerImagesRequest**](DeleteMemberServerImagesRequest.md)| deleteMemberServerImagesRequest | 

### Return type

[**DeleteMemberServerImagesResponse**](deleteMemberServerImagesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNasVolumeInstance**
> DeleteNasVolumeInstanceResponse DeleteNasVolumeInstance(ctx, deleteNasVolumeInstanceRequest)


NAS볼륨인스턴스삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deleteNasVolumeInstanceRequest** | [**DeleteNasVolumeInstanceRequest**](DeleteNasVolumeInstanceRequest.md)| deleteNasVolumeInstanceRequest | 

### Return type

[**DeleteNasVolumeInstanceResponse**](deleteNasVolumeInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePortForwardingRules**
> DeletePortForwardingRulesResponse DeletePortForwardingRules(ctx, deletePortForwardingRulesRequest)


포트포워딩Rule삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deletePortForwardingRulesRequest** | [**DeletePortForwardingRulesRequest**](DeletePortForwardingRulesRequest.md)| deletePortForwardingRulesRequest | 

### Return type

[**DeletePortForwardingRulesResponse**](deletePortForwardingRulesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePublicIpInstances**
> DeletePublicIpInstancesResponse DeletePublicIpInstances(ctx, deletePublicIpInstancesRequest)


공인IP인스턴스삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deletePublicIpInstancesRequest** | [**DeletePublicIpInstancesRequest**](DeletePublicIpInstancesRequest.md)| deletePublicIpInstancesRequest | 

### Return type

[**DeletePublicIpInstancesResponse**](deletePublicIpInstancesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisassociatePublicIpFromServerInstance**
> DisassociatePublicIpFromServerInstanceResponse DisassociatePublicIpFromServerInstance(ctx, disassociatePublicIpFromServerInstanceRequest)


공인IP를서버인스턴스에할당해제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **disassociatePublicIpFromServerInstanceRequest** | [**DisassociatePublicIpFromServerInstanceRequest**](DisassociatePublicIpFromServerInstanceRequest.md)| disassociatePublicIpFromServerInstanceRequest | 

### Return type

[**DisassociatePublicIpFromServerInstanceResponse**](disassociatePublicIpFromServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessControlGroupList**
> GetAccessControlGroupListResponse GetAccessControlGroupList(ctx, getAccessControlGroupListRequest)


접근제어그룹리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getAccessControlGroupListRequest** | [**GetAccessControlGroupListRequest**](GetAccessControlGroupListRequest.md)| getAccessControlGroupListRequest | 

### Return type

[**GetAccessControlGroupListResponse**](getAccessControlGroupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessControlGroupServerInstanceList**
> GetAccessControlGroupServerInstanceListResponse GetAccessControlGroupServerInstanceList(ctx, getAccessControlGroupServerInstanceListRequest)


접근제어그룹적용된서버인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getAccessControlGroupServerInstanceListRequest** | [**GetAccessControlGroupServerInstanceListRequest**](GetAccessControlGroupServerInstanceListRequest.md)| getAccessControlGroupServerInstanceListRequest | 

### Return type

[**GetAccessControlGroupServerInstanceListResponse**](getAccessControlGroupServerInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessControlRuleList**
> GetAccessControlRuleListResponse GetAccessControlRuleList(ctx, getAccessControlRuleListRequest)


접근제어규칙리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getAccessControlRuleListRequest** | [**GetAccessControlRuleListRequest**](GetAccessControlRuleListRequest.md)| getAccessControlRuleListRequest | 

### Return type

[**GetAccessControlRuleListResponse**](getAccessControlRuleListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockStorageInstanceList**
> GetBlockStorageInstanceListResponse GetBlockStorageInstanceList(ctx, getBlockStorageInstanceListRequest)


블록스토리지인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getBlockStorageInstanceListRequest** | [**GetBlockStorageInstanceListRequest**](GetBlockStorageInstanceListRequest.md)| getBlockStorageInstanceListRequest | 

### Return type

[**GetBlockStorageInstanceListResponse**](getBlockStorageInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockStorageSnapshotInstanceList**
> GetBlockStorageSnapshotInstanceListResponse GetBlockStorageSnapshotInstanceList(ctx, getBlockStorageSnapshotInstanceListRequest)


블록스토리지스냅샷인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getBlockStorageSnapshotInstanceListRequest** | [**GetBlockStorageSnapshotInstanceListRequest**](GetBlockStorageSnapshotInstanceListRequest.md)| getBlockStorageSnapshotInstanceListRequest | 

### Return type

[**GetBlockStorageSnapshotInstanceListResponse**](getBlockStorageSnapshotInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoginKeyList**
> GetLoginKeyListResponse GetLoginKeyList(ctx, getLoginKeyListRequest)


로그인키리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getLoginKeyListRequest** | [**GetLoginKeyListRequest**](GetLoginKeyListRequest.md)| getLoginKeyListRequest | 

### Return type

[**GetLoginKeyListResponse**](getLoginKeyListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMemberServerImageList**
> GetMemberServerImageListResponse GetMemberServerImageList(ctx, getMemberServerImageListRequest)


회원서버이미지리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getMemberServerImageListRequest** | [**GetMemberServerImageListRequest**](GetMemberServerImageListRequest.md)| getMemberServerImageListRequest | 

### Return type

[**GetMemberServerImageListResponse**](getMemberServerImageListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNasVolumeInstanceList**
> GetNasVolumeInstanceListResponse GetNasVolumeInstanceList(ctx, getNasVolumeInstanceListRequest)


NAS볼륨인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getNasVolumeInstanceListRequest** | [**GetNasVolumeInstanceListRequest**](GetNasVolumeInstanceListRequest.md)| getNasVolumeInstanceListRequest | 

### Return type

[**GetNasVolumeInstanceListResponse**](getNasVolumeInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNasVolumeInstanceRatingList**
> GetNasVolumeInstanceRatingListResponse GetNasVolumeInstanceRatingList(ctx, getNasVolumeInstanceRatingListRequest)


NAS볼륨인스턴스측정리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getNasVolumeInstanceRatingListRequest** | [**GetNasVolumeInstanceRatingListRequest**](GetNasVolumeInstanceRatingListRequest.md)| getNasVolumeInstanceRatingListRequest | 

### Return type

[**GetNasVolumeInstanceRatingListResponse**](getNasVolumeInstanceRatingListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortForwardingRuleList**
> GetPortForwardingRuleListResponse GetPortForwardingRuleList(ctx, getPortForwardingRuleListRequest)


포트포워딩Rule리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getPortForwardingRuleListRequest** | [**GetPortForwardingRuleListRequest**](GetPortForwardingRuleListRequest.md)| getPortForwardingRuleListRequest | 

### Return type

[**GetPortForwardingRuleListResponse**](getPortForwardingRuleListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPublicIpInstanceList**
> GetPublicIpInstanceListResponse GetPublicIpInstanceList(ctx, getPublicIpInstanceListRequest)


공인IP인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getPublicIpInstanceListRequest** | [**GetPublicIpInstanceListRequest**](GetPublicIpInstanceListRequest.md)| getPublicIpInstanceListRequest | 

### Return type

[**GetPublicIpInstanceListResponse**](getPublicIpInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPublicIpTargetServerInstanceList**
> GetPublicIpTargetServerInstanceListResponse GetPublicIpTargetServerInstanceList(ctx, getPublicIpTargetServerInstanceListRequest)


공인IP할당(가능)서버인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getPublicIpTargetServerInstanceListRequest** | [**GetPublicIpTargetServerInstanceListRequest**](GetPublicIpTargetServerInstanceListRequest.md)| getPublicIpTargetServerInstanceListRequest | 

### Return type

[**GetPublicIpTargetServerInstanceListResponse**](getPublicIpTargetServerInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRaidList**
> GetRaidListResponse GetRaidList(ctx, getRaidListRequest)


RAID리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getRaidListRequest** | [**GetRaidListRequest**](GetRaidListRequest.md)| getRaidListRequest | 

### Return type

[**GetRaidListResponse**](getRaidListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegionList**
> GetRegionListResponse GetRegionList(ctx, getRegionListRequest)


REGION리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getRegionListRequest** | [**GetRegionListRequest**](GetRegionListRequest.md)| getRegionListRequest | 

### Return type

[**GetRegionListResponse**](getRegionListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRootPassword**
> GetRootPasswordResponse GetRootPassword(ctx, getRootPasswordRequest)


루트패스워드조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getRootPasswordRequest** | [**GetRootPasswordRequest**](GetRootPasswordRequest.md)| getRootPasswordRequest | 

### Return type

[**GetRootPasswordResponse**](getRootPasswordResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerImageProductList**
> GetServerImageProductListResponse GetServerImageProductList(ctx, getServerImageProductListRequest)


서버이미지상품리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getServerImageProductListRequest** | [**GetServerImageProductListRequest**](GetServerImageProductListRequest.md)| getServerImageProductListRequest | 

### Return type

[**GetServerImageProductListResponse**](getServerImageProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerInstanceList**
> GetServerInstanceListResponse GetServerInstanceList(ctx, getServerInstanceListRequest)


서버인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getServerInstanceListRequest** | [**GetServerInstanceListRequest**](GetServerInstanceListRequest.md)| getServerInstanceListRequest | 

### Return type

[**GetServerInstanceListResponse**](getServerInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerProductList**
> GetServerProductListResponse GetServerProductList(ctx, getServerProductListRequest)


서버상품리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getServerProductListRequest** | [**GetServerProductListRequest**](GetServerProductListRequest.md)| getServerProductListRequest | 

### Return type

[**GetServerProductListResponse**](getServerProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetZoneList**
> GetZoneListResponse GetZoneList(ctx, getZoneListRequest)


ZONE리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getZoneListRequest** | [**GetZoneListRequest**](GetZoneListRequest.md)| getZoneListRequest | 

### Return type

[**GetZoneListResponse**](getZoneListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootServerInstances**
> RebootServerInstancesResponse RebootServerInstances(ctx, rebootServerInstancesRequest)


서버인스턴스재시작

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **rebootServerInstancesRequest** | [**RebootServerInstancesRequest**](RebootServerInstancesRequest.md)| rebootServerInstancesRequest | 

### Return type

[**RebootServerInstancesResponse**](rebootServerInstancesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecreateServerInstance**
> RecreateServerInstanceResponse RecreateServerInstance(ctx, recreateServerInstanceRequest)


서버인스턴스재생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **recreateServerInstanceRequest** | [**RecreateServerInstanceRequest**](RecreateServerInstanceRequest.md)| recreateServerInstanceRequest | 

### Return type

[**RecreateServerInstanceResponse**](recreateServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveNasVolumeAccessControl**
> RemoveNasVolumeAccessControlResponse RemoveNasVolumeAccessControl(ctx, removeNasVolumeAccessControlRequest)


NAS볼륨인스턴스접근제어제거

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **removeNasVolumeAccessControlRequest** | [**RemoveNasVolumeAccessControlRequest**](RemoveNasVolumeAccessControlRequest.md)| removeNasVolumeAccessControlRequest | 

### Return type

[**RemoveNasVolumeAccessControlResponse**](removeNasVolumeAccessControlResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetNasVolumeAccessControl**
> SetNasVolumeAccessControlResponse SetNasVolumeAccessControl(ctx, setNasVolumeAccessControlRequest)


NAS볼륨인스턴스접근제어설정

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **setNasVolumeAccessControlRequest** | [**SetNasVolumeAccessControlRequest**](SetNasVolumeAccessControlRequest.md)| setNasVolumeAccessControlRequest | 

### Return type

[**SetNasVolumeAccessControlResponse**](setNasVolumeAccessControlResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartServerInstances**
> StartServerInstancesResponse StartServerInstances(ctx, startServerInstancesRequest)


서버인스턴스시작

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **startServerInstancesRequest** | [**StartServerInstancesRequest**](StartServerInstancesRequest.md)| startServerInstancesRequest | 

### Return type

[**StartServerInstancesResponse**](startServerInstancesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopServerInstances**
> StopServerInstancesResponse StopServerInstances(ctx, stopServerInstancesRequest)


서버인스턴스종료

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **stopServerInstancesRequest** | [**StopServerInstancesRequest**](StopServerInstancesRequest.md)| stopServerInstancesRequest | 

### Return type

[**StopServerInstancesResponse**](stopServerInstancesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminateServerInstances**
> TerminateServerInstancesResponse TerminateServerInstances(ctx, terminateServerInstancesRequest)


서버인스턴스반납

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **terminateServerInstancesRequest** | [**TerminateServerInstancesRequest**](TerminateServerInstancesRequest.md)| terminateServerInstancesRequest | 

### Return type

[**TerminateServerInstancesResponse**](terminateServerInstancesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

