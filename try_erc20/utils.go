package try_erc20

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"try_rde/try_erc20/contract"
)

func getL1ContractAddress(erc20Factory *contract.OptimismMintableERC20Factory, receipt *types.Receipt, StandardL2TokenCreatedTopic common.Hash) (common.Address, error) {
	for _, log := range receipt.Logs {
		if len(log.Topics) == 0 || log.Topics[0] != StandardL2TokenCreatedTopic {
			continue
		}

		ev, err := erc20Factory.ParseStandardL2TokenCreated(*log)
		if err != nil {
			return [20]byte{}, err
		}
		return ev.LocalToken, nil
	}
	return [20]byte{}, errors.New("unable to find StandardL2TokenCreated event")
}
