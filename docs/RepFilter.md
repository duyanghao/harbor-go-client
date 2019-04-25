# RepFilter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The replication policy filter kind. The valid values are project, repository and tag. | [optional] [default to null]
**Value** | **string** | The value of replication policy filter. When creating repository and tag filter, filling it with the pattern as string. When creating label filter, filling it with label ID as integer. | [optional] [default to null]
**Pattern** | **string** | Depraceted, use value instead. The replication policy filter pattern. | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | This map object is the replication policy filter metadata. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


