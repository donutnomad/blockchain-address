package xwaves

import (
	"bytes"
	"github.com/mr-tron/base58"
)

func ValidateAddressBase58(address string) bool {
	bs, err := base58.Decode(address)
	if err != nil {
		return false
	}
	if len(bs) != len(WavesAddress{}) {
		return false
	}
	scheme := bs[1]
	if scheme != MainNetScheme && scheme != TestNetScheme && scheme != StageNetScheme && scheme != CustomNetScheme {
		return false
	}
	checksum, err := addressChecksum(bs[:wavesAddressHeaderSize+wavesAddressBodySize])
	if err != nil {
		return false
	}
	return bytes.Compare(checksum[:], bs[wavesAddressHeaderSize+wavesAddressBodySize:]) == 0
}
