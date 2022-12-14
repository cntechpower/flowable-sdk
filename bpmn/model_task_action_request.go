/*
 * Flowable REST API
 *
 * # flowable / flowəb(ə)l /    - a compact and highly efficient workflow and Business Process Management (BPM) platform for developers, system admins and business users.  - a lightning fast, tried and tested BPMN 2 process engine written in Java. It is Apache 2.0 licensed open source, with a committed community.  - can run embedded in a Java application, or as a service on a server, a cluster, and in the cloud. It integrates perfectly with Spring. With a rich Java and REST API, it is the ideal engine for orchestrating human or system activities.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bpmn

type TaskActionRequest struct {
	// Action to perform: Either complete, claim, delegate or resolve
	Action string `json:"action"`
	// If action is claim or delegate, you can use this parameter to set the assignee associated 
	Assignee string `json:"assignee,omitempty"`
	// Required when completing a task with a form
	FormDefinitionId string `json:"formDefinitionId,omitempty"`
	// Optional outcome value when completing a task with a form
	Outcome string `json:"outcome,omitempty"`
	// If action is complete, you can use this parameter to set variables 
	Variables []RestVariable `json:"variables,omitempty"`
	// If action is complete, you can use this parameter to set transient variables 
	TransientVariables []RestVariable `json:"transientVariables,omitempty"`
}
