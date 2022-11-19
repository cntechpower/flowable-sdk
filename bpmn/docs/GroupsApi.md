# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](GroupsApi.md#CreateGroup) | **Post** /identity/groups | Create a group
[**CreateMembership**](GroupsApi.md#CreateMembership) | **Post** /identity/groups/{groupId}/members | Add a member to a group
[**DeleteGroup**](GroupsApi.md#DeleteGroup) | **Delete** /identity/groups/{groupId} | Delete a group
[**DeleteMembership**](GroupsApi.md#DeleteMembership) | **Delete** /identity/groups/{groupId}/members/{userId} | Delete a member from a group
[**GetGroup**](GroupsApi.md#GetGroup) | **Get** /identity/groups/{groupId} | Get a single group
[**ListGroups**](GroupsApi.md#ListGroups) | **Get** /identity/groups | List groups
[**UpdateGroup**](GroupsApi.md#UpdateGroup) | **Put** /identity/groups/{groupId} | Update a group

# **CreateGroup**
> GroupResponse CreateGroup(ctx, optional)
Create a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiCreateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiCreateGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupRequest**](GroupRequest.md)|  | 

### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMembership**
> MembershipResponse CreateMembership(ctx, groupId, optional)
Add a member to a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***GroupsApiCreateMembershipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiCreateMembershipOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MembershipRequest**](MembershipRequest.md)|  | 

### Return type

[**MembershipResponse**](MembershipResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> DeleteGroup(ctx, groupId)
Delete a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMembership**
> DeleteMembership(ctx, groupId, userId)
Delete a member from a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroup**
> GroupResponse GetGroup(ctx, groupId)
Get a single group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroups**
> DataResponseGroupResponse ListGroups(ctx, optional)
List groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiListGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiListGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return group with the given id | 
 **name** | **optional.String**| Only return groups with the given name | 
 **type_** | **optional.String**| Only return groups with the given type | 
 **nameLike** | **optional.String**| Only return groups with a name like the given value. Use % as wildcard-character. | 
 **member** | **optional.String**| Only return groups which have a member with the given username. | 
 **potentialStarter** | **optional.String**| Only return groups which members are potential starters for a process-definition with the given id. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseGroupResponse**](DataResponseGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> GroupResponse UpdateGroup(ctx, groupId, optional)
Update a group

All request values are optional. For example, you can only include the name attribute in the request body JSON-object, only updating the name of the group, leaving all other fields unaffected. When an attribute is explicitly included and is set to null, the group-value will be updated to null.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***GroupsApiUpdateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiUpdateGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GroupRequest**](GroupRequest.md)|  | 

### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

