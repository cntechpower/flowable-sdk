# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEngineInfo**](EngineApi.md#GetEngineInfo) | **Get** /management/engine | Get engine info
[**GetProperties**](EngineApi.md#GetProperties) | **Get** /management/properties | List engine properties

# **GetEngineInfo**
> EngineInfoResponse GetEngineInfo(ctx, )
Get engine info

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EngineInfoResponse**](EngineInfoResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProperties**
> map[string]string GetProperties(ctx, )
List engine properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

**map[string]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

