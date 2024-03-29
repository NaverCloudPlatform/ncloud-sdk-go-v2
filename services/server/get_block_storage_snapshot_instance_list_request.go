/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetBlockStorageSnapshotInstanceListRequest struct {

	// 블록스토리지스냅샷인스턴스번호리스트
BlockStorageSnapshotInstanceNoList []*string `json:"blockStorageSnapshotInstanceNoList,omitempty"`

	// 원본블록스토리지인스턴스번호리스트
OriginalBlockStorageInstanceNoList []*string `json:"originalBlockStorageInstanceNoList,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`
}
