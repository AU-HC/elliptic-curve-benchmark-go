package main

import (
	"elliptic-curve-benchmark-go/ecc"
	"elliptic-curve-benchmark-go/elgamal"
	"elliptic-curve-benchmark-go/random"
	"fmt"
	"time"
)

func main() {
	// Benchmark values
	amountOfSamples := 1000

	// Encryption
	printLineSeparator()
	benchmarkEncryption(amountOfSamples)

	// Decryption
	printLineSeparator()
	benchmarkDecryption(amountOfSamples)
	printLineSeparator()
}

func benchmarkEncryption(amountOfSamples int) {
	// ElGamal Elliptic Curves
	curve, _, publicKey := ecc.GenerateKeyPair()
	eccTimeStart := time.Now()
	for i := 0; i < amountOfSamples; i++ {
		ecc.Encrypt(curve, publicKey, random.GetZeroOrOne())
	}
	eccTimeElapsed := time.Since(eccTimeStart)
	informationString := fmt.Sprintf("ElGamal Elliptic Curves encryption with %d samples took: %dms", amountOfSamples, eccTimeElapsed.Milliseconds())
	fmt.Println(informationString)

	// Exponential ElGamal
	elgamalTimeStart := time.Now()
	_, elgamalPublicKey := elgamal.GenerateKeyPair()
	for i := 0; i < amountOfSamples; i++ {
		elgamal.Encrypt(elgamalPublicKey, random.GetZeroOrOne())
	}
	elgamalTimeElapsed := time.Since(elgamalTimeStart)
	informationString = fmt.Sprintf("Exponential ElGamal encryption with %d samples took: %dms", amountOfSamples, elgamalTimeElapsed.Milliseconds())
	fmt.Println(informationString)
}

func benchmarkDecryption(amountOfSamples int) {
	// Setting up encryptions
	eccCiphertexts := make([]ecc.Ciphertext, amountOfSamples)
	elgamalCiphertexts := make([]elgamal.Ciphertext, amountOfSamples)

	curve, eccSecretKey, eccPublicKey := ecc.GenerateKeyPair()
	for i := 0; i < amountOfSamples; i++ {
		c1, c2 := ecc.Encrypt(curve, eccPublicKey, random.GetZeroOrOne())

		eccCiphertexts[i] = ecc.Ciphertext{C1: c1, C2: c2}
	}

	elgamalSecretKey, elgamalPublicKey := elgamal.GenerateKeyPair()
	for i := 0; i < amountOfSamples; i++ {
		alpha, beta := elgamal.Encrypt(elgamalPublicKey, random.GetZeroOrOne())
		elgamalCiphertexts[i] = elgamal.Ciphertext{Alpha: alpha, Beta: beta}
	}

	// ElGamal Elliptic Curves
	eccTimeStart := time.Now()
	for i := 0; i < amountOfSamples; i++ {
		cipherText := eccCiphertexts[i]
		ecc.Decrypt(curve, eccSecretKey, cipherText.C1, cipherText.C2)
	}
	eccTimeElapsed := time.Since(eccTimeStart)
	informationString := fmt.Sprintf("ElGamal Elliptic Curves decryption with %d samples took: %dms", amountOfSamples, eccTimeElapsed.Milliseconds())
	fmt.Println(informationString)

	// Exponential ElGamal
	elgamalTimeStart := time.Now()
	for i := 0; i < amountOfSamples; i++ {
		cipherText := elgamalCiphertexts[i]
		elgamal.Decrypt(elgamalSecretKey, cipherText.Alpha, cipherText.Beta)
	}
	elgamalTimeElapsed := time.Since(elgamalTimeStart)
	informationString = fmt.Sprintf("Exponential ElGamal encryption with %d samples took: %dms", amountOfSamples, elgamalTimeElapsed.Milliseconds())
	fmt.Println(informationString)
}

func printLineSeparator() {
	fmt.Println("----------------------------------------------------------------------------")
}
