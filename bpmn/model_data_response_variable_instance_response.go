/*
 * Flowable REST API
 *
 * # flowable / flowəb(ə)l /    - a compact and highly efficient workflow and Business Process Management (BPM) platform for developers, system admins and business users.  - a lightning fast, tried and tested BPMN 2 process engine written in Java. It is Apache 2.0 licensed open source, with a committed community.  - can run embedded in a Java application, or as a service on a server, a cluster, and in the cloud. It integrates perfectly with Spring. With a rich Java and REST API, it is the ideal engine for orchestrating human or system activities.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bpmn

type DataResponseVariableInstanceResponse struct {
	Data []VariableInstanceResponse `json:"data,omitempty"`
	Total int64 `json:"total,omitempty"`
	Start int32 `json:"start,omitempty"`
	Sort string `json:"sort,omitempty"`
	Order string `json:"order,omitempty"`
	Size int32 `json:"size,omitempty"`
}
