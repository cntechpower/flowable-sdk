# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /identity/users | Create a user
[**CreateUserInfo**](UsersApi.md#CreateUserInfo) | **Post** /identity/users/{userId}/info | Create a new user’s info entry
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /identity/users/{userId} | Delete a user
[**DeleteUserInfo**](UsersApi.md#DeleteUserInfo) | **Delete** /identity/users/{userId}/info/{key} | Delete a user’s info
[**GetUser**](UsersApi.md#GetUser) | **Get** /identity/users/{userId} | Get a single user
[**GetUserInfo**](UsersApi.md#GetUserInfo) | **Get** /identity/users/{userId}/info/{key} | Get a user’s info
[**GetUserPicture**](UsersApi.md#GetUserPicture) | **Get** /identity/users/{userId}/picture | Get a user’s picture
[**ListUserInfo**](UsersApi.md#ListUserInfo) | **Get** /identity/users/{userId}/info | List user’s info
[**ListUsers**](UsersApi.md#ListUsers) | **Get** /identity/users | List users
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /identity/users/{userId} | Update a user
[**UpdateUserInfo**](UsersApi.md#UpdateUserInfo) | **Put** /identity/users/{userId}/info/{key} | Update a user’s info
[**UpdateUserPicture**](UsersApi.md#UpdateUserPicture) | **Put** /identity/users/{userId}/picture | Updating a user’s picture

# **CreateUser**
> UserResponse CreateUser(ctx, optional)
Create a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiCreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiCreateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UserRequest**](UserRequest.md)|  | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUserInfo**
> UserInfoResponse CreateUserInfo(ctx, userId, optional)
Create a new user’s info entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UsersApiCreateUserInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiCreateUserInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserInfoRequest**](UserInfoRequest.md)|  | 

### Return type

[**UserInfoResponse**](UserInfoResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, userId)
Delete a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserInfo**
> DeleteUserInfo(ctx, userId, key)
Delete a user’s info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **key** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> UserResponse GetUser(ctx, userId)
Get a single user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserInfo**
> UserInfoResponse GetUserInfo(ctx, userId, key)
Get a user’s info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **key** | **string**|  | 

### Return type

[**UserInfoResponse**](UserInfoResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserPicture**
> []string GetUserPicture(ctx, userId)
Get a user’s picture

The response body contains the raw picture data, representing the user’s picture. The Content-type of the response corresponds to the mimeType that was set when creating the picture.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserInfo**
> []UserInfoResponse ListUserInfo(ctx, userId)
List user’s info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**[]UserInfoResponse**](UserInfoResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsers**
> DataResponseUserResponse ListUsers(ctx, optional)
List users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiListUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiListUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Only return group with the given id | 
 **firstName** | **optional.String**| Only return users with the given firstname | 
 **lastName** | **optional.String**| Only return users with the given lastname | 
 **displayName** | **optional.String**| Only return users with the given displayName | 
 **email** | **optional.String**| Only return users with the given email | 
 **firstNameLike** | **optional.String**| Only return userswith a firstname like the given value. Use % as wildcard-character. | 
 **lastNameLike** | **optional.String**| Only return users with a lastname like the given value. Use % as wildcard-character. | 
 **displayNameLike** | **optional.String**| Only return users with a displayName like the given value. Use % as wildcard-character. | 
 **emailLike** | **optional.String**| Only return users with an email like the given value. Use % as wildcard-character. | 
 **memberOfGroup** | **optional.String**| Only return users which are a member of the given group. | 
 **tenantId** | **optional.String**| Only return users which are a members of the given tenant. | 
 **potentialStarter** | **optional.String**| Only return users  which members are potential starters for a process-definition with the given id. | 
 **sort** | **optional.String**| Property to sort on, to be used together with the order. | 

### Return type

[**DataResponseUserResponse**](DataResponseUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UserResponse UpdateUser(ctx, userId, optional)
Update a user

All request values are optional. For example, you can only include the firstName attribute in the request body JSON-object, only updating the firstName of the user, leaving all other fields unaffected. When an attribute is explicitly included and is set to null, the user-value will be updated to null. Example: {\"firstName\" : null} will clear the firstName of the user).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UsersApiUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUpdateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserRequest**](UserRequest.md)|  | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserInfo**
> UserInfoResponse UpdateUserInfo(ctx, userId, key, optional)
Update a user’s info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **key** | **string**|  | 
 **optional** | ***UsersApiUpdateUserInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUpdateUserInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UserInfoRequest**](UserInfoRequest.md)|  | 

### Return type

[**UserInfoResponse**](UserInfoResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserPicture**
> UpdateUserPicture(ctx, file, userId)
Updating a user’s picture

The request should be of type multipart/form-data. There should be a single file-part included with the binary value of the picture. On top of that, the following additional form-fields can be present:  mimeType: Optional mime-type for the uploaded picture. If omitted, the default of image/jpeg is used as a mime-type for the picture.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

