/**
 * @Author: xueyanghan
 * @File: rsa.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/12 17:48
 */

package tests

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestRSAEncrypt1(t *testing.T) {
	// 生成密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 使用私钥导出公钥
	publicKey := &privateKey.PublicKey

	// 要加密的数据
	plainText := []byte("Hello, RSA!")

	// 使用公钥进行加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}

	// 使用私钥进行解密
	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}

	fmt.Println("Original data:", string(plainText))
	fmt.Println("Encrypted data:", ciphertext)
	fmt.Println("Decrypted data:", string(decryptedText))
}

func TestRSAEncrypt2(t *testing.T) {
	// 生成密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 使用私钥导出公钥
	publicKey := &privateKey.PublicKey

	// 要加密的数据
	plainText := []byte("Hello, RSA!")

	// 使用公钥进行加密
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, plainText, nil)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}

	// 使用私钥进行解密
	decryptedText, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}

	fmt.Println("Original data:", string(plainText))
	fmt.Println("Encrypted data:", ciphertext)
	fmt.Println("Decrypted data:", string(decryptedText))
}

func TestRSASign1(t *testing.T) {
	// 生成密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 使用私钥对数据进行签名
	plainText := []byte("Hello, RSA!")
	hashed := sha256.Sum256(plainText)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println("Error signing data:", err)
		return
	}

	// 使用公钥对签名进行验证
	publicKey := &privateKey.PublicKey
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Println("Error verifying signature:", err)
		return
	}

	fmt.Println("Signature verified")
}

func TestRSASign2(t *testing.T) {
	// 生成密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 使用私钥对数据进行签名
	plainText := []byte("Hello, RSA!")
	hashed := sha256.Sum256(plainText)
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], nil)
	if err != nil {
		fmt.Println("Error signing data:", err)
		return
	}

	// 使用公钥对签名进行验证
	publicKey := &privateKey.PublicKey
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, hashed[:], signature, nil)
	if err != nil {
		fmt.Println("Error verifying signature:", err)
		return
	}

	fmt.Println("Signature verified")
}
