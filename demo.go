package token

import (
	"context"
	"crypto/ecdsa"
	"ethereum-wallet/token"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"regexp"

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
	instance, err := token.NewToken(tokenAddress, client)
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
	//d, _ := new(big.Int).SetString("88dd1d995dff4568cefd78840e851784cd505578b90b5c10cb15e3ef0f107294", 16)
	//privateKey, err := crypto.ToECDSA(d.Bytes())

	privateKey, err := crypto.HexToECDSA("88dd1d995dff4568cefd78840e851784cd505578b90b5c10cb15e3ef0f107294")
	if err != nil {
		log.Fatal(err)
	}

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
	//myAddress := "0xE590AeBF062435C3fA7026AB6CC41CE7f9324db9"
	//fmt.Println(checkAddressValid(myAddress))

	//contractAddress := "0x0881DC670828Dc74Dc9AdE68ec328a579Dc1E660"
	//fmt.Println(isContract(contractAddress, client))
	//fmt.Println(isContract(myAddress, client))

	//Transferring ETH
	//signedTxHash := transferEth(privateKey, big.NewInt(100000000000000000) /* 0.1 Eth */, "0xE590AeBF062435C3fA7026AB6CC41CE7f9324db9", client)
	//fmt.Println(signedTxHash)

	//Transferring Token
	//amount := big.NewInt(1000000)// 1 USDT
	//signedTxHash := transferUSDT(privateKey, amount, "0xB1b766702D858681239787C7aDE9eFA7001938b2", client)
	//fmt.Println(signedTxHash)


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

func checkAddressValid(address string) (bool) {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func isContract(address string, client *ethclient.Client) (bool){
	ethAddress := common.HexToAddress(address)
	bytecode, err := client.CodeAt(context.Background(), ethAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	return len(bytecode) > 0
}

func transferEth(privateKey *ecdsa.PrivateKey, value *big.Int, toAddress string, client *ethclient.Client) string {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	to := common.HexToAddress(toAddress)
	var data []byte
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	return signedTx.Hash().Hex()
}

func transferUSDT(privateKey *ecdsa.PrivateKey, amount *big.Int, toAddress string, client *ethclient.Client) string {
	TehterTokenAddress := common.HexToAddress("0x0881DC670828Dc74Dc9AdE68ec328a579Dc1E660")
	to := common.HexToAddress(toAddress)


	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(to.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &TehterTokenAddress,
		Data: data,
	})
	if err != nil {
		log.Println("gas limit estimate error")
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	tx := types.NewTransaction(nonce, TehterTokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	return signedTx.Hash().Hex()

}