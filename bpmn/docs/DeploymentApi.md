# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeployment**](DeploymentApi.md#DeleteDeployment) | **Delete** /repository/deployments/{deploymentId} | Delete a deployment
[**GetDeployment**](DeploymentApi.md#GetDeployment) | **Get** /repository/deployments/{deploymentId} | Get a deployment
[**GetDeploymentResource**](DeploymentApi.md#GetDeploymentResource) | **Get** /repository/deployments/{deploymentId}/resources/** | Get a deployment resource
[**GetDeploymentResourceData**](DeploymentApi.md#GetDeploymentResourceData) | **Get** /repository/deployments/{deploymentId}/resourcedata/{resourceName} | Get a deployment resource content
[**ListDeploymentResources**](DeploymentApi.md#ListDeploymentResources) | **Get** /repository/deployments/{deploymentId}/resources | List resources in a deployment
[**ListDeployments**](DeploymentApi.md#ListDeployments) | **Get** /repository/deployments | List Deployments
[**UploadDeployment**](DeploymentApi.md#UploadDeployment) | **Post** /repository/deployments | Create a new deployment

# **DeleteDeployment**
> DeleteDeployment(ctx, deploymentId, optional)
Delete a deployment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentId** | **string**|  | 
 **optional** | ***DeploymentApiDeleteDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentApiDeleteDeploymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployment**
> DeploymentResponse GetDeployment(ctx, deploymentId)
Get a deployment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentId** | **string**| The id of the deployment to get. | 

### Return type

[**DeploymentResponse**](DeploymentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentResource**
> DeploymentResourceResponse GetDeploymentResource(ctx, deploymentId)
Get a deployment resource

Replace ** by ResourceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentId** | **string**|  | 

### Return type

[**DeploymentResourceResponse**](DeploymentResourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentResourceData**
> []string GetDeploymentResourceData(ctx, deploymentId, resourceName)
Get a deployment resource content

The response body will contain the binary resource-content for the requested resource. The response content-type will be the same as the type returned in the resources mimeType property. Also, a content-disposition header is set, allowing browsers to download the file instead of displaying it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentId** | **string**|  | 
  **resourceName** | **string**| The name of the resource to get. Make sure you URL-encode the resourceName in case it contains forward slashes. Eg: use diagrams%2Fmy-process.bpmn20.xml instead of diagrams/my-process.bpmn20.xml. | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeploymentResources**
> []DeploymentResourceResponse ListDeploymentResources(ctx, deploymentId)
List resources in a deployment

The dataUrl property in the resulting JSON for a single resource contains the actual URL to use for retrieving the binary resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentId** | **string**|  | 

### Return type

[**[]DeploymentResourceResponse**](DeploymentResourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeployments**
> DataResponseDeploymentResponse ListDeployments(ctx, optional)
List Deployments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeploymentApiListDeploymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentApiListDeploymentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Only return deployments with the given name. | 
 **nameLike** | **optional.String**| Only return deployments with a name like the given name. | 
 **category** | **optional.String**| Only return deployments with the given category. | 
 **categoryNotEquals** | **optional.String**| Only return deployments which do not have the given category. | 
 **parentDeploymentId** | **optional.String**| Only return deployments with the given parent deployment id. | 
 **parentDeploymentIdLike** | **optional.String**| Only return deployments with a parent deployment id like the given value. | 
 **tenantId** | **optional.String**| Only return deployments with the given tenantId. | 
 **tenantIdLike** | **optional.String**| Only return deployments with a tenantId like the given value. | 
 **withoutTenantId** | **optional.Bool**| If true, only returns deployments without a tenantId set. If false, the withoutTenantId parameter is ignored. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseDeploymentResponse**](DataResponseDeploymentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadDeployment**
> DeploymentResponse UploadDeployment(ctx, file, optional)
Create a new deployment

The request body should contain data of type multipart/form-data. There should be exactly one file in the request, any additional files will be ignored. The deployment name is the name of the file-field passed in. If multiple resources need to be deployed in a single deployment, compress the resources in a zip and make sure the file-name ends with .bar or .zip.  An additional parameter (form-field) can be passed in the request body with name tenantId. The value of this field will be used as the id of the tenant this deployment is done in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 
 **optional** | ***DeploymentApiUploadDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentApiUploadDeploymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentKey** | **optional.**|  | 
 **deploymentName** | **optional.**|  | 
 **tenantId** | **optional.**|  | 

### Return type

[**DeploymentResponse**](DeploymentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

