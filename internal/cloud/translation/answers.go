package translation

import (
	operatorservice "go.temporal.io/api/operatorservice/v1"
	workflowservice "go.temporal.io/api/workflowservice/v1"

	"github.com/temporalio/temporal-proxy/internal/services"
)

const (
	// getClusterInfoMethod is the WorkflowService method that describes the
	// Temporal Service. Cloud frontends refuse it, since it carries no namespace to
	// match the endpoint's, and no Cloud API reports what it returns.
	getClusterInfoMethod = "/" + services.WorkflowService + "/GetClusterInfo"

	// listSearchAttributesMethod is the OperatorService method that lists a
	// namespace's search attributes. Cloud frontends refuse every OperatorService
	// call from a customer credential, whatever its role.
	listSearchAttributesMethod = "/" + services.OperatorService + "/ListSearchAttributes"
)

// getClusterInfo answers WorkflowService.GetClusterInfo with an empty reply. An
// empty reply is what the Temporal UI itself uses when it knows it is talking to
// Cloud, so a client that tolerates absent fields keeps working, where Cloud's
// refusal fails it outright. Nothing is filled in, since Cloud reports no
// version, cluster id, or visibility store to fill it with.
func getClusterInfo() *Translation {
	return Answer(getClusterInfoMethod, clusterInfoReply)
}

// listSearchAttributes answers OperatorService.ListSearchAttributes with an empty
// reply rather than Cloud's PermissionDenied, which the Temporal UI treats as an
// expired login on every namespace page.
func listSearchAttributes() *Translation {
	return Answer(listSearchAttributesMethod, searchAttributesReply)
}

// clusterInfoReply clears the caller's reply, so a reused message cannot carry a
// value this answer never stated.
func clusterInfoReply(_ *workflowservice.GetClusterInfoRequest, reply *workflowservice.GetClusterInfoResponse) error {
	reply.Reset()
	return nil
}

// searchAttributesReply clears the caller's reply, for the same reason as
// clusterInfoReply.
func searchAttributesReply(
	_ *operatorservice.ListSearchAttributesRequest,
	reply *operatorservice.ListSearchAttributesResponse,
) error {
	reply.Reset()
	return nil
}
