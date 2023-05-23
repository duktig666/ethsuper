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

package eth1

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/pkg/errors"
	"math/big"
)

// NewSimulatedClient ETH client of constructing simulation transactions
func NewSimulatedClient(walletBalance *big.Int, gasLimit uint64, privateKeyStr string) (*backends.SimulatedBackend, error) {
	chainID := big.NewInt(1337)
	opts, err := KeyTransactOpts(chainID, privateKeyStr)
	if err != nil {
		return nil, errors.Wrapf(err, "")
	}

	address := opts.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: walletBalance,
		},
	}

	simulatedClient := backends.NewSimulatedBackend(genesisAlloc, gasLimit)

	return simulatedClient, nil
}
