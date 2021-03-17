package main

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)

type ParamsWithdrawContract struct {
	From *common.Address
	To *common.Address
	AmountToA float64
	AmountToB float64
	FeePercent int
	Passphrase *string
}

func HandleWithdrawContract(state UserState, params *ParamsWithdrawContract, newParam string) *ParamsWithdrawContract {
	if params == nil{
		params = new(ParamsWithdrawContract)
	}
	switch state {
	case withdrawUSDTContractState0:
		address := common.HexToAddress(newParam)
		params.From = &address
	case withdrawUSDTContractState1:
		address := common.HexToAddress(newParam)
		params.To = &address
	case withdrawUSDTContractState2:
		amount, err := strconv.ParseFloat(newParam, 64)
		if err != nil {
			return nil
		}
		params.AmountToA = amount
	case withdrawUSDTContractState3:
		amount, err := strconv.ParseFloat(newParam, 64)
		if err != nil {
			return nil
		}
		params.AmountToB = amount
	case withdrawUSDTContractState4:
		feePercent, err := strconv.Atoi(newParam)
		if err != nil {
			return nil
		}
		params.FeePercent = feePercent
	case withdrawUSDTContractState5:
		params.Passphrase = &newParam
	default:
		return nil
	}
	return params
}