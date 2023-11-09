package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"try_rde/abistr"
)

const (
	L1URL              = "http://localhost:8545"
	L2URL              = "http://localhost:9545"
	l1ContractAddr     = "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9" //proxy addr
	l2ContractAddr     = "0x4200000000000000000000000000000000000010"
	legacyERC20MNTAddr = "0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000"
	WETH9Addr          = "0xdEAddEaDdeadDEadDEADDEAddEADDEAddead1111" //"0x4200000000000000000000000000000000000006"
	//l2StandardBridgeAddr = "0xC0d3c0d3c0D3c0d3C0D3c0D3C0d3C0D3C0D30010"
	L2_CROSS_DOMAIN_MESSENGER_AddrHex         = "0x4200000000000000000000000000000000000007"
	Proxy__BVM_L1CrossDomainMessenger_AddrHex = "0x0165878A594ca255338adfa4d48449f69242Eb8F"
	L2ToL1MessagePasser                       = "0x4200000000000000000000000000000000000016"
	L1OptimismPortal                          = "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853"
	L2OutputOracleProxy                       = "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707"

	account1  = "0x784e50947Df23dBa8f91029089ef7B046257E544"
	account4  = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
	account20 = "0x00000500E87eE83A1BFa233512af25a4003836C8" // Account20

	account1SK  = "0d0c6dd2f25fc746bcec70aa27b31ec7fcd949ff5ec69dc58276d2d233f344c9"
	account4SK  = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	account20SK = "dd888cfabd6d3c3eeb683063657706fb660416ec4972bb5761204e0dbf59e33c"
)

func main() {
	l1Client, err := ethclient.Dial(L1URL)
	if err != nil {
		fmt.Printf("[err 00] %s\n", err.Error())
	}
	l1AccountAddress := common.HexToAddress(account4)
	l1ContractAddress := common.HexToAddress(l1ContractAddr)

	// abi
	l1ContractABI, err := abi.JSON(strings.NewReader(abistr.L1StandardBridgeABI))
	if err != nil {
		fmt.Printf("[err 0] %s\n", err.Error())
	}

	//// approve
	//amountToApprove := big.NewInt(10000000) //
	//approveData, err := l1ContractABI.Pack("approve", l1ContractAddress, amountToApprove)
	//if err != nil {
	//	fmt.Printf("[err 0.5] %s\n", err.Error())
	//}

	// depositMNT args
	amount := big.NewInt(10000000)
	gasPrice, err := l1Client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("[err 1] %s\n", err.Error())
	}
	gasLimit := uint32(2000000)
	callData, err := l1ContractABI.Pack("depositMNT", big.NewInt(100), gasLimit, []byte{}) // approveData)
	if err != nil {
		fmt.Printf("[err 2] %s\n", err.Error())
	}

	pendingNonce, err := l1Client.PendingNonceAt(context.Background(), l1AccountAddress)
	if err != nil {
		fmt.Printf("[err 3] %s\n", err.Error())
	}
	tx := types.NewTransaction(pendingNonce, l1ContractAddress, amount, uint64(gasLimit), gasPrice, callData)

	chainID, err := l1Client.ChainID(context.Background())
	if err != nil {
		fmt.Printf("[err 4] %s\n", err.Error())
	}
	privKey, err := crypto.HexToECDSA(account4SK)
	if err != nil {
		fmt.Printf("[err 5] %s\n", err.Error())
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		fmt.Printf("[err 6] %s\n", err.Error())
	}
	err = l1Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Printf("[err 7] %s\n", err.Error())
	}

	fmt.Println("finish")
}
