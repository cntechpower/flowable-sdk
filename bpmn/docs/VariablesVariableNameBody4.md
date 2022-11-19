# VariablesVariableNameBody4

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Required name of the variable | [optional] [default to null]
**Type_** | **string** | Type of variable that is updated. If omitted, reverts to raw JSON-value type (string, boolean, integer or double) | [optional] [default to null]
**Scope** | **string** | Scope of variable to be returned. When local, only task-local variable value is returned. When global, only variable value from the task’s parent execution-hierarchy are returned. When the parameter is omitted, a local variable will be returned if it exists, otherwise a global variable.. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

