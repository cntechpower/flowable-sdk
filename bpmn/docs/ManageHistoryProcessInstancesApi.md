# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteHistoricProcessInstances**](ManageHistoryProcessInstancesApi.md#BulkDeleteHistoricProcessInstances) | **Post** /history/historic-process-instances/delete | Post action request to delete a bulk of historic process instances

# **BulkDeleteHistoricProcessInstances**
> BulkDeleteHistoricProcessInstances(ctx, optional)
Post action request to delete a bulk of historic process instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManageHistoryProcessInstancesApiBulkDeleteHistoricProcessInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManageHistoryProcessInstancesApiBulkDeleteHistoricProcessInstancesOpts struct
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

