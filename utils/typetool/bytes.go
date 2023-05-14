package util

import (
	"encoding/hex"
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"strings"
)

// SplitBytes Splits a byte array into a two-dimensional byte array of a given size
func SplitBytes(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}

// HexToBytes converts a hex string to a byte array.
// This should only be used for pre-defined test strings; it will panic if the input is invalid.
func HexToBytes(input string) []byte {
	res, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic(err)
	}
	return res
}

// HexToPubKey converts a hex string to a spec public key.
// This should only be used for pre-defined test strings; it will panic if the input is invalid.
func HexToPubKey(input string) spec.BLSPubKey {
	data := HexToBytes(input)
	var res spec.BLSPubKey
	copy(res[:], data)
	return res
}

// HexToSignature converts a hex string to a spec signature.
// This should only be used for pre-defined test strings; it will panic if the input is invalid.
func HexToSignature(input string) spec.BLSSignature {
	data := HexToBytes(input)
	var res spec.BLSSignature
	copy(res[:], data)
	return res
}

// HexToDomainType converts a hex string to a spec domain type.
// This should only be used for pre-defined test strings; it will panic if the input is invalid.
func HexToDomainType(input string) spec.DomainType {
	data := HexToBytes(input)
	var res spec.DomainType
	copy(res[:], data)
	return res
}

// HexToDomain converts a hex string to a spec domain.
// This should only be used for pre-defined test strings; it will panic if the input is invalid.
func HexToDomain(input string) spec.Domain {
	data := HexToBytes(input)
	var res spec.Domain
	copy(res[:], data)
	return res
}

// HexToVersion converts a hex string to a spec version.
// This should only be used for pre-defined test strings; it will panic if the input is invalid.
func HexToVersion(input string) spec.Version {
	data := HexToBytes(input)
	var res spec.Version
	copy(res[:], data)
	return res
}

// HexToRoot converts a hex string to a spec root.
// This should only be used for pre-defined test strings; it will panic if the input is invalid.
func HexToRoot(input string) spec.Root {
	data := HexToBytes(input)
	var res spec.Root
	copy(res[:], data)
	return res
}
