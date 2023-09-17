/**
 * @Author: xueyanghan
 * @File: main.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/17 12:01
 */

package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

func main() {
	salt, passwordHash, err := GeneratePowCrypt(1, []byte("123456"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("passwordHash[%d]: %v\n", len(passwordHash), passwordHash)
	fmt.Printf("passwordHash binary: %08b\n", passwordHash)
	isOk, err := VerifyPowCrypt(passwordHash, salt, []byte("123456"))
	if err != nil {
		return
	}
	if isOk {
		fmt.Println("Password was correct!")
	}
}

func GeneratePowCrypt(difficult int, password []byte) ([]byte, []byte, error) {
	var (
		random []byte
		salt   []byte
	)
	for {
		salt = make([]byte, 32)
		_, err := rand.Read(salt)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "generate salt failed")
		}
		passwordWithSalt := sha256.Sum256(append(password, salt...))
		random, err = Pow(difficult, passwordWithSalt[:])
		if err != nil {
			continue
		}
		break
	}
	return salt, random, nil
}

func VerifyPowCrypt(passwordHash []byte, salt []byte, password []byte) (bool, error) {
	newPasswordHash := sha256.Sum256(append(password, salt...))
	if !bytes.Equal(passwordHash, newPasswordHash[:]) {
		return false, nil
	}
	return true, nil
}

func Pow(different int, random []byte) ([]byte, error) {
	max32BigIntString := fmt.Sprintf("%s", strings.Repeat("f", 64))
	max32BigInt, _ := new(big.Int).SetString(max32BigIntString, 16)
	maxWithN0Prefix := new(big.Int).Rsh(max32BigInt, uint(different))

	fmt.Printf("maxWithN0Prefix: %08b\n", maxWithN0Prefix.Bytes())

	randomHash := new(big.Int).SetBytes(random)
	if randomHash.Cmp(maxWithN0Prefix) == -1 || randomHash.Cmp(maxWithN0Prefix) == 0 {
		return random, nil
	}
	return nil, errors.New("random hash is too big")
}
