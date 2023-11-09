package txutils

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetGas(cli *ethclient.Client, contractABI abi.ABI, gasLimit uint32, from, to common.Address, value *big.Int) (*big.Int, *big.Int, uint32, error) {
	gasTipCap, err := cli.SuggestGasTipCap(context.Background())
	header, err := cli.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, nil, 0, err
	}
	gasFeeCap := new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(header.BaseFee, big.NewInt(2)),
	)

	estCallData, err := contractABI.Pack("depositETH", gasLimit, []byte{}) // approveData)
	if err != nil {
		return nil, nil, 0, err
	}
	fmt.Printf("callData is %s\n", hex.EncodeToString(estCallData))

	estMsg := ethereum.CallMsg{
		From:      from,
		To:        &to,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     value,
		Data:      estCallData,
	}
	estGasLimit, estErr := cli.EstimateGas(context.Background(), estMsg)
	if estErr != nil {
		fmt.Printf("estErr: %s\n", estErr.Error())
	} else {
		fmt.Printf("estGas: %d\n", estGasLimit)
	}

	return gasTipCap, gasFeeCap, gasLimit, nil
}
