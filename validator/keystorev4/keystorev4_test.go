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

package keystorev4

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
	"testing"
)

func TestGeMyKeystoreV4Json(t *testing.T) {
	keystore, pass, err := GenerateKeystoreV4ByRandomPass()
	require.Nil(t, err)
	keystoreJson, err := keystore.String()
	require.Nil(t, err)
	fmt.Println("keystoreJson:", keystoreJson)
	fmt.Println("pass:", pass)
}

func Test2MyKeystoreV4(t *testing.T) {
	key := keystorev4.New()

	cusPri, err := hex.DecodeString("07733d5deefeced8bbbe8df0e46e52eabcead8c01e10fd3a1f1cfe670af4dcc2")
	require.Nil(t, err)

	keystoreMap, err := key.Encrypt(cusPri, "ethsuper")
	require.Nil(t, err)

	// keystoreMap parse keystore.json
	marshal, err := json.Marshal(keystoreMap)
	fmt.Printf("keystoreMap json:%v\n", string(marshal))
	require.Nil(t, err)

	var keystore KeystoreV4
	err = json.Unmarshal(marshal, &keystore.Crypto)
	require.Nil(t, err)

	keystoreStr, err := keystore.String()
	require.Nil(t, err)

	fmt.Printf("keystoreStr:%+v\n", keystoreStr)
}

func TestDecrypt(t *testing.T) {
	// pbkdf2
	keystore2 := `
		{"crypto":{"kdf":{"function":"pbkdf2","params":{"salt":"57c8f524e2eb67cc2dd951a936d17518410e3db21802304af5a9207c8de3abd9","dklen":32,"c":262144,"prf":"hmac-sha256"},"message":""},"checksum":{"function":"sha256","params":{},"message":"e13513c9bd2d923ada0cd2f53966c907633e1c6d489b941c0bdd48d82688a1d6"},"cipher":{"function":"aes-128-ctr","params":{"iv":"83bdaf4a6a5e7c78bb0d60d119a0791e"},"message":"aef74389e1b5f632b49e2288e799554e61f2287c7ea1b081fb12b08a5f404e1f"}},"path":"","pubkey":"ab196b549773f6f9c1cae6541cde9e1a0ff8b1c4b5994aeee34b37a0beb5fd50a6cd8a73e5a427434defe100a49e2e43","uuid":"46dc22bc-20a3-11ed-9ab1-c234cac5c4e9","version":4}
	`
	pass2 := "28w6^HT$ukrj39iC`*W+7V1<5x=?4\\0_"

	privateKey := "0de561798bb5ff94cfa1499693e0749df77f16570958072278abd6bf4fbec692"

	secret2, err := DecryptFromJson(keystore2, pass2)
	require.Nil(t, err)

	require.Equal(t, privateKey, hex.EncodeToString(secret2))
}

func TestEncrypt(t *testing.T) {
	pri := "0de561798bb5ff94cfa1499693e0749df77f16570958072278abd6bf4fbec692"
	pass := "28w6^HT$ukrj39iC`*W+7V1<5x=?4\\0_"

	keystoreStr, err := EncryptForJsonStr(pri, pass)
	require.Nil(t, err)

	fmt.Println(keystoreStr)
}
