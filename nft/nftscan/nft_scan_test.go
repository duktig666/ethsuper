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
// Date: 2023/7/18

package nftscan

import (
	"context"
	"github.com/duktig666/ethsuper/common/logger"
	"github.com/duktig666/ethsuper/config"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestNFTScan_ScanNFT(t *testing.T) {
	var err error
	ctx := context.Background()
	config.InitConfig("../../conf/config-mainnet.dev.yaml")
	logger.InitLog("debug", "console")

	nftScan, err := NewNFTScan(ctx, config.GlobalConfig.Eth.ElAddr)
	require.NoError(t, err)

	startNumber := big.NewInt(17718800)
	endNumber := big.NewInt(17719800)
	err = nftScan.ScanNFT(ctx, startNumber, endNumber)
	require.NoError(t, err)
}
