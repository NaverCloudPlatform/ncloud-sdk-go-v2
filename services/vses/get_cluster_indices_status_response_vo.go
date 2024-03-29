/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type GetClusterIndicesStatusResponseVo struct {
	CurrentPage int32 `json:"currentPage,omitempty"`
	IndiceInfoList []CatIndice `json:"indiceInfoList,omitempty"`
	IsPaged bool `json:"isPaged,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
	TotalCount int32 `json:"totalCount,omitempty"`
	TotalPage int32 `json:"totalPage,omitempty"`
}
