package cores

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func HexEncodeToString(data []byte) string {
	return hex.EncodeToString(data)
}

func HexDecodeToBytes(data string) ([]byte, error) {
	return hex.DecodeString(data)
}

func HashSha256(data []byte) []byte {
	hash := sha256.New()
	buff := hash.Sum(data)
	return buff
}

func HashSha256Compare(data []byte, hash []byte) bool {
	temp := HashSha256(data)
	return BytesEquals(temp, hash)
}

func HashSha384(data []byte) []byte {
	hash := sha512.New384()
	buff := hash.Sum(data)
	return buff
}

func HashSha384Compare(data []byte, hash []byte) bool {
	temp := HashSha384(data)
	return BytesEquals(temp, hash)
}

func HashSha512(data []byte) []byte {
	hash := sha512.New()
	buff := hash.Sum(data)
	return buff
}

func HashSha512Compare(data []byte, hash []byte) bool {
	temp := HashSha512(data)
	return BytesEquals(temp, hash)
}

type BytesOrStringImpl interface {
	string | []byte
}

func BytesEquals[V BytesOrStringImpl](data V, buff V) bool {
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

func ViperJwtConfigUnmarshal() (*JwtConfig, error) {
	var err error
	KeepVoid(err)

	jwtConfig := NewJwtConfig()
	keyName := ToSnakeCase(GetNameType(jwtConfig))
	if err = viper.UnmarshalKey(keyName, jwtConfig); err != nil {
		return nil, err
	}

	return jwtConfig, nil
}

func EnsureDirAndFile(filePath string) error {
	var err error
	var fileInfo os.FileInfo
	var file *os.File
	KeepVoid(err, fileInfo, file)

	pathDir := filepath.Dir(filePath)
	pathFile := filepath.Base(filePath)

	// Check if the directory exists, and create it if it doesn't
	if fileInfo, err = os.Stat(pathDir); os.IsNotExist(err) {
		if err = os.MkdirAll(pathDir, os.ModePerm); err != nil {
			return NewThrow(fmt.Sprintf("failed to create directory: %s", pathDir), err)
		}
		fmt.Printf("Directory %s created\n", pathDir)
	} else {
		fmt.Printf("Directory %s already exists\n", pathDir)
	}

	// Check if the file exists, and create it if it doesn't
	if fileInfo, err = os.Stat(filePath); os.IsNotExist(err) {
		if file, err = os.Create(filePath); err != nil {
			return NewThrow(fmt.Sprintf("failed to create file: %s", pathFile), err)
		}
		NoErr(file.Close())
		fmt.Printf("File %s created\n", pathFile)
	} else {
		fmt.Printf("File %s already exists\n", pathFile)
	}

	return nil
}

func GetNameType(obj any) string {
	var ok bool
	var nameable NameTyped
	KeepVoid(ok, nameable)

	if IsNone(obj) {
		return "<nil>"
	}

	// try cast nameable and call method
	if nameable, ok = obj.(NameTyped); ok {
		return nameable.GetNameType()
	}

	// print type name by sprintf
	return fmt.Sprintf("%T", obj)
}
