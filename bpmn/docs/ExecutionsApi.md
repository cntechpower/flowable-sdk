# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeExecutionActivityState**](ExecutionsApi.md#ChangeExecutionActivityState) | **Post** /runtime/executions/{executionId}/change-state | Change the state of an execution
[**CreateExecutionVariable**](ExecutionsApi.md#CreateExecutionVariable) | **Post** /runtime/executions/{executionId}/variables | Create variables on an execution
[**CreateOrUpdateExecutionVariable**](ExecutionsApi.md#CreateOrUpdateExecutionVariable) | **Put** /runtime/executions/{executionId}/variables | Update variables on an execution
[**DeleteLocalVariables**](ExecutionsApi.md#DeleteLocalVariables) | **Delete** /runtime/executions/{executionId}/variables | Delete all variables for an execution
[**DeletedExecutionVariable**](ExecutionsApi.md#DeletedExecutionVariable) | **Delete** /runtime/executions/{executionId}/variables/{variableName} | Delete a variable for an execution
[**ExecuteExecutionAction**](ExecutionsApi.md#ExecuteExecutionAction) | **Put** /runtime/executions | Signal event received
[**GetExecution**](ExecutionsApi.md#GetExecution) | **Get** /runtime/executions/{executionId} | Get an execution
[**GetExecutionVariable**](ExecutionsApi.md#GetExecutionVariable) | **Get** /runtime/executions/{executionId}/variables/{variableName} | Get a variable for an execution
[**GetExecutionVariableData**](ExecutionsApi.md#GetExecutionVariableData) | **Get** /runtime/executions/{executionId}/variables/{variableName}/data | Get the binary data for an execution
[**ListExecutionActiveActivities**](ExecutionsApi.md#ListExecutionActiveActivities) | **Get** /runtime/executions/{executionId}/activities | List active activities in an execution
[**ListExecutionVariables**](ExecutionsApi.md#ListExecutionVariables) | **Get** /runtime/executions/{executionId}/variables | List variables for an execution
[**ListExecutions**](ExecutionsApi.md#ListExecutions) | **Get** /runtime/executions | List of executions
[**PerformExecutionAction**](ExecutionsApi.md#PerformExecutionAction) | **Put** /runtime/executions/{executionId} | Execute an action on an execution
[**QueryExecutions**](ExecutionsApi.md#QueryExecutions) | **Post** /query/executions | Query executions
[**UpdateExecutionVariable**](ExecutionsApi.md#UpdateExecutionVariable) | **Put** /runtime/executions/{executionId}/variables/{variableName} | Update a variable on an execution

# **ChangeExecutionActivityState**
> ChangeExecutionActivityState(ctx, executionId, optional)
Change the state of an execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 
 **optional** | ***ExecutionsApiChangeExecutionActivityStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiChangeExecutionActivityStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExecutionChangeActivityStateRequest**](ExecutionChangeActivityStateRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateExecutionVariable**
> interface{} CreateExecutionVariable(ctx, executionId, optional)
Create variables on an execution

This endpoint can be used in 2 ways: By passing a JSON Body (array of RestVariable) or by passing a multipart/form-data Object. Any number of variables can be passed into the request body array. NB: Swagger V2 specification does not support this use case that is why this endpoint might be buggy/incomplete if used with other tools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 
 **optional** | ***ExecutionsApiCreateExecutionVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiCreateExecutionVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExecutionIdVariablesBody2**](ExecutionIdVariablesBody2.md)|  | 
 **name** | **optional.**|  | 
 **type_** | **optional.**|  | 
 **scope** | **optional.**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateExecutionVariable**
> interface{} CreateOrUpdateExecutionVariable(ctx, executionId, optional)
Update variables on an execution

This endpoint can be used in 2 ways: By passing a JSON Body (array of RestVariable) or by passing a multipart/form-data Object. Any number of variables can be passed into the request body array. NB: Swagger V2 specification does not support this use case that is why this endpoint might be buggy/incomplete if used with other tools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 
 **optional** | ***ExecutionsApiCreateOrUpdateExecutionVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiCreateOrUpdateExecutionVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExecutionIdVariablesBody**](ExecutionIdVariablesBody.md)|  | 
 **name** | **optional.**|  | 
 **type_** | **optional.**|  | 
 **scope** | **optional.**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLocalVariables**
> DeleteLocalVariables(ctx, executionId)
Delete all variables for an execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletedExecutionVariable**
> DeletedExecutionVariable(ctx, executionId, variableName, optional)
Delete a variable for an execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***ExecutionsApiDeletedExecutionVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiDeletedExecutionVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteExecutionAction**
> ExecuteExecutionAction(ctx, optional)
Signal event received

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExecutionsApiExecuteExecutionActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiExecuteExecutionActionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ExecutionActionRequest**](ExecutionActionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExecution**
> ExecutionResponse GetExecution(ctx, executionId)
Get an execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 

### Return type

[**ExecutionResponse**](ExecutionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExecutionVariable**
> RestVariable GetExecutionVariable(ctx, executionId, variableName, optional)
Get a variable for an execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***ExecutionsApiGetExecutionVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiGetExecutionVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **optional.String**|  | 

### Return type

[**RestVariable**](RestVariable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExecutionVariableData**
> []string GetExecutionVariableData(ctx, executionId, variableName, optional)
Get the binary data for an execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***ExecutionsApiGetExecutionVariableDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiGetExecutionVariableDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **optional.String**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExecutionActiveActivities**
> []string ListExecutionActiveActivities(ctx, executionId)
List active activities in an execution

Returns all activities which are active in the execution and in all child-executions (and their children, recursively), if any.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExecutionVariables**
> []RestVariable ListExecutionVariables(ctx, executionId, optional)
List variables for an execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 
 **optional** | ***ExecutionsApiListExecutionVariablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiListExecutionVariablesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **optional.String**|  | 

### Return type

[**[]RestVariable**](RestVariable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExecutions**
> DataResponseExecutionResponse ListExecutions(ctx, optional)
List of executions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExecutionsApiListExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiListExecutionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return models with the given version. | 
 **activityId** | **optional.String**| Only return executions with the given activity id. | 
 **processDefinitionKey** | **optional.String**| Only return process instances with the given process definition key. | 
 **processDefinitionId** | **optional.String**| Only return process instances with the given process definition id. | 
 **processInstanceId** | **optional.String**| Only return executions which are part of the process instance with the given id. | 
 **messageEventSubscriptionName** | **optional.String**| Only return executions which are subscribed to a message with the given name. | 
 **signalEventSubscriptionName** | **optional.String**| Only return executions which are subscribed to a signal with the given name. | 
 **parentId** | **optional.String**| Only return executions which are a direct child of the given execution. | 
 **tenantId** | **optional.String**| Only return process instances with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return process instances with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns process instances without a tenantId set. If false, the withoutTenantId parameter is ignored. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseExecutionResponse**](DataResponseExecutionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformExecutionAction**
> ExecutionResponse PerformExecutionAction(ctx, executionId, optional)
Execute an action on an execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 
 **optional** | ***ExecutionsApiPerformExecutionActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiPerformExecutionActionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExecutionActionRequest**](ExecutionActionRequest.md)|  | 

### Return type

[**ExecutionResponse**](ExecutionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryExecutions**
> DataResponseExecutionResponse QueryExecutions(ctx, optional)
Query executions

The request body can contain all possible filters that can be used in the List executions URL query. On top of these, it’s possible to provide an array of variables and processInstanceVariables to include in the query, with their format described here.  The general paging and sorting query-parameters can be used for this URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExecutionsApiQueryExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiQueryExecutionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ExecutionQueryRequest**](ExecutionQueryRequest.md)|  | 

### Return type

[**DataResponseExecutionResponse**](DataResponseExecutionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExecutionVariable**
> RestVariable UpdateExecutionVariable(ctx, executionId, variableName, optional)
Update a variable on an execution

This endpoint can be used in 2 ways: By passing a JSON Body (RestVariable) or by passing a multipart/form-data Object. NB: Swagger V2 specification does not support this use case that is why this endpoint might be buggy/incomplete if used with other tools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **executionId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***ExecutionsApiUpdateExecutionVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionsApiUpdateExecutionVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of VariablesVariableNameBody**](VariablesVariableNameBody.md)|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **name** | **optional.**|  | 
 **type_** | **optional.**|  | 
 **scope** | **optional.**|  | 

### Return type

[**RestVariable**](RestVariable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

