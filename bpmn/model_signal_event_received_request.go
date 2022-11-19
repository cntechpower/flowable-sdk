/*
 * Flowable REST API
 *
 * # flowable / flowəb(ə)l /    - a compact and highly efficient workflow and Business Process Management (BPM) platform for developers, system admins and business users.  - a lightning fast, tried and tested BPMN 2 process engine written in Java. It is Apache 2.0 licensed open source, with a committed community.  - can run embedded in a Java application, or as a service on a server, a cluster, and in the cloud. It integrates perfectly with Spring. With a rich Java and REST API, it is the ideal engine for orchestrating human or system activities.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bpmn

type SignalEventReceivedRequest struct {
	// Name of the signal
	SignalName string `json:"signalName,omitempty"`
	// Array of variables (in the general variables format) to use as payload to pass along with the signal. Cannot be used in case async is set to true, this will result in an error.
	Variables []RestVariable `json:"variables,omitempty"`
	// ID of the tenant that the signal event should be processed in
	TenantId string `json:"tenantId,omitempty"`
	// If true, handling of the signal will happen asynchronously. Return code will be 202 - Accepted to indicate the request is accepted but not yet executed. If false,                     handling the signal will be done immediately and result (200 - OK) will only return after this completed successfully. Defaults to false if omitted.
	Async bool `json:"async,omitempty"`
}
