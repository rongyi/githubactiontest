package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ontology-layer-2/go-ethereum/crypto"
)

func main() {
	fmt.Println("hello    world")
	fmt.Println("hello no indent")
	fmt.Println("hello no indent")
	fmt.Println("hello no indent")
	fmt.Println("hello no indent")
	fmt.Println("hello no indent")
	fmt.Println("hello no indent")

	pri, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	b := crypto.FromECDSA(pri)
	log.Printf("get privatekey: %s", hex.EncodeToString(b))
}
