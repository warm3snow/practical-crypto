package crypto

import (
	"github.com/warm3snow/practical-crypto/crypto/hsmimpl"
	"github.com/warm3snow/practical-crypto/crypto/softimpl"
)

var (
	Csp CSP
)

// CSP - crypto service provider
type CSP interface {
	// Hash returns the hash data against the specified algorithm
	Hash(algo string, origin []byte) ([]byte, error)

	// Enc returns the encrypted data against the specified algorithm
	Enc(algo string, key, plainText []byte, mode string) ([]byte, error)

	// Dec returns the decrypted data against the specified algorithm
	Dec(algo string, key, cipherText []byte, mode string) ([]byte, error)

	// HMac returns the hmac data against the specified algorithm
	HMac(algo, key string, plain []byte) ([]byte, error)

	// Sign returns the signature data against the specified algorithm
	Sign(algo, key string, plain []byte) ([]byte, error)

	// Verify returns the verification result against the specified algorithm
	Verify(algo, key string, plain, sig []byte) (bool, error)

	// GenSymKey return a fixed length random key against algo
	GenSymKey(algo string, keySize int) ([]byte, error)
}

// InitCrypto returns a new crypto instance. support soft and hsm
func InitCrypto(soft bool, libPath string) (CSP, error) {
	if !soft {
		return hsmimpl.New(libPath)
	}
	return softimpl.New()
}
