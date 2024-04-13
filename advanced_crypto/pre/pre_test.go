/**
 * @Author: xueyanghan
 * @File: pre_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/13 23:17
 */

package pre

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestPre(t *testing.T) {
	msg := "hello PRE!"
	alpha := []byte("1234567890")

	// 1. KeyGen
	_, pkA, err := KeyGen()
	assert.NoError(t, err)

	// 2. Encrypt
	randNumber := make([]byte, 32)
	_, err = rand.Read(randNumber)
	assert.NoError(t, err)

	r := new(big.Int).SetBytes(randNumber)
	c, err := Encrypt(r, pkA, []byte(msg))
	assert.NoError(t, err)

	// 3. ReGenKey
	skB, pkB, err := KeyGen()
	assert.NoError(t, err)

	rkAB := ReGenKey(r, pkA, pkB, alpha)

	// 4. ReEncrypt
	cipher, err := ReEncrypt(c, rkAB)
	assert.NoError(t, err)

	// 5. Decrypt
	plain, err := Decrypt(cipher, skB, alpha)
	assert.NoError(t, err)
	assert.Equal(t, msg, string(plain))
}
