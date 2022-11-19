# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBatchPart**](BatchPartsApi.md#GetBatchPart) | **Get** /management/batch-parts/{batchPartId} | Get a single batch part

# **GetBatchPart**
> BatchPartResponse GetBatchPart(ctx, batchPartId)
Get a single batch part

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchPartId** | **string**|  | 

### Return type

[**BatchPartResponse**](BatchPartResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

