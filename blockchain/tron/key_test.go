package tron

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	privateKey, pubkey, err := GenerateKey()
	require.NoError(t, err)
	fmt.Printf("privateKey: %s\npubkey: %s\n", privateKey, pubkey)
}
