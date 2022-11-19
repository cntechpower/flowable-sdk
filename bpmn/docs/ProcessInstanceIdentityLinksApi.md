# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProcessInstanceIdentityLinks**](ProcessInstanceIdentityLinksApi.md#CreateProcessInstanceIdentityLinks) | **Post** /runtime/process-instances/{processInstanceId}/identitylinks | Add an involved user to a process instance
[**DeleteProcessInstanceIdentityLinks**](ProcessInstanceIdentityLinksApi.md#DeleteProcessInstanceIdentityLinks) | **Delete** /runtime/process-instances/{processInstanceId}/identitylinks/users/{identityId}/{type} | Remove an involved user to from process instance
[**GetProcessInstanceIdentityLinks**](ProcessInstanceIdentityLinksApi.md#GetProcessInstanceIdentityLinks) | **Get** /runtime/process-instances/{processInstanceId}/identitylinks/users/{identityId}/{type} | Get a specific involved people from process instance
[**ListProcessInstanceIdentityLinks**](ProcessInstanceIdentityLinksApi.md#ListProcessInstanceIdentityLinks) | **Get** /runtime/process-instances/{processInstanceId}/identitylinks | Get involved people for process instance

# **CreateProcessInstanceIdentityLinks**
> RestIdentityLink CreateProcessInstanceIdentityLinks(ctx, processInstanceId, optional)
Add an involved user to a process instance

Note that the groupId in Response Body will always be null, as it’s only possible to involve users with a process-instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
 **optional** | ***ProcessInstanceIdentityLinksApiCreateProcessInstanceIdentityLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstanceIdentityLinksApiCreateProcessInstanceIdentityLinksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RestIdentityLink**](RestIdentityLink.md)|  | 

### Return type

[**RestIdentityLink**](RestIdentityLink.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProcessInstanceIdentityLinks**
> DeleteProcessInstanceIdentityLinks(ctx, processInstanceId, identityId, type_)
Remove an involved user to from process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
  **identityId** | **string**|  | 
  **type_** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessInstanceIdentityLinks**
> RestIdentityLink GetProcessInstanceIdentityLinks(ctx, processInstanceId, identityId, type_)
Get a specific involved people from process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
  **identityId** | **string**|  | 
  **type_** | **string**|  | 

### Return type

[**RestIdentityLink**](RestIdentityLink.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProcessInstanceIdentityLinks**
> []RestIdentityLink ListProcessInstanceIdentityLinks(ctx, processInstanceId)
Get involved people for process instance

Note that the groupId in Response Body will always be null, as it’s only possible to involve users with a process-instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 

### Return type

[**[]RestIdentityLink**](RestIdentityLink.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

