package main

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)

type ParamsCreateContract struct {
	From *common.Address
	A *common.Address
	B *common.Address
	Judge *common.Address
	Content *string
	FeePercentLimit *int
	Passphrase *string
}

func HandleCreateContract(state UserState, params *ParamsCreateContract, newParam string) *ParamsCreateContract {
	if params == nil{
		params = new(ParamsCreateContract)
	}
	switch state {
	case createUSDTContractState0:
		address := common.HexToAddress(newParam)
		params.From = &address
	case createUSDTContractState1:
		address := common.HexToAddress(newParam)
		params.A = &address
	case createUSDTContractState2:
		address := common.HexToAddress(newParam)
		params.B = &address
	case createUSDTContractState3:
		address := common.HexToAddress(newParam)
		params.Judge = &address
	case createUSDTContractState4:
		params.Content = &newParam
	case createUSDTContractState5:
		feePercentLimit, err := strconv.Atoi(newParam)
		if err != nil {
			return nil
		}
		params.FeePercentLimit = &feePercentLimit
	case createUSDTContractState6:
		params.Passphrase = &newParam
	default:
		return nil
	}
	return params
}