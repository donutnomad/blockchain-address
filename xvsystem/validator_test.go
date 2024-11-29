package xvsystem

import "testing"

func TestValidateAddressBase58(t *testing.T) {
	ok := ValidateAddressBase58("ARCyXtXjjqkd5MWBGp6veqqDwR4gz1SapDo")
	if !ok {
		panic("invalid address")
	}
}
