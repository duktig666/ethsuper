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
// Date: 2023/6/30

package beacon

import (
	"context"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/duktig666/ethsuper/eth1"
	"github.com/pkg/errors"
	"math/big"
	"strconv"
)

// ValidatorIsFullExited Whether the validator has completed the lifecycle of the exit
func ValidatorIsFullExited(validator *consensusApi.Validator) bool {
	return validator.Status == consensusApi.ValidatorStateWithdrawalDone
}

// ValidatorIsRequestExiting Whether the validator request the exit
func ValidatorIsRequestExiting(validator *consensusApi.Validator) bool {
	return validator.Status == consensusApi.ValidatorStateActiveExiting ||
		validator.Status == consensusApi.ValidatorStateActiveSlashed ||
		validator.Status == consensusApi.ValidatorStateExitedUnslashed ||
		validator.Status == consensusApi.ValidatorStateExitedSlashed ||
		validator.Status == consensusApi.ValidatorStateWithdrawalPossible ||
		validator.Status == consensusApi.ValidatorStateWithdrawalDone
}

// ValidatorSlashedAmount 32ETh - ValidatorWithdrawAbleBalance
func (b *BeaconService) ValidatorSlashedAmount(ctx context.Context, validator *consensusApi.Validator) (*big.Int, error) {
	if !validator.Validator.Slashed {
		return nil, errors.New("validator not slashed.")
	}

	pubkey := validator.Validator.PublicKey.String()

	withdrawAbleSlot := epochToSlot(validator.Validator.WithdrawableEpoch)
	pubkeys := make([]string, 1)
	pubkeys[0] = pubkey
	slashedValidator, err := b.ValidatorsByPubKey(ctx, strconv.FormatInt(int64(withdrawAbleSlot), 10), pubkeys)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get validators info.")
	}

	slashAmount := new(big.Int).Sub(eth1.ETH32().BigInt(), eth1.GweiToWei(slashedValidator[pubkey].Balance))

	return slashAmount, nil
}

func epochToSlot(epoch phase0.Epoch) phase0.Slot {
	return phase0.Slot(uint64(epoch) * 12)
}
