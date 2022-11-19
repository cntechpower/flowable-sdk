# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeActivityState**](ProcessInstancesApi.md#ChangeActivityState) | **Post** /runtime/process-instances/{processInstanceId}/change-state | Change the state a process instance
[**CreateProcessInstance**](ProcessInstancesApi.md#CreateProcessInstance) | **Post** /runtime/process-instances | Start a process instance
[**DeleteProcessInstance**](ProcessInstancesApi.md#DeleteProcessInstance) | **Delete** /runtime/process-instances/{processInstanceId} | Delete a process instance
[**DeleteProcessInstances**](ProcessInstancesApi.md#DeleteProcessInstances) | **Post** /runtime/process-instances/delete | Bulk delete process instances
[**EvaluateConditions**](ProcessInstancesApi.md#EvaluateConditions) | **Post** /runtime/process-instances/{processInstanceId}/evaluate-conditions | Evaluate the conditions of a process instance
[**GetProcessInstance**](ProcessInstancesApi.md#GetProcessInstance) | **Get** /runtime/process-instances/{processInstanceId} | Get a process instance
[**GetProcessInstanceDiagram**](ProcessInstancesApi.md#GetProcessInstanceDiagram) | **Get** /runtime/process-instances/{processInstanceId}/diagram | Get diagram for a process instance
[**InjectActivityInProcessInstance**](ProcessInstancesApi.md#InjectActivityInProcessInstance) | **Post** /runtime/process-instances/{processInstanceId}/inject | Inject activity in a process instance
[**ListProcessInstances**](ProcessInstancesApi.md#ListProcessInstances) | **Get** /runtime/process-instances | List process instances
[**MigrateProcessInstance**](ProcessInstancesApi.md#MigrateProcessInstance) | **Post** /runtime/process-instances/{processInstanceId}/migrate | Migrate process instance
[**QueryProcessInstances**](ProcessInstancesApi.md#QueryProcessInstances) | **Post** /query/process-instances | Query process instances
[**UpdateProcessInstance**](ProcessInstancesApi.md#UpdateProcessInstance) | **Put** /runtime/process-instances/{processInstanceId} | Update process instance properties or execute an action on a process instance (body needs to contain an &#x27;action&#x27; property for the latter).

# **ChangeActivityState**
> ChangeActivityState(ctx, processInstanceId, optional)
Change the state a process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
 **optional** | ***ProcessInstancesApiChangeActivityStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstancesApiChangeActivityStateOpts struct
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

# **CreateProcessInstance**
> ProcessInstanceResponse CreateProcessInstance(ctx, optional)
Start a process instance

Note that also a *transientVariables* property is accepted as part of this json, that follows the same structure as the *variables* property.  Only one of *processDefinitionId*, *processDefinitionKey* or *message* can be used in the request body.   Parameters *businessKey*, *variables* and *tenantId* are optional.   If tenantId is omitted, the default tenant will be used. More information about the variable format can be found in the REST variables section.   Note that the variable-scope that is supplied is ignored, process-variables are always local.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProcessInstancesApiCreateProcessInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstancesApiCreateProcessInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProcessInstanceCreateRequest**](ProcessInstanceCreateRequest.md)|  | 

### Return type

[**ProcessInstanceResponse**](ProcessInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProcessInstance**
> DeleteProcessInstance(ctx, processInstanceId, optional)
Delete a process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
 **optional** | ***ProcessInstancesApiDeleteProcessInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstancesApiDeleteProcessInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteReason** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProcessInstances**
> DeleteProcessInstances(ctx, optional)
Bulk delete process instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProcessInstancesApiDeleteProcessInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstancesApiDeleteProcessInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BulkDeleteInstancesRestActionRequest**](BulkDeleteInstancesRestActionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateConditions**
> EvaluateConditions(ctx, processInstanceId)
Evaluate the conditions of a process instance

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

# **GetProcessInstance**
> ProcessInstanceResponse GetProcessInstance(ctx, processInstanceId)
Get a process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 

### Return type

[**ProcessInstanceResponse**](ProcessInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessInstanceDiagram**
> []string GetProcessInstanceDiagram(ctx, processInstanceId)
Get diagram for a process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InjectActivityInProcessInstance**
> InjectActivityInProcessInstance(ctx, processInstanceId, optional)
Inject activity in a process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
 **optional** | ***ProcessInstancesApiInjectActivityInProcessInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstancesApiInjectActivityInProcessInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of InjectActivityRequest**](InjectActivityRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProcessInstances**
> DataResponseProcessInstanceResponse ListProcessInstances(ctx, optional)
List process instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProcessInstancesApiListProcessInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstancesApiListProcessInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return models with the given version. | 
 **name** | **optional.String**| Only return models with the given name. | 
 **nameLike** | **optional.String**| Only return models like the given name. | 
 **nameLikeIgnoreCase** | **optional.String**| Only return models like the given name ignoring case. | 
 **processDefinitionKey** | **optional.String**| Only return process instances with the given process definition key. | 
 **processDefinitionId** | **optional.String**| Only return process instances with the given process definition id. | 
 **processDefinitionCategory** | **optional.String**| Only return process instances with the given process definition category. | 
 **processDefinitionVersion** | **optional.Int32**| Only return process instances with the given process definition version. | 
 **processDefinitionEngineVersion** | **optional.String**| Only return process instances with the given process definition engine version. | 
 **businessKey** | **optional.String**| Only return process instances with the given businessKey. | 
 **businessKeyLike** | **optional.String**| Only return process instances with the businessKey like the given key. | 
 **startedBy** | **optional.String**| Only return process instances started by the given user. | 
 **startedBefore** | **optional.Time**| Only return process instances started before the given date. | 
 **startedAfter** | **optional.Time**| Only return process instances started after the given date. | 
 **activeActivityId** | **optional.String**| Only return process instances which have an active activity instance with the provided activity id. | 
 **involvedUser** | **optional.String**| Only return process instances in which the given user is involved. | 
 **suspended** | **optional.Bool**| If true, only return process instance which are suspended. If false, only return process instances which are not suspended (active). | 
 **superProcessInstanceId** | **optional.String**| Only return process instances which have the given super process-instance id (for processes that have a call-activities). | 
 **subProcessInstanceId** | **optional.String**| Only return process instances which have the given sub process-instance id (for processes started as a call-activity). | 
 **excludeSubprocesses** | **optional.Bool**| Return only process instances which are not sub processes. | 
 **includeProcessVariables** | **optional.Bool**| Indication to include process variables in the result. | 
 **callbackId** | **optional.String**| Only return process instances with the given callbackId. | 
 **callbackType** | **optional.String**| Only return process instances with the given callbackType. | 
 **tenantId** | **optional.String**| Only return process instances with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return process instances with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns process instances without a tenantId set. If false, the withoutTenantId parameter is ignored. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseProcessInstanceResponse**](DataResponseProcessInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateProcessInstance**
> MigrateProcessInstance(ctx, processInstanceId, optional)
Migrate process instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
 **optional** | ***ProcessInstancesApiMigrateProcessInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstancesApiMigrateProcessInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryProcessInstances**
> DataResponseProcessInstanceResponse QueryProcessInstances(ctx, optional)
Query process instances

The request body can contain all possible filters that can be used in the List process instances URL query. On top of these, it’s possible to provide an array of variables to include in the query, with their format described here.  The general paging and sorting query-parameters can be used for this URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProcessInstancesApiQueryProcessInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstancesApiQueryProcessInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProcessInstanceQueryRequest**](ProcessInstanceQueryRequest.md)|  | 

### Return type

[**DataResponseProcessInstanceResponse**](DataResponseProcessInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProcessInstance**
> ProcessInstanceResponse UpdateProcessInstance(ctx, processInstanceId, optional)
Update process instance properties or execute an action on a process instance (body needs to contain an 'action' property for the latter).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processInstanceId** | **string**|  | 
 **optional** | ***ProcessInstancesApiUpdateProcessInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessInstancesApiUpdateProcessInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProcessInstanceUpdateRequest**](ProcessInstanceUpdateRequest.md)|  | 

### Return type

[**ProcessInstanceResponse**](ProcessInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

