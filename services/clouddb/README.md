# Go API client for clouddb

    Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2018-11-13T06:30:03Z
- Package version: 1.1.2
- Build package: io.swagger.codegen.languages.NcpGoForNcloudClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
"./clouddb"
```

## Documentation for API Endpoints

All URIs are relative to *https://ncloud.apigw.ntruss.com/clouddb/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V2Api* | [**CreateCloudDBInstance**](docs/V2Api.md#createclouddbinstance) | **Post** /createCloudDBInstance | 
*V2Api* | [**DeleteCloudDBServerInstance**](docs/V2Api.md#deleteclouddbserverinstance) | **Post** /deleteCloudDBServerInstance | 
*V2Api* | [**DownloadDmsFile**](docs/V2Api.md#downloaddmsfile) | **Post** /downloadDmsFile | 
*V2Api* | [**FlushCloudDBInstance**](docs/V2Api.md#flushclouddbinstance) | **Post** /flushCloudDBInstance | 
*V2Api* | [**GetBackupList**](docs/V2Api.md#getbackuplist) | **Post** /getBackupList | 
*V2Api* | [**GetCloudDBConfigGroupList**](docs/V2Api.md#getclouddbconfiggrouplist) | **Post** /getCloudDBConfigGroupList | 
*V2Api* | [**GetCloudDBImageProductList**](docs/V2Api.md#getclouddbimageproductlist) | **Post** /getCloudDBImageProductList | 
*V2Api* | [**GetCloudDBInstanceList**](docs/V2Api.md#getclouddbinstancelist) | **Post** /getCloudDBInstanceList | 
*V2Api* | [**GetCloudDBProductList**](docs/V2Api.md#getclouddbproductlist) | **Post** /getCloudDBProductList | 
*V2Api* | [**GetDmsOperation**](docs/V2Api.md#getdmsoperation) | **Post** /getDmsOperation | 
*V2Api* | [**GetObjectStorageBackupList**](docs/V2Api.md#getobjectstoragebackuplist) | **Post** /getObjectStorageBackupList | 
*V2Api* | [**RebootCloudDBServerInstance**](docs/V2Api.md#rebootclouddbserverinstance) | **Post** /rebootCloudDBServerInstance | 
*V2Api* | [**RestoreDmsDatabase**](docs/V2Api.md#restoredmsdatabase) | **Post** /restoreDmsDatabase | 
*V2Api* | [**RestoreDmsTransactionLog**](docs/V2Api.md#restoredmstransactionlog) | **Post** /restoreDmsTransactionLog | 
*V2Api* | [**SetObjectStorageInfo**](docs/V2Api.md#setobjectstorageinfo) | **Post** /setObjectStorageInfo | 
*V2Api* | [**UploadDmsFile**](docs/V2Api.md#uploaddmsfile) | **Post** /uploadDmsFile | 


## Documentation For Models

 - [AccessControlGroup](docs/AccessControlGroup.md)
 - [BackupFile](docs/BackupFile.md)
 - [CloudDbConfig](docs/CloudDbConfig.md)
 - [CloudDbConfigGroup](docs/CloudDbConfigGroup.md)
 - [CloudDbInstance](docs/CloudDbInstance.md)
 - [CloudDbServerInstance](docs/CloudDbServerInstance.md)
 - [CommonCode](docs/CommonCode.md)
 - [CreateCloudDbInstanceRequest](docs/CreateCloudDbInstanceRequest.md)
 - [CreateCloudDbInstanceResponse](docs/CreateCloudDbInstanceResponse.md)
 - [DeleteCloudDbServerInstanceRequest](docs/DeleteCloudDbServerInstanceRequest.md)
 - [DeleteCloudDbServerInstanceResponse](docs/DeleteCloudDbServerInstanceResponse.md)
 - [DmsFile](docs/DmsFile.md)
 - [DownloadDmsFileRequest](docs/DownloadDmsFileRequest.md)
 - [DownloadDmsFileResponse](docs/DownloadDmsFileResponse.md)
 - [FlushCloudDbInstanceRequest](docs/FlushCloudDbInstanceRequest.md)
 - [FlushCloudDbInstanceResponse](docs/FlushCloudDbInstanceResponse.md)
 - [GetBackupListRequest](docs/GetBackupListRequest.md)
 - [GetBackupListResponse](docs/GetBackupListResponse.md)
 - [GetCloudDbConfigGroupListRequest](docs/GetCloudDbConfigGroupListRequest.md)
 - [GetCloudDbConfigGroupListResponse](docs/GetCloudDbConfigGroupListResponse.md)
 - [GetCloudDbImageProductListRequest](docs/GetCloudDbImageProductListRequest.md)
 - [GetCloudDbImageProductListResponse](docs/GetCloudDbImageProductListResponse.md)
 - [GetCloudDbInstanceListRequest](docs/GetCloudDbInstanceListRequest.md)
 - [GetCloudDbInstanceListResponse](docs/GetCloudDbInstanceListResponse.md)
 - [GetCloudDbProductListRequest](docs/GetCloudDbProductListRequest.md)
 - [GetCloudDbProductListResponse](docs/GetCloudDbProductListResponse.md)
 - [GetDmsOperationRequest](docs/GetDmsOperationRequest.md)
 - [GetDmsOperationResponse](docs/GetDmsOperationResponse.md)
 - [GetObjectStorageBackupListRequest](docs/GetObjectStorageBackupListRequest.md)
 - [GetObjectStorageBackupListResponse](docs/GetObjectStorageBackupListResponse.md)
 - [Product](docs/Product.md)
 - [RebootCloudDbServerInstanceRequest](docs/RebootCloudDbServerInstanceRequest.md)
 - [RebootCloudDbServerInstanceResponse](docs/RebootCloudDbServerInstanceResponse.md)
 - [Region](docs/Region.md)
 - [RestoreDmsDatabaseRequest](docs/RestoreDmsDatabaseRequest.md)
 - [RestoreDmsDatabaseResponse](docs/RestoreDmsDatabaseResponse.md)
 - [RestoreDmsTransactionLogRequest](docs/RestoreDmsTransactionLogRequest.md)
 - [RestoreDmsTransactionLogResponse](docs/RestoreDmsTransactionLogResponse.md)
 - [SetObjectStorageInfoRequest](docs/SetObjectStorageInfoRequest.md)
 - [SetObjectStorageInfoResponse](docs/SetObjectStorageInfoResponse.md)
 - [UploadDmsFileRequest](docs/UploadDmsFileRequest.md)
 - [UploadDmsFileResponse](docs/UploadDmsFileResponse.md)
 - [Zone](docs/Zone.md)

