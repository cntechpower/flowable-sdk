# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateProcessVariable**](ProcessInstanceVariablesApi.md#CreateOrUpdateProcessVariable) | **Put** /runtime/process-instances/{processInstanceId}/variables | Update a multiple/single (non)binary variable on a process instance
[**CreateProcessInstanceVariable**](ProcessInstanceVariablesApi.md#CreateProcessInstanceVariable) | **Post** /runtime/process-instances/{processInstanceId}/variables | Create variables or new binary variable on a process instance
[**DeleteLocalProcessVariable**](ProcessInstanceVariablesApi.md#DeleteLocalProcessVariable) | **Delete** /runtime/process-instances/{processInstanceId}/variables | Delete all variables
[**DeleteProcessInstanceVariable**](ProcessInstanceVariablesApi.md#DeleteProcessInstanceVariable) | **Delete** /runtime/process-instances/{processInstanceId}/variables/{variableName} | Delete a variable
[**GetProcessInstanceVariable**](ProcessInstanceVariablesApi.md#GetProcessInstanceVariable) | **Get** /runtime/process-instances/{processInstanceId}/variables/{variableName} | Get a variable for a process instance
[**GetProcessInstanceVariableData**](ProcessInstanceVariablesApi.md#GetProcessInstanceVariableData) | **Get** /runtime/process-instances/{processInstanceId}/variables/{variableName}/data | Get the binary data for a variable
[**ListProcessInstanceVariables**](ProcessInstanceVariablesApi.md#ListProcessInstanceVariables) | **Get** /runtime/process-instances/{processInstanceId}/variables | List variables for a process instance
[**UpdateProcessInstanceVariable**](ProcessInstanceVariablesApi.md#UpdateProcessInstanceVariable) | **Put** /runtime/process-instances/{processInstanceId}/variables/{variableName} | Update a single variable on a process instance

# **CreateOrUpdateProcessVariable**
> interface{} CreateOrUpdateProcessVariable(ctx, processInstanceId, optional)
Update a multiple/single (non)binary variable on a process instance

This endpoint can be used in 2 ways: By passing a JSON Body (RestVariable or an array of RestVariable) or by passing a multipart/form-data Object. Nonexistent variables are created on the process-instance and existing ones are overridden without any error. Any number of variables can be passed into the request body array. Note that scope is ignored, only local variables can be set in a process instance. NB: Swagger V2 specification does not support this use case that is why this endpoint might be buggy/incomplete if used with other tools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
 **optional** | ***ProcessInstanceVariablesApiCreateOrUpdateProcessVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstanceVariablesApiCreateOrUpdateProcessVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProcessInstanceIdVariablesBody**](ProcessInstanceIdVariablesBody.md)|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **name** | **optional.**|  | 
 **type_** | **optional.**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProcessInstanceVariable**
> interface{} CreateProcessInstanceVariable(ctx, processInstanceId, optional)
Create variables or new binary variable on a process instance

This endpoint can be used in 2 ways: By passing a JSON Body (RestVariable or an array of RestVariable) or by passing a multipart/form-data Object. Nonexistent variables are created on the process-instance and existing ones are overridden without any error. Any number of variables can be passed into the request body array. Note that scope is ignored, only local variables can be set in a process instance. NB: Swagger V2 specification does not support this use case that is why this endpoint might be buggy/incomplete if used with other tools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
 **optional** | ***ProcessInstanceVariablesApiCreateProcessInstanceVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstanceVariablesApiCreateProcessInstanceVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProcessInstanceIdVariablesBody2**](ProcessInstanceIdVariablesBody2.md)|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **name** | **optional.**|  | 
 **type_** | **optional.**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLocalProcessVariable**
> DeleteLocalProcessVariable(ctx, processInstanceId)
Delete all variables

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

# **DeleteProcessInstanceVariable**
> DeleteProcessInstanceVariable(ctx, processInstanceId, variableName, optional)
Delete a variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***ProcessInstanceVariablesApiDeleteProcessInstanceVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstanceVariablesApiDeleteProcessInstanceVariableOpts struct
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

# **GetProcessInstanceVariable**
> RestVariable GetProcessInstanceVariable(ctx, processInstanceId, variableName, optional)
Get a variable for a process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***ProcessInstanceVariablesApiGetProcessInstanceVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstanceVariablesApiGetProcessInstanceVariableOpts struct
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

# **GetProcessInstanceVariableData**
> []string GetProcessInstanceVariableData(ctx, processInstanceId, variableName, optional)
Get the binary data for a variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***ProcessInstanceVariablesApiGetProcessInstanceVariableDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstanceVariablesApiGetProcessInstanceVariableDataOpts struct
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

# **ListProcessInstanceVariables**
> []RestVariable ListProcessInstanceVariables(ctx, processInstanceId, optional)
List variables for a process instance

In case the variable is a binary variable or serializable, the valueUrl points to an URL to fetch the raw value. If it’s a plain variable, the value is present in the response. Note that only local scoped variables are returned, as there is no global scope for process-instance variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
 **optional** | ***ProcessInstanceVariablesApiListProcessInstanceVariablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstanceVariablesApiListProcessInstanceVariablesOpts struct
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

# **UpdateProcessInstanceVariable**
> RestVariable UpdateProcessInstanceVariable(ctx, processInstanceId, variableName, optional)
Update a single variable on a process instance

This endpoint can be used in 2 ways: By passing a JSON Body (RestVariable) or by passing a multipart/form-data Object. Nonexistent variables are created on the process-instance and existing ones are overridden without any error. Note that scope is ignored, only local variables can be set in a process instance. NB: Swagger V2 specification does not support this use case that is why this endpoint might be buggy/incomplete if used with other tools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***ProcessInstanceVariablesApiUpdateProcessInstanceVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstanceVariablesApiUpdateProcessInstanceVariableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of VariablesVariableNameBody2**](VariablesVariableNameBody2.md)|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **name** | **optional.**|  | 
 **type_** | **optional.**|  | 

### Return type

[**RestVariable**](RestVariable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

