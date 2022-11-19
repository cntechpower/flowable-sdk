# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryActivityInstances**](QueryApi.md#QueryActivityInstances) | **Post** /query/activity-instances | Query for activity instances
[**QueryActivityInstances_0**](QueryApi.md#QueryActivityInstances_0) | **Post** /query/historic-activity-instances | Query for historic activity instances
[**QueryExecutions**](QueryApi.md#QueryExecutions) | **Post** /query/executions | Query executions
[**QueryHistoricDetail**](QueryApi.md#QueryHistoricDetail) | **Post** /query/historic-detail | Query for historic details
[**QueryHistoricProcessInstance**](QueryApi.md#QueryHistoricProcessInstance) | **Post** /query/historic-process-instances | Query for historic process instances
[**QueryHistoricTaskInstance**](QueryApi.md#QueryHistoricTaskInstance) | **Post** /query/historic-task-instances | Query for historic task instances
[**QueryProcessInstances**](QueryApi.md#QueryProcessInstances) | **Post** /query/process-instances | Query process instances
[**QueryTasks**](QueryApi.md#QueryTasks) | **Post** /query/tasks | Query for tasks
[**QueryVariableInstances**](QueryApi.md#QueryVariableInstances) | **Post** /query/historic-variable-instances | Query for historic variable instances
[**QueryVariableInstances_0**](QueryApi.md#QueryVariableInstances_0) | **Post** /query/variable-instances | Query for variable instances

# **QueryActivityInstances**
> DataResponseActivityInstanceResponse QueryActivityInstances(ctx, optional)
Query for activity instances

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of historic task instances, but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryApiQueryActivityInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryApiQueryActivityInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ActivityInstanceQueryRequest**](ActivityInstanceQueryRequest.md)|  | 

### Return type

[**DataResponseActivityInstanceResponse**](DataResponseActivityInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryActivityInstances_0**
> DataResponseHistoricActivityInstanceResponse QueryActivityInstances_0(ctx, optional)
Query for historic activity instances

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of historic task instances, but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryApiQueryActivityInstances_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryApiQueryActivityInstances_1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of HistoricActivityInstanceQueryRequest**](HistoricActivityInstanceQueryRequest.md)|  | 

### Return type

[**DataResponseHistoricActivityInstanceResponse**](DataResponseHistoricActivityInstanceResponse.md)

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
 **optional** | ***QueryApiQueryExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryApiQueryExecutionsOpts struct
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

# **QueryHistoricDetail**
> DataResponseHistoricDetailResponse QueryHistoricDetail(ctx, optional)
Query for historic details

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of historic process instances, but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryApiQueryHistoricDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryApiQueryHistoricDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of HistoricDetailQueryRequest**](HistoricDetailQueryRequest.md)|  | 

### Return type

[**DataResponseHistoricDetailResponse**](DataResponseHistoricDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
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
 **optional** | ***QueryApiQueryHistoricProcessInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryApiQueryHistoricProcessInstanceOpts struct
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

# **QueryHistoricTaskInstance**
> DataResponseHistoricTaskInstanceResponse QueryHistoricTaskInstance(ctx, optional)
Query for historic task instances

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of historic task instances, but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long. On top of that, the query allows for filtering based on process variables. The taskVariables and processVariables properties are JSON-arrays containing objects with the format as described here.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryApiQueryHistoricTaskInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryApiQueryHistoricTaskInstanceOpts struct
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

# **QueryProcessInstances**
> DataResponseProcessInstanceResponse QueryProcessInstances(ctx, optional)
Query process instances

The request body can contain all possible filters that can be used in the List process instances URL query. On top of these, it’s possible to provide an array of variables to include in the query, with their format described here.  The general paging and sorting query-parameters can be used for this URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryApiQueryProcessInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryApiQueryProcessInstancesOpts struct
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

# **QueryTasks**
> DataResponseTaskResponse QueryTasks(ctx, optional)
Query for tasks

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of tasks (except for candidateGroupIn which is only available in this POST task query REST service), but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long. On top of that, the query allows for filtering based on task and process variables. The taskVariables and processInstanceVariables are both JSON-arrays containing objects with the format as described here.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryApiQueryTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryApiQueryTasksOpts struct
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

# **QueryVariableInstances**
> DataResponseHistoricVariableInstanceResponse QueryVariableInstances(ctx, optional)
Query for historic variable instances

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of historic process instances, but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long. On top of that, the query allows for filtering based on process variables. The variables property is a JSON-array containing objects with the format as described here.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryApiQueryVariableInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryApiQueryVariableInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of HistoricVariableInstanceQueryRequest**](HistoricVariableInstanceQueryRequest.md)|  | 

### Return type

[**DataResponseHistoricVariableInstanceResponse**](DataResponseHistoricVariableInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryVariableInstances_0**
> DataResponseVariableInstanceResponse QueryVariableInstances_0(ctx, optional)
Query for variable instances

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of variable instances, but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryApiQueryVariableInstances_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryApiQueryVariableInstances_2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VariableInstanceQueryRequest**](VariableInstanceQueryRequest.md)|  | 

### Return type

[**DataResponseVariableInstanceResponse**](DataResponseVariableInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

