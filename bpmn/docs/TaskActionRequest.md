# TaskActionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action to perform: Either complete, claim, delegate or resolve | [default to null]
**Assignee** | **string** | If action is claim or delegate, you can use this parameter to set the assignee associated  | [optional] [default to null]
**FormDefinitionId** | **string** | Required when completing a task with a form | [optional] [default to null]
**Outcome** | **string** | Optional outcome value when completing a task with a form | [optional] [default to null]
**Variables** | [**[]RestVariable**](RestVariable.md) | If action is complete, you can use this parameter to set variables  | [optional] [default to null]
**TransientVariables** | [**[]RestVariable**](RestVariable.md) | If action is complete, you can use this parameter to set transient variables  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

