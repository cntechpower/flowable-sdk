# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaskInstance**](HistoryTaskApi.md#DeleteTaskInstance) | **Delete** /history/historic-task-instances/{taskId} | Delete a historic task instance
[**GetHistoricTaskForm**](HistoryTaskApi.md#GetHistoricTaskForm) | **Get** /history/historic-task-instances/{taskId}/form | Get a historic task instance form
[**GetHistoricTaskLogEntries**](HistoryTaskApi.md#GetHistoricTaskLogEntries) | **Get** /history/historic-task-log-entries | List historic task log entries
[**GetTaskInstance**](HistoryTaskApi.md#GetTaskInstance) | **Get** /history/historic-task-instances/{taskId} | Get a single historic task instance
[**ListHistoricTaskInstanceIdentityLinks**](HistoryTaskApi.md#ListHistoricTaskInstanceIdentityLinks) | **Get** /history/historic-task-instances/{taskId}/identitylinks | List identity links of a historic task instance
[**ListHistoricTaskInstances**](HistoryTaskApi.md#ListHistoricTaskInstances) | **Get** /history/historic-task-instances | List historic task instances
[**QueryHistoricTaskInstance**](HistoryTaskApi.md#QueryHistoricTaskInstance) | **Post** /query/historic-task-instances | Query for historic task instances

# **DeleteTaskInstance**
> DeleteTaskInstance(ctx, taskId)
Delete a historic task instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoricTaskForm**
> string GetHistoricTaskForm(ctx, taskId)
Get a historic task instance form

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoricTaskLogEntries**
> DataResponseHistoricTaskLogEntryResponse GetHistoricTaskLogEntries(ctx, optional)
List historic task log entries

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryTaskApiGetHistoricTaskLogEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryTaskApiGetHistoricTaskLogEntriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **optional.String**| An id of the historic task instance. | 
 **type_** | **optional.String**| The type of the log entry. | 
 **userId** | **optional.String**| The user who produced the task change. | 
 **processInstanceId** | **optional.String**| The process instance id of the historic task log entry. | 
 **processDefinitionId** | **optional.String**| The process definition id of the historic task log entry. | 
 **scopeId** | **optional.String**| Only return historic task log entries with the given scopeId. | 
 **scopeDefinitionId** | **optional.String**| Only return historic task log entries with the given scopeDefinitionId. | 
 **subScopeId** | **optional.String**| Only return historic task log entries with the given subScopeId | 
 **scopeType** | **optional.String**| Only return historic task log entries with the given scopeType. | 
 **from** | **optional.Time**| Return task log entries starting from a date. | 
 **to** | **optional.Time**| Return task log entries up to a date. | 
 **tenantId** | **optional.String**| Only return historic task log entries with the given tenantId. | 
 **fromLogNumber** | **optional.String**| Return task log entries starting from a log number | 
 **toLogNumber** | **optional.String**| Return task log entries up to specific a log number | 

### Return type

[**DataResponseHistoricTaskLogEntryResponse**](DataResponseHistoricTaskLogEntryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskInstance**
> HistoricTaskInstanceResponse GetTaskInstance(ctx, taskId)
Get a single historic task instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

[**HistoricTaskInstanceResponse**](HistoricTaskInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistoricTaskInstanceIdentityLinks**
> []HistoricIdentityLinkResponse ListHistoricTaskInstanceIdentityLinks(ctx, taskId)
List identity links of a historic task instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

[**[]HistoricIdentityLinkResponse**](HistoricIdentityLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistoricTaskInstances**
> DataResponseHistoricTaskInstanceResponse ListHistoricTaskInstances(ctx, optional)
List historic task instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryTaskApiListHistoricTaskInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryTaskApiListHistoricTaskInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **optional.String**| An id of the historic task instance. | 
 **processInstanceId** | **optional.String**| The process instance id of the historic task instance. | 
 **processInstanceIdWithChildren** | **optional.String**| Selects the historic task instances for the process instance and its children. | 
 **withoutProcessInstanceId** | **optional.Bool**| If true, only returns historic task instances without a process instance id set. If false, the withoutProcessInstanceId parameter is ignored. | 
 **processDefinitionKey** | **optional.String**| The process definition key of the historic task instance. | 
 **processDefinitionKeyLike** | **optional.String**| The process definition key of the historic task instance, which matches the given value. | 
 **processDefinitionId** | **optional.String**| The process definition id of the historic task instance. | 
 **processDefinitionName** | **optional.String**| The process definition name of the historic task instance. | 
 **processDefinitionNameLike** | **optional.String**| The process definition name of the historic task instance, which matches the given value. | 
 **processBusinessKey** | **optional.String**| The process instance business key of the historic task instance. | 
 **processBusinessKeyLike** | **optional.String**| The process instance business key of the historic task instance that matches the given value. | 
 **executionId** | **optional.String**| The execution id of the historic task instance. | 
 **taskDefinitionKey** | **optional.String**| The task definition key for tasks part of a process | 
 **taskDefinitionKeys** | **optional.String**| The task definition key for tasks part of a process | 
 **taskName** | **optional.String**| The task name of the historic task instance. | 
 **taskNameLike** | **optional.String**| The task name with like operator for the historic task instance. | 
 **taskDescription** | **optional.String**| The task description of the historic task instance. | 
 **taskDescriptionLike** | **optional.String**| The task description with like operator for the historic task instance. | 
 **taskCategory** | **optional.String**| Select tasks with the given category. Note that this is the task category, not the category of the process definition (namespace within the BPMN Xml). | 
 **taskCategoryIn** | **optional.String**| Select tasks with the given categories. Note that this is the task category, not the category of the process definition (namespace within the BPMN Xml). | 
 **taskCategoryNotIn** | **optional.String**| Select tasks not assigned to the given categories. Note that this is the task category, not the category of the process definition (namespace within the BPMN Xml). | 
 **taskWithoutCategory** | **optional.String**| Select tasks with no category assigned. Note that this is the task category, not the category of the process definition (namespace within the BPMN Xml). | 
 **taskDeleteReason** | **optional.String**| The task delete reason of the historic task instance. | 
 **taskDeleteReasonLike** | **optional.String**| The task delete reason with like operator for the historic task instance. | 
 **taskAssignee** | **optional.String**| The assignee of the historic task instance. | 
 **taskAssigneeLike** | **optional.String**| The assignee with like operator for the historic task instance. | 
 **taskOwner** | **optional.String**| The owner of the historic task instance. | 
 **taskOwnerLike** | **optional.String**| The owner with like operator for the historic task instance. | 
 **taskInvolvedUser** | **optional.String**| An involved user of the historic task instance | 
 **taskCandidateGroup** | **optional.String**| Only return tasks that can be claimed by a user in the given group. | 
 **taskPriority** | **optional.String**| The priority of the historic task instance. | 
 **finished** | **optional.Bool**| Indication if the historic task instance is finished. | 
 **processFinished** | **optional.Bool**| Indication if the process instance of the historic task instance is finished. | 
 **parentTaskId** | **optional.String**| An optional parent task id of the historic task instance. | 
 **dueDate** | **optional.Time**| Return only historic task instances that have a due date equal this date. | 
 **dueDateAfter** | **optional.Time**| Return only historic task instances that have a due date after this date. | 
 **dueDateBefore** | **optional.Time**| Return only historic task instances that have a due date before this date. | 
 **withoutDueDate** | **optional.Bool**| Return only historic task instances that have no due-date. When false is provided as value, this parameter is ignored. | 
 **taskCompletedOn** | **optional.Time**| Return only historic task instances that have been completed on this date. | 
 **taskCompletedAfter** | **optional.Time**| Return only historic task instances that have been completed after this date. | 
 **taskCompletedBefore** | **optional.Time**| Return only historic task instances that have been completed before this date. | 
 **taskCreatedOn** | **optional.Time**| Return only historic task instances that were created on this date. | 
 **taskCreatedBefore** | **optional.Time**| Return only historic task instances that were created before this date. | 
 **taskCreatedAfter** | **optional.Time**| Return only historic task instances that were created after this date. | 
 **includeTaskLocalVariables** | **optional.Bool**| An indication if the historic task instance local variables should be returned as well. | 
 **includeProcessVariables** | **optional.Bool**| An indication if the historic task instance global variables should be returned as well. | 
 **scopeDefinitionId** | **optional.String**| Only return historic task instances with the given scopeDefinitionId. | 
 **scopeId** | **optional.String**| Only return historic task instances with the given scopeId. | 
 **withoutScopeId** | **optional.Bool**| If true, only returns historic task instances without a scope id set. If false, the withoutScopeId parameter is ignored. | 
 **scopeType** | **optional.String**| Only return historic task instances with the given scopeType. | 
 **propagatedStageInstanceId** | **optional.String**| Only return tasks which have the given id as propagated stage instance id | 
 **tenantId** | **optional.String**| Only return historic task instances with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return historic task instances with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns historic task instances without a tenantId set. If false, the withoutTenantId parameter is ignored. | 

### Return type

[**DataResponseHistoricTaskInstanceResponse**](DataResponseHistoricTaskInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryHistoricTaskInstance**
> DataResponseHistoricTaskInstanceResponse QueryHistoricTaskInstance(ctx, optional)
Query for historic task instances

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of historic task instances, but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long. On top of that, the query allows for filtering based on process variables. The taskVariables and processVariables properties are JSON-arrays containing objects with the format as described here.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryTaskApiQueryHistoricTaskInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryTaskApiQueryHistoricTaskInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of HistoricTaskInstanceQueryRequest**](HistoricTaskInstanceQueryRequest.md)|  | 

### Return type

[**DataResponseHistoricTaskInstanceResponse**](DataResponseHistoricTaskInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

