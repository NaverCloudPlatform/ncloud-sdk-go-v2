# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/vpostgresql/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCloudPostgresqlDatabaseList**](V2Api.md#AddCloudPostgresqlDatabaseList) | **Post** /addCloudPostgresqlDatabaseList | 
[**AddCloudPostgresqlUserList**](V2Api.md#AddCloudPostgresqlUserList) | **Post** /addCloudPostgresqlUserList | 
[**ChangeCloudPostgresqlUserList**](V2Api.md#ChangeCloudPostgresqlUserList) | **Post** /changeCloudPostgresqlUserList | 
[**CreateCloudPostgresqlInstance**](V2Api.md#CreateCloudPostgresqlInstance) | **Post** /createCloudPostgresqlInstance | 
[**CreateCloudPostgresqlReadReplicaInstance**](V2Api.md#CreateCloudPostgresqlReadReplicaInstance) | **Post** /createCloudPostgresqlReadReplicaInstance | 
[**DeleteCloudPostgresqlDatabaseList**](V2Api.md#DeleteCloudPostgresqlDatabaseList) | **Post** /deleteCloudPostgresqlDatabaseList | 
[**DeleteCloudPostgresqlInstance**](V2Api.md#DeleteCloudPostgresqlInstance) | **Post** /deleteCloudPostgresqlInstance | 
[**DeleteCloudPostgresqlReadReplicaInstance**](V2Api.md#DeleteCloudPostgresqlReadReplicaInstance) | **Post** /deleteCloudPostgresqlReadReplicaInstance | 
[**DeleteCloudPostgresqlUserList**](V2Api.md#DeleteCloudPostgresqlUserList) | **Post** /deleteCloudPostgresqlUserList | 
[**ExportBackupToObjectStorage**](V2Api.md#ExportBackupToObjectStorage) | **Post** /exportBackupToObjectStorage | 
[**ExportDbServerLogToObjectStorage**](V2Api.md#ExportDbServerLogToObjectStorage) | **Post** /exportDbServerLogToObjectStorage | 
[**GetCloudPostgresqlBackupDetailList**](V2Api.md#GetCloudPostgresqlBackupDetailList) | **Post** /getCloudPostgresqlBackupDetailList | 
[**GetCloudPostgresqlBackupList**](V2Api.md#GetCloudPostgresqlBackupList) | **Post** /getCloudPostgresqlBackupList | 
[**GetCloudPostgresqlBucketList**](V2Api.md#GetCloudPostgresqlBucketList) | **Post** /getCloudPostgresqlBucketList | 
[**GetCloudPostgresqlDatabaseList**](V2Api.md#GetCloudPostgresqlDatabaseList) | **Post** /getCloudPostgresqlDatabaseList | 
[**GetCloudPostgresqlImageProductList**](V2Api.md#GetCloudPostgresqlImageProductList) | **Post** /getCloudPostgresqlImageProductList | 
[**GetCloudPostgresqlInstanceDetail**](V2Api.md#GetCloudPostgresqlInstanceDetail) | **Post** /getCloudPostgresqlInstanceDetail | 
[**GetCloudPostgresqlInstanceList**](V2Api.md#GetCloudPostgresqlInstanceList) | **Post** /getCloudPostgresqlInstanceList | 
[**GetCloudPostgresqlProductList**](V2Api.md#GetCloudPostgresqlProductList) | **Post** /getCloudPostgresqlProductList | 
[**GetCloudPostgresqlTargetSubnetList**](V2Api.md#GetCloudPostgresqlTargetSubnetList) | **Post** /getCloudPostgresqlTargetSubnetList | 
[**GetCloudPostgresqlTargetVpcList**](V2Api.md#GetCloudPostgresqlTargetVpcList) | **Post** /getCloudPostgresqlTargetVpcList | 
[**GetCloudPostgresqlUserList**](V2Api.md#GetCloudPostgresqlUserList) | **Post** /getCloudPostgresqlUserList | 
[**GetDbServerLogList**](V2Api.md#GetDbServerLogList) | **Post** /getDbServerLogList | 
[**RebootCloudPostgresqlServerInstance**](V2Api.md#RebootCloudPostgresqlServerInstance) | **Post** /rebootCloudPostgresqlServerInstance | 


# **AddCloudPostgresqlDatabaseList**
> AddCloudPostgresqlDatabaseListResponse AddCloudPostgresqlDatabaseList(addCloudPostgresqlDatabaseListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addCloudPostgresqlDatabaseListRequest** | **[\*AddCloudPostgresqlDatabaseListRequest](AddCloudPostgresqlDatabaseListRequest.md)** | addCloudPostgresqlDatabaseListRequest | 

### Return type

*[**AddCloudPostgresqlDatabaseListResponse**](AddCloudPostgresqlDatabaseListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCloudPostgresqlUserList**
> AddCloudPostgresqlUserListResponse AddCloudPostgresqlUserList(addCloudPostgresqlUserListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addCloudPostgresqlUserListRequest** | **[\*AddCloudPostgresqlUserListRequest](AddCloudPostgresqlUserListRequest.md)** | addCloudPostgresqlUserListRequest | 

### Return type

*[**AddCloudPostgresqlUserListResponse**](AddCloudPostgresqlUserListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeCloudPostgresqlUserList**
> ChangeCloudPostgresqlUserListResponse ChangeCloudPostgresqlUserList(changeCloudPostgresqlUserListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**changeCloudPostgresqlUserListRequest** | **[\*ChangeCloudPostgresqlUserListRequest](ChangeCloudPostgresqlUserListRequest.md)** | changeCloudPostgresqlUserListRequest | 

### Return type

*[**ChangeCloudPostgresqlUserListResponse**](ChangeCloudPostgresqlUserListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudPostgresqlInstance**
> CreateCloudPostgresqlInstanceResponse CreateCloudPostgresqlInstance(createCloudPostgresqlInstanceRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createCloudPostgresqlInstanceRequest** | **[\*CreateCloudPostgresqlInstanceRequest](CreateCloudPostgresqlInstanceRequest.md)** | createCloudPostgresqlInstanceRequest | 

### Return type

*[**CreateCloudPostgresqlInstanceResponse**](CreateCloudPostgresqlInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudPostgresqlReadReplicaInstance**
> CreateCloudPostgresqlReadReplicaInstanceResponse CreateCloudPostgresqlReadReplicaInstance(createCloudPostgresqlReadReplicaInstanceRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createCloudPostgresqlReadReplicaInstanceRequest** | **[\*CreateCloudPostgresqlReadReplicaInstanceRequest](CreateCloudPostgresqlReadReplicaInstanceRequest.md)** | createCloudPostgresqlReadReplicaInstanceRequest | 

### Return type

*[**CreateCloudPostgresqlReadReplicaInstanceResponse**](CreateCloudPostgresqlReadReplicaInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudPostgresqlDatabaseList**
> DeleteCloudPostgresqlDatabaseListResponse DeleteCloudPostgresqlDatabaseList(deleteCloudPostgresqlDatabaseListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteCloudPostgresqlDatabaseListRequest** | **[\*DeleteCloudPostgresqlDatabaseListRequest](DeleteCloudPostgresqlDatabaseListRequest.md)** | deleteCloudPostgresqlDatabaseListRequest | 

### Return type

*[**DeleteCloudPostgresqlDatabaseListResponse**](DeleteCloudPostgresqlDatabaseListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudPostgresqlInstance**
> DeleteCloudPostgresqlInstanceResponse DeleteCloudPostgresqlInstance(deleteCloudPostgresqlInstanceRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteCloudPostgresqlInstanceRequest** | **[\*DeleteCloudPostgresqlInstanceRequest](DeleteCloudPostgresqlInstanceRequest.md)** | deleteCloudPostgresqlInstanceRequest | 

### Return type

*[**DeleteCloudPostgresqlInstanceResponse**](DeleteCloudPostgresqlInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudPostgresqlReadReplicaInstance**
> DeleteCloudPostgresqlReadReplicaInstanceResponse DeleteCloudPostgresqlReadReplicaInstance(deleteCloudPostgresqlReadReplicaInstanceRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteCloudPostgresqlReadReplicaInstanceRequest** | **[\*DeleteCloudPostgresqlReadReplicaInstanceRequest](DeleteCloudPostgresqlReadReplicaInstanceRequest.md)** | deleteCloudPostgresqlReadReplicaInstanceRequest | 

### Return type

*[**DeleteCloudPostgresqlReadReplicaInstanceResponse**](DeleteCloudPostgresqlReadReplicaInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudPostgresqlUserList**
> DeleteCloudPostgresqlUserListResponse DeleteCloudPostgresqlUserList(deleteCloudPostgresqlUserListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteCloudPostgresqlUserListRequest** | **[\*DeleteCloudPostgresqlUserListRequest](DeleteCloudPostgresqlUserListRequest.md)** | deleteCloudPostgresqlUserListRequest | 

### Return type

*[**DeleteCloudPostgresqlUserListResponse**](DeleteCloudPostgresqlUserListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportBackupToObjectStorage**
> ExportBackupToObjectStorageResponse ExportBackupToObjectStorage(exportBackupToObjectStorageRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**exportBackupToObjectStorageRequest** | **[\*ExportBackupToObjectStorageRequest](ExportBackupToObjectStorageRequest.md)** | exportBackupToObjectStorageRequest | 

### Return type

*[**ExportBackupToObjectStorageResponse**](ExportBackupToObjectStorageResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportDbServerLogToObjectStorage**
> ExportDbServerLogToObjectStorageResponse ExportDbServerLogToObjectStorage(exportDbServerLogToObjectStorageRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**exportDbServerLogToObjectStorageRequest** | **[\*ExportDbServerLogToObjectStorageRequest](ExportDbServerLogToObjectStorageRequest.md)** | exportDbServerLogToObjectStorageRequest | 

### Return type

*[**ExportDbServerLogToObjectStorageResponse**](ExportDbServerLogToObjectStorageResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPostgresqlBackupDetailList**
> GetCloudPostgresqlBackupDetailListResponse GetCloudPostgresqlBackupDetailList(getCloudPostgresqlBackupDetailListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudPostgresqlBackupDetailListRequest** | **[\*GetCloudPostgresqlBackupDetailListRequest](GetCloudPostgresqlBackupDetailListRequest.md)** | getCloudPostgresqlBackupDetailListRequest | 

### Return type

*[**GetCloudPostgresqlBackupDetailListResponse**](GetCloudPostgresqlBackupDetailListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPostgresqlBackupList**
> GetCloudPostgresqlBackupListResponse GetCloudPostgresqlBackupList(getCloudPostgresqlBackupListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudPostgresqlBackupListRequest** | **[\*GetCloudPostgresqlBackupListRequest](GetCloudPostgresqlBackupListRequest.md)** | getCloudPostgresqlBackupListRequest | 

### Return type

*[**GetCloudPostgresqlBackupListResponse**](GetCloudPostgresqlBackupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPostgresqlBucketList**
> GetCloudPostgresqlBucketListResponse GetCloudPostgresqlBucketList(getCloudPostgresqlBucketListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudPostgresqlBucketListRequest** | **[\*GetCloudPostgresqlBucketListRequest](GetCloudPostgresqlBucketListRequest.md)** | getCloudPostgresqlBucketListRequest | 

### Return type

*[**GetCloudPostgresqlBucketListResponse**](GetCloudPostgresqlBucketListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPostgresqlDatabaseList**
> GetCloudPostgresqlDatabaseListResponse GetCloudPostgresqlDatabaseList(getCloudPostgresqlDatabaseListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudPostgresqlDatabaseListRequest** | **[\*GetCloudPostgresqlDatabaseListRequest](GetCloudPostgresqlDatabaseListRequest.md)** | getCloudPostgresqlDatabaseListRequest | 

### Return type

*[**GetCloudPostgresqlDatabaseListResponse**](GetCloudPostgresqlDatabaseListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPostgresqlImageProductList**
> GetCloudPostgresqlImageProductListResponse GetCloudPostgresqlImageProductList(getCloudPostgresqlImageProductListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudPostgresqlImageProductListRequest** | **[\*GetCloudPostgresqlImageProductListRequest](GetCloudPostgresqlImageProductListRequest.md)** | getCloudPostgresqlImageProductListRequest | 

### Return type

*[**GetCloudPostgresqlImageProductListResponse**](GetCloudPostgresqlImageProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPostgresqlInstanceDetail**
> GetCloudPostgresqlInstanceDetailResponse GetCloudPostgresqlInstanceDetail(getCloudPostgresqlInstanceDetailRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudPostgresqlInstanceDetailRequest** | **[\*GetCloudPostgresqlInstanceDetailRequest](GetCloudPostgresqlInstanceDetailRequest.md)** | getCloudPostgresqlInstanceDetailRequest | 

### Return type

*[**GetCloudPostgresqlInstanceDetailResponse**](GetCloudPostgresqlInstanceDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPostgresqlInstanceList**
> GetCloudPostgresqlInstanceListResponse GetCloudPostgresqlInstanceList(getCloudPostgresqlInstanceListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudPostgresqlInstanceListRequest** | **[\*GetCloudPostgresqlInstanceListRequest](GetCloudPostgresqlInstanceListRequest.md)** | getCloudPostgresqlInstanceListRequest | 

### Return type

*[**GetCloudPostgresqlInstanceListResponse**](GetCloudPostgresqlInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPostgresqlProductList**
> GetCloudPostgresqlProductListResponse GetCloudPostgresqlProductList(getCloudPostgresqlProductListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudPostgresqlProductListRequest** | **[\*GetCloudPostgresqlProductListRequest](GetCloudPostgresqlProductListRequest.md)** | getCloudPostgresqlProductListRequest | 

### Return type

*[**GetCloudPostgresqlProductListResponse**](GetCloudPostgresqlProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPostgresqlTargetSubnetList**
> GetCloudPostgresqlTargetSubnetListResponse GetCloudPostgresqlTargetSubnetList(getCloudPostgresqlTargetSubnetListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudPostgresqlTargetSubnetListRequest** | **[\*GetCloudPostgresqlTargetSubnetListRequest](GetCloudPostgresqlTargetSubnetListRequest.md)** | getCloudPostgresqlTargetSubnetListRequest | 

### Return type

*[**GetCloudPostgresqlTargetSubnetListResponse**](GetCloudPostgresqlTargetSubnetListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPostgresqlTargetVpcList**
> GetCloudPostgresqlTargetVpcListResponse GetCloudPostgresqlTargetVpcList(getCloudPostgresqlTargetVpcListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudPostgresqlTargetVpcListRequest** | **[\*GetCloudPostgresqlTargetVpcListRequest](GetCloudPostgresqlTargetVpcListRequest.md)** | getCloudPostgresqlTargetVpcListRequest | 

### Return type

*[**GetCloudPostgresqlTargetVpcListResponse**](GetCloudPostgresqlTargetVpcListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPostgresqlUserList**
> GetCloudPostgresqlUserListResponse GetCloudPostgresqlUserList(getCloudPostgresqlUserListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudPostgresqlUserListRequest** | **[\*GetCloudPostgresqlUserListRequest](GetCloudPostgresqlUserListRequest.md)** | getCloudPostgresqlUserListRequest | 

### Return type

*[**GetCloudPostgresqlUserListResponse**](GetCloudPostgresqlUserListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDbServerLogList**
> GetDbServerLogListResponse GetDbServerLogList(getDbServerLogListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getDbServerLogListRequest** | **[\*GetDbServerLogListRequest](GetDbServerLogListRequest.md)** | getDbServerLogListRequest | 

### Return type

*[**GetDbServerLogListResponse**](GetDbServerLogListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootCloudPostgresqlServerInstance**
> RebootCloudPostgresqlServerInstanceResponse RebootCloudPostgresqlServerInstance(rebootCloudPostgresqlServerInstanceRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**rebootCloudPostgresqlServerInstanceRequest** | **[\*RebootCloudPostgresqlServerInstanceRequest](RebootCloudPostgresqlServerInstanceRequest.md)** | rebootCloudPostgresqlServerInstanceRequest | 

### Return type

*[**RebootCloudPostgresqlServerInstanceResponse**](RebootCloudPostgresqlServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

