# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchMigrateInstancesOfProcessDefinition**](ProcessDefinitionsApi.md#BatchMigrateInstancesOfProcessDefinition) | **Post** /repository/process-definitions/{processDefinitionId}/batch-migrate | Batch migrate all instances of process definition
[**CreateIdentityLink**](ProcessDefinitionsApi.md#CreateIdentityLink) | **Post** /repository/process-definitions/{processDefinitionId}/identitylinks | Add a candidate starter to a process definition
[**DeleteIdentityLink**](ProcessDefinitionsApi.md#DeleteIdentityLink) | **Delete** /repository/process-definitions/{processDefinitionId}/identitylinks/{family}/{identityId} | Delete a candidate starter from a process definition
[**ExecuteProcessDefinitionAction**](ProcessDefinitionsApi.md#ExecuteProcessDefinitionAction) | **Put** /repository/process-definitions/{processDefinitionId} | Execute actions for a process definition
[**GetBpmnModelResource**](ProcessDefinitionsApi.md#GetBpmnModelResource) | **Get** /repository/process-definitions/{processDefinitionId}/model | Get a process definition BPMN model
[**GetIdentityLink**](ProcessDefinitionsApi.md#GetIdentityLink) | **Get** /repository/process-definitions/{processDefinitionId}/identitylinks/{family}/{identityId} | Get a candidate starter from a process definition
[**GetModelResource**](ProcessDefinitionsApi.md#GetModelResource) | **Get** /repository/process-definitions/{processDefinitionId}/image | Get a process definition image
[**GetProcessDefinition**](ProcessDefinitionsApi.md#GetProcessDefinition) | **Get** /repository/process-definitions/{processDefinitionId} | Get a process definition
[**GetProcessDefinitionResource**](ProcessDefinitionsApi.md#GetProcessDefinitionResource) | **Get** /repository/process-definitions/{processDefinitionId}/resourcedata | Get a process definition resource content
[**GetProcessDefinitionStartForm**](ProcessDefinitionsApi.md#GetProcessDefinitionStartForm) | **Get** /repository/process-definitions/{processDefinitionId}/start-form | Get a process definition start form
[**ListProcessDefinitionDecisionTables**](ProcessDefinitionsApi.md#ListProcessDefinitionDecisionTables) | **Get** /repository/process-definitions/{processDefinitionId}/decision-tables | List decision tables for a process-definition
[**ListProcessDefinitionDecisions**](ProcessDefinitionsApi.md#ListProcessDefinitionDecisions) | **Get** /repository/process-definitions/{processDefinitionId}/decisions | List decisions for a process-definition
[**ListProcessDefinitionFormDefinitions**](ProcessDefinitionsApi.md#ListProcessDefinitionFormDefinitions) | **Get** /repository/process-definitions/{processDefinitionId}/form-definitions | List form definitions for a process-definition
[**ListProcessDefinitionIdentityLinks**](ProcessDefinitionsApi.md#ListProcessDefinitionIdentityLinks) | **Get** /repository/process-definitions/{processDefinitionId}/identitylinks | List candidate starters for a process-definition
[**ListProcessDefinitions**](ProcessDefinitionsApi.md#ListProcessDefinitions) | **Get** /repository/process-definitions | List of process definitions
[**MigrateInstancesOfProcessDefinition**](ProcessDefinitionsApi.md#MigrateInstancesOfProcessDefinition) | **Post** /repository/process-definitions/{processDefinitionId}/migrate | Migrate all instances of process definition

# **BatchMigrateInstancesOfProcessDefinition**
> BatchMigrateInstancesOfProcessDefinition(ctx, processDefinitionId, optional)
Batch migrate all instances of process definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 
 **optional** | ***ProcessDefinitionsApiBatchMigrateInstancesOfProcessDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessDefinitionsApiBatchMigrateInstancesOfProcessDefinitionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIdentityLink**
> RestIdentityLink CreateIdentityLink(ctx, processDefinitionId, optional)
Add a candidate starter to a process definition

It is possible to add either a user or a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 
 **optional** | ***ProcessDefinitionsApiCreateIdentityLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessDefinitionsApiCreateIdentityLinkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RestIdentityLink**](RestIdentityLink.md)|  | 

### Return type

[**RestIdentityLink**](RestIdentityLink.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIdentityLink**
> DeleteIdentityLink(ctx, processDefinitionId, family, identityId)
Delete a candidate starter from a process definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 
  **family** | **string**|  | 
  **identityId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteProcessDefinitionAction**
> ProcessDefinitionResponse ExecuteProcessDefinitionAction(ctx, body, processDefinitionId)
Execute actions for a process definition

Execute actions for a process definition (Update category, Suspend or Activate)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProcessDefinitionActionRequest**](ProcessDefinitionActionRequest.md)|  | 
  **processDefinitionId** | **string**|  | 

### Return type

[**ProcessDefinitionResponse**](ProcessDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBpmnModelResource**
> BpmnModel GetBpmnModelResource(ctx, processDefinitionId)
Get a process definition BPMN model

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 

### Return type

[**BpmnModel**](BpmnModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdentityLink**
> RestIdentityLink GetIdentityLink(ctx, processDefinitionId, family, identityId)
Get a candidate starter from a process definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 
  **family** | **string**|  | 
  **identityId** | **string**|  | 

### Return type

[**RestIdentityLink**](RestIdentityLink.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetModelResource**
> []string GetModelResource(ctx, processDefinitionId)
Get a process definition image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessDefinition**
> ProcessDefinitionResponse GetProcessDefinition(ctx, processDefinitionId)
Get a process definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 

### Return type

[**ProcessDefinitionResponse**](ProcessDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessDefinitionResource**
> []string GetProcessDefinitionResource(ctx, processDefinitionId)
Get a process definition resource content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessDefinitionStartForm**
> string GetProcessDefinitionStartForm(ctx, processDefinitionId)
Get a process definition start form

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProcessDefinitionDecisionTables**
> []DecisionResponse ListProcessDefinitionDecisionTables(ctx, processDefinitionId)
List decision tables for a process-definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 

### Return type

[**[]DecisionResponse**](DecisionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProcessDefinitionDecisions**
> []DecisionResponse ListProcessDefinitionDecisions(ctx, processDefinitionId)
List decisions for a process-definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 

### Return type

[**[]DecisionResponse**](DecisionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProcessDefinitionFormDefinitions**
> []FormDefinitionResponse ListProcessDefinitionFormDefinitions(ctx, processDefinitionId)
List form definitions for a process-definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 

### Return type

[**[]FormDefinitionResponse**](FormDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProcessDefinitionIdentityLinks**
> []RestIdentityLink ListProcessDefinitionIdentityLinks(ctx, processDefinitionId)
List candidate starters for a process-definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 

### Return type

[**[]RestIdentityLink**](RestIdentityLink.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProcessDefinitions**
> DataResponseProcessDefinitionResponse ListProcessDefinitions(ctx, optional)
List of process definitions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProcessDefinitionsApiListProcessDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessDefinitionsApiListProcessDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **optional.Int32**| Only return process definitions with the given version. | 
 **name** | **optional.String**| Only return process definitions with the given name. | 
 **nameLike** | **optional.String**| Only return process definitions with a name like the given name. | 
 **nameLikeIgnoreCase** | **optional.String**| Only return process definitions with a name like the given name ignoring case. | 
 **key** | **optional.String**| Only return process definitions with the given key. | 
 **keyLike** | **optional.String**| Only return process definitions with a name like the given key. | 
 **resourceName** | **optional.String**| Only return process definitions with the given resource name. | 
 **resourceNameLike** | **optional.String**| Only return process definitions with a name like the given resource name. | 
 **category** | **optional.String**| Only return process definitions with the given category. | 
 **categoryLike** | **optional.String**| Only return process definitions with a category like the given name. | 
 **categoryNotEquals** | **optional.String**| Only return process definitions which do not have the given category. | 
 **deploymentId** | **optional.String**| Only return process definitions which are part of a deployment with the given deployment id. | 
 **parentDeploymentId** | **optional.String**| Only return process definitions which are part of a deployment with the given parent deployment id. | 
 **startableByUser** | **optional.String**| Only return process definitions which are part of a deployment with the given id. | 
 **latest** | **optional.Bool**| Only return the latest process definition versions. Can only be used together with key and keyLike parameters, using any other parameter will result in a 400-response. | 
 **suspended** | **optional.Bool**| If true, only returns process definitions which are suspended. If false, only active process definitions (which are not suspended) are returned. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseProcessDefinitionResponse**](DataResponseProcessDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateInstancesOfProcessDefinition**
> MigrateInstancesOfProcessDefinition(ctx, processDefinitionId, optional)
Migrate all instances of process definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processDefinitionId** | **string**|  | 
 **optional** | ***ProcessDefinitionsApiMigrateInstancesOfProcessDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessDefinitionsApiMigrateInstancesOfProcessDefinitionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

