package xvsystem

import "github.com/donutnomad/blockchain-address/xwaves"

type VsysAddress = xwaves.WavesAddress
type PublicKey = [32]byte

const (
	MainNetScheme Scheme = 'M'
	TestNetScheme Scheme = 'T'
)

type Scheme = byte

// NewAddressFromPublicKey produces an WavesAddress from given scheme and Public Key bytes.
func NewAddressFromPublicKey(scheme Scheme, publicKey PublicKey) (out VsysAddress, _ error) {
	res, err := xwaves.NewAddressFromPublicKeyWithVersion(5, scheme, publicKey)
	if err != nil {
		return VsysAddress{}, err
	}
	return res, nil
}
