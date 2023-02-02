package vmongodb

type CloudMongoDbUserParam struct {
	// User 이름
	UserName *string `json:"userName,omitempty"`
	// Database 이름
	DatabaseName *string `json:"databaseName,omitempty"`
	// DB 권한
	Authority *string `json:"authority,omitempty"`
	// User Password
	Password *string `json:"password,omitempty"`
}
