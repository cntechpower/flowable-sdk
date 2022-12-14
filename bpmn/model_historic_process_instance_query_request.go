/*
 * Flowable REST API
 *
 * # flowable / flowəb(ə)l /    - a compact and highly efficient workflow and Business Process Management (BPM) platform for developers, system admins and business users.  - a lightning fast, tried and tested BPMN 2 process engine written in Java. It is Apache 2.0 licensed open source, with a committed community.  - can run embedded in a Java application, or as a service on a server, a cluster, and in the cloud. It integrates perfectly with Spring. With a rich Java and REST API, it is the ideal engine for orchestrating human or system activities.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bpmn
import (
	"time"
)

type HistoricProcessInstanceQueryRequest struct {
	Start int32 `json:"start,omitempty"`
	Size int32 `json:"size,omitempty"`
	Sort string `json:"sort,omitempty"`
	Order string `json:"order,omitempty"`
	ProcessInstanceId string `json:"processInstanceId,omitempty"`
	ProcessInstanceIds []string `json:"processInstanceIds,omitempty"`
	ProcessInstanceName string `json:"processInstanceName,omitempty"`
	ProcessInstanceNameLike string `json:"processInstanceNameLike,omitempty"`
	ProcessInstanceNameLikeIgnoreCase string `json:"processInstanceNameLikeIgnoreCase,omitempty"`
	ProcessBusinessKey string `json:"processBusinessKey,omitempty"`
	ProcessBusinessKeyLike string `json:"processBusinessKeyLike,omitempty"`
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`
	ProcessDefinitionKey string `json:"processDefinitionKey,omitempty"`
	ProcessDefinitionKeyIn []string `json:"processDefinitionKeyIn,omitempty"`
	ProcessDefinitionKeyNotIn []string `json:"processDefinitionKeyNotIn,omitempty"`
	ProcessDefinitionName string `json:"processDefinitionName,omitempty"`
	ProcessDefinitionVersion int32 `json:"processDefinitionVersion,omitempty"`
	ProcessDefinitionCategory string `json:"processDefinitionCategory,omitempty"`
	DeploymentId string `json:"deploymentId,omitempty"`
	DeploymentIdIn []string `json:"deploymentIdIn,omitempty"`
	SuperProcessInstanceId string `json:"superProcessInstanceId,omitempty"`
	ExcludeSubprocesses bool `json:"excludeSubprocesses,omitempty"`
	Finished bool `json:"finished,omitempty"`
	ActiveActivityId string `json:"activeActivityId,omitempty"`
	ActiveActivityIds []string `json:"activeActivityIds,omitempty"`
	InvolvedUser string `json:"involvedUser,omitempty"`
	FinishedAfter time.Time `json:"finishedAfter,omitempty"`
	FinishedBefore time.Time `json:"finishedBefore,omitempty"`
	StartedAfter time.Time `json:"startedAfter,omitempty"`
	StartedBefore time.Time `json:"startedBefore,omitempty"`
	StartedBy string `json:"startedBy,omitempty"`
	IncludeProcessVariables bool `json:"includeProcessVariables,omitempty"`
	Variables []QueryVariable `json:"variables,omitempty"`
	CallbackId string `json:"callbackId,omitempty"`
	CallbackType string `json:"callbackType,omitempty"`
	WithoutCallbackId bool `json:"withoutCallbackId,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
	TenantIdLike string `json:"tenantIdLike,omitempty"`
	WithoutTenantId bool `json:"withoutTenantId,omitempty"`
}
