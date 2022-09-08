package random

import (
	"crypto/rand"
	"math/big"
)

func Number() (int64, error) {
	rand, err := rand.Int(rand.Reader, big.NewInt(9999))
	if err != nil {
		return 0, err
	}

	return rand.Int64(), nil
}
