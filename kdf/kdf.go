/**
 * @Author: xueyanghan
 * @File: kdf.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/19 14:15
 */

package kdf

import (
	"github.com/pkg/errors"
	"github.com/warm3snow/practical-crypto/kdf/argon2impl"
	"github.com/warm3snow/practical-crypto/kdf/bcryptimpl"
	"github.com/warm3snow/practical-crypto/kdf/pbkdf2impl"
	"github.com/warm3snow/practical-crypto/kdf/scryptimpl"
	"math/rand"
	"strings"
	"time"
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
func GenerateRandomPassword(length int, minClasses int) (string, error) {
	// 定义密码字符集
	digits := "0123456789"
	upperLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerLetters := "abcdefghijklmnopqrstuvwxyz"
	specialChars := "!@#$%^&*()_+[]{}|;:,.<>?`~"

	charSets := []string{digits, upperLetters, lowerLetters, specialChars}
	selectedSets := make([]byte, 0, minClasses)

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 从每个字符集中随机选择一个字符
	charSetsUsedFlag := make(map[int]bool)
	for i := 0; i < minClasses; i++ {
		charSetsIndex := rand.Intn(len(charSets))
		if charSetsUsedFlag[charSetsIndex] {
			i--
			continue
		} else {
			charSetsUsedFlag[charSetsIndex] = true
		}
		charSet := charSets[charSetsIndex]
		selectedSets = append(selectedSets, charSet[rand.Intn(len(charSet))])
	}

	// 剩余字符随机选择
	for i := 0; i < length-minClasses; i++ {
		charSet := charSets[rand.Intn(len(charSets))]
		selectedSets = append(selectedSets, charSet[rand.Intn(len(charSet))])
	}

	return string(selectedSets), nil
}
