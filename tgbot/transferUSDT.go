package main

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)

type ParamsTransferUSDT struct {
	From *common.Address
	To *common.Address
	Amount float64
	Passphrase *string
}

func HandleTransferUSDT(state UserState, params *ParamsTransferUSDT, newParam string) *ParamsTransferUSDT {
	if params == nil{
		params = new(ParamsTransferUSDT)
	}
	switch state {
	case transferUSDTState0:
		address := common.HexToAddress(newParam)
		params.From = &address
	case transferUSDTState1:
		address := common.HexToAddress(newParam)
		params.To = &address
	case transferUSDTState2:
		amount, err := strconv.ParseFloat(newParam, 64)
		if err != nil {
			return nil
		}
		params.Amount = amount
	case transferUSDTState3:
		params.Passphrase = &newParam
	default:
		return nil
	}
	return params
}
