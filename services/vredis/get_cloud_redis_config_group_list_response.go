/*
 * vredis
 *
 * <br/>https://ncloud.apigw.ntruss.com/vredis/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vredis

type GetCloudRedisConfigGroupListResponse struct {
	RequestId *string `json:"requestId,omitempty"`

	ReturnCode *string `json:"returnCode,omitempty"`

	ReturnMessage *string `json:"returnMessage,omitempty"`

	TotalRows *int32 `json:"totalRows,omitempty"`

	// CloudRedisConfigGroup리스트
	CloudRedisConfigGroupList []*CloudRedisConfigGroup `json:"cloudRedisConfigGroupList,omitempty"`
}
