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

package util

import (
	"encoding/json"
	"sort"
)

func EqualUint64Array(arr1, arr2 []uint64) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	arr1Json, _ := json.Marshal(arr1)
	arr2Json, _ := json.Marshal(arr2)
	return string(arr1Json) == string(arr2Json)
}

func EqualUint64ArraySorted(arr1, arr2 []uint64) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	arr1Json, _ := json.Marshal(arr1)
	arr2Json, _ := json.Marshal(arr2)
	return string(arr1Json) == string(arr2Json)
}
