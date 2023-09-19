package softimpl

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"github.com/warm3snow/gmsm/sm2"
	"github.com/warm3snow/gmsm/sm3"
	"github.com/warm3snow/gmsm/sm4"
	"github.com/warm3snow/gmsm/x509"
	"hash"
	"strings"
)

type softimpl struct {
}

func (s softimpl) GenSymKey(algo string, keySize int) ([]byte, error) {
	var key []byte
	switch strings.ToUpper(algo) {
	case "SM4":
		keySize = 16
		key = make([]byte, keySize)
	case "AES":
		if keySize != 16 && keySize != 24 && keySize != 32 {
			return nil, fmt.Errorf("wrong aes keySize, want 16/24/32, got %d", keySize)
		}
		key = make([]byte, keySize)
	}
	if _, err := rand.Read(key[:]); err != nil {
		return nil, err
	}
	return key, nil
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
		switch mode {
		case "C1C2C3":
			ciphertext, err = sm2.Encrypt(pub, plainText, rand.Reader, sm2.C1C2C3)
		case "C1C3C2":
			ciphertext, err = sm2.Encrypt(pub, plainText, rand.Reader, sm2.C1C3C2)
		default:
			ciphertext, err = sm2.EncryptAsn1(pub, plainText, rand.Reader)
		}
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
		switch mode {
		case "C1C2C3":
			plainText, err = sm2.Decrypt(priv, cipherText, sm2.C1C2C3)
		case "C1C3C2":
			plainText, err = sm2.Decrypt(priv, cipherText, sm2.C1C3C2)
		default:
			plainText, err = sm2.DecryptAsn1(priv, cipherText)
		}
		if err != nil {
			return nil, err
		}
	}
	return plainText, nil
}

func (s softimpl) Sign(algo, key string, plain []byte) ([]byte, error) {
	priv, err := ParsePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}
	switch privkey := priv.(type) {
	case *sm2.PrivateKey:
		return privkey.Sign(rand.Reader, plain, nil)
	case *ecdsa.PrivateKey:
		return ecdsa.SignASN1(rand.Reader, privkey, plain)
	case *rsa.PrivateKey:
		return rsa.SignPSS(rand.Reader, privkey, crypto.SHA256, plain, nil)
	}
	return nil, fmt.Errorf("unsupported private key type")
}

func (s softimpl) Verify(algo, key string, plain, sig []byte) (bool, error) {
	pub, err := ParsePublicKey([]byte(key))
	if err != nil {
		return false, err
	}
	switch pubkey := pub.(type) {
	case *sm2.PublicKey:
		return pubkey.Verify(plain, sig), nil
	case *ecdsa.PublicKey:
		return ecdsa.VerifyASN1(pubkey, plain, sig), nil
	case *rsa.PublicKey:
		err = rsa.VerifyPSS(pub.(*rsa.PublicKey), crypto.SHA256, plain, sig, nil)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, fmt.Errorf("unsupported public key type")
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

func New() (*softimpl, error) {
	return &softimpl{}, nil
}
