# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/vpc/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptOrRejectVpcPeering**](V2Api.md#AcceptOrRejectVpcPeering) | **Post** /acceptOrRejectVpcPeering | 
[**AddNetworkAclInboundRule**](V2Api.md#AddNetworkAclInboundRule) | **Post** /addNetworkAclInboundRule | 
[**AddNetworkAclOutboundRule**](V2Api.md#AddNetworkAclOutboundRule) | **Post** /addNetworkAclOutboundRule | 
[**AddRoute**](V2Api.md#AddRoute) | **Post** /addRoute | 
[**AddRouteTableSubnet**](V2Api.md#AddRouteTableSubnet) | **Post** /addRouteTableSubnet | 
[**CreateNatGatewayInstance**](V2Api.md#CreateNatGatewayInstance) | **Post** /createNatGatewayInstance | 
[**CreateNetworkAcl**](V2Api.md#CreateNetworkAcl) | **Post** /createNetworkAcl | 
[**CreateRouteTable**](V2Api.md#CreateRouteTable) | **Post** /createRouteTable | 
[**CreateSubnet**](V2Api.md#CreateSubnet) | **Post** /createSubnet | 
[**CreateVpc**](V2Api.md#CreateVpc) | **Post** /createVpc | 
[**CreateVpcPeeringInstance**](V2Api.md#CreateVpcPeeringInstance) | **Post** /createVpcPeeringInstance | 
[**DeleteNatGatewayInstance**](V2Api.md#DeleteNatGatewayInstance) | **Post** /deleteNatGatewayInstance | 
[**DeleteNetworkAcl**](V2Api.md#DeleteNetworkAcl) | **Post** /deleteNetworkAcl | 
[**DeleteRouteTable**](V2Api.md#DeleteRouteTable) | **Post** /deleteRouteTable | 
[**DeleteSubnet**](V2Api.md#DeleteSubnet) | **Post** /deleteSubnet | 
[**DeleteVpc**](V2Api.md#DeleteVpc) | **Post** /deleteVpc | 
[**DeleteVpcPeeringInstance**](V2Api.md#DeleteVpcPeeringInstance) | **Post** /deleteVpcPeeringInstance | 
[**GetNatGatewayInstanceDetail**](V2Api.md#GetNatGatewayInstanceDetail) | **Post** /getNatGatewayInstanceDetail | 
[**GetNatGatewayInstanceList**](V2Api.md#GetNatGatewayInstanceList) | **Post** /getNatGatewayInstanceList | 
[**GetNetworkAclDetail**](V2Api.md#GetNetworkAclDetail) | **Post** /getNetworkAclDetail | 
[**GetNetworkAclList**](V2Api.md#GetNetworkAclList) | **Post** /getNetworkAclList | 
[**GetNetworkAclRuleList**](V2Api.md#GetNetworkAclRuleList) | **Post** /getNetworkAclRuleList | 
[**GetRouteList**](V2Api.md#GetRouteList) | **Post** /getRouteList | 
[**GetRouteTableDetail**](V2Api.md#GetRouteTableDetail) | **Post** /getRouteTableDetail | 
[**GetRouteTableList**](V2Api.md#GetRouteTableList) | **Post** /getRouteTableList | 
[**GetRouteTableSubnetList**](V2Api.md#GetRouteTableSubnetList) | **Post** /getRouteTableSubnetList | 
[**GetSubnetDetail**](V2Api.md#GetSubnetDetail) | **Post** /getSubnetDetail | 
[**GetSubnetList**](V2Api.md#GetSubnetList) | **Post** /getSubnetList | 
[**GetVpcDetail**](V2Api.md#GetVpcDetail) | **Post** /getVpcDetail | 
[**GetVpcList**](V2Api.md#GetVpcList) | **Post** /getVpcList | 
[**GetVpcPeeringInstanceDetail**](V2Api.md#GetVpcPeeringInstanceDetail) | **Post** /getVpcPeeringInstanceDetail | 
[**GetVpcPeeringInstanceList**](V2Api.md#GetVpcPeeringInstanceList) | **Post** /getVpcPeeringInstanceList | 
[**RemoveNetworkAclInboundRule**](V2Api.md#RemoveNetworkAclInboundRule) | **Post** /removeNetworkAclInboundRule | 
[**RemoveNetworkAclOutboundRule**](V2Api.md#RemoveNetworkAclOutboundRule) | **Post** /removeNetworkAclOutboundRule | 
[**RemoveRoute**](V2Api.md#RemoveRoute) | **Post** /removeRoute | 
[**RemoveRouteTableSubnet**](V2Api.md#RemoveRouteTableSubnet) | **Post** /removeRouteTableSubnet | 
[**SetNatGatewayDescription**](V2Api.md#SetNatGatewayDescription) | **Post** /setNatGatewayDescription | 
[**SetNetworkAclDescription**](V2Api.md#SetNetworkAclDescription) | **Post** /setNetworkAclDescription | 
[**SetRouteTableDescription**](V2Api.md#SetRouteTableDescription) | **Post** /setRouteTableDescription | 
[**SetSubnetNetworkAcl**](V2Api.md#SetSubnetNetworkAcl) | **Post** /setSubnetNetworkAcl | 
[**SetVpcPeeringDescription**](V2Api.md#SetVpcPeeringDescription) | **Post** /setVpcPeeringDescription | 


# **AcceptOrRejectVpcPeering**
> AcceptOrRejectVpcPeeringResponse AcceptOrRejectVpcPeering(acceptOrRejectVpcPeeringRequest)


VPCPeering요청수락거절

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**acceptOrRejectVpcPeeringRequest** | **[\*AcceptOrRejectVpcPeeringRequest](AcceptOrRejectVpcPeeringRequest.md)** | acceptOrRejectVpcPeeringRequest | 

### Return type

*[**AcceptOrRejectVpcPeeringResponse**](AcceptOrRejectVpcPeeringResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddNetworkAclInboundRule**
> AddNetworkAclInboundRuleResponse AddNetworkAclInboundRule(addNetworkAclInboundRuleRequest)


네트워크ACLInboundRule추가

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addNetworkAclInboundRuleRequest** | **[\*AddNetworkAclInboundRuleRequest](AddNetworkAclInboundRuleRequest.md)** | addNetworkAclInboundRuleRequest | 

### Return type

*[**AddNetworkAclInboundRuleResponse**](AddNetworkAclInboundRuleResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddNetworkAclOutboundRule**
> AddNetworkAclOutboundRuleResponse AddNetworkAclOutboundRule(addNetworkAclOutboundRuleRequest)


네트워크ACLOutboundRule추가

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addNetworkAclOutboundRuleRequest** | **[\*AddNetworkAclOutboundRuleRequest](AddNetworkAclOutboundRuleRequest.md)** | addNetworkAclOutboundRuleRequest | 

### Return type

*[**AddNetworkAclOutboundRuleResponse**](AddNetworkAclOutboundRuleResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRoute**
> AddRouteResponse AddRoute(addRouteRequest)


라우트추가

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addRouteRequest** | **[\*AddRouteRequest](AddRouteRequest.md)** | addRouteRequest | 

### Return type

*[**AddRouteResponse**](AddRouteResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRouteTableSubnet**
> AddRouteTableSubnetResponse AddRouteTableSubnet(addRouteTableSubnetRequest)


라우트테이블의연관서브넷추가

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addRouteTableSubnetRequest** | **[\*AddRouteTableSubnetRequest](AddRouteTableSubnetRequest.md)** | addRouteTableSubnetRequest | 

### Return type

*[**AddRouteTableSubnetResponse**](AddRouteTableSubnetResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNatGatewayInstance**
> CreateNatGatewayInstanceResponse CreateNatGatewayInstance(createNatGatewayInstanceRequest)


NATGateway인스턴스생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createNatGatewayInstanceRequest** | **[\*CreateNatGatewayInstanceRequest](CreateNatGatewayInstanceRequest.md)** | createNatGatewayInstanceRequest | 

### Return type

*[**CreateNatGatewayInstanceResponse**](CreateNatGatewayInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNetworkAcl**
> CreateNetworkAclResponse CreateNetworkAcl(createNetworkAclRequest)


네트워크ACL생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createNetworkAclRequest** | **[\*CreateNetworkAclRequest](CreateNetworkAclRequest.md)** | createNetworkAclRequest | 

### Return type

*[**CreateNetworkAclResponse**](CreateNetworkAclResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRouteTable**
> CreateRouteTableResponse CreateRouteTable(createRouteTableRequest)


라우트테이블생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createRouteTableRequest** | **[\*CreateRouteTableRequest](CreateRouteTableRequest.md)** | createRouteTableRequest | 

### Return type

*[**CreateRouteTableResponse**](CreateRouteTableResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubnet**
> CreateSubnetResponse CreateSubnet(createSubnetRequest)


서브넷생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createSubnetRequest** | **[\*CreateSubnetRequest](CreateSubnetRequest.md)** | createSubnetRequest | 

### Return type

*[**CreateSubnetResponse**](CreateSubnetResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVpc**
> CreateVpcResponse CreateVpc(createVpcRequest)


VPC생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createVpcRequest** | **[\*CreateVpcRequest](CreateVpcRequest.md)** | createVpcRequest | 

### Return type

*[**CreateVpcResponse**](CreateVpcResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVpcPeeringInstance**
> CreateVpcPeeringInstanceResponse CreateVpcPeeringInstance(createVpcPeeringInstanceRequest)


VPCPeering인스턴스생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createVpcPeeringInstanceRequest** | **[\*CreateVpcPeeringInstanceRequest](CreateVpcPeeringInstanceRequest.md)** | createVpcPeeringInstanceRequest | 

### Return type

*[**CreateVpcPeeringInstanceResponse**](CreateVpcPeeringInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNatGatewayInstance**
> DeleteNatGatewayInstanceResponse DeleteNatGatewayInstance(deleteNatGatewayInstanceRequest)


NATGateway인스턴스삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteNatGatewayInstanceRequest** | **[\*DeleteNatGatewayInstanceRequest](DeleteNatGatewayInstanceRequest.md)** | deleteNatGatewayInstanceRequest | 

### Return type

*[**DeleteNatGatewayInstanceResponse**](DeleteNatGatewayInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkAcl**
> DeleteNetworkAclResponse DeleteNetworkAcl(deleteNetworkAclRequest)


네트워크ACL삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteNetworkAclRequest** | **[\*DeleteNetworkAclRequest](DeleteNetworkAclRequest.md)** | deleteNetworkAclRequest | 

### Return type

*[**DeleteNetworkAclResponse**](DeleteNetworkAclResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouteTable**
> DeleteRouteTableResponse DeleteRouteTable(deleteRouteTableRequest)


라우트테이블삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteRouteTableRequest** | **[\*DeleteRouteTableRequest](DeleteRouteTableRequest.md)** | deleteRouteTableRequest | 

### Return type

*[**DeleteRouteTableResponse**](DeleteRouteTableResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubnet**
> DeleteSubnetResponse DeleteSubnet(deleteSubnetRequest)


서브넷삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteSubnetRequest** | **[\*DeleteSubnetRequest](DeleteSubnetRequest.md)** | deleteSubnetRequest | 

### Return type

*[**DeleteSubnetResponse**](DeleteSubnetResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVpc**
> DeleteVpcResponse DeleteVpc(deleteVpcRequest)


VPC삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteVpcRequest** | **[\*DeleteVpcRequest](DeleteVpcRequest.md)** | deleteVpcRequest | 

### Return type

*[**DeleteVpcResponse**](DeleteVpcResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVpcPeeringInstance**
> DeleteVpcPeeringInstanceResponse DeleteVpcPeeringInstance(deleteVpcPeeringInstanceRequest)


VPCPeering인스턴스삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteVpcPeeringInstanceRequest** | **[\*DeleteVpcPeeringInstanceRequest](DeleteVpcPeeringInstanceRequest.md)** | deleteVpcPeeringInstanceRequest | 

### Return type

*[**DeleteVpcPeeringInstanceResponse**](DeleteVpcPeeringInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNatGatewayInstanceDetail**
> GetNatGatewayInstanceDetailResponse GetNatGatewayInstanceDetail(getNatGatewayInstanceDetailRequest)


NATGateway인스턴스상세조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getNatGatewayInstanceDetailRequest** | **[\*GetNatGatewayInstanceDetailRequest](GetNatGatewayInstanceDetailRequest.md)** | getNatGatewayInstanceDetailRequest | 

### Return type

*[**GetNatGatewayInstanceDetailResponse**](GetNatGatewayInstanceDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNatGatewayInstanceList**
> GetNatGatewayInstanceListResponse GetNatGatewayInstanceList(getNatGatewayInstanceListRequest)


NATGateway인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getNatGatewayInstanceListRequest** | **[\*GetNatGatewayInstanceListRequest](GetNatGatewayInstanceListRequest.md)** | getNatGatewayInstanceListRequest | 

### Return type

*[**GetNatGatewayInstanceListResponse**](GetNatGatewayInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkAclDetail**
> GetNetworkAclDetailResponse GetNetworkAclDetail(getNetworkAclDetailRequest)


네트워크ACL상세조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getNetworkAclDetailRequest** | **[\*GetNetworkAclDetailRequest](GetNetworkAclDetailRequest.md)** | getNetworkAclDetailRequest | 

### Return type

*[**GetNetworkAclDetailResponse**](GetNetworkAclDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkAclList**
> GetNetworkAclListResponse GetNetworkAclList(getNetworkAclListRequest)


네트워크ACL리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getNetworkAclListRequest** | **[\*GetNetworkAclListRequest](GetNetworkAclListRequest.md)** | getNetworkAclListRequest | 

### Return type

*[**GetNetworkAclListResponse**](GetNetworkAclListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkAclRuleList**
> GetNetworkAclRuleListResponse GetNetworkAclRuleList(getNetworkAclRuleListRequest)


네트워크ACLRule리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getNetworkAclRuleListRequest** | **[\*GetNetworkAclRuleListRequest](GetNetworkAclRuleListRequest.md)** | getNetworkAclRuleListRequest | 

### Return type

*[**GetNetworkAclRuleListResponse**](GetNetworkAclRuleListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteList**
> GetRouteListResponse GetRouteList(getRouteListRequest)


라우트리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getRouteListRequest** | **[\*GetRouteListRequest](GetRouteListRequest.md)** | getRouteListRequest | 

### Return type

*[**GetRouteListResponse**](GetRouteListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteTableDetail**
> GetRouteTableDetailResponse GetRouteTableDetail(getRouteTableDetailRequest)


라우트테이블상세조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getRouteTableDetailRequest** | **[\*GetRouteTableDetailRequest](GetRouteTableDetailRequest.md)** | getRouteTableDetailRequest | 

### Return type

*[**GetRouteTableDetailResponse**](GetRouteTableDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteTableList**
> GetRouteTableListResponse GetRouteTableList(getRouteTableListRequest)


라우트테이블리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getRouteTableListRequest** | **[\*GetRouteTableListRequest](GetRouteTableListRequest.md)** | getRouteTableListRequest | 

### Return type

*[**GetRouteTableListResponse**](GetRouteTableListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteTableSubnetList**
> GetRouteTableSubnetListResponse GetRouteTableSubnetList(getRouteTableSubnetListRequest)


라우트테이블에연관된서브넷리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getRouteTableSubnetListRequest** | **[\*GetRouteTableSubnetListRequest](GetRouteTableSubnetListRequest.md)** | getRouteTableSubnetListRequest | 

### Return type

*[**GetRouteTableSubnetListResponse**](GetRouteTableSubnetListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubnetDetail**
> GetSubnetDetailResponse GetSubnetDetail(getSubnetDetailRequest)


서브넷상세조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getSubnetDetailRequest** | **[\*GetSubnetDetailRequest](GetSubnetDetailRequest.md)** | getSubnetDetailRequest | 

### Return type

*[**GetSubnetDetailResponse**](GetSubnetDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubnetList**
> GetSubnetListResponse GetSubnetList(getSubnetListRequest)


서브넷리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getSubnetListRequest** | **[\*GetSubnetListRequest](GetSubnetListRequest.md)** | getSubnetListRequest | 

### Return type

*[**GetSubnetListResponse**](GetSubnetListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpcDetail**
> GetVpcDetailResponse GetVpcDetail(getVpcDetailRequest)


VPC상세조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getVpcDetailRequest** | **[\*GetVpcDetailRequest](GetVpcDetailRequest.md)** | getVpcDetailRequest | 

### Return type

*[**GetVpcDetailResponse**](GetVpcDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpcList**
> GetVpcListResponse GetVpcList(getVpcListRequest)


VPC리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getVpcListRequest** | **[\*GetVpcListRequest](GetVpcListRequest.md)** | getVpcListRequest | 

### Return type

*[**GetVpcListResponse**](GetVpcListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpcPeeringInstanceDetail**
> GetVpcPeeringInstanceDetailResponse GetVpcPeeringInstanceDetail(getVpcPeeringInstanceDetailRequest)


VPCPeering인스턴스상세조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getVpcPeeringInstanceDetailRequest** | **[\*GetVpcPeeringInstanceDetailRequest](GetVpcPeeringInstanceDetailRequest.md)** | getVpcPeeringInstanceDetailRequest | 

### Return type

*[**GetVpcPeeringInstanceDetailResponse**](GetVpcPeeringInstanceDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpcPeeringInstanceList**
> GetVpcPeeringInstanceListResponse GetVpcPeeringInstanceList(getVpcPeeringInstanceListRequest)


VPCPeering인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getVpcPeeringInstanceListRequest** | **[\*GetVpcPeeringInstanceListRequest](GetVpcPeeringInstanceListRequest.md)** | getVpcPeeringInstanceListRequest | 

### Return type

*[**GetVpcPeeringInstanceListResponse**](GetVpcPeeringInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveNetworkAclInboundRule**
> RemoveNetworkAclInboundRuleResponse RemoveNetworkAclInboundRule(removeNetworkAclInboundRuleRequest)


네트워크ACLInboundRule제거

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**removeNetworkAclInboundRuleRequest** | **[\*RemoveNetworkAclInboundRuleRequest](RemoveNetworkAclInboundRuleRequest.md)** | removeNetworkAclInboundRuleRequest | 

### Return type

*[**RemoveNetworkAclInboundRuleResponse**](RemoveNetworkAclInboundRuleResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveNetworkAclOutboundRule**
> RemoveNetworkAclOutboundRuleResponse RemoveNetworkAclOutboundRule(removeNetworkAclOutboundRuleRequest)


네트워크ACLOutboundRule제거

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**removeNetworkAclOutboundRuleRequest** | **[\*RemoveNetworkAclOutboundRuleRequest](RemoveNetworkAclOutboundRuleRequest.md)** | removeNetworkAclOutboundRuleRequest | 

### Return type

*[**RemoveNetworkAclOutboundRuleResponse**](RemoveNetworkAclOutboundRuleResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRoute**
> RemoveRouteResponse RemoveRoute(removeRouteRequest)


라우트제거

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**removeRouteRequest** | **[\*RemoveRouteRequest](RemoveRouteRequest.md)** | removeRouteRequest | 

### Return type

*[**RemoveRouteResponse**](RemoveRouteResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRouteTableSubnet**
> RemoveRouteTableSubnetResponse RemoveRouteTableSubnet(removeRouteTableSubnetRequest)


라우트테이블의연관서브넷제거

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**removeRouteTableSubnetRequest** | **[\*RemoveRouteTableSubnetRequest](RemoveRouteTableSubnetRequest.md)** | removeRouteTableSubnetRequest | 

### Return type

*[**RemoveRouteTableSubnetResponse**](RemoveRouteTableSubnetResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetNatGatewayDescription**
> SetNatGatewayDescriptionResponse SetNatGatewayDescription(setNatGatewayDescriptionRequest)


NATGateway설명설정

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setNatGatewayDescriptionRequest** | **[\*SetNatGatewayDescriptionRequest](SetNatGatewayDescriptionRequest.md)** | setNatGatewayDescriptionRequest | 

### Return type

*[**SetNatGatewayDescriptionResponse**](SetNatGatewayDescriptionResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetNetworkAclDescription**
> SetNetworkAclDescriptionResponse SetNetworkAclDescription(setNetworkAclDescriptionRequest)


네트워크ACL설명설정

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setNetworkAclDescriptionRequest** | **[\*SetNetworkAclDescriptionRequest](SetNetworkAclDescriptionRequest.md)** | setNetworkAclDescriptionRequest | 

### Return type

*[**SetNetworkAclDescriptionResponse**](SetNetworkAclDescriptionResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRouteTableDescription**
> SetRouteTableDescriptionResponse SetRouteTableDescription(setRouteTableDescriptionRequest)


라우트테이블설명설정

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setRouteTableDescriptionRequest** | **[\*SetRouteTableDescriptionRequest](SetRouteTableDescriptionRequest.md)** | setRouteTableDescriptionRequest | 

### Return type

*[**SetRouteTableDescriptionResponse**](SetRouteTableDescriptionResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSubnetNetworkAcl**
> SetSubnetNetworkAclResponse SetSubnetNetworkAcl(setSubnetNetworkAclRequest)


서브넷의네트워크ACL설정

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setSubnetNetworkAclRequest** | **[\*SetSubnetNetworkAclRequest](SetSubnetNetworkAclRequest.md)** | setSubnetNetworkAclRequest | 

### Return type

*[**SetSubnetNetworkAclResponse**](SetSubnetNetworkAclResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetVpcPeeringDescription**
> SetVpcPeeringDescriptionResponse SetVpcPeeringDescription(setVpcPeeringDescriptionRequest)


VPCPeering설명설정

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setVpcPeeringDescriptionRequest** | **[\*SetVpcPeeringDescriptionRequest](SetVpcPeeringDescriptionRequest.md)** | setVpcPeeringDescriptionRequest | 

### Return type

*[**SetVpcPeeringDescriptionResponse**](SetVpcPeeringDescriptionResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

