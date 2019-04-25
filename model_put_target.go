/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PutTarget struct {
	// The target name.
	Name string `json:"name,omitempty"`
	// The target address URL string.
	Endpoint string `json:"endpoint,omitempty"`
	// The target server username.
	Username string `json:"username,omitempty"`
	// The target server password.
	Password string `json:"password,omitempty"`
	// Whether or not the certificate will be verified when Harbor tries to access the server.
	Insecure bool `json:"insecure,omitempty"`
}
