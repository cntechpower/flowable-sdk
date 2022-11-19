package flowable_sdk

import (
	"context"
	"github.com/antihax/optional"
	"github.com/cntechpower/flowable-sdk/bpmn"
	"testing"
)

func TestBPMN(t *testing.T) {
	ctx := context.Background()
	cfg := bpmn.NewConfiguration()
	cfg.Host = "127.0.0.1"
	cli := bpmn.NewAPIClient(cfg)

	cli.ProcessInstancesApi.CreateProcessInstance(ctx, &bpmn.ProcessInstancesApiCreateProcessInstanceOpts{Body: optional.NewInterface(bpmn.ProcessInstanceCreateRequest{
		ProcessDefinitionId:        "",
		ProcessDefinitionKey:       "",
		Message:                    "",
		Name:                       "",
		BusinessKey:                "",
		Variables:                  nil,
		TransientVariables:         nil,
		StartFormVariables:         nil,
		Outcome:                    "",
		TenantId:                   "",
		OverrideDefinitionTenantId: "",
		ReturnVariables:            false,
	})})
}
