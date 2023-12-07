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

func TestHsmSM3HMac(t *testing.T) {
	csp, err := New(libPath())
	assert.NoError(t, err)

	msgHmac, err := csp.HMac("SM3", "1", msg)
	assert.NoError(t, err)

	log.Println(hex.EncodeToString(msgHmac))
}

//func TestEncAndDec(t *testing.T) {
//	csp, err := New(libPath())
//	//Test SM2 enc and dec
//
//	cipherText, err := csp.Enc("SM2", "1", msg, "")
//	assert.NoError(t, err)
//	plainText, err = csp.Dec("SM2", priKey, cipherText, "")
//	assert.NoError(t, err)
//	assert.Equal(t, msg, plainText)
//
//	//Test SM4 enc and dec
//	key = []byte("1234567890123456")
//	cipherText, err = csp.Enc("SM4", key, msg, "")
//	assert.NoError(t, err)
//	plainText, err = csp.Dec("SM4", key, cipherText, "")
//	assert.NoError(t, err)
//	assert.Equal(t, msg, plainText)
//}

func TestSM2SignAndVerify(t *testing.T) {
	csp, err := New(libPath())
	//Test SM2 enc and dec

	signature, err := csp.Sign("SM2", "1", KeyPwdForKey1, msg)
	assert.NoError(t, err)

	fmt.Printf("signature: %s\n", hex.EncodeToString(signature))

	pass, err := csp.Verify("SM2", "1", msg, signature)
	assert.NoError(t, err)
	assert.True(t, pass)
}

func TestSM2EncAndDec(t *testing.T) {
	csp, err := New(libPath())
	assert.NoError(t, err)

	cipherText, err := csp.Enc("SM2", "1", KeyPwdForKey1, msg, "")
	assert.NoError(t, err)

	plainText, err := csp.Dec("SM2", "1", KeyPwdForKey1, cipherText, "")
	assert.NoError(t, err)
	assert.Equal(t, msg, plainText)
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

	num := 500
	doneChan := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			//v, err := csp.HMac("SM3", "1", msg)
			v, err := csp.Hash("SM3", msg)

			assert.NoError(t, err)

			log.Println(hex.EncodeToString(v))
			doneChan <- struct{}{}
		}()
	}
	for i := 0; i < num; i++ {
		<-doneChan
	}
}

func BenchmarkSM4(b *testing.B) {
	csp, err := New(libPath())
	assert.NoError(b, err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := csp.Enc("SM4", "1", "", msg, "CBC_PKCS5")
			assert.NoError(b, err)
		}
	})
}

func TestSM4Parallel(t *testing.T) {
	csp, err := New(libPath())
	assert.NoError(t, err)

	num := 50
	//doneChan := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		msg = []byte(strings.Repeat("hello world", i+1))
		//go func() {
		cipherText, err := csp.Enc("SM4", "1", "12321", msg, "ECB")
		assert.NoError(t, err)

		plainText, err := csp.Dec("SM4", "1", "12321", cipherText, "ECB")
		assert.NoError(t, err)
		assert.Equal(t, msg, plainText)

		//time.Sleep(10 * time.Second)
		//doneChan <- struct{}{}
		//}()
	}
	//for i := 0; i < num; i++ {
	//	<-doneChan
	//}
}
