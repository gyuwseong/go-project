package coupon

import (
	"crypto/rand"
	"math/big"
)

const (
	maxCodeLength = 10
	maxRetries    = 3
	charset       = "가나다라마바사아자차카타파하0123456789"
)

type CouponGenerator struct{}

func NewCouponGenerator() *CouponGenerator {
	return &CouponGenerator{}
}

func (g *CouponGenerator) GenerateCode() string {
	code := make([]byte, maxCodeLength)
	for i := 0; i < maxCodeLength; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		code[i] = charset[n.Int64()]
	}
	return string(code)
}
