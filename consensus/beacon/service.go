// desc:
// @author renshiwei
// Date: 2023/4/6 13:10

package beacon

import (
	"context"
	consensusClient "github.com/attestantio/go-eth2-client"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"math/big"
	"time"
)

//type BeaconError struct {
//	Message string `json:"message"`
//	Code    int    `json:"code"`
//}
//
//var beaconError *BeaconError

// BeaconBlocksProvider is the interface for providing beacon blocks.
type BeaconBlocksProvider interface {
	// BeaconBlock provides the block header of a given block ID.
	BeaconBlock(ctx context.Context, blockID string) (*BeaconBlock, bool, error)

	ExecutionPayload(ctx context.Context, blockID string) (*ExecutionPayload, error)

	ExecutionBlock(ctx context.Context, blockID string) (*ExecutionBlock, error)

	HeadSlot(ctx context.Context) (*big.Int, error)
}

// ExecutionInfoProvider is the interface for get ExecutionPayload.
type ExecutionInfoProvider interface {
	// ExecutionPayload provides the block's Execution info of a given block ID.
	ExecutionPayload(ctx context.Context, blockID string) (*ExecutionPayload, error)
}

type ValidatorsProvider interface {
	ValidatorsByPubKey(ctx context.Context, stateID string, validatorPubKeys []string) (map[string]*consensusApi.Validator, error)
}

type BeaconProvider interface {
	consensusClient.Service

	// Timeout returns the timeout of the client.
	Timeout() time.Duration
}

// BeaconService is the service for providing consensus layer.
type BeaconService struct {
	baseUrl string
	timeout time.Duration
}

// Name provides the name of the service.
func (b *BeaconService) Name() string {
	return "Customize Standard (HTTP)"
}

// Address provides the address for the connection.
func (b *BeaconService) Address() string {
	return b.baseUrl
}

// Timeout provides the timeout for the connection.
func (b *BeaconService) Timeout() time.Duration {
	return b.timeout
}

func New(ctx context.Context, addr string, timeout time.Duration) (BeaconProvider, error) {
	return &BeaconService{
		baseUrl: addr,
		timeout: timeout,
	}, nil
}
