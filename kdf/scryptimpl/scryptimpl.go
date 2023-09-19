/**
 * @Author: xueyanghan
 * @File: pbkdf2impl.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/19 16:02
 */

package pbkdf2impl

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/scrypt"
	"strconv"
	"strings"
)

var (
	defaultN = 32768
	defaultR = 8
	defaultP = 1

	defaultKeyLen  = 32
	defaultSaltLen = 16
)

type ScryptImpl struct {
	n, r, p int
	keyLen  int

	salt      []byte
	deriveKey []byte
}

func (s *ScryptImpl) DeriveKeyByPassword(password string) (deriveKey []byte, err error) {
	salt := make([]byte, defaultSaltLen)
	_, err = rand.Read(salt)
	if err != nil {
		return nil, errors.Wrap(err, "rand.Read error")
	}
	s.salt = salt
	s.deriveKey, err = scrypt.Key([]byte(password), salt, s.n, s.r, s.p, s.keyLen)
	if err != nil {
		return nil, errors.Wrap(err, "scrypt.Key error")
	}
	return s.deriveKey, nil
}

func (s *ScryptImpl) VerifyDeriveKeyStr(kdfKeyStr string, password []byte) (isOk bool, err error) {
	kdfKeyStr = strings.TrimLeft(kdfKeyStr, "$")
	kdfKeyStrs := strings.Split(kdfKeyStr, "$")
	if len(kdfKeyStrs) != 4 {
		return false, errors.New("kdfKeyStr format error, not 4 parts")
	}
	if kdfKeyStrs[0] != s.KDFName() {
		return false, errors.New("kdfKeyStr format error, not scrypt")
	}
	salt, err := base64.StdEncoding.DecodeString(kdfKeyStrs[1])
	if err != nil {
		return false, errors.Wrap(err, "base64.StdEncoding.DecodeString salt error")
	}
	deriveKey, err := base64.StdEncoding.DecodeString(kdfKeyStrs[2])
	if err != nil {
		return false, errors.Wrap(err, "base64.StdEncoding.DecodeString deriveKey error")
	}
	params := strings.Split(kdfKeyStrs[3], ":")
	if len(params) != 4 {
		return false, errors.New("kdfKeyStr format error, not 4 parts")
	}
	n, err := strconv.Atoi(params[0])
	if err != nil {
		return false, errors.Wrap(err, "strconv.Atoi n error")
	}
	r, err := strconv.Atoi(params[1])
	if err != nil {
		return false, errors.Wrap(err, "strconv.Atoi r error")
	}
	p, err := strconv.Atoi(params[2])
	if err != nil {
		return false, errors.Wrap(err, "strconv.Atoi p error")
	}
	keyLen, err := strconv.Atoi(params[3])
	if err != nil {
		return false, errors.Wrap(err, "strconv.Atoi keyLen error")
	}
	deriveKey2, err := scrypt.Key(password, salt, n, r, p, keyLen)
	if err != nil {
		return false, errors.Wrap(err, "scrypt.Key error")
	}
	if !bytes.Equal(deriveKey, deriveKey2) {
		return false, nil
	}
	return true, nil
}

func (s *ScryptImpl) GetDeriveKeyStr() string {
	// kdf key format: $scrypt$$salt$key$n:r:p:keyLen
	kdfKeyStrs := make([]string, 0)
	kdfKeyStrs = append(kdfKeyStrs, s.KDFName())
	encodedSalt := base64.StdEncoding.EncodeToString(s.salt)
	kdfKeyStrs = append(kdfKeyStrs, encodedSalt)
	encodedDK := base64.StdEncoding.EncodeToString(s.deriveKey)
	kdfKeyStrs = append(kdfKeyStrs, encodedDK)
	kdfKeyStrs = append(kdfKeyStrs, fmt.Sprintf("%d:%d:%d:%d", s.n, s.r, s.p, s.keyLen))
	return "$" + strings.Join(kdfKeyStrs, "$")
}

func (s *ScryptImpl) checkParams() {
	if s.n <= 0 {
		s.n = defaultN
	}
	if s.r <= 0 {
		s.r = defaultR
	}
	if s.p <= 0 {
		s.p = defaultP
	}
	if s.keyLen <= 0 {
		s.keyLen = defaultKeyLen
	}
}

func (s *ScryptImpl) KDFName() string {
	return "scrypt"
}

func New(n, r, p, keyLen int) *ScryptImpl {
	impl := &ScryptImpl{
		n:      n,
		r:      r,
		p:      p,
		keyLen: keyLen,
	}
	impl.checkParams()
	return impl
}
