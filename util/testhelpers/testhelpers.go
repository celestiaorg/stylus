//
// Copyright 2021-2022, Offchain Labs, Inc. All rights reserved.
//

package testhelpers

import (
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/nitro/util/colors"
)

// Fail a test should an error occur
func RequireImpl(t *testing.T, err error, text ...string) {
	t.Helper()
	if err != nil {
		t.Fatal(colors.Red, text, err, colors.Clear)
	}
}

func FailImpl(t *testing.T, printables ...interface{}) {
	t.Helper()
	t.Fatal(colors.Red, printables, colors.Clear)
}

func RandomizeSlice(slice []byte) []byte {
	_, err := rand.Read(slice)
	if err != nil {
		panic(err)
	}
	return slice
}

func RandomAddress() common.Address {
	var address common.Address
	RandomizeSlice(address[:])
	return address
}
