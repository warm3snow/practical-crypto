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

	t.Logf("Proxy重加密密钥 rk: %s", hex.EncodeToString(rk.Bytes()))
	t.Log("Proxy重加密数据 Capsule':", hex.EncodeToString(reCapsule))


	// ReEncrypt again and check if the result is the same
	//reCapsule2, err := ReEncrypt(reCapsule, rk)
	//assert.NoError(t, err)
	//
	//t.Logf("Proxy重加密密钥 rk: %s", hex.EncodeToString(rk.Bytes()))
	//t.Log("Proxy重加密数据 Capsule2':", hex.EncodeToString(reCapsule))
	//
	//reCapsuleStruct := new(Capsule)
	//err = json.Unmarshal(reCapsule, reCapsuleStruct)
	//assert.NoError(t, err)
	//
	//reCapsule2Struct := new(Capsule)
	//err = json.Unmarshal(reCapsule2, reCapsule2Struct)
	//assert.NoError(t, err)

	//isEqual := reflect.DeepEqual(reCapsuleStruct, reCapsule2Struct)
	//if !isEqual {
	//	t.Errorf("ReEncrypt result is not the same")
	//}else{
	//	t.Logf("ReEncrypt result is the same")
	//}
	//



	// 5. Decrypt
	plain, err := Decrypt(reCapsule, skB, pkA, C, random)
	assert.NoError(t, err)

	t.Log("Bob解密数据 M:", new(big.Int).SetBytes(plain).String())
	t.Log("明文Msg：", new(big.Int).SetBytes([]byte(msg)).String())

	assert.Equal(t, msg, string(plain))

	t.Log("Bob解密数据 M:", string(plain))
}

func BenchmarkPre(b *testing.B) {
	msg := "hello PRE!"
	// 1. KeyGen
	// alice
	skA, pkA, err := KeyGen()
	assert.NoError(b, err)

	// bob
	skB, pkB, err := KeyGen()
	assert.NoError(b, err)

	// 2. Encrypt
	capsule, C, err := Encrypt(pkA, []byte(msg))
	assert.NoError(b, err)


	// 3. ReGenKey
	rk, random := ReGenKey(skA, pkB)

	b.ResetTimer()
	b.ReportAllocs()
	// 4. ReEncrypt and Decrypt
	for i := 0; i < b.N; i++ {
		//reCapsule, err := ReEncrypt(capsule, rk)
		_, err = ReEncrypt(capsule, rk)
		//assert.NoError(b, err)

		//plain, err := Decrypt(reCapsule, skB, pkA, C, random)
		//assert.NoError(b, err)
		//assert.Equal(b, msg, string(plain))
	}

	_, _, _ = skB, C, random
}