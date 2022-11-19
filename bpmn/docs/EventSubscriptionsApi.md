# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventSubscription**](EventSubscriptionsApi.md#GetEventSubscription) | **Get** /runtime/event-subscriptions/{eventSubscriptionId} | Get a single event subscription
[**ListEventSubscriptions**](EventSubscriptionsApi.md#ListEventSubscriptions) | **Get** /runtime/event-subscriptions | List of event subscriptions

# **GetEventSubscription**
> EventSubscriptionResponse GetEventSubscription(ctx, eventSubscriptionId)
Get a single event subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventSubscriptionId** | **string**|  | 

### Return type

[**EventSubscriptionResponse**](EventSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEventSubscriptions**
> DataResponseEventSubscriptionResponse ListEventSubscriptions(ctx, optional)
List of event subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventSubscriptionsApiListEventSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventSubscriptionsApiListEventSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return event subscriptions with the given id | 
 **eventType** | **optional.String**| Only return event subscriptions with the given event type | 
 **eventName** | **optional.String**| Only return event subscriptions with the given event name | 
 **activityId** | **optional.String**| Only return event subscriptions with the given activity id | 
 **executionId** | **optional.String**| Only return event subscriptions with the given execution id | 
 **processInstanceId** | **optional.String**| Only return event subscriptions part of a process with the given id | 
 **withoutProcessInstanceId** | **optional.Bool**| Only return event subscriptions that have no process instance id | 
 **processDefinitionId** | **optional.String**| Only return event subscriptions with the given process definition id | 
 **withoutProcessDefinitionId** | **optional.Bool**| Only return event subscriptions that have no process definition id | 
 **scopeId** | **optional.String**| Only return event subscriptions part of a scope with the given id | 
 **subScopeId** | **optional.String**| Only return event subscriptions part of a sub scope with the given id | 
 **withoutScopeId** | **optional.Bool**| Only return event subscriptions that have no scope id | 
 **scopeDefinitionId** | **optional.String**| Only return event subscriptions with the given scope definition id | 
 **withoutScopeDefinitionId** | **optional.Bool**| Only return event subscriptions that have no scope definition id | 
 **createdBefore** | **optional.Time**| Only return event subscriptions which are created before the given date. | 
 **createdAfter** | **optional.Time**| Only return event subscriptions which are created after the given date. | 
 **tenantId** | **optional.String**| Only return event subscriptions with the given tenant id. | 
 **withoutTenantId** | **optional.Bool**| Only return event subscriptions that have no tenant id | 
 **configuration** | **optional.String**| Only return event subscriptions with the given configuration value. | 
 **withoutConfiguration** | **optional.Bool**| Only return event subscriptions that have no configuration value | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseEventSubscriptionResponse**](DataResponseEventSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

