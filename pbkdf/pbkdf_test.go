/**
 * @Author: xueyanghan
 * @File: pbkdf_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/17 16:48
 */

package pbkdf

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
	"testing"
)

var (
	password = []byte("123456")
)

func TestArgon2(t *testing.T) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	key := argon2.Key(password, salt, 3, 32*1024, 4, 32)
	key2 := argon2.IDKey(password, salt, 4, 32*1024, 4, 32)

	fmt.Printf("key[%d]: %v\n key2[%d]: %v\n", len(key), key, len(key2), key2)
}

func TestPbkdf2(t *testing.T) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	iters := 1000
	keyLen := 32

	deriveKey := pbkdf2.Key(password, salt, iters, keyLen, sha256.New)
	fmt.Println(deriveKey)

	deriveKey2 := pbkdf2.Key(password, salt, iters, keyLen, sha256.New)
	fmt.Println(deriveKey2)
}

func TestBcrypt(t *testing.T) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("hashedPassword[%d]: %v\n", len(hashedPassword), hashedPassword)

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
	fmt.Println(deriveKey)
}
