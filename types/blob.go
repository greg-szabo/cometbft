package types

import (
	"encoding/hex"

	"github.com/cometbft/cometbft/crypto/tmhash"
)

// Blob represents an Ethereum blob, a binary data structure introduced in
// EIP-4844 (Proto-Danksharding).
// Blobs are primarily used by Layer 2 rollups to post batched transaction data
// to Ethereum in a cost-efficient manner. Each blob is referenced in a blob
// transaction via a KZG commitment, which ensures data integrity without
// requiring full storage in execution clients.
//
// A "Blob" is typically an opaque byte slice up to ~128 KB. In our case, however,
// a single Blob slice can hold multiple blobs, potentially exceeding that size.
// The exact number of contained blobs is encoded within the block's transactions,
// but CometBFT neither parses nor uses this information. It simply gossips the
// blobs alongside their corresponding blocks.
//
// We introduce support for blobs in CometBFT to serve use-cases where CometBFT
// acts purely as a finalization gadget, so storing blobs in CometBFT's storage
// layer is unnecessary and undesirable.
type Blob []byte

// Hash returns the SHA-256 hash of the blob.
func (b Blob) Hash() []byte {
	return tmhash.Sum(b)
}

// String returns a hex-encoded representation of the blob.
func (b Blob) String() string {
	if b == nil {
		return "nil-Blob"
	}
	return "Blob#" + hex.EncodeToString(b)
}
