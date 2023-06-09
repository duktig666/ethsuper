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

// description:KeystoreV4 加解密
// @author renshiwei
// Date: 2022/8/20 22:37

package keystorev4

import (
	e2types "github.com/wealdtech/go-eth2-types/v2"
	encryptorKeystore "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
)

var (
	// Pbkdf2Encryptor Default (less resource consumption, wider usage)
	Pbkdf2Encryptor *encryptorKeystore.Encryptor
	ScryptEncryptor *encryptorKeystore.Encryptor
)

func init() {
	err := e2types.InitBLS()
	if err != nil {
		panic("InitBLS err.")
	}

	Pbkdf2Encryptor = encryptorKeystore.New()
	ScryptEncryptor = encryptorKeystore.New(encryptorKeystore.WithCipher("scrypt"))
}

// SwitchEncryptor Switch KeystoreV4 Encryptor.
// Default pbkdf2: less resource consumption, wider usage.
func SwitchEncryptor(keystore *KeystoreV4) *encryptorKeystore.Encryptor {
	var encryptor *encryptorKeystore.Encryptor
	switch keystore.Crypto.KDF.Function {
	case "pbkdf2":
		encryptor = Pbkdf2Encryptor
	case "scrypt":
		encryptor = ScryptEncryptor
	default:
		encryptor = Pbkdf2Encryptor
	}
	return encryptor
}
