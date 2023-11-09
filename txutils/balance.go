package txutils

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"try_rde/abistr/abijson"
)

//func getMNTBalanceFromL1(t *testing.T, address string) *big.Int {
//	client, err := ethclient.Dial("http://localhost:8545")
//	require.NoError(t, err)
//	require.NotNil(t, client)
//
//	l1MntInstance, err := bindings.NewL1MantleToken(common.HexToAddress(l1MntAddress), client)
//	require.NoError(t, err)
//	bal, err := l1MntInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
//	require.NoError(t, err)
//	require.NotNil(t, bal)
//	return bal
//}

//func getMNTBalanceFromL2(t *testing.T, address string) *big.Int {
//	client, err := ethclient.Dial("http://localhost:9545")
//	require.NoError(t, err)
//	require.NotNil(t, client)
//
//	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
//	require.NoError(t, err)
//	require.NotNil(t, balance)
//	return balance
//}

func GetETHBalanceFromL2(t *testing.T, address string) *big.Int {
	client, err := ethclient.Dial("http://localhost:9545")
	require.NoError(t, err)
	require.NotNil(t, client)
	WETH9Addr := "0xdEAddEaDdeadDEadDEADDEAddEADDEAddead1111" //"0x4200000000000000000000000000000000000006"

	l2ETHInst, err := abijson.NewL2TestToken(common.HexToAddress(WETH9Addr), client)
	assert.NoError(t, err)

	bal, err := l2ETHInst.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	assert.NoError(t, err)

	if err == nil {
		//t.Logf("getETHBalanceFromL2 balance is %d\n", bal)
		return bal
	}
	return nil
}
