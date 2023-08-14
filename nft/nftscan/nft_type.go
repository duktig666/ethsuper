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

// description:
// @author renshiwei
// Date: 2023/7/18

package nftscan

import (
	"github.com/duktig666/ethsuper/contracts/erc721"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const (
	Erc721 = iota
	Erc1155
	Unknown
)

func JudgmentErcType(contract common.Address, backend bind.ContractBackend) (int, error) {
	instance, err := erc721.NewErc721(contract, backend)
	if err != nil {
		return 0, err
	}

	if ok, err := instance.SupportsInterface(&bind.CallOpts{}, [4]byte{0x80, 0xac, 0x58, 0xcd}); err == nil && ok {
		return Erc721, nil
	}

	if ok, err := instance.SupportsInterface(&bind.CallOpts{}, [4]byte{0xd9, 0xb6, 0x7a, 0x26}); err == nil && ok {
		return Erc1155, nil
	}

	return Unknown, nil
}
