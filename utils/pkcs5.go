/**
 * @Author: xueyanghan
 * @File: pkcs5.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/20 14:26
 */

package utils

import "bytes"

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
