/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetServerImageListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 서버이미지이름
ServerImageName *string `json:"serverImageName,omitempty"`

	// 서버이미지상태코드
ServerImageStatusCode *string `json:"serverImageStatusCode,omitempty"`

	// 서버이미지번호리스트
ServerImageNoList []*string `json:"serverImageNoList,omitempty"`

	// 하이퍼바이저유형코드리스트
HypervisorCodeList []*string `json:"hypervisorCodeList,omitempty"`

	// 서버이미지유형코드리스트
ServerImageTypeCodeList []*string `json:"serverImageTypeCodeList,omitempty"`

	// OS유형코드리스트
OsTypeCodeList []*string `json:"osTypeCodeList,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 정렬대상
SortedBy *string `json:"sortedBy,omitempty"`

	// 정렬순서
SortingOrder *string `json:"sortingOrder,omitempty"`
}