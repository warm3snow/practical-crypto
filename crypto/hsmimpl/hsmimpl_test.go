package hsmimpl

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/warm3snow/gmsm/sm3"
	"log"
	"os"
	"runtime"
	"testing"
)

var (
	msg = []byte("hello world")
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

	signature, err := csp.Sign("SM2", "1", msg, []byte("11111111"))
	assert.NoError(t, err)

	fmt.Printf("signature: %s\n", hex.EncodeToString(signature))

	pass, err := csp.Verify("SM2", "1", msg, signature)
	assert.NoError(t, err)
	assert.True(t, pass)
}

func TestSM4(t *testing.T) {
	csp, err := New(libPath())
	assert.NoError(t, err)

	//ecb mode
	cipherText, err := csp.Enc("SM4", "1", msg, "ECB")
	assert.NoError(t, err)

	plainText, err := csp.Dec("SM4", "1", cipherText, "ECB")
	assert.NoError(t, err)
	assert.Equal(t, msg, plainText)

	// cbc mode
	cipherText, err = csp.Enc("SM4", "1", msg, "CBC_PKCS5")
	assert.NoError(t, err)

	plainText, err = csp.Dec("SM4", "1", cipherText, "CBC_PKCS5")
	assert.NoError(t, err)
	assert.Equal(t, msg, plainText)
}
