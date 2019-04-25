/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The chart search result item
type SearchResult struct {
	// The chart name with repo name
	Name string `json:"name,omitempty"`
	// The matched level
	Score int32 `json:"score,omitempty"`
	Chart *ChartVersion `json:"chart,omitempty"`
}
