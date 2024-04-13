/**
 * @Author: xueyanghan
 * @File: pre.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/13 21:07
 */

package pre

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/warm3snow/gmsm/sm2"
	"math/big"
)

/*
1. $本地公私钥：<pk_A, sk_A>, 其中：pk_A = sk_A*G$
2. 待共享数据：$M$
3. Alice加密数据：$ C = Encrypt(M, pk_A) $
4. Alice生成重加密密钥：$	 sk_{A->B} = ReGenKey(r, pk_A, pk_B, \alpha)$
5. Proxy重加密：$C^{'} = ReEncrypt(C, sk_{A->B})$
6. Bob本地解密： $ M = Decrypt(C^{'}, sk_B, \alpha) $

Bob本地公私钥:$<pk_B, sk_B>$, 其中：$pk_B = sk_B*G$
*/

//KeyGen, Encrypt, ReGenKey, ReEncrypt, Decrypt

func KeyGen() (*sm2.PrivateKey, *sm2.PublicKey, error) {
	// 生成SM2密钥对
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Failed to generate sender's SM2 key pair:", err)
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func Encrypt(r *big.Int, pk *sm2.PublicKey, m []byte) ([]byte, error) {
	//获取随机数
	//randNumber := make([]byte, 32)
	//_, err := rand.Read(randNumber)
	//if err != nil {
	//	fmt.Println("Failed to generate random number:", err)
	//	return nil, err
	//}
	var c C
	// calculate C1
	c.C1.X, c.C1.Y = sm2.P256Sm2().ScalarBaseMult(r.Bytes())
	xa, ya := sm2.P256Sm2().ScalarMult(pk.X, pk.Y, r.Bytes())
	t := sha256.Sum256(append(xa.Bytes(), ya.Bytes()...))

	// calculate C2
	c.C2 = make([]byte, len(m))
	for i := 0; i < len(m); i++ {
		c.C2[i] = m[i] ^ t[i]
	}

	// calculate C3
	c3Bytes := sha256.Sum256(append(append(xa.Bytes(), m...), ya.Bytes()...))
	c.C3 = c3Bytes[:]

	// calculate C4
	c1Bytes := append(c.C1.X.Bytes(), c.C1.Y.Bytes()...)
	c4Bytes := sha256.Sum256(append(append(m, c1Bytes...), c.C3...))
	c.C4 = c4Bytes[:]

	return json.Marshal(c)
}

func ReGenKey(r *big.Int, pkA, pkB *sm2.PublicKey, alpha []byte) *big.Int {
	// calculate t
	xa, ya := sm2.P256Sm2().ScalarMult(pkA.X, pkA.Y, r.Bytes())
	t := sha256.Sum256(append(xa.Bytes(), ya.Bytes()...))

	// calculate rk_{A->B}
	xb, yb := sm2.P256Sm2().ScalarMult(pkB.X, pkB.Y, r.Bytes())
	s := sha256.Sum256(append(append(xb.Bytes(), yb.Bytes()...), alpha...))

	// xor t and s
	rXorS := make([]byte, len(t))
	for i := 0; i < len(t); i++ {
		rXorS[i] = t[i] ^ s[i]
	}

	return new(big.Int).SetBytes(rXorS)
}

func ReEncrypt(cipher []byte, skAB *big.Int) ([]byte, error) {
	var c C
	err := json.Unmarshal(cipher, &c)
	if err != nil {
		fmt.Println("Failed to unmarshal the ciphertext:", err)
		return nil, err
	}
	// re-calculate C2
	newC2 := make([]byte, len(c.C2))
	for i := 0; i < len(c.C2); i++ {
		newC2[i] = skAB.Bytes()[i] ^ c.C2[i]
	}
	c.C2 = newC2

	return json.Marshal(c)
}

func Decrypt(cipher []byte, skB *sm2.PrivateKey, alpha []byte) ([]byte, error) {
	var c C
	err := json.Unmarshal(cipher, &c)
	if err != nil {
		fmt.Println("Failed to unmarshal the ciphertext:", err)
		return nil, err
	}

	// calculate M1
	xb, yb := sm2.P256Sm2().ScalarMult(c.C1.X, c.C1.Y, skB.D.Bytes())
	s := sha256.Sum256(append(append(xb.Bytes(), yb.Bytes()...), alpha...))
	m1 := make([]byte, len(c.C2))
	for i := 0; i < len(c.C2); i++ {
		m1[i] = c.C2[i] ^ s[i]
	}

	// calculate k
	c1Bytes := append(c.C1.X.Bytes(), c.C1.Y.Bytes()...)
	kBytes := sha256.Sum256(append(append(m1, c1Bytes...), c.C3...))
	k := kBytes[:]

	if bytes.Equal(k, c.C4) {
		return m1, nil
	}
	return nil, errors.New("Failed to decrypt the ciphertext")
}
