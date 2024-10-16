package cores

import (
	crypto "crypto/rand"
	"math/big"
)

func CryptoRandomRange(min, max int) (int, error) {
	var n *big.Int
	var err error
	if n, err = crypto.Int(crypto.Reader, big.NewInt(int64(max-min))); err != nil {
		return 0, err
	}
	return min + int(n.Int64()), nil
}

func CryptoRandomRange32(min, max int32) (int32, error) {
	var n *big.Int
	var err error
	if n, err = crypto.Int(crypto.Reader, big.NewInt(int64(max-min))); err != nil {
		return 0, err
	}
	return min + int32(n.Int64()), nil
}

func CryptoRandomRange64(min, max int64) (int64, error) {
	var n *big.Int
	var err error
	if n, err = crypto.Int(crypto.Reader, big.NewInt(max-min)); err != nil {
		return 0, err
	}
	return min + n.Int64(), nil
}
