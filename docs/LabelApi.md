# \LabelApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChartreporepoChartsnameversionLabelsGet**](LabelApi.md#ChartreporepoChartsnameversionLabelsGet) | **Get** /chartrepo/:repo/charts/:name/:version/labels | Return the attahced labels of chart.
[**ChartreporepoChartsnameversionLabelsPost**](LabelApi.md#ChartreporepoChartsnameversionLabelsPost) | **Post** /chartrepo/:repo/charts/:name/:version/labels | Mark label to chart.
[**ChartreporepoChartsnameversionLabelsidDelete**](LabelApi.md#ChartreporepoChartsnameversionLabelsidDelete) | **Delete** /chartrepo/:repo/charts/:name/:version/labels/:id | Remove label from chart.


# **ChartreporepoChartsnameversionLabelsGet**
> ChartreporepoChartsnameversionLabelsGet(ctx, repo, name, version)
Return the attahced labels of chart.

Return the attahced labels of the specified chart version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartreporepoChartsnameversionLabelsPost**
> ChartreporepoChartsnameversionLabelsPost(ctx, repo, name, version, label)
Mark label to chart.

Mark label to the specified chart version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 
  **label** | [**Label**](Label.md)| The label being marked to the chart version | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartreporepoChartsnameversionLabelsidDelete**
> ChartreporepoChartsnameversionLabelsidDelete(ctx, repo, name, version, id)
Remove label from chart.

Remove label from the specified chart version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 
  **id** | **int32**| The label ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

