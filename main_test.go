package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"strings"
	"testing"
	"try_rde/abistr"
)

func Test_L1DepositMNT(t *testing.T) {
	l1Client, err := ethclient.Dial(l1URL)
	require.NoError(t, err)
	if err != nil {
		t.Logf("[err 00] %s\n", err.Error())
	}
	l1AccountAddress := common.HexToAddress(account4)
	l1ContractAddress := common.HexToAddress(l1ContractAddr)

	// abi
	l1ContractABI, err := abi.JSON(strings.NewReader(abistr.L1StandardBridgeABI))
	if err != nil {
		t.Logf("[err 0] %s\n", err.Error())
	}

	//// approve
	//amountToApprove := big.NewInt(10000000) //
	//approveData, err := l1ContractABI.Pack("approve", l1ContractAddress, amountToApprove)
	//if err != nil {
	//	t.Logf("[err 0.5] %s\n", err.Error())
	//}

	// depositMNT args
	amount := big.NewInt(10000000)
	gasPrice, err := l1Client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Logf("[err 1] %s\n", err.Error())
	}
	gasLimit := uint32(200000)
	callData, err := l1ContractABI.Pack("depositMNT", big.NewInt(100), gasLimit, []byte{}) // approveData)
	if err != nil {
		t.Logf("[err 2] %s\n", err.Error())
	}

	pendingNonce, err := l1Client.PendingNonceAt(context.Background(), l1AccountAddress)
	if err != nil {
		t.Logf("[err 3] %s\n", err.Error())
	}
	tx := types.NewTransaction(pendingNonce, l1ContractAddress, amount, uint64(gasLimit), gasPrice, callData)

	chainID, err := l1Client.ChainID(context.Background())
	if err != nil {
		t.Logf("[err 4] %s\n", err.Error())
	}
	privKey, err := crypto.HexToECDSA(account4SK)
	if err != nil {
		t.Logf("[err 5] %s\n", err.Error())
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		t.Logf("[err 6] %s\n", err.Error())
	}

	err = l1Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Logf("[err 7] %s\n", err.Error())
	}

	fmt.Println("finish")
}

func Test_L1DepositEth(t *testing.T) {
	ast := assert.New(t)

	l1Client, err := ethclient.Dial(l1URL)
	if err != nil {
		t.Logf("[err 00] %s\n", err.Error())
	}
	l1AccountAddress := common.HexToAddress(account20)
	l1ContractAddress := common.HexToAddress(l1ContractAddr)

	// abi
	l1ContractABI, err := abi.JSON(strings.NewReader(abistr.L1StandardBridgeABI))
	ast.NoError(err)

	// depositMNT args
	amount := big.NewInt(100)
	gasPrice, err := l1Client.SuggestGasPrice(context.Background())
	//gasPrice := big.NewInt(10007)
	ast.NoError(err)
	t.Logf("suggestGasPrice is %s\n", gasPrice.String())

	gasLimit := uint32(2000000) //uint32(2000000000)

	callData, err := l1ContractABI.Pack("depositETH", gasLimit, []byte{}) // approveData)
	ast.NoError(err)
	t.Logf("callData is %s\n", common.ToHex(callData))

	pendingNonce, err := l1Client.PendingNonceAt(context.Background(), l1AccountAddress)
	ast.NoError(err)

	tx := types.NewTransaction(pendingNonce, l1ContractAddress, amount, uint64(gasLimit), gasPrice, callData)

	chainID, err := l1Client.ChainID(context.Background())
	ast.NoError(err)

	privKey, err := crypto.HexToECDSA(account20SK)
	ast.NoError(err)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	ast.NoError(err)
	t.Logf("tx hash is %s\n", signedTx.Hash().Hex())

	estMsg := ethereum.CallMsg{
		From:     l1AccountAddress,
		To:       &l1ContractAddress,
		Gas:      uint64(gasLimit),
		GasPrice: gasPrice,
		Value:    amount,
		Data:     callData,
	}
	estGas, estErr := l1Client.EstimateGas(context.Background(), estMsg)
	if estErr != nil {
		t.Logf("estErr: %s\n", estErr.Error())
	} else {
		t.Logf("estGas: %d\n", estGas)
	}

	err = l1Client.SendTransaction(context.Background(), signedTx)
	ast.NoError(err)

	fmt.Println("finish")
}

func Test_L2Withdraw(t *testing.T) {
	l2Client, err := ethclient.Dial(l2URL)
	if err != nil {
		t.Logf("[err 00] %s\n", err.Error())
	}
	l2Account := account20
	l2AccountAddress := common.HexToAddress(l2Account)
	l2AccountSK := account20SK
	l2AccountSKEcdsa, err := crypto.HexToECDSA(l2AccountSK)
	l2ContractABI, err := abi.JSON(strings.NewReader(abistr.L1StandardBridgeABI))
	chainID, err := l2Client.ChainID(context.Background())

	pendingNonce, err := l2Client.PendingNonceAt(context.Background(), l2AccountAddress)
	amount := big.NewInt(1000000)
	gasLimit := uint64(10000)
	callData, err := l2ContractABI.Pack("withdraw", l2AccountAddress, uint32(gasLimit), []byte{})
	gasPrice, err := l2Client.SuggestGasPrice(context.Background())
	tx := types.NewTransaction(pendingNonce, common.HexToAddress(l2ContractAddr), amount, gasLimit, gasPrice, callData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), l2AccountSKEcdsa)

	sendTxErr := l2Client.SendTransaction(context.Background(), signedTx)
	require.NoError(t, sendTxErr, "sendTransaction should has no err")
}
