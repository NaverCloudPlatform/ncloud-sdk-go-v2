/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetNetworkInterfaceListRequest struct {

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`

	// ZONE번호
ZoneNo *string `json:"zoneNo,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`
}
