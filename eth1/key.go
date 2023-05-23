/*
 * // Copyright © 2022-2023. duktig666 Limited.
 * // Licensed under the Apache License, Version 2.0 (the "License");
 * // you may not use this file except in compliance with the License.
 * // You may obtain a copy of the License at
 * //
 * //     http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing, software
 * // distributed under the License is distributed on an "AS IS" BASIS,
 * // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * // See the License for the specific language governing permissions and
 * // limitations under the License.
 */

// desc:
// @author renshiwei
// Date: 2023/4/11 17:02

package eth1

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"strings"
)

// CreateEth1Key Create ETH1 private key and pubkey
func CreateEth1Key() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", errors.Wrap(err, "publicKey is not *ecdsa.PublicKey")
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	priv := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("publicKey is not *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return priv, address, nil
}

// SignHexStr ETH private key signature
func SignHexStr(str, priv string) (string, error) {
	decodeStr, err := hex.DecodeString(str)
	if err != nil {
		return "", errors.Wrapf(err, "sign param:%s DecodeString err.", str)
	}

	priKey, err := crypto.HexToECDSA(priv)
	if err != nil {
		return "", errors.Wrap(err, "priKey read err.")
	}
	signStr, err := crypto.Sign(decodeStr, priKey)
	if err != nil {
		return "", errors.Wrapf(err, "sign param:%s err.", str)
	}
	return hex.EncodeToString(signStr), nil
}

// AddressFromPrivateKey Push out the address from the private key
func AddressFromPrivateKey(pri string) (common.Address, error) {
	if !strings.HasPrefix(pri, "0x") {
		pri = fmt.Sprintf("0x%s", pri)
	}
	priHex, err := hexutil.Decode(pri)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "Failed to Decode")
	}
	privateKey, err := crypto.ToECDSA(priHex)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "Failed to ToECDSA")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("publicKey is not a *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, nil
}
