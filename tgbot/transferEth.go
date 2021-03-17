package main

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)

type ParamsTransferEth struct {
	From *common.Address
	To *common.Address
	Amount float64
	Passphrase *string
}

func HandleTransferEth(state UserState, params *ParamsTransferEth, newParam string) *ParamsTransferEth {
	if params == nil{
		params = new(ParamsTransferEth)
	}
	switch state {
	case transferEthState0:
		address := common.HexToAddress(newParam)
		params.From = &address
	case transferEthState1:
		address := common.HexToAddress(newParam)
		params.To = &address
	case transferEthState2:
		amount, err := strconv.ParseFloat(newParam, 64)
		if err != nil {
			return nil
		}
		params.Amount = amount
	case transferEthState3:
		params.Passphrase = &newParam
	default:
		return nil
	}
	return params
}
