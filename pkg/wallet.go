package wallet

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"ethereum-wallet/contracts"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

const keyStorePath = "../keystores/"

type TokenAddress string

type EthNet string

const (
	Ropsten EthNet = "https://ropsten.infura.io/v3/cbf6f482ffa6444c88c16c67aebbd738"
	Main EthNet = "https://mainnet.infura.io/v3/cbf6f482ffa6444c88c16c67aebbd738"
)

const (
	RopsternTetherTokenAddress TokenAddress = "0x0881DC670828Dc74Dc9AdE68ec328a579Dc1E660"
)

const RopstenServerAddress = "0xb390dCA0dA832a8Ff93f6Ee10835352f3321286d"

type Wallet struct {
	userId string
	ks *keystore.KeyStore
	client *ethclient.Client
}

func NewWallet(userId, passphrase string) (*Wallet, error){
	//Set up default ethereum network
	client, err := ethclient.Dial(string(Main))
	if err != nil{
		return nil, err
	}

	 newWallet := &Wallet{
		userId: userId,
		ks: keystore.NewKeyStore(keyStorePath + userId, keystore.StandardScryptN, keystore.StandardScryptP),
		client: client,
	 }

	 err = newWallet.importAccounts(passphrase)
	 if(err != nil){
	 	return nil, err
	 }
	 return newWallet, nil
}

func (w *Wallet) importAccounts(passphrase string) error{
	root := keyStorePath + w.userId

	err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if info.IsDir(){
			return nil
		}

		jsonBytes, err := ioutil.ReadFile(file)
		if err != nil{
			return err
		}

		w.ks.Import(jsonBytes, passphrase, passphrase)

		return nil
	})

	return err
}

func (w *Wallet) GenerateAccount(passphrase string) error {
	_, err := w.ks.NewAccount(passphrase)
	if err != nil {
		return  err
	}

	return nil
}

func (w *Wallet) ImportAccount(privKey, passphrase string) error {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return err
	}

	_, err = w.ks.ImportECDSA(privateKey, passphrase)
	if err != nil {
		return err
	}

	return nil
}

func (w *Wallet) GetAllAccounts() []accounts.Account{
	return w.ks.Accounts()
}

func (w *Wallet) DeleteAccount(accountAddress common.Address, passphrase string) error {
	for _, account := range(w.ks.Accounts()) {
		if(account.Address == accountAddress){
			w.ks.Delete(account, passphrase)
			return nil
		}
	}

	return errors.New("account not found.")
}

func (w *Wallet) ChangePassword(passphrase, newPassphrase string) error{
	for _, account := range(w.ks.Accounts()) {
		err := w.ks.Update(account, passphrase, newPassphrase)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *Wallet) GetBalance(accountAddress common.Address, passphrase string) (*big.Float, error){
	balanceAt, err := w.client.BalanceAt(context.Background(), accountAddress, nil)
	if err != nil {
		return nil, err
	}

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18 /* Ether Decimals */)))
	return ethValue, nil
}

func(w *Wallet) GetTokenBalance(token TokenAddress, accountAddress common.Address) (*big.Float, error){
	tokenAddress := common.HexToAddress(string(token))
	instance, err := contracts.NewToken(tokenAddress, w.client)
	if err != nil {
		return nil, err
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, accountAddress)
	if err != nil {
		return nil, err
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	return value, nil
}

func(w *Wallet) ChangeNetwork(newNet EthNet) error {
	client, err := ethclient.Dial(string(newNet))
	if err != nil {
		return err
	}
	w.client = client
	return nil
}

func(w *Wallet) TransferEther(from, to common.Address, value float64, passphrase string) (txHash string, err error) {
	for _, account := range(w.ks.Accounts()) {
		if(account.Address == from){


			root := keyStorePath + w.userId

			err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
				if info.IsDir(){
					return nil
				}

				if !strings.HasSuffix(file, strings.ToLower(from.String()[2:])){
					return nil
				}

				jsonBytes, err := ioutil.ReadFile(file)
				if err != nil{
					return err
				}

				Key, err := keystore.DecryptKey(jsonBytes, passphrase)
				if err != nil {
					return err
				}

				nonce, err := w.client.PendingNonceAt(context.Background(), from)
				if err != nil {
					return err
				}

				gasPrice, err := w.client.SuggestGasPrice(context.Background())
				if err != nil {
					return err
				}

				valueInWei, _ := new(big.Float).Mul(big.NewFloat(value) ,big.NewFloat(math.Pow10(18))).Int(nil)

				gasLimit := uint64(21000)

				var data []byte
				tx := types.NewTransaction(nonce, to, valueInWei, gasLimit, gasPrice, data)

				chainID, err := w.client.NetworkID(context.Background())
				if err != nil {
					return err
				}

				signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), Key.PrivateKey)
				if err != nil {
					return err
				}

				err = w.client.SendTransaction(context.Background(), signedTx)
				if err != nil {
					return err
				}

				txHash = signedTx.Hash().Hex()
				return nil
			})

			return txHash, err
		}
	}

	return "", errors.New("_from account not found.")

}

func(w *Wallet) TransferToken(token TokenAddress, from, to common.Address, value float64, passphrase string) (txHash string, err error) {
	for _, account := range(w.ks.Accounts()) {
		if(account.Address == from){


			root := keyStorePath + w.userId

			err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
				if info.IsDir(){
					return nil
				}

				if !strings.HasSuffix(file, strings.ToLower(from.String()[2:])){
					return nil
				}

				jsonBytes, err := ioutil.ReadFile(file)
				if err != nil{
					return err
				}

				Key, err := keystore.DecryptKey(jsonBytes, passphrase)
				if err != nil {
					return err
				}

				tokenAddress := common.HexToAddress(string(token))
				instance, err := contracts.NewToken(tokenAddress, w.client)
				if err != nil {
					return err
				}

				decimals, err := instance.Decimals(&bind.CallOpts{})
				if err != nil {
					return err
				}


				valueInWei, _ := new(big.Float).Mul(big.NewFloat(value) ,big.NewFloat(math.Pow10(int(decimals)))).Int(nil)

				txHash, err = transferToken(tokenAddress, Key.PrivateKey, valueInWei, to, w.client)
				return err
			})

			return txHash, err
		}
	}

	return "", errors.New("_from account not found.")

}

func(w *Wallet) ApproveToken(token TokenAddress, from, to common.Address, value float64, passphrase string) (txHash string, err error) {
	for _, account := range(w.ks.Accounts()) {
		if(account.Address == from){


			root := keyStorePath + w.userId

			err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
				if info.IsDir(){
					return nil
				}

				if !strings.HasSuffix(file, strings.ToLower(from.String()[2:])){
					return nil
				}

				jsonBytes, err := ioutil.ReadFile(file)
				if err != nil{
					return err
				}

				Key, err := keystore.DecryptKey(jsonBytes, passphrase)
				if err != nil {
					return err
				}

				tokenAddress := common.HexToAddress(string(token))
				instance, err := contracts.NewToken(tokenAddress, w.client)
				if err != nil {
					return err
				}

				decimals, err := instance.Decimals(&bind.CallOpts{})
				if err != nil {
					return err
				}


				valueInWei, _ := new(big.Float).Mul(big.NewFloat(value) ,big.NewFloat(math.Pow10(int(decimals)))).Int(nil)

				txHash, err = approveToken(tokenAddress, Key.PrivateKey, valueInWei, to, w.client)
				return err
			})

			return txHash, err
		}
	}

	return "", errors.New("_from account not found.")

}

func(w *Wallet) USDTContractDeposit(from, contractAdr common.Address, value float64, passphrase string) (txHash string, err error){
	const USDTDecimals = 6

	// Should be after approving confirmed
	for _, account := range(w.ks.Accounts()) {
		if(account.Address == from){


			root := keyStorePath + w.userId

			err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
				if info.IsDir(){
					return nil
				}

				if !strings.HasSuffix(file, strings.ToLower(from.String()[2:])){
					return nil
				}

				jsonBytes, err := ioutil.ReadFile(file)
				if err != nil{
					return err
				}

				Key, err := keystore.DecryptKey(jsonBytes, passphrase)
				if err != nil {
					return err
				}

				instance, err := contracts.NewContract(contractAdr, w.client)
				if err != nil {
					return err
				}

				nonce, err := w.client.PendingNonceAt(context.Background(), from)
				if err != nil {
					return err
				}

				gasPrice, err := w.client.SuggestGasPrice(context.Background())
				if err != nil {
					return err
				}

				auth := bind.NewKeyedTransactor(Key.PrivateKey)
				auth.Nonce = big.NewInt(int64(nonce))
				auth.Value = big.NewInt(0)     // in wei
				// TODO: gas estimate
				auth.GasLimit = uint64(150000) // 146656
				auth.GasPrice = gasPrice

				amount, _ := big.NewFloat(value * math.Pow10(USDTDecimals)).Int(nil)

				tx, err := instance.Deposit(auth, amount)
				txHash = tx.Hash().Hex()

				return err
			})

			return txHash, err
		}
	}

	return "", errors.New("_from account not found.")






}

func(w *Wallet) USDTContractWithdraw(from, contractAdr common.Address, valueToA, valueToB float64, feePercent int64, passphrase string) (txHash string, err error){
	const USDTDecimals = 6

	// Should be after approving confirmed
	for _, account := range(w.ks.Accounts()) {
		if(account.Address == from){


			root := keyStorePath + w.userId

			err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
				if info.IsDir(){
					return nil
				}

				if !strings.HasSuffix(file, strings.ToLower(from.String()[2:])){
					return nil
				}

				jsonBytes, err := ioutil.ReadFile(file)
				if err != nil{
					return err
				}

				Key, err := keystore.DecryptKey(jsonBytes, passphrase)
				if err != nil {
					return err
				}

				instance, err := contracts.NewContract(contractAdr, w.client)
				if err != nil {
					return err
				}

				nonce, err := w.client.PendingNonceAt(context.Background(), from)
				if err != nil {
					return err
				}

				gasPrice, err := w.client.SuggestGasPrice(context.Background())
				if err != nil {
					return err
				}

				auth := bind.NewKeyedTransactor(Key.PrivateKey)
				auth.Nonce = big.NewInt(int64(nonce))
				auth.Value = big.NewInt(0)     // in wei
				// TODO: gas estimate
				auth.GasLimit = uint64(150000) // 114618
				auth.GasPrice = gasPrice

				amountToA, _ := big.NewFloat(valueToA * math.Pow10(USDTDecimals)).Int(nil)
				amountToB, _ := big.NewFloat(valueToB * math.Pow10(USDTDecimals)).Int(nil)

				tx, err := instance.Withdraw(auth, amountToA, amountToB, big.NewInt(feePercent))
				txHash = tx.Hash().Hex()

				return err
			})

			return txHash, err
		}
	}

	return "", errors.New("_from account not found.")






}

func(w *Wallet) USDTContractCreate(from, addressA, addressB, addressJudge common.Address, content string, feePercentLimit int64,  passphrase string) (txHash string, err error){
	const USDTDecimals = 6

	// Should be after approving confirmed
	for _, account := range(w.ks.Accounts()) {
		if(account.Address == from){


			root := keyStorePath + w.userId

			err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
				if info.IsDir(){
					return nil
				}

				if !strings.HasSuffix(file, strings.ToLower(from.String()[2:])){
					return nil
				}

				jsonBytes, err := ioutil.ReadFile(file)
				if err != nil{
					return err
				}

				Key, err := keystore.DecryptKey(jsonBytes, passphrase)
				if err != nil {
					return err
				}

				ServerAdr := common.HexToAddress(RopstenServerAddress)
				instance, err := contracts.NewSever(ServerAdr, w.client)
				if err != nil {
					return err
				}

				nonce, err := w.client.PendingNonceAt(context.Background(), from)
				if err != nil {
					return err
				}

				gasPrice, err := w.client.SuggestGasPrice(context.Background())
				if err != nil {
					return err
				}

				auth := bind.NewKeyedTransactor(Key.PrivateKey)
				auth.Nonce = big.NewInt(int64(nonce))
				auth.Value = big.NewInt(0)     // in wei
				// TODO: gas estimate
				auth.GasLimit = uint64(2000000) // 1898437
				auth.GasPrice = gasPrice


				tx, err := instance.CreateContract(auth, addressA, addressB, addressJudge, content ,big.NewInt(feePercentLimit))
				txHash = tx.Hash().Hex()

				return err
			})

			return txHash, err
		}
	}

	return "", errors.New("_from account not found.")






}

func transferToken(tokenAdress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int, toAddress common.Address, client *ethclient.Client) (string, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	value := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAdress,
		Data: data,
	})
	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(nonce, tokenAdress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil

}

func approveToken(tokenAdress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int, toAddress common.Address, client *ethclient.Client) (string, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	value := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	transferFnSignature := []byte("approve(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAdress,
		Data: data,
	})
	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(nonce, tokenAdress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil

}