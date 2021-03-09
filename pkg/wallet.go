package wallet

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"os"
	"path/filepath"
)

const keyStorePath = "./keystores/"

type Wallet struct {
	userId string
	accounts []accounts.Account
	ks *keystore.KeyStore
}

func NewWallet(userId, passphrase string) *Wallet{
	 newWallet := &Wallet{
		userId: userId,
		ks: keystore.NewKeyStore(keyStorePath + userId, keystore.StandardScryptN, keystore.StandardScryptP),
	 }
	 newWallet.importAccounts(passphrase)
	 return newWallet
}

func (w *Wallet) importAccounts(passphrase string) error{
	root := keyStorePath + w.userId

	err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
		jsonBytes, err := ioutil.ReadFile(file)
		if err != nil{
			return err
		}

		account, _ := w.ks.Import(jsonBytes, passphrase, passphrase)
		if err != nil{
			return err
		}

		w.accounts = append(w.accounts, account)
		return nil
	})

	if err != nil{
		return err
	}

	return nil
}

func (w *Wallet) generateAccount(passphrase string) error {
	account, err := w.ks.NewAccount(passphrase)
	if err != nil {
		return  err
	}

	w.accounts = append(w.accounts, account)
	return nil
}

func (w *Wallet) importAccount(privKey, passphrase string) error {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return err
	}

	newAccount, err := w.ks.ImportECDSA(privateKey, passphrase)
	if err != nil {
		return err
	}

	w.accounts = append(w.accounts, newAccount)
	return nil
}


