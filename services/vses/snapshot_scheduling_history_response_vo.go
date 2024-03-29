/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type SnapshotSchedulingHistoryResponseVo struct {
	BucketName string `json:"bucketName,omitempty"`
	SnapshotName string `json:"snapshotName,omitempty"`
	SnapshotSchedulingConfigSetUpDate string `json:"snapshotSchedulingConfigSetUpDate,omitempty"`
	SnapshotSchedulingSetUpDay string `json:"snapshotSchedulingSetUpDay,omitempty"`
	SnapshotSchedulingSetUpTime string `json:"snapshotSchedulingSetUpTime,omitempty"`
	SnapshotSchedulingStatus string `json:"snapshotSchedulingStatus,omitempty"`
}
