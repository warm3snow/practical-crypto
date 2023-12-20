/**
 * @Author: xueyanghan
 * @File: sm4.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/11/20 14:25
 */

package base

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

const (
	BLOCK_SIZE = 16
)

func SM4Encrypt(c *Ctx, s SessionHandle, keyIndex uint, origin []byte, blockMode string) ([]byte, error) {
	keyHandle, err := c.SDFGetSymmKeyHandle(s, keyIndex)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to get sym keyHandle, keyIndex = %d", keyIndex)
	}
	defer func() {
		err := c.SDFDestroyKey(s, keyHandle)
		if err != nil {
			log.Printf("failed to destroy key, %s\n", err)
			return
		}
	}()
	iv := make([]byte, BLOCK_SIZE)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	var cipherWithPad []byte
	switch blockMode {
	case "CBC_PKCS5":
		plainWithPad := PKCS5Padding(origin, BLOCK_SIZE)
		// must copy iv, SDFEncrypt of cbc mode will change iv in place
		iv2 := make([]byte, len(iv))
		copy(iv2, iv)
		out, outLen, err := c.SDFEncrypt(s, keyHandle, SGD_SMS4_CBC, iv2, plainWithPad, uint(len(plainWithPad)))
		if err != nil {
			return nil, err
		}
		cipherWithPad = append(iv, out[:outLen]...)
		//fmt.Printf("iv: %x\n", iv)
		//fmt.Printf("iv2: %x\n", iv2)

	case "ECB":
		plainWithPad := PKCS5Padding(origin, BLOCK_SIZE)
		iv2 := make([]byte, len(iv))
		copy(iv2, iv)
		out, outLen, err := c.SDFEncrypt(s, keyHandle, SGD_SMS4_ECB, iv, plainWithPad, uint(len(plainWithPad)))
		if err != nil {
			return nil, err
		}
		cipherWithPad = out[:outLen]
	default:
		return nil, fmt.Errorf("sm4 encryption fails: unknown cipher block mode [%s]", blockMode)
	}

	return cipherWithPad, nil
}

func SM4Decrypt(c *Ctx, s SessionHandle, keyIndex uint, ciphertext []byte, blockMode string) ([]byte, error) {
	if len(ciphertext) < BLOCK_SIZE {
		return nil, errors.New("invalid ciphertext length")
	}
	keyHandle, err := c.SDFGetSymmKeyHandle(s, keyIndex)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to get sym keyHandle, keyIndex = %d", keyIndex)
	}
	defer func() {
		err := c.SDFDestroyKey(s, keyHandle)
		if err != nil {
			log.Printf("failed to destroy key, %s\n", err)
			return
		}
	}()

	switch blockMode {
	case "CBC_PKCS5":
		// must copy iv, SDFEncrypt of cbc mode will change iv in place
		iv := ciphertext[:BLOCK_SIZE]
		iv2 := make([]byte, len(iv))
		copy(iv2, iv)
		out, outLen, err := c.SDFDecrypt(s, keyHandle, SGD_SMS4_CBC, iv2, ciphertext[BLOCK_SIZE:], uint(len(ciphertext[BLOCK_SIZE:])))
		if err != nil {
			return nil, err
		}
		return PKCS5UnPadding(out[:outLen])

	case "ECB":
		out, outLen, err := c.SDFDecrypt(s, keyHandle, SGD_SMS4_ECB, nil, ciphertext, uint(len(ciphertext)))
		if err != nil {
			return nil, err
		}
		return PKCS5UnPadding(out[:outLen])
	default:
		return nil, fmt.Errorf("sm4 encryption fails: unknown cipher block mode [%s]", blockMode)
	}
}

// SM4GenKey generates a random key with the given key size.
func SM4GenKey(c *Ctx, s SessionHandle, keySize uint) ([]byte, error) {
	if keySize != 16 {
		keySize = 16
	}
	return c.SDFGenerateRandom(s, keySize)
}

// PKCS5Padding padding with pkcs5
func PKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS5UnPadding un padding with pkcs5
func PKCS5UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	paddingLen := int(data[length-1])
	if paddingLen > length {
		return nil, errors.New("decrypt failed,please check it")
	}
	return data[:(length - paddingLen)], nil
}
