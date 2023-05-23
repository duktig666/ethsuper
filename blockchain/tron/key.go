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

package tron

import (
	"github.com/duktig666/ethsuper/eth1"
	"github.com/pkg/errors"
	"strings"
)

// GenerateKey Generate tron private key and pubkey
func GenerateKey() (string, string, error) {
	privateKey, ethPubkey, err := eth1.CreateEth1Key()
	if err != nil {
		return "", "", errors.Wrap(err, "generate tron key err.")
	}
	tronPubkey := strings.Replace(ethPubkey, "0x", "41", 1)
	return privateKey, tronPubkey, nil
}
