/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GcResult struct {
	// the id of gc job.
	Id int32 `json:"id,omitempty"`
	// the job name of gc job.
	JobName string `json:"job_name,omitempty"`
	// the job kind of gc job.
	JobKind string `json:"job_kind,omitempty"`
	Schedule *GcScheduleSchedule `json:"schedule,omitempty"`
	// the status of gc job.
	JobStatus string `json:"job_status,omitempty"`
	// if gc job was deleted.
	Deleted bool `json:"deleted,omitempty"`
	// the creation time of gc job.
	CreationTime string `json:"creation_time,omitempty"`
	// the update time of gc job.
	UpdateTime string `json:"update_time,omitempty"`
}
