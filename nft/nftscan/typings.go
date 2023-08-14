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
// Date: 2023/7/19

package nftscan

type NFTContract struct {
	ID               int64   `gorm:"primary key ; auto increment" json:"id"`
	NftContractHash  string  `json:"nft_contract_hash"`
	CollectionName   string  `json:"collection_name"`
	CollectionSymbol string  `json:"collection_symbol"`
	TransactionHash  string  `json:"transaction_hash"`
	Blockchain       string  `json:"blockchain"`
	BlockHeight      uint64  `json:"block_height"` // 发现块高
	CreateTime       int64   `json:"create_time"`
	ErcType          string  `json:"erc_type"`
	LowestPrice24h   float64 `json:"lowest_price_24h"`
	HighestPrice24h  float64 `json:"highest_price_24h"`
	AveragePrice24h  float64 `json:"average_price_24h"`
	Volume24h        float64 `json:"volume_24h"`
	VolumeTotal      float64 `json:"volume_total"`
}

type NFTAsset struct {
	ID                  int64  `gorm:"primary key ; auto increment" json:"id"`
	NftAssetId          string `json:"nft_asset_id"`
	NftContractHash     string `json:"nft_contract_hash"`
	MintTransactionHash string `json:"mint_transaction_hash"`
	MintBlockHeight     uint64 `json:"mint_block_height"`
	MintTimeStamp       uint64 `json:"mint_time_stamp"`
	Creator             string `json:"creator"` // 发现块高
	Holder              string `json:"holder"`
	ImageUrl            string `json:"image_url"`
	Format              string `json:"format"`
	// todo burn
}

type NFTTransaction struct {
	ID                   int64   `gorm:"primary key ; auto increment" json:"id"`
	TransactionHash      string  `json:"transaction_hash"`
	BlockHeight          uint64  `json:"block_height"`
	TimeStamp            uint64  `json:"timestamp"`
	From                 string  `json:"from"`
	To                   string  `json:"to"`
	Value                uint64  `json:"value"` // 发现块高
	GasPrice             uint64  `json:"gas_price"`
	GasLimit             uint64  `json:"gas_limit"`
	GasUsedByTransaction uint64  `json:"gas_used_by_transaction"`
	TransactionFee       float64 `json:"transaction_fee"`
}

type NFTTransfer struct {
	ID              int64  `gorm:"primary key ; auto increment" json:"id"`
	NFTContractHash string `json:"nft_contract_hash"`
	NFTAssetId      string `json:"nft_asset_id"`
	TransactionHash string `json:"transaction_hash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Type            string `json:"type"`
}
