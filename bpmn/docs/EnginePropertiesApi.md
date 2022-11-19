# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEngineProperty**](EnginePropertiesApi.md#CreateEngineProperty) | **Post** /management/engine-properties | Create a new engine property
[**DeleteEngineProperty**](EnginePropertiesApi.md#DeleteEngineProperty) | **Delete** /management/engine-properties/{engineProperty} | Delete an engine property
[**GetEngineProperties**](EnginePropertiesApi.md#GetEngineProperties) | **Get** /management/engine-properties | Get all engine properties
[**UpdateEngineProperty**](EnginePropertiesApi.md#UpdateEngineProperty) | **Put** /management/engine-properties/{engineProperty} | Update an engine property

# **CreateEngineProperty**
> CreateEngineProperty(ctx, optional)
Create a new engine property

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnginePropertiesApiCreateEnginePropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnginePropertiesApiCreateEnginePropertyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PropertyRequestBody**](PropertyRequestBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEngineProperty**
> DeleteEngineProperty(ctx, engineProperty)
Delete an engine property

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineProperty** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEngineProperties**
> map[string]string GetEngineProperties(ctx, )
Get all engine properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

**map[string]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEngineProperty**
> UpdateEngineProperty(ctx, engineProperty, optional)
Update an engine property

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineProperty** | **string**|  | 
 **optional** | ***EnginePropertiesApiUpdateEnginePropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnginePropertiesApiUpdateEnginePropertyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PropertyRequestBody**](PropertyRequestBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

