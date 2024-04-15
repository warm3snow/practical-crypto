/**
 * @Author: xueyanghan
 * @File: pre_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/13 23:17
 */

package pre

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestPre(t *testing.T) {
	msg := "hello PRE!"

	t.Log("Test PRE")
	t.Log("待共享数据明文 M:", msg)

	// 1. KeyGen
	// alice
	skA, pkA, err := KeyGen()
	assert.NoError(t, err)

	// bob
	skB, pkB, err := KeyGen()
	assert.NoError(t, err)

	// 2. Encrypt
	capsule, C, err := Encrypt(pkA, []byte(msg))
	assert.NoError(t, err)

	t.Log("Alice加密数据:", hex.EncodeToString(C))
	t.Log("Alice加密数据 Capsule:", hex.EncodeToString(capsule))

	// 3. ReGenKey
	rk, random := ReGenKey(skA, pkB)

	// 4. ReEncrypt
	reCapsule, err := ReEncrypt(capsule, rk)
	assert.NoError(t, err)

	t.Log("Proxy重加密数据 Capsule':", hex.EncodeToString(reCapsule))

	// 5. Decrypt
	plain, err := Decrypt(reCapsule, skB, pkA, C, random)
	assert.NoError(t, err)

	t.Log("Bob解密数据 M:", new(big.Int).SetBytes(plain).String())
	t.Log("明文Msg：", new(big.Int).SetBytes([]byte(msg)).String())

	assert.Equal(t, msg, string(plain))

	t.Log("Bob解密数据 M:", string(plain))
}
