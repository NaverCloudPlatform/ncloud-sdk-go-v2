/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type CloudMysqlRecoveryTime struct {

	// 시점복원가능한시작시간
RecoveryStartTime *string `json:"recoveryStartTime,omitempty"`

	// 시점복원가능한마지막시간
RecoveryEndTime *string `json:"recoveryEndTime,omitempty"`
}
