package main

import (
	"ethereum-wallet/pkg"
	"fmt"
	"strconv"
)

func main(){
	userId := "testId"
	passphrase := "testPass"
	//newPassphrase := "newPss"

	w, err := wallet.NewWallet(userId, passphrase)
	if err != nil {
		fmt.Println(err)
	}

	err = w.ChangeNetwork(wallet.Ropsten)
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
	for i, account := range(accounts) {
		fmt.Print("Address of Account " + strconv.Itoa(i) +":")
		fmt.Println(account.Address)
	}

	//delete account
	//err = w.DeleteAccount(common.HexToAddress("0x0B4690f6A73dA3d795A43499a872b6484866AF46"), passphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//err = w.ChangePassword(passphrase, newPassphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println("change password successfully.")
	//}

	fmt.Print("Balance of Account 0: ")
	Balance, err := w.GetBalance(accounts[0].Address, passphrase)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(Balance)
	}

	fmt.Print("Balance of Account 1: ")
	Balance, err = w.GetBalance(accounts[1].Address, passphrase)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(Balance)
	}

	tokenBalance, err := w.GetTokenBalance(wallet.TetherTokenAddress, accounts[0].Address)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(tokenBalance)
	}

	tokenBalance, err = w.GetTokenBalance(wallet.TetherTokenAddress, accounts[1].Address)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(tokenBalance)
	}

	//txHash, err := w.TransferEther(accounts[0].Address, accounts[1].Address, 0.1, passphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(txHash)
	//}


}