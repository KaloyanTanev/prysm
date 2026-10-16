package eth

import (
	"encoding/hex"
	"testing"

	"github.com/OffchainLabs/prysm/v7/testing/require"
)

// TestBuilderRequestAuth_HashTreeRoot pins the roots of the builder request
// authorization. The builder API defines both types as regular SSZ containers,
// not progressive ones, so the roots are plain Merkle roots over the fields.
func TestBuilderRequestAuth_HashTreeRoot(t *testing.T) {
	signature := make([]byte, 96)
	copy(signature, []byte{0x01, 0x02, 0x03})
	auth := &SignedBuilderRequestAuth{
		Message: &BuilderRequestAuth{
			Data: []byte("builder.example.com"),
			Slot: 12345,
		},
		Signature: signature,
	}

	messageRoot, err := auth.Message.HashTreeRoot()
	require.NoError(t, err)
	require.Equal(t, "6ad33c944510322cf6c8f8148a2e60a5b199367c6ef7ab4f0a98133f9c9fd6c9", hex.EncodeToString(messageRoot[:]))

	signedRoot, err := auth.HashTreeRoot()
	require.NoError(t, err)
	require.Equal(t, "f15386db9a7a1961f761390f93b14eebbe9c81e29b32cc3307d6fade2fa5dbe0", hex.EncodeToString(signedRoot[:]))
}
