package main

import "github.com/ethereum/go-ethereum/common"

type ParamsCheckRelativeContract struct {
	Address *common.Address
}

func HandleCheckRelativeContract(state UserState, params *ParamsCheckRelativeContract, newParam string) *ParamsCheckRelativeContract {
	if params == nil{
		params = new(ParamsCheckRelativeContract)
	}

	switch state {
	case checkRelativeContractState:
		address := common.HexToAddress(newParam)
		params.Address = &address
	default:
		return nil
	}
	return params
}
