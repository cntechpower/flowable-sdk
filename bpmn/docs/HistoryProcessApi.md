# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComment**](HistoryProcessApi.md#CreateComment) | **Post** /history/historic-process-instances/{processInstanceId}/comments | Create a new comment on a historic process instance
[**DeleteComment**](HistoryProcessApi.md#DeleteComment) | **Delete** /history/historic-process-instances/{processInstanceId}/comments/{commentId} | Delete a comment on a historic process instance
[**DeleteHistoricProcessInstance**](HistoryProcessApi.md#DeleteHistoricProcessInstance) | **Delete** /history/historic-process-instances/{processInstanceId} |  Delete a historic process instance
[**GetComment**](HistoryProcessApi.md#GetComment) | **Get** /history/historic-process-instances/{processInstanceId}/comments/{commentId} | Get a comment on a historic process instance
[**GetHistoricProcessInstance**](HistoryProcessApi.md#GetHistoricProcessInstance) | **Get** /history/historic-process-instances/{processInstanceId} | Get a historic process instance
[**GetHistoricProcessInstanceVariableData**](HistoryProcessApi.md#GetHistoricProcessInstanceVariableData) | **Get** /history/historic-process-instances/{processInstanceId}/variables/{variableName}/data | Get the binary data for a historic process instance variable
[**ListHistoricProcessInstanceComments**](HistoryProcessApi.md#ListHistoricProcessInstanceComments) | **Get** /history/historic-process-instances/{processInstanceId}/comments | List comments on a historic process instance
[**ListHistoricProcessInstanceIdentityLinks**](HistoryProcessApi.md#ListHistoricProcessInstanceIdentityLinks) | **Get** /history/historic-process-instances/{processInstanceId}/identitylinks | List identity links of a historic process instance
[**ListHistoricProcessInstances**](HistoryProcessApi.md#ListHistoricProcessInstances) | **Get** /history/historic-process-instances | List of historic process instances
[**QueryHistoricProcessInstance**](HistoryProcessApi.md#QueryHistoricProcessInstance) | **Post** /query/historic-process-instances | Query for historic process instances

# **CreateComment**
> CommentResponse CreateComment(ctx, processInstanceId, optional)
Create a new comment on a historic process instance

Parameter saveProcessInstanceId is optional, if true save process instance id of task with comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
 **optional** | ***HistoryProcessApiCreateCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryProcessApiCreateCommentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CommentResponse**](CommentResponse.md)|  | 

### Return type

[**CommentResponse**](CommentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteComment**
> DeleteComment(ctx, processInstanceId, commentId)
Delete a comment on a historic process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
  **commentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHistoricProcessInstance**
> DeleteHistoricProcessInstance(ctx, processInstanceId)
 Delete a historic process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComment**
> CommentResponse GetComment(ctx, processInstanceId, commentId)
Get a comment on a historic process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
  **commentId** | **string**|  | 

### Return type

[**CommentResponse**](CommentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoricProcessInstance**
> HistoricProcessInstanceResponse GetHistoricProcessInstance(ctx, processInstanceId)
Get a historic process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 

### Return type

[**HistoricProcessInstanceResponse**](HistoricProcessInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoricProcessInstanceVariableData**
> []string GetHistoricProcessInstanceVariableData(ctx, processInstanceId, variableName)
Get the binary data for a historic process instance variable

The response body contains the binary value of the variable. When the variable is of type binary, the content-type of the response is set to application/octet-stream, regardless of the content of the variable or the request accept-type header. In case of serializable, application/x-java-serialized-object is used as content-type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
  **variableName** | **string**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistoricProcessInstanceComments**
> []CommentResponse ListHistoricProcessInstanceComments(ctx, processInstanceId)
List comments on a historic process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 

### Return type

[**[]CommentResponse**](CommentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistoricProcessInstanceIdentityLinks**
> []HistoricIdentityLinkResponse ListHistoricProcessInstanceIdentityLinks(ctx, processInstanceId)
List identity links of a historic process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 

### Return type

[**[]HistoricIdentityLinkResponse**](HistoricIdentityLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistoricProcessInstances**
> DataResponseHistoricProcessInstanceResponse ListHistoricProcessInstances(ctx, optional)
List of historic process instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryProcessApiListHistoricProcessInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryProcessApiListHistoricProcessInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processInstanceId** | **optional.String**| An id of the historic process instance. | 
 **processInstanceName** | **optional.String**| A name of the historic process instance. | 
 **processInstanceNameLike** | **optional.String**| A name of the historic process instance used in a like query. | 
 **processInstanceNameLikeIgnoreCase** | **optional.String**| A name of the historic process instance used in a like query ignoring case. | 
 **processDefinitionKey** | **optional.String**| The process definition key of the historic process instance. | 
 **processDefinitionId** | **optional.String**| The process definition id of the historic process instance. | 
 **processDefinitionName** | **optional.String**| The process definition name of the historic process instance. | 
 **processDefinitionCategory** | **optional.String**| The process definition category of the historic process instance. | 
 **processDefinitionVersion** | **optional.String**| The process definition version of the historic process instance. | 
 **deploymentId** | **optional.String**| The deployment id of the historic process instance. | 
 **businessKey** | **optional.String**| The business key of the historic process instance. | 
 **businessKeyLike** | **optional.String**| Only return instances with a businessKey like this key. | 
 **activeActivityId** | **optional.String**| Only return instances which have an active activity instance with the provided activity id. | 
 **involvedUser** | **optional.String**| An involved user of the historic process instance. | 
 **finished** | **optional.Bool**| Indication if the historic process instance is finished. | 
 **superProcessInstanceId** | **optional.String**| An optional parent process id of the historic process instance. | 
 **excludeSubprocesses** | **optional.Bool**| Return only historic process instances which are not sub processes. | 
 **finishedAfter** | **optional.Time**| Return only historic process instances that were finished after this date. | 
 **finishedBefore** | **optional.Time**| Return only historic process instances that were finished before this date. | 
 **startedAfter** | **optional.Time**| Return only historic process instances that were started after this date. | 
 **startedBefore** | **optional.Time**| Return only historic process instances that were started before this date. | 
 **startedBy** | **optional.String**| Return only historic process instances that were started by this user. | 
 **includeProcessVariables** | **optional.Bool**| An indication if the historic process instance variables should be returned as well. | 
 **callbackId** | **optional.String**| Only return instances with the given callbackId. | 
 **callbackType** | **optional.String**| Only return instances with the given callbackType. | 
 **withoutCallbackId** | **optional.Bool**| Only return instances that do not have a callbackId. | 
 **tenantId** | **optional.String**| Only return instances with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return instances with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns instances without a tenantId set. If false, the withoutTenantId parameter is ignored.  | 

### Return type

[**DataResponseHistoricProcessInstanceResponse**](DataResponseHistoricProcessInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryHistoricProcessInstance**
> DataResponseHistoricProcessInstanceResponse QueryHistoricProcessInstance(ctx, optional)
Query for historic process instances

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of historic process instances, but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long. On top of that, the query allows for filtering based on process variables. The variables property is a JSON-array containing objects with the format as described here.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryProcessApiQueryHistoricProcessInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryProcessApiQueryHistoricProcessInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of HistoricProcessInstanceQueryRequest**](HistoricProcessInstanceQueryRequest.md)|  | 

### Return type

[**DataResponseHistoricProcessInstanceResponse**](DataResponseHistoricProcessInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

