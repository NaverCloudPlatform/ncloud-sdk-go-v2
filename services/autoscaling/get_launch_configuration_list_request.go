/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type GetLaunchConfigurationListRequest struct {

	// 론치설정명리스트
LaunchConfigurationNameList []*string `json:"launchConfigurationNameList,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 소팅대상
SortedBy *string `json:"sortedBy,omitempty"`

	// 소팅순서
SortingOrder *string `json:"sortingOrder,omitempty"`

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`
}
