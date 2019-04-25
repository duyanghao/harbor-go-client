/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VulnNamespaceTimestamp struct {
	// The namespace of the Vulnerability
	Namespace string `json:"namespace,omitempty"`
	// The UTC timestamp in miliseconds of last successful update for vulnerability data.
	LastUpdate int32 `json:"last_update,omitempty"`
}
