package lock

import refs "github.com/epicchainlabs/epicchain-api-go/v2/refs/grpc"

// SetMembers sets `members` field.
func (x *Lock) SetMembers(ids []*refs.ObjectID) {
	x.Members = ids
}
