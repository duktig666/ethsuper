/*
 * // Copyright © 2022-2023. duktig666 Limited.
 * // Licensed under the Apache License, Version 2.0 (the "License");
 * // you may not use this file except in compliance with the License.
 * // You may obtain a copy of the License at
 * //
 * //     http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing, software
 * // distributed under the License is distributed on an "AS IS" BASIS,
 * // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * // See the License for the specific language governing permissions and
 * // limitations under the License.
 */

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
