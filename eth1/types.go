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

// description: 只能合约相关
// @author renshiwei
// Date: 2022/11/15 17:23

package eth1

type Tx struct {
	From    string `json:"from,omitempty"`
	To      string `json:"to"`
	Value   string `json:"value"`
	Data    string `json:"data"`
	ChainId int    `json:"chain_id"`
}
