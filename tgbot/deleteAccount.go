package main

import "github.com/ethereum/go-ethereum/common"

type ParamsDeleteAccount struct {
	Address *common.Address
	Passphrase *string
}

func HandleDeleteAccount(state UserState, params *ParamsDeleteAccount, newParam string) *ParamsDeleteAccount {
	if params == nil{
		params = new(ParamsDeleteAccount)
	}
	switch state {
	case deleteState0:
		address := common.HexToAddress(newParam)
		params.Address = &address
	case deleteState1:
		params.Passphrase = &newParam
	default:
		return nil
	}
	return params
}
