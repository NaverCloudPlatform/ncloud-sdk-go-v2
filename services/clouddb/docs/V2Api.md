# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/clouddb/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudDBInstance**](V2Api.md#CreateCloudDBInstance) | **Post** /createCloudDBInstance | 
[**DeleteCloudDBServerInstance**](V2Api.md#DeleteCloudDBServerInstance) | **Post** /deleteCloudDBServerInstance | 
[**DownloadDmsFile**](V2Api.md#DownloadDmsFile) | **Post** /downloadDmsFile | 
[**FlushCloudDBInstance**](V2Api.md#FlushCloudDBInstance) | **Post** /flushCloudDBInstance | 
[**GetBackupList**](V2Api.md#GetBackupList) | **Post** /getBackupList | 
[**GetCloudDBConfigGroupList**](V2Api.md#GetCloudDBConfigGroupList) | **Post** /getCloudDBConfigGroupList | 
[**GetCloudDBImageProductList**](V2Api.md#GetCloudDBImageProductList) | **Post** /getCloudDBImageProductList | 
[**GetCloudDBInstanceList**](V2Api.md#GetCloudDBInstanceList) | **Post** /getCloudDBInstanceList | 
[**GetCloudDBProductList**](V2Api.md#GetCloudDBProductList) | **Post** /getCloudDBProductList | 
[**GetDmsOperation**](V2Api.md#GetDmsOperation) | **Post** /getDmsOperation | 
[**GetObjectStorageBackupList**](V2Api.md#GetObjectStorageBackupList) | **Post** /getObjectStorageBackupList | 
[**RebootCloudDBServerInstance**](V2Api.md#RebootCloudDBServerInstance) | **Post** /rebootCloudDBServerInstance | 
[**RestoreDmsDatabase**](V2Api.md#RestoreDmsDatabase) | **Post** /restoreDmsDatabase | 
[**RestoreDmsTransactionLog**](V2Api.md#RestoreDmsTransactionLog) | **Post** /restoreDmsTransactionLog | 
[**SetObjectStorageInfo**](V2Api.md#SetObjectStorageInfo) | **Post** /setObjectStorageInfo | 
[**UploadDmsFile**](V2Api.md#UploadDmsFile) | **Post** /uploadDmsFile | 


# **CreateCloudDBInstance**
> CreateCloudDbInstanceResponse CreateCloudDBInstance(createCloudDBInstanceRequest)


CloudDB인스턴스생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createCloudDBInstanceRequest** | **[\*CreateCloudDbInstanceRequest](CreateCloudDbInstanceRequest.md)** | createCloudDBInstanceRequest | 

### Return type

*[**CreateCloudDbInstanceResponse**](CreateCloudDbInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudDBServerInstance**
> DeleteCloudDbServerInstanceResponse DeleteCloudDBServerInstance(deleteCloudDBServerInstanceRequest)


CloudDB서버인스턴스삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteCloudDBServerInstanceRequest** | **[\*DeleteCloudDbServerInstanceRequest](DeleteCloudDbServerInstanceRequest.md)** | deleteCloudDBServerInstanceRequest | 

### Return type

*[**DeleteCloudDbServerInstanceResponse**](DeleteCloudDbServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadDmsFile**
> DownloadDmsFileResponse DownloadDmsFile(downloadDmsFileRequest)


DMS파일다운로드

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

# **FlushCloudDBInstance**
> FlushCloudDbInstanceResponse FlushCloudDBInstance(flushCloudDBInstanceRequest)


CloudDB Flush

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**flushCloudDBInstanceRequest** | **[\*FlushCloudDbInstanceRequest](FlushCloudDbInstanceRequest.md)** | flushCloudDBInstanceRequest | 

### Return type

*[**FlushCloudDbInstanceResponse**](FlushCloudDbInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupList**
> GetBackupListResponse GetBackupList(getBackupListRequest)


백업리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getBackupListRequest** | **[\*GetBackupListRequest](GetBackupListRequest.md)** | getBackupListRequest | 

### Return type

*[**GetBackupListResponse**](GetBackupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudDBConfigGroupList**
> GetCloudDbConfigGroupListResponse GetCloudDBConfigGroupList(getCloudDBConfigGroupListRequest)


CloudDB설정그룹리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudDBConfigGroupListRequest** | **[\*GetCloudDbConfigGroupListRequest](GetCloudDbConfigGroupListRequest.md)** | getCloudDBConfigGroupListRequest | 

### Return type

*[**GetCloudDbConfigGroupListResponse**](GetCloudDbConfigGroupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudDBImageProductList**
> GetCloudDbImageProductListResponse GetCloudDBImageProductList(getCloudDBImageProductListRequest)


CloudDB이미지상품리스트

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudDBImageProductListRequest** | **[\*GetCloudDbImageProductListRequest](GetCloudDbImageProductListRequest.md)** | getCloudDBImageProductListRequest | 

### Return type

*[**GetCloudDbImageProductListResponse**](GetCloudDbImageProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudDBInstanceList**
> GetCloudDbInstanceListResponse GetCloudDBInstanceList(getCloudDBInstanceListRequest)


CloudDB인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudDBInstanceListRequest** | **[\*GetCloudDbInstanceListRequest](GetCloudDbInstanceListRequest.md)** | getCloudDBInstanceListRequest | 

### Return type

*[**GetCloudDbInstanceListResponse**](GetCloudDbInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudDBProductList**
> GetCloudDbProductListResponse GetCloudDBProductList(getCloudDBProductListRequest)


CloudDB상품리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudDBProductListRequest** | **[\*GetCloudDbProductListRequest](GetCloudDbProductListRequest.md)** | getCloudDBProductListRequest | 

### Return type

*[**GetCloudDbProductListResponse**](GetCloudDbProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmsOperation**
> GetDmsOperationResponse GetDmsOperation(getDmsOperationRequest)


DMS상태조회

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

# **GetObjectStorageBackupList**
> GetObjectStorageBackupListResponse GetObjectStorageBackupList(getObjectStorageBackupListRequest)


오브젝트스토리지백업리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getObjectStorageBackupListRequest** | **[\*GetObjectStorageBackupListRequest](GetObjectStorageBackupListRequest.md)** | getObjectStorageBackupListRequest | 

### Return type

*[**GetObjectStorageBackupListResponse**](GetObjectStorageBackupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootCloudDBServerInstance**
> RebootCloudDbServerInstanceResponse RebootCloudDBServerInstance(rebootCloudDBServerInstanceRequest)


CloudDB서버인스턴스재부팅

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**rebootCloudDBServerInstanceRequest** | **[\*RebootCloudDbServerInstanceRequest](RebootCloudDbServerInstanceRequest.md)** | rebootCloudDBServerInstanceRequest | 

### Return type

*[**RebootCloudDbServerInstanceResponse**](RebootCloudDbServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreDmsDatabase**
> RestoreDmsDatabaseResponse RestoreDmsDatabase(restoreDmsDatabaseRequest)


DMS데이터베이스복구

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


DMS트랜잭션로그복구

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

# **SetObjectStorageInfo**
> SetObjectStorageInfoResponse SetObjectStorageInfo(setObjectStorageInfoRequest)


오브젝트스토리지정보설정

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setObjectStorageInfoRequest** | **[\*SetObjectStorageInfoRequest](SetObjectStorageInfoRequest.md)** | setObjectStorageInfoRequest | 

### Return type

*[**SetObjectStorageInfoResponse**](SetObjectStorageInfoResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadDmsFile**
> UploadDmsFileResponse UploadDmsFile(uploadDmsFileRequest)


DMS파일업로드

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

