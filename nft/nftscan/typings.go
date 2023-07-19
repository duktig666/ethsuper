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

// description:
// @author renshiwei
// Date: 2023/7/19

package nftscan

type NFTContract struct {
	ID               int64   `gorm:"primary key ; auto increment" json:"id"`
	NftContractHash  string  `json:"nftContractHash"`
	CollectionName   string  `json:"collection_name"`
	CollectionSymbol string  `json:"collection_symbol"`
	TransactionHash  string  `json:"transaction_hash"`
	Blockchain       string  `json:"blockchain"`
	BlockHeight      uint64  `json:"block_height"` // 发现块高
	CreateTime       int64   `json:"create_time"`
	ErcType          string  `json:"erc_type"`
	LowestPrice24h   float64 `json:"lowestPrice24H"`
	HighestPrice24h  float64 `json:"highestPrice24H"`
	AveragePrice24h  float64 `json:"averagePrice24H"`
	Volume24h        float64 `json:"volume24H"`
	VolumeTotal      float64 `json:"volumeTotal"`
}

type NFTAsset struct {
	ID                  int64  `gorm:"primary key ; auto increment" json:"id"`
	NftAssetId          string `json:"nftAssetId"`
	NftContractHash     string `json:"nftContractHash"`
	MintTransactionHash string `json:"mintTransactionHash"`
	MintBlockHeight     uint64 `json:"mintBlockHeight"`
	MintTimeStamp       uint64 `json:"mintTimeStamp"`
	Creator             string `json:"creator"` // 发现块高
	Holder              string `json:"holder"`
	ImageUrl            string `json:"imageUrl"`
	Format              string `json:"format"`
	// todo burn
}

type NFTTransaction struct {
	ID                   int64   `gorm:"primary key ; auto increment" json:"id"`
	TransactionHash      string  `json:"transactionHash"`
	BlockHeight          uint64  `json:"blockHeight"`
	TimeStamp            uint64  `json:"timeStamp"`
	From                 string  `json:"from"`
	To                   string  `json:"to"`
	Value                uint64  `json:"value"` // 发现块高
	GasPrice             uint64  `json:"gasPrice"`
	GasLimit             uint64  `json:"gasLimit"`
	GasUsedByTransaction uint64  `json:"gasUsedByTransaction"`
	TransactionFee       float64 `json:"transactionFee"`
}

type NFTTransfer struct {
	ID              int64  `gorm:"primary key ; auto increment" json:"id"`
	NFTContractHash string `json:"NFTContractHash"`
	NFTAssetId      string `json:"NFTAssetId"`
	TransactionHash string `json:"transactionHash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Type            string `json:"type"`
}
