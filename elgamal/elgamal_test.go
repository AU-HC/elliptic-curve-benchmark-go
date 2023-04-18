package elgamal

import (
	"math/big"
	"testing"
)

func TestElGamalEncryptionOfZero(t *testing.T) {
	secretKey, publicKey := GenerateKeyPair()

	plaintext := big.NewInt(1)

	want := big.NewInt(0).Exp(publicKey.G, plaintext, publicKey.P)

	alpha, beta := Encrypt(publicKey, plaintext)
	decryptedCiphertext := Decrypt(secretKey, alpha, beta)

	if want.Cmp(decryptedCiphertext) != 0 {
		t.Fatalf("Error got %s, but wanted %s", want.String(), plaintext.String())
	}
}

func TestElGamalEncryptionOfOne(t *testing.T) {
	secretKey, publicKey := GenerateKeyPair()

	plaintext := big.NewInt(1)

	want := big.NewInt(0).Exp(publicKey.G, plaintext, publicKey.P)

	alpha, beta := Encrypt(publicKey, plaintext)
	decryptedCiphertext := Decrypt(secretKey, alpha, beta)

	if want.Cmp(decryptedCiphertext) != 0 {
		t.Fatalf("Error got %s, but wanted %s", want.String(), plaintext.String())
	}
}
