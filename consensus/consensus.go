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
// Date: 2023/4/6 20:02

package consensus

import (
	"context"
	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/duktig666/ethsuper/consensus/beacon"
	"github.com/duktig666/ethsuper/consensus/chaintime/standard"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"time"
)

type Consensus struct {
	ConsensusClient        eth2client.Service
	ChainTimeService       *standard.Service
	CustomizeBeaconService beacon.BeaconProvider

	Timeout time.Duration
}

func New(ctx context.Context, addr string, timeout time.Duration) (*Consensus, error) {
	beaconNode, err := beacon.ConnectToBeaconNode(ctx, addr, timeout, true)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	beaconService, err := beacon.New(ctx, addr, timeout)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to connect to beacon service")
	}

	chainTimeService, err := standard.New(ctx,
		standard.WithGenesisTimeProvider(beaconNode.(eth2client.GenesisTimeProvider)),
		standard.WithSpecProvider(beaconNode.(eth2client.SpecProvider)),
		standard.WithLogLevel(zerolog.Disabled),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create chaintime service")
	}

	return &Consensus{
		ConsensusClient:        beaconNode,
		ChainTimeService:       chainTimeService,
		CustomizeBeaconService: beaconService,
		Timeout:                timeout,
	}, nil

}
