/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type SetSnapshotScheduling struct {
	BucketName string `json:"bucketName"`
	SnapshotName string `json:"snapshotName"`
	ScheduledDay string `json:"scheduledDay"`
	ScheduledHour int32 `json:"scheduledHour"`
	ScheduledMinute int32 `json:"scheduledMinute"`
}
