package main

type ParamsImportAccount struct {
	PrivKey *string
	Passphrase *string
}

func HandleImportAccount(state UserState, params *ParamsImportAccount, newParam string) *ParamsImportAccount {
	if params == nil{
		params = new(ParamsImportAccount)
	}
	switch state {
	case importState0:
		params.PrivKey = &newParam
	case importState1:
		params.Passphrase = &newParam
	default:
		return nil
	}
	return params
}
