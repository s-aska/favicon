package secure

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"strconv"
)

// GenerateRandomBytes ...
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomString ...
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

// GenerateRandomNumber ... 擬似乱数でstring生成
func GenerateRandomNumber(max int64) (string, error) {
	v, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(v.Uint64(), 10), nil
}
