package filcrypto_helpers_test

import (
	helpers "../rust-filcrypto-helpers"
	"fmt"
	"testing"
)

func TestSimplePedersenHash(t *testing.T) {
	var input [64]byte
	fmt.Println("input: ", input)
	output := helpers.PedersenHash(input[:])
	fmt.Println("output: ", output);
}
