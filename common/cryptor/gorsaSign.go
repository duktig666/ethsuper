package cryptor

// SignMd5WithRsa Signed using the RSA With MD5 algorithm
func SignMd5WithRsa(data string, privateKey string) (string, error) {
	grsa := RSASecurity{}
	if err := grsa.SetPrivateKey(privateKey); err != nil {
		return "", err
	}

	sign, err := grsa.SignMd5WithRsa(data)
	if err != nil {
		return "", err
	}

	return sign, err
}

// SignSha1WithRsa Sign with the RSA With SHA1 algorithm
func SignSha1WithRsa(data string, privateKey string) (string, error) {
	grsa := RSASecurity{}
	if err := grsa.SetPrivateKey(privateKey); err != nil {
		return "", err
	}

	sign, err := grsa.SignSha1WithRsa(data)
	if err != nil {
		return "", err
	}

	return sign, err
}

// SignSha256WithRsa Sign with the RSA With SHA256 algorithm
func SignSha256WithRsa(data string, privateKey string) (string, error) {
	grsa := RSASecurity{}
	if err := grsa.SetPrivateKey(privateKey); err != nil {
		return "", err
	}

	sign, err := grsa.SignSha256WithRsa(data)
	if err != nil {
		return "", err
	}
	return sign, err
}

// VerifySignMd5WithRsa Verify signatures using RSA With MD 5
func VerifySignMd5WithRsa(data string, signData string, publicKey string) error {
	grsa := RSASecurity{}
	if err := grsa.SetPublicKey(publicKey); err != nil {
		return err
	}
	return grsa.VerifySignMd5WithRsa(data, signData)
}

// VerifySignSha1WithRsa Verify the signature using RSA With SHA 1
func VerifySignSha1WithRsa(data string, signData string, publicKey string) error {
	grsa := RSASecurity{}
	if err := grsa.SetPublicKey(publicKey); err != nil {
		return err
	}
	return grsa.VerifySignSha1WithRsa(data, signData)
}

// VerifySignSha256WithRsa Verify the signature using RSA With SHA 256
func VerifySignSha256WithRsa(data string, signData string, publicKey string) error {
	grsa := RSASecurity{}
	if err := grsa.SetPublicKey(publicKey); err != nil {
		return err
	}
	return grsa.VerifySignSha256WithRsa(data, signData)
}
