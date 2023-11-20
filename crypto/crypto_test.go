/**
 * @Author: xueyanghan
 * @File: crypto_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/11/20 16:30
 */

package crypto

import (
	"github.com/stretchr/testify/assert"
	"github.com/warm3snow/gmsm/sm3"
	"testing"
)

var (
	msg      = []byte("hello world")
	keyIndex = "1"
)

func TestInitCrypto(t *testing.T) {
	soft := false
	libPath := "./hsmimpl/lib/libswsds.dylib"
	csp, err := InitCrypto(soft, libPath)
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

	sig, err := csp.Sign("SM2", keyIndex, msg, []byte("11111111"))
	assert.NoError(t, err)
	assert.NotNil(t, sig)
	ok, err := csp.Verify("SM2", keyIndex, msg, sig)
	assert.NoError(t, err)
	assert.True(t, ok)

	cipherText, err := csp.Enc("SM4", keyIndex, msg, "CBC_PKCS5")
	assert.NoError(t, err)
	assert.NotNil(t, cipherText)
	plainText, err := csp.Dec("SM4", keyIndex, cipherText, "CBC_PKCS5")
	assert.NoError(t, err)
	assert.Equal(t, msg, plainText)
}
