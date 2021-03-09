package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"log"
	"math"
	"math/big"
)

const ETHDecimals = 18
const USDTDecimals = 6

func main(){
	coinbase := "0xE590AeBF062435C3fA7026AB6CC41CE7f9324db9"
	TehterTokenAddress := "0x0881DC670828Dc74Dc9AdE68ec328a579Dc1E660"

	client, err := ethclient.Dial("https://ropsten.infura.io/v3/cbf6f482ffa6444c88c16c67aebbd738")
	if err != nil {
		log.Fatal(err)
	}

	//Account Balances
	account := common.HexToAddress(coinbase)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(ETHDecimals)))
	fmt.Printf("Eth value: %s\n", ethValue.String())

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Printf("Pending eth balance: %s\n", pendingBalance.String())


	// Account Token Balance
	tokenAddress := common.HexToAddress(TehterTokenAddress)
	instance, err := NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(coinbase)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("usdt wei: %s\n", bal.String())

	usdtFBalance := new(big.Float)
	usdtFBalance.SetString(bal.String())
	usdtValue := new(big.Float).Quo(usdtFBalance, big.NewFloat(math.Pow10(USDTDecimals)))
	fmt.Printf("Usdt value : %s\n", usdtValue.String())

	//Generating New Wallets

	// ToECDSA creates a private key with the given D value.
	d, _ := new(big.Int).SetString("88dd1d995dff4568cefd78840e851784cd505578b90b5c10cb15e3ef0f107294", 16)
	privateKey, err := crypto.ToECDSA(d.Bytes())

	//privateKey, err := crypto.GenerateKey()
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Print("Private key: ")
	fmt.Println(privateKey)

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	hexAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(hexAddress)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))

	//Keystores
	//file := "./wallets/UTC--2021-03-09T03-42-18.244284000Z--95acf4fac4eed3d1270264aadcd7c63a1da888b9"
	//importKs(file)

	//TODO: HD Wallet?

	//Address check
}

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs(file string) {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())

	key, err := keystore.DecryptKey(jsonBytes, password)
	fmt.Println(key.PrivateKey)
}