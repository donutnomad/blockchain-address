package adrethereum

import (
	"testing"
)

func TestToChecksumAddress(t *testing.T) {
	var a1 = "0xaeaca00632fc155ba82c77f10cb678b0dcb5475c"
	if !ValidateAddressHex(a1, true) {
		panic("invalid address")
	}
}
