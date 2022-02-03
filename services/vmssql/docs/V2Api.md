# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/vmssql/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudMssqlInstance**](V2Api.md#CreateCloudMssqlInstance) | **Post** /createCloudMssqlInstance | 
[**CreateCloudMssqlSlaveInstance**](V2Api.md#CreateCloudMssqlSlaveInstance) | **Post** /createCloudMssqlSlaveInstance | 
[**DeleteCloudMssqlInstance**](V2Api.md#DeleteCloudMssqlInstance) | **Post** /deleteCloudMssqlInstance | 
[**DeleteCloudMssqlServerInstance**](V2Api.md#DeleteCloudMssqlServerInstance) | **Post** /deleteCloudMssqlServerInstance | 
[**DownloadDmsFile**](V2Api.md#DownloadDmsFile) | **Post** /downloadDmsFile | 
[**GetCloudMssqlBackupDetailList**](V2Api.md#GetCloudMssqlBackupDetailList) | **Post** /getCloudMssqlBackupDetailList | 
[**GetCloudMssqlBackupList**](V2Api.md#GetCloudMssqlBackupList) | **Post** /getCloudMssqlBackupList | 
[**GetCloudMssqlCharacterSetList**](V2Api.md#GetCloudMssqlCharacterSetList) | **Post** /getCloudMssqlCharacterSetList | 
[**GetCloudMssqlConfigGroupList**](V2Api.md#GetCloudMssqlConfigGroupList) | **Post** /getCloudMssqlConfigGroupList | 
[**GetCloudMssqlImageProductList**](V2Api.md#GetCloudMssqlImageProductList) | **Post** /getCloudMssqlImageProductList | 
[**GetCloudMssqlInstanceDetail**](V2Api.md#GetCloudMssqlInstanceDetail) | **Post** /getCloudMssqlInstanceDetail | 
[**GetCloudMssqlInstanceList**](V2Api.md#GetCloudMssqlInstanceList) | **Post** /getCloudMssqlInstanceList | 
[**GetCloudMssqlProductList**](V2Api.md#GetCloudMssqlProductList) | **Post** /getCloudMssqlProductList | 
[**GetDmsBackupList**](V2Api.md#GetDmsBackupList) | **Post** /getDmsBackupList | 
[**GetDmsObjectStorageBackupList**](V2Api.md#GetDmsObjectStorageBackupList) | **Post** /getDmsObjectStorageBackupList | 
[**GetDmsOperation**](V2Api.md#GetDmsOperation) | **Post** /getDmsOperation | 
[**RebootCloudMssqlServerInstance**](V2Api.md#RebootCloudMssqlServerInstance) | **Post** /rebootCloudMssqlServerInstance | 
[**RestoreDmsDatabase**](V2Api.md#RestoreDmsDatabase) | **Post** /restoreDmsDatabase | 
[**RestoreDmsTransactionLog**](V2Api.md#RestoreDmsTransactionLog) | **Post** /restoreDmsTransactionLog | 
[**SetDmsObjectStorageInfo**](V2Api.md#SetDmsObjectStorageInfo) | **Post** /setDmsObjectStorageInfo | 
[**UploadDmsFile**](V2Api.md#UploadDmsFile) | **Post** /uploadDmsFile | 


# **CreateCloudMssqlInstance**
> CreateCloudMssqlInstanceResponse CreateCloudMssqlInstance(createCloudMssqlInstanceRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createCloudMssqlInstanceRequest** | **[\*CreateCloudMssqlInstanceRequest](CreateCloudMssqlInstanceRequest.md)** | createCloudMssqlInstanceRequest | 

### Return type

*[**CreateCloudMssqlInstanceResponse**](CreateCloudMssqlInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudMssqlSlaveInstance**
> CreateCloudMssqlSlaveInstanceResponse CreateCloudMssqlSlaveInstance(createCloudMssqlSlaveInstanceRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createCloudMssqlSlaveInstanceRequest** | **[\*CreateCloudMssqlSlaveInstanceRequest](CreateCloudMssqlSlaveInstanceRequest.md)** | createCloudMssqlSlaveInstanceRequest | 

### Return type

*[**CreateCloudMssqlSlaveInstanceResponse**](CreateCloudMssqlSlaveInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudMssqlInstance**
> DeleteCloudMssqlInstanceResponse DeleteCloudMssqlInstance(deleteCloudMssqlInstanceRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteCloudMssqlInstanceRequest** | **[\*DeleteCloudMssqlInstanceRequest](DeleteCloudMssqlInstanceRequest.md)** | deleteCloudMssqlInstanceRequest | 

### Return type

*[**DeleteCloudMssqlInstanceResponse**](DeleteCloudMssqlInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudMssqlServerInstance**
> DeleteCloudMssqlServerInstanceResponse DeleteCloudMssqlServerInstance(deleteCloudMssqlServerInstanceRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteCloudMssqlServerInstanceRequest** | **[\*DeleteCloudMssqlServerInstanceRequest](DeleteCloudMssqlServerInstanceRequest.md)** | deleteCloudMssqlServerInstanceRequest | 

### Return type

*[**DeleteCloudMssqlServerInstanceResponse**](DeleteCloudMssqlServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadDmsFile**
> DownloadDmsFileResponse DownloadDmsFile(downloadDmsFileRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**downloadDmsFileRequest** | **[\*DownloadDmsFileRequest](DownloadDmsFileRequest.md)** | downloadDmsFileRequest | 

### Return type

*[**DownloadDmsFileResponse**](DownloadDmsFileResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMssqlBackupDetailList**
> GetCloudMssqlBackupDetailListResponse GetCloudMssqlBackupDetailList(getCloudMssqlBackupDetailListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMssqlBackupDetailListRequest** | **[\*GetCloudMssqlBackupDetailListRequest](GetCloudMssqlBackupDetailListRequest.md)** | getCloudMssqlBackupDetailListRequest | 

### Return type

*[**GetCloudMssqlBackupDetailListResponse**](GetCloudMssqlBackupDetailListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMssqlBackupList**
> GetCloudMssqlBackupListResponse GetCloudMssqlBackupList(getCloudMssqlBackupListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMssqlBackupListRequest** | **[\*GetCloudMssqlBackupListRequest](GetCloudMssqlBackupListRequest.md)** | getCloudMssqlBackupListRequest | 

### Return type

*[**GetCloudMssqlBackupListResponse**](GetCloudMssqlBackupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMssqlCharacterSetList**
> GetCloudMssqlCharacterSetListResponse GetCloudMssqlCharacterSetList(getCloudMssqlCharacterSetListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMssqlCharacterSetListRequest** | **[\*GetCloudMssqlCharacterSetListRequest](GetCloudMssqlCharacterSetListRequest.md)** | getCloudMssqlCharacterSetListRequest | 

### Return type

*[**GetCloudMssqlCharacterSetListResponse**](GetCloudMssqlCharacterSetListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMssqlConfigGroupList**
> GetCloudMssqlConfigGroupListResponse GetCloudMssqlConfigGroupList(getCloudMssqlConfigGroupListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMssqlConfigGroupListRequest** | **[\*GetCloudMssqlConfigGroupListRequest](GetCloudMssqlConfigGroupListRequest.md)** | getCloudMssqlConfigGroupListRequest | 

### Return type

*[**GetCloudMssqlConfigGroupListResponse**](GetCloudMssqlConfigGroupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMssqlImageProductList**
> GetCloudMssqlImageProductListResponse GetCloudMssqlImageProductList(getCloudMssqlImageProductListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMssqlImageProductListRequest** | **[\*GetCloudMssqlImageProductListRequest](GetCloudMssqlImageProductListRequest.md)** | getCloudMssqlImageProductListRequest | 

### Return type

*[**GetCloudMssqlImageProductListResponse**](GetCloudMssqlImageProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMssqlInstanceDetail**
> GetCloudMssqlInstanceDetailResponse GetCloudMssqlInstanceDetail(getCloudMssqlInstanceDetailRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMssqlInstanceDetailRequest** | **[\*GetCloudMssqlInstanceDetailRequest](GetCloudMssqlInstanceDetailRequest.md)** | getCloudMssqlInstanceDetailRequest | 

### Return type

*[**GetCloudMssqlInstanceDetailResponse**](GetCloudMssqlInstanceDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMssqlInstanceList**
> GetCloudMssqlInstanceListResponse GetCloudMssqlInstanceList(getCloudMssqlInstanceListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMssqlInstanceListRequest** | **[\*GetCloudMssqlInstanceListRequest](GetCloudMssqlInstanceListRequest.md)** | getCloudMssqlInstanceListRequest | 

### Return type

*[**GetCloudMssqlInstanceListResponse**](GetCloudMssqlInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMssqlProductList**
> GetCloudMssqlProductListResponse GetCloudMssqlProductList(getCloudMssqlProductListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMssqlProductListRequest** | **[\*GetCloudMssqlProductListRequest](GetCloudMssqlProductListRequest.md)** | getCloudMssqlProductListRequest | 

### Return type

*[**GetCloudMssqlProductListResponse**](GetCloudMssqlProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmsBackupList**
> GetDmsBackupListResponse GetDmsBackupList(getDmsBackupListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getDmsBackupListRequest** | **[\*GetDmsBackupListRequest](GetDmsBackupListRequest.md)** | getDmsBackupListRequest | 

### Return type

*[**GetDmsBackupListResponse**](GetDmsBackupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmsObjectStorageBackupList**
> GetDmsObjectStorageBackupListResponse GetDmsObjectStorageBackupList(getDmsObjectStorageBackupListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getDmsObjectStorageBackupListRequest** | **[\*GetDmsObjectStorageBackupListRequest](GetDmsObjectStorageBackupListRequest.md)** | getDmsObjectStorageBackupListRequest | 

### Return type

*[**GetDmsObjectStorageBackupListResponse**](GetDmsObjectStorageBackupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmsOperation**
> GetDmsOperationResponse GetDmsOperation(getDmsOperationRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getDmsOperationRequest** | **[\*GetDmsOperationRequest](GetDmsOperationRequest.md)** | getDmsOperationRequest | 

### Return type

*[**GetDmsOperationResponse**](GetDmsOperationResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootCloudMssqlServerInstance**
> RebootCloudMssqlServerInstanceResponse RebootCloudMssqlServerInstance(rebootCloudMssqlServerInstanceRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**rebootCloudMssqlServerInstanceRequest** | **[\*RebootCloudMssqlServerInstanceRequest](RebootCloudMssqlServerInstanceRequest.md)** | rebootCloudMssqlServerInstanceRequest | 

### Return type

*[**RebootCloudMssqlServerInstanceResponse**](RebootCloudMssqlServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreDmsDatabase**
> RestoreDmsDatabaseResponse RestoreDmsDatabase(restoreDmsDatabaseRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**restoreDmsDatabaseRequest** | **[\*RestoreDmsDatabaseRequest](RestoreDmsDatabaseRequest.md)** | restoreDmsDatabaseRequest | 

### Return type

*[**RestoreDmsDatabaseResponse**](RestoreDmsDatabaseResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreDmsTransactionLog**
> RestoreDmsTransactionLogResponse RestoreDmsTransactionLog(restoreDmsTransactionLogRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**restoreDmsTransactionLogRequest** | **[\*RestoreDmsTransactionLogRequest](RestoreDmsTransactionLogRequest.md)** | restoreDmsTransactionLogRequest | 

### Return type

*[**RestoreDmsTransactionLogResponse**](RestoreDmsTransactionLogResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetDmsObjectStorageInfo**
> SetDmsObjectStorageInfoResponse SetDmsObjectStorageInfo(setDmsObjectStorageInfoRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setDmsObjectStorageInfoRequest** | **[\*SetDmsObjectStorageInfoRequest](SetDmsObjectStorageInfoRequest.md)** | setDmsObjectStorageInfoRequest | 

### Return type

*[**SetDmsObjectStorageInfoResponse**](SetDmsObjectStorageInfoResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadDmsFile**
> UploadDmsFileResponse UploadDmsFile(uploadDmsFileRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uploadDmsFileRequest** | **[\*UploadDmsFileRequest](UploadDmsFileRequest.md)** | uploadDmsFileRequest | 

### Return type

*[**UploadDmsFileResponse**](UploadDmsFileResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

