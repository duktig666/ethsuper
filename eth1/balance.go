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

// desc:
// @author renshiwei
// Date: 2023/4/6 16:28

package eth1

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
)

func (e *EthClient) BalanceAt(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error) {
	balance, err := e.Client.BalanceAt(ctx, common.HexToAddress(address), blockNumber)
	if err != nil {
		return decimal.Zero.BigInt(), errors.Wrapf(err, "Fail to get balance address:%s blockNumber:%s", address, blockNumber.String())
	}
	return balance, nil
}
