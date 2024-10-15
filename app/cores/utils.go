package cores

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func EncodeHexToString(data []byte) string {
	return hex.EncodeToString(data)
}

func DecodeHexToBytes(data string) ([]byte, error) {
	return hex.DecodeString(data)
}

func GetHashSha256(data []byte) []byte {
	hash := sha256.New()
	buff := hash.Sum(data)
	return buff
}

type StrOrBuffImpl interface {
	~[]byte | ~string
}

func BytesEquals[V StrOrBuffImpl](data V, buff V) bool {
	size := len(data)
	if size != len(buff) {
		return false
	}
	for i := 0; i < size; i++ {
		if data[i] != buff[i] {
			return false
		}
	}
	return true
}

func StringEquals(data string, buff string) bool {
	return BytesEquals(data, buff)
}

func HashSha256Compare(data []byte, hash []byte) bool {
	temp := GetHashSha256(data)
	return BytesEquals(temp, hash)
}

func GetHashSha256ToString(data []byte) string {
	temp := GetHashSha256(data)
	return EncodeHexToString(temp)
}

func HashSha256StringCompare(data []byte, hash string) bool {
	temp := GetHashSha256ToString(data)
	return StringEquals(temp, hash)
}

func GetHashSha512(data []byte) []byte {
	hash := sha512.New()
	buff := hash.Sum(data)
	return buff
}

func HashSha512Compare(data []byte, hash []byte) bool {
	temp := GetHashSha512(data)
	return BytesEquals(temp, hash)
}

func GetHashSha512ToString(data []byte) string {
	temp := GetHashSha512(data)
	return EncodeHexToString(temp)
}

func HashSha512StringCompare(data []byte, hash string) bool {
	temp := GetHashSha512ToString(data)
	return StringEquals(temp, hash)
}
