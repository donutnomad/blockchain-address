package xwaves

import "testing"

func TestValidateAddressBase58(t *testing.T) {
	ok := ValidateAddressBase58("ARCyXtXjjqkd5MWBGp6veqqDwR4gz1SapDo")
	if ok {
		panic("check")
	}
	ok = ValidateAddressBase58("3PCFuURVbtpke5zoRzhBV8hEDDZSVAbPebj")
	if !ok {
		panic("invalid address")
	}
}
