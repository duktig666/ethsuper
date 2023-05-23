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
// Date: 2023/4/6 13:06

package beacon

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetBeaconBlock_ByLocal(t *testing.T) {
	initClAddr()
	b := &BeaconService{
		baseUrl: clAddr,
		timeout: 1 * time.Minute,
	}

	beaconBlock, _, err := b.BeaconBlock(context.Background(), "head")
	require.NoError(t, err)

	marshal, err := json.Marshal(beaconBlock)
	require.NoError(t, err)

	fmt.Printf("beaconBlock:%s\n", string(marshal))
}

func TestGetExecutionPayload_ByLocal(t *testing.T) {
	initClAddr()
	b := &BeaconService{
		baseUrl: clAddr,
		timeout: 1 * time.Minute,
	}

	executionPayload, err := b.ExecutionPayload(context.Background(), "head")
	require.NoError(t, err)

	fmt.Printf("BlockNumber:%s\n", executionPayload.BlockNumber)
	fmt.Printf("Timestamp:%s\n", executionPayload.Timestamp)
	fmt.Printf("BlockHash:%s\n", executionPayload.BlockHash)
}

func TestGetExecutionBlock_ByLocal(t *testing.T) {
	initClAddr()
	b := &BeaconService{
		baseUrl: clAddr,
		timeout: 1 * time.Minute,
	}

	executionBlock, err := b.ExecutionBlock(context.Background(), "5354527")
	require.NoError(t, err)

	fmt.Printf("BlockNumber:%s\n", executionBlock.BlockNumber)
	fmt.Printf("Timestamp:%s\n", executionBlock.Timestamp)
	fmt.Printf("BlockHash:%s\n", executionBlock.BlockHash)
}
