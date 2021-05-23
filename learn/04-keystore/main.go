package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "passwordd"
	// a, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(a.Address)

	b, err := ioutil.ReadFile("./wallet/UTC--2021-05-23T16-37-05.963879488Z--527a16d4d5a0bf229e6add125d32866783df7c3c")
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Priv", hexutil.Encode(pData))

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Pub", hexutil.Encode(pData))

	fmt.Println("Add", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())
}
