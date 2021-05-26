package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url  = "https://kovan.infura.io/v3/0c7b3f204f37416388610fb274b0452c"
	murl = "https://mainnet.infura.io/v3/0c7b3f204f37416388610fb274b0452c"
)

func main() {
	// ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// _, err := ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// "1f7ecea2fa83cc4a7de969f11d16a40edf9023d7"
	// "1e41ca1ccfc06597525c966a986b35a09e22358d"

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	a1 := common.HexToAddress("c393967d7b4b7fd02e697d13085d645c9412af11")
	a2 := common.HexToAddress("1e41ca1ccfc06597525c966a986b35a09e22358d")

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance 1:", b1)
	fmt.Println("Balance 2:", b2)
	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal(err)
	}
	// 1 ether = 1000000000000000000 wei
	amount := big.NewInt(100000000000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, a2, amount, 21000, gasPrice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("wallet/UTC--2021-05-24T16-47-26.459903259Z--c393967d7b4b7fd02e697d13085d645c9412af11")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
