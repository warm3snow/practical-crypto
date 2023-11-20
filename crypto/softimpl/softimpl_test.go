/**
 * @Author: xueyanghan
 * @File: softimpl_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/17 17:16
 */

package softimpl

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"github.com/warm3snow/gmsm/sm2"
	"github.com/warm3snow/gmsm/x509"
	"testing"
)

func TestEncAndDec(t *testing.T) {
	var (
		msg        = []byte("hello softimpl")
		key        []byte
		plainText  []byte
		cipherText []byte
		err        error
	)
	csp := softimpl{}
	//Test SM2 enc and dec
	sm2Key, err := sm2.GenerateKey(rand.Reader)
	assert.NoError(t, err)
	pubKey, err := x509.MarshalSm2PublicKey(&sm2Key.PublicKey)
	assert.NoError(t, err)
	priKey, err := x509.MarshalSm2PrivateKey(sm2Key, nil)
	assert.NoError(t, err)

	cipherText, err = csp.Enc("SM2", string(pubKey), msg, "")
	assert.NoError(t, err)
	plainText, err = csp.Dec("SM2", string(priKey), cipherText, "")
	assert.NoError(t, err)
	assert.Equal(t, msg, plainText)

	//Test SM4 enc and dec
	key = []byte("1234567890123456")
	cipherText, err = csp.Enc("SM4", string(key), msg, "")
	assert.NoError(t, err)
	plainText, err = csp.Dec("SM4", string(key), cipherText, "")
	assert.NoError(t, err)
	assert.Equal(t, msg, plainText)
}

func TestEnc(t *testing.T) {
	//TODO
}
