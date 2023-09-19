/**
 * @Author: xueyanghan
 * @File: pbkdf_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/17 16:48
 */

package pbkdf2impl

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
	"testing"
)

var (
	password = []byte("123456")
)

func TestPbkdf2(t *testing.T) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	iters := 1000
	keyLen := 32

	deriveKey := pbkdf2.Key(password, salt, iters, keyLen, sha256.New)
	fmt.Printf("deriveKey hex[%d] = %s\n", len(deriveKey), hex.EncodeToString(deriveKey))

	deriveKey = pbkdf2.Key(password, salt, iters, keyLen, sha256.New)
	fmt.Printf("deriveKey hex[%d] = %s\n", len(deriveKey), hex.EncodeToString(deriveKey))
}

func TestBcrypt(t *testing.T) {
	// not key derivation function
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("hashedPassword[%d]: %v\n", len(hashedPassword), hashedPassword)
	fmt.Printf("hashedPassword str: %s\n", string(hashedPassword))

	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return
	}
	fmt.Println("Password was correct!")
}

func TestScrypt(t *testing.T) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}

	N := 16384
	r := 8
	p := 1
	keyLen := 32

	deriveKey, err := scrypt.Key(password, salt, N, r, p, keyLen)
	if err != nil {
		panic(err)
	}
	fmt.Printf("deriveKey hex[%d] = %s\n", len(deriveKey), hex.EncodeToString(deriveKey))

	deriveKey, err = scrypt.Key(password, salt, N, r, p, keyLen)
	if err != nil {
		panic(err)
	}
	fmt.Printf("deriveKey hex[%d] = %s\n", len(deriveKey), hex.EncodeToString(deriveKey))
}
