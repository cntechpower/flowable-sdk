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

type JobResponse struct {
	Id string `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	CorrelationId string `json:"correlationId,omitempty"`
	ProcessInstanceId string `json:"processInstanceId,omitempty"`
	ProcessInstanceUrl string `json:"processInstanceUrl,omitempty"`
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`
	ProcessDefinitionUrl string `json:"processDefinitionUrl,omitempty"`
	ExecutionId string `json:"executionId,omitempty"`
	ExecutionUrl string `json:"executionUrl,omitempty"`
	ElementId string `json:"elementId,omitempty"`
	ElementName string `json:"elementName,omitempty"`
	Retries int32 `json:"retries,omitempty"`
	ExceptionMessage string `json:"exceptionMessage,omitempty"`
	DueDate time.Time `json:"dueDate,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
}