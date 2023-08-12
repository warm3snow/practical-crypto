/**
 * @Author: xueyanghan
 * @File: ecdsa_test.go
 * @Version: 1.0.0
 * @Description: ecdsa主要用于签名，无加解密功能
 * @Date: 2023/8/12 17:57
 */

package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"testing"
)

func TestECDSASign(t *testing.T) {
	// 生成密钥对,(ecdsa有很多曲线，在elliptic包中，我们这里使用p256曲线）
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 要加密的数据
	plainText := []byte("Hello, ECDSA!")

	// 使用公钥进行加密（这里使用私钥对数据签名，不是加密）
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, plainText)
	if err != nil {
		fmt.Println("Error signing data:", err)
		return
	}

	// 验证签名
	valid := ecdsa.Verify(&privateKey.PublicKey, plainText, r, s)
	if valid {
		fmt.Println("Signature is valid.")
	} else {
		fmt.Println("Signature is NOT valid.")
	}
}
