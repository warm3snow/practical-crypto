/**
 * @Author: xueyanghan
 * @File: kdf.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/19 14:15
 */

package kdf

import (
	"crypto/rand"
	"github.com/pkg/errors"
	"github.com/warm3snow/practical-crypto/kdf/argon2impl"
	"github.com/warm3snow/practical-crypto/kdf/bcryptimpl"
	"github.com/warm3snow/practical-crypto/kdf/pbkdf2impl"
	"github.com/warm3snow/practical-crypto/kdf/scryptimpl"
	"math/big"
	"strings"
)

const (
	BCRYPT = "bcrypt"
	PBKDF2 = "pbkdf2"
	SCRYPT = "scrypt"
	ARGON2 = "argon2"
)

// KDF interface
type KDF interface {
	// DeriveKeyByPassword derive key by password
	DeriveKeyByPassword(password string) (deriveKey []byte, err error)

	// VerifyDeriveKeyStr verify deriveKeyStr
	VerifyDeriveKeyStr(kdfKeyStr string, password []byte) (isOk bool, err error)

	// GetDeriveKeyStr get deriveKeyStr
	GetDeriveKeyStr() string

	// KDFName get kdf name
	KDFName() string
}

// InitKdf init kdf by kdfName. Note: param '-1' means use default value. TODO: refactor this
func InitKdf(kdfName string, keyLen int) (KDF, error) {
	switch strings.ToLower(kdfName) {
	case BCRYPT:
		return bcryptimpl.New(-1), nil
	case PBKDF2:
		return pbkdf2impl.New(-1, keyLen), nil
	case SCRYPT:
		return scryptimpl.New(-1, -1, -1, keyLen), nil
	case ARGON2:
		return argon2impl.New(-1, -1, -1, keyLen), nil
	default:
		return nil, errors.Errorf("kdfName %s not supported", kdfName)
	}
}

// GenerateRandomPassword generate random password
func GenerateRandomPassword(length int) (string, error) {
	// 定义密码字符集
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" +
		"0123456789!@#$%^&*()_+[]{}|;:,.<>?`~"

	password := make([]byte, length)
	for i := range password {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
		password[i] = characters[index.Int64()]
	}
	return string(password), nil
}
