/**
 * @Author: xueyanghan
 * @File: key.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/19 17:13
 */

package softimpl

import (
	"crypto"
	"crypto/x509"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/pkg/errors"
	tjx509 "github.com/tjfoc/gmsm/x509"
)

// ParsePrivateKey parse bytes to a private key.
func ParsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}

	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, nil
	}

	if key, err := tjx509.ParsePKCS8UnecryptedPrivateKey(der); err == nil {
		return key, nil
	}

	// Serialization for bitcoin signature key: encode ECC numbers with hex
	Secp256k1Key, _ := btcec.PrivKeyFromBytes(der)
	key := Secp256k1Key.ToECDSA()
	return key, nil
}

func ParsePublicKey(der []byte) (crypto.PublicKey, error) {
	if key, err := x509.ParsePKCS1PublicKey(der); err == nil {
		return key, nil
	}

	if key, err := x509.ParsePKIXPublicKey(der); err == nil {
		return key, nil
	}

	if key, err := tjx509.ParseSm2PublicKey(der); err == nil {
		return key, nil
	}

	// Serialization for bitcoin signature key: encode ECC numbers with hex
	if key, err := btcec.ParsePubKey(der); err == nil {
		return key.ToECDSA(), nil
	}

	return nil, errors.New("failed to parse public key")
}
