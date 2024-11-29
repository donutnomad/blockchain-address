package adrethereum

import (
	"golang.org/x/crypto/sha3"
	"strings"
)

func ValidateAddressHex(address string, validateChecksum bool) bool {
	if len(address) < 2 {
		return false
	}
	if address[0] == '0' && (address[1] == 'x' || address[1] == 'X') {
		address = address[2:]
	}
	if len(address) != 40 {
		return false
	}
	if validateChecksum {
		lower := strings.ToLower(address)
		if address != lower {
			out := ToChecksumAddress([40]byte([]byte(strings.ToLower(address))))
			if string(out[:]) != address {
				return false
			}
		}
	}
	return isHex(address)
}

func ToChecksumAddress(hexBs [40]byte) [40]byte {
	sha := sha3.NewLegacyKeccak256()
	sha.Write(hexBs[:])
	hash := sha.Sum(nil)
	for i := 0; i < len(hexBs); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if hexBs[i] > '9' && hashByte > 7 {
			hexBs[i] -= 32
		}
	}
	return hexBs
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// isHex validates whether each byte is valid hexadecimal string.
func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}
