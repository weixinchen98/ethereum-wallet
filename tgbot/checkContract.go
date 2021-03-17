package main

import "github.com/ethereum/go-ethereum/common"

type ParamsCheckContract struct {
	Address *common.Address
}

func HandleCheckContract(state UserState, params *ParamsCheckContract, newParam string) *ParamsCheckContract {
	if params == nil{
		params = new(ParamsCheckContract)
	}

	switch state {
	case checkContractState:
		address := common.HexToAddress(newParam)
		params.Address = &address
	default:
		return nil
	}
	return params
}