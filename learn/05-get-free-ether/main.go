package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
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

	client, err := ethclient.Dial(murl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	a1 := common.HexToAddress("1f7ecea2fa83cc4a7de969f11d16a40edf9023d7")
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
}
