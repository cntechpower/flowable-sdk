# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBatch**](BatchesApi.md#DeleteBatch) | **Delete** /management/batches/{batchId} | Delete a batch
[**GetBatch**](BatchesApi.md#GetBatch) | **Get** /management/batches/{batchId} | Get a single batch
[**GetBatchDocument**](BatchesApi.md#GetBatchDocument) | **Get** /management/batches/{batchId}/batch-document | Get the batch document
[**GetBatchPartDocument**](BatchesApi.md#GetBatchPartDocument) | **Get** /management/batch-parts/{batchPartId}/batch-part-document | Get the batch part document
[**ListBatches**](BatchesApi.md#ListBatches) | **Get** /management/batches | List batches
[**ListBatchesPart**](BatchesApi.md#ListBatchesPart) | **Get** /management/batches/{batchId}/batch-parts | List batch parts

# **DeleteBatch**
> DeleteBatch(ctx, batchId)
Delete a batch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBatch**
> BatchResponse GetBatch(ctx, batchId)
Get a single batch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchId** | **string**|  | 

### Return type

[**BatchResponse**](BatchResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBatchDocument**
> string GetBatchDocument(ctx, batchId)
Get the batch document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchId** | **string**|  | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBatchPartDocument**
> string GetBatchPartDocument(ctx, batchPartId)
Get the batch part document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchPartId** | **string**|  | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBatches**
> DataResponseBatchResponse ListBatches(ctx, optional)
List batches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchesApiListBatchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchesApiListBatchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return batch with the given id | 
 **batchType** | **optional.String**| Only return batches for the given type | 
 **searchKey** | **optional.String**| Only return batches for the given search key | 
 **searchKey2** | **optional.String**| Only return batches for the given search key2 | 
 **createTimeBefore** | **optional.Time**| Only return batches created before the given date | 
 **createTimeAfter** | **optional.Time**| Only batches batches created after the given date | 
 **completeTimeBefore** | **optional.Time**| Only return batches completed before the given date | 
 **completeTimeAfter** | **optional.Time**| Only batches batches completed after the given date | 
 **status** | **optional.String**| Only return batches for the given status | 
 **tenantId** | **optional.String**| Only return batches for the given tenant id | 
 **tenantIdLike** | **optional.String**| Only return batches like given search key | 
 **withoutTenantId** | **optional.Bool**| If true, only returns batches without a tenantId set. If false, the withoutTenantId parameter is ignored. | 

### Return type

[**DataResponseBatchResponse**](DataResponseBatchResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBatchesPart**
> []BatchPartResponse ListBatchesPart(ctx, batchId, optional)
List batch parts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchId** | **string**|  | 
 **optional** | ***BatchesApiListBatchesPartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchesApiListBatchesPartOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| Only return batch parts for the given status | 

### Return type

[**[]BatchPartResponse**](BatchPartResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

