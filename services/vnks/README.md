# Go API client for vnks

    <br/>https://nks.apigw.ntruss.com/vnks/v2

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2024-09-26T09:01:52Z
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.NcpGoForVnksClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
"./vnks"
```

## Documentation for API Endpoints

All URIs are relative to *https://nks.apigw.ntruss.com/vnks/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V2Api* | [**ClustersGet**](docs/V2Api.md#clustersget) | **Get** /clusters | 
*V2Api* | [**ClustersPost**](docs/V2Api.md#clusterspost) | **Post** /clusters | 
*V2Api* | [**ClustersUuidAccessEntriesEntryUuidDelete**](docs/V2Api.md#clustersuuidaccessentriesentryuuiddelete) | **Delete** /clusters/{uuid}/access-entries/{entryUuid} | 
*V2Api* | [**ClustersUuidAccessEntriesEntryUuidGet**](docs/V2Api.md#clustersuuidaccessentriesentryuuidget) | **Get** /clusters/{uuid}/access-entries/{entryUuid} | 
*V2Api* | [**ClustersUuidAccessEntriesEntryUuidPut**](docs/V2Api.md#clustersuuidaccessentriesentryuuidput) | **Put** /clusters/{uuid}/access-entries/{entryUuid} | 
*V2Api* | [**ClustersUuidAccessEntriesGet**](docs/V2Api.md#clustersuuidaccessentriesget) | **Get** /clusters/{uuid}/access-entries | 
*V2Api* | [**ClustersUuidAccessEntriesPost**](docs/V2Api.md#clustersuuidaccessentriespost) | **Post** /clusters/{uuid}/access-entries | 
*V2Api* | [**ClustersUuidAddSubnetPatch**](docs/V2Api.md#clustersuuidaddsubnetpatch) | **Patch** /clusters/{uuid}/add-subnet | 
*V2Api* | [**ClustersUuidAuthTypePatch**](docs/V2Api.md#clustersuuidauthtypepatch) | **Patch** /clusters/{uuid}/auth-type | 
*V2Api* | [**ClustersUuidDelete**](docs/V2Api.md#clustersuuiddelete) | **Delete** /clusters/{uuid} | 
*V2Api* | [**ClustersUuidGet**](docs/V2Api.md#clustersuuidget) | **Get** /clusters/{uuid} | 
*V2Api* | [**ClustersUuidIpAclGet**](docs/V2Api.md#clustersuuidipaclget) | **Get** /clusters/{uuid}/ip-acl | 
*V2Api* | [**ClustersUuidIpAclPatch**](docs/V2Api.md#clustersuuidipaclpatch) | **Patch** /clusters/{uuid}/ip-acl | 
*V2Api* | [**ClustersUuidKubeconfigGet**](docs/V2Api.md#clustersuuidkubeconfigget) | **Get** /clusters/{uuid}/kubeconfig | 
*V2Api* | [**ClustersUuidKubeconfigResetPatch**](docs/V2Api.md#clustersuuidkubeconfigresetpatch) | **Patch** /clusters/{uuid}/kubeconfig/reset | 
*V2Api* | [**ClustersUuidLbSubnetPatch**](docs/V2Api.md#clustersuuidlbsubnetpatch) | **Patch** /clusters/{uuid}/lb-subnet | 
*V2Api* | [**ClustersUuidLogPatch**](docs/V2Api.md#clustersuuidlogpatch) | **Patch** /clusters/{uuid}/log | 
*V2Api* | [**ClustersUuidNodePoolGet**](docs/V2Api.md#clustersuuidnodepoolget) | **Get** /clusters/{uuid}/node-pool | 
*V2Api* | [**ClustersUuidNodePoolInstanceNoDelete**](docs/V2Api.md#clustersuuidnodepoolinstancenodelete) | **Delete** /clusters/{uuid}/node-pool/{instanceNo} | 
*V2Api* | [**ClustersUuidNodePoolInstanceNoLabelsPut**](docs/V2Api.md#clustersuuidnodepoolinstancenolabelsput) | **Put** /clusters/{uuid}/node-pool/{instanceNo}/labels | 
*V2Api* | [**ClustersUuidNodePoolInstanceNoPatch**](docs/V2Api.md#clustersuuidnodepoolinstancenopatch) | **Patch** /clusters/{uuid}/node-pool/{instanceNo} | 
*V2Api* | [**ClustersUuidNodePoolInstanceNoSubnetsPatch**](docs/V2Api.md#clustersuuidnodepoolinstancenosubnetspatch) | **Patch** /clusters/{uuid}/node-pool/{instanceNo}/subnets | 
*V2Api* | [**ClustersUuidNodePoolInstanceNoTaintsPut**](docs/V2Api.md#clustersuuidnodepoolinstancenotaintsput) | **Put** /clusters/{uuid}/node-pool/{instanceNo}/taints | 
*V2Api* | [**ClustersUuidNodePoolInstanceNoUpgradePatch**](docs/V2Api.md#clustersuuidnodepoolinstancenoupgradepatch) | **Patch** /clusters/{uuid}/node-pool/{instanceNo}/upgrade | 
*V2Api* | [**ClustersUuidNodePoolPost**](docs/V2Api.md#clustersuuidnodepoolpost) | **Post** /clusters/{uuid}/node-pool | 
*V2Api* | [**ClustersUuidNodesGet**](docs/V2Api.md#clustersuuidnodesget) | **Get** /clusters/{uuid}/nodes | 
*V2Api* | [**ClustersUuidNodesInstanceNoDelete**](docs/V2Api.md#clustersuuidnodesinstancenodelete) | **Delete** /clusters/{uuid}/nodes/{instanceNo} | 
*V2Api* | [**ClustersUuidOidcGet**](docs/V2Api.md#clustersuuidoidcget) | **Get** /clusters/{uuid}/oidc | 
*V2Api* | [**ClustersUuidOidcPatch**](docs/V2Api.md#clustersuuidoidcpatch) | **Patch** /clusters/{uuid}/oidc | 
*V2Api* | [**ClustersUuidReturnProtectionPatch**](docs/V2Api.md#clustersuuidreturnprotectionpatch) | **Patch** /clusters/{uuid}/return-protection | 
*V2Api* | [**ClustersUuidSecretEncryptionPatch**](docs/V2Api.md#clustersuuidsecretencryptionpatch) | **Patch** /clusters/{uuid}/secret-encryption | 
*V2Api* | [**ClustersUuidUpgradePatch**](docs/V2Api.md#clustersuuidupgradepatch) | **Patch** /clusters/{uuid}/upgrade | 
*V2Api* | [**OptionServerImageGet**](docs/V2Api.md#optionserverimageget) | **Get** /option/server-image | 
*V2Api* | [**OptionServerProductCodeGet**](docs/V2Api.md#optionserverproductcodeget) | **Get** /option/server-product-code | 
*V2Api* | [**OptionVersionGet**](docs/V2Api.md#optionversionget) | **Get** /option/version | 
*V2Api* | [**RootGet**](docs/V2Api.md#rootget) | **Get** / | 


## Documentation For Models

 - [AccessEntryPolicyRes](docs/AccessEntryPolicyRes.md)
 - [AccessEntryRes](docs/AccessEntryRes.md)
 - [AddSubnetDto](docs/AddSubnetDto.md)
 - [AuditLogDto](docs/AuditLogDto.md)
 - [AutoscaleOption](docs/AutoscaleOption.md)
 - [AutoscalerUpdate](docs/AutoscalerUpdate.md)
 - [Cluster](docs/Cluster.md)
 - [ClusterAccessEntriesRes](docs/ClusterAccessEntriesRes.md)
 - [ClusterAccessEntryRes](docs/ClusterAccessEntryRes.md)
 - [ClusterInputBody](docs/ClusterInputBody.md)
 - [ClusterLogInput](docs/ClusterLogInput.md)
 - [ClusterRes](docs/ClusterRes.md)
 - [ClustersRes](docs/ClustersRes.md)
 - [CreateAccessEntryDto](docs/CreateAccessEntryDto.md)
 - [CreateAccessEntryPolicyDto](docs/CreateAccessEntryPolicyDto.md)
 - [CreateAccessEntryRes](docs/CreateAccessEntryRes.md)
 - [CreateClusterRes](docs/CreateClusterRes.md)
 - [DefaultNodePoolParam](docs/DefaultNodePoolParam.md)
 - [DeleteAccessEntryRes](docs/DeleteAccessEntryRes.md)
 - [IpAclsDto](docs/IpAclsDto.md)
 - [IpAclsEntriesDto](docs/IpAclsEntriesDto.md)
 - [IpAclsEntriesRes](docs/IpAclsEntriesRes.md)
 - [IpAclsRes](docs/IpAclsRes.md)
 - [KubeconfigRes](docs/KubeconfigRes.md)
 - [NodePool](docs/NodePool.md)
 - [NodePoolCreationBody](docs/NodePoolCreationBody.md)
 - [NodePoolDto](docs/NodePoolDto.md)
 - [NodePoolLabel](docs/NodePoolLabel.md)
 - [NodePoolRes](docs/NodePoolRes.md)
 - [NodePoolTaint](docs/NodePoolTaint.md)
 - [NodePoolUpdateBody](docs/NodePoolUpdateBody.md)
 - [OidcRes](docs/OidcRes.md)
 - [OptionRes](docs/OptionRes.md)
 - [OptionResForServerProduct](docs/OptionResForServerProduct.md)
 - [OptionsRes](docs/OptionsRes.md)
 - [OptionsResForServerProduct](docs/OptionsResForServerProduct.md)
 - [ReturnProtectionDto](docs/ReturnProtectionDto.md)
 - [ServerProduct](docs/ServerProduct.md)
 - [SubnetDto](docs/SubnetDto.md)
 - [UpdateAccessEntryDto](docs/UpdateAccessEntryDto.md)
 - [UpdateAccessEntryRes](docs/UpdateAccessEntryRes.md)
 - [UpdateAuthTypeDto](docs/UpdateAuthTypeDto.md)
 - [UpdateClusterLbSubnetRes](docs/UpdateClusterLbSubnetRes.md)
 - [UpdateClusterRes](docs/UpdateClusterRes.md)
 - [UpdateNodePoolRes](docs/UpdateNodePoolRes.md)
 - [UpdateNodepoolLabelDto](docs/UpdateNodepoolLabelDto.md)
 - [UpdateNodepoolSubnetDto](docs/UpdateNodepoolSubnetDto.md)
 - [UpdateNodepoolTaintDto](docs/UpdateNodepoolTaintDto.md)
 - [UpdateOidcDto](docs/UpdateOidcDto.md)
 - [UpdateSecretEncryptionDto](docs/UpdateSecretEncryptionDto.md)
 - [WorkerNode](docs/WorkerNode.md)
 - [WorkerNodeRes](docs/WorkerNodeRes.md)

