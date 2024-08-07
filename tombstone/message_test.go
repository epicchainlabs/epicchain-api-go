package tombstone_test

import (
	"testing"

	"github.com/epicchainlabs/epicchain-api-go/v2/rpc/message"
	messagetest "github.com/epicchainlabs/epicchain-api-go/v2/rpc/message/test"
	tombstonetest "github.com/epicchainlabs/epicchain-api-go/v2/tombstone/test"
)

func TestMessageConvert(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return tombstonetest.GenerateTombstone(empty) },
	)
}
