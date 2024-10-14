package xwaves

import (
	"encoding/hex"
	"github.com/mr-tron/base58/base58"
	"github.com/pkg/errors"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
	"strings"
)

const (
	DigestSize = 32
)

type Digest [DigestSize]byte

func (d Digest) String() string {
	return base58.Encode(d[:])
}

func (d Digest) Hex() string {
	return hex.EncodeToString(d[:])
}

func (d Digest) ShortString() string {
	str := base58.Encode(d[:])
	sb := new(strings.Builder)
	sb.WriteString(str[:6])
	sb.WriteRune(0x2026) //22ef
	sb.WriteString(str[len(str)-6:])
	return sb.String()
}

func (d Digest) Bytes() []byte {
	return d[:]
}

func (d Digest) MarshalBinary() ([]byte, error) {
	return d[:], nil
}

func (d *Digest) UnmarshalBinary(data []byte) error {
	if l := len(data); l < DigestSize {
		return errors.Errorf("failed unmarshal Digest, required %d bytes, got %d", DigestSize, l)
	}
	copy(d[:], data[:DigestSize])
	return nil
}

func SecureHash(data []byte) (Digest, error) {
	var d Digest
	fh, err := blake2b.New256(nil)
	if err != nil {
		return d, err
	}
	if _, err := fh.Write(data); err != nil {
		return d, err
	}
	fh.Sum(d[:0])
	h := sha3.NewLegacyKeccak256()
	if _, err := h.Write(d[:DigestSize]); err != nil {
		return d, err
	}
	h.Sum(d[:0])
	return d, nil
}
