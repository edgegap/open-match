/*
 * Codema
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.41.174.post.dev2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Monitor struct {
	// API Name
	Name string `json:"name"`
	// API Version
	Version string `json:"version"`
	// API Host
	Host string `json:"host"`
	// API Host URL
	HostUrl string `json:"host_url"`
	// API Swagger Specification Location
	SpecUrl string `json:"spec_url"`
	// API Messages
	Messages []string `json:"messages,omitempty"`
	// API Errors
	Errors []string `json:"errors,omitempty"`
}
