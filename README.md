# Go API client for swagger

These APIs provide services for manipulating Harbor project.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.7.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import sw "./harbor-go-client"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ChartRepositoryApi* | [**ChartrepoChartsPost**](docs/ChartRepositoryApi.md#chartrepochartspost) | **Post** /chartrepo/charts | Upload a chart file to the defult &#39;library&#39; project.
*ChartRepositoryApi* | [**ChartrepoHealthGet**](docs/ChartRepositoryApi.md#chartrepohealthget) | **Get** /chartrepo/health | Check the health of chart repository service.
*ChartRepositoryApi* | [**ChartrepoRepoChartsGet**](docs/ChartRepositoryApi.md#chartreporepochartsget) | **Get** /chartrepo/{repo}/charts | Get all the charts under the specified project
*ChartRepositoryApi* | [**ChartrepoRepoChartsNameDelete**](docs/ChartRepositoryApi.md#chartreporepochartsnamedelete) | **Delete** /chartrepo/{repo}/charts/{name} | Delete all the versions of the specified chart
*ChartRepositoryApi* | [**ChartrepoRepoChartsNameGet**](docs/ChartRepositoryApi.md#chartreporepochartsnameget) | **Get** /chartrepo/{repo}/charts/{name} | Get all the versions of the specified chart
*ChartRepositoryApi* | [**ChartrepoRepoChartsNameVersionDelete**](docs/ChartRepositoryApi.md#chartreporepochartsnameversiondelete) | **Delete** /chartrepo/{repo}/charts/{name}/{version} | Delete the specified chart version
*ChartRepositoryApi* | [**ChartrepoRepoChartsNameVersionGet**](docs/ChartRepositoryApi.md#chartreporepochartsnameversionget) | **Get** /chartrepo/{repo}/charts/{name}/{version} | Get the specified chart version
*ChartRepositoryApi* | [**ChartrepoRepoChartsPost**](docs/ChartRepositoryApi.md#chartreporepochartspost) | **Post** /chartrepo/{repo}/charts | Upload a chart file to the specified project.
*ChartRepositoryApi* | [**ChartrepoRepoProvPost**](docs/ChartRepositoryApi.md#chartreporepoprovpost) | **Post** /chartrepo/{repo}/prov | Upload a provance file to the specified project.
*ChartRepositoryApi* | [**ChartreporepoChartsnameversionLabelsGet**](docs/ChartRepositoryApi.md#chartreporepochartsnameversionlabelsget) | **Get** /chartrepo/:repo/charts/:name/:version/labels | Return the attahced labels of chart.
*ChartRepositoryApi* | [**ChartreporepoChartsnameversionLabelsPost**](docs/ChartRepositoryApi.md#chartreporepochartsnameversionlabelspost) | **Post** /chartrepo/:repo/charts/:name/:version/labels | Mark label to chart.
*ChartRepositoryApi* | [**ChartreporepoChartsnameversionLabelsidDelete**](docs/ChartRepositoryApi.md#chartreporepochartsnameversionlabelsiddelete) | **Delete** /chartrepo/:repo/charts/:name/:version/labels/:id | Remove label from chart.
*LabelApi* | [**ChartreporepoChartsnameversionLabelsGet**](docs/LabelApi.md#chartreporepochartsnameversionlabelsget) | **Get** /chartrepo/:repo/charts/:name/:version/labels | Return the attahced labels of chart.
*LabelApi* | [**ChartreporepoChartsnameversionLabelsPost**](docs/LabelApi.md#chartreporepochartsnameversionlabelspost) | **Post** /chartrepo/:repo/charts/:name/:version/labels | Mark label to chart.
*LabelApi* | [**ChartreporepoChartsnameversionLabelsidDelete**](docs/LabelApi.md#chartreporepochartsnameversionlabelsiddelete) | **Delete** /chartrepo/:repo/charts/:name/:version/labels/:id | Remove label from chart.
*ProductsApi* | [**ChartrepoChartsPost**](docs/ProductsApi.md#chartrepochartspost) | **Post** /chartrepo/charts | Upload a chart file to the defult &#39;library&#39; project.
*ProductsApi* | [**ChartrepoHealthGet**](docs/ProductsApi.md#chartrepohealthget) | **Get** /chartrepo/health | Check the health of chart repository service.
*ProductsApi* | [**ChartrepoRepoChartsGet**](docs/ProductsApi.md#chartreporepochartsget) | **Get** /chartrepo/{repo}/charts | Get all the charts under the specified project
*ProductsApi* | [**ChartrepoRepoChartsNameDelete**](docs/ProductsApi.md#chartreporepochartsnamedelete) | **Delete** /chartrepo/{repo}/charts/{name} | Delete all the versions of the specified chart
*ProductsApi* | [**ChartrepoRepoChartsNameGet**](docs/ProductsApi.md#chartreporepochartsnameget) | **Get** /chartrepo/{repo}/charts/{name} | Get all the versions of the specified chart
*ProductsApi* | [**ChartrepoRepoChartsNameVersionDelete**](docs/ProductsApi.md#chartreporepochartsnameversiondelete) | **Delete** /chartrepo/{repo}/charts/{name}/{version} | Delete the specified chart version
*ProductsApi* | [**ChartrepoRepoChartsNameVersionGet**](docs/ProductsApi.md#chartreporepochartsnameversionget) | **Get** /chartrepo/{repo}/charts/{name}/{version} | Get the specified chart version
*ProductsApi* | [**ChartrepoRepoChartsPost**](docs/ProductsApi.md#chartreporepochartspost) | **Post** /chartrepo/{repo}/charts | Upload a chart file to the specified project.
*ProductsApi* | [**ChartrepoRepoProvPost**](docs/ProductsApi.md#chartreporepoprovpost) | **Post** /chartrepo/{repo}/prov | Upload a provance file to the specified project.
*ProductsApi* | [**ChartreporepoChartsnameversionLabelsGet**](docs/ProductsApi.md#chartreporepochartsnameversionlabelsget) | **Get** /chartrepo/:repo/charts/:name/:version/labels | Return the attahced labels of chart.
*ProductsApi* | [**ChartreporepoChartsnameversionLabelsPost**](docs/ProductsApi.md#chartreporepochartsnameversionlabelspost) | **Post** /chartrepo/:repo/charts/:name/:version/labels | Mark label to chart.
*ProductsApi* | [**ChartreporepoChartsnameversionLabelsidDelete**](docs/ProductsApi.md#chartreporepochartsnameversionlabelsiddelete) | **Delete** /chartrepo/:repo/charts/:name/:version/labels/:id | Remove label from chart.
*ProductsApi* | [**ConfigurationsGet**](docs/ProductsApi.md#configurationsget) | **Get** /configurations | Get system configurations.
*ProductsApi* | [**ConfigurationsPut**](docs/ProductsApi.md#configurationsput) | **Put** /configurations | Modify system configurations.
*ProductsApi* | [**ConfigurationsResetPost**](docs/ProductsApi.md#configurationsresetpost) | **Post** /configurations/reset | Reset system configurations.
*ProductsApi* | [**EmailPingPost**](docs/ProductsApi.md#emailpingpost) | **Post** /email/ping | Test connection and authentication with email server.
*ProductsApi* | [**InternalSyncregistryPost**](docs/ProductsApi.md#internalsyncregistrypost) | **Post** /internal/syncregistry | Sync repositories from registry to DB.
*ProductsApi* | [**JobsReplicationGet**](docs/ProductsApi.md#jobsreplicationget) | **Get** /jobs/replication | List filters jobs according to the policy and repository
*ProductsApi* | [**JobsReplicationIdDelete**](docs/ProductsApi.md#jobsreplicationiddelete) | **Delete** /jobs/replication/{id} | Delete specific ID job.
*ProductsApi* | [**JobsReplicationIdLogGet**](docs/ProductsApi.md#jobsreplicationidlogget) | **Get** /jobs/replication/{id}/log | Get job logs.
*ProductsApi* | [**JobsReplicationPut**](docs/ProductsApi.md#jobsreplicationput) | **Put** /jobs/replication | Update status of jobs. Only stop is supported for now.
*ProductsApi* | [**JobsScanIdLogGet**](docs/ProductsApi.md#jobsscanidlogget) | **Get** /jobs/scan/{id}/log | Get job logs.
*ProductsApi* | [**LabelsGet**](docs/ProductsApi.md#labelsget) | **Get** /labels | List labels according to the query strings.
*ProductsApi* | [**LabelsIdDelete**](docs/ProductsApi.md#labelsiddelete) | **Delete** /labels/{id} | Delete the label specified by ID.
*ProductsApi* | [**LabelsIdGet**](docs/ProductsApi.md#labelsidget) | **Get** /labels/{id} | Get the label specified by ID.
*ProductsApi* | [**LabelsIdPut**](docs/ProductsApi.md#labelsidput) | **Put** /labels/{id} | Update the label properties.
*ProductsApi* | [**LabelsIdResourcesGet**](docs/ProductsApi.md#labelsidresourcesget) | **Get** /labels/{id}/resources | Get the resources that the label is referenced by.
*ProductsApi* | [**LabelsPost**](docs/ProductsApi.md#labelspost) | **Post** /labels | Post creates a label
*ProductsApi* | [**LdapGroupsSearchGet**](docs/ProductsApi.md#ldapgroupssearchget) | **Get** /ldap/groups/search | Search available ldap groups.
*ProductsApi* | [**LdapPingPost**](docs/ProductsApi.md#ldappingpost) | **Post** /ldap/ping | Ping available ldap service.
*ProductsApi* | [**LdapUsersImportPost**](docs/ProductsApi.md#ldapusersimportpost) | **Post** /ldap/users/import | Import selected available ldap users.
*ProductsApi* | [**LdapUsersSearchGet**](docs/ProductsApi.md#ldapuserssearchget) | **Get** /ldap/users/search | Search available ldap users.
*ProductsApi* | [**LogsGet**](docs/ProductsApi.md#logsget) | **Get** /logs | Get recent logs of the projects which the user is a member of
*ProductsApi* | [**PoliciesReplicationGet**](docs/ProductsApi.md#policiesreplicationget) | **Get** /policies/replication | List filters policies by name and project_id
*ProductsApi* | [**PoliciesReplicationIdDelete**](docs/ProductsApi.md#policiesreplicationiddelete) | **Delete** /policies/replication/{id} | Delete the replication policy specified by ID.
*ProductsApi* | [**PoliciesReplicationIdGet**](docs/ProductsApi.md#policiesreplicationidget) | **Get** /policies/replication/{id} | Get replication policy.
*ProductsApi* | [**PoliciesReplicationIdPut**](docs/ProductsApi.md#policiesreplicationidput) | **Put** /policies/replication/{id} | Put modifies name, description, target and enablement of policy.
*ProductsApi* | [**PoliciesReplicationPost**](docs/ProductsApi.md#policiesreplicationpost) | **Post** /policies/replication | Post creates a policy
*ProductsApi* | [**ProjectsGet**](docs/ProductsApi.md#projectsget) | **Get** /projects | List projects
*ProductsApi* | [**ProjectsHead**](docs/ProductsApi.md#projectshead) | **Head** /projects | Check if the project name user provided already exists.
*ProductsApi* | [**ProjectsPost**](docs/ProductsApi.md#projectspost) | **Post** /projects | Create a new project.
*ProductsApi* | [**ProjectsProjectIdDelete**](docs/ProductsApi.md#projectsprojectiddelete) | **Delete** /projects/{project_id} | Delete project by projectID
*ProductsApi* | [**ProjectsProjectIdGet**](docs/ProductsApi.md#projectsprojectidget) | **Get** /projects/{project_id} | Return specific project detail infomation
*ProductsApi* | [**ProjectsProjectIdLogsGet**](docs/ProductsApi.md#projectsprojectidlogsget) | **Get** /projects/{project_id}/logs | Get access logs accompany with a relevant project.
*ProductsApi* | [**ProjectsProjectIdMembersGet**](docs/ProductsApi.md#projectsprojectidmembersget) | **Get** /projects/{project_id}/members | Get all project member information
*ProductsApi* | [**ProjectsProjectIdMembersMidDelete**](docs/ProductsApi.md#projectsprojectidmembersmiddelete) | **Delete** /projects/{project_id}/members/{mid} | Delete project member
*ProductsApi* | [**ProjectsProjectIdMembersMidGet**](docs/ProductsApi.md#projectsprojectidmembersmidget) | **Get** /projects/{project_id}/members/{mid} | Get the project member information
*ProductsApi* | [**ProjectsProjectIdMembersMidPut**](docs/ProductsApi.md#projectsprojectidmembersmidput) | **Put** /projects/{project_id}/members/{mid} | Update project member
*ProductsApi* | [**ProjectsProjectIdMembersPost**](docs/ProductsApi.md#projectsprojectidmemberspost) | **Post** /projects/{project_id}/members | Create project member
*ProductsApi* | [**ProjectsProjectIdMetadatasGet**](docs/ProductsApi.md#projectsprojectidmetadatasget) | **Get** /projects/{project_id}/metadatas | Get project metadata.
*ProductsApi* | [**ProjectsProjectIdMetadatasMetaNameDelete**](docs/ProductsApi.md#projectsprojectidmetadatasmetanamedelete) | **Delete** /projects/{project_id}/metadatas/{meta_name} | Delete metadata of a project
*ProductsApi* | [**ProjectsProjectIdMetadatasMetaNameGet**](docs/ProductsApi.md#projectsprojectidmetadatasmetanameget) | **Get** /projects/{project_id}/metadatas/{meta_name} | Get project metadata
*ProductsApi* | [**ProjectsProjectIdMetadatasMetaNamePut**](docs/ProductsApi.md#projectsprojectidmetadatasmetanameput) | **Put** /projects/{project_id}/metadatas/{meta_name} | Update metadata of a project.
*ProductsApi* | [**ProjectsProjectIdMetadatasPost**](docs/ProductsApi.md#projectsprojectidmetadataspost) | **Post** /projects/{project_id}/metadatas | Add metadata for the project.
*ProductsApi* | [**ProjectsProjectIdPut**](docs/ProductsApi.md#projectsprojectidput) | **Put** /projects/{project_id} | Update properties for a selected project.
*ProductsApi* | [**ReplicationsPost**](docs/ProductsApi.md#replicationspost) | **Post** /replications | Trigger the replication according to the specified policy.
*ProductsApi* | [**RepositoriesGet**](docs/ProductsApi.md#repositoriesget) | **Get** /repositories | Get repositories accompany with relevant project and repo name.
*ProductsApi* | [**RepositoriesRepoNameDelete**](docs/ProductsApi.md#repositoriesreponamedelete) | **Delete** /repositories/{repo_name} | Delete a repository.
*ProductsApi* | [**RepositoriesRepoNameLabelsGet**](docs/ProductsApi.md#repositoriesreponamelabelsget) | **Get** /repositories/{repo_name}/labels | Get labels of a repository.
*ProductsApi* | [**RepositoriesRepoNameLabelsLabelIdDelete**](docs/ProductsApi.md#repositoriesreponamelabelslabeliddelete) | **Delete** /repositories/{repo_name}/labels/{label_id} | Delete label from the repository.
*ProductsApi* | [**RepositoriesRepoNameLabelsPost**](docs/ProductsApi.md#repositoriesreponamelabelspost) | **Post** /repositories/{repo_name}/labels | Add a label to the repository.
*ProductsApi* | [**RepositoriesRepoNamePut**](docs/ProductsApi.md#repositoriesreponameput) | **Put** /repositories/{repo_name} | Update description of the repository.
*ProductsApi* | [**RepositoriesRepoNameSignaturesGet**](docs/ProductsApi.md#repositoriesreponamesignaturesget) | **Get** /repositories/{repo_name}/signatures | Get signature information of a repository
*ProductsApi* | [**RepositoriesRepoNameTagsGet**](docs/ProductsApi.md#repositoriesreponametagsget) | **Get** /repositories/{repo_name}/tags | Get tags of a relevant repository.
*ProductsApi* | [**RepositoriesRepoNameTagsPost**](docs/ProductsApi.md#repositoriesreponametagspost) | **Post** /repositories/{repo_name}/tags | Retag an image
*ProductsApi* | [**RepositoriesRepoNameTagsTagDelete**](docs/ProductsApi.md#repositoriesreponametagstagdelete) | **Delete** /repositories/{repo_name}/tags/{tag} | Delete a tag in a repository.
*ProductsApi* | [**RepositoriesRepoNameTagsTagGet**](docs/ProductsApi.md#repositoriesreponametagstagget) | **Get** /repositories/{repo_name}/tags/{tag} | Get the tag of the repository.
*ProductsApi* | [**RepositoriesRepoNameTagsTagLabelsGet**](docs/ProductsApi.md#repositoriesreponametagstaglabelsget) | **Get** /repositories/{repo_name}/tags/{tag}/labels | Get labels of an image.
*ProductsApi* | [**RepositoriesRepoNameTagsTagLabelsLabelIdDelete**](docs/ProductsApi.md#repositoriesreponametagstaglabelslabeliddelete) | **Delete** /repositories/{repo_name}/tags/{tag}/labels/{label_id} | Delete label from the image.
*ProductsApi* | [**RepositoriesRepoNameTagsTagLabelsPost**](docs/ProductsApi.md#repositoriesreponametagstaglabelspost) | **Post** /repositories/{repo_name}/tags/{tag}/labels | Add a label to image.
*ProductsApi* | [**RepositoriesRepoNameTagsTagManifestGet**](docs/ProductsApi.md#repositoriesreponametagstagmanifestget) | **Get** /repositories/{repo_name}/tags/{tag}/manifest | Get manifests of a relevant repository.
*ProductsApi* | [**RepositoriesRepoNameTagsTagScanPost**](docs/ProductsApi.md#repositoriesreponametagstagscanpost) | **Post** /repositories/{repo_name}/tags/{tag}/scan | Scan the image.
*ProductsApi* | [**RepositoriesRepoNameTagsTagVulnerabilityDetailsGet**](docs/ProductsApi.md#repositoriesreponametagstagvulnerabilitydetailsget) | **Get** /repositories/{repo_name}/tags/{tag}/vulnerability/details | Get vulnerability details of the image.
*ProductsApi* | [**RepositoriesScanAllPost**](docs/ProductsApi.md#repositoriesscanallpost) | **Post** /repositories/scanAll | Scan all images of the registry.
*ProductsApi* | [**RepositoriesTopGet**](docs/ProductsApi.md#repositoriestopget) | **Get** /repositories/top | Get public repositories which are accessed most.
*ProductsApi* | [**SearchGet**](docs/ProductsApi.md#searchget) | **Get** /search | Search for projects, repositories and helm charts
*ProductsApi* | [**StatisticsGet**](docs/ProductsApi.md#statisticsget) | **Get** /statistics | Get projects number and repositories number relevant to the user
*ProductsApi* | [**SystemGcGet**](docs/ProductsApi.md#systemgcget) | **Get** /system/gc | Get gc results.
*ProductsApi* | [**SystemGcIdGet**](docs/ProductsApi.md#systemgcidget) | **Get** /system/gc/{id} | Get gc status.
*ProductsApi* | [**SystemGcIdLogGet**](docs/ProductsApi.md#systemgcidlogget) | **Get** /system/gc/{id}/log | Get gc job log.
*ProductsApi* | [**SystemGcScheduleGet**](docs/ProductsApi.md#systemgcscheduleget) | **Get** /system/gc/schedule | Get gc&#39;s schedule.
*ProductsApi* | [**SystemGcSchedulePost**](docs/ProductsApi.md#systemgcschedulepost) | **Post** /system/gc/schedule | Create a gc schedule.
*ProductsApi* | [**SystemGcSchedulePut**](docs/ProductsApi.md#systemgcscheduleput) | **Put** /system/gc/schedule | Update gc&#39;s schedule.
*ProductsApi* | [**SysteminfoGet**](docs/ProductsApi.md#systeminfoget) | **Get** /systeminfo | Get general system info
*ProductsApi* | [**SysteminfoGetcertGet**](docs/ProductsApi.md#systeminfogetcertget) | **Get** /systeminfo/getcert | Get default root certificate.
*ProductsApi* | [**SysteminfoVolumesGet**](docs/ProductsApi.md#systeminfovolumesget) | **Get** /systeminfo/volumes | Get system volume info (total/free size).
*ProductsApi* | [**TargetsGet**](docs/ProductsApi.md#targetsget) | **Get** /targets | List filters targets by name.
*ProductsApi* | [**TargetsIdDelete**](docs/ProductsApi.md#targetsiddelete) | **Delete** /targets/{id} | Delete specific replication&#39;s target.
*ProductsApi* | [**TargetsIdGet**](docs/ProductsApi.md#targetsidget) | **Get** /targets/{id} | Get replication&#39;s target.
*ProductsApi* | [**TargetsIdPoliciesGet**](docs/ProductsApi.md#targetsidpoliciesget) | **Get** /targets/{id}/policies/ | List the target relevant policies.
*ProductsApi* | [**TargetsIdPut**](docs/ProductsApi.md#targetsidput) | **Put** /targets/{id} | Update replication&#39;s target.
*ProductsApi* | [**TargetsPingPost**](docs/ProductsApi.md#targetspingpost) | **Post** /targets/ping | Ping validates target.
*ProductsApi* | [**TargetsPost**](docs/ProductsApi.md#targetspost) | **Post** /targets | Create a new replication target.
*ProductsApi* | [**UsergroupsGet**](docs/ProductsApi.md#usergroupsget) | **Get** /usergroups | Get all user groups information
*ProductsApi* | [**UsergroupsGroupIdDelete**](docs/ProductsApi.md#usergroupsgroupiddelete) | **Delete** /usergroups/{group_id} | Delete user group
*ProductsApi* | [**UsergroupsGroupIdGet**](docs/ProductsApi.md#usergroupsgroupidget) | **Get** /usergroups/{group_id} | Get user group information
*ProductsApi* | [**UsergroupsGroupIdPut**](docs/ProductsApi.md#usergroupsgroupidput) | **Put** /usergroups/{group_id} | Update group information
*ProductsApi* | [**UsergroupsPost**](docs/ProductsApi.md#usergroupspost) | **Post** /usergroups | Create user group
*ProductsApi* | [**UsersCurrentGet**](docs/ProductsApi.md#userscurrentget) | **Get** /users/current | Get current user info.
*ProductsApi* | [**UsersGet**](docs/ProductsApi.md#usersget) | **Get** /users | Get registered users of Harbor.
*ProductsApi* | [**UsersPost**](docs/ProductsApi.md#userspost) | **Post** /users | Creates a new user account.
*ProductsApi* | [**UsersUserIdDelete**](docs/ProductsApi.md#usersuseriddelete) | **Delete** /users/{user_id} | Mark a registered user as be removed.
*ProductsApi* | [**UsersUserIdGet**](docs/ProductsApi.md#usersuseridget) | **Get** /users/{user_id} | Get a user&#39;s profile.
*ProductsApi* | [**UsersUserIdPasswordPut**](docs/ProductsApi.md#usersuseridpasswordput) | **Put** /users/{user_id}/password | Change the password on a user that already exists.
*ProductsApi* | [**UsersUserIdPut**](docs/ProductsApi.md#usersuseridput) | **Put** /users/{user_id} | Update a registered user to change his profile.
*ProductsApi* | [**UsersUserIdSysadminPut**](docs/ProductsApi.md#usersuseridsysadminput) | **Put** /users/{user_id}/sysadmin | Update a registered user to change to be an administrator of Harbor.


## Documentation For Models

 - [AccessLog](docs/AccessLog.md)
 - [BadRequestFormatedError](docs/BadRequestFormatedError.md)
 - [BoolConfigItem](docs/BoolConfigItem.md)
 - [ChartApiError](docs/ChartApiError.md)
 - [ChartInfoEntry](docs/ChartInfoEntry.md)
 - [ChartInfoList](docs/ChartInfoList.md)
 - [ChartMetadata](docs/ChartMetadata.md)
 - [ChartVersion](docs/ChartVersion.md)
 - [ChartVersionDetails](docs/ChartVersionDetails.md)
 - [ChartVersions](docs/ChartVersions.md)
 - [ComponentOverviewEntry](docs/ComponentOverviewEntry.md)
 - [Configurations](docs/Configurations.md)
 - [ConfigurationsResponse](docs/ConfigurationsResponse.md)
 - [ConfigurationsScanAllPolicy](docs/ConfigurationsScanAllPolicy.md)
 - [ConfigurationsScanAllPolicyParameter](docs/ConfigurationsScanAllPolicyParameter.md)
 - [ConflictFormatedError](docs/ConflictFormatedError.md)
 - [Dependency](docs/Dependency.md)
 - [DetailedTag](docs/DetailedTag.md)
 - [DetailedTagScanOverview](docs/DetailedTagScanOverview.md)
 - [DetailedTagScanOverviewComponents](docs/DetailedTagScanOverviewComponents.md)
 - [DigitalSignature](docs/DigitalSignature.md)
 - [EmailServerSetting](docs/EmailServerSetting.md)
 - [ForbiddenChartApiError](docs/ForbiddenChartApiError.md)
 - [GcResult](docs/GcResult.md)
 - [GcSchedule](docs/GcSchedule.md)
 - [GcScheduleSchedule](docs/GcScheduleSchedule.md)
 - [GeneralInfo](docs/GeneralInfo.md)
 - [GeneralInfoClairVulnerabilityStatus](docs/GeneralInfoClairVulnerabilityStatus.md)
 - [HasAdminRole](docs/HasAdminRole.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InsufficientStorageChartApiError](docs/InsufficientStorageChartApiError.md)
 - [IntegerConfigItem](docs/IntegerConfigItem.md)
 - [InternalChartApiError](docs/InternalChartApiError.md)
 - [JobStatus](docs/JobStatus.md)
 - [Label](docs/Label.md)
 - [Labels](docs/Labels.md)
 - [LdapConf](docs/LdapConf.md)
 - [LdapFailedImportUsers](docs/LdapFailedImportUsers.md)
 - [LdapImportUsers](docs/LdapImportUsers.md)
 - [LdapUsers](docs/LdapUsers.md)
 - [Manifest](docs/Manifest.md)
 - [NotFoundChartApiError](docs/NotFoundChartApiError.md)
 - [Password](docs/Password.md)
 - [PingTarget](docs/PingTarget.md)
 - [Project](docs/Project.md)
 - [ProjectMember](docs/ProjectMember.md)
 - [ProjectMemberEntity](docs/ProjectMemberEntity.md)
 - [ProjectMetadata](docs/ProjectMetadata.md)
 - [ProjectReq](docs/ProjectReq.md)
 - [PutTarget](docs/PutTarget.md)
 - [RepFilter](docs/RepFilter.md)
 - [RepPolicy](docs/RepPolicy.md)
 - [RepTarget](docs/RepTarget.md)
 - [RepTargetPost](docs/RepTargetPost.md)
 - [RepTrigger](docs/RepTrigger.md)
 - [Replication](docs/Replication.md)
 - [ReplicationResponse](docs/ReplicationResponse.md)
 - [RepoSignature](docs/RepoSignature.md)
 - [Repository](docs/Repository.md)
 - [RepositoryDescription](docs/RepositoryDescription.md)
 - [Resource](docs/Resource.md)
 - [RetagReq](docs/RetagReq.md)
 - [Role](docs/Role.md)
 - [RoleParam](docs/RoleParam.md)
 - [RoleRequest](docs/RoleRequest.md)
 - [ScheduleParam](docs/ScheduleParam.md)
 - [Search](docs/Search.md)
 - [SearchRepository](docs/SearchRepository.md)
 - [SearchResult](docs/SearchResult.md)
 - [SecurityReport](docs/SecurityReport.md)
 - [StatisticMap](docs/StatisticMap.md)
 - [Storage](docs/Storage.md)
 - [StringConfigItem](docs/StringConfigItem.md)
 - [SystemInfo](docs/SystemInfo.md)
 - [Tags](docs/Tags.md)
 - [UnauthorizedChartApiError](docs/UnauthorizedChartApiError.md)
 - [UpdateJobs](docs/UpdateJobs.md)
 - [User](docs/User.md)
 - [UserEntity](docs/UserEntity.md)
 - [UserGroup](docs/UserGroup.md)
 - [UserProfile](docs/UserProfile.md)
 - [VulnNamespaceTimestamp](docs/VulnNamespaceTimestamp.md)
 - [VulnerabilityItem](docs/VulnerabilityItem.md)


## Usage

```golang
import sw "github.com/jeremyxu2010/harbor-go-client"

cfg := &sw.Configuration{
	BasePath:      "http://${harbor_svc_ip}/api",
	DefaultHeader: make(map[string]string),
	UserAgent:     "Swagger-Codegen/1.0.0/go",
}
client := sw.NewAPIClient(cfg)
// client.ProductsApi.SomeOperation(...)
// client.LabelApi.SomeOperation(...)
// client.ChartRepositoryApi.SomeOperation(...)
```

## Documentation For Authorization
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author

[jeremyxu2010](https://github.com/jeremyxu2010)

