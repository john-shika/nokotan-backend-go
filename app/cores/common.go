package cores

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
)

func NewRandomBytes(size int) ([]byte, error) {
	var n int
	var err error
	KeepVoid(n, err)
	buff := make([]byte, 32)
	if n, err = rand.Read(buff); err != nil {
		return nil, err
	}
	return buff, nil
}

func NewBase64EncodeToString(key []byte) string {
	return base64.StdEncoding.EncodeToString(key)
}

func NewBase64DecodeToBytes(key string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(key)
}

func NewUuid() string {
	var err error
	var temp uuid.UUID
	KeepVoid(err, temp)
	if temp, err = uuid.NewV7(); err != nil {
		panic(err)
	}
	return temp.String()
}
