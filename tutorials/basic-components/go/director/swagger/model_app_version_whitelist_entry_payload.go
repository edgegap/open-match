/*
 * Codema
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.41.174.post.dev2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AppVersionWhitelistEntryPayload struct {
	// CIDR to allow
	Cidr string `json:"cidr"`
	// Label to organized your entries
	Label string `json:"label,omitempty"`
	// If the Rule will be applied on runtime
	IsActive bool `json:"is_active,omitempty"`
}
