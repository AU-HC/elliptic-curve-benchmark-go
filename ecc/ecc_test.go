package ecc

import (
	"math/big"
	"testing"
)

func TestECCEncryptionOfZero(t *testing.T) {
	curve, secretKey, publicKey := GenerateKeyPair()

	plaintext := big.NewInt(0)

	want := EncodeMessageAsPoint(curve, plaintext)

	ciphertext1, ciphertext2 := Encrypt(curve, publicKey, plaintext)
	decryptedCiphertext := Decrypt(curve, secretKey, ciphertext1, ciphertext2)

	if !want.Compare(decryptedCiphertext) {
		t.Fatalf("Error got %s, but wanted %s", want.String(), plaintext.String())
	}
}

func TestECCEncryptionOfOne(t *testing.T) {
	curve, secretKey, publicKey := GenerateKeyPair()

	plaintext := big.NewInt(1)

	want := EncodeMessageAsPoint(curve, plaintext)

	ciphertext1, ciphertext2 := Encrypt(curve, publicKey, plaintext)
	decryptedCiphertext := Decrypt(curve, secretKey, ciphertext1, ciphertext2)

	if !want.Compare(decryptedCiphertext) {
		t.Fatalf("Error got %s, but wanted %s", want.String(), plaintext.String())
	}
}
