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

package cryptor

import (
	"encoding/base64"
	"github.com/pkg/errors"
)

// PubkeyEncrypt public encrypt
func PubkeyEncrypt(data, publicKey string) (string, error) {

	rsaSecurity := RSASecurity{}
	err := rsaSecurity.SetPublicKey(publicKey)
	if err != nil {
		return "", errors.Wrap(err, "set public key err.")
	}

	rsaData, err := rsaSecurity.PubKeyEncrypt([]byte(data))
	if err != nil {
		return "", errors.Wrap(err, "PubKeyEncrypt err.")
	}

	return base64.StdEncoding.EncodeToString(rsaData), nil
}

// PrivateKeyEncrypt private encrypt
func PrivateKeyEncrypt(data, privateKey string) (string, error) {

	rsaSecurity := RSASecurity{}
	err := rsaSecurity.SetPrivateKey(privateKey)
	if err != nil {
		return "", errors.Wrap(err, "set public key err.")
	}

	rsaData, err := rsaSecurity.PriKeyEncrypt([]byte(data))
	if err != nil {
		return "", errors.Wrap(err, "PriKeyEncrypt err.")
	}

	return base64.StdEncoding.EncodeToString(rsaData), nil
}

// PubkeyDecrypt public decrypt
func PubkeyDecrypt(data, publicKey string) (string, error) {

	databs, _ := base64.StdEncoding.DecodeString(data)

	rsaSecurity := RSASecurity{}
	if err := rsaSecurity.SetPublicKey(publicKey); err != nil {
		return "", err
	}

	rsaData, err := rsaSecurity.PubKeyDecrypt(databs)
	if err != nil {
		return "", err
	}
	return string(rsaData), nil
}

// PrivateKeyDecrypt private decrypt
func PrivateKeyDecrypt(data, privateKey string) (string, error) {

	databs, _ := base64.StdEncoding.DecodeString(data)

	rsaSecurity := RSASecurity{}

	if err := rsaSecurity.SetPrivateKey(privateKey); err != nil {
		return "", err
	}

	rsaData, err := rsaSecurity.PriKeyDecrypt(databs)
	if err != nil {
		return "", err
	}

	return string(rsaData), nil
}
