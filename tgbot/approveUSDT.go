package main

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)

type ParamsApproveUSDT struct {
	From *common.Address
	To *common.Address
	Amount float64
	Passphrase *string
}

func HandleApproveUSDT(state UserState, params *ParamsApproveUSDT, newParam string) *ParamsApproveUSDT {
	if params == nil{
		params = new(ParamsApproveUSDT)
	}

	switch state {
	case approveUSDTState0:
		address := common.HexToAddress(newParam)
		params.From = &address
	case approveUSDTState1:
		address := common.HexToAddress(newParam)
		params.To = &address
	case approveUSDTState2:
		amount, err := strconv.ParseFloat(newParam, 64)
		if err != nil {
			return nil
		}
		params.Amount = amount
	case approveUSDTState3:
		params.Passphrase = &newParam
	default:
		return nil
	}
	return params
}
