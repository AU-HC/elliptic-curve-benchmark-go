package random

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomModQ() *big.Int {
	q, _ := big.NewInt(0).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF43", 16)
	number, err := rand.Int(rand.Reader, q)

	if err != nil {
		panic(err)
	}

	return number
}

func GetZeroOrOne() *big.Int {
	number, err := rand.Int(rand.Reader, big.NewInt(2))

	if err != nil {
		panic(err)
	}

	return number
}
