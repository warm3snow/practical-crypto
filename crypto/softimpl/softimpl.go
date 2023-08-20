package softimpl

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"github.com/warm3snow/gmsm/sm2"
	"github.com/warm3snow/gmsm/sm3"
	"github.com/warm3snow/gmsm/sm4"
	"github.com/warm3snow/gmsm/x509"
	"hash"
	"strings"
)

type softimpl struct {
}

func New() (*softimpl, error) {
	return &softimpl{}, nil
}

func (s softimpl) Hash(algo string, origin []byte) ([]byte, error) {
	switch strings.ToUpper(algo) {
	case "SM3":
		return sm3.Sm3Sum(origin), nil
	default:
		v := sha256.Sum256(origin)
		return v[:], nil
	}
}

func (s softimpl) Enc(algo string, key, plainText []byte, mode string) ([]byte, error) {
	var (
		ciphertext []byte
		err        error
	)
	switch strings.ToUpper(algo) {
	case "SM4":
		ciphertext, err = sm4.Sm4Cbc(key, plainText, true)
		if err != nil {
			return nil, err
		}
	case "SM2":
		pub, err := x509.ParseSm2PublicKey(key)
		if err != nil {
			return nil, err
		}
		ciphertext, err = sm2.Encrypt(pub, plainText, rand.Reader, sm2.C1C2C3)
		if err != nil {
			return nil, err
		}
	}
	return ciphertext, nil
}

func (s softimpl) Dec(algo string, key, cipherText []byte, mode string) ([]byte, error) {
	var (
		plainText []byte
		err       error
	)
	switch strings.ToUpper(algo) {
	case "SM4":
		plainText, err = sm4.Sm4Cbc(key, cipherText, false)
		if err != nil {
			return nil, err
		}
	case "SM2":
		priv, err := x509.ParsePKCS8PrivateKey(key, nil)
		if err != nil {
			return nil, err
		}
		plainText, err = sm2.Decrypt(priv, cipherText, sm2.C1C2C3)
		if err != nil {
			return nil, err
		}
	}
	return plainText, nil
}

func (s softimpl) Sign(algo, key string, plain []byte) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (s softimpl) Verify(algo, key string, plain, sig []byte) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s softimpl) HMac(algo, key string, plain []byte) ([]byte, error) {
	var h hash.Hash
	switch strings.ToUpper(algo) {
	case "SM3":
		h = hmac.New(sm3.New, []byte(key))
	default:
		h = hmac.New(sha256.New, []byte(key))
	}
	return h.Sum(plain), nil
}
