package main

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)

type ParamsDepositContract struct {
	From *common.Address
	To *common.Address
	Amount float64
	Passphrase *string
}

func HandleDepositContract(state UserState, params *ParamsDepositContract, newParam string) *ParamsDepositContract {
	if params == nil{
		params = new(ParamsDepositContract)
	}
	switch state {
	case depositUSDTContractState0:
		address := common.HexToAddress(newParam)
		params.From = &address
	case depositUSDTContractState1:
		address := common.HexToAddress(newParam)
		params.To = &address
	case depositUSDTContractState2:
		amount, err := strconv.ParseFloat(newParam, 64)
		if err != nil {
			return nil
		}
		params.Amount = amount
	case depositUSDTContractState3:
		params.Passphrase = &newParam
	default:
		return nil
	}
	return params
}

