/*
 * vnks
 *
 * <br/>https://nks.apigw.ntruss.com/vnks/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnks

type UpdateOidcDto struct {
	ClientId *string `json:"clientId"`

	IssuerURL *string `json:"issuerURL"`

	Status *bool `json:"status"`

	GroupsClaim *string `json:"groupsClaim,omitempty"`

	GroupsPrefix *string `json:"groupsPrefix,omitempty"`

	UsernameClaim *string `json:"usernameClaim,omitempty"`

	UsernamePrefix *string `json:"usernamePrefix,omitempty"`

	RequiredClaim *string `json:"requiredClaim,omitempty"`
}
