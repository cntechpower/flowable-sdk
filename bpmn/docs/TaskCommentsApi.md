# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskComments**](TaskCommentsApi.md#CreateTaskComments) | **Post** /runtime/tasks/{taskId}/comments | Create a new comment on a task
[**DeleteTaskComment**](TaskCommentsApi.md#DeleteTaskComment) | **Delete** /runtime/tasks/{taskId}/comments/{commentId} | Delete a comment on a task
[**GetTaskComment**](TaskCommentsApi.md#GetTaskComment) | **Get** /runtime/tasks/{taskId}/comments/{commentId} |  Get a comment on a task
[**ListTaskComments**](TaskCommentsApi.md#ListTaskComments) | **Get** /runtime/tasks/{taskId}/comments | List comments on a task

# **CreateTaskComments**
> CommentResponse CreateTaskComments(ctx, taskId, optional)
Create a new comment on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
 **optional** | ***TaskCommentsApiCreateTaskCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskCommentsApiCreateTaskCommentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CommentRequest**](CommentRequest.md)|  | 

### Return type

[**CommentResponse**](CommentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTaskComment**
> DeleteTaskComment(ctx, taskId, commentId)
Delete a comment on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **commentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskComment**
> CommentResponse GetTaskComment(ctx, taskId, commentId)
 Get a comment on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **commentId** | **string**|  | 

### Return type

[**CommentResponse**](CommentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTaskComments**
> []CommentResponse ListTaskComments(ctx, taskId)
List comments on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

[**[]CommentResponse**](CommentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

