//  Copyright © 2022-2023. duktig666 Limited.
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package eth1

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateEth1Key(t *testing.T) {
	priKey, pubkey, err := CreateEth1Key()
	require.Nil(t, err)
	fmt.Println(priKey)
	fmt.Println(pubkey)
}

func TestSignStr(t *testing.T) {
	out := "5377ed989297d9e42944cc152c29990c260064979d1effc1acd6890c1d513e8b19bee1f7416dde9544600cc8ec2c6d4c620870433a4971ab6684cdd1a80ca42100"

	str := hex.EncodeToString(crypto.Keccak256([]byte("123456")))
	fmt.Println(str)

	sign, err := SignHexStr(str, "4e043906b895ca7ce3a64a2661927c915d8e7ca5c30b00a3bec4b6ff0746d957")
	require.Nil(t, err)
	require.Equal(t, sign, out)
}

func TestPubkeyFromPrivateKey(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	privateKeyBytes := crypto.FromECDSA(privateKey)
	priv := hexutil.Encode(privateKeyBytes)

	address, err := AddressFromPrivateKey(priv)
	require.NoError(t, err)
	fmt.Println(address)

	address2, err := AddressFromPrivateKey(priv[2:])
	require.NoError(t, err)
	require.Equal(t, address, address2)
}
