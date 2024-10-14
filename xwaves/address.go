package xwaves

import (
	"github.com/mr-tron/base58"
	"github.com/pkg/errors"
)

const (
	AddressIDSize                 = 20
	wavesAddressHeaderSize        = 2
	wavesAddressBodySize          = AddressIDSize
	wavesAddressChecksumSize      = 4
	WavesAddressSize              = wavesAddressHeaderSize + wavesAddressBodySize + wavesAddressChecksumSize
	wavesAddressVersion      byte = 0x01
)

const (
	MainNetScheme   Scheme = 'W'
	TestNetScheme   Scheme = 'T'
	StageNetScheme  Scheme = 'S'
	CustomNetScheme Scheme = 'E'
)

type PublicKey [32]byte
type Scheme = byte

// WavesAddress is the transformed Public Key with additional bytes of the version, a blockchain scheme and a checksum.
type WavesAddress [WavesAddressSize]byte

// String produces the BASE58 string representation of the WavesAddress.
func (a WavesAddress) String() string {
	return base58.Encode(a[:])
}

// NewAddressFromPublicKey produces an WavesAddress from given scheme and Public Key bytes.
func NewAddressFromPublicKey(scheme Scheme, publicKey PublicKey) (WavesAddress, error) {
	return NewAddressFromPublicKeyWithVersion(wavesAddressVersion, scheme, publicKey)
}

func NewAddressFromEthereumAddress(scheme Scheme, address [20]byte) (WavesAddress, error) {
	return newWavesAddress(wavesAddressVersion, scheme, address)
}

func NewAddressFromPublicKeyWithVersion(version, scheme Scheme, publicKey PublicKey) (WavesAddress, error) {
	h, err := SecureHash(publicKey[:])
	if err != nil {
		return WavesAddress{}, errors.Wrap(err, "failed to produce Digest from PublicKey")
	}
	return newAddressFromPublicKeyHash(version, scheme, h)
}

// newAddressFromPublicKeyHash produces an WavesAddress from given public key hash.
func newAddressFromPublicKeyHash(version byte, scheme Scheme, pubKeyHash Digest) (WavesAddress, error) {
	var body [wavesAddressBodySize]byte
	copy(body[:], pubKeyHash[:])
	return newWavesAddress(version, scheme, body)
}

// newWavesAddress produces an WavesAddress from given body (AddressID).
func newWavesAddress(version byte, scheme Scheme, body [wavesAddressBodySize]byte) (WavesAddress, error) {
	var addr WavesAddress
	addr[0] = version
	addr[1] = scheme
	copy(addr[wavesAddressHeaderSize:], body[:])
	checksum, err := addressChecksum(addr[:wavesAddressHeaderSize+wavesAddressBodySize])
	if err != nil {
		return addr, errors.Wrap(err, "failed to calculate WavesAddress checksum")
	}
	copy(addr[wavesAddressHeaderSize+wavesAddressBodySize:], checksum[:])
	return addr, nil
}

func addressChecksum(b []byte) (cs [wavesAddressChecksumSize]byte, err error) {
	h, err := SecureHash(b)
	if err != nil {
		return cs, err
	}
	copy(cs[:], h[:wavesAddressChecksumSize])
	return cs, nil
}
