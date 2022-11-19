# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateModel**](ModelsApi.md#CreateModel) | **Post** /repository/models | Create a model
[**DeleteModel**](ModelsApi.md#DeleteModel) | **Delete** /repository/models/{modelId} | Delete a model
[**GetExtraEditorSource**](ModelsApi.md#GetExtraEditorSource) | **Get** /repository/models/{modelId}/source-extra | Get the extra editor source for a model
[**GetModel**](ModelsApi.md#GetModel) | **Get** /repository/models/{modelId} | Get a model
[**GetModelBytes**](ModelsApi.md#GetModelBytes) | **Get** /repository/models/{modelId}/source | Get the editor source for a model
[**ListModels**](ModelsApi.md#ListModels) | **Get** /repository/models | List models
[**SetExtraEditorSource**](ModelsApi.md#SetExtraEditorSource) | **Put** /repository/models/{modelId}/source-extra | Set the extra editor source for a model
[**SetModelSource**](ModelsApi.md#SetModelSource) | **Put** /repository/models/{modelId}/source | Set the editor source for a model
[**UpdateModel**](ModelsApi.md#UpdateModel) | **Put** /repository/models/{modelId} | Update a model

# **CreateModel**
> ModelResponse CreateModel(ctx, optional)
Create a model

All request values are optional. For example, you can only include the name attribute in the request body JSON-object, only setting the name of the model, leaving all other fields null.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ModelsApiCreateModelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModelsApiCreateModelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ModelRequest**](ModelRequest.md)|  | 

### Return type

[**ModelResponse**](ModelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteModel**
> DeleteModel(ctx, modelId)
Delete a model

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modelId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtraEditorSource**
> []string GetExtraEditorSource(ctx, modelId)
Get the extra editor source for a model

Response body contains the model’s raw editor source. The response’s content-type is set to application/octet-stream, regardless of the content of the source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modelId** | **string**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetModel**
> ModelResponse GetModel(ctx, modelId)
Get a model

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modelId** | **string**|  | 

### Return type

[**ModelResponse**](ModelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetModelBytes**
> []string GetModelBytes(ctx, modelId)
Get the editor source for a model

Response body contains the model’s raw editor source. The response’s content-type is set to application/octet-stream, regardless of the content of the source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modelId** | **string**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListModels**
> DataResponseModelResponse ListModels(ctx, optional)
List models

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ModelsApiListModelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModelsApiListModelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return models with the given version. | 
 **category** | **optional.String**| Only return models with the given category. | 
 **categoryLike** | **optional.String**| Only return models with a category like the given name. | 
 **categoryNotEquals** | **optional.String**| Only return models which do not have the given category. | 
 **name** | **optional.String**| Only return models with the given name. | 
 **nameLike** | **optional.String**| Only return models with a name like the given name. | 
 **key** | **optional.String**| Only return models with the given key. | 
 **deploymentId** | **optional.String**| Only return models with the given category. | 
 **version** | **optional.Int32**| Only return models with the given version. | 
 **latestVersion** | **optional.Bool**| If true, only return models which are the latest version. Best used in combination with key. If false is passed in as value, this is ignored and all versions are returned. | 
 **deployed** | **optional.Bool**| If true, only deployed models are returned. If false, only undeployed models are returned (deploymentId is null). | 
 **tenantId** | **optional.String**| Only return models with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return models with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns models without a tenantId set. If false, the withoutTenantId parameter is ignored. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseModelResponse**](DataResponseModelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetExtraEditorSource**
> SetExtraEditorSource(ctx, file, modelId)
Set the extra editor source for a model

Response body contains the model’s raw editor source. The response’s content-type is set to application/octet-stream, regardless of the content of the source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 
  **modelId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetModelSource**
> SetModelSource(ctx, file, modelId)
Set the editor source for a model

Response body contains the model’s raw editor source. The response’s content-type is set to application/octet-stream, regardless of the content of the source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 
  **modelId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateModel**
> ModelResponse UpdateModel(ctx, modelId, optional)
Update a model

All request values are optional. For example, you can only include the name attribute in the request body JSON-object, only updating the name of the model, leaving all other fields unaffected. When an attribute is explicitly included and is set to null, the model-value will be updated to null. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modelId** | **string**|  | 
 **optional** | ***ModelsApiUpdateModelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModelsApiUpdateModelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ModelRequest**](ModelRequest.md)|  | 

### Return type

[**ModelResponse**](ModelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

