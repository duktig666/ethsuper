package tron

import (
	"github.com/duktig666/ethsuper/eth1"
	"github.com/pkg/errors"
	"strings"
)

// GenerateKey Generate tron private key and pubkey
func GenerateKey() (string, string, error) {
	privateKey, ethPubkey, err := eth1.CreateEth1Key()
	if err != nil {
		return "", "", errors.Wrap(err, "generate tron key err.")
	}
	tronPubkey := strings.Replace(ethPubkey, "0x", "41", 1)
	return privateKey, tronPubkey, nil
}
