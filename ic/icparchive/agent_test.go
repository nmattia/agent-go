// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package icparchive_test

import (
	"github.com/aviate-labs/agent-go"

	"github.com/aviate-labs/agent-go/mock"
	"github.com/aviate-labs/agent-go/principal"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/aviate-labs/agent-go/ic/icparchive"
)

// Test_GetBlocks tests the "get_blocks" method on the "icparchive" canister.
func Test_GetBlocks(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "get_blocks",
			Arguments: []any{new(icparchive.GetBlocksArgs)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(icparchive.GetBlocksResult)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 icparchive.GetBlocksArgs
	if _, err := a.GetBlocks(a0); err != nil {
		t.Fatal(err)
	}

}

// newAgent creates a new agent with the given (mock) methods.
// Runs a mock replica in the background.
func newAgent(methods []mock.Method) (*icparchive.Agent, error) {
	replica := mock.NewReplica()
	canisterId := principal.Principal{Raw: []byte("icparchive")}
	replica.AddCanister(canisterId, methods)
	s := httptest.NewServer(replica)
	u, _ := url.Parse(s.URL)
	a, err := icparchive.NewAgent(canisterId, agent.Config{
		ClientConfig: &agent.ClientConfig{Host: u},
		FetchRootKey: true,
	})
	if err != nil {
		return nil, err
	}
	return a, nil
}