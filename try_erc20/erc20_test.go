package try_erc20

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"try_rde/abistr/abijson"
	"try_rde/try_erc20/contract"
	"try_rde/txutils"
)

var (
	OtimismMintableERC20FactoryAddr = common.HexToAddress("0x4200000000000000000000000000000000000012")
	L1URL                           = "http://localhost:8545"
	L2URL                           = "http://localhost:9545"
	account20SK                     = "dd888cfabd6d3c3eeb683063657706fb660416ec4972bb5761204e0dbf59e33c"
	priv, _                         = crypto.HexToECDSA(account20SK)
	wwqTokenAddrL1                  = common.HexToAddress("0xfaeEBf311d135C7918430542e612b77033c4CA14")
	wwqTokenAddrL2                  = common.HexToAddress("0x7c6b91D9Be155A6Db01f749217d76fF02A7227F2")
	StandardL2TokenCreatedTopic     = crypto.Keccak256Hash([]byte("StandardL2TokenCreated(address,address)"))

	l1BridgeAddr = common.HexToAddress("0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9") //proxy addr
	l2BridgeAddr = common.HexToAddress("0x4200000000000000000000000000000000000010")
)

func Test_deployContractToL2(t *testing.T) {
	ast := assert.New(t)
	//l1cli, err := ethclient.Dial(L1URL)
	//ast.NoError(err)
	l2cli, err := ethclient.Dial(L2URL)
	ast.NoError(err)

	erc20Factory, err := contract.NewOptimismMintableERC20Factory(OtimismMintableERC20FactoryAddr, l2cli)
	ast.NoError(err)
	ast.NotNil(erc20Factory)

	chainID, err := l2cli.ChainID(context.Background())
	ast.NoError(err)

	opt, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	ast.NoError(err)
	tx, err := erc20Factory.CreateStandardL2Token(opt, wwqTokenAddrL1, "wwqToken", "WWQT")
	ast.NoError(err)
	if err == nil {
		t.Logf("L2 create contract tx hash is %s\n", tx.Hash().Hex()) // 0x95957a45ac14205053b3ae633be11355ef20843d038b6696c9d3a8c2899e7cf8
	}

	time.Sleep(3 * time.Second)

	receipt, err := l2cli.TransactionReceipt(context.Background(), tx.Hash())
	ast.NoError(err)
	ast.Equal(types.ReceiptStatusSuccessful, receipt.Status, "receipt status should be 0x01")

	l2ContractAddr, err := getL1ContractAddress(erc20Factory, receipt, StandardL2TokenCreatedTopic) // 0x7c6b91D9Be155A6Db01f749217d76fF02A7227F2
	ast.NoError(err)
	if err == nil {
		t.Logf("L2 contract address is %s\n", l2ContractAddr.Hex())
	}
}

func Test_getReceiptAndL2ContractAddr(t *testing.T) {
	ast := assert.New(t)
	l2cli, err := ethclient.Dial(L2URL)
	ast.NoError(err)

	erc20Factory, err := contract.NewOptimismMintableERC20Factory(OtimismMintableERC20FactoryAddr, l2cli)
	ast.NoError(err)
	ast.NotNil(erc20Factory)

	txHash := common.HexToHash("0x95957a45ac14205053b3ae633be11355ef20843d038b6696c9d3a8c2899e7cf8")
	receipt, err := l2cli.TransactionReceipt(context.Background(), txHash)
	ast.NoError(err)
	ast.Equal(types.ReceiptStatusSuccessful, receipt.Status, "receipt status should be 0x01")

	l2ContractAddr, err := getL1ContractAddress(erc20Factory, receipt, StandardL2TokenCreatedTopic)
	ast.NoError(err)
	if err == nil {
		t.Logf("L2 contract address is %s\n", l2ContractAddr.Hex())
	}
}

func Test_approveForDepositWWQT(t *testing.T) {
	ast := assert.New(t)
	l1cli, err := ethclient.Dial(L1URL)
	ast.NoError(err)

	l1Bridge, err := abijson.NewL1StandardBridge(l1BridgeAddr, l1cli)
	ast.NoError(err)
	ast.NotNil(l1Bridge)
	l1WWQTContract, err := contract.NewTokenWwqERC20(wwqTokenAddrL1, l1cli)
	ast.NoError(err)
	ast.NotNil(l1WWQTContract)

	// approve
	t.Log("========================== approve =========================")
	chainID, err := l1cli.ChainID(context.Background())
	ast.NoError(err)
	opt, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	ast.NoError(err)
	//tx, err := l1WWQTContract.Approve(opt, crypto.PubkeyToAddress(priv.PublicKey), big.NewInt(1000000))
	tx, err := l1WWQTContract.Approve(opt, l1BridgeAddr, big.NewInt(1000000))
	ast.NoError(err)
	if err == nil {
		t.Logf("approve tx hash is %s\n", tx.Hash().Hex())
	}

	// approve receipt
	time.Sleep(10 * time.Second)
	receipt, err := l1cli.TransactionReceipt(context.Background(), tx.Hash())
	ast.NoError(err)
	ast.Equal(types.ReceiptStatusSuccessful, receipt.Status, "receipt status should be 0x01")
}

func Test_depositWWQT(t *testing.T) {
	ast := assert.New(t)
	l1cli, err := ethclient.Dial(L1URL)
	ast.NoError(err)

	l1Bridge, err := abijson.NewL1StandardBridge(l1BridgeAddr, l1cli)
	ast.NoError(err)
	ast.NotNil(l1Bridge)
	l1WWQTContract, err := contract.NewTokenWwqERC20(wwqTokenAddrL1, l1cli)
	ast.NoError(err)
	ast.NotNil(l1WWQTContract)

	// deposit
	t.Log("========================== deposit =========================")
	chainID, err := l1cli.ChainID(context.Background())
	ast.NoError(err)
	opt, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	ast.NoError(err)
	amount, minGasLimit := big.NewInt(10000), uint32(200000)
	tx, err := l1Bridge.DepositERC20(opt, wwqTokenAddrL1, wwqTokenAddrL2, amount, minGasLimit, []byte{})
	ast.NoError(err)
	if err == nil {
		t.Logf("deposit tx hash is %s\n", tx.Hash().Hex())
	}

	// deposit receipt
	time.Sleep(5 * time.Second)
	receipt, err := l1cli.TransactionReceipt(context.Background(), tx.Hash())
	ast.NoError(err)
	ast.Equal(types.ReceiptStatusSuccessful, receipt.Status, "receipt status should be 0x01")
}

func Test_approveAndDepositWWQT(t *testing.T) {
	ast := assert.New(t)
	l1cli, err := ethclient.Dial(L1URL)
	ast.NoError(err)

	l1Bridge, err := abijson.NewL1StandardBridge(l1BridgeAddr, l1cli)
	ast.NoError(err)
	ast.NotNil(l1Bridge)
	l1WWQTContract, err := contract.NewTokenWwqERC20(wwqTokenAddrL1, l1cli)
	ast.NoError(err)
	ast.NotNil(l1WWQTContract)

	// approve
	t.Log("========================== approve =========================")
	chainID, err := l1cli.ChainID(context.Background())
	ast.NoError(err)
	opt, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	ast.NoError(err)
	//tx, err := l1WWQTContract.Approve(opt, crypto.PubkeyToAddress(priv.PublicKey), big.NewInt(1000000))
	tx, err := l1WWQTContract.Approve(opt, l1BridgeAddr, big.NewInt(1000000))
	ast.NoError(err)
	if err == nil {
		t.Logf("approve tx hash is %s\n", tx.Hash().Hex())
	}

	// approve receipt
	time.Sleep(10 * time.Second)
	receipt, err := l1cli.TransactionReceipt(context.Background(), tx.Hash())
	ast.NoError(err)
	ast.Equal(types.ReceiptStatusSuccessful, receipt.Status, "receipt status should be 0x01")

	// deposit
	t.Log("========================== deposit =========================")
	amount, minGasLimit := big.NewInt(10000), uint32(200000)
	tx, err = l1Bridge.DepositERC20(opt, wwqTokenAddrL1, wwqTokenAddrL2, amount, minGasLimit, []byte{})
	ast.NoError(err)
	if err == nil {
		t.Logf("deposit tx hash is %s\n", tx.Hash().Hex())
	}

	// deposit receipt
	time.Sleep(5 * time.Second)
	receipt, err = l1cli.TransactionReceipt(context.Background(), tx.Hash())
	ast.NoError(err)
	ast.Equal(types.ReceiptStatusSuccessful, receipt.Status, "receipt status should be 0x01")
}

func Test_approveForWithdrawWWQT(t *testing.T) {
	ast := assert.New(t)
	l2cli, err := ethclient.Dial(L2URL)
	ast.NoError(err)

	l2WWQTContract, err := contract.NewTokenWwqERC20(wwqTokenAddrL2, l2cli)
	ast.NoError(err)
	ast.NotNil(l2WWQTContract)

	// approve
	t.Log("========================== approve =========================")
	chainID, err := l2cli.ChainID(context.Background())
	ast.NoError(err)
	opt, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	ast.NoError(err)
	tx, err := l2WWQTContract.Approve(opt, l2BridgeAddr, big.NewInt(1000000))
	ast.NoError(err)
	if err == nil {
		t.Logf("approve tx hash is %s\n", tx.Hash().Hex())
	}

	// approve receipt
	time.Sleep(10 * time.Second)
	receipt, err := l2cli.TransactionReceipt(context.Background(), tx.Hash())
	ast.NoError(err)
	ast.Equal(types.ReceiptStatusSuccessful, receipt.Status, "receipt status should be 0x01")
}

func Test_withdrawWWQT(t *testing.T) {
	ast := assert.New(t)
	l2cli, err := ethclient.Dial(L2URL)
	ast.NoError(err)

	l2WWQTContract, err := contract.NewTokenWwqERC20(wwqTokenAddrL2, l2cli)
	ast.NoError(err)
	ast.NotNil(l2WWQTContract)
	l2Bridge, err := abijson.NewL2StandardBridge(l2BridgeAddr, l2cli)
	ast.NoError(err)
	ast.NotNil(l2Bridge)

	// withdraw
	t.Log("========================== withdraw =========================")
	chainID, err := l2cli.ChainID(context.Background())
	ast.NoError(err)
	opt, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	ast.NoError(err)
	ast.NotNil(opt)

	amount, minGasLimit := big.NewInt(10000), uint32(200000)
	tx, err := l2Bridge.Withdraw(opt, wwqTokenAddrL2, amount, minGasLimit, []byte{})
	ast.NoError(err)
	if err == nil {
		t.Logf("withdraw txhash is %s\n", tx.Hash().Hex()) // 0xb224e220e11eecb16da98a3540ffb2969a849c80a2b311fdc8351c7adb0a25a3
	}

	// receipt
	time.Sleep(5 * time.Second)
	receipt, err := l2cli.TransactionReceipt(context.Background(), tx.Hash())
	ast.NoError(err)
	ast.Equal(types.ReceiptStatusSuccessful, receipt.Status, "receipt status should be 0x1")

}

func Test_provenAndFinalize(t *testing.T) {
	ast := assert.New(t)
	l1cli, err := ethclient.Dial(L1URL)
	ast.NoError(err)
	l2cli, err := ethclient.Dial(L2URL)
	ast.NoError(err)

	var (
		L2OutputOracleProxy     = "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707"
		L1OptimismPortal        = "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853"
		wdTx                    = "0xb224e220e11eecb16da98a3540ffb2969a849c80a2b311fdc8351c7adb0a25a3"
		l2ToL1MessagePasserAddr = "0x4200000000000000000000000000000000000016"
		L1OptimismPortalAddr    = "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853"
	)

	withdrawReceipt, err := l2cli.TransactionReceipt(context.Background(), common.HexToHash(wdTx))
	ast.NoError(err)

	// prepare withdraw tx
	L2OutputOracleProxyAddr := common.HexToAddress(L2OutputOracleProxy)
	L2OutputOracleProxyContract, err := abijson.NewL2OutputOracleProxy(L2OutputOracleProxyAddr, l1cli)
	l2ToL1MessagePasserContract, err := abijson.NewL2ToL1MessagePasser(common.HexToAddress(l2ToL1MessagePasserAddr), l2cli)
	ast.NoError(err)

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
	blkNum, err := txutils.WaitForFinalizationPeriod(context.Background(), t, l1cli, common.HexToAddress(L1OptimismPortal), withdrawReceipt.BlockNumber, finalizationPeriod)
	ast.NoError(err)
	header, err := receiptCli.HeaderByNumber(context.Background(), big.NewInt(int64(blkNum)))
	ast.NoError(err)
	ast.NotNil(header)
	t.Logf("blkNum is %d, header.Num is %d\n", blkNum, header.Number)

	params, err := txutils.ProveWithdrawalParameters(context.Background(), t, l2GethCli, receiptCli, common.HexToHash(wdTx), header, &L2OutputOracleProxyContract.L2OutputOracleProxyCaller, l2ToL1MessagePasserContract)
	ast.NoError(err)

	// // prepare withdraw transaction
	wd := abijson.TypesWithdrawalTransaction{
		Nonce:    params.Nonce,
		Sender:   params.Sender,
		Target:   params.Target,
		MntValue: params.MNTValue,
		EthValue: params.ETHValue,
		GasLimit: params.GasLimit,
		Data:     params.Data,
	}
	t.Logf("wd is %+v\n", wd)

	// prepare proven
	t.Log("================= prepare proven =================")
	opPortal, err := abijson.NewL1OptimismPortal(common.HexToAddress(L1OptimismPortalAddr), l1cli)
	ast.NoError(err)

	t.Logf("wd %+v\n", wd)
	mytypeWd := txutils.GetMyWithdraw(&wd)
	slot, err := mytypeWd.StorageSlot()
	ast.NoError(err)

	proof, err := l2GethCli.GetProof(context.Background(), predeploys.L2ToL1MessagePasserAddr, []string{slot.String()}, big.NewInt(int64(blkNum)))
	ast.NoError(err)
	ast.NotNil(proof)
	t.Logf("proof.storage.len is %d", len(proof.StorageProof))

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
	wdHash, err := txutils.GetWdHash(&wd)
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

	l1ChainID, err := l1cli.ChainID(context.Background())
	if !alreadyProve {
		// proven
		t.Log("================= proven =================")
		ast.NoError(err)
		proveOpt, err := bind.NewKeyedTransactorWithChainID(priv, l1ChainID)
		ast.NoError(err)
		ast.NotNil(proveOpt)
		proveTx, err := opPortal.ProveWithdrawalTransaction(
			proveOpt,
			wd,
			params.L2OutputIndex,
			params.OutputRootProof,
			params.WithdrawalProof,
		)
		ast.NoError(err)
		ast.NotNil(proveTx)

		// waitMined
		receipt, err := bind.WaitMined(context.Background(), l1cli, proveTx)
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
	finOpt, err := bind.NewKeyedTransactorWithChainID(priv, l1ChainID)
	ast.NoError(err)
	fTx, err := opPortal.FinalizeWithdrawalTransaction(finOpt, wd)
	ast.NoError(err)
	ast.NotNil(fTx)
	receipt, err := bind.WaitMined(context.Background(), l1cli, fTx)
	ast.NoError(err)
	ast.Equal(types.ReceiptStatusSuccessful, receipt.Status)
}
