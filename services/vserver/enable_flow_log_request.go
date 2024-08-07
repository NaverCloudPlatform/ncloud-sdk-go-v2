/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type EnableFlowLogRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 수집액션유형코드
CollectActionTypeCode *string `json:"collectActionTypeCode"`

	// 수집주기(분)
CollectIntervalMinute *int32 `json:"collectIntervalMinute,omitempty"`

	// 네트워크인터페이스번호
NetworkInterfaceNo *string `json:"networkInterfaceNo"`

	// 저장소유형코드
StorageTypeCode *string `json:"storageTypeCode,omitempty"`

	// 저장소버킷이름
StorageBucketName *string `json:"storageBucketName"`

	// 저장소버킷디렉토리이름
StorageBucketDirectoryName *string `json:"storageBucketDirectoryName,omitempty"`
}
