/*
 * sourcepipeline
 *
 * <br/>https://sourcepipeline.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcepipeline

type GetHistoryListResponseHistoryList struct {
	ProjectId *int32 `json:"projectId,omitempty"`

	Id *int32 `json:"id,omitempty"`

	RequestType *string `json:"requestType,omitempty"`

	RequestId *string `json:"requestId,omitempty"`

	Begin *int32 `json:"begin,omitempty"`

	End *int32 `json:"end,omitempty"`

	Status *string `json:"status,omitempty"`
}