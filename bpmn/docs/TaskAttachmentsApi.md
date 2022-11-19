# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttachment**](TaskAttachmentsApi.md#CreateAttachment) | **Post** /runtime/tasks/{taskId}/attachments | Create a new attachment on a task, containing a link to an external resource or an attached file
[**DeleteAttachment**](TaskAttachmentsApi.md#DeleteAttachment) | **Delete** /runtime/tasks/{taskId}/attachments/{attachmentId} | Delete an attachment on a task
[**GetAttachment**](TaskAttachmentsApi.md#GetAttachment) | **Get** /runtime/tasks/{taskId}/attachments/{attachmentId} | Get an attachment on a task
[**GetAttachmentContent**](TaskAttachmentsApi.md#GetAttachmentContent) | **Get** /runtime/tasks/{taskId}/attachments/{attachmentId}/content | Get the content for an attachment
[**ListTaskAttachments**](TaskAttachmentsApi.md#ListTaskAttachments) | **Get** /runtime/tasks/{taskId}/attachments | List attachments on a task

# **CreateAttachment**
> AttachmentResponse CreateAttachment(ctx, taskId, optional)
Create a new attachment on a task, containing a link to an external resource or an attached file

This endpoint can be used in 2 ways: By passing a JSON Body (AttachmentRequest) to link an external resource or by passing a multipart/form-data Object to attach a file. NB: Swagger V2 specification does not support this use case that is why this endpoint might be buggy/incomplete if used with other tools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
 **optional** | ***TaskAttachmentsApiCreateAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskAttachmentsApiCreateAttachmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TaskIdAttachmentsBody**](TaskIdAttachmentsBody.md)|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **name** | **optional.**|  | 
 **description** | **optional.**|  | 
 **type_** | **optional.**|  | 

### Return type

[**AttachmentResponse**](AttachmentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAttachment**
> DeleteAttachment(ctx, taskId, attachmentId)
Delete an attachment on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **attachmentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAttachment**
> AttachmentResponse GetAttachment(ctx, taskId, attachmentId)
Get an attachment on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **attachmentId** | **string**|  | 

### Return type

[**AttachmentResponse**](AttachmentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAttachmentContent**
> []string GetAttachmentContent(ctx, taskId, attachmentId)
Get the content for an attachment

The response body contains the binary content. By default, the content-type of the response is set to application/octet-stream unless the attachment type contains a valid Content-type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **attachmentId** | **string**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTaskAttachments**
> []AttachmentResponse ListTaskAttachments(ctx, taskId)
List attachments on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

[**[]AttachmentResponse**](AttachmentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

