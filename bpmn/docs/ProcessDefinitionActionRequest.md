# ProcessDefinitionActionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action to perform: Either activate or suspend | [default to null]
**IncludeProcessInstances** | **bool** | Whether or not to suspend/activate running process-instances for this process-definition. If omitted, the process-instances are left in the state they are | [optional] [default to null]
**Date** | [**time.Time**](time.Time.md) | Date (ISO-8601) when the suspension/activation should be executed. If omitted, the suspend/activation is effective immediately. | [optional] [default to null]
**Category** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

