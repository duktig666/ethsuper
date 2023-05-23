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
// Date: 2023/4/6 19:50

package mock

import (
	"context"
	eth2client "github.com/attestantio/go-eth2-client"
	"time"
)

// GenesisTimeProvider is a mock for eth2client.GenesisTimeProvider.
type GenesisTimeProvider struct {
	genesisTime time.Time
}

// NewGenesisTimeProvider returns a mock genesis time provider with the provided value.
func NewGenesisTimeProvider(genesisTime time.Time) eth2client.GenesisTimeProvider {
	return &GenesisTimeProvider{
		genesisTime: genesisTime,
	}
}

// GenesisTime is a mock.
func (m *GenesisTimeProvider) GenesisTime(ctx context.Context) (time.Time, error) {
	return m.genesisTime, nil
}

// SpecProvider is a mock for eth2client.SpecProvider.
type SpecProvider struct {
	spec map[string]interface{}
}

// NewSpecProvider returns a mock spec provider with the provided values.
func NewSpecProvider(slotDuration time.Duration,
	slotsPerEpoch uint64,
	epochsPerSyncCommitteePeriod uint64,
) eth2client.SpecProvider {
	return &SpecProvider{
		spec: map[string]interface{}{
			"SECONDS_PER_SLOT":                 slotDuration,
			"SLOTS_PER_EPOCH":                  slotsPerEpoch,
			"EPOCHS_PER_SYNC_COMMITTEE_PERIOD": epochsPerSyncCommitteePeriod,
		},
	}
}

// Spec is a mock.
func (m *SpecProvider) Spec(ctx context.Context) (map[string]interface{}, error) {
	return m.spec, nil
}
