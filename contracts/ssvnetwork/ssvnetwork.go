// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ssvnetwork

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ISSVNetworkCoreCluster is an auto generated low-level Go binding around an user-defined struct.
type ISSVNetworkCoreCluster struct {
	ValidatorCount  uint32
	NetworkFeeIndex uint64
	Index           uint64
	Balance         *big.Int
	Active          bool
}

// ISSVNetworkCoreSnapshot is an auto generated low-level Go binding around an user-defined struct.
type ISSVNetworkCoreSnapshot struct {
	Block   uint64
	Index   uint64
	Balance uint64
}

// SsvnetworkMetaData contains all meta data concerning the Ssvnetwork contract.
var SsvnetworkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ApprovalNotWithinTimeframe\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterDoesNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterIsLiquidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterNotLiquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedValidatorLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeExceedsIncreaseLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeIncreaseNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectClusterState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorIdsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewBlockPeriodIsBelowMinimum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeDelcared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameFeeChangeNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsortedOperatorsList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorOwnedByOtherAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterReactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"DeclareOperatorFeePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"ExecuteOperatorFeePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"FeeRecipientAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"LiquidationThresholdPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinimumLiquidationCollateralUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"NetworkEarningsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"NetworkFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"OperatorFeeCancellationDeclared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeDeclared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"OperatorFeeIncreaseLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whitelisted\",\"type\":\"address\"}],\"name\":\"OperatorWhitelistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"OperatorWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"shares\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"cancelDeclaredOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"clusters\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"balance\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"block\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"declareOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"declareOperatorFeePeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"executeOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executeOperatorFeePeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"initialVersion_\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"operatorMaxFeeIncrease_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"declareOperatorFeePeriod_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executeOperatorFeePeriod_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minimumBlocksBeforeLiquidation_\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minimumLiquidationCollateral_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumBlocksBeforeLiquidation\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumLiquidationCollateral\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"network\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"networkFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndexBlockNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"operatorFeeChangeRequests\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"approvalBeginTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"approvalEndTime\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorMaxFeeIncrease\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"block\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"balance\",\"type\":\"uint64\"}],\"internalType\":\"structISSVNetworkCore.Snapshot\",\"name\":\"snapshot\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"operatorsWhitelist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"reactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"reduceOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"registerOperator\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes\",\"name\":\"shares\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"setFeeRecipientAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"whitelisted\",\"type\":\"address\"}],\"name\":\"setOperatorWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDeclareOperatorFeePeriod\",\"type\":\"uint64\"}],\"name\":\"updateDeclareOperatorFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newExecuteOperatorFeePeriod\",\"type\":\"uint64\"}],\"name\":\"updateExecuteOperatorFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blocks\",\"type\":\"uint64\"}],\"name\":\"updateLiquidationThresholdPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinimumLiquidationCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateNetworkFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newOperatorMaxFeeIncrease\",\"type\":\"uint64\"}],\"name\":\"updateOperatorFeeIncreaseLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"validatorPKs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorsPerOperatorLimit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNetworkEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawOperatorEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"withdrawOperatorEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SsvnetworkABI is the input ABI used to generate the binding from.
// Deprecated: Use SsvnetworkMetaData.ABI instead.
var SsvnetworkABI = SsvnetworkMetaData.ABI

// Ssvnetwork is an auto generated Go binding around an Ethereum contract.
type Ssvnetwork struct {
	SsvnetworkCaller     // Read-only binding to the contract
	SsvnetworkTransactor // Write-only binding to the contract
	SsvnetworkFilterer   // Log filterer for contract events
}

// SsvnetworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type SsvnetworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvnetworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SsvnetworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvnetworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SsvnetworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvnetworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SsvnetworkSession struct {
	Contract     *Ssvnetwork       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SsvnetworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SsvnetworkCallerSession struct {
	Contract *SsvnetworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SsvnetworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SsvnetworkTransactorSession struct {
	Contract     *SsvnetworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SsvnetworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type SsvnetworkRaw struct {
	Contract *Ssvnetwork // Generic contract binding to access the raw methods on
}

// SsvnetworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SsvnetworkCallerRaw struct {
	Contract *SsvnetworkCaller // Generic read-only contract binding to access the raw methods on
}

// SsvnetworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SsvnetworkTransactorRaw struct {
	Contract *SsvnetworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSsvnetwork creates a new instance of Ssvnetwork, bound to a specific deployed contract.
func NewSsvnetwork(address common.Address, backend bind.ContractBackend) (*Ssvnetwork, error) {
	contract, err := bindSsvnetwork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ssvnetwork{SsvnetworkCaller: SsvnetworkCaller{contract: contract}, SsvnetworkTransactor: SsvnetworkTransactor{contract: contract}, SsvnetworkFilterer: SsvnetworkFilterer{contract: contract}}, nil
}

// NewSsvnetworkCaller creates a new read-only instance of Ssvnetwork, bound to a specific deployed contract.
func NewSsvnetworkCaller(address common.Address, caller bind.ContractCaller) (*SsvnetworkCaller, error) {
	contract, err := bindSsvnetwork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkCaller{contract: contract}, nil
}

// NewSsvnetworkTransactor creates a new write-only instance of Ssvnetwork, bound to a specific deployed contract.
func NewSsvnetworkTransactor(address common.Address, transactor bind.ContractTransactor) (*SsvnetworkTransactor, error) {
	contract, err := bindSsvnetwork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkTransactor{contract: contract}, nil
}

// NewSsvnetworkFilterer creates a new log filterer instance of Ssvnetwork, bound to a specific deployed contract.
func NewSsvnetworkFilterer(address common.Address, filterer bind.ContractFilterer) (*SsvnetworkFilterer, error) {
	contract, err := bindSsvnetwork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkFilterer{contract: contract}, nil
}

// bindSsvnetwork binds a generic wrapper to an already deployed contract.
func bindSsvnetwork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SsvnetworkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ssvnetwork *SsvnetworkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ssvnetwork.Contract.SsvnetworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ssvnetwork *SsvnetworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.SsvnetworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ssvnetwork *SsvnetworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.SsvnetworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ssvnetwork *SsvnetworkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ssvnetwork.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ssvnetwork *SsvnetworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ssvnetwork *SsvnetworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.contract.Transact(opts, method, params...)
}

// Clusters is a free data retrieval call binding the contract method 0x2e41e2e8.
//
// Solidity: function clusters(bytes32 ) view returns(bytes32)
func (_Ssvnetwork *SsvnetworkCaller) Clusters(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "clusters", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Clusters is a free data retrieval call binding the contract method 0x2e41e2e8.
//
// Solidity: function clusters(bytes32 ) view returns(bytes32)
func (_Ssvnetwork *SsvnetworkSession) Clusters(arg0 [32]byte) ([32]byte, error) {
	return _Ssvnetwork.Contract.Clusters(&_Ssvnetwork.CallOpts, arg0)
}

// Clusters is a free data retrieval call binding the contract method 0x2e41e2e8.
//
// Solidity: function clusters(bytes32 ) view returns(bytes32)
func (_Ssvnetwork *SsvnetworkCallerSession) Clusters(arg0 [32]byte) ([32]byte, error) {
	return _Ssvnetwork.Contract.Clusters(&_Ssvnetwork.CallOpts, arg0)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(uint32 validatorCount, uint64 balance, uint64 block)
func (_Ssvnetwork *SsvnetworkCaller) Dao(opts *bind.CallOpts) (struct {
	ValidatorCount uint32
	Balance        uint64
	Block          uint64
}, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "dao")

	outstruct := new(struct {
		ValidatorCount uint32
		Balance        uint64
		Block          uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidatorCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Balance = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Block = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(uint32 validatorCount, uint64 balance, uint64 block)
func (_Ssvnetwork *SsvnetworkSession) Dao() (struct {
	ValidatorCount uint32
	Balance        uint64
	Block          uint64
}, error) {
	return _Ssvnetwork.Contract.Dao(&_Ssvnetwork.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(uint32 validatorCount, uint64 balance, uint64 block)
func (_Ssvnetwork *SsvnetworkCallerSession) Dao() (struct {
	ValidatorCount uint32
	Balance        uint64
	Block          uint64
}, error) {
	return _Ssvnetwork.Contract.Dao(&_Ssvnetwork.CallOpts)
}

// DeclareOperatorFeePeriod is a free data retrieval call binding the contract method 0xb2802f4f.
//
// Solidity: function declareOperatorFeePeriod() view returns(uint64)
func (_Ssvnetwork *SsvnetworkCaller) DeclareOperatorFeePeriod(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "declareOperatorFeePeriod")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DeclareOperatorFeePeriod is a free data retrieval call binding the contract method 0xb2802f4f.
//
// Solidity: function declareOperatorFeePeriod() view returns(uint64)
func (_Ssvnetwork *SsvnetworkSession) DeclareOperatorFeePeriod() (uint64, error) {
	return _Ssvnetwork.Contract.DeclareOperatorFeePeriod(&_Ssvnetwork.CallOpts)
}

// DeclareOperatorFeePeriod is a free data retrieval call binding the contract method 0xb2802f4f.
//
// Solidity: function declareOperatorFeePeriod() view returns(uint64)
func (_Ssvnetwork *SsvnetworkCallerSession) DeclareOperatorFeePeriod() (uint64, error) {
	return _Ssvnetwork.Contract.DeclareOperatorFeePeriod(&_Ssvnetwork.CallOpts)
}

// ExecuteOperatorFeePeriod is a free data retrieval call binding the contract method 0xac4f2873.
//
// Solidity: function executeOperatorFeePeriod() view returns(uint64)
func (_Ssvnetwork *SsvnetworkCaller) ExecuteOperatorFeePeriod(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "executeOperatorFeePeriod")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ExecuteOperatorFeePeriod is a free data retrieval call binding the contract method 0xac4f2873.
//
// Solidity: function executeOperatorFeePeriod() view returns(uint64)
func (_Ssvnetwork *SsvnetworkSession) ExecuteOperatorFeePeriod() (uint64, error) {
	return _Ssvnetwork.Contract.ExecuteOperatorFeePeriod(&_Ssvnetwork.CallOpts)
}

// ExecuteOperatorFeePeriod is a free data retrieval call binding the contract method 0xac4f2873.
//
// Solidity: function executeOperatorFeePeriod() view returns(uint64)
func (_Ssvnetwork *SsvnetworkCallerSession) ExecuteOperatorFeePeriod() (uint64, error) {
	return _Ssvnetwork.Contract.ExecuteOperatorFeePeriod(&_Ssvnetwork.CallOpts)
}

// MinimumBlocksBeforeLiquidation is a free data retrieval call binding the contract method 0xa8e89916.
//
// Solidity: function minimumBlocksBeforeLiquidation() view returns(uint64)
func (_Ssvnetwork *SsvnetworkCaller) MinimumBlocksBeforeLiquidation(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "minimumBlocksBeforeLiquidation")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinimumBlocksBeforeLiquidation is a free data retrieval call binding the contract method 0xa8e89916.
//
// Solidity: function minimumBlocksBeforeLiquidation() view returns(uint64)
func (_Ssvnetwork *SsvnetworkSession) MinimumBlocksBeforeLiquidation() (uint64, error) {
	return _Ssvnetwork.Contract.MinimumBlocksBeforeLiquidation(&_Ssvnetwork.CallOpts)
}

// MinimumBlocksBeforeLiquidation is a free data retrieval call binding the contract method 0xa8e89916.
//
// Solidity: function minimumBlocksBeforeLiquidation() view returns(uint64)
func (_Ssvnetwork *SsvnetworkCallerSession) MinimumBlocksBeforeLiquidation() (uint64, error) {
	return _Ssvnetwork.Contract.MinimumBlocksBeforeLiquidation(&_Ssvnetwork.CallOpts)
}

// MinimumLiquidationCollateral is a free data retrieval call binding the contract method 0xd70cae7e.
//
// Solidity: function minimumLiquidationCollateral() view returns(uint64)
func (_Ssvnetwork *SsvnetworkCaller) MinimumLiquidationCollateral(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "minimumLiquidationCollateral")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinimumLiquidationCollateral is a free data retrieval call binding the contract method 0xd70cae7e.
//
// Solidity: function minimumLiquidationCollateral() view returns(uint64)
func (_Ssvnetwork *SsvnetworkSession) MinimumLiquidationCollateral() (uint64, error) {
	return _Ssvnetwork.Contract.MinimumLiquidationCollateral(&_Ssvnetwork.CallOpts)
}

// MinimumLiquidationCollateral is a free data retrieval call binding the contract method 0xd70cae7e.
//
// Solidity: function minimumLiquidationCollateral() view returns(uint64)
func (_Ssvnetwork *SsvnetworkCallerSession) MinimumLiquidationCollateral() (uint64, error) {
	return _Ssvnetwork.Contract.MinimumLiquidationCollateral(&_Ssvnetwork.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint64 networkFee, uint64 networkFeeIndex, uint64 networkFeeIndexBlockNumber)
func (_Ssvnetwork *SsvnetworkCaller) Network(opts *bind.CallOpts) (struct {
	NetworkFee                 uint64
	NetworkFeeIndex            uint64
	NetworkFeeIndexBlockNumber uint64
}, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "network")

	outstruct := new(struct {
		NetworkFee                 uint64
		NetworkFeeIndex            uint64
		NetworkFeeIndexBlockNumber uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NetworkFee = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.NetworkFeeIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.NetworkFeeIndexBlockNumber = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint64 networkFee, uint64 networkFeeIndex, uint64 networkFeeIndexBlockNumber)
func (_Ssvnetwork *SsvnetworkSession) Network() (struct {
	NetworkFee                 uint64
	NetworkFeeIndex            uint64
	NetworkFeeIndexBlockNumber uint64
}, error) {
	return _Ssvnetwork.Contract.Network(&_Ssvnetwork.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint64 networkFee, uint64 networkFeeIndex, uint64 networkFeeIndexBlockNumber)
func (_Ssvnetwork *SsvnetworkCallerSession) Network() (struct {
	NetworkFee                 uint64
	NetworkFeeIndex            uint64
	NetworkFeeIndexBlockNumber uint64
}, error) {
	return _Ssvnetwork.Contract.Network(&_Ssvnetwork.CallOpts)
}

// OperatorFeeChangeRequests is a free data retrieval call binding the contract method 0x1bc1d9d3.
//
// Solidity: function operatorFeeChangeRequests(uint64 ) view returns(uint64 fee, uint64 approvalBeginTime, uint64 approvalEndTime)
func (_Ssvnetwork *SsvnetworkCaller) OperatorFeeChangeRequests(opts *bind.CallOpts, arg0 uint64) (struct {
	Fee               uint64
	ApprovalBeginTime uint64
	ApprovalEndTime   uint64
}, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "operatorFeeChangeRequests", arg0)

	outstruct := new(struct {
		Fee               uint64
		ApprovalBeginTime uint64
		ApprovalEndTime   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fee = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ApprovalBeginTime = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ApprovalEndTime = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// OperatorFeeChangeRequests is a free data retrieval call binding the contract method 0x1bc1d9d3.
//
// Solidity: function operatorFeeChangeRequests(uint64 ) view returns(uint64 fee, uint64 approvalBeginTime, uint64 approvalEndTime)
func (_Ssvnetwork *SsvnetworkSession) OperatorFeeChangeRequests(arg0 uint64) (struct {
	Fee               uint64
	ApprovalBeginTime uint64
	ApprovalEndTime   uint64
}, error) {
	return _Ssvnetwork.Contract.OperatorFeeChangeRequests(&_Ssvnetwork.CallOpts, arg0)
}

// OperatorFeeChangeRequests is a free data retrieval call binding the contract method 0x1bc1d9d3.
//
// Solidity: function operatorFeeChangeRequests(uint64 ) view returns(uint64 fee, uint64 approvalBeginTime, uint64 approvalEndTime)
func (_Ssvnetwork *SsvnetworkCallerSession) OperatorFeeChangeRequests(arg0 uint64) (struct {
	Fee               uint64
	ApprovalBeginTime uint64
	ApprovalEndTime   uint64
}, error) {
	return _Ssvnetwork.Contract.OperatorFeeChangeRequests(&_Ssvnetwork.CallOpts, arg0)
}

// OperatorMaxFeeIncrease is a free data retrieval call binding the contract method 0x9ccfd315.
//
// Solidity: function operatorMaxFeeIncrease() view returns(uint64)
func (_Ssvnetwork *SsvnetworkCaller) OperatorMaxFeeIncrease(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "operatorMaxFeeIncrease")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OperatorMaxFeeIncrease is a free data retrieval call binding the contract method 0x9ccfd315.
//
// Solidity: function operatorMaxFeeIncrease() view returns(uint64)
func (_Ssvnetwork *SsvnetworkSession) OperatorMaxFeeIncrease() (uint64, error) {
	return _Ssvnetwork.Contract.OperatorMaxFeeIncrease(&_Ssvnetwork.CallOpts)
}

// OperatorMaxFeeIncrease is a free data retrieval call binding the contract method 0x9ccfd315.
//
// Solidity: function operatorMaxFeeIncrease() view returns(uint64)
func (_Ssvnetwork *SsvnetworkCallerSession) OperatorMaxFeeIncrease() (uint64, error) {
	return _Ssvnetwork.Contract.OperatorMaxFeeIncrease(&_Ssvnetwork.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x97f988dc.
//
// Solidity: function operators(uint64 ) view returns(address owner, uint64 fee, uint32 validatorCount, (uint64,uint64,uint64) snapshot)
func (_Ssvnetwork *SsvnetworkCaller) Operators(opts *bind.CallOpts, arg0 uint64) (struct {
	Owner          common.Address
	Fee            uint64
	ValidatorCount uint32
	Snapshot       ISSVNetworkCoreSnapshot
}, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "operators", arg0)

	outstruct := new(struct {
		Owner          common.Address
		Fee            uint64
		ValidatorCount uint32
		Snapshot       ISSVNetworkCoreSnapshot
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ValidatorCount = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Snapshot = *abi.ConvertType(out[3], new(ISSVNetworkCoreSnapshot)).(*ISSVNetworkCoreSnapshot)

	return *outstruct, err

}

// Operators is a free data retrieval call binding the contract method 0x97f988dc.
//
// Solidity: function operators(uint64 ) view returns(address owner, uint64 fee, uint32 validatorCount, (uint64,uint64,uint64) snapshot)
func (_Ssvnetwork *SsvnetworkSession) Operators(arg0 uint64) (struct {
	Owner          common.Address
	Fee            uint64
	ValidatorCount uint32
	Snapshot       ISSVNetworkCoreSnapshot
}, error) {
	return _Ssvnetwork.Contract.Operators(&_Ssvnetwork.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x97f988dc.
//
// Solidity: function operators(uint64 ) view returns(address owner, uint64 fee, uint32 validatorCount, (uint64,uint64,uint64) snapshot)
func (_Ssvnetwork *SsvnetworkCallerSession) Operators(arg0 uint64) (struct {
	Owner          common.Address
	Fee            uint64
	ValidatorCount uint32
	Snapshot       ISSVNetworkCoreSnapshot
}, error) {
	return _Ssvnetwork.Contract.Operators(&_Ssvnetwork.CallOpts, arg0)
}

// OperatorsWhitelist is a free data retrieval call binding the contract method 0x3ee7dd4e.
//
// Solidity: function operatorsWhitelist(uint64 ) view returns(address)
func (_Ssvnetwork *SsvnetworkCaller) OperatorsWhitelist(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "operatorsWhitelist", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorsWhitelist is a free data retrieval call binding the contract method 0x3ee7dd4e.
//
// Solidity: function operatorsWhitelist(uint64 ) view returns(address)
func (_Ssvnetwork *SsvnetworkSession) OperatorsWhitelist(arg0 uint64) (common.Address, error) {
	return _Ssvnetwork.Contract.OperatorsWhitelist(&_Ssvnetwork.CallOpts, arg0)
}

// OperatorsWhitelist is a free data retrieval call binding the contract method 0x3ee7dd4e.
//
// Solidity: function operatorsWhitelist(uint64 ) view returns(address)
func (_Ssvnetwork *SsvnetworkCallerSession) OperatorsWhitelist(arg0 uint64) (common.Address, error) {
	return _Ssvnetwork.Contract.OperatorsWhitelist(&_Ssvnetwork.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ssvnetwork *SsvnetworkCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ssvnetwork *SsvnetworkSession) Owner() (common.Address, error) {
	return _Ssvnetwork.Contract.Owner(&_Ssvnetwork.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ssvnetwork *SsvnetworkCallerSession) Owner() (common.Address, error) {
	return _Ssvnetwork.Contract.Owner(&_Ssvnetwork.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ssvnetwork *SsvnetworkCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ssvnetwork *SsvnetworkSession) PendingOwner() (common.Address, error) {
	return _Ssvnetwork.Contract.PendingOwner(&_Ssvnetwork.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ssvnetwork *SsvnetworkCallerSession) PendingOwner() (common.Address, error) {
	return _Ssvnetwork.Contract.PendingOwner(&_Ssvnetwork.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ssvnetwork *SsvnetworkCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ssvnetwork *SsvnetworkSession) ProxiableUUID() ([32]byte, error) {
	return _Ssvnetwork.Contract.ProxiableUUID(&_Ssvnetwork.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ssvnetwork *SsvnetworkCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Ssvnetwork.Contract.ProxiableUUID(&_Ssvnetwork.CallOpts)
}

// ValidatorPKs is a free data retrieval call binding the contract method 0x350cfc55.
//
// Solidity: function validatorPKs(bytes32 ) view returns(address owner, bool active)
func (_Ssvnetwork *SsvnetworkCaller) ValidatorPKs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner  common.Address
	Active bool
}, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "validatorPKs", arg0)

	outstruct := new(struct {
		Owner  common.Address
		Active bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Active = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ValidatorPKs is a free data retrieval call binding the contract method 0x350cfc55.
//
// Solidity: function validatorPKs(bytes32 ) view returns(address owner, bool active)
func (_Ssvnetwork *SsvnetworkSession) ValidatorPKs(arg0 [32]byte) (struct {
	Owner  common.Address
	Active bool
}, error) {
	return _Ssvnetwork.Contract.ValidatorPKs(&_Ssvnetwork.CallOpts, arg0)
}

// ValidatorPKs is a free data retrieval call binding the contract method 0x350cfc55.
//
// Solidity: function validatorPKs(bytes32 ) view returns(address owner, bool active)
func (_Ssvnetwork *SsvnetworkCallerSession) ValidatorPKs(arg0 [32]byte) (struct {
	Owner  common.Address
	Active bool
}, error) {
	return _Ssvnetwork.Contract.ValidatorPKs(&_Ssvnetwork.CallOpts, arg0)
}

// ValidatorsPerOperatorLimit is a free data retrieval call binding the contract method 0x08ee2e9b.
//
// Solidity: function validatorsPerOperatorLimit() view returns(uint32)
func (_Ssvnetwork *SsvnetworkCaller) ValidatorsPerOperatorLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "validatorsPerOperatorLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ValidatorsPerOperatorLimit is a free data retrieval call binding the contract method 0x08ee2e9b.
//
// Solidity: function validatorsPerOperatorLimit() view returns(uint32)
func (_Ssvnetwork *SsvnetworkSession) ValidatorsPerOperatorLimit() (uint32, error) {
	return _Ssvnetwork.Contract.ValidatorsPerOperatorLimit(&_Ssvnetwork.CallOpts)
}

// ValidatorsPerOperatorLimit is a free data retrieval call binding the contract method 0x08ee2e9b.
//
// Solidity: function validatorsPerOperatorLimit() view returns(uint32)
func (_Ssvnetwork *SsvnetworkCallerSession) ValidatorsPerOperatorLimit() (uint32, error) {
	return _Ssvnetwork.Contract.ValidatorsPerOperatorLimit(&_Ssvnetwork.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(bytes32)
func (_Ssvnetwork *SsvnetworkCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ssvnetwork.contract.Call(opts, &out, "version")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(bytes32)
func (_Ssvnetwork *SsvnetworkSession) Version() ([32]byte, error) {
	return _Ssvnetwork.Contract.Version(&_Ssvnetwork.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(bytes32)
func (_Ssvnetwork *SsvnetworkCallerSession) Version() ([32]byte, error) {
	return _Ssvnetwork.Contract.Version(&_Ssvnetwork.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ssvnetwork *SsvnetworkTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ssvnetwork *SsvnetworkSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ssvnetwork.Contract.AcceptOwnership(&_Ssvnetwork.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ssvnetwork.Contract.AcceptOwnership(&_Ssvnetwork.TransactOpts)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkTransactor) CancelDeclaredOperatorFee(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "cancelDeclaredOperatorFee", operatorId)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkSession) CancelDeclaredOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.CancelDeclaredOperatorFee(&_Ssvnetwork.TransactOpts, operatorId)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) CancelDeclaredOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.CancelDeclaredOperatorFee(&_Ssvnetwork.TransactOpts, operatorId)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssvnetwork *SsvnetworkTransactor) DeclareOperatorFee(opts *bind.TransactOpts, operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "declareOperatorFee", operatorId, fee)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssvnetwork *SsvnetworkSession) DeclareOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.DeclareOperatorFee(&_Ssvnetwork.TransactOpts, operatorId, fee)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) DeclareOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.DeclareOperatorFee(&_Ssvnetwork.TransactOpts, operatorId, fee)
}

// Deposit is a paid mutator transaction binding the contract method 0xc3dcc9f4.
//
// Solidity: function deposit(address owner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactor) Deposit(opts *bind.TransactOpts, owner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "deposit", owner, operatorIds, amount, cluster)
}

// Deposit is a paid mutator transaction binding the contract method 0xc3dcc9f4.
//
// Solidity: function deposit(address owner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkSession) Deposit(owner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.Deposit(&_Ssvnetwork.TransactOpts, owner, operatorIds, amount, cluster)
}

// Deposit is a paid mutator transaction binding the contract method 0xc3dcc9f4.
//
// Solidity: function deposit(address owner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) Deposit(owner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.Deposit(&_Ssvnetwork.TransactOpts, owner, operatorIds, amount, cluster)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkTransactor) ExecuteOperatorFee(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "executeOperatorFee", operatorId)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkSession) ExecuteOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.ExecuteOperatorFee(&_Ssvnetwork.TransactOpts, operatorId)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) ExecuteOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.ExecuteOperatorFee(&_Ssvnetwork.TransactOpts, operatorId)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb7e6ec5.
//
// Solidity: function initialize(string initialVersion_, address token_, uint64 operatorMaxFeeIncrease_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_) returns()
func (_Ssvnetwork *SsvnetworkTransactor) Initialize(opts *bind.TransactOpts, initialVersion_ string, token_ common.Address, operatorMaxFeeIncrease_ uint64, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "initialize", initialVersion_, token_, operatorMaxFeeIncrease_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb7e6ec5.
//
// Solidity: function initialize(string initialVersion_, address token_, uint64 operatorMaxFeeIncrease_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_) returns()
func (_Ssvnetwork *SsvnetworkSession) Initialize(initialVersion_ string, token_ common.Address, operatorMaxFeeIncrease_ uint64, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.Initialize(&_Ssvnetwork.TransactOpts, initialVersion_, token_, operatorMaxFeeIncrease_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb7e6ec5.
//
// Solidity: function initialize(string initialVersion_, address token_, uint64 operatorMaxFeeIncrease_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) Initialize(initialVersion_ string, token_ common.Address, operatorMaxFeeIncrease_ uint64, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.Initialize(&_Ssvnetwork.TransactOpts, initialVersion_, token_, operatorMaxFeeIncrease_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_)
}

// Liquidate is a paid mutator transaction binding the contract method 0xe6a77817.
//
// Solidity: function liquidate(address owner, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactor) Liquidate(opts *bind.TransactOpts, owner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "liquidate", owner, operatorIds, cluster)
}

// Liquidate is a paid mutator transaction binding the contract method 0xe6a77817.
//
// Solidity: function liquidate(address owner, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkSession) Liquidate(owner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.Liquidate(&_Ssvnetwork.TransactOpts, owner, operatorIds, cluster)
}

// Liquidate is a paid mutator transaction binding the contract method 0xe6a77817.
//
// Solidity: function liquidate(address owner, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) Liquidate(owner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.Liquidate(&_Ssvnetwork.TransactOpts, owner, operatorIds, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0xe2baf194.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactor) Reactivate(opts *bind.TransactOpts, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "reactivate", operatorIds, amount, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0xe2baf194.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkSession) Reactivate(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.Reactivate(&_Ssvnetwork.TransactOpts, operatorIds, amount, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0xe2baf194.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) Reactivate(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.Reactivate(&_Ssvnetwork.TransactOpts, operatorIds, amount, cluster)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssvnetwork *SsvnetworkTransactor) ReduceOperatorFee(opts *bind.TransactOpts, operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "reduceOperatorFee", operatorId, fee)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssvnetwork *SsvnetworkSession) ReduceOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.ReduceOperatorFee(&_Ssvnetwork.TransactOpts, operatorId, fee)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) ReduceOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.ReduceOperatorFee(&_Ssvnetwork.TransactOpts, operatorId, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_Ssvnetwork *SsvnetworkTransactor) RegisterOperator(opts *bind.TransactOpts, publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "registerOperator", publicKey, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_Ssvnetwork *SsvnetworkSession) RegisterOperator(publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.RegisterOperator(&_Ssvnetwork.TransactOpts, publicKey, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_Ssvnetwork *SsvnetworkTransactorSession) RegisterOperator(publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.RegisterOperator(&_Ssvnetwork.TransactOpts, publicKey, fee)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x967ef807.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes shares, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactor) RegisterValidator(opts *bind.TransactOpts, publicKey []byte, operatorIds []uint64, shares []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "registerValidator", publicKey, operatorIds, shares, amount, cluster)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x967ef807.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes shares, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkSession) RegisterValidator(publicKey []byte, operatorIds []uint64, shares []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.RegisterValidator(&_Ssvnetwork.TransactOpts, publicKey, operatorIds, shares, amount, cluster)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x967ef807.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes shares, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) RegisterValidator(publicKey []byte, operatorIds []uint64, shares []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.RegisterValidator(&_Ssvnetwork.TransactOpts, publicKey, operatorIds, shares, amount, cluster)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkTransactor) RemoveOperator(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "removeOperator", operatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkSession) RemoveOperator(operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.RemoveOperator(&_Ssvnetwork.TransactOpts, operatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) RemoveOperator(operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.RemoveOperator(&_Ssvnetwork.TransactOpts, operatorId)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0xfc2ecd80.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactor) RemoveValidator(opts *bind.TransactOpts, publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "removeValidator", publicKey, operatorIds, cluster)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0xfc2ecd80.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkSession) RemoveValidator(publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.RemoveValidator(&_Ssvnetwork.TransactOpts, publicKey, operatorIds, cluster)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0xfc2ecd80.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) RemoveValidator(publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.RemoveValidator(&_Ssvnetwork.TransactOpts, publicKey, operatorIds, cluster)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ssvnetwork *SsvnetworkTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ssvnetwork *SsvnetworkSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ssvnetwork.Contract.RenounceOwnership(&_Ssvnetwork.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ssvnetwork.Contract.RenounceOwnership(&_Ssvnetwork.TransactOpts)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_Ssvnetwork *SsvnetworkTransactor) SetFeeRecipientAddress(opts *bind.TransactOpts, recipientAddress common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "setFeeRecipientAddress", recipientAddress)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_Ssvnetwork *SsvnetworkSession) SetFeeRecipientAddress(recipientAddress common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.SetFeeRecipientAddress(&_Ssvnetwork.TransactOpts, recipientAddress)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) SetFeeRecipientAddress(recipientAddress common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.SetFeeRecipientAddress(&_Ssvnetwork.TransactOpts, recipientAddress)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_Ssvnetwork *SsvnetworkTransactor) SetOperatorWhitelist(opts *bind.TransactOpts, operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "setOperatorWhitelist", operatorId, whitelisted)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_Ssvnetwork *SsvnetworkSession) SetOperatorWhitelist(operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.SetOperatorWhitelist(&_Ssvnetwork.TransactOpts, operatorId, whitelisted)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) SetOperatorWhitelist(operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.SetOperatorWhitelist(&_Ssvnetwork.TransactOpts, operatorId, whitelisted)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ssvnetwork *SsvnetworkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ssvnetwork *SsvnetworkSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.TransferOwnership(&_Ssvnetwork.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.TransferOwnership(&_Ssvnetwork.TransactOpts, newOwner)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 newDeclareOperatorFeePeriod) returns()
func (_Ssvnetwork *SsvnetworkTransactor) UpdateDeclareOperatorFeePeriod(opts *bind.TransactOpts, newDeclareOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "updateDeclareOperatorFeePeriod", newDeclareOperatorFeePeriod)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 newDeclareOperatorFeePeriod) returns()
func (_Ssvnetwork *SsvnetworkSession) UpdateDeclareOperatorFeePeriod(newDeclareOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateDeclareOperatorFeePeriod(&_Ssvnetwork.TransactOpts, newDeclareOperatorFeePeriod)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 newDeclareOperatorFeePeriod) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) UpdateDeclareOperatorFeePeriod(newDeclareOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateDeclareOperatorFeePeriod(&_Ssvnetwork.TransactOpts, newDeclareOperatorFeePeriod)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 newExecuteOperatorFeePeriod) returns()
func (_Ssvnetwork *SsvnetworkTransactor) UpdateExecuteOperatorFeePeriod(opts *bind.TransactOpts, newExecuteOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "updateExecuteOperatorFeePeriod", newExecuteOperatorFeePeriod)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 newExecuteOperatorFeePeriod) returns()
func (_Ssvnetwork *SsvnetworkSession) UpdateExecuteOperatorFeePeriod(newExecuteOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateExecuteOperatorFeePeriod(&_Ssvnetwork.TransactOpts, newExecuteOperatorFeePeriod)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 newExecuteOperatorFeePeriod) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) UpdateExecuteOperatorFeePeriod(newExecuteOperatorFeePeriod uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateExecuteOperatorFeePeriod(&_Ssvnetwork.TransactOpts, newExecuteOperatorFeePeriod)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_Ssvnetwork *SsvnetworkTransactor) UpdateLiquidationThresholdPeriod(opts *bind.TransactOpts, blocks uint64) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "updateLiquidationThresholdPeriod", blocks)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_Ssvnetwork *SsvnetworkSession) UpdateLiquidationThresholdPeriod(blocks uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateLiquidationThresholdPeriod(&_Ssvnetwork.TransactOpts, blocks)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) UpdateLiquidationThresholdPeriod(blocks uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateLiquidationThresholdPeriod(&_Ssvnetwork.TransactOpts, blocks)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_Ssvnetwork *SsvnetworkTransactor) UpdateMinimumLiquidationCollateral(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "updateMinimumLiquidationCollateral", amount)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_Ssvnetwork *SsvnetworkSession) UpdateMinimumLiquidationCollateral(amount *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateMinimumLiquidationCollateral(&_Ssvnetwork.TransactOpts, amount)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) UpdateMinimumLiquidationCollateral(amount *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateMinimumLiquidationCollateral(&_Ssvnetwork.TransactOpts, amount)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_Ssvnetwork *SsvnetworkTransactor) UpdateNetworkFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "updateNetworkFee", fee)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_Ssvnetwork *SsvnetworkSession) UpdateNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateNetworkFee(&_Ssvnetwork.TransactOpts, fee)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) UpdateNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateNetworkFee(&_Ssvnetwork.TransactOpts, fee)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 newOperatorMaxFeeIncrease) returns()
func (_Ssvnetwork *SsvnetworkTransactor) UpdateOperatorFeeIncreaseLimit(opts *bind.TransactOpts, newOperatorMaxFeeIncrease uint64) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "updateOperatorFeeIncreaseLimit", newOperatorMaxFeeIncrease)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 newOperatorMaxFeeIncrease) returns()
func (_Ssvnetwork *SsvnetworkSession) UpdateOperatorFeeIncreaseLimit(newOperatorMaxFeeIncrease uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateOperatorFeeIncreaseLimit(&_Ssvnetwork.TransactOpts, newOperatorMaxFeeIncrease)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 newOperatorMaxFeeIncrease) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) UpdateOperatorFeeIncreaseLimit(newOperatorMaxFeeIncrease uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpdateOperatorFeeIncreaseLimit(&_Ssvnetwork.TransactOpts, newOperatorMaxFeeIncrease)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Ssvnetwork *SsvnetworkTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Ssvnetwork *SsvnetworkSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpgradeTo(&_Ssvnetwork.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpgradeTo(&_Ssvnetwork.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ssvnetwork *SsvnetworkTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ssvnetwork *SsvnetworkSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpgradeToAndCall(&_Ssvnetwork.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.UpgradeToAndCall(&_Ssvnetwork.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4f6d984a.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactor) Withdraw(opts *bind.TransactOpts, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "withdraw", operatorIds, amount, cluster)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4f6d984a.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkSession) Withdraw(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.Withdraw(&_Ssvnetwork.TransactOpts, operatorIds, amount, cluster)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4f6d984a.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,uint256,bool) cluster) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) Withdraw(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.Withdraw(&_Ssvnetwork.TransactOpts, operatorIds, amount, cluster)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_Ssvnetwork *SsvnetworkTransactor) WithdrawNetworkEarnings(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "withdrawNetworkEarnings", amount)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_Ssvnetwork *SsvnetworkSession) WithdrawNetworkEarnings(amount *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.WithdrawNetworkEarnings(&_Ssvnetwork.TransactOpts, amount)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) WithdrawNetworkEarnings(amount *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.WithdrawNetworkEarnings(&_Ssvnetwork.TransactOpts, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_Ssvnetwork *SsvnetworkTransactor) WithdrawOperatorEarnings(opts *bind.TransactOpts, operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "withdrawOperatorEarnings", operatorId, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_Ssvnetwork *SsvnetworkSession) WithdrawOperatorEarnings(operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.WithdrawOperatorEarnings(&_Ssvnetwork.TransactOpts, operatorId, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) WithdrawOperatorEarnings(operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.WithdrawOperatorEarnings(&_Ssvnetwork.TransactOpts, operatorId, amount)
}

// WithdrawOperatorEarnings0 is a paid mutator transaction binding the contract method 0xc8c876c6.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkTransactor) WithdrawOperatorEarnings0(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.contract.Transact(opts, "withdrawOperatorEarnings0", operatorId)
}

// WithdrawOperatorEarnings0 is a paid mutator transaction binding the contract method 0xc8c876c6.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkSession) WithdrawOperatorEarnings0(operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.WithdrawOperatorEarnings0(&_Ssvnetwork.TransactOpts, operatorId)
}

// WithdrawOperatorEarnings0 is a paid mutator transaction binding the contract method 0xc8c876c6.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId) returns()
func (_Ssvnetwork *SsvnetworkTransactorSession) WithdrawOperatorEarnings0(operatorId uint64) (*types.Transaction, error) {
	return _Ssvnetwork.Contract.WithdrawOperatorEarnings0(&_Ssvnetwork.TransactOpts, operatorId)
}

// SsvnetworkAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Ssvnetwork contract.
type SsvnetworkAdminChangedIterator struct {
	Event *SsvnetworkAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkAdminChanged represents a AdminChanged event raised by the Ssvnetwork contract.
type SsvnetworkAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Ssvnetwork *SsvnetworkFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*SsvnetworkAdminChangedIterator, error) {

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SsvnetworkAdminChangedIterator{contract: _Ssvnetwork.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Ssvnetwork *SsvnetworkFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SsvnetworkAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkAdminChanged)
				if err := _Ssvnetwork.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Ssvnetwork *SsvnetworkFilterer) ParseAdminChanged(log types.Log) (*SsvnetworkAdminChanged, error) {
	event := new(SsvnetworkAdminChanged)
	if err := _Ssvnetwork.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Ssvnetwork contract.
type SsvnetworkBeaconUpgradedIterator struct {
	Event *SsvnetworkBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkBeaconUpgraded represents a BeaconUpgraded event raised by the Ssvnetwork contract.
type SsvnetworkBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Ssvnetwork *SsvnetworkFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SsvnetworkBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkBeaconUpgradedIterator{contract: _Ssvnetwork.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Ssvnetwork *SsvnetworkFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SsvnetworkBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkBeaconUpgraded)
				if err := _Ssvnetwork.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Ssvnetwork *SsvnetworkFilterer) ParseBeaconUpgraded(log types.Log) (*SsvnetworkBeaconUpgraded, error) {
	event := new(SsvnetworkBeaconUpgraded)
	if err := _Ssvnetwork.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkClusterDepositedIterator is returned from FilterClusterDeposited and is used to iterate over the raw logs and unpacked data for ClusterDeposited events raised by the Ssvnetwork contract.
type SsvnetworkClusterDepositedIterator struct {
	Event *SsvnetworkClusterDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkClusterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkClusterDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkClusterDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkClusterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkClusterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkClusterDeposited represents a ClusterDeposited event raised by the Ssvnetwork contract.
type SsvnetworkClusterDeposited struct {
	Owner       common.Address
	OperatorIds []uint64
	Value       *big.Int
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterDeposited is a free log retrieval operation binding the contract event 0x8ad175cb337f2878bf0c228738f102d387d7be2d47dc1953da383d9fbf1ed55b.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) FilterClusterDeposited(opts *bind.FilterOpts, owner []common.Address) (*SsvnetworkClusterDepositedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "ClusterDeposited", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkClusterDepositedIterator{contract: _Ssvnetwork.contract, event: "ClusterDeposited", logs: logs, sub: sub}, nil
}

// WatchClusterDeposited is a free log subscription operation binding the contract event 0x8ad175cb337f2878bf0c228738f102d387d7be2d47dc1953da383d9fbf1ed55b.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) WatchClusterDeposited(opts *bind.WatchOpts, sink chan<- *SsvnetworkClusterDeposited, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "ClusterDeposited", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkClusterDeposited)
				if err := _Ssvnetwork.contract.UnpackLog(event, "ClusterDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClusterDeposited is a log parse operation binding the contract event 0x8ad175cb337f2878bf0c228738f102d387d7be2d47dc1953da383d9fbf1ed55b.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) ParseClusterDeposited(log types.Log) (*SsvnetworkClusterDeposited, error) {
	event := new(SsvnetworkClusterDeposited)
	if err := _Ssvnetwork.contract.UnpackLog(event, "ClusterDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkClusterLiquidatedIterator is returned from FilterClusterLiquidated and is used to iterate over the raw logs and unpacked data for ClusterLiquidated events raised by the Ssvnetwork contract.
type SsvnetworkClusterLiquidatedIterator struct {
	Event *SsvnetworkClusterLiquidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkClusterLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkClusterLiquidated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkClusterLiquidated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkClusterLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkClusterLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkClusterLiquidated represents a ClusterLiquidated event raised by the Ssvnetwork contract.
type SsvnetworkClusterLiquidated struct {
	Owner       common.Address
	OperatorIds []uint64
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterLiquidated is a free log retrieval operation binding the contract event 0xa1934b9fd1b01348c2c4add7fdf17276e2169476b9a3448a3b2005492a6284b6.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) FilterClusterLiquidated(opts *bind.FilterOpts, owner []common.Address) (*SsvnetworkClusterLiquidatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "ClusterLiquidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkClusterLiquidatedIterator{contract: _Ssvnetwork.contract, event: "ClusterLiquidated", logs: logs, sub: sub}, nil
}

// WatchClusterLiquidated is a free log subscription operation binding the contract event 0xa1934b9fd1b01348c2c4add7fdf17276e2169476b9a3448a3b2005492a6284b6.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) WatchClusterLiquidated(opts *bind.WatchOpts, sink chan<- *SsvnetworkClusterLiquidated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "ClusterLiquidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkClusterLiquidated)
				if err := _Ssvnetwork.contract.UnpackLog(event, "ClusterLiquidated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClusterLiquidated is a log parse operation binding the contract event 0xa1934b9fd1b01348c2c4add7fdf17276e2169476b9a3448a3b2005492a6284b6.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) ParseClusterLiquidated(log types.Log) (*SsvnetworkClusterLiquidated, error) {
	event := new(SsvnetworkClusterLiquidated)
	if err := _Ssvnetwork.contract.UnpackLog(event, "ClusterLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkClusterReactivatedIterator is returned from FilterClusterReactivated and is used to iterate over the raw logs and unpacked data for ClusterReactivated events raised by the Ssvnetwork contract.
type SsvnetworkClusterReactivatedIterator struct {
	Event *SsvnetworkClusterReactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkClusterReactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkClusterReactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkClusterReactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkClusterReactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkClusterReactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkClusterReactivated represents a ClusterReactivated event raised by the Ssvnetwork contract.
type SsvnetworkClusterReactivated struct {
	Owner       common.Address
	OperatorIds []uint64
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterReactivated is a free log retrieval operation binding the contract event 0x4728af0c29bfcecc7b936d4b253139a1abb52e9d00b65bc7661a58356301bd52.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) FilterClusterReactivated(opts *bind.FilterOpts, owner []common.Address) (*SsvnetworkClusterReactivatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "ClusterReactivated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkClusterReactivatedIterator{contract: _Ssvnetwork.contract, event: "ClusterReactivated", logs: logs, sub: sub}, nil
}

// WatchClusterReactivated is a free log subscription operation binding the contract event 0x4728af0c29bfcecc7b936d4b253139a1abb52e9d00b65bc7661a58356301bd52.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) WatchClusterReactivated(opts *bind.WatchOpts, sink chan<- *SsvnetworkClusterReactivated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "ClusterReactivated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkClusterReactivated)
				if err := _Ssvnetwork.contract.UnpackLog(event, "ClusterReactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClusterReactivated is a log parse operation binding the contract event 0x4728af0c29bfcecc7b936d4b253139a1abb52e9d00b65bc7661a58356301bd52.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) ParseClusterReactivated(log types.Log) (*SsvnetworkClusterReactivated, error) {
	event := new(SsvnetworkClusterReactivated)
	if err := _Ssvnetwork.contract.UnpackLog(event, "ClusterReactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkClusterWithdrawnIterator is returned from FilterClusterWithdrawn and is used to iterate over the raw logs and unpacked data for ClusterWithdrawn events raised by the Ssvnetwork contract.
type SsvnetworkClusterWithdrawnIterator struct {
	Event *SsvnetworkClusterWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkClusterWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkClusterWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkClusterWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkClusterWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkClusterWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkClusterWithdrawn represents a ClusterWithdrawn event raised by the Ssvnetwork contract.
type SsvnetworkClusterWithdrawn struct {
	Owner       common.Address
	OperatorIds []uint64
	Value       *big.Int
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterWithdrawn is a free log retrieval operation binding the contract event 0x27434ba64e254f0675ae2adca8d3a485bb16ce3134f988859e747df217ee7f13.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) FilterClusterWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*SsvnetworkClusterWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "ClusterWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkClusterWithdrawnIterator{contract: _Ssvnetwork.contract, event: "ClusterWithdrawn", logs: logs, sub: sub}, nil
}

// WatchClusterWithdrawn is a free log subscription operation binding the contract event 0x27434ba64e254f0675ae2adca8d3a485bb16ce3134f988859e747df217ee7f13.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) WatchClusterWithdrawn(opts *bind.WatchOpts, sink chan<- *SsvnetworkClusterWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "ClusterWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkClusterWithdrawn)
				if err := _Ssvnetwork.contract.UnpackLog(event, "ClusterWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClusterWithdrawn is a log parse operation binding the contract event 0x27434ba64e254f0675ae2adca8d3a485bb16ce3134f988859e747df217ee7f13.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) ParseClusterWithdrawn(log types.Log) (*SsvnetworkClusterWithdrawn, error) {
	event := new(SsvnetworkClusterWithdrawn)
	if err := _Ssvnetwork.contract.UnpackLog(event, "ClusterWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkDeclareOperatorFeePeriodUpdatedIterator is returned from FilterDeclareOperatorFeePeriodUpdated and is used to iterate over the raw logs and unpacked data for DeclareOperatorFeePeriodUpdated events raised by the Ssvnetwork contract.
type SsvnetworkDeclareOperatorFeePeriodUpdatedIterator struct {
	Event *SsvnetworkDeclareOperatorFeePeriodUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkDeclareOperatorFeePeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkDeclareOperatorFeePeriodUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkDeclareOperatorFeePeriodUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkDeclareOperatorFeePeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkDeclareOperatorFeePeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkDeclareOperatorFeePeriodUpdated represents a DeclareOperatorFeePeriodUpdated event raised by the Ssvnetwork contract.
type SsvnetworkDeclareOperatorFeePeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeclareOperatorFeePeriodUpdated is a free log retrieval operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) FilterDeclareOperatorFeePeriodUpdated(opts *bind.FilterOpts) (*SsvnetworkDeclareOperatorFeePeriodUpdatedIterator, error) {

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "DeclareOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvnetworkDeclareOperatorFeePeriodUpdatedIterator{contract: _Ssvnetwork.contract, event: "DeclareOperatorFeePeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDeclareOperatorFeePeriodUpdated is a free log subscription operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) WatchDeclareOperatorFeePeriodUpdated(opts *bind.WatchOpts, sink chan<- *SsvnetworkDeclareOperatorFeePeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "DeclareOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkDeclareOperatorFeePeriodUpdated)
				if err := _Ssvnetwork.contract.UnpackLog(event, "DeclareOperatorFeePeriodUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeclareOperatorFeePeriodUpdated is a log parse operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) ParseDeclareOperatorFeePeriodUpdated(log types.Log) (*SsvnetworkDeclareOperatorFeePeriodUpdated, error) {
	event := new(SsvnetworkDeclareOperatorFeePeriodUpdated)
	if err := _Ssvnetwork.contract.UnpackLog(event, "DeclareOperatorFeePeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkExecuteOperatorFeePeriodUpdatedIterator is returned from FilterExecuteOperatorFeePeriodUpdated and is used to iterate over the raw logs and unpacked data for ExecuteOperatorFeePeriodUpdated events raised by the Ssvnetwork contract.
type SsvnetworkExecuteOperatorFeePeriodUpdatedIterator struct {
	Event *SsvnetworkExecuteOperatorFeePeriodUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkExecuteOperatorFeePeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkExecuteOperatorFeePeriodUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkExecuteOperatorFeePeriodUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkExecuteOperatorFeePeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkExecuteOperatorFeePeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkExecuteOperatorFeePeriodUpdated represents a ExecuteOperatorFeePeriodUpdated event raised by the Ssvnetwork contract.
type SsvnetworkExecuteOperatorFeePeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExecuteOperatorFeePeriodUpdated is a free log retrieval operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) FilterExecuteOperatorFeePeriodUpdated(opts *bind.FilterOpts) (*SsvnetworkExecuteOperatorFeePeriodUpdatedIterator, error) {

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "ExecuteOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvnetworkExecuteOperatorFeePeriodUpdatedIterator{contract: _Ssvnetwork.contract, event: "ExecuteOperatorFeePeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchExecuteOperatorFeePeriodUpdated is a free log subscription operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) WatchExecuteOperatorFeePeriodUpdated(opts *bind.WatchOpts, sink chan<- *SsvnetworkExecuteOperatorFeePeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "ExecuteOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkExecuteOperatorFeePeriodUpdated)
				if err := _Ssvnetwork.contract.UnpackLog(event, "ExecuteOperatorFeePeriodUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuteOperatorFeePeriodUpdated is a log parse operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) ParseExecuteOperatorFeePeriodUpdated(log types.Log) (*SsvnetworkExecuteOperatorFeePeriodUpdated, error) {
	event := new(SsvnetworkExecuteOperatorFeePeriodUpdated)
	if err := _Ssvnetwork.contract.UnpackLog(event, "ExecuteOperatorFeePeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkFeeRecipientAddressUpdatedIterator is returned from FilterFeeRecipientAddressUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientAddressUpdated events raised by the Ssvnetwork contract.
type SsvnetworkFeeRecipientAddressUpdatedIterator struct {
	Event *SsvnetworkFeeRecipientAddressUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkFeeRecipientAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkFeeRecipientAddressUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkFeeRecipientAddressUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkFeeRecipientAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkFeeRecipientAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkFeeRecipientAddressUpdated represents a FeeRecipientAddressUpdated event raised by the Ssvnetwork contract.
type SsvnetworkFeeRecipientAddressUpdated struct {
	Owner            common.Address
	RecipientAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientAddressUpdated is a free log retrieval operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_Ssvnetwork *SsvnetworkFilterer) FilterFeeRecipientAddressUpdated(opts *bind.FilterOpts, owner []common.Address) (*SsvnetworkFeeRecipientAddressUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "FeeRecipientAddressUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkFeeRecipientAddressUpdatedIterator{contract: _Ssvnetwork.contract, event: "FeeRecipientAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientAddressUpdated is a free log subscription operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_Ssvnetwork *SsvnetworkFilterer) WatchFeeRecipientAddressUpdated(opts *bind.WatchOpts, sink chan<- *SsvnetworkFeeRecipientAddressUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "FeeRecipientAddressUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkFeeRecipientAddressUpdated)
				if err := _Ssvnetwork.contract.UnpackLog(event, "FeeRecipientAddressUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeRecipientAddressUpdated is a log parse operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_Ssvnetwork *SsvnetworkFilterer) ParseFeeRecipientAddressUpdated(log types.Log) (*SsvnetworkFeeRecipientAddressUpdated, error) {
	event := new(SsvnetworkFeeRecipientAddressUpdated)
	if err := _Ssvnetwork.contract.UnpackLog(event, "FeeRecipientAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Ssvnetwork contract.
type SsvnetworkInitializedIterator struct {
	Event *SsvnetworkInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkInitialized represents a Initialized event raised by the Ssvnetwork contract.
type SsvnetworkInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ssvnetwork *SsvnetworkFilterer) FilterInitialized(opts *bind.FilterOpts) (*SsvnetworkInitializedIterator, error) {

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SsvnetworkInitializedIterator{contract: _Ssvnetwork.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ssvnetwork *SsvnetworkFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SsvnetworkInitialized) (event.Subscription, error) {

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkInitialized)
				if err := _Ssvnetwork.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ssvnetwork *SsvnetworkFilterer) ParseInitialized(log types.Log) (*SsvnetworkInitialized, error) {
	event := new(SsvnetworkInitialized)
	if err := _Ssvnetwork.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkLiquidationThresholdPeriodUpdatedIterator is returned from FilterLiquidationThresholdPeriodUpdated and is used to iterate over the raw logs and unpacked data for LiquidationThresholdPeriodUpdated events raised by the Ssvnetwork contract.
type SsvnetworkLiquidationThresholdPeriodUpdatedIterator struct {
	Event *SsvnetworkLiquidationThresholdPeriodUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkLiquidationThresholdPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkLiquidationThresholdPeriodUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkLiquidationThresholdPeriodUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkLiquidationThresholdPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkLiquidationThresholdPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkLiquidationThresholdPeriodUpdated represents a LiquidationThresholdPeriodUpdated event raised by the Ssvnetwork contract.
type SsvnetworkLiquidationThresholdPeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidationThresholdPeriodUpdated is a free log retrieval operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) FilterLiquidationThresholdPeriodUpdated(opts *bind.FilterOpts) (*SsvnetworkLiquidationThresholdPeriodUpdatedIterator, error) {

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "LiquidationThresholdPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvnetworkLiquidationThresholdPeriodUpdatedIterator{contract: _Ssvnetwork.contract, event: "LiquidationThresholdPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidationThresholdPeriodUpdated is a free log subscription operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) WatchLiquidationThresholdPeriodUpdated(opts *bind.WatchOpts, sink chan<- *SsvnetworkLiquidationThresholdPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "LiquidationThresholdPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkLiquidationThresholdPeriodUpdated)
				if err := _Ssvnetwork.contract.UnpackLog(event, "LiquidationThresholdPeriodUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidationThresholdPeriodUpdated is a log parse operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) ParseLiquidationThresholdPeriodUpdated(log types.Log) (*SsvnetworkLiquidationThresholdPeriodUpdated, error) {
	event := new(SsvnetworkLiquidationThresholdPeriodUpdated)
	if err := _Ssvnetwork.contract.UnpackLog(event, "LiquidationThresholdPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkMinimumLiquidationCollateralUpdatedIterator is returned from FilterMinimumLiquidationCollateralUpdated and is used to iterate over the raw logs and unpacked data for MinimumLiquidationCollateralUpdated events raised by the Ssvnetwork contract.
type SsvnetworkMinimumLiquidationCollateralUpdatedIterator struct {
	Event *SsvnetworkMinimumLiquidationCollateralUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkMinimumLiquidationCollateralUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkMinimumLiquidationCollateralUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkMinimumLiquidationCollateralUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkMinimumLiquidationCollateralUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkMinimumLiquidationCollateralUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkMinimumLiquidationCollateralUpdated represents a MinimumLiquidationCollateralUpdated event raised by the Ssvnetwork contract.
type SsvnetworkMinimumLiquidationCollateralUpdated struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinimumLiquidationCollateralUpdated is a free log retrieval operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_Ssvnetwork *SsvnetworkFilterer) FilterMinimumLiquidationCollateralUpdated(opts *bind.FilterOpts) (*SsvnetworkMinimumLiquidationCollateralUpdatedIterator, error) {

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "MinimumLiquidationCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvnetworkMinimumLiquidationCollateralUpdatedIterator{contract: _Ssvnetwork.contract, event: "MinimumLiquidationCollateralUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumLiquidationCollateralUpdated is a free log subscription operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_Ssvnetwork *SsvnetworkFilterer) WatchMinimumLiquidationCollateralUpdated(opts *bind.WatchOpts, sink chan<- *SsvnetworkMinimumLiquidationCollateralUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "MinimumLiquidationCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkMinimumLiquidationCollateralUpdated)
				if err := _Ssvnetwork.contract.UnpackLog(event, "MinimumLiquidationCollateralUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinimumLiquidationCollateralUpdated is a log parse operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_Ssvnetwork *SsvnetworkFilterer) ParseMinimumLiquidationCollateralUpdated(log types.Log) (*SsvnetworkMinimumLiquidationCollateralUpdated, error) {
	event := new(SsvnetworkMinimumLiquidationCollateralUpdated)
	if err := _Ssvnetwork.contract.UnpackLog(event, "MinimumLiquidationCollateralUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkNetworkEarningsWithdrawnIterator is returned from FilterNetworkEarningsWithdrawn and is used to iterate over the raw logs and unpacked data for NetworkEarningsWithdrawn events raised by the Ssvnetwork contract.
type SsvnetworkNetworkEarningsWithdrawnIterator struct {
	Event *SsvnetworkNetworkEarningsWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkNetworkEarningsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkNetworkEarningsWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkNetworkEarningsWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkNetworkEarningsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkNetworkEarningsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkNetworkEarningsWithdrawn represents a NetworkEarningsWithdrawn event raised by the Ssvnetwork contract.
type SsvnetworkNetworkEarningsWithdrawn struct {
	Value     *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNetworkEarningsWithdrawn is a free log retrieval operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_Ssvnetwork *SsvnetworkFilterer) FilterNetworkEarningsWithdrawn(opts *bind.FilterOpts) (*SsvnetworkNetworkEarningsWithdrawnIterator, error) {

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "NetworkEarningsWithdrawn")
	if err != nil {
		return nil, err
	}
	return &SsvnetworkNetworkEarningsWithdrawnIterator{contract: _Ssvnetwork.contract, event: "NetworkEarningsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNetworkEarningsWithdrawn is a free log subscription operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_Ssvnetwork *SsvnetworkFilterer) WatchNetworkEarningsWithdrawn(opts *bind.WatchOpts, sink chan<- *SsvnetworkNetworkEarningsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "NetworkEarningsWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkNetworkEarningsWithdrawn)
				if err := _Ssvnetwork.contract.UnpackLog(event, "NetworkEarningsWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNetworkEarningsWithdrawn is a log parse operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_Ssvnetwork *SsvnetworkFilterer) ParseNetworkEarningsWithdrawn(log types.Log) (*SsvnetworkNetworkEarningsWithdrawn, error) {
	event := new(SsvnetworkNetworkEarningsWithdrawn)
	if err := _Ssvnetwork.contract.UnpackLog(event, "NetworkEarningsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkNetworkFeeUpdatedIterator is returned from FilterNetworkFeeUpdated and is used to iterate over the raw logs and unpacked data for NetworkFeeUpdated events raised by the Ssvnetwork contract.
type SsvnetworkNetworkFeeUpdatedIterator struct {
	Event *SsvnetworkNetworkFeeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkNetworkFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkNetworkFeeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkNetworkFeeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkNetworkFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkNetworkFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkNetworkFeeUpdated represents a NetworkFeeUpdated event raised by the Ssvnetwork contract.
type SsvnetworkNetworkFeeUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNetworkFeeUpdated is a free log retrieval operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_Ssvnetwork *SsvnetworkFilterer) FilterNetworkFeeUpdated(opts *bind.FilterOpts) (*SsvnetworkNetworkFeeUpdatedIterator, error) {

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "NetworkFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvnetworkNetworkFeeUpdatedIterator{contract: _Ssvnetwork.contract, event: "NetworkFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchNetworkFeeUpdated is a free log subscription operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_Ssvnetwork *SsvnetworkFilterer) WatchNetworkFeeUpdated(opts *bind.WatchOpts, sink chan<- *SsvnetworkNetworkFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "NetworkFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkNetworkFeeUpdated)
				if err := _Ssvnetwork.contract.UnpackLog(event, "NetworkFeeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNetworkFeeUpdated is a log parse operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_Ssvnetwork *SsvnetworkFilterer) ParseNetworkFeeUpdated(log types.Log) (*SsvnetworkNetworkFeeUpdated, error) {
	event := new(SsvnetworkNetworkFeeUpdated)
	if err := _Ssvnetwork.contract.UnpackLog(event, "NetworkFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the Ssvnetwork contract.
type SsvnetworkOperatorAddedIterator struct {
	Event *SsvnetworkOperatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkOperatorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkOperatorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkOperatorAdded represents a OperatorAdded event raised by the Ssvnetwork contract.
type SsvnetworkOperatorAdded struct {
	OperatorId uint64
	Owner      common.Address
	PublicKey  []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_Ssvnetwork *SsvnetworkFilterer) FilterOperatorAdded(opts *bind.FilterOpts, operatorId []uint64, owner []common.Address) (*SsvnetworkOperatorAddedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "OperatorAdded", operatorIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkOperatorAddedIterator{contract: _Ssvnetwork.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_Ssvnetwork *SsvnetworkFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *SsvnetworkOperatorAdded, operatorId []uint64, owner []common.Address) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "OperatorAdded", operatorIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkOperatorAdded)
				if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorAdded is a log parse operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_Ssvnetwork *SsvnetworkFilterer) ParseOperatorAdded(log types.Log) (*SsvnetworkOperatorAdded, error) {
	event := new(SsvnetworkOperatorAdded)
	if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkOperatorFeeCancellationDeclaredIterator is returned from FilterOperatorFeeCancellationDeclared and is used to iterate over the raw logs and unpacked data for OperatorFeeCancellationDeclared events raised by the Ssvnetwork contract.
type SsvnetworkOperatorFeeCancellationDeclaredIterator struct {
	Event *SsvnetworkOperatorFeeCancellationDeclared // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkOperatorFeeCancellationDeclaredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkOperatorFeeCancellationDeclared)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkOperatorFeeCancellationDeclared)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkOperatorFeeCancellationDeclaredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkOperatorFeeCancellationDeclaredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkOperatorFeeCancellationDeclared represents a OperatorFeeCancellationDeclared event raised by the Ssvnetwork contract.
type SsvnetworkOperatorFeeCancellationDeclared struct {
	Owner      common.Address
	OperatorId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeCancellationDeclared is a free log retrieval operation binding the contract event 0x269e5b30b1f6e64e296fff57d56c77a4eb77fda7421146f1497f567539ba8906.
//
// Solidity: event OperatorFeeCancellationDeclared(address indexed owner, uint64 indexed operatorId)
func (_Ssvnetwork *SsvnetworkFilterer) FilterOperatorFeeCancellationDeclared(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvnetworkOperatorFeeCancellationDeclaredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "OperatorFeeCancellationDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkOperatorFeeCancellationDeclaredIterator{contract: _Ssvnetwork.contract, event: "OperatorFeeCancellationDeclared", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeCancellationDeclared is a free log subscription operation binding the contract event 0x269e5b30b1f6e64e296fff57d56c77a4eb77fda7421146f1497f567539ba8906.
//
// Solidity: event OperatorFeeCancellationDeclared(address indexed owner, uint64 indexed operatorId)
func (_Ssvnetwork *SsvnetworkFilterer) WatchOperatorFeeCancellationDeclared(opts *bind.WatchOpts, sink chan<- *SsvnetworkOperatorFeeCancellationDeclared, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "OperatorFeeCancellationDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkOperatorFeeCancellationDeclared)
				if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorFeeCancellationDeclared", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorFeeCancellationDeclared is a log parse operation binding the contract event 0x269e5b30b1f6e64e296fff57d56c77a4eb77fda7421146f1497f567539ba8906.
//
// Solidity: event OperatorFeeCancellationDeclared(address indexed owner, uint64 indexed operatorId)
func (_Ssvnetwork *SsvnetworkFilterer) ParseOperatorFeeCancellationDeclared(log types.Log) (*SsvnetworkOperatorFeeCancellationDeclared, error) {
	event := new(SsvnetworkOperatorFeeCancellationDeclared)
	if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorFeeCancellationDeclared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkOperatorFeeDeclaredIterator is returned from FilterOperatorFeeDeclared and is used to iterate over the raw logs and unpacked data for OperatorFeeDeclared events raised by the Ssvnetwork contract.
type SsvnetworkOperatorFeeDeclaredIterator struct {
	Event *SsvnetworkOperatorFeeDeclared // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkOperatorFeeDeclaredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkOperatorFeeDeclared)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkOperatorFeeDeclared)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkOperatorFeeDeclaredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkOperatorFeeDeclaredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkOperatorFeeDeclared represents a OperatorFeeDeclared event raised by the Ssvnetwork contract.
type SsvnetworkOperatorFeeDeclared struct {
	Owner       common.Address
	OperatorId  uint64
	BlockNumber *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeDeclared is a free log retrieval operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssvnetwork *SsvnetworkFilterer) FilterOperatorFeeDeclared(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvnetworkOperatorFeeDeclaredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "OperatorFeeDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkOperatorFeeDeclaredIterator{contract: _Ssvnetwork.contract, event: "OperatorFeeDeclared", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeDeclared is a free log subscription operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssvnetwork *SsvnetworkFilterer) WatchOperatorFeeDeclared(opts *bind.WatchOpts, sink chan<- *SsvnetworkOperatorFeeDeclared, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "OperatorFeeDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkOperatorFeeDeclared)
				if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorFeeDeclared", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorFeeDeclared is a log parse operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssvnetwork *SsvnetworkFilterer) ParseOperatorFeeDeclared(log types.Log) (*SsvnetworkOperatorFeeDeclared, error) {
	event := new(SsvnetworkOperatorFeeDeclared)
	if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorFeeDeclared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkOperatorFeeExecutedIterator is returned from FilterOperatorFeeExecuted and is used to iterate over the raw logs and unpacked data for OperatorFeeExecuted events raised by the Ssvnetwork contract.
type SsvnetworkOperatorFeeExecutedIterator struct {
	Event *SsvnetworkOperatorFeeExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkOperatorFeeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkOperatorFeeExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkOperatorFeeExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkOperatorFeeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkOperatorFeeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkOperatorFeeExecuted represents a OperatorFeeExecuted event raised by the Ssvnetwork contract.
type SsvnetworkOperatorFeeExecuted struct {
	Owner       common.Address
	OperatorId  uint64
	BlockNumber *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeExecuted is a free log retrieval operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssvnetwork *SsvnetworkFilterer) FilterOperatorFeeExecuted(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvnetworkOperatorFeeExecutedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "OperatorFeeExecuted", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkOperatorFeeExecutedIterator{contract: _Ssvnetwork.contract, event: "OperatorFeeExecuted", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeExecuted is a free log subscription operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssvnetwork *SsvnetworkFilterer) WatchOperatorFeeExecuted(opts *bind.WatchOpts, sink chan<- *SsvnetworkOperatorFeeExecuted, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "OperatorFeeExecuted", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkOperatorFeeExecuted)
				if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorFeeExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorFeeExecuted is a log parse operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssvnetwork *SsvnetworkFilterer) ParseOperatorFeeExecuted(log types.Log) (*SsvnetworkOperatorFeeExecuted, error) {
	event := new(SsvnetworkOperatorFeeExecuted)
	if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorFeeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkOperatorFeeIncreaseLimitUpdatedIterator is returned from FilterOperatorFeeIncreaseLimitUpdated and is used to iterate over the raw logs and unpacked data for OperatorFeeIncreaseLimitUpdated events raised by the Ssvnetwork contract.
type SsvnetworkOperatorFeeIncreaseLimitUpdatedIterator struct {
	Event *SsvnetworkOperatorFeeIncreaseLimitUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkOperatorFeeIncreaseLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkOperatorFeeIncreaseLimitUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkOperatorFeeIncreaseLimitUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkOperatorFeeIncreaseLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkOperatorFeeIncreaseLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkOperatorFeeIncreaseLimitUpdated represents a OperatorFeeIncreaseLimitUpdated event raised by the Ssvnetwork contract.
type SsvnetworkOperatorFeeIncreaseLimitUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeIncreaseLimitUpdated is a free log retrieval operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) FilterOperatorFeeIncreaseLimitUpdated(opts *bind.FilterOpts) (*SsvnetworkOperatorFeeIncreaseLimitUpdatedIterator, error) {

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "OperatorFeeIncreaseLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvnetworkOperatorFeeIncreaseLimitUpdatedIterator{contract: _Ssvnetwork.contract, event: "OperatorFeeIncreaseLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeIncreaseLimitUpdated is a free log subscription operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) WatchOperatorFeeIncreaseLimitUpdated(opts *bind.WatchOpts, sink chan<- *SsvnetworkOperatorFeeIncreaseLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "OperatorFeeIncreaseLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkOperatorFeeIncreaseLimitUpdated)
				if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorFeeIncreaseLimitUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorFeeIncreaseLimitUpdated is a log parse operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_Ssvnetwork *SsvnetworkFilterer) ParseOperatorFeeIncreaseLimitUpdated(log types.Log) (*SsvnetworkOperatorFeeIncreaseLimitUpdated, error) {
	event := new(SsvnetworkOperatorFeeIncreaseLimitUpdated)
	if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorFeeIncreaseLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the Ssvnetwork contract.
type SsvnetworkOperatorRemovedIterator struct {
	Event *SsvnetworkOperatorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkOperatorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkOperatorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkOperatorRemoved represents a OperatorRemoved event raised by the Ssvnetwork contract.
type SsvnetworkOperatorRemoved struct {
	OperatorId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_Ssvnetwork *SsvnetworkFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, operatorId []uint64) (*SsvnetworkOperatorRemovedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "OperatorRemoved", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkOperatorRemovedIterator{contract: _Ssvnetwork.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_Ssvnetwork *SsvnetworkFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *SsvnetworkOperatorRemoved, operatorId []uint64) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "OperatorRemoved", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkOperatorRemoved)
				if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRemoved is a log parse operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_Ssvnetwork *SsvnetworkFilterer) ParseOperatorRemoved(log types.Log) (*SsvnetworkOperatorRemoved, error) {
	event := new(SsvnetworkOperatorRemoved)
	if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkOperatorWhitelistUpdatedIterator is returned from FilterOperatorWhitelistUpdated and is used to iterate over the raw logs and unpacked data for OperatorWhitelistUpdated events raised by the Ssvnetwork contract.
type SsvnetworkOperatorWhitelistUpdatedIterator struct {
	Event *SsvnetworkOperatorWhitelistUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkOperatorWhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkOperatorWhitelistUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkOperatorWhitelistUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkOperatorWhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkOperatorWhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkOperatorWhitelistUpdated represents a OperatorWhitelistUpdated event raised by the Ssvnetwork contract.
type SsvnetworkOperatorWhitelistUpdated struct {
	OperatorId  uint64
	Whitelisted common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorWhitelistUpdated is a free log retrieval operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_Ssvnetwork *SsvnetworkFilterer) FilterOperatorWhitelistUpdated(opts *bind.FilterOpts, operatorId []uint64) (*SsvnetworkOperatorWhitelistUpdatedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "OperatorWhitelistUpdated", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkOperatorWhitelistUpdatedIterator{contract: _Ssvnetwork.contract, event: "OperatorWhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorWhitelistUpdated is a free log subscription operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_Ssvnetwork *SsvnetworkFilterer) WatchOperatorWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *SsvnetworkOperatorWhitelistUpdated, operatorId []uint64) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "OperatorWhitelistUpdated", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkOperatorWhitelistUpdated)
				if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorWhitelistUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorWhitelistUpdated is a log parse operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_Ssvnetwork *SsvnetworkFilterer) ParseOperatorWhitelistUpdated(log types.Log) (*SsvnetworkOperatorWhitelistUpdated, error) {
	event := new(SsvnetworkOperatorWhitelistUpdated)
	if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorWhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkOperatorWithdrawnIterator is returned from FilterOperatorWithdrawn and is used to iterate over the raw logs and unpacked data for OperatorWithdrawn events raised by the Ssvnetwork contract.
type SsvnetworkOperatorWithdrawnIterator struct {
	Event *SsvnetworkOperatorWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkOperatorWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkOperatorWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkOperatorWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkOperatorWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkOperatorWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkOperatorWithdrawn represents a OperatorWithdrawn event raised by the Ssvnetwork contract.
type SsvnetworkOperatorWithdrawn struct {
	Owner      common.Address
	OperatorId uint64
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorWithdrawn is a free log retrieval operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_Ssvnetwork *SsvnetworkFilterer) FilterOperatorWithdrawn(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvnetworkOperatorWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "OperatorWithdrawn", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkOperatorWithdrawnIterator{contract: _Ssvnetwork.contract, event: "OperatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchOperatorWithdrawn is a free log subscription operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_Ssvnetwork *SsvnetworkFilterer) WatchOperatorWithdrawn(opts *bind.WatchOpts, sink chan<- *SsvnetworkOperatorWithdrawn, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "OperatorWithdrawn", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkOperatorWithdrawn)
				if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorWithdrawn is a log parse operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_Ssvnetwork *SsvnetworkFilterer) ParseOperatorWithdrawn(log types.Log) (*SsvnetworkOperatorWithdrawn, error) {
	event := new(SsvnetworkOperatorWithdrawn)
	if err := _Ssvnetwork.contract.UnpackLog(event, "OperatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Ssvnetwork contract.
type SsvnetworkOwnershipTransferStartedIterator struct {
	Event *SsvnetworkOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Ssvnetwork contract.
type SsvnetworkOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ssvnetwork *SsvnetworkFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SsvnetworkOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkOwnershipTransferStartedIterator{contract: _Ssvnetwork.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ssvnetwork *SsvnetworkFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *SsvnetworkOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkOwnershipTransferStarted)
				if err := _Ssvnetwork.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ssvnetwork *SsvnetworkFilterer) ParseOwnershipTransferStarted(log types.Log) (*SsvnetworkOwnershipTransferStarted, error) {
	event := new(SsvnetworkOwnershipTransferStarted)
	if err := _Ssvnetwork.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ssvnetwork contract.
type SsvnetworkOwnershipTransferredIterator struct {
	Event *SsvnetworkOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkOwnershipTransferred represents a OwnershipTransferred event raised by the Ssvnetwork contract.
type SsvnetworkOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ssvnetwork *SsvnetworkFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SsvnetworkOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkOwnershipTransferredIterator{contract: _Ssvnetwork.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ssvnetwork *SsvnetworkFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SsvnetworkOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkOwnershipTransferred)
				if err := _Ssvnetwork.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ssvnetwork *SsvnetworkFilterer) ParseOwnershipTransferred(log types.Log) (*SsvnetworkOwnershipTransferred, error) {
	event := new(SsvnetworkOwnershipTransferred)
	if err := _Ssvnetwork.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Ssvnetwork contract.
type SsvnetworkUpgradedIterator struct {
	Event *SsvnetworkUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkUpgraded represents a Upgraded event raised by the Ssvnetwork contract.
type SsvnetworkUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Ssvnetwork *SsvnetworkFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SsvnetworkUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkUpgradedIterator{contract: _Ssvnetwork.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Ssvnetwork *SsvnetworkFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SsvnetworkUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkUpgraded)
				if err := _Ssvnetwork.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Ssvnetwork *SsvnetworkFilterer) ParseUpgraded(log types.Log) (*SsvnetworkUpgraded, error) {
	event := new(SsvnetworkUpgraded)
	if err := _Ssvnetwork.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the Ssvnetwork contract.
type SsvnetworkValidatorAddedIterator struct {
	Event *SsvnetworkValidatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkValidatorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkValidatorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkValidatorAdded represents a ValidatorAdded event raised by the Ssvnetwork contract.
type SsvnetworkValidatorAdded struct {
	Owner       common.Address
	OperatorIds []uint64
	PublicKey   []byte
	Shares      []byte
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xd432d53fd23906d0269fb9864b3832f32b83148df4315eab6ca914225c7fa2f4.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) FilterValidatorAdded(opts *bind.FilterOpts, owner []common.Address) (*SsvnetworkValidatorAddedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "ValidatorAdded", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkValidatorAddedIterator{contract: _Ssvnetwork.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xd432d53fd23906d0269fb9864b3832f32b83148df4315eab6ca914225c7fa2f4.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *SsvnetworkValidatorAdded, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "ValidatorAdded", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkValidatorAdded)
				if err := _Ssvnetwork.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorAdded is a log parse operation binding the contract event 0xd432d53fd23906d0269fb9864b3832f32b83148df4315eab6ca914225c7fa2f4.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) ParseValidatorAdded(log types.Log) (*SsvnetworkValidatorAdded, error) {
	event := new(SsvnetworkValidatorAdded)
	if err := _Ssvnetwork.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvnetworkValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the Ssvnetwork contract.
type SsvnetworkValidatorRemovedIterator struct {
	Event *SsvnetworkValidatorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SsvnetworkValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvnetworkValidatorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SsvnetworkValidatorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SsvnetworkValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvnetworkValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvnetworkValidatorRemoved represents a ValidatorRemoved event raised by the Ssvnetwork contract.
type SsvnetworkValidatorRemoved struct {
	Owner       common.Address
	OperatorIds []uint64
	PublicKey   []byte
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xf80b0a8cdbf1aa38914ab10cc803eba9406d6f8ed18780f2fe7b2e8b8e55f2b3.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, owner []common.Address) (*SsvnetworkValidatorRemovedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.FilterLogs(opts, "ValidatorRemoved", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvnetworkValidatorRemovedIterator{contract: _Ssvnetwork.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xf80b0a8cdbf1aa38914ab10cc803eba9406d6f8ed18780f2fe7b2e8b8e55f2b3.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *SsvnetworkValidatorRemoved, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssvnetwork.contract.WatchLogs(opts, "ValidatorRemoved", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvnetworkValidatorRemoved)
				if err := _Ssvnetwork.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRemoved is a log parse operation binding the contract event 0xf80b0a8cdbf1aa38914ab10cc803eba9406d6f8ed18780f2fe7b2e8b8e55f2b3.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,uint256,bool) cluster)
func (_Ssvnetwork *SsvnetworkFilterer) ParseValidatorRemoved(log types.Log) (*SsvnetworkValidatorRemoved, error) {
	event := new(SsvnetworkValidatorRemoved)
	if err := _Ssvnetwork.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
