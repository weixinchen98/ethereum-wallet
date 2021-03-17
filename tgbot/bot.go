package main

import (
	"context"
	"errors"
	"ethereum-wallet/contracts"
	"ethereum-wallet/pkg"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"

	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"strconv"
	"strings"
	"time"
)

const USDTDecimals = 6

type LogDeposit struct {
	Time  	*big.Int
	From    common.Address
	Amount 	*big.Int
}

type LogWithdraw struct {
	Time    	*big.Int
	AmountToA  	*big.Int
	AmountToB  	*big.Int
	FeePercent  *big.Int
}

type LogIntervene struct {
	Time    	*big.Int
	AmountToA  	*big.Int
	AmountToB  	*big.Int
	FeePercent  *big.Int
}

type UserState int

const(
	noState UserState = iota
	createState
	deleteState0
	deleteState1
	importState0
	importState1
	transferEthState0
	transferEthState1
	transferEthState2
	transferEthState3
	transferUSDTState0
	transferUSDTState1
	transferUSDTState2
	transferUSDTState3
	approveUSDTState0
	approveUSDTState1
	approveUSDTState2
	approveUSDTState3
	createUSDTContractState0
	createUSDTContractState1
	createUSDTContractState2
	createUSDTContractState3
	createUSDTContractState4
	createUSDTContractState5
	createUSDTContractState6
	depositUSDTContractState0
	depositUSDTContractState1
	depositUSDTContractState2
	depositUSDTContractState3
	withdrawUSDTContractState0
	withdrawUSDTContractState1
	withdrawUSDTContractState2
	withdrawUSDTContractState3
	withdrawUSDTContractState4
	withdrawUSDTContractState5
	checkRelativeContractState
	checkContractState
)

var userNetwork map[int]wallet.EthNet
var userState map[int]UserState
var userDeleteAccountParams map[int]*ParamsDeleteAccount
var userImportAccountParams map[int]*ParamsImportAccount
var userTransferEthParams map[int]*ParamsTransferEth
var userTransferUSDTParams map[int]*ParamsTransferUSDT
var userApproveUSDTParams map[int]*ParamsApproveUSDT
var userCreateContractParams map[int]*ParamsCreateContract
var userDepositContractParams map[int]*ParamsDepositContract
var userWithdrawContractParams map[int]*ParamsWithdrawContract
var userCheckRelativeContractParams map[int]*ParamsCheckRelativeContract
var userCheckContractParams map[int]*ParamsCheckContract




func main() {
	userNetwork = make(map[int]wallet.EthNet)
	userState = make(map[int]UserState)
	userDeleteAccountParams = make(map[int]*ParamsDeleteAccount)
	userImportAccountParams = make(map[int]*ParamsImportAccount)
	userTransferEthParams = make(map[int]*ParamsTransferEth)
	userTransferUSDTParams = make(map[int]*ParamsTransferUSDT)
	userApproveUSDTParams = make(map[int]*ParamsApproveUSDT)
	userCreateContractParams = make(map[int]*ParamsCreateContract)
	userDepositContractParams = make(map[int]*ParamsDepositContract)
	userWithdrawContractParams = make(map[int]*ParamsWithdrawContract)
	userCheckRelativeContractParams = make(map[int]*ParamsCheckRelativeContract)
	userCheckContractParams = make(map[int]*ParamsCheckContract)

	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		Token:  "1672120660:AAEburjQTE_B7RLPB2SFfLF406208C-6IOw",
		Poller: &tb.LongPoller{Timeout: 60 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}


	b.Handle("/help", func(m *tb.Message) {
		b.Send(m.Sender, "Hello dear user!")
		b.Send(m.Sender, "You can use command as follow \n" +
			"/help: to get commands list\n" +
			"/changeNetwork ethNet: to change ethereum network to ethNet (using Ropsten as defualt network)\n" +
			"/checkAccounts: to check accounts relative to your telegram id\n" +
			"/checkBalance: to check your accounts' Ether balance\n" +
			"/createAccount password: to create account relative to your telegram id by password\n" +
			"/deleteAccount accountAddress password: to delete account relative to your telegram id by password\n" +
			"/importAccount privateKey password: to import account relative to your telegram id by private key and password\n" +
			"/transferEth fromAddress toAddress amount password: \n" +
			"/transferUSDT fromAddress toAddress amount password: \n" +
			"/approveUSDT fromAddress toAddress amount password: \n" +
			"/createUSDTContract fromAddress addressA addressB addressJudge content feePercentLimit password: \n" +
			"/depositUSDTContract fromAddress contractAddress amount password: (should be after approving token confirmed) \n" +
			"/withdrawUSDTContract fromAddress contractAddress amountToA amountToB feePercent password: \n" +
			"/checkRelativeContract accountAddress: \n" +
			"/checkContract contractAddress: \n" )
	})

	b.Handle("/changeNetwork", func(m *tb.Message){
		userId := m.Sender.ID

		if userNetwork[userId] == "" {
			userNetwork[userId] = wallet.Ropsten
		}

		returnStr := "change network successfully!"

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/checkAccounts", func(m *tb.Message){
		userId := m.Sender.ID

		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		returnStr := getAccountsStr(w)

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/checkBalance", func(m *tb.Message){
		userId := m.Sender.ID

		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		returnStr := ""
		for _, account := range(w.GetAllAccounts()) {
			//get account ether balance
			returnStr += "Account address: " + account.Address.String()

			Balance, err := w.GetBalance(account.Address)
			if err != nil {
				b.Send(m.Sender, err.Error())
				return
			}
			returnStr += " Ether balance: " + Balance.String()

			//get account erc-20 token balance
			tokenBalance, err := w.GetTokenBalance(wallet.RopsternTetherTokenAddress, account.Address)
			if err != nil {
				b.Send(m.Sender, err.Error())
				return
			}
			returnStr += " Tether balance: " + tokenBalance.String() + "\n"
		}


		b.Send(m.Sender, returnStr)
	})

	b.Handle("/createAccount", func(m *tb.Message) {
		userId := m.Sender.ID

		passphrase := strings.Trim(strings.TrimPrefix(m.Text, "/createAccount"), " ")

		if len(passphrase) == 0 {
			userState[userId] = createState
			b.Send(m.Sender, "Please enter password for creating account:)")
			return
		}

			w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		err = w.GenerateAccount(passphrase)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		returnStr := "Create account successfully! \n" + getAccountsStr(w)

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/deleteAccount", func(m *tb.Message) {
		userId := m.Sender.ID

		text := strings.Split(strings.TrimPrefix(strings.TrimPrefix(m.Text, "/deleteAccount"), " "), " ")

		if err := checkParameters(2, text); err != nil  {
			userState[userId] = deleteState0
			b.Send(m.Sender, "Please enter address for deleting account:)")
			return
		}

		accountAddress := text[0]
		passphrase := text[1]

		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		err = w.DeleteAccount(common.HexToAddress(accountAddress), passphrase)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		accounts := w.GetAllAccounts()
		returnStr := "Create account successfully! \n" + accounts[len(accounts)-1].Address.String()

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/importAccount", func(m *tb.Message) {
		userId := m.Sender.ID

		text := strings.Split(strings.TrimPrefix(strings.TrimPrefix(m.Text, "/importAccount"), " "), " ")

		if err := checkParameters(2, text); err != nil  {
			userState[userId] = importState0
			b.Send(m.Sender, "Please enter private key for importing account :)")
			return
		}

		privateKey := text[0]
		passphrase := text[1]

		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		err = w.ImportAccount(privateKey, passphrase)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		returnStr := "Import account successfully! \n" + getAccountsStr(w)

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/transferEth", func(m *tb.Message) {
		userId := m.Sender.ID

		text := strings.Split(strings.TrimPrefix(strings.TrimPrefix(m.Text, "/transferEth"), " "), " ")

		if err := checkParameters(4, text); err != nil  {
			userState[userId] = transferEthState0
			b.Send(m.Sender, "Please enter from address for transferring eth :)")
			return
		}

		fromAddress := text[0]
		toAddress := text[1]
		amount, err := strconv.ParseFloat(text[2], 64)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		passphrase := text[3]


		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		txHash, err := w.TransferEther(common.HexToAddress(fromAddress), common.HexToAddress(toAddress), amount, passphrase)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		returnStr := "Transfer Ether successfully! \n" + txHash

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/transferUSDT", func(m *tb.Message) {
		userId := m.Sender.ID

		text := strings.Split(strings.TrimPrefix(strings.TrimPrefix(m.Text, "/transferUSDT"), " "), " ")

		if err := checkParameters(4, text); err != nil  {
			userState[userId] = transferUSDTState0
			b.Send(m.Sender, "Please enter from address for transferring USDT :)")
			return
		}

		fromAddress := text[0]
		toAddress := text[1]
		amount, err := strconv.ParseFloat(text[2], 64)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		passphrase := text[3]


		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		txHash, err := w.TransferToken(wallet.RopsternTetherTokenAddress, common.HexToAddress(fromAddress), common.HexToAddress(toAddress), amount, passphrase)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		returnStr := "Transfer USDT successfully! \n" + txHash

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/approveUSDT", func(m *tb.Message) {
		userId := m.Sender.ID

		text := strings.Split(strings.TrimPrefix(strings.TrimPrefix(m.Text, "/approveUSDT"), " "), " ")

		if err := checkParameters(4, text); err != nil  {
			userState[userId] = approveUSDTState0
			b.Send(m.Sender, "Please enter from address for approving USDT :)")
			return
		}

		fromAddress := text[0]
		toAddress := text[1]
		amount, err := strconv.ParseFloat(text[2], 64)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		passphrase := text[3]


		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		txHash, err := w.ApproveToken(wallet.RopsternTetherTokenAddress, common.HexToAddress(fromAddress), common.HexToAddress(toAddress), amount, passphrase)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		returnStr := "Approve USDT successfully! \n" + txHash

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/createUSDTContract", func(m *tb.Message) {
		userId := m.Sender.ID

		text := strings.Split(strings.TrimPrefix(strings.TrimPrefix(m.Text, "/createUSDTContract"), " "), " ")

		if err := checkParameters(7, text); err != nil  {
			userState[userId] = createUSDTContractState0
			b.Send(m.Sender, "Please enter from address for creating contract :)")
			return
		}

		fromAddress := common.HexToAddress(text[0])
		addressA := common.HexToAddress(text[1])
		addressB := common.HexToAddress(text[2])
		addressJudge := common.HexToAddress(text[3])
		content := text[4]
		feePercentLimit, err := strconv.Atoi(text[5])
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		passphrase := text[6]


		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		txHash, err := w.USDTContractCreate(fromAddress, addressA, addressB, addressJudge, content, int64(feePercentLimit), passphrase)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		returnStr := "Create USDT contract successfully! \n" + txHash

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/depositUSDTContract", func(m *tb.Message) {
		userId := m.Sender.ID

		text := strings.Split(strings.TrimPrefix(strings.TrimPrefix(m.Text, "/depositUSDTContract"), " "), " ")

		if err := checkParameters(4, text); err != nil  {
			userState[userId] = depositUSDTContractState0
			b.Send(m.Sender, "Please enter from address for depositing contract :)")
			return
		}

		fromAddress := common.HexToAddress(text[0])
		contractAddress := common.HexToAddress(text[1])
		amount, err := strconv.ParseFloat(text[2], 64)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		passphrase := text[3]


		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		txHash, err := w.USDTContractDeposit(fromAddress, contractAddress, amount, passphrase)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		returnStr := "Deposit USDT contract successfully! \n" + txHash

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/withdrawUSDTContract", func(m *tb.Message) {
		userId := m.Sender.ID

		text := strings.Split(strings.TrimPrefix(strings.TrimPrefix(m.Text, "/withdrawUSDTContract"), " "), " ")

		if err := checkParameters(6, text); err != nil  {
			userState[userId] = withdrawUSDTContractState0
			b.Send(m.Sender, "Please enter from address for withdrawing contract :)")
			return
		}

		fromAddress := common.HexToAddress(text[0])
		contractAddress := common.HexToAddress(text[1])
		amountToA, err := strconv.ParseFloat(text[2], 64)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		amountToB, err := strconv.ParseFloat(text[3], 64)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		feePercent, err := strconv.Atoi(text[4])
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		passphrase := text[5]


		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		txHash, err := w.USDTContractWithdraw(fromAddress, contractAddress, amountToA, amountToB, int64(feePercent), passphrase)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		returnStr := "Withdraw USDT contract successfully! \n" + txHash

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/checkRelativeContract", func(m *tb.Message){
		userId := m.Sender.ID

		text := strings.Trim(strings.TrimPrefix(m.Text, "/checkRelativeContract"), " ")

		if len(text) == 0 {
			userState[userId] = checkRelativeContractState
			b.Send(m.Sender, "Please enter from address for checking relative contract :)")
			return
		}

		accountAddress := common.HexToAddress(text)

		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		contractsAddress, err := getRelativeContracts(accountAddress, w.GetClient())
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		returnStr := "查询到" + strconv.Itoa(len(contractsAddress)) + "个合约: \n"
		for i, contractAddress := range(contractsAddress){
			returnStr += "第" + strconv.Itoa(i) + "个合约: "
			contractStr, err := getContractInfo(contractAddress, w.GetClient())
			if err != nil {
				b.Send(m.Sender, err.Error())
				return
			}
			returnStr += contractStr + "\n"
		}

		b.Send(m.Sender, returnStr)
	})

	b.Handle("/checkContract", func(m *tb.Message){
		userId := m.Sender.ID

		text := strings.Trim(strings.TrimPrefix(m.Text, "/checkContract"), " ")

		if len(text) == 0 {
			userState[userId] = checkContractState
			b.Send(m.Sender, "Please enter contract address for checking contract :)")
			return
		}

		contractAddress := common.HexToAddress(text)

		w, err := wallet.NewWallet(strconv.Itoa(userId))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		initNetwork(userId, w)

		returnStr, err := getContractInfo(contractAddress, w.GetClient())
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		b.Send(m.Sender, returnStr)

	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		userId := m.Sender.ID
		newParam := strings.Trim(m.Text, " ")

		w, err := getUserWallet(userId)
		if err != nil{
			b.Send(m.Sender, err.Error())
			return
		}

		switch userState[userId] {
		case noState:
			b.Send(m.Sender, "Enter '/help' to get help.")

		case createState:
			createAccountParams := HandleCreateAccount(createState, nil, newParam)
			err = w.GenerateAccount(*createAccountParams.Passphrase)
			if err != nil {
				b.Send(m.Sender, err.Error())
				userState[userId] = noState
				return
			}
			accounts := w.GetAllAccounts()
			returnStr := "Create account successfully! \n" + accounts[len(accounts)-1].Address.String()
			b.Send(m.Sender, returnStr)
			userState[userId] = noState

		case deleteState0:
			userDeleteAccountParams[userId] = HandleDeleteAccount(deleteState0, nil, newParam)
			b.Send(m.Sender, "Please enter password for deleting account:)")
			userState[userId] = deleteState1
		case deleteState1:
			userDeleteAccountParams[userId] = HandleDeleteAccount(deleteState1, userDeleteAccountParams[userId], newParam)
			err = w.DeleteAccount(*userDeleteAccountParams[userId].Address, *userDeleteAccountParams[userId].Passphrase)
			if err != nil {
				b.Send(m.Sender, err.Error())
				userState[userId] = noState
				return
			}
			returnStr := "Delete account successfully! \n" + getAccountsStr(w)
			b.Send(m.Sender, returnStr)
			userState[userId] = noState

		case importState0:
			userImportAccountParams[userId] = HandleImportAccount(importState0, nil, newParam)
			b.Send(m.Sender, "Please enter password for importing account:)")
			userState[userId] = importState1
		case importState1:
			userImportAccountParams[userId] = HandleImportAccount(importState1, userImportAccountParams[userId], newParam)
			err := w.ImportAccount(*userImportAccountParams[userId].PrivKey, *userImportAccountParams[userId].Passphrase)
			if err != nil {
				b.Send(m.Sender, err.Error())
				userState[userId] = noState
				return
			}
			accounts := w.GetAllAccounts()
			returnStr := "Import account successfully! \n" + accounts[len(accounts)-1].Address.String()
			b.Send(m.Sender, returnStr)
			userState[userId] = noState

		case transferEthState0:
			userTransferEthParams[userId] = HandleTransferEth(transferEthState0, nil, newParam)
			b.Send(m.Sender, "Please enter to address for transferring eth :)")
			userState[userId] = transferEthState1
		case transferEthState1:
			userTransferEthParams[userId] = HandleTransferEth(transferEthState1, userTransferEthParams[userId], newParam)
			b.Send(m.Sender, "Please enter amount for transferring eth :)")
			userState[userId] = transferEthState2
		case transferEthState2:
			userTransferEthParams[userId] = HandleTransferEth(transferEthState2, userTransferEthParams[userId], newParam)
			b.Send(m.Sender, "Please enter password for transferring eth account :)")
			userState[userId] = transferEthState3
		case transferEthState3:
			userTransferEthParams[userId] = HandleTransferEth(transferEthState3, userTransferEthParams[userId], newParam)

			txHash, err := w.TransferEther(*userTransferEthParams[userId].From, *userTransferEthParams[userId].To, userTransferEthParams[userId].Amount, *userTransferEthParams[userId].Passphrase)
			if err != nil {
				b.Send(m.Sender, err.Error())
				userState[userId] = noState
				return
			}

			returnStr := "Transfer Ether successfully! \n" + txHash

			b.Send(m.Sender, returnStr)
			userState[userId] = noState

		case transferUSDTState0:
			userTransferUSDTParams[userId] = HandleTransferUSDT(transferUSDTState0, nil, newParam)
			b.Send(m.Sender, "Please enter to address for transferring USDT :)")
			userState[userId] = transferUSDTState1
		case transferUSDTState1:
			userTransferUSDTParams[userId] = HandleTransferUSDT(transferUSDTState1, userTransferUSDTParams[userId], newParam)
			b.Send(m.Sender, "Please enter amount for transferring USDT :)")
			userState[userId] = transferUSDTState2
		case transferUSDTState2:
			userTransferUSDTParams[userId] = HandleTransferUSDT(transferUSDTState2, userTransferUSDTParams[userId], newParam)
			b.Send(m.Sender, "Please enter password for transferring USDT account:)")
			userState[userId] = transferUSDTState3
		case transferUSDTState3:
			userTransferUSDTParams[userId] = HandleTransferUSDT(transferUSDTState3, userTransferUSDTParams[userId], newParam)

			txHash, err := w.TransferToken(wallet.RopsternTetherTokenAddress, *userTransferUSDTParams[userId].From, *userTransferUSDTParams[userId].To, userTransferUSDTParams[userId].Amount, *userTransferUSDTParams[userId].Passphrase)
			if err != nil {
				b.Send(m.Sender, err.Error())
				userState[userId] = noState
				return
			}

			returnStr := "Transfer USDT successfully! \n" + txHash

			b.Send(m.Sender, returnStr)
			userState[userId] = noState

		case approveUSDTState0:
			userApproveUSDTParams[userId] = HandleApproveUSDT(approveUSDTState0, nil, newParam)
			b.Send(m.Sender, "Please enter to address for approving USDT :)")
			userState[userId] = approveUSDTState1
		case approveUSDTState1:
			userApproveUSDTParams[userId] = HandleApproveUSDT(approveUSDTState0, nil, newParam)
			b.Send(m.Sender, "Please enter amount for approving USDT :)")
			userState[userId] = approveUSDTState2
		case approveUSDTState2:
			userApproveUSDTParams[userId] = HandleApproveUSDT(approveUSDTState0, nil, newParam)
			b.Send(m.Sender, "Please enter password for approving USDT account:)")
			userState[userId] = approveUSDTState3
		case approveUSDTState3:
			userApproveUSDTParams[userId] = HandleApproveUSDT(approveUSDTState0, nil, newParam)

			txHash, err := w.ApproveToken(wallet.RopsternTetherTokenAddress, *userApproveUSDTParams[userId].From, *userApproveUSDTParams[userId].To, userApproveUSDTParams[userId].Amount, *userApproveUSDTParams[userId].Passphrase)
			if err != nil {
				b.Send(m.Sender, err.Error())
				userState[userId] = noState
				return
			}

			returnStr := "Approve USDT successfully! \n" + txHash

			b.Send(m.Sender, returnStr)
			userState[userId] = noState

		case createUSDTContractState0:
			userCreateContractParams[userId] = HandleCreateContract(createUSDTContractState0, nil, newParam)
			b.Send(m.Sender, "Please enter address A for creating contract :)")
			userState[userId] = createUSDTContractState1
		case createUSDTContractState1:
			userCreateContractParams[userId] = HandleCreateContract(createUSDTContractState1, userCreateContractParams[userId], newParam)
			b.Send(m.Sender, "Please enter address B for creating contract :)")
			userState[userId] = createUSDTContractState2
		case createUSDTContractState2:
			userCreateContractParams[userId] = HandleCreateContract(createUSDTContractState2, userCreateContractParams[userId], newParam)
			b.Send(m.Sender, "Please enter address Judge for creating contract :)")
			userState[userId] = createUSDTContractState3
		case createUSDTContractState3:
			userCreateContractParams[userId] = HandleCreateContract(createUSDTContractState3, userCreateContractParams[userId], newParam)
			b.Send(m.Sender, "Please enter content for creating contract :)")
			userState[userId] = createUSDTContractState4
		case createUSDTContractState4:
			userCreateContractParams[userId] = HandleCreateContract(createUSDTContractState4, userCreateContractParams[userId], newParam)
			b.Send(m.Sender, "Please enter fee percent limit(%) for creating contract :)")
			userState[userId] = createUSDTContractState5
		case createUSDTContractState5:
			userCreateContractParams[userId] = HandleCreateContract(createUSDTContractState5, userCreateContractParams[userId], newParam)
			b.Send(m.Sender, "Please enter password of from address for creating contract :)")
			userState[userId] = createUSDTContractState6
		case createUSDTContractState6:
			userCreateContractParams[userId] = HandleCreateContract(createUSDTContractState6, userCreateContractParams[userId], newParam)

			fromAddress := *userCreateContractParams[userId].From
			addressA := *userCreateContractParams[userId].A
			addressB := *userCreateContractParams[userId].B
			addressJudge := *userCreateContractParams[userId].Judge
			content := *userCreateContractParams[userId].Content
			feePercentLimit := *userCreateContractParams[userId].FeePercentLimit
			passphrase := *userCreateContractParams[userId].Passphrase
			txHash, err := w.USDTContractCreate(fromAddress, addressA, addressB, addressJudge, content, int64(feePercentLimit), passphrase)
			if err != nil {
				b.Send(m.Sender, err.Error())
				userState[userId] = noState
				return
			}

			returnStr := "Create USDT contract successfully! \n" + txHash

			b.Send(m.Sender, returnStr)

			userState[userId] = noState

		case depositUSDTContractState0:
			userDepositContractParams[userId] = HandleDepositContract(depositUSDTContractState0, nil, newParam)
			b.Send(m.Sender, "Please enter contract address for depositing contract :)")
			userState[userId] = depositUSDTContractState1
		case depositUSDTContractState1:
			userDepositContractParams[userId] = HandleDepositContract(depositUSDTContractState1, userDepositContractParams[userId], newParam)
			b.Send(m.Sender, "Please enter amount for depositing contract :)")
			userState[userId] = depositUSDTContractState2
		case depositUSDTContractState2:
			userDepositContractParams[userId] = HandleDepositContract(depositUSDTContractState1, userDepositContractParams[userId], newParam)
			b.Send(m.Sender, "Please enter password of from address for depositing contract :)")
			userState[userId] = depositUSDTContractState3
		case depositUSDTContractState3:
			userDepositContractParams[userId] = HandleDepositContract(depositUSDTContractState3, nil, newParam)
			fromAddress := *userDepositContractParams[userId].From
			contractAddress := *userDepositContractParams[userId].To
			amount := userDepositContractParams[userId].Amount
			passphrase := *userDepositContractParams[userId].Passphrase
			txHash, err := w.USDTContractDeposit(fromAddress, contractAddress, amount, passphrase)
			if err != nil {
				b.Send(m.Sender, err.Error())
				userState[userId] = noState
				return
			}

			returnStr := "Deposit USDT contract successfully! \n" + txHash

			b.Send(m.Sender, returnStr)
			userState[userId] = noState

		case withdrawUSDTContractState0:
			userWithdrawContractParams[userId] = HandleWithdrawContract(withdrawUSDTContractState0, nil, newParam)
			b.Send(m.Sender, "Please enter contract address for withdrawing contract :)")
			userState[userId] = withdrawUSDTContractState1
		case withdrawUSDTContractState1:
			userWithdrawContractParams[userId] = HandleWithdrawContract(withdrawUSDTContractState1, userWithdrawContractParams[userId], newParam)
			b.Send(m.Sender, "Please enter amount to A for withdrawing contract :)")
			userState[userId] = withdrawUSDTContractState2
		case withdrawUSDTContractState2:
			userWithdrawContractParams[userId] = HandleWithdrawContract(withdrawUSDTContractState2, userWithdrawContractParams[userId], newParam)
			b.Send(m.Sender, "Please enter amount to B for withdrawing contract :)")
			userState[userId] = withdrawUSDTContractState3
		case withdrawUSDTContractState3:
			userWithdrawContractParams[userId] = HandleWithdrawContract(withdrawUSDTContractState3, userWithdrawContractParams[userId], newParam)
			b.Send(m.Sender, "Please enter fee percent(int, %) for withdrawing contract :)")
			userState[userId] = withdrawUSDTContractState4
		case withdrawUSDTContractState4:
			userWithdrawContractParams[userId] = HandleWithdrawContract(withdrawUSDTContractState4, userWithdrawContractParams[userId], newParam)
			b.Send(m.Sender, "Please enter password of from account for withdrawing contract :)")
			userState[userId] = withdrawUSDTContractState5
		case withdrawUSDTContractState5:
			userWithdrawContractParams[userId] = HandleWithdrawContract(withdrawUSDTContractState4, userWithdrawContractParams[userId], newParam)

			fromAddress := *userWithdrawContractParams[userId].From
			contractAddress := *userWithdrawContractParams[userId].To
			amountToA := userWithdrawContractParams[userId].AmountToA
			amountToB := userWithdrawContractParams[userId].AmountToB
			feePercent:= userWithdrawContractParams[userId].FeePercent
			passphrase:= *userWithdrawContractParams[userId].Passphrase

			txHash, err := w.USDTContractWithdraw(fromAddress, contractAddress, amountToA, amountToB, int64(feePercent), passphrase)
			if err != nil {
				b.Send(m.Sender, err.Error())
				userState[userId] = noState
				return
			}

			returnStr := "Withdraw USDT contract successfully! \n" + txHash

			b.Send(m.Sender, returnStr)

			userState[userId] = noState

		case checkRelativeContractState:
			userCheckRelativeContractParams[userId] = HandleCheckRelativeContract(checkRelativeContractState, nil, newParam)
			accountAddress := *userCheckRelativeContractParams[userId].Address
			contractsAddress, err := getRelativeContracts(accountAddress, w.GetClient())
			if err != nil {
				b.Send(m.Sender, err.Error())
				userState[userId] = noState
				return
			}

			returnStr := "查询到" + strconv.Itoa(len(contractsAddress)) + "个合约: \n"
			for i, contractAddress := range(contractsAddress){
				returnStr += "第" + strconv.Itoa(i) + "个合约: "
				contractStr, err := getContractInfo(contractAddress, w.GetClient())
				if err != nil {
					b.Send(m.Sender, err.Error())
					userState[userId] = noState
					return
				}
				returnStr += contractStr + "\n"
			}

			b.Send(m.Sender, returnStr)
			userState[userId] = noState

		case checkContractState:
			userCheckContractParams[userId] = HandleCheckContract(checkContractState, nil, newParam)

			contractAddress := *userCheckContractParams[userId].Address
			returnStr, err := getContractInfo(contractAddress, w.GetClient())
			if err != nil {
				b.Send(m.Sender, err.Error())
				userState[userId] = noState
				return
			}

			b.Send(m.Sender, returnStr)

			userState[userId] = noState
		}



	})

	b.Start()
}

func getUserWallet(userId int) (*wallet.Wallet, error) {
	w, err := wallet.NewWallet(strconv.Itoa(userId))
	if err != nil {
		return nil, err
	}
	initNetwork(userId, w)
	return w, nil
}

func getRelativeContracts(accountAddress common.Address, client *ethclient.Client) ([]common.Address, error) {
	ServerAdr := common.HexToAddress(wallet.RopstenServerAddress)
	instance, err := contracts.NewSever(ServerAdr, client)
	if err != nil {
		return nil, err
	}

	contractNum, err := instance.RelationsNum(nil, accountAddress)
	if err != nil {
		return nil, err
	}

	res := make([]common.Address, contractNum.Int64())

	for i := 0; i < int(contractNum.Int64()); i++ {
		contractId, err := instance.Relations(nil, accountAddress, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		contractAddress, err := instance.Contracts(nil, contractId)
		if err != nil {
			return nil, err
		}

		res[i] = contractAddress
	}

	return res, nil
}

func getContractInfo(contractAddress common.Address, client *ethclient.Client) (string, error) {
	returnStr := ""
	instance, err := contracts.NewContract(contractAddress, client)
	if err != nil {
		return "", err
	}

	balance, err := instance.Balance(nil)
	if err != nil {
		return "", err
	}
	addressA, err := instance.A(nil)
	if err != nil {
		return "", err
	}
	addressB, err := instance.B(nil)
	if err != nil {
		return "", err
	}
	addressJudge, err := instance.Judge(nil)
	if err != nil {
		return "", err
	}
	content, err := instance.Content(nil)
	if err != nil {
		return "", err
	}
	feePercentLimit, err := instance.FeePercentLimit(nil)
	if err != nil {
		return "", err
	}

	returnStr += "address A: " + addressA.String() +
		" \naddress B: " + addressB.String() +
		" \naddress Judge: " + addressJudge.String() +
		" \n内容: " + content + ", 担保费率上限: " + feePercentLimit.String() + "%, 合约余额: " + balance.String() + " USDT \n"

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return "", err
	}

	contractAbi, err := abi.JSON(strings.NewReader(contracts.ContractABI))
	if err != nil {
		log.Fatal(err)
	}

	logDepositSig := []byte("Deposit(uint256,address,uint256)")
	logWithdrawSig := []byte("Withdraw(uint256,uint256,uint256,uint256)")
	logInterveneSig := []byte("Intervene(uint256,uint256,uint256,uint256)")

	logDepositSigHash := crypto.Keccak256Hash(logDepositSig)
	logWithdrawSigHash := crypto.Keccak256Hash(logWithdrawSig)
	logInterveneSigHash := crypto.Keccak256Hash(logInterveneSig)

	for _, vLog := range logs{
		switch vLog.Topics[0].Hex() {
		case logDepositSigHash.Hex():
			var depositEvent LogDeposit
			err := contractAbi.UnpackIntoInterface(&depositEvent, "Deposit", vLog.Data)
			if err != nil {
				return "", err
			}
			amount := new(big.Float).Quo(new(big.Float).SetInt(depositEvent.Amount), big.NewFloat(math.Pow10(USDTDecimals)))

			returnStr += formatTimestamp(depositEvent.Time) + " " + depositEvent.From.String() + " 转入 " + amount.String() + " USDT \n"
		case logWithdrawSigHash.Hex():
			var withdrawEvent LogWithdraw
			err := contractAbi.UnpackIntoInterface(&withdrawEvent, "Withdraw", vLog.Data)
			if err != nil {
				return "", err
			}

			amountToA := new(big.Float).Quo(new(big.Float).SetInt(withdrawEvent.AmountToA), big.NewFloat(math.Pow10(USDTDecimals)))
			amountToB := new(big.Float).Quo(new(big.Float).SetInt(withdrawEvent.AmountToB), big.NewFloat(math.Pow10(USDTDecimals)))

			feePercent := new(big.Float).Quo(new(big.Float).SetInt(withdrawEvent.FeePercent), big.NewFloat(100.0))
			feeA := new(big.Float).Mul(amountToA, feePercent)
			feeB := new(big.Float).Mul(amountToB, feePercent)

			amountToA = new(big.Float).Sub(amountToA, feeA)
			amountToB = new(big.Float).Sub(amountToB, feeB)
			fee := new(big.Float).Add(feeA, feeB)

			returnStr += formatTimestamp(withdrawEvent.Time)  + " " +
				amountToA.String() + " USDT 转出到 A, " +
				amountToB.String()  + " USDT 转出到 B, " +
				fee.String() + " USDT 转出到 Judge \n"
		case logInterveneSigHash.Hex():
			var interveneEvent LogIntervene
			err := contractAbi.UnpackIntoInterface(&interveneEvent, "Withdraw", vLog.Data)
			if err != nil {
				return "", err
			}


			amountToA := new(big.Float).Quo(new(big.Float).SetInt(interveneEvent.AmountToA), big.NewFloat(math.Pow10(USDTDecimals)))
			amountToB := new(big.Float).Quo(new(big.Float).SetInt(interveneEvent.AmountToB), big.NewFloat(math.Pow10(USDTDecimals)))

			feePercent := new(big.Float).Quo(new(big.Float).SetInt(interveneEvent.FeePercent), big.NewFloat(100.0))
			feeA := new(big.Float).Mul(amountToA, feePercent)
			feeB := new(big.Float).Mul(amountToB, feePercent)

			amountToA = new(big.Float).Sub(amountToA, feeA)
			amountToB = new(big.Float).Sub(amountToB, feeB)
			fee := new(big.Float).Add(feeA, feeB)

			returnStr += formatTimestamp(interveneEvent.Time) + " " +
				amountToA.String() + " USDT 转出到 A, " +
				amountToB.String()  + " USDT 转出到 B, " +
				fee.String() + " USDT 转出到 Server \n"
		}
	}
	return returnStr, nil
}

func getAccountsStr(w *wallet.Wallet) (res string) {
	accounts := w.GetAllAccounts()
	for i, account := range(accounts) {
		res += "Address of Account " + strconv.Itoa(i) +":" + account.Address.String() + "\n"
	}
	return res
}

func initNetwork(userId int, w *wallet.Wallet){
	if userNetwork[userId] == "" {
		userNetwork[userId] = wallet.Ropsten
	}
	w.ChangeNetwork(userNetwork[userId])
}

func checkParameters(num int, text []string) error {
	if len(text) < num {
		return errors.New("Not enough parameters ")
	}
	return nil
}

func formatTimestamp(timestamp *big.Int) string{
	tm := time.Unix(timestamp.Int64(), 0)
	return tm.String()
}

