/*
 * Codema
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.41.174.post.dev2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Application struct {
	// Application name
	Name string `json:"name"`
	// If the application can be deployed
	IsActive bool `json:"is_active"`
	// Image base64 string
	Image string `json:"image,omitempty"`
	// Creation date
	CreateTime string `json:"create_time"`
	// Date of the last update
	LastUpdated string `json:"last_updated"`
}