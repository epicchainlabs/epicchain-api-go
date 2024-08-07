package accounting_test

import (
	"testing"

	accountingtest "github.com/epicchainlabs/epicchain-api-go/v2/accounting/test"
	"github.com/epicchainlabs/epicchain-api-go/v2/rpc/message"
	messagetest "github.com/epicchainlabs/epicchain-api-go/v2/rpc/message/test"
)

func TestMessage(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return accountingtest.GenerateDecimal(empty) },
		func(empty bool) message.Message { return accountingtest.GenerateBalanceRequestBody(empty) },
		func(empty bool) message.Message { return accountingtest.GenerateBalanceRequest(empty) },
		func(empty bool) message.Message { return accountingtest.GenerateBalanceResponseBody(empty) },
		func(empty bool) message.Message { return accountingtest.GenerateBalanceResponse(empty) },
	)
}
