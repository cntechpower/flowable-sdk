# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskInstanceIdentityLinks**](TaskIdentityLinksApi.md#CreateTaskInstanceIdentityLinks) | **Post** /runtime/tasks/{taskId}/identitylinks | Create an identity link on a task
[**DeleteTaskInstanceIdentityLinks**](TaskIdentityLinksApi.md#DeleteTaskInstanceIdentityLinks) | **Delete** /runtime/tasks/{taskId}/identitylinks/{family}/{identityId}/{type} | Delete an identity link on a task
[**GetTaskInstanceIdentityLinks**](TaskIdentityLinksApi.md#GetTaskInstanceIdentityLinks) | **Get** /runtime/tasks/{taskId}/identitylinks/{family}/{identityId}/{type} | Get a single identity link on a task
[**ListIdentityLinksForFamily**](TaskIdentityLinksApi.md#ListIdentityLinksForFamily) | **Get** /runtime/tasks/{taskId}/identitylinks/{family} | List identity links for a task for either groups or users
[**ListTasksInstanceIdentityLinks**](TaskIdentityLinksApi.md#ListTasksInstanceIdentityLinks) | **Get** /runtime/tasks/{taskId}/identitylinks | List identity links for a task

# **CreateTaskInstanceIdentityLinks**
> RestIdentityLink CreateTaskInstanceIdentityLinks(ctx, taskId, optional)
Create an identity link on a task

It is possible to add either a user or a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
 **optional** | ***TaskIdentityLinksApiCreateTaskInstanceIdentityLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskIdentityLinksApiCreateTaskInstanceIdentityLinksOpts struct
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

# **DeleteTaskInstanceIdentityLinks**
> DeleteTaskInstanceIdentityLinks(ctx, taskId, family, identityId, type_)
Delete an identity link on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **family** | **string**|  | 
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

# **GetTaskInstanceIdentityLinks**
> RestIdentityLink GetTaskInstanceIdentityLinks(ctx, taskId, family, identityId, type_)
Get a single identity link on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **family** | **string**|  | 
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

# **ListIdentityLinksForFamily**
> []RestIdentityLink ListIdentityLinksForFamily(ctx, taskId, family)
List identity links for a task for either groups or users

Returns only identity links targeting either users or groups. Response body and status-codes are exactly the same as when getting the full list of identity links for a task.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **family** | **string**|  | 

### Return type

[**[]RestIdentityLink**](RestIdentityLink.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTasksInstanceIdentityLinks**
> []RestIdentityLink ListTasksInstanceIdentityLinks(ctx, taskId)
List identity links for a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

[**[]RestIdentityLink**](RestIdentityLink.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

