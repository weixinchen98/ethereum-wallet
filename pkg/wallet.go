package wallet

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"os"
	"path/filepath"
)

const keyStorePath = "../keystores/"

type Wallet struct {
	userId string
	ks *keystore.KeyStore
}

func NewWallet(userId, passphrase string) (*Wallet, error){
	 newWallet := &Wallet{
		userId: userId,
		ks: keystore.NewKeyStore(keyStorePath + userId, keystore.StandardScryptN, keystore.StandardScryptP),
	 }
	 err := newWallet.importAccounts(passphrase)
	 if(err != nil){
	 	return nil, err
	 }
	 return newWallet, nil
}

func (w *Wallet) importAccounts(passphrase string) error{
	root := keyStorePath + w.userId

	err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
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



