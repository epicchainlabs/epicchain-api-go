package subnettest

import (
	refstest "github.com/epicchainlabs/epicchain-api-go/v2/refs/test"
	"github.com/epicchainlabs/epicchain-api-go/v2/subnet"
)

func GenerateSubnetInfo(empty bool) *subnet.Info {
	m := new(subnet.Info)

	if !empty {
		m.SetID(refstest.GenerateSubnetID(false))
		m.SetOwner(refstest.GenerateOwnerID(false))
	}

	return m
}
