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
