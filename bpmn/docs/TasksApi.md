# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateTasks**](TasksApi.md#BulkUpdateTasks) | **Put** /runtime/tasks | Update Tasks
[**CreateTask**](TasksApi.md#CreateTask) | **Post** /runtime/tasks | Create Task
[**CreateTaskVariable**](TasksApi.md#CreateTaskVariable) | **Post** /runtime/tasks/{taskId}/variables | Create new variables on a task
[**DeleteAllLocalTaskVariables**](TasksApi.md#DeleteAllLocalTaskVariables) | **Delete** /runtime/tasks/{taskId}/variables | Delete all local variables on a task
[**DeleteEvent**](TasksApi.md#DeleteEvent) | **Delete** /runtime/tasks/{taskId}/events/{eventId} | Delete an event on a task
[**DeleteTask**](TasksApi.md#DeleteTask) | **Delete** /runtime/tasks/{taskId} | Delete a task
[**ExecuteTaskAction**](TasksApi.md#ExecuteTaskAction) | **Post** /runtime/tasks/{taskId} | Tasks actions
[**GetEvent**](TasksApi.md#GetEvent) | **Get** /runtime/tasks/{taskId}/events/{eventId} | Get an event on a task
[**GetTask**](TasksApi.md#GetTask) | **Get** /runtime/tasks/{taskId} | Get a task
[**GetTaskForm**](TasksApi.md#GetTaskForm) | **Get** /runtime/tasks/{taskId}/form | Get a task form
[**ListTaskEvents**](TasksApi.md#ListTaskEvents) | **Get** /runtime/tasks/{taskId}/events | List events for a task
[**ListTaskSubtasks**](TasksApi.md#ListTaskSubtasks) | **Get** /runtime/tasks/{taskId}/subtasks | List of sub tasks for a task
[**ListTasks**](TasksApi.md#ListTasks) | **Get** /runtime/tasks | List of tasks
[**QueryTasks**](TasksApi.md#QueryTasks) | **Post** /query/tasks | Query for tasks
[**UpdateTask**](TasksApi.md#UpdateTask) | **Put** /runtime/tasks/{taskId} | Update a task

# **BulkUpdateTasks**
> DataResponseTaskResponse BulkUpdateTasks(ctx, optional)
Update Tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TasksApiBulkUpdateTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TasksApiBulkUpdateTasksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BulkTasksRequest**](BulkTasksRequest.md)|  | 

### Return type

[**DataResponseTaskResponse**](DataResponseTaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTask**
> TaskResponse CreateTask(ctx, optional)
Create Task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TasksApiCreateTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TasksApiCreateTaskOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TaskRequest**](TaskRequest.md)|  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTaskVariable**
> interface{} CreateTaskVariable(ctx, taskId, optional)
Create new variables on a task

This endpoint can be used in 2 ways: By passing a JSON Body (RestVariable or an Array of RestVariable) or by passing a multipart/form-data Object. It is possible to create simple (non-binary) variable or list of variables or new binary variable  Any number of variables can be passed into the request body array. NB: Swagger V2 specification does not support this use case that is why this endpoint might be buggy/incomplete if used with other tools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
 **optional** | ***TasksApiCreateTaskVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TasksApiCreateTaskVariableOpts struct
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

# **DeleteAllLocalTaskVariables**
> DeleteAllLocalTaskVariables(ctx, taskId)
Delete all local variables on a task

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

# **DeleteEvent**
> DeleteEvent(ctx, taskId, eventId)
Delete an event on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **eventId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTask**
> DeleteTask(ctx, taskId, optional)
Delete a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
 **optional** | ***TasksApiDeleteTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TasksApiDeleteTaskOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascadeHistory** | **optional.String**| Whether or not to delete the HistoricTask instance when deleting the task (if applicable). If not provided, this value defaults to false. | 
 **deleteReason** | **optional.String**| Reason why the task is deleted. This value is ignored when cascadeHistory is true. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteTaskAction**
> ExecuteTaskAction(ctx, taskId, optional)
Tasks actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
 **optional** | ***TasksApiExecuteTaskActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TasksApiExecuteTaskActionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TaskActionRequest**](TaskActionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvent**
> EventResponse GetEvent(ctx, taskId, eventId)
Get an event on a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **eventId** | **string**|  | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTask**
> TaskResponse GetTask(ctx, taskId)
Get a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskForm**
> string GetTaskForm(ctx, taskId)
Get a task form

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

# **ListTaskEvents**
> []EventResponse ListTaskEvents(ctx, taskId)
List events for a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

[**[]EventResponse**](EventResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTaskSubtasks**
> []TaskResponse ListTaskSubtasks(ctx, taskId)
List of sub tasks for a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

[**[]TaskResponse**](TaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTasks**
> DataResponseTaskResponse ListTasks(ctx, optional)
List of tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TasksApiListTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TasksApiListTasksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Only return models with the given version. | 
 **nameLike** | **optional.String**| Only return tasks with a name like the given name. | 
 **description** | **optional.String**| Only return tasks with the given description. | 
 **priority** | **optional.String**| Only return tasks with the given priority. | 
 **minimumPriority** | **optional.String**| Only return tasks with a priority greater than the given value. | 
 **maximumPriority** | **optional.String**| Only return tasks with a priority lower than the given value. | 
 **assignee** | **optional.String**| Only return tasks assigned to the given user. | 
 **assigneeLike** | **optional.String**| Only return tasks assigned with an assignee like the given value. | 
 **owner** | **optional.String**| Only return tasks owned by the given user. | 
 **ownerLike** | **optional.String**| Only return tasks assigned with an owner like the given value. | 
 **unassigned** | **optional.String**| Only return tasks that are not assigned to anyone. If false is passed, the value is ignored. | 
 **delegationState** | **optional.String**| Only return tasks that have the given delegation state. Possible values are pending and resolved. | 
 **candidateUser** | **optional.String**| Only return tasks that can be claimed by the given user. This includes both tasks where the user is an explicit candidate for and task that are claimable by a group that the user is a member of. | 
 **candidateGroup** | **optional.String**| Only return tasks that can be claimed by a user in the given group. | 
 **candidateGroups** | **optional.String**| Only return tasks that can be claimed by a user in the given groups. Values split by comma. | 
 **involvedUser** | **optional.String**| Only return tasks in which the given user is involved. | 
 **taskDefinitionKey** | **optional.String**| Only return tasks with the given task definition id. | 
 **taskDefinitionKeyLike** | **optional.String**| Only return tasks with a given task definition id like the given value. | 
 **taskDefinitionKeys** | **optional.String**| Only return tasks with the given task definition ids. | 
 **processInstanceId** | **optional.String**| Only return tasks which are part of the process instance with the given id. | 
 **processInstanceIdWithChildren** | **optional.String**| Only return tasks which are part of the process instance and its children with the given id. | 
 **withoutProcessInstanceId** | **optional.Bool**| If true, only returns tasks without a process instance id set. If false, the withoutProcessInstanceId parameter is ignored. | 
 **processInstanceBusinessKey** | **optional.String**| Only return tasks which are part of the process instance with the given business key. | 
 **processInstanceBusinessKeyLike** | **optional.String**| Only return tasks which are part of the process instance which has a business key like the given value. | 
 **processDefinitionId** | **optional.String**| Only return tasks which are part of a process instance which has a process definition with the given id. | 
 **processDefinitionKey** | **optional.String**| Only return tasks which are part of a process instance which has a process definition with the given key. | 
 **processDefinitionKeyLike** | **optional.String**| Only return tasks which are part of a process instance which has a process definition with a key like the given value. | 
 **processDefinitionName** | **optional.String**| Only return tasks which are part of a process instance which has a process definition with the given name. | 
 **processDefinitionNameLike** | **optional.String**| Only return tasks which are part of a process instance which has a process definition with a name like the given value. | 
 **executionId** | **optional.String**| Only return tasks which are part of the execution with the given id. | 
 **createdOn** | **optional.Time**| Only return tasks which are created on the given date. | 
 **createdBefore** | **optional.Time**| Only return tasks which are created before the given date. | 
 **createdAfter** | **optional.Time**| Only return tasks which are created after the given date. | 
 **dueDate** | **optional.Time**| Only return tasks which are due on the given date. | 
 **dueBefore** | **optional.Time**| Only return tasks which are due before the given date. | 
 **dueAfter** | **optional.Time**| Only return tasks which are due after the given date. | 
 **withoutDueDate** | **optional.Bool**| Only return tasks which do not have a due date. The property is ignored if the value is false. | 
 **excludeSubTasks** | **optional.Bool**| Only return tasks that are not a subtask of another task. | 
 **active** | **optional.Bool**| If true, only return tasks that are not suspended (either part of a process that is not suspended or not part of a process at all). If false, only tasks that are part of suspended process instances are returned. | 
 **includeTaskLocalVariables** | **optional.Bool**| Indication to include task local variables in the result. | 
 **includeProcessVariables** | **optional.Bool**| Indication to include process variables in the result. | 
 **scopeDefinitionId** | **optional.String**| Only return tasks with the given scopeDefinitionId. | 
 **scopeId** | **optional.String**| Only return tasks with the given scopeId. | 
 **withoutScopeId** | **optional.Bool**| If true, only returns tasks without a scope id set. If false, the withoutScopeId parameter is ignored. | 
 **scopeType** | **optional.String**| Only return tasks with the given scopeType. | 
 **propagatedStageInstanceId** | **optional.String**| Only return tasks which have the given id as propagated stage instance id | 
 **tenantId** | **optional.String**| Only return tasks with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return tasks with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns tasks without a tenantId set. If false, the withoutTenantId parameter is ignored. | 
 **candidateOrAssigned** | **optional.String**| Select tasks that has been claimed or assigned to user or waiting to claim by user (candidate user or groups). | 
 **category** | **optional.String**| Select tasks with the given category. Note that this is the task category, not the category of the process definition (namespace within the BPMN Xml).  | 
 **categoryIn** | **optional.String**| Select tasks for the given categories. Note that this is the task category, not the category of the process definition (namespace within the BPMN Xml).  | 
 **categoryNotIn** | **optional.String**| Select tasks which are not assigned to the given categories. Does not return tasks without categories. Note that this is the task category, not the category of the process definition (namespace within the BPMN Xml).  | 
 **withoutCategory** | **optional.String**| Select tasks without a category assigned. Note that this is the task category, not the category of the process definition (namespace within the BPMN Xml).  | 

### Return type

[**DataResponseTaskResponse**](DataResponseTaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryTasks**
> DataResponseTaskResponse QueryTasks(ctx, optional)
Query for tasks

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of tasks (except for candidateGroupIn which is only available in this POST task query REST service), but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long. On top of that, the query allows for filtering based on task and process variables. The taskVariables and processInstanceVariables are both JSON-arrays containing objects with the format as described here.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TasksApiQueryTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TasksApiQueryTasksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TaskQueryRequest**](TaskQueryRequest.md)|  | 

### Return type

[**DataResponseTaskResponse**](DataResponseTaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTask**
> TaskResponse UpdateTask(ctx, taskId, optional)
Update a task

All request values are optional. For example, you can only include the assignee attribute in the request body JSON-object, only updating the assignee of the task, leaving all other fields unaffected. When an attribute is explicitly included and is set to null, the task-value will be updated to null. Example: {\"dueDate\" : null} will clear the duedate of the task).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
 **optional** | ***TasksApiUpdateTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TasksApiUpdateTaskOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TaskRequest**](TaskRequest.md)|  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

