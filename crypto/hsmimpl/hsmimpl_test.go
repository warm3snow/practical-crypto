package hsmimpl

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/warm3snow/gmsm/sm3"
	"log"
	"os"
	"runtime"
	"strings"
	"testing"
)

var (
	msg           = []byte("hello world")
	KeyPwdForKey1 = "11111111"
)

func libPath() string {
	wd, _ := os.Getwd()
	if runtime.GOOS == "darwin" {
		return wd + "/lib/libswsds.dylib"
	} else if runtime.GOOS == "linux" {
		return wd + "/lib/libswsds.so"
	}
	return wd
}

func TestSM3(t *testing.T) {
	csp, err := New(libPath())
	assert.NoError(t, err)

	msgHash, err := csp.Hash("SM3", msg)
	assert.NoError(t, err)

	assert.Equal(t, msgHash, sm3.Sm3Sum(msg))
}

func TestSM3HMac(t *testing.T) {
	csp, err := New(libPath())
	assert.NoError(t, err)

	msgHmac, err := csp.HMac("SM3", "1", msg)
	assert.NoError(t, err)

	log.Println(hex.EncodeToString(msgHmac))
}

func TestSM2(t *testing.T) {
	csp, err := New(libPath())
	assert.NoError(t, err)
	//Test SM2 enc and dec

	for i := 0; i < 100; i++ {
		signature, err := csp.Sign("SM2", "1", KeyPwdForKey1, msg)
		assert.NoError(t, err)
		pass, err := csp.Verify("SM2", "1", msg, signature)
		assert.NoError(t, err)
		assert.True(t, pass)
		fmt.Printf("signature: %s\n", hex.EncodeToString(signature))
	}

	//cipherText, err := csp.Enc("SM2", "1", KeyPwdForKey1, msg, "")
	//assert.NoError(t, err)
	//plainText, err := csp.Dec("SM2", "1", KeyPwdForKey1, cipherText, "")
	//assert.NoError(t, err)
	//assert.Equal(t, msg, plainText)
}

func TestSM4(t *testing.T) {
	csp, err := New(libPath())
	assert.NoError(t, err)

	//ecb mode
	cipherText, err := csp.Enc("SM4", "1", "", msg, "ECB")
	assert.NoError(t, err)
	plainText, err := csp.Dec("SM4", "1", "", cipherText, "ECB")
	assert.NoError(t, err)
	assert.Equal(t, msg, plainText)

	// cbc mode
	cipherText, err = csp.Enc("SM4", "1", "", msg, "CBC_PKCS5")
	assert.NoError(t, err)
	plainText, err = csp.Dec("SM4", "1", "", cipherText, "CBC_PKCS5")
	assert.NoError(t, err)
	assert.Equal(t, msg, plainText)
}

func TestHMacParallel(t *testing.T) {
	csp, err := New(libPath())
	assert.NoError(t, err)

	num := 100
	doneChan := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			v, err := csp.HMac("SM3", "1", msg)
			assert.NoError(t, err)
			log.Println(hex.EncodeToString(v))

			doneChan <- struct{}{}
		}()
	}
	for i := 0; i < num; i++ {
		<-doneChan
	}
}

func TestSM4Parallel(t *testing.T) {
	csp, err := New(libPath())
	assert.NoError(t, err)

	num := 20
	doneChan := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		msg = []byte(strings.Repeat("hello world", i+1))
		go func() {
			cipherText, err := csp.Enc("SM4", "1", "12321", msg, "CBC_PKCS5")
			assert.NoError(t, err)

			plainText, err := csp.Dec("SM4", "1", "12321", cipherText, "CBC_PKCS5")
			assert.NoError(t, err)
			assert.Equal(t, msg, plainText)

			fmt.Printf("cipherText: %s\tplainText: %s\n", hex.EncodeToString(cipherText), string(plainText))

			doneChan <- struct{}{}
		}()
	}
	for i := 0; i < num; i++ {
		<-doneChan
	}
}

func TestSM4Sequential(t *testing.T) {
	csp, err := New(libPath())
	assert.NoError(t, err)

	num := 50
	for i := 0; i < num; i++ {
		msg = []byte(strings.Repeat("hello world", i+1))
		cipherText, err := csp.Enc("SM4", "1", "12321", msg, "ECB")
		assert.NoError(t, err)

		plainText, err := csp.Dec("SM4", "1", "12321", cipherText, "ECB")
		assert.NoError(t, err)
		assert.Equal(t, msg, plainText)
	}
}

func TestSM2Parallel(t *testing.T) {
	csp, err := New("")
	assert.NoError(t, err)

	num := 100
	doneChan := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			signature, err := csp.Sign("SM2", "1", KeyPwdForKey1, msg)
			assert.NoError(t, err)

			pass, err := csp.Verify("SM2", "1", msg, signature)
			assert.NoError(t, err)
			assert.True(t, pass)

			doneChan <- struct{}{}
		}()
	}
	for i := 0; i < num; i++ {
		<-doneChan
	}
}

func TestSM2NTimes(t *testing.T) {
	csp, err := New("")
	assert.NoError(t, err)

	num := 1000

	for i := 0; i < num; i++ {
		signature, err := csp.Sign("SM2", "1", KeyPwdForKey1, msg)
		assert.NoError(t, err)

		pass, err := csp.Verify("SM2", "1", msg, signature)
		assert.NoError(t, err)
		assert.True(t, pass)
		if err != nil || !pass {
			fmt.Printf("signature: %s\n", hex.EncodeToString(signature))
			fmt.Printf("err: %s\n", err.Error())
		}
	}
}
