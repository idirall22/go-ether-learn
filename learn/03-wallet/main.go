package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(pData))

	puData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(puData))

	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
