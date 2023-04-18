package ecc

import (
	"crypto/elliptic"
	"crypto/rand"
	"elliptic-curve-benchmark-go/random"
	"math/big"
)

func GenerateKeyPair() (elliptic.Curve, []byte, Point) {
	curve := elliptic.P256()
	privateKey, x, y, err := elliptic.GenerateKey(curve, rand.Reader)

	if err != nil {
		panic(err)
	}

	return curve, privateKey, Point{x: x, y: y}
}

func Encrypt(curve elliptic.Curve, publicKey Point, m *big.Int) (Point, Point) {
	// Generate nonce
	nonce := random.GenerateRandomModQ()

	// Encode the message as a point
	encodedMessage := EncodeMessageAsPoint(curve, m)

	// Calculate M1
	xM1, yM1 := curve.ScalarBaseMult(nonce.Bytes())

	// Calculate M2
	xKB, yKB := curve.ScalarMult(publicKey.x, publicKey.y, nonce.Bytes())
	xM2, yM2 := curve.Add(encodedMessage.x, encodedMessage.y, xKB, yKB)

	// Return the points
	return Point{x: xM1, y: yM1}, Point{x: xM2, y: yM2}
}

func Decrypt(curve elliptic.Curve, secretKey []byte, m1, m2 Point) *Point {
	xSM1, ySM1 := curve.ScalarMult(m1.x, m1.y, secretKey) // calculating sM1
	ySM1.Neg(ySM1)                                        // finding inverse of y value
	ySM1.Mod(ySM1, curve.Params().P)

	x, y := curve.Add(xSM1, ySM1, m2.x, m2.y)
	return &Point{x: x, y: y}
}

func EncodeMessageAsPoint(curve elliptic.Curve, m *big.Int) Point {
	x, y := curve.ScalarBaseMult(m.Bytes())
	return Point{x: x, y: y}
}

type Point struct {
	x, y *big.Int
}

func (p *Point) Compare(point *Point) bool {
	if p.x.Cmp(point.x) != 0 {
		return false
	}

	if p.y.Cmp(point.y) != 0 {
		return false
	}

	return true
}

func (p *Point) String() string {
	return p.x.String() + ", " + p.y.String()
}
