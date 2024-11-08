package cores

import (
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
	"strings"
)

const (
	SALT_SIZE      = 16
	N_ITERATIONS   = 10000
	KEY_LENGTH     = 32
	SEPARATOR_CHAR = "$"
	PREFIX_WORD    = "PDKF2_"
)

func HashPassword(password string) (string, error) {
	var n int
	var err error
	KeepVoid(n, err)

	salt := make([]byte, SALT_SIZE)
	if n, err = rand.Read(salt); err != nil {
		return EmptyString, err
	}

	buff := []byte(password)
	key := pbkdf2.Key(buff, salt, N_ITERATIONS, KEY_LENGTH, sha256.New)
	temp := PREFIX_WORD + NewEncodeBase64URLSafe(salt) + SEPARATOR_CHAR + NewEncodeBase64URLSafe(key)
	return temp, nil
}

func CompareHashPassword(hash string, password string) bool {
	var n int
	var ok bool
	var err error
	KeepVoid(n, ok, err)

	if hash, ok = strings.CutPrefix(hash, PREFIX_WORD); !ok {
		return false
	}

	if !strings.Contains(hash, SEPARATOR_CHAR) {
		return false
	}

	tokens := strings.Split(hash, SEPARATOR_CHAR)
	if len(tokens) != 2 {
		return false
	}

	salt := Unwrap(NewDecodeBase64URLSafe(tokens[0]))
	pass := Unwrap(NewDecodeBase64URLSafe(tokens[1]))

	buff := []byte(password)
	key := pbkdf2.Key(buff, salt, 10000, 32, sha256.New)
	return BytesEquals(pass, key)
}
