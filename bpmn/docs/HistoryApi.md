# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricDetailVariableData**](HistoryApi.md#GetHistoricDetailVariableData) | **Get** /history/historic-detail/{detailId}/data | Get the binary data for a historic detail variable
[**GetHistoricInstanceVariableData**](HistoryApi.md#GetHistoricInstanceVariableData) | **Get** /history/historic-variable-instances/{varInstanceId}/data | Get the binary data for a historic task instance variable
[**GetHistoricTaskInstanceVariableData**](HistoryApi.md#GetHistoricTaskInstanceVariableData) | **Get** /history/historic-task-instances/{taskId}/variables/{variableName}/data | Get the binary data for a historic task instance variable
[**ListActivityInstances**](HistoryApi.md#ListActivityInstances) | **Get** /runtime/activity-instances | List activity instances
[**ListHistoricActivityInstances**](HistoryApi.md#ListHistoricActivityInstances) | **Get** /history/historic-activity-instances | List historic activity instances
[**ListHistoricDetails**](HistoryApi.md#ListHistoricDetails) | **Get** /history/historic-detail | Get historic detail
[**ListHistoricVariableInstances**](HistoryApi.md#ListHistoricVariableInstances) | **Get** /history/historic-variable-instances | List of historic variable instances
[**ListVariableInstances**](HistoryApi.md#ListVariableInstances) | **Get** /runtime/variable-instances | List of variable instances
[**QueryActivityInstances**](HistoryApi.md#QueryActivityInstances) | **Post** /query/historic-activity-instances | Query for historic activity instances
[**QueryHistoricDetail**](HistoryApi.md#QueryHistoricDetail) | **Post** /query/historic-detail | Query for historic details
[**QueryVariableInstances**](HistoryApi.md#QueryVariableInstances) | **Post** /query/historic-variable-instances | Query for historic variable instances
[**QueryVariableInstances_0**](HistoryApi.md#QueryVariableInstances_0) | **Post** /query/variable-instances | Query for variable instances

# **GetHistoricDetailVariableData**
> []string GetHistoricDetailVariableData(ctx, detailId)
Get the binary data for a historic detail variable

The response body contains the binary value of the variable. When the variable is of type binary, the content-type of the response is set to application/octet-stream, regardless of the content of the variable or the request accept-type header. In case of serializable, application/x-java-serialized-object is used as content-type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **detailId** | **string**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoricInstanceVariableData**
> []string GetHistoricInstanceVariableData(ctx, varInstanceId)
Get the binary data for a historic task instance variable

The response body contains the binary value of the variable. When the variable is of type binary, the content-type of the response is set to application/octet-stream, regardless of the content of the variable or the request accept-type header. In case of serializable, application/x-java-serialized-object is used as content-type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **varInstanceId** | **string**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoricTaskInstanceVariableData**
> []string GetHistoricTaskInstanceVariableData(ctx, taskId, variableName, optional)
Get the binary data for a historic task instance variable

The response body contains the binary value of the variable. When the variable is of type binary, the content-type of the response is set to application/octet-stream, regardless of the content of the variable or the request accept-type header. In case of serializable, application/x-java-serialized-object is used as content-type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **variableName** | **string**|  | 
 **optional** | ***HistoryApiGetHistoricTaskInstanceVariableDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiGetHistoricTaskInstanceVariableDataOpts struct
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

# **ListActivityInstances**
> DataResponseActivityInstanceResponse ListActivityInstances(ctx, optional)
List activity instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryApiListActivityInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiListActivityInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityId** | **optional.String**| An id of the activity instance. | 
 **activityInstanceId** | **optional.String**| An id of the activity instance. | 
 **activityName** | **optional.String**| The name of the activity instance. | 
 **activityType** | **optional.String**| The element type of the activity instance. | 
 **executionId** | **optional.String**| The execution id of the activity instance. | 
 **finished** | **optional.Bool**| Indication if the activity instance is finished. | 
 **taskAssignee** | **optional.String**| The assignee of the activity instance. | 
 **processInstanceId** | **optional.String**| The process instance id of the activity instance. | 
 **processDefinitionId** | **optional.String**| The process definition id of the activity instance. | 
 **tenantId** | **optional.String**| Only return instances with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return instances with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns instances without a tenantId set. If false, the withoutTenantId parameter is ignored. | 

### Return type

[**DataResponseActivityInstanceResponse**](DataResponseActivityInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistoricActivityInstances**
> DataResponseHistoricActivityInstanceResponse ListHistoricActivityInstances(ctx, optional)
List historic activity instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryApiListHistoricActivityInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiListHistoricActivityInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityId** | **optional.String**| An id of the activity instance. | 
 **activityInstanceId** | **optional.String**| An id of the historic activity instance. | 
 **activityName** | **optional.String**| The name of the historic activity instance. | 
 **activityType** | **optional.String**| The element type of the historic activity instance. | 
 **executionId** | **optional.String**| The execution id of the historic activity instance. | 
 **finished** | **optional.Bool**| Indication if the historic activity instance is finished. | 
 **taskAssignee** | **optional.String**| The assignee of the historic activity instance. | 
 **processInstanceId** | **optional.String**| The process instance id of the historic activity instance. | 
 **processDefinitionId** | **optional.String**| The process definition id of the historic activity instance. | 
 **tenantId** | **optional.String**| Only return instances with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return instances with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns instances without a tenantId set. If false, the withoutTenantId parameter is ignored. | 

### Return type

[**DataResponseHistoricActivityInstanceResponse**](DataResponseHistoricActivityInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistoricDetails**
> DataResponseHistoricDetailResponse ListHistoricDetails(ctx, optional)
Get historic detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryApiListHistoricDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiListHistoricDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| The id of the historic detail. | 
 **processInstanceId** | **optional.String**| The process instance id of the historic detail. | 
 **executionId** | **optional.String**| The execution id of the historic detail. | 
 **activityInstanceId** | **optional.String**| The activity instance id of the historic detail. | 
 **taskId** | **optional.String**| The task id of the historic detail. | 
 **selectOnlyFormProperties** | **optional.Bool**| Indication to only return form properties in the result. | 
 **selectOnlyVariableUpdates** | **optional.Bool**| Indication to only return variable updates in the result. | 

### Return type

[**DataResponseHistoricDetailResponse**](DataResponseHistoricDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistoricVariableInstances**
> DataResponseHistoricVariableInstanceResponse ListHistoricVariableInstances(ctx, optional)
List of historic variable instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryApiListHistoricVariableInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiListHistoricVariableInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processInstanceId** | **optional.String**| The process instance id of the historic variable instance. | 
 **taskId** | **optional.String**| The task id of the historic variable instance. | 
 **excludeTaskVariables** | **optional.Bool**| Indication to exclude the task variables from the result. | 
 **excludeLocalVariables** | **optional.Bool**| Indication to exclude local variables or not. | 
 **variableName** | **optional.String**| The variable name of the historic variable instance. | 
 **variableNameLike** | **optional.String**| The variable name using the like operator for the historic variable instance. | 

### Return type

[**DataResponseHistoricVariableInstanceResponse**](DataResponseHistoricVariableInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVariableInstances**
> DataResponseVariableInstanceResponse ListVariableInstances(ctx, optional)
List of variable instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryApiListVariableInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiListVariableInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processInstanceId** | **optional.String**| The process instance id of the variable instance. | 
 **taskId** | **optional.String**| The task id of the variable instance. | 
 **excludeTaskVariables** | **optional.Bool**| Indication to exclude the task variables from the result. | 
 **excludeLocalVariables** | **optional.Bool**| Indication to exclude local variables or not. | 
 **variableName** | **optional.String**| The variable name of the variable instance. | 
 **variableNameLike** | **optional.String**| The variable name using the like operator for the variable instance. | 

### Return type

[**DataResponseVariableInstanceResponse**](DataResponseVariableInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryActivityInstances**
> DataResponseHistoricActivityInstanceResponse QueryActivityInstances(ctx, optional)
Query for historic activity instances

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of historic task instances, but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryApiQueryActivityInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiQueryActivityInstancesOpts struct
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

# **QueryHistoricDetail**
> DataResponseHistoricDetailResponse QueryHistoricDetail(ctx, optional)
Query for historic details

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of historic process instances, but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryApiQueryHistoricDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiQueryHistoricDetailOpts struct
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

# **QueryVariableInstances**
> DataResponseHistoricVariableInstanceResponse QueryVariableInstances(ctx, optional)
Query for historic variable instances

All supported JSON parameter fields allowed are exactly the same as the parameters found for getting a collection of historic process instances, but passed in as JSON-body arguments rather than URL-parameters to allow for more advanced querying and preventing errors with request-uri’s that are too long. On top of that, the query allows for filtering based on process variables. The variables property is a JSON-array containing objects with the format as described here.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HistoryApiQueryVariableInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiQueryVariableInstancesOpts struct
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
 **optional** | ***HistoryApiQueryVariableInstances_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HistoryApiQueryVariableInstances_1Opts struct
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

