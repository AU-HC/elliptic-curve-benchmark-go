package main

import (
	"elliptic-curve-benchmark-go/elgamal"
	"fmt"
)

func main() {
	/*
		curve, _, publicKey := ecc.GenerateKeyPair()

		eccTimeStart := time.Now()
		for i := 0; i < 1000; i++ {
			ecc.Encrypt(curve, publicKey, big.NewInt(1321233121548349453), big.NewInt(1))
		}

		eccTimeElapsed := time.Since(eccTimeStart)
		fmt.Println(eccTimeElapsed.String())
	*/

	fmt.Println(elgamal.GenerateKey().K)

}
