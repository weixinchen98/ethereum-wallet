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
	//err = w.ImportAccount("913900ac615ff4e1ac5bce20f863434df4d1d2e104767b7a7e94bf664b544a3c", passphrase)
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

	//change password
	//err = w.ChangePassword(passphrase, newPassphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println("change password successfully.")
	//}

	//get account ether balance
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

	//get account erc-20 token balance
	tokenBalance, err := w.GetTokenBalance(wallet.RopsternTetherTokenAddress, accounts[0].Address)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(tokenBalance)
	}

	tokenBalance, err = w.GetTokenBalance(wallet.RopsternTetherTokenAddress, accounts[1].Address)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(tokenBalance)
	}

	//Transfer ether
	//txHash, err := w.TransferEther(accounts[0].Address, accounts[1].Address, 0.1, passphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println("Transaction hash: " + txHash)
	//}

	//Transfer erc-20 token
	//txHash, err := w.TransferToken(wallet.RopsternTetherTokenAddress, accounts[0].Address, accounts[1].Address, 1, passphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println("Transaction hash: " + txHash)
	//}

	//Approve ERC-20 token
	//contractAdr := common.HexToAddress("0x2E28EFADc79eB784a162eEF7F5ff393710232646")
	//txHash, err := w.ApproveToken(wallet.RopsternTetherTokenAddress, accounts[0].Address, contractAdr, 1, passphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println("Transaction hash: " + txHash)
	//}

	//Create contract
	txHash, err := w.USDTContractCreate(accounts[0].Address, accounts[0].Address, accounts[1].Address, accounts[2].Address, "test content", 20, passphrase)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Transaction hash: " + txHash)
	}


	//contract deposit
	//txHash, err := w.USDTContractDeposit(accounts[0].Address, contractAdr, 1, passphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println("Transaction hash: " + txHash)
	//}

	//contract withdraw
	//txHash, err := w.USDTContractWithdraw(accounts[2].Address, contractAdr, 0.5, 0.5, 10, passphrase)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println("Transaction hash: " + txHash)
	//}
}