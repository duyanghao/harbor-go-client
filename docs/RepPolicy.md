# RepPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The policy ID. | [optional] [default to null]
**Name** | **string** | The policy name. | [optional] [default to null]
**Description** | **string** | The description of the policy. | [optional] [default to null]
**Projects** | [**[]Project**](Project.md) | The project list that the policy applys to. | [optional] [default to null]
**Targets** | [**[]RepTarget**](RepTarget.md) | The target list. | [optional] [default to null]
**Trigger** | [***RepTrigger**](RepTrigger.md) |  | [optional] [default to null]
**Filters** | [**[]RepFilter**](RepFilter.md) | The replication policy filter array. | [optional] [default to null]
**ReplicateExistingImageNow** | **bool** | Whether to replicate the existing images now. | [optional] [default to null]
**ReplicateDeletion** | **bool** | Whether to replicate the deletion operation. | [optional] [default to null]
**CreationTime** | **string** | The create time of the policy. | [optional] [default to null]
**UpdateTime** | **string** | The update time of the policy. | [optional] [default to null]
**ErrorJobCount** | **int32** | The error job count number for the policy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


