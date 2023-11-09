package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"try_rde/abistr"
	"try_rde/abistr/abijson"
	"try_rde/txutils"
)

func Test_L1DepositEth_failed_pending(t *testing.T) {
	ast := assert.New(t)

	l1Client, err := ethclient.Dial(L1URL)
	if err != nil {
		t.Fatalf("[err 00] %s\n", err.Error())
	}
	l1AccountAddress := common.HexToAddress(account20)
	l1ContractAddress := common.HexToAddress(l1ContractAddr)

	// abi
	l1ContractABI, err := abi.JSON(strings.NewReader(abistr.L1StandardBridgeABI))
	ast.NoError(err)

	// depositETH args
	amount := new(big.Int).Mul(big.NewInt(5), big.NewInt(params.Ether)) //big.NewInt(1000000000000000000)

	//gasLimit := uint32(3000000) //uint32(2000000000)

	//gas
	//gasTipCap, gasFeeCap, estGasLimit, err := txutils.GetGas(l1Client, l1ContractABI, gasLimit, l1AccountAddress, l1ContractAddress, amount)
	ast.NoError(err)

	callData, err := l1ContractABI.Pack("depositETH", uint32(100000), []byte{}) // approveData)
	ast.NoError(err)
	t.Logf("callData is %s\n", hex.EncodeToString(callData))

	pendingNonce, err := l1Client.PendingNonceAt(context.Background(), l1AccountAddress)
	ast.NoError(err)

	chainID, err := l1Client.ChainID(context.Background())
	ast.NoError(err)

	//tx := types.NewTransaction(pendingNonce, l1ContractAddress, amount, uint64(gasLimit), gasPrice, callData)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     pendingNonce,
		GasTipCap: big.NewInt(99),
		GasFeeCap: big.NewInt(99),
		Gas:       uint64(2500000),
		To:        &l1ContractAddress,
		Value:     amount,
		Data:      callData,
	})

	privKey, err := crypto.HexToECDSA(account20SK)
	ast.NoError(err)

	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	signer := types.LatestSignerForChainID(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privKey)
	ast.NoError(err)
	signedTx, err := tx.WithSignature(signer, signature)
	ast.NoError(err)
	t.Logf("tx hash is %s\n", signedTx.Hash().Hex())

	err = l1Client.SendTransaction(context.Background(), signedTx)
	ast.NoError(err)

	fmt.Println("finish")
}

func Test_L1DepositETHWithAbiInst_nogasPrice(t *testing.T) {
	ast := assert.New(t)
	cli, err := ethclient.Dial(L1URL)
	ast.NoError(err)

	l1AccountAddress := common.HexToAddress(account20)
	l1ContractAddress := common.HexToAddress(l1ContractAddr)
	chainID, err := cli.ChainID(context.Background())
	ast.NoError(err)

	contract, err := abijson.NewL1StandardBridge(l1ContractAddress, cli)
	ast.NoError(err)

	header, err := cli.HeaderByNumber(context.Background(), nil)
	ast.NoError(err)

	privKey, err := crypto.HexToECDSA(account20SK)
	ast.NoError(err)

	gasTipCap, err := cli.SuggestGasTipCap(context.Background())
	ast.NoError(err)

	gasFeeCap := new(big.Int).Add(gasTipCap, new(big.Int).Mul(header.BaseFee, big.NewInt(2)))

	opt, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	ast.NoError(err)

	opt.GasFeeCap = gasFeeCap
	opt.GasTipCap = gasTipCap
	opt.GasLimit = uint64(2000000)
	//opt.GasPrice = big.NewInt(100)
	opt.Value = new(big.Int).Mul(big.NewInt(100), big.NewInt(params.Ether)) //big.NewInt(200000)

	pendingNonce, err := cli.PendingNonceAt(context.Background(), l1AccountAddress)
	ast.NoError(err)

	opt.Nonce = big.NewInt(int64(pendingNonce))

	minGasLimit := uint32(200000)
	tx, err := contract.DepositETH(opt, minGasLimit, []byte{})
	ast.NoError(err)

	t.Logf("tx is nil? %v\n", tx == nil)
	signedTx, err := opt.Signer(l1AccountAddress, tx)
	ast.NoError(err)

	err = cli.SendTransaction(context.Background(), signedTx)
	ast.NoError(err)
	t.Logf("tx.hash is %s\n", tx.Hash())
	t.Logf("最后的err %s\n", err.Error())
}

func Test_L1DepositETHWithAbiInst_gasPrice(t *testing.T) {
	ast := assert.New(t)
	cli, err := ethclient.Dial(L1URL)
	ast.NoError(err)

	l1AccountAddress := common.HexToAddress(account20)
	l1ContractAddress := common.HexToAddress(l1ContractAddr)
	chainID, err := cli.ChainID(context.Background())
	ast.NoError(err)

	contract, err := abijson.NewL1StandardBridge(l1ContractAddress, cli)
	ast.NoError(err)

	//header, err := cli.HeaderByNumber(context.Background(), nil)
	//ast.NoError(err)

	privKey, err := crypto.HexToECDSA(account20SK)
	ast.NoError(err)

	//gasTipCap, err := cli.SuggestGasTipCap(context.Background())
	//ast.NoError(err)

	//gasFeeCap := new(big.Int).Add(gasTipCap, new(big.Int).Mul(header.BaseFee, big.NewInt(2)))

	opt, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	ast.NoError(err)

	//opt.GasFeeCap = gasFeeCap
	//opt.GasTipCap = gasTipCap
	opt.GasLimit = uint64(2500000)
	opt.GasPrice = big.NewInt(99)
	opt.Value = new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.Ether)) //big.NewInt(1000000000000000000)

	pendingNonce, err := cli.PendingNonceAt(context.Background(), l1AccountAddress)
	ast.NoError(err)

	opt.Nonce = big.NewInt(int64(pendingNonce))

	minGasLimit := uint32(100000)
	tx, err := contract.DepositETH(opt, minGasLimit, []byte{})
	ast.NoError(err)

	t.Logf("tx is nil? %v\n", tx == nil)
	signedTx, err := opt.Signer(l1AccountAddress, tx)
	ast.NoError(err)

	err = cli.SendTransaction(context.Background(), signedTx)
	ast.NoError(err)
	t.Logf("tx.value is %d\n", tx.Value())
	t.Logf("tx.nonce is %d\n", tx.Nonce())
	t.Logf("tx.hash is %s\n", tx.Hash())
	t.Logf("tx.gas = %d, tx.gasTip = %d, tx.gasFee = %d\n", tx.Gas(), tx.GasTipCap(), tx.GasFeeCap())
	t.Logf("最后的err %s\n", err.Error())
}

func Test_L1DepositMNT(t *testing.T) {
	ast := assert.New(t)

	l1Client, err := ethclient.Dial(L1URL)
	require.NoError(t, err)
	if err != nil {
		t.Logf("[err 00] %s\n", err.Error())
	}
	l1AccountAddress := common.HexToAddress(account20)
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
	//amount := big.NewInt(10000000)
	gasLimit := uint32(200000)
	callData, err := l1ContractABI.Pack("depositMNT", big.NewInt(100), gasLimit, []byte{}) // approveData)
	if err != nil {
		t.Logf("[err 2] %s\n", err.Error())
	}

	pendingNonce, err := l1Client.PendingNonceAt(context.Background(), l1AccountAddress)
	if err != nil {
		t.Logf("[err 3] %s\n", err.Error())
	}

	chainID, err := l1Client.ChainID(context.Background())
	if err != nil {
		t.Logf("[err 4] %s\n", err.Error())
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     pendingNonce,
		GasTipCap: big.NewInt(99),
		GasFeeCap: big.NewInt(99),
		Gas:       uint64(2500000),
		To:        &l1ContractAddress,
		Value:     nil,
		Data:      callData,
	})

	privKey, err := crypto.HexToECDSA(account20SK)
	ast.NoError(err)

	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	signer := types.LatestSignerForChainID(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privKey)
	ast.NoError(err)
	signedTx, err := tx.WithSignature(signer, signature)
	ast.NoError(err)
	t.Logf("tx hash is %s\n", signedTx.Hash().Hex())

	err = l1Client.SendTransaction(context.Background(), signedTx)
	ast.NoError(err)

	fmt.Println("finish")
}
func Test_L1DepositMNTWithAbiInst_succ(t *testing.T) {
	ast := assert.New(t)
	cli, err := ethclient.Dial(L1URL)
	ast.NoError(err)

	l1AccountAddress := common.HexToAddress(account20)
	l1ContractAddress := common.HexToAddress(l1ContractAddr)
	chainID, err := cli.ChainID(context.Background())
	ast.NoError(err)
	privKey, err := crypto.HexToECDSA(account20SK)
	ast.NoError(err)

	contract, err := abijson.NewL1StandardBridge(l1ContractAddress, cli)
	ast.NoError(err)

	// approve
	l1MNTokenAddr := common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512") //proxy_l1MantleToken
	l1MNTContract, err := abijson.NewL1MantleToken(l1MNTokenAddr, cli)                 // l1ContractAddress
	ast.NoError(err)

	opt, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	ast.NoError(err)
	tx, err := l1MNTContract.Approve(opt, l1AccountAddress, big.NewInt(100000))
	ast.NoError(err)

	signedTx, err := opt.Signer(l1AccountAddress, tx)
	ast.NoError(err)
	err = cli.SendTransaction(context.Background(), signedTx)
	ast.NoError(err)

	// deposit MNT
	opt, err = bind.NewKeyedTransactorWithChainID(privKey, chainID)
	ast.NoError(err)

	opt.Value = big.NewInt(1000)

	minGasLimit := uint32(200000)
	tx, err = contract.DepositMNT(opt, big.NewInt(10000), minGasLimit, []byte{})
	ast.NoError(err)

	t.Logf("tx is nil? %v\n", tx == nil)
	signedTx, err = opt.Signer(l1AccountAddress, tx)
	ast.NoError(err)

	err = cli.SendTransaction(context.Background(), signedTx)
	ast.NoError(err)
	t.Logf("tx.hash is %s\n", tx.Hash())
	t.Logf("最后的err %s\n", err.Error())
}

func Test_L2Withdraw_0(t *testing.T) {
	l2Client, err := ethclient.Dial(L2URL)
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

func Test_L2WithdrawETH(t *testing.T) {
	ast := assert.New(t)
	l1cli, err := ethclient.Dial(L1URL)
	l2cli, err := ethclient.Dial(L2URL)
	ast.NoError(err)
	//l1StandardBridge := common.HexToAddress(l1ContractAddr)
	l2StandardBridge := common.HexToAddress(l2ContractAddr)

	//l1ChainID, err := l1cli.ChainID(context.Background())
	ast.NoError(err)
	l2ChainID, err := l2cli.ChainID(context.Background())
	ast.NoError(err)

	//l2TokenAddr := common.HexToAddress("")
	//BVM_ETH_Addr := common.HexToAddress("0xdEAddEaDdeadDEadDEADDEAddEADDEAddead1111")
	account20Addr := common.HexToAddress("0x00000500E87eE83A1BFa233512af25a4003836C8")
	privKey, err := crypto.HexToECDSA(account20SK)
	ast.NoError(err)

	contract, err := abijson.NewL2StandardBridge(l2StandardBridge, l2cli)
	ast.NoError(err)

	approveAmount := new(big.Int).Mul(big.NewInt(60), big.NewInt(params.Ether))
	amount := new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether))

	// check balance
	t.Log("================= check balance =================")
	l1Bal, err := l1cli.BalanceAt(context.Background(), account20Addr, nil)
	ast.NoError(err)
	if err == nil {
		t.Logf("l1bal is %d\n", l1Bal) // 9998899998809314943436246
	}

	l2Bal := txutils.GetETHBalanceFromL2(t, account20)
	ast.NotNil(l2Bal)
	t.Logf("l2bal is %d\n", l2Bal) // 10000000000000200000

	// approve
	t.Log("================= approve =================")
	l2ETHTokenAddr := common.HexToAddress(WETH9Addr)
	L2WETHContract, err := abijson.NewL2TestToken(l2ETHTokenAddr, l2cli)
	ast.NoError(err)
	opt, err := bind.NewKeyedTransactorWithChainID(privKey, l2ChainID)
	ast.NoError(err)
	//tx, err := L2WETHContract.Approve(opt, account20Addr, approveAmount)
	tx, err := L2WETHContract.Approve(opt, l2StandardBridge, approveAmount)
	ast.NoError(err)
	signedTx, err := opt.Signer(account20Addr, tx)
	ast.NoError(err)
	err = l2cli.SendTransaction(context.Background(), signedTx)
	ast.NoError(err)
	t.Logf("approve tx hash is %s\n", tx.Hash().Hex())

	time.Sleep(5 * time.Second)

	// allowance
	t.Log("================= allowance =================")
	owner, spender := account20Addr, l2StandardBridge
	allawance, err := L2WETHContract.Allowance(&bind.CallOpts{}, owner, spender)
	ast.NoError(err)
	t.Logf("allowance is %d\n", allawance) // 60000000000000000000

	// withdraw
	t.Log("================= withdraw =================")
	opts, err := bind.NewKeyedTransactorWithChainID(privKey, l2ChainID)
	ast.NoError(err)

	wdTx, err := contract.Withdraw(opts, common.HexToAddress(WETH9Addr), amount, 200000, []byte{})
	ast.NoError(err)
	if err == nil {
		t.Logf("wdtx.hash is %s\n", wdTx.Hash()) // 0x37665cddfa9afb8ff481abfc757623b271317b40b0ddfbf0a4b766145ceff8b4
	}

	withdrawNonce := wdTx.Nonce()
	t.Logf("withdraw nonce is %d\n", withdrawNonce)
	// filter event
	l2ToL1MessagePasserAddr := common.HexToAddress("0x4200000000000000000000000000000000000016")
	l2ToL1MessagePasser, err := abijson.NewL2ToL1MessagePasser(l2ToL1MessagePasserAddr, l2cli)

	L2_CROSS_DOMAIN_MESSENGER_Addr := common.HexToAddress(L2_CROSS_DOMAIN_MESSENGER_AddrHex)
	Proxy__BVM_L1CrossDomainMessenger := common.HexToAddress("0x0165878A594ca255338adfa4d48449f69242Eb8F")
	wdNonce, bigbool := new(big.Int).SetString("01000000000000000000000000000000000000000000000000000000000006", 16) // 27
	t.Logf("big.setString bigbool %v, wdNonce %d\n", bigbool, wdNonce)
	iter, err := l2ToL1MessagePasser.FilterMessagePassed(&bind.FilterOpts{}, []*big.Int{wdNonce}, []common.Address{L2_CROSS_DOMAIN_MESSENGER_Addr}, []common.Address{Proxy__BVM_L1CrossDomainMessenger})
	ast.NoError(err)
	//event MessagePassed(
	//	uint256 indexed nonce,
	//	address indexed sender,
	//	address indexed target,
	//	uint256 mntValue,
	//	uint256 ethValue,
	//	uint256 gasLimit,
	//	bytes data,
	//	bytes32 withdrawalHash
	//);
	for iter.Next() {
		t.Logf("nonce: %d, sender: %s, target: %s, mntValue: %d, ethValue: %d, gasLimit: %d, data: %s, wdHash: %s",
			iter.Event.Nonce, iter.Event.Sender.Hex(), iter.Event.Target.Hex(), iter.Event.MntValue, iter.Event.EthValue, iter.Event.GasLimit, iter.Event.Data, iter.Event.WithdrawalHash)
	}

	time.Sleep(3 * time.Second)
	receipt, err := l2cli.TransactionReceipt(context.Background(), wdTx.Hash())
	ast.NoError(err)
	ast.NotNil(receipt)

}

func Test_ProvenAndFinalize(t *testing.T) {
	ast := assert.New(t)

	// env
	l1cli, err := ethclient.Dial(L1URL)
	l2cli, err := ethclient.Dial(L2URL)
	ast.NoError(err)
	//l1StandardBridge := common.HexToAddress(l1ContractAddr)
	//l2StandardBridge := common.HexToAddress(l2ContractAddr)

	l1ChainID, err := l1cli.ChainID(context.Background())
	ast.NoError(err)
	//l2ChainID, err := l2cli.ChainID(context.Background())
	//ast.NoError(err)

	privKey, err := crypto.HexToECDSA(account20SK)
	ast.NoError(err)

	l2ToL1MessagePasser, err := abijson.NewL2ToL1MessagePasser(common.HexToAddress(L2ToL1MessagePasser), l2cli)
	ast.NoError(err)

	//l2Contract, err := abijson.NewL2StandardBridge(l2StandardBridge, l2cli)
	//ast.NoError(err)

	// Receipt
	t.Log("================= receipt =================")
	wdTxHex := "0xe8b09e3c4956a5d26357502a8c0e36ed623f1343479294be470342eee801bb3f" //"0x37665cddfa9afb8ff481abfc757623b271317b40b0ddfbf0a4b766145ceff8b4"
	receipt, err := l2cli.TransactionReceipt(context.Background(), common.HexToHash(wdTxHex))
	ast.NoError(err)

	// prepare wd args
	t.Log("================= prepare withdraw tx =================")
	L2OutputOracleProxyAddr := common.HexToAddress(L2OutputOracleProxy)
	L2OutputOracleProxyContract, err := abijson.NewL2OutputOracleProxy(L2OutputOracleProxyAddr, l1cli)

	l2RpcClient, err := rpc.DialContext(context.Background(), L2URL)
	ast.NoError(err)
	l2GethCli := gethclient.New(l2RpcClient)
	receiptCli := ethclient.NewClient(l2RpcClient)

	opt := &bind.CallOpts{}
	period, err := L2OutputOracleProxyContract.FINALIZATIONPERIODSECONDS(opt)
	ast.NoError(err)
	if err == nil {
		t.Logf("proven period is %d\n", period)
	}

	finalizationPeriod, err := L2OutputOracleProxyContract.FINALIZATIONPERIODSECONDS(opt)
	ast.NoError(err)
	if err == nil {
		t.Logf("finalizationPeriod is %d\n", finalizationPeriod)
	}
	//blkNum := txutils.GetBlockNum(t, submissionInterval, startingBlockNumber, receipt.BlockNumber)
	blkNum, err := txutils.WaitForFinalizationPeriod(context.Background(), t, l1cli, common.HexToAddress(L1OptimismPortal), receipt.BlockNumber, finalizationPeriod)
	ast.NoError(err)
	header, err := receiptCli.HeaderByNumber(context.Background(), big.NewInt(int64(blkNum)))
	ast.NoError(err)
	ast.NotNil(header)
	t.Logf("blkNum is %d, header.Num is %d\n", blkNum, header.Number)

	params, err := txutils.ProveWithdrawalParameters(context.Background(), t, l2GethCli, receiptCli, common.HexToHash(wdTxHex), header, &L2OutputOracleProxyContract.L2OutputOracleProxyCaller, l2ToL1MessagePasser)
	ast.NoError(err)

	// // prepare withdraw transaction
	wd := &abijson.TypesWithdrawalTransaction{
		Nonce:    params.Nonce,
		Sender:   params.Sender,
		Target:   params.Target,
		MntValue: big.NewInt(0), //params.MNTValue,
		EthValue: params.ETHValue,
		GasLimit: params.GasLimit,
		Data:     params.Data,
	}

	// prepare proven
	t.Log("================= prepare proven =================")
	L1OptimismPortalAddr := common.HexToAddress("0xa513E6E4b8f2a923D98304ec87F64353C4D5C853")
	opPortal, err := abijson.NewL1OptimismPortal(L1OptimismPortalAddr, l1cli)
	ast.NoError(err)

	t.Logf("wd %+v\n", *wd)
	mytypeWd := txutils.GetMyWithdraw(wd)
	slot, err := mytypeWd.StorageSlot()
	ast.NoError(err)

	proof, err := l2GethCli.GetProof(context.Background(), predeploys.L2ToL1MessagePasserAddr, []string{slot.String()}, big.NewInt(int64(blkNum)))
	ast.NoError(err)
	ast.NotNil(proof)
	t.Logf("proof.storage.len is %d", len(proof.StorageProof))

	//outputRootProof := abijson.TypesOutputRootProof{ //bindings.TypesOutputRootProof{
	//	Version:                  [32]byte{},
	//	StateRoot:                header.Root,
	//	MessagePasserStorageRoot: proof.StorageHash,
	//	LatestBlockhash:          header.Hash(),
	//}

	//trieNodes := make([][]byte, len(proof.StorageProof[0].Proof))
	//for i, s := range proof.StorageProof[0].Proof {
	//	trieNodes[i] = common.FromHex(s)
	//}

	//// Compute the output root locally
	//l2OutputRoot, err := txutils.ComputeL2OutputRoot(&outputRootProof)
	//localOutputRootHash := common.Hash(l2OutputRoot)
	//ast.NoError(err)
	//// ensure that the locally computed hash matches
	////ast.Equal(l2Output.OutputRoot, localOutputRootHash, fmt.Sprintf("mismatch in output root hashes, got 0x%x expected 0x%x", localOutputRootHash, l2Output.OutputRoot))
	//if params.OutputRootProof.OutputRoot == localOutputRootHash {
	//	t.Log("l2Output.OutputRoot == localOutputRootHash")
	//}
	//
	// check storage value
	// abiTrue represents the storage representation of the boolean
	// value true.
	var abiTrue = common.Hash{31: 0x01}

	storageValue, err := l2cli.StorageAt(context.Background(), predeploys.L2ToL1MessagePasserAddr, slot, big.NewInt(int64(blkNum)))
	ast.NoError(err)
	if err == nil {
		t.Logf("L2ToL1MessagePasser status value %s\n", common.Bytes2Hex(storageValue))
	}

	// the value should be set to a boolean in storage
	ast.Equal(abiTrue.Bytes(), storageValue, fmt.Sprintf("storage slot %x not found in state", slot.Hex()))

	// check proven
	t.Log("================= check proven =================")
	wdHash, err := txutils.GetWdHash(wd)
	ast.NoError(err)

	var alreadyProve bool
	proven, err := opPortal.ProvenWithdrawals(&bind.CallOpts{}, wdHash)
	ast.NoError(err)
	ast.NotNil(proven)
	if proven.Timestamp.Cmp(common.Big0) == 0 {
		t.Log("check proven not proven yet")
	} else {
		alreadyProve = true
		t.Log("already proven")
	}

	if !alreadyProve {
		// proven
		t.Log("================= proven =================")
		proveOpt, err := bind.NewKeyedTransactorWithChainID(privKey, l1ChainID)
		ast.NoError(err)
		ast.NotNil(proveOpt)
		//proveTx, err := opPortal.ProveWithdrawalTransaction(opt, *wd, params.L2OutputIndex, outputRootProof, trieNodes)
		proveTx, err := opPortal.ProveWithdrawalTransaction(
			proveOpt,
			abijson.TypesWithdrawalTransaction{
				Nonce:    params.Nonce,
				Sender:   params.Sender,
				Target:   params.Target,
				MntValue: params.MNTValue,
				EthValue: params.ETHValue,
				GasLimit: params.GasLimit,
				Data:     params.Data,
			},
			params.L2OutputIndex,
			params.OutputRootProof,
			params.WithdrawalProof,
		)
		ast.NoError(err)
		ast.NotNil(proveTx)

		// waitMined
		receipt, err = bind.WaitMined(context.Background(), l1cli, proveTx)
		ast.NoError(err)
		if receipt.Status != types.ReceiptStatusSuccessful {
			t.Log("withdrawal proof unsuccessful")
		} else {
			t.Logf("withdraw prove succ")
		}
	}
	// check finalize
	t.Log("================= check finalize =================")

	isFinalized, err := opPortal.FinalizedWithdrawals(&bind.CallOpts{}, wdHash)
	ast.NoError(err)
	if err == nil {
		t.Logf("FinalizedWithdrawals isFinalized = %v \n", isFinalized)
	}

	// opPortal
	t.Log("================= finalize optimismPortal =================")
	finOpt, err := bind.NewKeyedTransactorWithChainID(privKey, l1ChainID)
	ast.NoError(err)
	fTx, err := opPortal.FinalizeWithdrawalTransaction(finOpt, *wd)
	ast.NoError(err)
	ast.NotNil(fTx)
	receipt, err = bind.WaitMined(context.Background(), l1cli, fTx)
	ast.NoError(err)
	ast.Equal(types.ReceiptStatusSuccessful, receipt.Status)
	//
	//wdHash := "0xa7ef171339c1c81b16107529355098b8ed26c7d02d0c83db7738ddafd32a1b52"
	//wdHashByte, err := hex.DecodeString(wdHash)
	//isFinalized, err := opPortal.FinalizedWithdrawals(&bind.CallOpts{})
	//ast.NoError(err)
	//t.Logf("opPortal is finalized? %v\n", isFinalized)

	//// l1StandardBridge finalize
	//t.Log("================= finalize =================")
	//l1Contract, err := abijson.NewL1StandardBridge(l1StandardBridge, l1cli)
	//ast.NoError(err)
	//opt, err = bind.NewKeyedTransactorWithChainID(privKey, l1ChainID)
	//ast.NoError(err)
	//from, to := l2StandardBridge, account20Addr
	////tx, err = contract.FinalizeBridgeETH(opt, from, to, amount, []byte{})
	//tx, err = l1Contract.FinalizeETHWithdrawal(opt, from, to, amount, []byte{})
	//ast.NoError(err)
	//signedTx, err = opt.Signer(account20Addr, tx)
	//ast.NoError(err)
	//t.Logf("finalizeETHWithdraw tx hash is %s\n", tx.Hash())
}
