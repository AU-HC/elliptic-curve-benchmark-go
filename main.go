package main

import (
	"elliptic-curve-benchmark-go/ecc"
	"elliptic-curve-benchmark-go/elgamal"
	"fmt"
	"math/big"
	"time"
)

func main() {
	// Benchmark values
	amountOfSamples := 1000

	// ElGamal Elliptic Curves
	curve, _, publicKey := ecc.GenerateKeyPair()
	eccTimeStart := time.Now()
	for i := 0; i < 1000; i++ {
		ecc.Encrypt(curve, publicKey, big.NewInt(1))
	}
	eccTimeElapsed := time.Since(eccTimeStart)
	informationString := fmt.Sprintf("ElGamal Elliptic Curves with %d samples took: %s", amountOfSamples, eccTimeElapsed.String())
	fmt.Println(informationString)

	// Exponential ElGamal
	elgamalTimeStart := time.Now()
	_, elgamalPublicKey := elgamal.GenerateKeyPair()
	for i := 0; i < 1000; i++ {
		elgamal.Encrypt(elgamalPublicKey, big.NewInt(1))
	}
	elgamalTimeElapsed := time.Since(elgamalTimeStart)
	informationString = fmt.Sprintf("Exponential ElGamal with %d samples took: %s", amountOfSamples, elgamalTimeElapsed.String())
	fmt.Println(informationString)
}
