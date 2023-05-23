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
// Date: 2023/4/7 16:37

package beacon

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestValidatorsByPubKey_ByLocal(t *testing.T) {
	initClAddr()
	b := &BeaconService{
		baseUrl: clAddr,
		timeout: 1 * time.Minute,
	}

	validatorMap, err := b.ValidatorsByPubKey(context.Background(), "head", []string{
		"0x8000091c2ae64ee414a54c1cc1fc67dec663408bc636cb86756e0200e41a75c8f86603f104f02c856983d2783116be13",
		"0x80000e851c0f53c3246ff726d7ff7766661ca5e12a07c45c114d208d54f0f8233d4380b2e9aff759d69795d1df905526",
	})
	require.NoError(t, err)

	fmt.Printf("pubkey1:%+v\n", validatorMap["0x8000091c2ae64ee414a54c1cc1fc67dec663408bc636cb86756e0200e41a75c8f86603f104f02c856983d2783116be13"])
	fmt.Printf("pubkey2:%+v\n", validatorMap["0x80000e851c0f53c3246ff726d7ff7766661ca5e12a07c45c114d208d54f0f8233d4380b2e9aff759d69795d1df905526"])
}
