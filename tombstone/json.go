package tombstone

import (
	"github.com/epicchainlabs/epicchain-api-go/v2/rpc/message"
	tombstone "github.com/epicchainlabs/epicchain-api-go/v2/tombstone/grpc"
)

func (s *Tombstone) MarshalJSON() ([]byte, error) {
	return message.MarshalJSON(s)
}

func (s *Tombstone) UnmarshalJSON(data []byte) error {
	return message.UnmarshalJSON(s, data, new(tombstone.Tombstone))
}
