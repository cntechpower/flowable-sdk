# {{classname}}

All URIs are relative to */flowable-rest/service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTable**](DatabaseTablesApi.md#GetTable) | **Get** /management/tables/{tableName} | Get a single table
[**GetTableData**](DatabaseTablesApi.md#GetTableData) | **Get** /management/tables/{tableName}/data | Get row data for a single table
[**GetTableMetaData**](DatabaseTablesApi.md#GetTableMetaData) | **Get** /management/tables/{tableName}/columns | Get column info for a single table
[**ListTables**](DatabaseTablesApi.md#ListTables) | **Get** /management/tables |  List tables

# **GetTable**
> TableResponse GetTable(ctx, tableName)
Get a single table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableName** | **string**|  | 

### Return type

[**TableResponse**](TableResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTableData**
> DataResponseListMapStringObject GetTableData(ctx, tableName, optional)
Get row data for a single table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableName** | **string**|  | 
 **optional** | ***DatabaseTablesApiGetTableDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DatabaseTablesApiGetTableDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| Index of the first row to fetch. Defaults to 0. | 
 **size** | **optional.Int32**| Number of rows to fetch, starting from start. Defaults to 10. | 
 **orderAscendingColumn** | **optional.String**| Name of the column to sort the resulting rows on, ascending. | 
 **orderDescendingColumn** | **optional.String**| Name of the column to sort the resulting rows on, descending. | 

### Return type

[**DataResponseListMapStringObject**](DataResponseListMapStringObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTableMetaData**
> TableMetaData GetTableMetaData(ctx, tableName)
Get column info for a single table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tableName** | **string**|  | 

### Return type

[**TableMetaData**](TableMetaData.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTables**
> []TableResponse ListTables(ctx, )
 List tables

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TableResponse**](TableResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

