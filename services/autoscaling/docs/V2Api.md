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
> CreateAutoScalingGroupResponse CreateAutoScalingGroup(createAutoScalingGroupRequest)


B.오토스케일링그룹생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createAutoScalingGroupRequest** | **[\*CreateAutoScalingGroupRequest](CreateAutoScalingGroupRequest.md)** | createAutoScalingGroupRequest | 

### Return type

*[**CreateAutoScalingGroupResponse**](CreateAutoScalingGroupResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLaunchConfiguration**
> CreateLaunchConfigurationResponse CreateLaunchConfiguration(createLaunchConfigurationRequest)


A.론치설정생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createLaunchConfigurationRequest** | **[\*CreateLaunchConfigurationRequest](CreateLaunchConfigurationRequest.md)** | createLaunchConfigurationRequest | 

### Return type

*[**CreateLaunchConfigurationResponse**](CreateLaunchConfigurationResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAutoScalingGroup**
> DeleteAutoScalingGroupResponse DeleteAutoScalingGroup(deleteAutoScalingGroupRequest)


B.오토스케일링그룹삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteAutoScalingGroupRequest** | **[\*DeleteAutoScalingGroupRequest](DeleteAutoScalingGroupRequest.md)** | deleteAutoScalingGroupRequest | 

### Return type

*[**DeleteAutoScalingGroupResponse**](DeleteAutoScalingGroupResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAutoScalingLaunchConfiguration**
> DeleteAutoScalingLaunchConfigurationRequest DeleteAutoScalingLaunchConfiguration(deleteAutoScalingLaunchConfigurationRequest)


A.오토스케일링론치설정삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteAutoScalingLaunchConfigurationRequest** | **[\*DeleteAutoScalingLaunchConfigurationRequest](DeleteAutoScalingLaunchConfigurationRequest.md)** | deleteAutoScalingLaunchConfigurationRequest | 

### Return type

*[**DeleteAutoScalingLaunchConfigurationRequest**](DeleteAutoScalingLaunchConfigurationRequest.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicy**
> DeletePolicyResponse DeletePolicy(deletePolicyRequest)


F.스케일링정책삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deletePolicyRequest** | **[\*DeletePolicyRequest](DeletePolicyRequest.md)** | deletePolicyRequest | 

### Return type

*[**DeletePolicyResponse**](DeletePolicyResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScheduledAction**
> DeleteScheduledActionResponse DeleteScheduledAction(deleteScheduledActionRequest)


C.스케쥴액션삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteScheduledActionRequest** | **[\*DeleteScheduledActionRequest](DeleteScheduledActionRequest.md)** | deleteScheduledActionRequest | 

### Return type

*[**DeleteScheduledActionResponse**](DeleteScheduledActionResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutePolicy**
> ExecutePolicyResponse ExecutePolicy(executePolicyRequest)


F.스케일링정책수행

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**executePolicyRequest** | **[\*ExecutePolicyRequest](ExecutePolicyRequest.md)** | executePolicyRequest | 

### Return type

*[**ExecutePolicyResponse**](ExecutePolicyResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdjustmentTypeList**
> GetAdjustmentTypeListResponse GetAdjustmentTypeList(getAdjustmentTypeListRequest)


F.조정유형리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getAdjustmentTypeListRequest** | **[\*GetAdjustmentTypeListRequest](GetAdjustmentTypeListRequest.md)** | getAdjustmentTypeListRequest | 

### Return type

*[**GetAdjustmentTypeListResponse**](GetAdjustmentTypeListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoScalingActivityLogList**
> GetAutoScalingActivityLogListResponse GetAutoScalingActivityLogList(getAutoScalingActivityLogListRequest)


E.액티비티로그리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getAutoScalingActivityLogListRequest** | **[\*GetAutoScalingActivityLogListRequest](GetAutoScalingActivityLogListRequest.md)** | getAutoScalingActivityLogListRequest | 

### Return type

*[**GetAutoScalingActivityLogListResponse**](GetAutoScalingActivityLogListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoScalingConfigurationLogList**
> GetAutoScalingConfigurationLogListResponse GetAutoScalingConfigurationLogList(getAutoScalingConfigurationLogListRequest)


E.오토스케일링설정로그리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getAutoScalingConfigurationLogListRequest** | **[\*GetAutoScalingConfigurationLogListRequest](GetAutoScalingConfigurationLogListRequest.md)** | getAutoScalingConfigurationLogListRequest | 

### Return type

*[**GetAutoScalingConfigurationLogListResponse**](GetAutoScalingConfigurationLogListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoScalingGroupList**
> GetAutoScalingGroupListResponse GetAutoScalingGroupList(getAutoScalingGroupListRequest)


B.오토스케일링그룹리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getAutoScalingGroupListRequest** | **[\*GetAutoScalingGroupListRequest](GetAutoScalingGroupListRequest.md)** | getAutoScalingGroupListRequest | 

### Return type

*[**GetAutoScalingGroupListResponse**](GetAutoScalingGroupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoScalingPolicyList**
> GetAutoScalingPolicyListResponse GetAutoScalingPolicyList(getAutoScalingPolicyListRequest)


F.오토스케일링정책리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getAutoScalingPolicyListRequest** | **[\*GetAutoScalingPolicyListRequest](GetAutoScalingPolicyListRequest.md)** | getAutoScalingPolicyListRequest | 

### Return type

*[**GetAutoScalingPolicyListResponse**](GetAutoScalingPolicyListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLaunchConfigurationList**
> GetLaunchConfigurationListResponse GetLaunchConfigurationList(getLaunchConfigurationListRequest)


A.론치설정리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getLaunchConfigurationListRequest** | **[\*GetLaunchConfigurationListRequest](GetLaunchConfigurationListRequest.md)** | getLaunchConfigurationListRequest | 

### Return type

*[**GetLaunchConfigurationListResponse**](GetLaunchConfigurationListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScalingProcessTypeList**
> GetScalingProcessTypeListResponse GetScalingProcessTypeList(getScalingProcessTypeListRequest)


D.프로세스구분리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getScalingProcessTypeListRequest** | **[\*GetScalingProcessTypeListRequest](GetScalingProcessTypeListRequest.md)** | getScalingProcessTypeListRequest | 

### Return type

*[**GetScalingProcessTypeListResponse**](GetScalingProcessTypeListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScheduledActionList**
> GetScheduledActionListResponse GetScheduledActionList(getScheduledActionListRequest)


C.스케쥴액션리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getScheduledActionListRequest** | **[\*GetScheduledActionListRequest](GetScheduledActionListRequest.md)** | getScheduledActionListRequest | 

### Return type

*[**GetScheduledActionListResponse**](GetScheduledActionListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutScalingPolicy**
> PutScalingPolicyResponse PutScalingPolicy(putScalingPolicyRequest)


F.스케일링정책생성/변경

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**putScalingPolicyRequest** | **[\*PutScalingPolicyRequest](PutScalingPolicyRequest.md)** | putScalingPolicyRequest | 

### Return type

*[**PutScalingPolicyResponse**](PutScalingPolicyResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutScheduledUpdateGroupAction**
> PutScheduledUpdateGroupActionResponse PutScheduledUpdateGroupAction(putScheduledUpdateGroupActionRequest)


C.스케쥴액션생성|수정

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**putScheduledUpdateGroupActionRequest** | **[\*PutScheduledUpdateGroupActionRequest](PutScheduledUpdateGroupActionRequest.md)** | putScheduledUpdateGroupActionRequest | 

### Return type

*[**PutScheduledUpdateGroupActionResponse**](PutScheduledUpdateGroupActionResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResumeProcesses**
> ResumeProcessesResponse ResumeProcesses(resumeProcessesRequest)


D.프로세스재개

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**resumeProcessesRequest** | **[\*ResumeProcessesRequest](ResumeProcessesRequest.md)** | resumeProcessesRequest | 

### Return type

*[**ResumeProcessesResponse**](ResumeProcessesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetDesiredCapacity**
> SetDesiredCapacityResponse SetDesiredCapacity(setDesiredCapacityRequest)


B.기대용량치갱신

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setDesiredCapacityRequest** | **[\*SetDesiredCapacityRequest](SetDesiredCapacityRequest.md)** | setDesiredCapacityRequest | 

### Return type

*[**SetDesiredCapacityResponse**](SetDesiredCapacityResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetServerInstanceHealth**
> SetServerInstanceHealthResponse SetServerInstanceHealth(setServerInstanceHealthRequest)


B.서버인스턴스헬스상태갱신

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setServerInstanceHealthRequest** | **[\*SetServerInstanceHealthRequest](SetServerInstanceHealthRequest.md)** | setServerInstanceHealthRequest | 

### Return type

*[**SetServerInstanceHealthResponse**](SetServerInstanceHealthResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuspendProcesses**
> SuspendProcessesResponse SuspendProcesses(suspendProcessesRequest)


D.프로세스보류

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**suspendProcessesRequest** | **[\*SuspendProcessesRequest](SuspendProcessesRequest.md)** | suspendProcessesRequest | 

### Return type

*[**SuspendProcessesResponse**](SuspendProcessesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminateServerInstanceInAutoScalingGroup**
> TerminateServerInstanceInAutoScalingGroupResponse TerminateServerInstanceInAutoScalingGroup(terminateServerInstanceInAutoScalingGroupRequest)


B.오토스케일링그룹에속한서버인스턴스삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**terminateServerInstanceInAutoScalingGroupRequest** | **[\*TerminateServerInstanceInAutoScalingGroupRequest](TerminateServerInstanceInAutoScalingGroupRequest.md)** | terminateServerInstanceInAutoScalingGroupRequest | 

### Return type

*[**TerminateServerInstanceInAutoScalingGroupResponse**](TerminateServerInstanceInAutoScalingGroupResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAutoScalingGroup**
> UpdateAutoScalingGroupResponse UpdateAutoScalingGroup(updateAutoScalingGroupRequest)


B.오토스케일링그룹수정

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**updateAutoScalingGroupRequest** | **[\*UpdateAutoScalingGroupRequest](UpdateAutoScalingGroupRequest.md)** | updateAutoScalingGroupRequest | 

### Return type

*[**UpdateAutoScalingGroupResponse**](UpdateAutoScalingGroupResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

