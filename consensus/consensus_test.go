// desc:
// @author renshiwei
// Date: 2023/4/6 20:29

package consensus

import (
	"context"
	"fmt"
	"github.com/duktig666/ethsuper/config"
	"github.com/duktig666/ethsuper/consensus/beacon"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetConsensusInfo_ByLocal(t *testing.T) {
	config.InitConfig("../conf/config-goerli.dev.yaml")
	clAddr := config.GlobalConfig.Eth.ClAddr
	timeout := 1 * time.Minute

	background := context.Background()

	consensusClient, err := New(background, clAddr, timeout)
	require.NoError(t, err)

	currentEpoch := consensusClient.ChainTimeService.CurrentEpoch()
	firstSlotOfCurrentEpoch := consensusClient.ChainTimeService.FirstSlotOfEpoch(currentEpoch)
	fmt.Printf("current epoch: %v\n", currentEpoch)
	fmt.Printf("first slot of current epoch: %v\n", firstSlotOfCurrentEpoch)
	require.Equal(t, currentEpoch, consensusClient.ChainTimeService.SlotToEpoch(firstSlotOfCurrentEpoch))

	block, _, err := consensusClient.CustomizeBeaconService.(beacon.BeaconBlocksProvider).BeaconBlock(background, "head")
	require.NoError(t, err)
	fmt.Println("head slot:", block.Data.Message.Slot)
}
