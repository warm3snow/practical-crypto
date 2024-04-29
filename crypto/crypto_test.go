/**
 * @Author: xueyanghan
 * @File: crypto_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/11/20 16:30
 */

package crypto

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"github.com/warm3snow/gmsm/sm2"
	"github.com/warm3snow/gmsm/sm3"
	"github.com/warm3snow/gmsm/x509"
	"testing"
)

var (
	msg      = []byte("hello world")
	keyIndex = "1"
	symKey   = []byte("1234567890123456")

	sm2Key, _ = sm2.GenerateKey(rand.Reader)
	pubKey, _ = x509.MarshalSm2PublicKey(&sm2Key.PublicKey)
	priKey, _ = x509.MarshalSm2PrivateKey(sm2Key, nil)
)

func TestInitCryptoForHsmimpl(t *testing.T) {
	csp, err := InitCrypto(&CryptoConfig{
		Soft:    false,
		LibPath: "./hsmimpl/lib/libswsds.dylib",
	})
	assert.NoError(t, err)
	assert.NotNil(t, csp)

	// Test SM3
	dgst, err := csp.Hash("SM3", msg)
	assert.NoError(t, err)
	assert.Equal(t, dgst, sm3.Sm3Sum(msg))

	// Test HMAC_SM3
	dgst, err = csp.HMac("SM3", keyIndex, msg)
	assert.NoError(t, err)
	assert.NotNil(t, dgst)
	dgst2, err := csp.HMac("SM3", keyIndex, msg)
	assert.NoError(t, err)
	assert.NotNil(t, dgst2)
	assert.Equal(t, dgst, dgst2)

	sig, err := csp.Sign("SM2", keyIndex, "", msg)
	assert.NoError(t, err)
	assert.NotNil(t, sig)
	ok, err := csp.Verify("SM2", keyIndex, msg, sig)
	assert.NoError(t, err)
	assert.True(t, ok)

	cipherText, err := csp.Enc("SM4", keyIndex, "", msg, "CBC_PKCS5")
	assert.NoError(t, err)
	assert.NotNil(t, cipherText)
	plainText, err := csp.Dec("SM4", keyIndex, "", cipherText, "CBC_PKCS5")
	assert.NoError(t, err)
	assert.Equal(t, msg, plainText)
}

func TestInitCryptoForSoftimpl(t *testing.T) {
	csp, err := InitCrypto(&CryptoConfig{
		Soft: true,
	})
	assert.NoError(t, err)
	assert.NotNil(t, csp)

	// Test SM3
	dgst, err := csp.Hash("SM3", msg)
	assert.NoError(t, err)
	assert.Equal(t, dgst, sm3.Sm3Sum(msg))

	// Test HMAC_SM3
	dgst, err = csp.HMac("SM3", string(symKey), msg)
	assert.NoError(t, err)
	assert.NotNil(t, dgst)
	dgst2, err := csp.HMac("SM3", string(symKey), msg)
	assert.NoError(t, err)
	assert.NotNil(t, dgst2)
	assert.Equal(t, dgst, dgst2)

	sig, err := csp.Sign("SM2", string(priKey), "", msg)
	assert.NoError(t, err)
	assert.NotNil(t, sig)
	ok, err := csp.Verify("SM2", string(pubKey), msg, sig)
	assert.NoError(t, err)
	assert.True(t, ok)

	cipherText, err := csp.Enc("SM4", string(symKey), "", msg, "CBC_PKCS5")
	assert.NoError(t, err)
	assert.NotNil(t, cipherText)
	plainText, err := csp.Dec("SM4", string(symKey), "", cipherText, "CBC_PKCS5")
	assert.NoError(t, err)
	assert.Equal(t, msg, plainText)
}
