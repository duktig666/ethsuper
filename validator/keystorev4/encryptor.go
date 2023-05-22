// description:KeystoreV4 加解密
// @author renshiwei
// Date: 2022/8/20 22:37

package keystorev4

import (
	encryptorKeystore "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
)

var (
	// Pbkdf2Encryptor Default (less resource consumption, wider usage)
	Pbkdf2Encryptor *encryptorKeystore.Encryptor
	ScryptEncryptor *encryptorKeystore.Encryptor
)

func init() {
	Pbkdf2Encryptor = encryptorKeystore.New()
	ScryptEncryptor = encryptorKeystore.New(encryptorKeystore.WithCipher("scrypt"))
}

// SwitchEncryptor 选择匹配的 KeystoreV4 Encryptor
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
