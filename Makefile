# https://documentation.flowable.com/latest/assets/core-swagger/bpmn.html
bpmn_api:
	swagger-codegen generate -i https://developer-docs.flowable.com/swagger-docs/3.12.0/process/flowable-swagger-process.yaml \
	-l go -DpackageName=bpmn -o bpmn/