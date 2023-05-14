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
