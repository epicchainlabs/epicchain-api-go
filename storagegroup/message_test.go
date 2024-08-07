package storagegroup_test

import (
	"testing"

	"github.com/epicchainlabs/epicchain-api-go/v2/rpc/message"
	messagetest "github.com/epicchainlabs/epicchain-api-go/v2/rpc/message/test"
	storagegrouptest "github.com/epicchainlabs/epicchain-api-go/v2/storagegroup/test"
)

func TestMessageConvert(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return storagegrouptest.GenerateStorageGroup(empty) },
	)
}
