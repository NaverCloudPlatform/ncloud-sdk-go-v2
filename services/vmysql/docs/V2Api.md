# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/vmysql/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCloudMysqlDatabaseList**](V2Api.md#AddCloudMysqlDatabaseList) | **Post** /addCloudMysqlDatabaseList | 
[**AddCloudMysqlUserList**](V2Api.md#AddCloudMysqlUserList) | **Post** /addCloudMysqlUserList | 
[**ChangeCloudMysqlUserList**](V2Api.md#ChangeCloudMysqlUserList) | **Post** /changeCloudMysqlUserList | 
[**CreateCloudMysqlInstance**](V2Api.md#CreateCloudMysqlInstance) | **Post** /createCloudMysqlInstance | 
[**CreateCloudMysqlRecoveryInstance**](V2Api.md#CreateCloudMysqlRecoveryInstance) | **Post** /createCloudMysqlRecoveryInstance | 
[**CreateCloudMysqlSlaveInstance**](V2Api.md#CreateCloudMysqlSlaveInstance) | **Post** /createCloudMysqlSlaveInstance | 
[**DeleteCloudMysqlDatabaseList**](V2Api.md#DeleteCloudMysqlDatabaseList) | **Post** /deleteCloudMysqlDatabaseList | 
[**DeleteCloudMysqlInstance**](V2Api.md#DeleteCloudMysqlInstance) | **Post** /deleteCloudMysqlInstance | 
[**DeleteCloudMysqlServerInstance**](V2Api.md#DeleteCloudMysqlServerInstance) | **Post** /deleteCloudMysqlServerInstance | 
[**DeleteCloudMysqlUserList**](V2Api.md#DeleteCloudMysqlUserList) | **Post** /deleteCloudMysqlUserList | 
[**ExportBackupToObjectStorage**](V2Api.md#ExportBackupToObjectStorage) | **Post** /exportBackupToObjectStorage | 
[**ExportDbServerLogToObjectStorage**](V2Api.md#ExportDbServerLogToObjectStorage) | **Post** /exportDbServerLogToObjectStorage | 
[**GetCloudMysqlBackupDetailList**](V2Api.md#GetCloudMysqlBackupDetailList) | **Post** /getCloudMysqlBackupDetailList | 
[**GetCloudMysqlBackupList**](V2Api.md#GetCloudMysqlBackupList) | **Post** /getCloudMysqlBackupList | 
[**GetCloudMysqlDatabaseList**](V2Api.md#GetCloudMysqlDatabaseList) | **Post** /getCloudMysqlDatabaseList | 
[**GetCloudMysqlImageProductList**](V2Api.md#GetCloudMysqlImageProductList) | **Post** /getCloudMysqlImageProductList | 
[**GetCloudMysqlInstanceDetail**](V2Api.md#GetCloudMysqlInstanceDetail) | **Post** /getCloudMysqlInstanceDetail | 
[**GetCloudMysqlInstanceList**](V2Api.md#GetCloudMysqlInstanceList) | **Post** /getCloudMysqlInstanceList | 
[**GetCloudMysqlProductList**](V2Api.md#GetCloudMysqlProductList) | **Post** /getCloudMysqlProductList | 
[**GetCloudMysqlRecoveryTime**](V2Api.md#GetCloudMysqlRecoveryTime) | **Post** /getCloudMysqlRecoveryTime | 
[**GetCloudMysqlTargetSubnetList**](V2Api.md#GetCloudMysqlTargetSubnetList) | **Post** /getCloudMysqlTargetSubnetList | 
[**GetCloudMysqlTargetVpcList**](V2Api.md#GetCloudMysqlTargetVpcList) | **Post** /getCloudMysqlTargetVpcList | 
[**GetCloudMysqlUserList**](V2Api.md#GetCloudMysqlUserList) | **Post** /getCloudMysqlUserList | 
[**GetDbServerLogList**](V2Api.md#GetDbServerLogList) | **Post** /getDbServerLogList | 
[**RebootCloudMysqlServerInstance**](V2Api.md#RebootCloudMysqlServerInstance) | **Post** /rebootCloudMysqlServerInstance | 


# **AddCloudMysqlDatabaseList**
> AddCloudMysqlDatabaseListResponse AddCloudMysqlDatabaseList(addCloudMysqlDatabaseListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addCloudMysqlDatabaseListRequest** | **[\*AddCloudMysqlDatabaseListRequest](AddCloudMysqlDatabaseListRequest.md)** | addCloudMysqlDatabaseListRequest | 

### Return type

*[**AddCloudMysqlDatabaseListResponse**](AddCloudMysqlDatabaseListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCloudMysqlUserList**
> AddCloudMysqlUserListResponse AddCloudMysqlUserList(addCloudMysqlUserListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addCloudMysqlUserListRequest** | **[\*AddCloudMysqlUserListRequest](AddCloudMysqlUserListRequest.md)** | addCloudMysqlUserListRequest | 

### Return type

*[**AddCloudMysqlUserListResponse**](AddCloudMysqlUserListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeCloudMysqlUserList**
> ChangeCloudMysqlUserListResponse ChangeCloudMysqlUserList(changeCloudMysqlUserListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**changeCloudMysqlUserListRequest** | **[\*ChangeCloudMysqlUserListRequest](ChangeCloudMysqlUserListRequest.md)** | changeCloudMysqlUserListRequest | 

### Return type

*[**ChangeCloudMysqlUserListResponse**](ChangeCloudMysqlUserListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudMysqlInstance**
> CreateCloudMysqlInstanceResponse CreateCloudMysqlInstance(createCloudMysqlInstanceRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createCloudMysqlInstanceRequest** | **[\*CreateCloudMysqlInstanceRequest](CreateCloudMysqlInstanceRequest.md)** | createCloudMysqlInstanceRequest | 

### Return type

*[**CreateCloudMysqlInstanceResponse**](CreateCloudMysqlInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudMysqlRecoveryInstance**
> CreateCloudMysqlRecoveryInstanceResponse CreateCloudMysqlRecoveryInstance(createCloudMysqlRecoveryInstanceRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createCloudMysqlRecoveryInstanceRequest** | **[\*CreateCloudMysqlRecoveryInstanceRequest](CreateCloudMysqlRecoveryInstanceRequest.md)** | createCloudMysqlRecoveryInstanceRequest | 

### Return type

*[**CreateCloudMysqlRecoveryInstanceResponse**](CreateCloudMysqlRecoveryInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudMysqlSlaveInstance**
> CreateCloudMysqlSlaveInstanceResponse CreateCloudMysqlSlaveInstance(createCloudMysqlSlaveInstanceRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createCloudMysqlSlaveInstanceRequest** | **[\*CreateCloudMysqlSlaveInstanceRequest](CreateCloudMysqlSlaveInstanceRequest.md)** | createCloudMysqlSlaveInstanceRequest | 

### Return type

*[**CreateCloudMysqlSlaveInstanceResponse**](CreateCloudMysqlSlaveInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudMysqlDatabaseList**
> DeleteCloudMysqlDatabaseListResponse DeleteCloudMysqlDatabaseList(deleteCloudMysqlDatabaseListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteCloudMysqlDatabaseListRequest** | **[\*DeleteCloudMysqlDatabaseListRequest](DeleteCloudMysqlDatabaseListRequest.md)** | deleteCloudMysqlDatabaseListRequest | 

### Return type

*[**DeleteCloudMysqlDatabaseListResponse**](DeleteCloudMysqlDatabaseListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudMysqlInstance**
> DeleteCloudMysqlInstanceResponse DeleteCloudMysqlInstance(deleteCloudMysqlInstanceRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteCloudMysqlInstanceRequest** | **[\*DeleteCloudMysqlInstanceRequest](DeleteCloudMysqlInstanceRequest.md)** | deleteCloudMysqlInstanceRequest | 

### Return type

*[**DeleteCloudMysqlInstanceResponse**](DeleteCloudMysqlInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudMysqlServerInstance**
> DeleteCloudMysqlServerInstanceResponse DeleteCloudMysqlServerInstance(deleteCloudMysqlServerInstanceRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteCloudMysqlServerInstanceRequest** | **[\*DeleteCloudMysqlServerInstanceRequest](DeleteCloudMysqlServerInstanceRequest.md)** | deleteCloudMysqlServerInstanceRequest | 

### Return type

*[**DeleteCloudMysqlServerInstanceResponse**](DeleteCloudMysqlServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudMysqlUserList**
> DeleteCloudMysqlUserListResponse DeleteCloudMysqlUserList(deleteCloudMysqlUserListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteCloudMysqlUserListRequest** | **[\*DeleteCloudMysqlUserListRequest](DeleteCloudMysqlUserListRequest.md)** | deleteCloudMysqlUserListRequest | 

### Return type

*[**DeleteCloudMysqlUserListResponse**](DeleteCloudMysqlUserListResponse.md)

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

# **GetCloudMysqlBackupDetailList**
> GetCloudMysqlBackupDetailListResponse GetCloudMysqlBackupDetailList(getCloudMysqlBackupDetailListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMysqlBackupDetailListRequest** | **[\*GetCloudMysqlBackupDetailListRequest](GetCloudMysqlBackupDetailListRequest.md)** | getCloudMysqlBackupDetailListRequest | 

### Return type

*[**GetCloudMysqlBackupDetailListResponse**](GetCloudMysqlBackupDetailListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMysqlBackupList**
> GetCloudMysqlBackupListResponse GetCloudMysqlBackupList(getCloudMysqlBackupListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMysqlBackupListRequest** | **[\*GetCloudMysqlBackupListRequest](GetCloudMysqlBackupListRequest.md)** | getCloudMysqlBackupListRequest | 

### Return type

*[**GetCloudMysqlBackupListResponse**](GetCloudMysqlBackupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMysqlDatabaseList**
> GetCloudMysqlDatabaseListResponse GetCloudMysqlDatabaseList(getCloudMysqlDatabaseListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMysqlDatabaseListRequest** | **[\*GetCloudMysqlDatabaseListRequest](GetCloudMysqlDatabaseListRequest.md)** | getCloudMysqlDatabaseListRequest | 

### Return type

*[**GetCloudMysqlDatabaseListResponse**](GetCloudMysqlDatabaseListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMysqlImageProductList**
> GetCloudMysqlImageProductListResponse GetCloudMysqlImageProductList(getCloudMysqlImageProductListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMysqlImageProductListRequest** | **[\*GetCloudMysqlImageProductListRequest](GetCloudMysqlImageProductListRequest.md)** | getCloudMysqlImageProductListRequest | 

### Return type

*[**GetCloudMysqlImageProductListResponse**](GetCloudMysqlImageProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMysqlInstanceDetail**
> GetCloudMysqlInstanceDetailResponse GetCloudMysqlInstanceDetail(getCloudMysqlInstanceDetailRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMysqlInstanceDetailRequest** | **[\*GetCloudMysqlInstanceDetailRequest](GetCloudMysqlInstanceDetailRequest.md)** | getCloudMysqlInstanceDetailRequest | 

### Return type

*[**GetCloudMysqlInstanceDetailResponse**](GetCloudMysqlInstanceDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMysqlInstanceList**
> GetCloudMysqlInstanceListResponse GetCloudMysqlInstanceList(getCloudMysqlInstanceListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMysqlInstanceListRequest** | **[\*GetCloudMysqlInstanceListRequest](GetCloudMysqlInstanceListRequest.md)** | getCloudMysqlInstanceListRequest | 

### Return type

*[**GetCloudMysqlInstanceListResponse**](GetCloudMysqlInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMysqlProductList**
> GetCloudMysqlProductListResponse GetCloudMysqlProductList(getCloudMysqlProductListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMysqlProductListRequest** | **[\*GetCloudMysqlProductListRequest](GetCloudMysqlProductListRequest.md)** | getCloudMysqlProductListRequest | 

### Return type

*[**GetCloudMysqlProductListResponse**](GetCloudMysqlProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMysqlRecoveryTime**
> GetCloudMysqlRecoveryTimeResponse GetCloudMysqlRecoveryTime(getCloudMysqlRecoveryTimeRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMysqlRecoveryTimeRequest** | **[\*GetCloudMysqlRecoveryTimeRequest](GetCloudMysqlRecoveryTimeRequest.md)** | getCloudMysqlRecoveryTimeRequest | 

### Return type

*[**GetCloudMysqlRecoveryTimeResponse**](GetCloudMysqlRecoveryTimeResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMysqlTargetSubnetList**
> GetCloudMysqlTargetSubnetListResponse GetCloudMysqlTargetSubnetList(getCloudMysqlTargetSubnetListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMysqlTargetSubnetListRequest** | **[\*GetCloudMysqlTargetSubnetListRequest](GetCloudMysqlTargetSubnetListRequest.md)** | getCloudMysqlTargetSubnetListRequest | 

### Return type

*[**GetCloudMysqlTargetSubnetListResponse**](GetCloudMysqlTargetSubnetListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMysqlTargetVpcList**
> GetCloudMysqlTargetVpcListResponse GetCloudMysqlTargetVpcList(getCloudMysqlTargetVpcListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMysqlTargetVpcListRequest** | **[\*GetCloudMysqlTargetVpcListRequest](GetCloudMysqlTargetVpcListRequest.md)** | getCloudMysqlTargetVpcListRequest | 

### Return type

*[**GetCloudMysqlTargetVpcListResponse**](GetCloudMysqlTargetVpcListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudMysqlUserList**
> GetCloudMysqlUserListResponse GetCloudMysqlUserList(getCloudMysqlUserListRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getCloudMysqlUserListRequest** | **[\*GetCloudMysqlUserListRequest](GetCloudMysqlUserListRequest.md)** | getCloudMysqlUserListRequest | 

### Return type

*[**GetCloudMysqlUserListResponse**](GetCloudMysqlUserListResponse.md)

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

# **RebootCloudMysqlServerInstance**
> RebootCloudMysqlServerInstanceResponse RebootCloudMysqlServerInstance(rebootCloudMysqlServerInstanceRequest)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**rebootCloudMysqlServerInstanceRequest** | **[\*RebootCloudMysqlServerInstanceRequest](RebootCloudMysqlServerInstanceRequest.md)** | rebootCloudMysqlServerInstanceRequest | 

### Return type

*[**RebootCloudMysqlServerInstanceResponse**](RebootCloudMysqlServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

