package txutils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/stretchr/testify/assert"
	"try_rde/abistr/abijson"
)

var (
	// Standard ABI types
	Uint256Type, _ = abi.NewType("uint256", "", nil)
	BytesType, _   = abi.NewType("bytes", "", nil)
	AddressType, _ = abi.NewType("address", "", nil)
	Bytes32Type, _ = abi.NewType("bytes32", "", nil)
)

type Withdrawal struct {
	Nonce    *big.Int        `json:"nonce"`
	Sender   *common.Address `json:"sender"`
	Target   *common.Address `json:"target"`
	MntValue *big.Int        `json:"mntValue"`
	EthValue *big.Int        `json:"ethValue"`
	GasLimit *big.Int        `json:"gasLimit"`
	Data     hexutil.Bytes   `json:"data"`
}

func (w *Withdrawal) Encode() ([]byte, error) {
	args := abi.Arguments{
		{Name: "nonce", Type: Uint256Type},
		{Name: "sender", Type: AddressType},
		{Name: "target", Type: AddressType},
		{Name: "mntValue", Type: Uint256Type},
		{Name: "ethValue", Type: Uint256Type},
		{Name: "gasLimit", Type: Uint256Type},
		{Name: "data", Type: BytesType},
	}
	enc, err := args.Pack(w.Nonce, w.Sender, w.Target, w.MntValue, w.EthValue, w.GasLimit, []byte(w.Data))
	if err != nil {
		return nil, fmt.Errorf("cannot encode withdrawal: %w", err)
	}
	return enc, nil
}

func (w *Withdrawal) Hash() (common.Hash, error) {
	encoded, err := w.Encode()
	if err != nil {
		return common.Hash{}, err
	}
	hash := crypto.Keccak256(encoded)
	return common.BytesToHash(hash), nil
}

func GetMyWithdraw(wd *abijson.TypesWithdrawalTransaction) Withdrawal {
	return Withdrawal{
		Nonce:    wd.Nonce,
		Sender:   &wd.Sender,
		Target:   &wd.Target,
		MntValue: wd.MntValue,
		EthValue: wd.EthValue,
		GasLimit: wd.GasLimit,
		Data:     wd.Data,
	}
}
func GetWdHash(wd *abijson.TypesWithdrawalTransaction) (common.Hash, error) {
	w := &Withdrawal{
		Nonce:    wd.Nonce,
		Sender:   &wd.Sender,
		Target:   &wd.Target,
		MntValue: wd.MntValue,
		EthValue: wd.EthValue,
		GasLimit: wd.GasLimit,
		Data:     wd.Data,
	}
	return w.Hash()
}

func (w *Withdrawal) StorageSlot() (common.Hash, error) {
	hash, err := w.Hash()
	if err != nil {
		return common.Hash{}, err
	}
	preimage := make([]byte, 64)
	copy(preimage, hash.Bytes())

	slot := crypto.Keccak256(preimage)
	return common.BytesToHash(slot), nil
}

// ComputeL2OutputRoot computes the L2 output root by hashing an output root proof.
type Bytes32 [32]byte

func ComputeL2OutputRoot(proofElements *abijson.TypesOutputRootProof) (Bytes32, error) {
	if proofElements == nil {
		return Bytes32{}, errors.New("no proof elements")
	}

	digest := crypto.Keccak256Hash(
		proofElements.Version[:],
		proofElements.StateRoot[:],
		proofElements.MessagePasserStorageRoot[:],
		proofElements.LatestBlockhash[:],
	)
	return Bytes32(digest), nil
}

// ProvenWithdrawalParameters is the set of parameters to pass to the ProveWithdrawalTransaction
// and FinalizeWithdrawalTransaction functions
type ProvenWithdrawalParameters struct {
	Nonce           *big.Int
	Sender          common.Address
	Target          common.Address
	MNTValue        *big.Int
	ETHValue        *big.Int
	GasLimit        *big.Int
	L2OutputIndex   *big.Int
	Data            []byte
	OutputRootProof abijson.TypesOutputRootProof
	WithdrawalProof [][]byte // List of trie nodes to prove L2 storage
}

var MessagePassedTopic = crypto.Keccak256Hash([]byte("MessagePassed(uint256,address,address,uint256,uint256,uint256,bytes,bytes32)"))

// ParseMessagePassed parses MessagePassed events from
// a transaction receipt. It does not support multiple withdrawals
// per receipt.
func ParseMessagePassed(contract *abijson.L2ToL1MessagePasser, receipt *types.Receipt) (*abijson.L2ToL1MessagePasserMessagePassed, error) {
	//contract, err := abijson.NewL2ToL1MessagePasser(common.HexToAddress("0x4200000000000000000000000000000000000016"), nil)
	//if err != nil {
	//	return nil, err
	//}

	for _, log := range receipt.Logs {
		if len(log.Topics) == 0 || log.Topics[0] != MessagePassedTopic {
			continue
		}

		ev, err := contract.ParseMessagePassed(*log)
		if err != nil {
			return nil, fmt.Errorf("failed to parse log: %w", err)
		}
		return ev, nil
	}
	return nil, errors.New("Unable to find MessagePassed event")
}

func GetBlockNum(t *testing.T, submissionInterval *big.Int, startingBlockNumber *big.Int, l2BlockNumber *big.Int) *big.Int {
	t.Logf("GetBlockNum input blkNum is %d\n", l2BlockNumber)
	rem := new(big.Int)
	l2BlockNumber = l2BlockNumber.Sub(l2BlockNumber, startingBlockNumber)
	l2BlockNumber, rem = l2BlockNumber.DivMod(l2BlockNumber, submissionInterval, rem)
	if rem.Cmp(common.Big0) != 0 {
		l2BlockNumber = l2BlockNumber.Add(l2BlockNumber, common.Big1)
	}
	l2BlockNumber = l2BlockNumber.Mul(l2BlockNumber, submissionInterval)
	l2BlockNumber = l2BlockNumber.Add(l2BlockNumber, startingBlockNumber)
	t.Logf("GetBlockNum output blkNum is %d\n", l2BlockNumber)

	return l2BlockNumber
}

func WithdrawalHash(ev *abijson.L2ToL1MessagePasserMessagePassed) (common.Hash, error) {
	//  abi.encode(nonce, msg.sender, _target, msg.value, _gasLimit, _data)
	args := abi.Arguments{
		{Name: "nonce", Type: Uint256Type},
		{Name: "sender", Type: AddressType},
		{Name: "target", Type: AddressType},
		{Name: "mntValue", Type: Uint256Type},
		{Name: "ethValue", Type: Uint256Type},
		{Name: "gasLimit", Type: Uint256Type},
		{Name: "data", Type: BytesType},
	}
	enc, err := args.Pack(ev.Nonce, ev.Sender, ev.Target, ev.MntValue, ev.EthValue, ev.GasLimit, ev.Data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack for withdrawal hash: %w", err)
	}
	return crypto.Keccak256Hash(enc), nil
}

// ProveWithdrawalParameters queries L1 & L2 to generate all withdrawal parameters and proof necessary to prove a withdrawal on L1.
// The header provided is very important. It should be a block (timestamp) for which there is a submitted output in the L2 Output Oracle
// contract. If not, the withdrawal will fail as it the storage proof cannot be verified if there is no submitted state root.
func ProveWithdrawalParameters(ctx context.Context, t *testing.T, proofCl *gethclient.Client, l2ReceiptCl *ethclient.Client, txHash common.Hash, header *types.Header, l2OutputOracleContract *abijson.L2OutputOracleProxyCaller, l2ToL1MessagePasser *abijson.L2ToL1MessagePasser) (ProvenWithdrawalParameters, error) {
	ast := assert.New(t)
	// Transaction receipt
	receipt, err := l2ReceiptCl.TransactionReceipt(ctx, txHash)
	ast.NoError(err)
	if err != nil {
		return ProvenWithdrawalParameters{}, err
	}
	// Parse the receipt
	ev, err := ParseMessagePassed(l2ToL1MessagePasser, receipt)
	ast.NoError(err)
	if err != nil {
		return ProvenWithdrawalParameters{}, err
	}
	// Generate then verify the withdrawal proof
	withdrawalHash, err := WithdrawalHash(ev)
	ast.NoError(err)
	if !bytes.Equal(withdrawalHash[:], ev.WithdrawalHash[:]) {
		return ProvenWithdrawalParameters{}, errors.New("Computed withdrawal hash incorrectly")
	}
	if err != nil {
		return ProvenWithdrawalParameters{}, err
	}
	slot := StorageSlotOfWithdrawalHash(withdrawalHash)
	p, err := proofCl.GetProof(ctx, predeploys.L2ToL1MessagePasserAddr, []string{slot.String()}, header.Number)
	ast.NoError(err)
	if err != nil {
		return ProvenWithdrawalParameters{}, err
	}

	// Fetch the L2OutputIndex from the L2 Output Oracle caller (on L1)
	l2OutputIndex, err := l2OutputOracleContract.GetL2OutputIndexAfter(&bind.CallOpts{}, header.Number)
	ast.NoError(err)
	if err != nil {
		return ProvenWithdrawalParameters{}, fmt.Errorf("failed to get l2OutputIndex: %w", err)
	}
	//// TODO: Could skip this step, but it's nice to double check it
	//err = VerifyProof(header.Root, p)
	//if err != nil {
	//	return ProvenWithdrawalParameters{}, err
	//}
	//if len(p.StorageProof) != 1 {
	//	return ProvenWithdrawalParameters{}, errors.New("invalid amount of storage proofs")
	//}

	// Encode it as expected by the contract
	trieNodes := make([][]byte, len(p.StorageProof[0].Proof))
	for i, s := range p.StorageProof[0].Proof {
		trieNodes[i] = common.FromHex(s)
	}

	return ProvenWithdrawalParameters{
		Nonce:         ev.Nonce,
		Sender:        ev.Sender,
		Target:        ev.Target,
		MNTValue:      ev.MntValue,
		ETHValue:      ev.EthValue,
		GasLimit:      ev.GasLimit,
		L2OutputIndex: l2OutputIndex,
		Data:          ev.Data,
		OutputRootProof: abijson.TypesOutputRootProof{
			Version:                  [32]byte{}, // Empty for version 1
			StateRoot:                header.Root,
			MessagePasserStorageRoot: p.StorageHash,
			LatestBlockhash:          header.Hash(),
		},
		WithdrawalProof: trieNodes,
	}, nil
}

func StorageSlotOfWithdrawalHash(hash common.Hash) common.Hash {
	buf := make([]byte, 64)
	copy(buf, hash[:])
	return crypto.Keccak256Hash(buf)
}

func WaitForFinalizationPeriod(ctx context.Context, t *testing.T, client *ethclient.Client, portalAddr common.Address, l2BlockNumber *big.Int, finalizationPeriod *big.Int) (uint64, error) {
	l2BlockNumber = new(big.Int).Set(l2BlockNumber) // Don't clobber caller owned l2BlockNumber
	opts := &bind.CallOpts{Context: ctx}

	portal, err := abijson.NewL1OptimismPortal(portalAddr, client) //NewOptimismPortalCaller(portalAddr, client)
	if err != nil {
		return 0, err
	}
	l2OOAddress, err := portal.L2ORACLE(opts) // L2OutputOracleAddr
	if err != nil {
		return 0, err
	}
	t.Logf("[WaitForFinalizationPeriod] l2OOAddress is %s\n", l2OOAddress)
	l2OO, err := bindings.NewL2OutputOracleCaller(l2OOAddress, client)
	if err != nil {
		return 0, err
	}
	submissionInterval, err := l2OO.SUBMISSIONINTERVAL(opts)
	if err != nil {
		return 0, err
	}
	t.Logf("[WaitForFinalizationPeriod] submissionInterval is %d\n", submissionInterval)
	startingBlockNumber, err := l2OO.StartingBlockNumber(opts)
	if err != nil {
		return 0, err
	}
	t.Logf("[WaitForFinalizationPeriod] startingBlockNumber is %d\n", startingBlockNumber)

	// Convert blockNumber to submission interval boundary
	t.Logf("[WaitForFinalizationPeriod] before calculate, l2BlockNumber is %d\n", l2BlockNumber)
	rem := new(big.Int)
	l2BlockNumber = l2BlockNumber.Sub(l2BlockNumber, startingBlockNumber)             // -
	l2BlockNumber, rem = l2BlockNumber.DivMod(l2BlockNumber, submissionInterval, rem) // z = z(a, b, c);  z = a/b  c = a%b
	if rem.Cmp(common.Big0) != 0 {
		l2BlockNumber = l2BlockNumber.Add(l2BlockNumber, common.Big1)
	}
	l2BlockNumber = l2BlockNumber.Mul(l2BlockNumber, submissionInterval) // *
	l2BlockNumber = l2BlockNumber.Add(l2BlockNumber, startingBlockNumber)
	t.Logf("[WaitForFinalizationPeriod] after calculate, l2BlockNumber is %d\n", l2BlockNumber)

	//finalizationPeriod, err := l2OO.FINALIZATIONPERIODSECONDS(opts)  // proxy才有这个方法
	//if err != nil {
	//	return 0, err
	//}

	latest, err := l2OO.LatestBlockNumber(opts)
	if err != nil {
		return 0, err
	}
	t.Logf("[WaitForFinalizationPeriod] latest is %d\n", latest)

	// Now poll for the output to be submitted on chain
	var ticker *time.Ticker
	diff := new(big.Int).Sub(l2BlockNumber, latest)
	if diff.Cmp(big.NewInt(10)) > 0 {
		ticker = time.NewTicker(time.Minute)
	} else {
		ticker = time.NewTicker(time.Second)
	}

loop:
	for {
		select {
		case <-ticker.C:
			latest, err = l2OO.LatestBlockNumber(opts)
			if err != nil {
				return 0, err
			}
			// Already passed the submitted block (likely just equals rather than >= here).
			if latest.Cmp(l2BlockNumber) >= 0 {
				break loop
			}
		case <-ctx.Done():
			return 0, ctx.Err()
		}
	}

	// Now wait for it to be finalized
	output, err := l2OO.GetL2OutputAfter(opts, l2BlockNumber)
	if err != nil {
		return 0, err
	}
	if output.OutputRoot == [32]byte{} {
		return 0, errors.New("empty output root. likely no proposal at timestamp")
	}
	targetTimestamp := new(big.Int).Add(output.Timestamp, finalizationPeriod)
	targetTime := time.Unix(targetTimestamp.Int64(), 0)
	// Assume clock is relatively correct
	time.Sleep(time.Until(targetTime))
	// Poll for L1 Block to have a time greater than the target time
	ticker = time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			header, err := client.HeaderByNumber(ctx, nil)
			if err != nil {
				return 0, err
			}
			if header.Time > targetTimestamp.Uint64() {
				return l2BlockNumber.Uint64(), nil
			}
		case <-ctx.Done():
			return 0, ctx.Err()
		}
	}

}
