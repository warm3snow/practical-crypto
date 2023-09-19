/**
 * @Author: xueyanghan
 * @File: argon2impl.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/19 14:18
 */

package pbkdf2impl

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/argon2"
	"strconv"
	"strings"
)

var (
	defaultTime   = 3
	defaultMemory = 32 * 1024
	defaultThread = 4
	defaultKeyLen = 32
)

type Argon2impl struct {
	Time    int
	Memory  int //Memory the memory parameter specifies the size of the memory in KiB
	Threads int
	KeyLen  int

	salt      []byte
	deriveKey []byte
}

func (a *Argon2impl) DeriveKeyByPassword(password string) (deriveKey []byte, err error) {
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return nil, errors.Wrap(err, "rand.Read failed")
	}
	a.salt = salt
	a.deriveKey = argon2.Key([]byte(password), a.salt, uint32(a.Time), uint32(a.Memory), uint8(a.Threads), uint32(a.KeyLen))
	return a.deriveKey, nil
}

func (a *Argon2impl) VerifyDeriveKeyStr(kdfKeyStr string, password []byte) (isOk bool, err error) {
	kdfKeyStr = strings.TrimLeft(kdfKeyStr, "$")
	kdfKeyStrs := strings.Split(kdfKeyStr, "$")
	if len(kdfKeyStrs) != 4 {
		return false, errors.New("kdfKeyStr format error, not 4 parts")
	}
	if kdfKeyStrs[0] != "argon2" {
		return false, errors.New("kdfKeyStr format error, not argon2")
	}
	salt, err := base64.StdEncoding.DecodeString(kdfKeyStrs[1])
	if err != nil {
		return false, errors.Wrap(err, "base64.StdEncoding.DecodeString salt failed")
	}
	deriveKey, err := base64.StdEncoding.DecodeString(kdfKeyStrs[2])
	if err != nil {
		return false, errors.Wrap(err, "base64.StdEncoding.DecodeString deriveKey failed")
	}

	params := strings.Split(kdfKeyStrs[3], ":")
	if len(params) != 4 {
		return false, errors.New("kdfKeyStr format error, not 4 params")
	}
	timeStr, memoryStr, threadsStr, keyLenStr := params[0], params[1], params[2], params[3]
	if timeStr == "" || memoryStr == "" || threadsStr == "" || keyLenStr == "" {
		return false, errors.New("kdfKeyStr format error, params is empty")
	}
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return false, errors.Wrap(err, "strconv.Atoi time failed")
	}
	memory, err := strconv.Atoi(memoryStr)
	if err != nil {
		return false, errors.Wrap(err, "strconv.Atoi memory failed")
	}
	threads, err := strconv.Atoi(threadsStr)
	if err != nil {
		return false, errors.Wrap(err, "strconv.Atoi threads failed")
	}
	keyLen, err := strconv.Atoi(keyLenStr)
	if err != nil {
		return false, errors.Wrap(err, "strconv.Atoi keyLen failed")
	}
	regenDeriveKey := argon2.Key(password, salt, uint32(time), uint32(memory), uint8(threads), uint32(keyLen))
	if bytes.Equal(deriveKey, regenDeriveKey) {
		return true, nil
	}
	return false, nil
}

func (a *Argon2impl) GetDeriveKeyStr() string {
	// format: $argon2$<params>$salt$key
	kdfKeyStrs := make([]string, 0)
	kdfKeyStrs = append(kdfKeyStrs, "$argon2")
	encodedSalt := base64.StdEncoding.EncodeToString(a.salt)
	kdfKeyStrs = append(kdfKeyStrs, fmt.Sprintf("%s", encodedSalt))
	encodedDK := base64.StdEncoding.EncodeToString(a.deriveKey)
	kdfKeyStrs = append(kdfKeyStrs, fmt.Sprintf("%s", encodedDK))
	kdfKeyStrs = append(kdfKeyStrs, fmt.Sprintf("%d:%d:%d:%d", a.Time, a.Memory, a.Threads, a.KeyLen))
	return strings.Join(kdfKeyStrs, "$")
}

func (a *Argon2impl) checkParams() {
	if a.Time <= 0 {
		a.Time = defaultTime
	}
	if a.Memory <= 0 {
		a.Memory = defaultMemory
	}
	if a.Threads <= 0 {
		a.Threads = defaultThread
	}
	if a.KeyLen <= 0 {
		a.KeyLen = defaultKeyLen
	}
}

// New - new a Argon2impl
func New(time, memory, threads, keyLen int) *Argon2impl {
	argon2Impl := &Argon2impl{
		Time:    time,
		Memory:  memory,
		Threads: threads,
		KeyLen:  keyLen,
	}
	argon2Impl.checkParams()

	return argon2Impl
}
