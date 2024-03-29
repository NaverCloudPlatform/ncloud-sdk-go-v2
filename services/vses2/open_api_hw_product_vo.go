/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v2)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses2

type OpenApiHwProductVo struct {
	CpuCount              string `json:"cpuCount,omitempty"`
	DiskType2Code         string `json:"diskType2Code,omitempty"`
	InfraResourceTypeCode string `json:"infraResourceTypeCode,omitempty"`
	MemorySize            int64  `json:"memorySize,omitempty"`
	ProductCode           string `json:"productCode,omitempty"`
	ProductEnglishDesc    string `json:"productEnglishDesc,omitempty"`
	ProductName           string `json:"productName,omitempty"`
	ProductType2Code      string `json:"productType2Code,omitempty"`
}
