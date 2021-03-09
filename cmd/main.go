package main

import (
	"ethereum-wallet/pkg"
	"fmt"
)

func main(){
	userId := "testId"
	passphrase := "testPass"

	w, err := wallet.NewWallet(userId, passphrase)
	if err != nil {
		fmt.Println(err)
	}

	//generate new account
	//err = w.GenerateAccount(passphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//import account
	//err := w.ImportAccount("88dd1d995dff4568cefd78840e851784cd505578b90b5c10cb15e3ef0f107294", passphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//get account address
	accounts := w.GetAllAccounts()
	for _, account := range(accounts) {
		fmt.Println(account.Address)
	}

	//delete account
	//err = w.DeleteAccount(common.HexToAddress("0x0B4690f6A73dA3d795A43499a872b6484866AF46"), passphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}


}