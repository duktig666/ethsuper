// description:
// @author renshiwei
// Date: 2022/8/18 19:15

package cryptor

import "github.com/sethvargo/go-password/password"

// GenerateRandomPass Randomly generate secure passwords
func GenerateRandomPass(length, numDigits, numSymbols int, noUpper, allowRepeat bool) (string, error) {
	return password.Generate(length, numDigits, numSymbols, noUpper, allowRepeat)
}

// Generate32LenPass Randomly generate secure passwords
func Generate32LenPass() (string, error) {
	return GenerateRandomPass(32, 10, 10, false, false)
}
