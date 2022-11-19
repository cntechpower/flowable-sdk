# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeadLetterJob**](JobsApi.md#DeleteDeadLetterJob) | **Delete** /management/deadletter-jobs/{jobId} | Delete a deadletter job
[**DeleteHistoryJob**](JobsApi.md#DeleteHistoryJob) | **Delete** /management/history-jobs/{jobId} | Delete a history job
[**DeleteJob**](JobsApi.md#DeleteJob) | **Delete** /management/jobs/{jobId} | Delete a job
[**DeleteSuspendedJob**](JobsApi.md#DeleteSuspendedJob) | **Delete** /management/suspended-jobs/{jobId} | Delete a suspended job
[**DeleteTimerJob**](JobsApi.md#DeleteTimerJob) | **Delete** /management/timer-jobs/{jobId} | Delete a timer job
[**ExecuteDeadLetterJobAction**](JobsApi.md#ExecuteDeadLetterJobAction) | **Post** /management/deadletter-jobs | Move a bulk of deadletter jobs. Accepts &#x27;move&#x27; and &#x27;moveToHistoryJob&#x27; as action.
[**ExecuteDeadLetterJobAction_0**](JobsApi.md#ExecuteDeadLetterJobAction_0) | **Post** /management/deadletter-jobs/{jobId} | Move a single deadletter job. Accepts &#x27;move&#x27; and &#x27;moveToHistoryJob&#x27; as action.
[**ExecuteHistoryJob**](JobsApi.md#ExecuteHistoryJob) | **Post** /management/history-jobs/{jobId} | Execute a history job
[**ExecuteJobAction**](JobsApi.md#ExecuteJobAction) | **Post** /management/jobs/{jobId} | Execute a single job
[**ExecuteTimerJobAction**](JobsApi.md#ExecuteTimerJobAction) | **Post** /management/timer-jobs/{jobId} | Move a single timer job
[**GetDeadLetterJobStacktrace**](JobsApi.md#GetDeadLetterJobStacktrace) | **Get** /management/deadletter-jobs/{jobId}/exception-stacktrace | Get the exception stacktrace for a deadletter job
[**GetDeadletterJob**](JobsApi.md#GetDeadletterJob) | **Get** /management/deadletter-jobs/{jobId} | Get a single deadletter job
[**GetHistoryJob**](JobsApi.md#GetHistoryJob) | **Get** /management/history-jobs/{jobId} | Get a single history job job
[**GetJob**](JobsApi.md#GetJob) | **Get** /management/jobs/{jobId} | Get a single job
[**GetJobStacktrace**](JobsApi.md#GetJobStacktrace) | **Get** /management/jobs/{jobId}/exception-stacktrace | Get the exception stacktrace for a job
[**GetSuspendedJob**](JobsApi.md#GetSuspendedJob) | **Get** /management/suspended-jobs/{jobId} | Get a single suspended job
[**GetSuspendedJobStacktrace**](JobsApi.md#GetSuspendedJobStacktrace) | **Get** /management/suspended-jobs/{jobId}/exception-stacktrace | Get the exception stacktrace for a suspended job
[**GetTimerJob**](JobsApi.md#GetTimerJob) | **Get** /management/timer-jobs/{jobId} | Get a single timer job
[**GetTimerJobStacktrace**](JobsApi.md#GetTimerJobStacktrace) | **Get** /management/timer-jobs/{jobId}/exception-stacktrace | Get the exception stacktrace for a timer job
[**ListDeadLetterJobs**](JobsApi.md#ListDeadLetterJobs) | **Get** /management/deadletter-jobs | List deadletter jobs
[**ListDeadLetterJobs_0**](JobsApi.md#ListDeadLetterJobs_0) | **Get** /management/history-jobs | List history jobs
[**ListJobs**](JobsApi.md#ListJobs) | **Get** /management/jobs | List jobs
[**ListSuspendedJobs**](JobsApi.md#ListSuspendedJobs) | **Get** /management/suspended-jobs | List suspended jobs
[**ListTimerJobs**](JobsApi.md#ListTimerJobs) | **Get** /management/timer-jobs | List timer jobs

# **DeleteDeadLetterJob**
> DeleteDeadLetterJob(ctx, jobId)
Delete a deadletter job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHistoryJob**
> DeleteHistoryJob(ctx, jobId)
Delete a history job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJob**
> DeleteJob(ctx, jobId)
Delete a job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSuspendedJob**
> DeleteSuspendedJob(ctx, jobId)
Delete a suspended job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTimerJob**
> DeleteTimerJob(ctx, jobId)
Delete a timer job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteDeadLetterJobAction**
> ExecuteDeadLetterJobAction(ctx, optional)
Move a bulk of deadletter jobs. Accepts 'move' and 'moveToHistoryJob' as action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobsApiExecuteDeadLetterJobActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiExecuteDeadLetterJobActionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BulkMoveDeadLetterActionRequest**](BulkMoveDeadLetterActionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteDeadLetterJobAction_0**
> ExecuteDeadLetterJobAction_0(ctx, jobId, optional)
Move a single deadletter job. Accepts 'move' and 'moveToHistoryJob' as action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 
 **optional** | ***JobsApiExecuteDeadLetterJobAction_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiExecuteDeadLetterJobAction_1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RestActionRequest**](RestActionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteHistoryJob**
> ExecuteHistoryJob(ctx, jobId, optional)
Execute a history job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 
 **optional** | ***JobsApiExecuteHistoryJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiExecuteHistoryJobOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RestActionRequest**](RestActionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteJobAction**
> ExecuteJobAction(ctx, jobId, optional)
Execute a single job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 
 **optional** | ***JobsApiExecuteJobActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiExecuteJobActionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RestActionRequest**](RestActionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteTimerJobAction**
> ExecuteTimerJobAction(ctx, jobId, optional)
Move a single timer job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 
 **optional** | ***JobsApiExecuteTimerJobActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiExecuteTimerJobActionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RestActionRequest**](RestActionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeadLetterJobStacktrace**
> string GetDeadLetterJobStacktrace(ctx, jobId)
Get the exception stacktrace for a deadletter job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeadletterJob**
> JobResponse GetDeadletterJob(ctx, jobId)
Get a single deadletter job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryJob**
> HistoryJobResponse GetHistoryJob(ctx, jobId)
Get a single history job job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

[**HistoryJobResponse**](HistoryJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJob**
> JobResponse GetJob(ctx, jobId)
Get a single job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobStacktrace**
> string GetJobStacktrace(ctx, jobId)
Get the exception stacktrace for a job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSuspendedJob**
> JobResponse GetSuspendedJob(ctx, jobId)
Get a single suspended job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSuspendedJobStacktrace**
> string GetSuspendedJobStacktrace(ctx, jobId)
Get the exception stacktrace for a suspended job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimerJob**
> JobResponse GetTimerJob(ctx, jobId)
Get a single timer job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimerJobStacktrace**
> string GetTimerJobStacktrace(ctx, jobId)
Get the exception stacktrace for a timer job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeadLetterJobs**
> DataResponseJobResponse ListDeadLetterJobs(ctx, optional)
List deadletter jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobsApiListDeadLetterJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiListDeadLetterJobsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return job with the given id | 
 **processInstanceId** | **optional.String**| Only return jobs part of a process with the given id | 
 **withoutProcessInstanceId** | **optional.Bool**| If true, only returns jobs without a process instance id set. If false, the withoutProcessInstanceId parameter is ignored. | 
 **executionId** | **optional.String**| Only return jobs part of an execution with the given id | 
 **processDefinitionId** | **optional.String**| Only return jobs with the given process definition id | 
 **elementId** | **optional.String**| Only return jobs with the given element id | 
 **elementName** | **optional.String**| Only return jobs with the given element name | 
 **executable** | **optional.Bool**| If true, only return jobs which are executable. If false, this parameter is ignored. | 
 **timersOnly** | **optional.Bool**| If true, only return jobs which are timers. If false, this parameter is ignored. Cannot be used together with &#x27;messagesOnly&#x27;. | 
 **messagesOnly** | **optional.Bool**| If true, only return jobs which are messages. If false, this parameter is ignored. Cannot be used together with &#x27;timersOnly&#x27; | 
 **withException** | **optional.Bool**| If true, only return jobs for which an exception occurred while executing it. If false, this parameter is ignored. | 
 **dueBefore** | **optional.Time**| Only return jobs which are due to be executed before the given date. Jobs without duedate are never returned using this parameter. | 
 **dueAfter** | **optional.Time**| Only return jobs which are due to be executed after the given date. Jobs without duedate are never returned using this parameter. | 
 **exceptionMessage** | **optional.String**| Only return jobs with the given exception message | 
 **tenantId** | **optional.String**| Only return jobs with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return jobs with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns jobs without a tenantId set. If false, the withoutTenantId parameter is ignored. | 
 **locked** | **optional.Bool**| If true, only return jobs which are locked.  If false, this parameter is ignored. | 
 **unlocked** | **optional.Bool**| If true, only return jobs which are unlocked. If false, this parameter is ignored. | 
 **withoutScopeId** | **optional.Bool**| If true, only returns jobs without a scope id set. If false, the withoutScopeId parameter is ignored. | 
 **withoutScopeType** | **optional.Bool**| If true, only returns jobs without a scope type set. If false, the withoutScopeType parameter is ignored. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseJobResponse**](DataResponseJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeadLetterJobs_0**
> DataResponseHistoryJobResponse ListDeadLetterJobs_0(ctx, optional)
List history jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobsApiListDeadLetterJobs_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiListDeadLetterJobs_2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return the job with the given id | 
 **withException** | **optional.Bool**| If true, only return jobs for which an exception occurred while executing it. If false, this parameter is ignored. | 
 **exceptionMessage** | **optional.String**| Only return jobs with the given exception message | 
 **scopeType** | **optional.String**| Only return jobs with the given scope type | 
 **tenantId** | **optional.String**| Only return jobs with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return jobs with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns jobs without a tenantId set. If false, the withoutTenantId parameter is ignored. | 
 **lockOwner** | **optional.String**| If true, only return jobs which are owned by the given lockOwner. | 
 **locked** | **optional.Bool**| If true, only return jobs which are locked.  If false, this parameter is ignored. | 
 **unlocked** | **optional.Bool**| If true, only return jobs which are unlocked. If false, this parameter is ignored. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseHistoryJobResponse**](DataResponseHistoryJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJobs**
> DataResponseJobResponse ListJobs(ctx, optional)
List jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobsApiListJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiListJobsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return job with the given id | 
 **processInstanceId** | **optional.String**| Only return jobs part of a process with the given id | 
 **withoutProcessInstanceId** | **optional.Bool**| If true, only returns jobs without a process instance id set. If false, the withoutProcessInstanceId parameter is ignored. | 
 **executionId** | **optional.String**| Only return jobs part of an execution with the given id | 
 **processDefinitionId** | **optional.String**| Only return jobs with the given process definition id | 
 **elementId** | **optional.String**| Only return jobs with the given element id | 
 **elementName** | **optional.String**| Only return jobs with the given element name | 
 **timersOnly** | **optional.Bool**| If true, only return jobs which are timers. If false, this parameter is ignored. Cannot be used together with &#x27;messagesOnly&#x27;. | 
 **messagesOnly** | **optional.Bool**| If true, only return jobs which are messages. If false, this parameter is ignored. Cannot be used together with &#x27;timersOnly&#x27; | 
 **withException** | **optional.Bool**| If true, only return jobs for which an exception occurred while executing it. If false, this parameter is ignored. | 
 **dueBefore** | **optional.Time**| Only return jobs which are due to be executed before the given date. Jobs without duedate are never returned using this parameter. | 
 **dueAfter** | **optional.Time**| Only return jobs which are due to be executed after the given date. Jobs without duedate are never returned using this parameter. | 
 **exceptionMessage** | **optional.String**| Only return jobs with the given exception message | 
 **tenantId** | **optional.String**| Only return jobs with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return jobs with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns jobs without a tenantId set. If false, the withoutTenantId parameter is ignored. | 
 **locked** | **optional.Bool**| If true, only return jobs which are locked.  If false, this parameter is ignored. | 
 **unlocked** | **optional.Bool**| If true, only return jobs which are unlocked. If false, this parameter is ignored. | 
 **withoutScopeId** | **optional.Bool**| If true, only returns jobs without a scope id set. If false, the withoutScopeId parameter is ignored. | 
 **withoutScopeType** | **optional.Bool**| If true, only returns jobs without a scope type set. If false, the withoutScopeType parameter is ignored. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseJobResponse**](DataResponseJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSuspendedJobs**
> DataResponseJobResponse ListSuspendedJobs(ctx, optional)
List suspended jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobsApiListSuspendedJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiListSuspendedJobsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return job with the given id | 
 **processInstanceId** | **optional.String**| Only return jobs part of a process with the given id | 
 **withoutProcessInstanceId** | **optional.Bool**| If true, only returns jobs without a process instance id set. If false, the withoutProcessInstanceId parameter is ignored. | 
 **executionId** | **optional.String**| Only return jobs part of an execution with the given id | 
 **processDefinitionId** | **optional.String**| Only return jobs with the given process definition id | 
 **elementId** | **optional.String**| Only return jobs with the given element id | 
 **elementName** | **optional.String**| Only return jobs with the given element name | 
 **executable** | **optional.Bool**| If true, only return jobs which are executable. If false, this parameter is ignored. | 
 **timersOnly** | **optional.Bool**| If true, only return jobs which are timers. If false, this parameter is ignored. Cannot be used together with &#x27;messagesOnly&#x27;. | 
 **messagesOnly** | **optional.Bool**| If true, only return jobs which are messages. If false, this parameter is ignored. Cannot be used together with &#x27;timersOnly&#x27; | 
 **withException** | **optional.Bool**| If true, only return jobs for which an exception occurred while executing it. If false, this parameter is ignored. | 
 **dueBefore** | **optional.Time**| Only return jobs which are due to be executed before the given date. Jobs without duedate are never returned using this parameter. | 
 **dueAfter** | **optional.Time**| Only return jobs which are due to be executed after the given date. Jobs without duedate are never returned using this parameter. | 
 **exceptionMessage** | **optional.String**| Only return jobs with the given exception message | 
 **tenantId** | **optional.String**| Only return jobs with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return jobs with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns jobs without a tenantId set. If false, the withoutTenantId parameter is ignored. | 
 **locked** | **optional.Bool**| If true, only return jobs which are locked.  If false, this parameter is ignored. | 
 **unlocked** | **optional.Bool**| If true, only return jobs which are unlocked. If false, this parameter is ignored. | 
 **withoutScopeId** | **optional.Bool**| If true, only returns jobs without a scope id set. If false, the withoutScopeId parameter is ignored. | 
 **withoutScopeType** | **optional.Bool**| If true, only returns jobs without a scope type set. If false, the withoutScopeType parameter is ignored. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseJobResponse**](DataResponseJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTimerJobs**
> DataResponseJobResponse ListTimerJobs(ctx, optional)
List timer jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobsApiListTimerJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiListTimerJobsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return job with the given id | 
 **processInstanceId** | **optional.String**| Only return jobs part of a process with the given id | 
 **withoutProcessInstanceId** | **optional.Bool**| If true, only returns jobs without a process instance id set. If false, the withoutProcessInstanceId parameter is ignored. | 
 **executionId** | **optional.String**| Only return jobs part of an execution with the given id | 
 **processDefinitionId** | **optional.String**| Only return jobs with the given process definition id | 
 **elementId** | **optional.String**| Only return jobs with the given element id | 
 **elementName** | **optional.String**| Only return jobs with the given element name | 
 **executable** | **optional.Bool**| If true, only return jobs which are executable. If false, this parameter is ignored. | 
 **timersOnly** | **optional.Bool**| If true, only return jobs which are timers. If false, this parameter is ignored. Cannot be used together with &#x27;messagesOnly&#x27;. | 
 **messagesOnly** | **optional.Bool**| If true, only return jobs which are messages. If false, this parameter is ignored. Cannot be used together with &#x27;timersOnly&#x27; | 
 **withException** | **optional.Bool**| If true, only return jobs for which an exception occurred while executing it. If false, this parameter is ignored. | 
 **dueBefore** | **optional.Time**| Only return jobs which are due to be executed before the given date. Jobs without duedate are never returned using this parameter. | 
 **dueAfter** | **optional.Time**| Only return jobs which are due to be executed after the given date. Jobs without duedate are never returned using this parameter. | 
 **exceptionMessage** | **optional.String**| Only return jobs with the given exception message | 
 **tenantId** | **optional.String**| Only return jobs with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return jobs with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns jobs without a tenantId set. If false, the withoutTenantId parameter is ignored. | 
 **locked** | **optional.Bool**| If true, only return jobs which are locked.  If false, this parameter is ignored. | 
 **unlocked** | **optional.Bool**| If true, only return jobs which are unlocked. If false, this parameter is ignored. | 
 **withoutScopeId** | **optional.Bool**| If true, only returns jobs without a scope id set. If false, the withoutScopeId parameter is ignored. | 
 **withoutScopeType** | **optional.Bool**| If true, only returns jobs without a scope type set. If false, the withoutScopeType parameter is ignored. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseJobResponse**](DataResponseJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

