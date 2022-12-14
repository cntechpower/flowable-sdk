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

type HistoricTaskLogEntryResponse struct {
	LogNumber int64 `json:"logNumber,omitempty"`
	Type_ string `json:"type,omitempty"`
	TaskId string `json:"taskId,omitempty"`
	TimeStamp time.Time `json:"timeStamp,omitempty"`
	UserId string `json:"userId,omitempty"`
	Data string `json:"data,omitempty"`
	ExecutionId string `json:"executionId,omitempty"`
	ProcessInstanceId string `json:"processInstanceId,omitempty"`
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`
	ScopeId string `json:"scopeId,omitempty"`
	ScopeDefinitionId string `json:"scopeDefinitionId,omitempty"`
	SubScopeId string `json:"subScopeId,omitempty"`
	ScopeType string `json:"scopeType,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
}
