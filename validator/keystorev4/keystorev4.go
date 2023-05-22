// encryptorKeystore v4 Encryptor @see https://www.cnblogs.com/wanghui-garcia/p/10007768.html

package keystorev4

import (
	"encoding/hex"
	"encoding/json"
	"github.com/duktig666/ethsuper/common/cryptor"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	e2types "github.com/wealdtech/go-eth2-types/v2"
)

type ksKDFParams struct {
	// Shared parameters
	Salt  string `json:"salt"`
	DKLen int    `json:"dklen"`
	// Scrypt-specific parameters
	N int `json:"n,omitempty"`
	P int `json:"p,omitempty"`
	R int `json:"r,omitempty"`
	// PBKDF2-specific parameters
	C   int    `json:"c,omitempty"`
	PRF string `json:"prf,omitempty"`
}

type ksKDF struct {
	Function string       `json:"function"`
	Params   *ksKDFParams `json:"params"`
	Message  string       `json:"message"`
}

type ksChecksum struct {
	Function string                 `json:"function"`
	Params   map[string]interface{} `json:"params"`
	Message  string                 `json:"message"`
}

type ksCipherParams struct {
	// AES-128-CTR-specific parameters
	IV string `json:"iv,omitempty"`
}

type ksCipher struct {
	Function string          `json:"function"`
	Params   *ksCipherParams `json:"params"`
	Message  string          `json:"message"`
}

type ksCrypto struct {
	KDF      *ksKDF      `json:"kdf"`
	Checksum *ksChecksum `json:"checksum"`
	Cipher   *ksCipher   `json:"cipher"`
}

type KeystoreV4 struct {
	Crypto  *ksCrypto `json:"crypto"`
	Path    string    `json:"path"`
	Pubkey  string    `json:"pubkey"`
	Uuid    string    `json:"uuid"`
	Version int       `json:"version" default:"4"`
}

func (k *KeystoreV4) Marshal() ([]byte, error) {
	return json.Marshal(k)
}

func (k *KeystoreV4) Unmarshal(input []byte) error {
	err := json.Unmarshal(input, k)
	if err != nil {
		return errors.Wrap(err, "json Unmarshal err.")
	}
	return nil
}

func (k *KeystoreV4) String() (string, error) {
	data, err := k.Marshal()
	return string(data), err
}

// Decrypt decrypt Keystore V4
func (k *KeystoreV4) Decrypt(pass string) ([]byte, error) {
	pubkey := k.Pubkey
	encryptor := SwitchEncryptor(k)

	marshal, err := json.Marshal(k.Crypto)
	var cryptoMap map[string]interface{}
	err = json.Unmarshal(marshal, &cryptoMap)
	if err != nil {
		return nil, errors.Errorf("The keystore format does not conform to err. pubkey:%v", pubkey)
	}

	deSecret, err := encryptor.Decrypt(cryptoMap, pass)
	if err != nil {
		return nil, errors.Errorf("keystore decrypt error. pubkey:%v", pubkey)
	}

	return deSecret, nil
}

// GenerateKeystoreV4ByRandomPass @return KeystoreV4;pass;err
func GenerateKeystoreV4ByRandomPass() (*KeystoreV4, string, error) {
	// Generate a password
	pass, err := cryptor.Generate32LenPass()
	if err != nil {
		return nil, "", errors.New("generate pass error")
	}

	keystore, err := GenerateKeystoreV4(pass)
	if err != nil {
		return nil, "", errors.Wrap(err, "GenerateKeystoreV4 error")
	}

	return keystore, pass, nil
}

// GenerateRandomBLSPrivateKey Generate BLS Private Key
func GenerateRandomBLSPrivateKey() (*e2types.BLSPrivateKey, error) {
	blsPrivateKey, err := e2types.GenerateBLSPrivateKey()
	if err != nil {
		return nil, errors.Wrap(err, "GenerateBLSPrivateKey err.")
	}
	return blsPrivateKey, nil
}

// GenerateKeystoreV4 Generate keystoreV4
func GenerateKeystoreV4(pass string) (*KeystoreV4, error) {
	blsPrivateKey, err := GenerateRandomBLSPrivateKey()
	if err != nil {
		return nil, errors.Wrap(err, "GenerateKeystoreV4 err.")
	}

	return GenerateKeystoreV4FromPrivateKey(blsPrivateKey.Marshal(), pass)
}

// GenerateKeystoreV4FromPrivateKey Generate keystoreV4 from the private key
func GenerateKeystoreV4FromPrivateKey(privateKey []byte, pass string) (*KeystoreV4, error) {
	blsPrivateKey, err := e2types.BLSPrivateKeyFromBytes(privateKey)

	pubkey := hex.EncodeToString(blsPrivateKey.PublicKey().Marshal())

	secret := blsPrivateKey.Marshal()
	keystoreMap, err := Pbkdf2Encryptor.Encrypt(secret, pass)
	if err != nil {
		return nil, errors.Wrap(err, "encrypt BLSPrivateKey error")
	}

	// keystoreMap resolves to keystore.json
	marshal, err := json.Marshal(keystoreMap)
	if err != nil {
		return nil, errors.Wrap(err, "privateKey to json error")
	}

	var keystore KeystoreV4

	// keystore Crypto
	err = json.Unmarshal(marshal, &keystore.Crypto)
	if err != nil {
		return nil, errors.Wrap(err, "privateKey json to KeystoreV4 struct error")
	}

	// uuid
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.Wrap(err, "uuid error")
	}

	// Set other properties of Keystore
	keystore.Version = 4
	keystore.Uuid = newUUID.String()
	keystore.Pubkey = pubkey

	return &keystore, nil
}

// DecryptFromJson Decrypt KeystoreV4 json str
func DecryptFromJson(keystoreV4Str string, pass string) ([]byte, error) {
	var keystore *KeystoreV4
	err := keystore.Unmarshal([]byte(keystoreV4Str))
	if err != nil {
		return nil, errors.Wrap(err, "The keystore JSON str format does not conform to err.")
	}
	return keystore.Decrypt(pass)
}

// Encrypt KeystoreV4 privateSecret encrypt
func Encrypt(privateSecret, pass string) (*KeystoreV4, error) {
	decodeString, err := hex.DecodeString(privateSecret)
	if err != nil {
		return nil, errors.Wrap(err, "privateSecret DecodeString err.")
	}
	return GenerateKeystoreV4FromPrivateKey(decodeString, pass)
}

// EncryptForJsonStr KeystoreV4 privateSecret encrypt
func EncryptForJsonStr(privateSecret, pass string) (string, error) {
	keystore, err := Encrypt(privateSecret, pass)
	if err != nil {
		return "", errors.Wrap(err, "EncryptForJsonStr err.")
	}

	s, err := keystore.String()
	if err != nil {
		return "", errors.Wrap(err, "EncryptForJsonStr err.")
	}
	return s, nil
}
