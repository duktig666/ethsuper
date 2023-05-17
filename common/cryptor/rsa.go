package cryptor

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

var RSA = &RSASecurity{}

type RSASecurity struct {
	pubStr     string
	priStr     string
	pubkey     *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// SetPublicKey set public key
func (rsas *RSASecurity) SetPublicKey(pubStr string) error {
	var err error
	rsas.pubStr = pubStr
	rsas.pubkey, err = rsas.PublicKey()
	return err
}

// SetPrivateKey set private key
func (rsas *RSASecurity) SetPrivateKey(priStr string) error {
	var err error
	rsas.priStr = priStr
	rsas.privateKey, err = rsas.PrivateKey()
	return err
}

// PrivateKey *rsa.PrivateKey
func (rsas *RSASecurity) PrivateKey() (*rsa.PrivateKey, error) {
	return RsaPrivateKeyFromBytes([]byte(rsas.priStr))
}

// PublicKey *rsa.PublicKey
func (rsas *RSASecurity) PublicKey() (*rsa.PublicKey, error) {
	return RsaPubkeyFromBytes([]byte(rsas.pubStr))
}

// PubKeyEncrypt pubkey encrypt
func (rsas *RSASecurity) PubKeyEncrypt(input []byte) ([]byte, error) {
	if rsas.pubkey == nil {
		return []byte(""), errors.New(`Please set the public key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(rsas.pubkey, bytes.NewReader(input), output, true)
	if err != nil {
		return []byte(""), err
	}
	return io.ReadAll(output)
}

// PubKeyDecrypt pubkey decrypt
func (rsas *RSASecurity) PubKeyDecrypt(input []byte) ([]byte, error) {
	if rsas.pubkey == nil {
		return []byte(""), errors.New(`Please set the public key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(rsas.pubkey, bytes.NewReader(input), output, false)
	if err != nil {
		return []byte(""), err
	}
	return io.ReadAll(output)
}

// PriKeyEncrypt private key encrypt
func (rsas *RSASecurity) PriKeyEncrypt(input []byte) ([]byte, error) {
	if rsas.privateKey == nil {
		return []byte(""), errors.New(`Please set the private key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(rsas.privateKey, bytes.NewReader(input), output, true)
	if err != nil {
		return []byte(""), err
	}
	return io.ReadAll(output)
}

// PriKeyDecrypt private key decrypt
func (rsas *RSASecurity) PriKeyDecrypt(input []byte) ([]byte, error) {
	if rsas.privateKey == nil {
		return []byte(""), errors.New(`Please set the private key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(rsas.privateKey, bytes.NewReader(input), output, false)
	if err != nil {
		return []byte(""), err
	}

	return io.ReadAll(output)
}

// SignMd5WithRsa PriKeySign private key sign by md5
func (rsas *RSASecurity) SignMd5WithRsa(data string) (string, error) {
	md5Hash := md5.New()
	s_data := []byte(data)
	md5Hash.Write(s_data)
	hashed := md5Hash.Sum(nil)

	signByte, err := rsa.SignPKCS1v15(rand.Reader, rsas.privateKey, crypto.MD5, hashed)
	sign := base64.StdEncoding.EncodeToString(signByte)
	return sign, err
}

// SignSha1WithRsa PriKeySign private key sign by sha1
func (rsas *RSASecurity) SignSha1WithRsa(data string) (string, error) {
	sha1Hash := sha1.New()
	s_data := []byte(data)
	sha1Hash.Write(s_data)
	hashed := sha1Hash.Sum(nil)

	signByte, err := rsa.SignPKCS1v15(rand.Reader, rsas.privateKey, crypto.SHA1, hashed)
	sign := base64.StdEncoding.EncodeToString(signByte)
	return sign, err
}

// SignSha256WithRsa PriKeySign private key sign by sha256
func (rsas *RSASecurity) SignSha256WithRsa(data string) (string, error) {
	sha256Hash := sha256.New()
	s_data := []byte(data)
	sha256Hash.Write(s_data)
	hashed := sha256Hash.Sum(nil)

	signByte, err := rsa.SignPKCS1v15(rand.Reader, rsas.privateKey, crypto.SHA256, hashed)
	sign := base64.StdEncoding.EncodeToString(signByte)
	return sign, err
}

// VerifySignMd5WithRsa PubKeyVerifySign public key verify sign by md5
func (rsas *RSASecurity) VerifySignMd5WithRsa(data string, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	hash := md5.New()
	hash.Write([]byte(data))
	return rsa.VerifyPKCS1v15(rsas.pubkey, crypto.MD5, hash.Sum(nil), sign)
}

// VerifySignSha1WithRsa PubKeyVerifySign public key verify sign by sha1
func (rsas *RSASecurity) VerifySignSha1WithRsa(data string, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write([]byte(data))
	return rsa.VerifyPKCS1v15(rsas.pubkey, crypto.SHA1, hash.Sum(nil), sign)
}

// VerifySignSha256WithRsa PubKeyVerifySign public key verify sign by sha256
func (rsas *RSASecurity) VerifySignSha256WithRsa(data string, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	hash := sha256.New()
	hash.Write([]byte(data))

	return rsa.VerifyPKCS1v15(rsas.pubkey, crypto.SHA256, hash.Sum(nil), sign)
}
