/**
 * @Author: xueyanghan
 * @File: sm2.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/12 17:34
 */

package tests

import (
	"crypto/rand"
	"fmt"
	"github.com/warm3snow/gmsm/sm2"
	"testing"
)

func TestSM2Encrypt1(t *testing.T) {
	// 生成密钥对
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 要加密的数据
	plainText := []byte("Hello, SM2-256!")

	// 使用公钥进行加密, SM2加解密有两种模式：C1C2C3和C1C3C2，这里使用C1C2C3（国密标准推荐使用C1C2C3）
	ciphertext, err := sm2.Encrypt(&privateKey.PublicKey, plainText, rand.Reader, sm2.C1C2C3)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}

	// 使用私钥进行解密
	decryptedText, err := sm2.Decrypt(privateKey, ciphertext, sm2.C1C2C3)
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}

	fmt.Println("Original data:", string(plainText))
	fmt.Println("Encrypted data:", ciphertext)
	fmt.Println("Decrypted data:", string(decryptedText))
}

func TestSM2Encrypt2(t *testing.T) {
	// 生成密钥对
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 要加密的数据
	plainText := []byte("Hello, SM2-256!")

	// 使用公钥进行加密
	ciphertext, err := sm2.EncryptAsn1(&privateKey.PublicKey, plainText, rand.Reader)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}

	// 使用私钥进行解密
	decryptedText, err := sm2.DecryptAsn1(privateKey, ciphertext)
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}

	fmt.Println("Original data:", string(plainText))
	fmt.Println("Encrypted data:", ciphertext)
	fmt.Println("Decrypted data:", string(decryptedText))
}

func TestSM2Sign(t *testing.T) {
	// 生成密钥对
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 要签名的数据
	plainText := []byte("Hello, SM2-256!")

	// 使用私钥进行签名
	sign, err := privateKey.Sign(rand.Reader, plainText, nil)
	if err != nil {
		fmt.Println("Error signing data:", err)
		return
	}

	// 使用公钥进行验签
	ok := privateKey.PublicKey.Verify(plainText, sign)
	if !ok {
		fmt.Println("Verify failed")
		return
	}

	fmt.Println("Verify success")
}
