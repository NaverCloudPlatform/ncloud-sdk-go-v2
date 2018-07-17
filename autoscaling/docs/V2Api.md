# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/autoscaling/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoScalingGroup**](V2Api.md#CreateAutoScalingGroup) | **Post** /createAutoScalingGroup | 
[**CreateLaunchConfiguration**](V2Api.md#CreateLaunchConfiguration) | **Post** /createLaunchConfiguration | 
[**DeleteAutoScalingGroup**](V2Api.md#DeleteAutoScalingGroup) | **Post** /deleteAutoScalingGroup | 
[**DeleteAutoScalingLaunchConfiguration**](V2Api.md#DeleteAutoScalingLaunchConfiguration) | **Post** /deleteAutoScalingLaunchConfiguration | 
[**DeletePolicy**](V2Api.md#DeletePolicy) | **Post** /deletePolicy | 
[**DeleteScheduledAction**](V2Api.md#DeleteScheduledAction) | **Post** /deleteScheduledAction | 
[**ExecutePolicy**](V2Api.md#ExecutePolicy) | **Post** /executePolicy | 
[**GetAdjustmentTypeList**](V2Api.md#GetAdjustmentTypeList) | **Post** /getAdjustmentTypeList | 
[**GetAutoScalingActivityLogList**](V2Api.md#GetAutoScalingActivityLogList) | **Post** /getAutoScalingActivityLogList | 
[**GetAutoScalingConfigurationLogList**](V2Api.md#GetAutoScalingConfigurationLogList) | **Post** /getAutoScalingConfigurationLogList | 
[**GetAutoScalingGroupList**](V2Api.md#GetAutoScalingGroupList) | **Post** /getAutoScalingGroupList | 
[**GetAutoScalingPolicyList**](V2Api.md#GetAutoScalingPolicyList) | **Post** /getAutoScalingPolicyList | 
[**GetLaunchConfigurationList**](V2Api.md#GetLaunchConfigurationList) | **Post** /getLaunchConfigurationList | 
[**GetScalingProcessTypeList**](V2Api.md#GetScalingProcessTypeList) | **Post** /getScalingProcessTypeList | 
[**GetScheduledActionList**](V2Api.md#GetScheduledActionList) | **Post** /getScheduledActionList | 
[**PutScalingPolicy**](V2Api.md#PutScalingPolicy) | **Post** /putScalingPolicy | 
[**PutScheduledUpdateGroupAction**](V2Api.md#PutScheduledUpdateGroupAction) | **Post** /putScheduledUpdateGroupAction | 
[**ResumeProcesses**](V2Api.md#ResumeProcesses) | **Post** /resumeProcesses | 
[**SetDesiredCapacity**](V2Api.md#SetDesiredCapacity) | **Post** /setDesiredCapacity | 
[**SetServerInstanceHealth**](V2Api.md#SetServerInstanceHealth) | **Post** /setServerInstanceHealth | 
[**SuspendProcesses**](V2Api.md#SuspendProcesses) | **Post** /suspendProcesses | 
[**TerminateServerInstanceInAutoScalingGroup**](V2Api.md#TerminateServerInstanceInAutoScalingGroup) | **Post** /terminateServerInstanceInAutoScalingGroup | 
[**UpdateAutoScalingGroup**](V2Api.md#UpdateAutoScalingGroup) | **Post** /updateAutoScalingGroup | 


# **CreateAutoScalingGroup**
> CreateAutoScalingGroupResponse CreateAutoScalingGroup(ctx, createAutoScalingGroupRequest)


B.오토스케일링그룹생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **createAutoScalingGroupRequest** | [**CreateAutoScalingGroupRequest**](CreateAutoScalingGroupRequest.md)| createAutoScalingGroupRequest | 

### Return type

[**CreateAutoScalingGroupResponse**](createAutoScalingGroupResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLaunchConfiguration**
> CreateLaunchConfigurationResponse CreateLaunchConfiguration(ctx, createLaunchConfigurationRequest)


A.론치설정생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **createLaunchConfigurationRequest** | [**CreateLaunchConfigurationRequest**](CreateLaunchConfigurationRequest.md)| createLaunchConfigurationRequest | 

### Return type

[**CreateLaunchConfigurationResponse**](createLaunchConfigurationResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAutoScalingGroup**
> DeleteAutoScalingGroupResponse DeleteAutoScalingGroup(ctx, deleteAutoScalingGroupRequest)


B.오토스케일링그룹삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deleteAutoScalingGroupRequest** | [**DeleteAutoScalingGroupRequest**](DeleteAutoScalingGroupRequest.md)| deleteAutoScalingGroupRequest | 

### Return type

[**DeleteAutoScalingGroupResponse**](deleteAutoScalingGroupResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAutoScalingLaunchConfiguration**
> DeleteAutoScalingLaunchConfigurationRequest DeleteAutoScalingLaunchConfiguration(ctx, deleteAutoScalingLaunchConfigurationRequest)


A.오토스케일링론치설정삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deleteAutoScalingLaunchConfigurationRequest** | [**DeleteAutoScalingLaunchConfigurationRequest**](DeleteAutoScalingLaunchConfigurationRequest.md)| deleteAutoScalingLaunchConfigurationRequest | 

### Return type

[**DeleteAutoScalingLaunchConfigurationRequest**](deleteAutoScalingLaunchConfigurationRequest.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicy**
> DeletePolicyResponse DeletePolicy(ctx, deletePolicyRequest)


F.스케일링정책삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deletePolicyRequest** | [**DeletePolicyRequest**](DeletePolicyRequest.md)| deletePolicyRequest | 

### Return type

[**DeletePolicyResponse**](deletePolicyResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScheduledAction**
> DeleteScheduledActionResponse DeleteScheduledAction(ctx, deleteScheduledActionRequest)


C.스케쥴액션삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deleteScheduledActionRequest** | [**DeleteScheduledActionRequest**](DeleteScheduledActionRequest.md)| deleteScheduledActionRequest | 

### Return type

[**DeleteScheduledActionResponse**](deleteScheduledActionResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutePolicy**
> ExecutePolicyResponse ExecutePolicy(ctx, executePolicyRequest)


F.스케일링정책수행

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **executePolicyRequest** | [**ExecutePolicyRequest**](ExecutePolicyRequest.md)| executePolicyRequest | 

### Return type

[**ExecutePolicyResponse**](executePolicyResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdjustmentTypeList**
> GetAdjustmentTypeListResponse GetAdjustmentTypeList(ctx, getAdjustmentTypeListRequest)


F.조정유형리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getAdjustmentTypeListRequest** | [**GetAdjustmentTypeListRequest**](GetAdjustmentTypeListRequest.md)| getAdjustmentTypeListRequest | 

### Return type

[**GetAdjustmentTypeListResponse**](getAdjustmentTypeListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoScalingActivityLogList**
> GetAutoScalingActivityLogListResponse GetAutoScalingActivityLogList(ctx, getAutoScalingActivityLogListRequest)


E.액티비티로그리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getAutoScalingActivityLogListRequest** | [**GetAutoScalingActivityLogListRequest**](GetAutoScalingActivityLogListRequest.md)| getAutoScalingActivityLogListRequest | 

### Return type

[**GetAutoScalingActivityLogListResponse**](getAutoScalingActivityLogListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoScalingConfigurationLogList**
> GetAutoScalingConfigurationLogListResponse GetAutoScalingConfigurationLogList(ctx, getAutoScalingConfigurationLogListRequest)


E.오토스케일링설정로그리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getAutoScalingConfigurationLogListRequest** | [**GetAutoScalingConfigurationLogListRequest**](GetAutoScalingConfigurationLogListRequest.md)| getAutoScalingConfigurationLogListRequest | 

### Return type

[**GetAutoScalingConfigurationLogListResponse**](getAutoScalingConfigurationLogListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoScalingGroupList**
> GetAutoScalingGroupListResponse GetAutoScalingGroupList(ctx, getAutoScalingGroupListRequest)


B.오토스케일링그룹리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getAutoScalingGroupListRequest** | [**GetAutoScalingGroupListRequest**](GetAutoScalingGroupListRequest.md)| getAutoScalingGroupListRequest | 

### Return type

[**GetAutoScalingGroupListResponse**](getAutoScalingGroupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoScalingPolicyList**
> GetAutoScalingPolicyListResponse GetAutoScalingPolicyList(ctx, getAutoScalingPolicyListRequest)


F.오토스케일링정책리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getAutoScalingPolicyListRequest** | [**GetAutoScalingPolicyListRequest**](GetAutoScalingPolicyListRequest.md)| getAutoScalingPolicyListRequest | 

### Return type

[**GetAutoScalingPolicyListResponse**](getAutoScalingPolicyListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLaunchConfigurationList**
> GetLaunchConfigurationListResponse GetLaunchConfigurationList(ctx, getLaunchConfigurationListRequest)


A.론치설정리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getLaunchConfigurationListRequest** | [**GetLaunchConfigurationListRequest**](GetLaunchConfigurationListRequest.md)| getLaunchConfigurationListRequest | 

### Return type

[**GetLaunchConfigurationListResponse**](getLaunchConfigurationListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScalingProcessTypeList**
> GetScalingProcessTypeListResponse GetScalingProcessTypeList(ctx, getScalingProcessTypeListRequest)


D.프로세스구분리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getScalingProcessTypeListRequest** | [**GetScalingProcessTypeListRequest**](GetScalingProcessTypeListRequest.md)| getScalingProcessTypeListRequest | 

### Return type

[**GetScalingProcessTypeListResponse**](getScalingProcessTypeListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScheduledActionList**
> GetScheduledActionListResponse GetScheduledActionList(ctx, getScheduledActionListRequest)


C.스케쥴액션리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getScheduledActionListRequest** | [**GetScheduledActionListRequest**](GetScheduledActionListRequest.md)| getScheduledActionListRequest | 

### Return type

[**GetScheduledActionListResponse**](getScheduledActionListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutScalingPolicy**
> PutScalingPolicyResponse PutScalingPolicy(ctx, putScalingPolicyRequest)


F.스케일링정책생성/변경

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **putScalingPolicyRequest** | [**PutScalingPolicyRequest**](PutScalingPolicyRequest.md)| putScalingPolicyRequest | 

### Return type

[**PutScalingPolicyResponse**](putScalingPolicyResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutScheduledUpdateGroupAction**
> PutScheduledUpdateGroupActionResponse PutScheduledUpdateGroupAction(ctx, putScheduledUpdateGroupActionRequest)


C.스케쥴액션생성|수정

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **putScheduledUpdateGroupActionRequest** | [**PutScheduledUpdateGroupActionRequest**](PutScheduledUpdateGroupActionRequest.md)| putScheduledUpdateGroupActionRequest | 

### Return type

[**PutScheduledUpdateGroupActionResponse**](putScheduledUpdateGroupActionResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResumeProcesses**
> ResumeProcessesResponse ResumeProcesses(ctx, resumeProcessesRequest)


D.프로세스재개

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **resumeProcessesRequest** | [**ResumeProcessesRequest**](ResumeProcessesRequest.md)| resumeProcessesRequest | 

### Return type

[**ResumeProcessesResponse**](resumeProcessesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetDesiredCapacity**
> SetDesiredCapacityResponse SetDesiredCapacity(ctx, setDesiredCapacityRequest)


B.기대용량치갱신

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **setDesiredCapacityRequest** | [**SetDesiredCapacityRequest**](SetDesiredCapacityRequest.md)| setDesiredCapacityRequest | 

### Return type

[**SetDesiredCapacityResponse**](setDesiredCapacityResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetServerInstanceHealth**
> SetServerInstanceHealthResponse SetServerInstanceHealth(ctx, setServerInstanceHealthRequest)


B.서버인스턴스헬스상태갱신

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **setServerInstanceHealthRequest** | [**SetServerInstanceHealthRequest**](SetServerInstanceHealthRequest.md)| setServerInstanceHealthRequest | 

### Return type

[**SetServerInstanceHealthResponse**](setServerInstanceHealthResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuspendProcesses**
> SuspendProcessesResponse SuspendProcesses(ctx, suspendProcessesRequest)


D.프로세스보류

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **suspendProcessesRequest** | [**SuspendProcessesRequest**](SuspendProcessesRequest.md)| suspendProcessesRequest | 

### Return type

[**SuspendProcessesResponse**](suspendProcessesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminateServerInstanceInAutoScalingGroup**
> TerminateServerInstanceInAutoScalingGroupResponse TerminateServerInstanceInAutoScalingGroup(ctx, terminateServerInstanceInAutoScalingGroupRequest)


B.오토스케일링그룹에속한서버인스턴스삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **terminateServerInstanceInAutoScalingGroupRequest** | [**TerminateServerInstanceInAutoScalingGroupRequest**](TerminateServerInstanceInAutoScalingGroupRequest.md)| terminateServerInstanceInAutoScalingGroupRequest | 

### Return type

[**TerminateServerInstanceInAutoScalingGroupResponse**](terminateServerInstanceInAutoScalingGroupResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAutoScalingGroup**
> UpdateAutoScalingGroupResponse UpdateAutoScalingGroup(ctx, updateAutoScalingGroupRequest)


B.오토스케일링그룹수정

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **updateAutoScalingGroupRequest** | [**UpdateAutoScalingGroupRequest**](UpdateAutoScalingGroupRequest.md)| updateAutoScalingGroupRequest | 

### Return type

[**UpdateAutoScalingGroupResponse**](updateAutoScalingGroupResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

