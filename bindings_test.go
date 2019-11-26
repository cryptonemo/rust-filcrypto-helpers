package pedersen_hash_test

import (
	ped "../rust-pedersen-hash"
	"fmt"
	"testing"
)

func TestPedersenHash(t *testing.T) {
	var input [64]byte
	fmt.Println("input: ", input)
	output := ped.PedersenHashFFI(input[:])
	fmt.Println("output: ", output);
}
