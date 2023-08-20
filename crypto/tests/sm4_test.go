/**
 * @Author: xueyanghan
 * @File: sm4_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/16 17:03
 */

package tests

import (
	"fmt"
	"github.com/warm3snow/gmsm/sm4"
	"testing"
)

func TestSM4Encrypt(t *testing.T) {
	// 生成密钥对
	var key = []byte("1234567890abcdef")
	// 要加密的数据
	plainText := []byte("Hello, SM4!")
	ciphertext, err := sm4.Sm4Cbc(key, plainText, true)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}

	// 使用私钥进行解密
	decryptedText, err := sm4.Sm4Cbc(key, ciphertext, false)
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}

	fmt.Println("Original data:", string(plainText))
	fmt.Println("Encrypted data:", ciphertext)
	fmt.Println("Decrypted data:", string(decryptedText))
}
