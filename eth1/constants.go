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

// desc:
// @author renshiwei
// Date: 2023/4/7 19:31

package eth1

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/params"
	"github.com/shopspring/decimal"
	"math/big"
)

const (
	MAINNET = "mainnet"
	GOERLI  = "goerli"
	PRATER  = "prater"
	SEPOLIA = "sepolia"
)

var ZERO_HASH = [32]byte{}

const (
	ZERO_HASH_STR    = "0x0000000000000000000000000000000000000000000000000000000000000000"
	ZERO_ADDRESS_STR = "0x0000000000000000000000000000000000000000"
)

func GWEIToWEI(value *big.Int) *big.Int {
	return new(big.Int).Mul(value, big.NewInt(params.GWei))
}

func ETH32() decimal.Decimal {
	eth32, _ := decimal.NewFromString("32000000000000000000")
	return eth32
}

func GweiToWei(gwei phase0.Gwei) *big.Int {
	return new(big.Int).Mul(big.NewInt(int64(gwei)), big.NewInt(params.GWei))
}

func ETH(eth int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(eth))
}
