# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskVariable**](TaskVariablesApi.md#CreateTaskVariable) | **Post** /runtime/tasks/{taskId}/variables | Create new variables on a task
[**DeleteTaskInstanceVariable**](TaskVariablesApi.md#DeleteTaskInstanceVariable) | **Delete** /runtime/tasks/{taskId}/variables/{variableName} | Delete a variable on a task
[**GetTaskInstanceVariable**](TaskVariablesApi.md#GetTaskInstanceVariable) | **Get** /runtime/tasks/{taskId}/variables/{variableName} | Get a variable from a task
[**GetTaskVariableData**](TaskVariablesApi.md#GetTaskVariableData) | **Get** /runtime/tasks/{taskId}/variables/{variableName}/data | Get the binary data for a variable
[**ListTaskVariables**](TaskVariablesApi.md#ListTaskVariables) | **Get** /runtime/tasks/{taskId}/variables | List variables for a task
[**UpdateTaskInstanceVariable**](TaskVariablesApi.md#UpdateTaskInstanceVariable) | **Put** /runtime/tasks/{taskId}/variables/{variableName} | Update an existing variable on a task

# **CreateTaskVariable**
> interface{} CreateTaskVariable(ctx, taskId, optional)
Create new variables on a task

This endpoint can be used in 2 ways: By passing a JSON Body (RestVariable or an Array of RestVariable) or by passing a multipart/form-data Object. It is possible to create simple (non-binary) variable or list of variables or new binary variable  Any number of variables can be passed into the request body array. NB: Swagger V2 specification does not support this use case that is why this endpoint might be buggy/incomplete if used with other tools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
 **optional** | ***TaskVariablesApiCreateTaskVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskVariablesApiCreateTaskVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TaskIdVariablesBody**](TaskIdVariablesBody.md)|  | 
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

# **DeleteTaskInstanceVariable**
> DeleteTaskInstanceVariable(ctx, taskId, variableName, optional)
Delete a variable on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***TaskVariablesApiDeleteTaskInstanceVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskVariablesApiDeleteTaskInstanceVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **optional.String**| Scope of variable to be returned. When local, only task-local variable value is returned. When global, only variable value from the task’s parent execution-hierarchy are returned. When the parameter is omitted, a local variable will be returned if it exists, otherwise a global variable. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskInstanceVariable**
> RestVariable GetTaskInstanceVariable(ctx, taskId, variableName, optional)
Get a variable from a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***TaskVariablesApiGetTaskInstanceVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskVariablesApiGetTaskInstanceVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **optional.String**| Scope of variable to be returned. When local, only task-local variable value is returned. When global, only variable value from the task’s parent execution-hierarchy are returned. When the parameter is omitted, a local variable will be returned if it exists, otherwise a global variable. | 

### Return type

[**RestVariable**](RestVariable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskVariableData**
> []string GetTaskVariableData(ctx, taskId, variableName, optional)
Get the binary data for a variable

The response body contains the binary value of the variable. When the variable is of type binary, the content-type of the response is set to application/octet-stream, regardless of the content of the variable or the request accept-type header. In case of serializable, application/x-java-serialized-object is used as content-type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***TaskVariablesApiGetTaskVariableDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskVariablesApiGetTaskVariableDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **optional.String**| Scope of variable to be returned. When local, only task-local variable value is returned. When global, only variable value from the task’s parent execution-hierarchy are returned. When the parameter is omitted, a local variable will be returned if it exists, otherwise a global variable. | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTaskVariables**
> []RestVariable ListTaskVariables(ctx, taskId, optional)
List variables for a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
 **optional** | ***TaskVariablesApiListTaskVariablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskVariablesApiListTaskVariablesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **optional.String**| Scope of variable to be returned. When local, only task-local variable value is returned. When global, only variable value from the task’s parent execution-hierarchy are returned. When the parameter is omitted, a local variable will be returned if it exists, otherwise a global variable. | 

### Return type

[**[]RestVariable**](RestVariable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTaskInstanceVariable**
> RestVariable UpdateTaskInstanceVariable(ctx, taskId, variableName, optional)
Update an existing variable on a task

This endpoint can be used in 2 ways: By passing a JSON Body (RestVariable) or by passing a multipart/form-data Object. It is possible to update simple (non-binary) variable or  binary variable  Any number of variables can be passed into the request body array. NB: Swagger V2 specification does not support this use case that is why this endpoint might be buggy/incomplete if used with other tools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***TaskVariablesApiUpdateTaskInstanceVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskVariablesApiUpdateTaskInstanceVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of VariablesVariableNameBody4**](VariablesVariableNameBody4.md)|  | 
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

