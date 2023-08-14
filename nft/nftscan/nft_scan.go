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
// Date: 2023/7/18

package nftscan

import (
	"context"
	"fmt"
	"github.com/duktig666/ethsuper/common/logger"
	"github.com/duktig666/ethsuper/contracts/erc721"
	"github.com/duktig666/ethsuper/eth1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math"
	"math/big"
	"strconv"
)

type NFTScan struct {
	ethClient *eth1.EthClient
	chainID   *big.Int
}

func NewNFTScan(ctx context.Context, elRpc string) (*NFTScan, error) {
	ethClient, err := eth1.NewEthClient(ctx, elRpc)
	if err != nil {
		return nil, errors.Wrap(err, "NewEthClient err.")
	}

	chainID, err := ethClient.Client.ChainID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "ChainID err.")
	}
	return &NFTScan{
		ethClient: ethClient,
		chainID:   chainID,
	}, nil
}

// ScanNFT Scan NFT data based on block intervals
func (v *NFTScan) ScanNFT(ctx context.Context, startBlockNumber, endBlockNumber *big.Int) error {
	curBlockNumber, err := v.ethClient.Client.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "curBlockNumber err.")
	}
	if curBlockNumber < startBlockNumber.Uint64() {
		return errors.Errorf("curBlockNumber:%v < startBlockNumber:%v", curBlockNumber, startBlockNumber)
	} else if curBlockNumber < endBlockNumber.Uint64() {
		endBlockNumber = big.NewInt(int64(curBlockNumber))
		logger.Warn("curBlockNumber > endBlockNumber. reset endBlockNumber.",
			zap.Uint64("curBlockNumber", curBlockNumber),
			zap.Uint64("endBlockNumber", endBlockNumber.Uint64()),
		)
	}

	err = v.scanBlock(ctx, startBlockNumber, endBlockNumber)
	if err != nil {
		return errors.Wrap(err, "scanBlock err.")
	}
	return nil
}

func (v *NFTScan) scanBlock(ctx context.Context, startBlockNumber, endBlockNumber *big.Int) error {
	scanBlockNumber := new(big.Int).Set(startBlockNumber)
	for {
		block, err := v.ethClient.Client.BlockByNumber(ctx, scanBlockNumber)
		if err != nil {
			return errors.Wrapf(err, "BlockByNumber err. blockNumber:%s", scanBlockNumber.String())
		}
		logger.Debug("==========Scan block tx...==========", zap.String("blockNumber", scanBlockNumber.String()))

		// Scan transactions for each block
		for _, tx := range block.Transactions() {
			// 如果交易的to是nil，则是合约创建的交易
			if tx.To() == nil {
				err = v.analyzeContractDeployTx(ctx, tx)
				if err != nil {
					logger.Errorf("analyzeContractDeployTx err:%+v", err)
					continue
				}
			}

			err = v.analyzeNftTransfer(ctx, tx, block)
			if err != nil {
				return errors.Wrap(err, "AnalyzeNftTransfer err")
			}
		}

		logger.Debug("==========Scan block tx end.==========", zap.String("blockNumber", scanBlockNumber.String()))

		// todo save processedBlockHeight
		scanBlockNumber = new(big.Int).Add(scanBlockNumber, big.NewInt(1))
		if scanBlockNumber.Cmp(endBlockNumber) > 0 {
			break
		}
	}

	return nil
}

func (v *NFTScan) analyzeContractDeployTx(ctx context.Context, tx *types.Transaction) error {
	receipt, err := v.ethClient.Client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return errors.Wrapf(err, "Get TransactionReceipt err. txHash:%s", tx.Hash().String())
	}

	// If the receipt's Contract Address is not address(0), it is a transaction created by the contract
	if receipt.ContractAddress.Hex() != eth1.ZERO_ADDRESS_STR {
		// Determine the ERC type of the contract address (ERC 721, ERC 1155)
		ercType, err := JudgmentErcType(receipt.ContractAddress, v.ethClient.Client)
		if err != nil {
			return errors.Wrapf(err, "JudgmentErcType err. contractAddress:%s txHash:%s", receipt.ContractAddress.Hex(), tx.Hash().String())
		}

		switch ercType {
		case Erc721:
			err := v.recordErc721(receipt.BlockNumber, receipt.ContractAddress, tx.Hash().Hex())
			if err != nil {
				return errors.Wrapf(err, "recordErc721 err. contractAddress:%s txHash:%s", receipt.ContractAddress.Hex(), tx.Hash().String())
			}
		}
	}

	return nil
}

// todo 扫描到 合约的创建
func (v *NFTScan) recordErc721(blockNumber *big.Int, address common.Address, tx string) error {
	erc721Contract, err := erc721.NewErc721(address, v.ethClient.Client)
	if err != nil {
		return errors.Wrap(err, "failed to NewErc721 contract.")
	}
	tokenName, err := erc721Contract.Name(nil)
	if err != nil {
		return errors.Wrap(err, "failed to get token name.")
	}

	tokenSymbol, err := erc721Contract.Symbol(nil)
	if err != nil {
		return errors.Wrap(err, "failed to get token symbol.")
	}

	logger.Debug("Scan erc contract deploy success.",
		zap.String("blockNumber", blockNumber.String()),
		zap.String("contractAddress", address.Hex()),
		zap.String("tokenName", tokenName),
		zap.String("tokenSymbol", tokenSymbol),
		zap.String("tx", tx),
	)

	return nil
}

func (v *NFTScan) analyzeNftTransfer(ctx context.Context, tx *types.Transaction, block *types.Block) error {
	receipt, err := v.ethClient.Client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return errors.Wrapf(err, "Get TransactionReceipt err. txHash:%s", tx.Hash().String())
	}

	flag := false
	for _, log := range receipt.Logs {
		if checkErc721Transfer(log) {
			flag = true
			logger.Debug("===Scan nft transfer success===")
			_ = v.recordNftTransfer(log)

			if checkType("0x"+log.Topics[1].String()[26:], "0x"+log.Topics[2].String()[26:]) == "mint" {
				_ = v.recordNft(log, tx, block)
			}
		}
	}

	if flag {
		_ = v.recordNftTransaction(receipt, block, tx)
	}

	return nil
}

// todo nft 交易特殊处理
func (v *NFTScan) recordNftTransfer(log *types.Log) error {
	nftContractHash := log.Address.String()
	nftAssetId := getTokenId(log.Topics[3].String()).String()
	txHash := log.TxHash.String()
	from := "0x" + log.Topics[1].String()[26:]
	to := "0x" + log.Topics[2].String()[26:]
	transferType := checkType(from, to)

	logger.Debug("Scan nft transfer success.",
		zap.String("transferType", transferType),
		zap.Uint64("blockNumber", log.BlockNumber),
		zap.String("nftContractHash", nftContractHash),
		zap.String("nftAssetId", nftAssetId),
		zap.String("txHash", txHash),
		zap.String("from", from),
		zap.String("to", to),
	)
	return nil
}

func (v *NFTScan) recordNft(log *types.Log, tx *types.Transaction, block *types.Block) error {
	nftAssetId := getTokenId(log.Topics[3].String()).String()
	mintTxHash := tx.Hash().String()
	mintBlockHeight := log.BlockNumber
	mintTimeStamp := block.Time()
	creator := "0x" + log.Topics[2].String()[26:]
	nftContractHash := log.Address.Hex()

	logger.Debug("Scan nft mint success.",
		zap.Uint64("blockNumber", log.BlockNumber),
		zap.String("nftAssetId", nftAssetId),
		zap.String("mintTxHash", mintTxHash),
		zap.Uint64("mintBlockHeight", mintBlockHeight),
		zap.Uint64("mintTimeStamp", mintTimeStamp),
		zap.String("creator", creator),
		zap.String("nftContractHash", nftContractHash),
	)

	return nil
}

func (v *NFTScan) recordNftTransaction(receipt *types.Receipt, block *types.Block, tx *types.Transaction) error {
	transactionHash := receipt.TxHash.String()
	blockHeight := receipt.BlockNumber.Uint64()
	timeStamp := block.Time()
	to := tx.To().String()
	value := tx.Value().Uint64()
	gasPrice := tx.GasPrice().Uint64()
	gasLimit := tx.Gas()
	gasUsedByTransaction := receipt.GasUsed
	transactionFee, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", math.Pow(10, -18)*float64(gasUsedByTransaction)*float64(gasPrice)), 64)

	logger.Debug("Scan nft transaction success.",
		zap.String("transactionHash", transactionHash),
		zap.Uint64("blockHeight", blockHeight),
		zap.Uint64("timeStamp", timeStamp),
		zap.String("to", to),
		zap.Uint64("value", value),
		zap.Uint64("gasPrice", gasPrice),
		zap.Uint64("gasLimit", gasLimit),
		zap.Uint64("gasUsedByTransaction", gasUsedByTransaction),
		zap.Float64("transactionFee", transactionFee),
	)

	return nil
}

func checkErc721Transfer(log *types.Log) bool {
	if len(log.Topics) == 4 && log.Topics[0].Hex() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {
		return true
	} else {
		return false
	}
}

func getTokenId(string16Hex string) *big.Int {
	string16 := string16Hex[2:]
	i := new(big.Int)
	i.SetString(string16, 16)
	return i
}

func checkType(from string, to string) string {
	if from == "0x0000000000000000000000000000000000000000" {
		return "mint"
	} else if to == "0x0000000000000000000000000000000000000000" {
		return "burn"
	} else {
		return "transfer"
	}
}
