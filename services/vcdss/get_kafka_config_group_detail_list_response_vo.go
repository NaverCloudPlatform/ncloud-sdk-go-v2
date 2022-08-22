/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type GetKafkaConfigGroupDetailListResponseVo struct {
	ConfigGroupNo              int32               `json:"configGroupNo,omitempty"`
	KafkaConfigGroupDetailList []KafkaConfigDetail `json:"kafkaConfigGroupDetailList,omitempty"`
	KafkaVersionCode           string              `json:"kafkaVersionCode,omitempty"`
	MemberNo                   string              `json:"memberNo,omitempty"`
}
