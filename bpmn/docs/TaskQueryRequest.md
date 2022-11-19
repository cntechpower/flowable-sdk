# TaskQueryRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **int32** |  | [optional] [default to null]
**Size** | **int32** |  | [optional] [default to null]
**Sort** | **string** |  | [optional] [default to null]
**Order** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**NameLike** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**DescriptionLike** | **string** |  | [optional] [default to null]
**Priority** | **int32** |  | [optional] [default to null]
**MinimumPriority** | **int32** |  | [optional] [default to null]
**MaximumPriority** | **int32** |  | [optional] [default to null]
**Assignee** | **string** |  | [optional] [default to null]
**AssigneeLike** | **string** |  | [optional] [default to null]
**Owner** | **string** |  | [optional] [default to null]
**OwnerLike** | **string** |  | [optional] [default to null]
**Unassigned** | **bool** |  | [optional] [default to null]
**DelegationState** | **string** |  | [optional] [default to null]
**CandidateUser** | **string** |  | [optional] [default to null]
**CandidateGroup** | **string** |  | [optional] [default to null]
**CandidateGroupIn** | **[]string** |  | [optional] [default to null]
**IgnoreAssignee** | **bool** |  | [optional] [default to null]
**InvolvedUser** | **string** |  | [optional] [default to null]
**ProcessInstanceId** | **string** |  | [optional] [default to null]
**ProcessInstanceIdWithChildren** | **string** |  | [optional] [default to null]
**WithoutProcessInstanceId** | **bool** |  | [optional] [default to null]
**ProcessInstanceBusinessKey** | **string** |  | [optional] [default to null]
**ProcessInstanceBusinessKeyLike** | **string** |  | [optional] [default to null]
**ProcessDefinitionId** | **string** |  | [optional] [default to null]
**ProcessDefinitionKey** | **string** |  | [optional] [default to null]
**ProcessDefinitionName** | **string** |  | [optional] [default to null]
**ProcessDefinitionKeyLike** | **string** |  | [optional] [default to null]
**ProcessDefinitionNameLike** | **string** |  | [optional] [default to null]
**ExecutionId** | **string** |  | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**CreatedBefore** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**CreatedAfter** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**ExcludeSubTasks** | **bool** |  | [optional] [default to null]
**TaskDefinitionKey** | **string** |  | [optional] [default to null]
**TaskDefinitionKeyLike** | **string** |  | [optional] [default to null]
**TaskDefinitionKeys** | **[]string** |  | [optional] [default to null]
**DueDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**DueBefore** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**DueAfter** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**WithoutDueDate** | **bool** |  | [optional] [default to null]
**Active** | **bool** |  | [optional] [default to null]
**IncludeTaskLocalVariables** | **bool** |  | [optional] [default to null]
**IncludeProcessVariables** | **bool** |  | [optional] [default to null]
**ScopeDefinitionId** | **string** |  | [optional] [default to null]
**ScopeId** | **string** |  | [optional] [default to null]
**WithoutScopeId** | **bool** |  | [optional] [default to null]
**ScopeType** | **string** |  | [optional] [default to null]
**PropagatedStageInstanceId** | **string** |  | [optional] [default to null]
**TenantId** | **string** |  | [optional] [default to null]
**TenantIdLike** | **string** |  | [optional] [default to null]
**WithoutTenantId** | **bool** |  | [optional] [default to null]
**CandidateOrAssigned** | **string** |  | [optional] [default to null]
**Category** | **string** |  | [optional] [default to null]
**CategoryIn** | **[]string** |  | [optional] [default to null]
**CategoryNotIn** | **[]string** |  | [optional] [default to null]
**WithoutCategory** | **bool** |  | [optional] [default to null]
**TaskVariables** | [**[]QueryVariable**](QueryVariable.md) |  | [optional] [default to null]
**ProcessInstanceVariables** | [**[]QueryVariable**](QueryVariable.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
