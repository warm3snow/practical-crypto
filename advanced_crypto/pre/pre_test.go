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
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestPre(t *testing.T) {
	msg := "hello PRE!"
	alpha := []byte("1234567890")

	t.Log("Test PRE")
	t.Log("待共享数据明文 M:", msg)
	t.Log("Alice的自定时授权参数 alpha:", string(alpha))

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

	t.Log("Alice加密数据 C:", hex.EncodeToString(c))

	// 3. ReGenKey
	skB, pkB, err := KeyGen()
	assert.NoError(t, err)

	rkAB := ReGenKey(r, pkA, pkB, alpha)

	// 4. ReEncrypt
	cipher, err := ReEncrypt(c, rkAB)
	assert.NoError(t, err)

	t.Log("Proxy重加密数据 C':", hex.EncodeToString(cipher))

	//assert.NotEqual(t, c, cipher)
	//for i := 0; i < len(c); i++ {
	//	if c[i] != cipher[i] {
	//		t.Log("C[i] != C'[i], i:", i)
	//	}
	//}

	// 5. Decrypt
	plain, err := Decrypt(cipher, skB, alpha)
	assert.NoError(t, err)
	assert.Equal(t, msg, string(plain))

	t.Log("Bob解密数据 M:", string(plain))
}
