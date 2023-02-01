package vmongodb

type AddCloudMongoDbUserListRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// CloudMongoDb 인스턴스번호
	CloudMongoDbInstanceNo *string `json:"cloudMongoDbInstanceNo,omitempty"`

	CloudMongoDbUserList []*CloudMongoDbUserParam `json:"cloudMongoDbUserList,omitempty"`
}
