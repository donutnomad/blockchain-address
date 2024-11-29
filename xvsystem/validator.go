package xvsystem

import (
	"bytes"
	"github.com/donutnomad/blockchain-address/xwaves"
	"github.com/mr-tron/base58"
)

func ValidateAddressBase58(address string) bool {
	bs, err := base58.Decode(address)
	if err != nil {
		return false
	}
	if len(bs) != len(VsysAddress{}) {
		return false
	}
	scheme := bs[1]
	if scheme != MainNetScheme && scheme != TestNetScheme {
		return false
	}
	checksum, err := xwaves.AddressChecksum(bs[:2+20])
	if err != nil {
		return false
	}
	return bytes.Compare(checksum[:], bs[2+20:]) == 0
}
