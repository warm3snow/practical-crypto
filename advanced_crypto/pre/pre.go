/**
 * @Author: xueyanghan
 * @File: pre.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/13 21:07
 */

package pre

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/warm3snow/gmsm/sm2"
	"github.com/warm3snow/gmsm/sm4"
	"io"
	"math/big"
)

/*
1. $本地公私钥：<pk_A, sk_A>, 其中：pk_A = sk_A*G$
2. 待共享数据：$M$
3. Alice加密数据：$ Capsule = Encrypt(M, pk_A) $
4. Alice生成重加密密钥：$	 sk_{A->B} = ReGenKey(r, pk_A, pk_B, \alpha)$
5. Proxy重加密：$Capsule^{'} = ReEncrypt(Capsule, sk_{A->B})$
6. Bob本地解密： $ M = Decrypt(Capsule^{'}, sk_B, \alpha) $

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

func Encrypt(pk *sm2.PublicKey, m []byte) ([]byte, []byte, error) {
	// 生成E，V, r
	randE, _ := sm2.GenerateKey(rand.Reader)
	eBytes := append(pk.X.Bytes(), pk.Y.Bytes()...)
	randV, _ := sm2.GenerateKey(rand.Reader)
	vBytes := append(randV.X.Bytes(), randV.Y.Bytes()...)

	params := pk.Curve.Params()
	b := make([]byte, params.BitSize/8+8)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return nil, nil, err
	}
	r := new(big.Int).SetBytes(b)

	// 计算s
	s := new(big.Int).Set(randV.D)
	h2 := new(big.Int).SetBytes(Hash(append(eBytes, vBytes...)))
	r = r.Mul(r, h2)
	s = s.Add(s, r)

	// 生成K
	x, y := pk.ScalarMult(pk.X, pk.Y, new(big.Int).Add(randE.D, randV.D).Bytes())
	K := Hash(append(x.Bytes(), y.Bytes()...))[0:sm4.BlockSize]

	//  数据加密
	C, err := sm4.Sm4Ecb(K, m, true)
	if err != nil {
		fmt.Println("Failed to encrypt the plaintext:", err)
		return nil, nil, err
	}

	var capsule Capsule
	capsule.E.X, capsule.E.Y = randE.X, randE.Y
	capsule.V.X, capsule.V.Y = randV.X, randV.Y
	capsule.S = s

	capBytes, err := json.Marshal(capsule)
	if err != nil {
		fmt.Println("Failed to marshal the ciphertext:", err)
		return nil, nil, err
	}

	return capBytes, C, nil
}

func ReGenKey(skA *sm2.PrivateKey, pkB *sm2.PublicKey) (*big.Int, []byte) {
	// 生成随机数x
	randX, _ := sm2.GenerateKey(rand.Reader)
	xBytes := append(randX.X.Bytes(), randX.Y.Bytes()...)

	// 计算dInv
	pkBBytes := append(pkB.X.Bytes(), pkB.Y.Bytes()...)
	abDHX, abDHY := Curve.ScalarMult(pkB.X, pkB.Y, skA.D.Bytes())
	abDHBytes := append(abDHX.Bytes(), abDHY.Bytes()...)
	d := HashToCurve(append(xBytes, append(pkBBytes, abDHBytes...)...))
	dInv := new(big.Int).Set(d)
	dInv.ModInverse(dInv, skA.Curve.Params().N)

	//fmt.Printf("ReGenKey d: %s\n", d.String())
	//fmt.Printf("ReGenKey dInv: %s\n", dInv.String())

	// 计算rk
	rk := new(big.Int).Set(skA.D)
	rk = rk.Mul(rk, dInv)

	return rk, xBytes
}

func ReEncrypt(capsule []byte, rk *big.Int) ([]byte, error) {
	var c Capsule
	err := json.Unmarshal(capsule, &c)
	if err != nil {
		fmt.Println("Failed to unmarshal the ciphertext:", err)
		return nil, err
	}
	// 增加条件判断， FIXME
	E1x, E1y := Curve.ScalarMult(c.E.X, c.E.Y, rk.Bytes())
	V1x, V1y := Curve.ScalarMult(c.V.X, c.V.Y, rk.Bytes())

	c.E.X, c.E.Y = E1x, E1y
	c.V.X, c.V.Y = V1x, V1y

	cBytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Failed to marshal the ciphertext:", err)
		return nil, err
	}

	return cBytes, nil
}

func Decrypt(capsule []byte, skB *sm2.PrivateKey, pkA *sm2.PublicKey, cipherText, random []byte) ([]byte, error) {
	var c Capsule
	err := json.Unmarshal(capsule, &c)
	if err != nil {
		fmt.Println("Failed to unmarshal the ciphertext:", err)
		return nil, err
	}

	// 计算K
	pkBBytes := append(skB.X.Bytes(), skB.Y.Bytes()...)

	abDHx, abDHy := Curve.ScalarMult(pkA.X, pkA.Y, skB.D.Bytes())
	abDHBytes := append(abDHx.Bytes(), abDHy.Bytes()...)

	d := HashToCurve(append(random, append(pkBBytes, abDHBytes...)...))
	//fmt.Printf("Decrypt d: %s\n", d.String())

	evAddx, evAddy := Curve.Add(c.E.X, c.E.Y, c.V.X, c.V.Y)
	x, y := Curve.ScalarMult(evAddx, evAddy, d.Bytes())
	K := Hash(append(x.Bytes(), y.Bytes()...))[0:sm4.BlockSize]

	// 解密
	return sm4.Sm4Ecb(K, cipherText, false)
}
