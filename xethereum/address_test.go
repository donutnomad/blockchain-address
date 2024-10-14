package adrethereum

import (
	"encoding/hex"
	"testing"
)

func TestAddress(t *testing.T) {
	pubBs, err := hex.DecodeString("04617b558022262f7fe898aef9535070d4806ad06910183824b53d5f0a63f6a9d02544c07fb4c655e31434cec5d58892c0b845e9522d235b3c03017648bb799eb6")
	if err != nil {
		panic(err)
	}
	address := NewAddressFromPublicKey([65]byte(pubBs[:]))
	if address.String() != "0x50EFEde5e295E029E434c0807D869Cb296Cc7Dd8" {
		t.Fatalf("invalid address")
	}
}
