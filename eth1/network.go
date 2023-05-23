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

package eth1

var chainIDMap map[string]int

func init() {
	chainIDMap = make(map[string]int)
	chainIDMap[MAINNET] = 1
	chainIDMap[GOERLI] = 5
	chainIDMap[PRATER] = 5
	chainIDMap[SEPOLIA] = 11155111
}

func IsSameNetwork(network string, chainID int) bool {
	return chainIDMap[network] == chainID
}
