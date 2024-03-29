/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type MetricInfoResponseVo struct {
	Aggregation string `json:"aggregation,omitempty"`
	Dimensions *DimensionsVo `json:"dimensions,omitempty"`
	Dps [][]interface{} `json:"dps,omitempty"`
	GraphName string `json:"graphName,omitempty"`
	Interval string `json:"interval,omitempty"`
	Metric string `json:"metric,omitempty"`
	ProductName string `json:"productName,omitempty"`
}
