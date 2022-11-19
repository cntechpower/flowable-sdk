# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFormData**](FormsApi.md#GetFormData) | **Get** /form/form-data | Get form data
[**SubmitForm**](FormsApi.md#SubmitForm) | **Post** /form/form-data | Submit task form data

# **GetFormData**
> FormDataResponse GetFormData(ctx, optional)
Get form data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FormsApiGetFormDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FormsApiGetFormDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **optional.String**|  | 
 **processDefinitionId** | **optional.String**|  | 

### Return type

[**FormDataResponse**](FormDataResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitForm**
> ProcessInstanceResponse SubmitForm(ctx, optional)
Submit task form data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FormsApiSubmitFormOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FormsApiSubmitFormOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SubmitFormRequest**](SubmitFormRequest.md)|  | 

### Return type

[**ProcessInstanceResponse**](ProcessInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

