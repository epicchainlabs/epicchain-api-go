package accounting

import (
	accounting "github.com/epicchainlabs/epicchain-api-go/v2/accounting/grpc"
	"github.com/epicchainlabs/epicchain-api-go/v2/rpc/message"
)

func (d *Decimal) MarshalJSON() ([]byte, error) {
	return message.MarshalJSON(d)
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	return message.UnmarshalJSON(d, data, new(accounting.Decimal))
}
