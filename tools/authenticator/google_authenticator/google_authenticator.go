/**
 * @Author: xueyanghan
 * @File: google_authenticator.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/12 18:32
 */

package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"time"
)

func generateTOTP(secret string, counter int64) (string, error) {
	// 1. 共享密钥解码：将Base32编码的密钥解码为字节数组
	decodedKey, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	// 2. HMAC-SHA-1：计算哈希值
	h := hmac.New(sha1.New, decodedKey)
	binary.Write(h, binary.BigEndian, counter)

	// 3. Truncate：对哈希值进行截断，只保留最后 4 个字节
	offset := h.Sum(nil)[19] & 0x0f
	truncatedHash := h.Sum(nil)[offset : offset+4]
	truncatedHash[0] = truncatedHash[0] & 0x7f
	otp := (int32(truncatedHash[0])<<24 | int32(truncatedHash[1])<<16 | int32(truncatedHash[2])<<8 | int32(truncatedHash[3])) % 1000000

	// 4. 格式化：将结果转换为十进制6位的字符串
	return fmt.Sprintf("%06d", otp), nil
}

func main() {
	// 设置共享密钥（这里使用Base32编码）
	secret := "ABCDEFGH234567MN"

	// 获取当前时间步长C
	currentTime := time.Now().Unix()
	timeStep := currentTime / 30

	// 生成TOTP
	otp, err := generateTOTP(secret, timeStep)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated OTP:", otp)
}
