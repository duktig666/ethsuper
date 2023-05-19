// desc:
// @author renshiwei
// Date: 2023/4/6 15:19

package eth1

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
)

type EthClient struct {
	Client *ethclient.Client
}

func NewEthClient(ctx context.Context, rpcHost string) (*EthClient, error) {
	elClient, err := getEthClient(ctx, rpcHost)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return &EthClient{
		Client: elClient,
	}, nil
}

func getEthClient(ctx context.Context, rpcHost string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcHost)
	if err != nil {
		return nil, errors.Wrap(err, "Fail to connect to Executive layer Client")
	}
	return client, nil
}

func KeyTransactOpts(chainID *big.Int, privateKeyStr string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, errors.Wrap(err, "PrivateKey hex err.")
	}
	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "eth client config privateKey provider err.")
	}
	return transactOpts, nil
}
