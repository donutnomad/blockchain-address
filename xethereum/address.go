package adrethereum

import (
	"golang.org/x/crypto/sha3"
)

const AddressSize = 20

type Address [AddressSize]byte

func (a Address) String() string {
	return string(checksumHex(a))
}

func NewAddressFromPublicKey(key [65]byte) (out Address) {
	return newAddress(keccak256(key[1:]))
}

func newAddress(s [32]byte) (out Address) {
	copy(out[:], s[12:])
	return
}

func keccak256(data ...[]byte) (out [32]byte) {
	h := sha3.NewLegacyKeccak256()
	for _, b := range data {
		h.Write(b)
	}
	h.Sum(out[:0])
	return
}

func encode(dst, src []byte) int {
	const hextable = "0123456789abcdef"
	j := 0
	for _, v := range src {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		j += 2
	}
	return len(src) * 2
}

func checksumHex(a [20]byte) []byte {
	var buf [len(a)*2 + 2]byte
	copy(buf[:2], "0x")
	encode(buf[2:], a[:])
	// compute checksum
	sha := sha3.NewLegacyKeccak256()
	sha.Write(buf[2:])
	hash := sha.Sum(nil)
	for i := 2; i < len(buf); i++ {
		hashByte := hash[(i-2)/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if buf[i] > '9' && hashByte > 7 {
			buf[i] -= 32
		}
	}
	return buf[:]
}
