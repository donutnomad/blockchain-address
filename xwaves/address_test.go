package xwaves

import (
	"encoding/hex"
	"testing"
)

func TestAddress(t *testing.T) {
	bs, err := hex.DecodeString("50afa06841a3d1d1969a0e8aa55f36a50adbbea1abc7bc6b06275914eb91fd26")
	if err != nil {
		panic(err)
	}
	key, err := NewAddressFromPublicKey(MainNetScheme, PublicKey(bs))
	if err != nil {
		panic(err)
	}
	if key.String() != "3PCFuURVbtpke5zoRzhBV8hEDDZSVAbPebj" {
		t.Fatalf("invalid address")
	}
}
