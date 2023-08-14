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
// Date: 2023/8/14

package eth1

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEthClient_AddressType_ByLocal(t *testing.T) {
	ctx := context.Background()
	client, err := NewEthClient(ctx, "https://rpc.ankr.com/eth")
	require.NoError(t, err)

	addressType, err := client.AddressType(ctx, common.HexToAddress("0x892e7c8C5E716e17891ABf9395a0de1f2fc84786"))
	require.NoError(t, err)
	require.Equal(t, EOA, addressType)

	addressType2, err := client.AddressType(ctx, common.HexToAddress("0x58553F5c5a6AEE89EaBFd42c231A18aB0872700d"))
	require.NoError(t, err)
	require.Equal(t, CONTRACT, addressType2)
}
