package main

type ParamsCreateAccount struct {
	Passphrase *string
}

func HandleCreateAccount(state UserState, params *ParamsCreateAccount, newParam string) *ParamsCreateAccount {
	if params == nil{
		params = new(ParamsCreateAccount)
	}

	switch state {
	case createState:
		params.Passphrase = &newParam
	default:
		return nil
	}
	return params
}
