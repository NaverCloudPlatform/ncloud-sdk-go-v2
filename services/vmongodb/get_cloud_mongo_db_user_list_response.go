package vmongodb

type GetCloudMongoDbUserListResponse struct {
	RequestId *string `json:"requestId,omitempty"`

	ReturnCode *string `json:"returnCode,omitempty"`

	ReturnMessage *string `json:"returnMessage,omitempty"`

	TotalRows *int32 `json:"totalRows,omitempty"`

	// CloudMongoDb User 리스트
	CloudMongoDbUserList []*CloudMongoDbUser `json:"cloudMongoDbUserList,omitempty"`
}
