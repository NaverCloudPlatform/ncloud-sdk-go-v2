package vmongodb

type GetCloudMongoDbUserListRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// CloudMongoDb인스턴스번호
	CloudMongoDbInstanceNo *string `json:"cloudMongoDbInstanceNo,omitempty"`

	// CloudMongoDb User 리스트
	CloudMongoDbUserList []*CloudMongoDbUser `json:"cloudMongoDbUserList,omitempty"`
}
