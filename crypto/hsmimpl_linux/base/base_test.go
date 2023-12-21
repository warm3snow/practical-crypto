package base

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/warm3snow/gmsm/sm3"
	"github.com/warm3snow/practical-crypto/utils"
	"log"
	"os"
	"runtime"
	"testing"
)

var (
	plain    = []byte("helloworld")
	keyIndex = uint(1)
)

func libPath() string {
	wd, _ := os.Getwd()
	if runtime.GOOS == "darwin" {
		return wd + "/../lib/libswsds.dylib"
	} else if runtime.GOOS == "linux" {
		return wd + "/../lib/libswsds.so"
	}
	return wd
}

func Connect(t *testing.T) (c *Ctx, deviceHandle SessionHandle, sessionHandle SessionHandle) {
	p := libPath()
	c = New(p)
	deviceHandle, err := c.SDFOpenDevice()
	if err != nil {
		t.Fatal("open device error: ", err)
	}

	sessionHandle, err = c.SDFOpenSession(deviceHandle)
	if err != nil {
		t.Fatal("open session error: ", err)
	}
	return
}
func Release(t *testing.T, c *Ctx, deviceHandle SessionHandle, sessionHandle SessionHandle) {
	fmt.Println("defer func: Close Session...")
	err := c.SDFCloseSession(sessionHandle)
	if err != nil {
		t.Fatal("close session error: ", err)
	}

	fmt.Println("defer func: Close Device...")
	err = c.SDFCloseDevice(deviceHandle)
	if err != nil {
		t.Fatal("close device error: ", err)
	}
	c.Destroy()
}

// 基础函数测试
func TestGetDeviceInfo(t *testing.T) {
	//t.Skip()
	c, d, s := Connect(t)
	defer Release(t, c, d, s)

	randomNum, err := c.SDFGenerateRandom(s, 16)
	if err != nil {
		t.Fatal("generate random error: ", err)
	}
	fmt.Println("random: ", randomNum)
	var info DeviceInfo
	info, err = c.SDFGetDeviceInfo(s)
	if err != nil {
		t.Fatal("get device information error: ", err)
	}
	fmt.Println("deviceInfo IssuerName: ", info.IssuerName)
	fmt.Println("deviceInfo DeviceName: ", info.DeviceName)
	fmt.Println("deviceInfo DeviceSerial: ", info.DeviceSerial)
	fmt.Println("deviceInfo DeviceVersion: ", info.DeviceVersion)
	fmt.Println("deviceInfo StandardVersion: ", info.StandardVersion)
	fmt.Println("deviceInfo AsymAlgAbility: ", info.AsymAlgAbility)
	fmt.Println("deviceInfo SymAlgAbility: ", info.SymAlgAbility)
	fmt.Println("deviceInfo HashAlgAbility: ", info.HashAlgAbility)
	fmt.Println("deviceInfo BufferSize: ", info.BufferSize)
}

func TestExportECCPulicKey(t *testing.T) {
	t.Skip()

	c, d, s := Connect(t)
	defer Release(t, c, d, s)

	fmt.Println("===SDFExportSignPublicKey_ECC===")
	var signPublicKey ECCrefPublicKey
	signPublicKey, err := c.SDFExportSignPublicKey_ECC(s, 1)
	assert.NoError(t, err)
	fmt.Println("SignPublic Key Bits", signPublicKey.Bits)
	fmt.Println("SignPublic Key X", []byte(signPublicKey.X))
	fmt.Println("SignPublic Key Y", []byte(signPublicKey.Y))

	fmt.Println("===SDFExportEncPublicKey_ECC===")
	var encPublicKey ECCrefPublicKey
	encPublicKey, err = c.SDFExportEncPublicKey_ECC(s, keyIndex)
	assert.NoError(t, err)
	fmt.Println("EncPublic Key Bits", encPublicKey.Bits)
	fmt.Println("EncPublic Key X", []byte(encPublicKey.X))
	fmt.Println("EncPublic Key Y", []byte(encPublicKey.Y))

}

func TestInternalECCSign(t *testing.T) {
	t.Skip()

	c, d, s := Connect(t)
	defer Release(t, c, d, s)

	err := c.SDFGetPrivateKeyAccessRight(s, keyIndex+10000, []byte("111"), 3)
	assert.NoError(t, err)

	inHashData := sm3.Sm3Sum(plain)
	fmt.Println("===SDFInternalSign_ECC===")
	signature, err := c.SDFInternalSign_ECC(s, keyIndex, inHashData, uint(len(inHashData)))
	assert.NoError(t, err)

	fmt.Println("===SDFInternalVerify_ECC===")
	err = c.SDFInternalVerify_ECC(s, keyIndex, inHashData, uint(len(inHashData)), signature)
	assert.NoError(t, err)
}

func TestInternalECCEnc(t *testing.T) {
	t.Skip()

	c, d, s := Connect(t)
	defer Release(t, c, d, s)

	err := c.SDFGetPrivateKeyAccessRight(s, keyIndex+10000, []byte("111"), 3)
	assert.NoError(t, err)
	defer c.SDFReleasePrivateKeyAccessRight(s, keyIndex)

	inHashData := sm3.Sm3Sum(plain)
	fmt.Println("===SDFInternalEncrypt_ECC===")
	encData, err := c.SDFInternalEncrypt_ECC(s, keyIndex, SGD_SM2_3, inHashData, uint(len(inHashData)))
	assert.NoError(t, err)

	fmt.Println("===SDFInternalDecrypt_ECC===")
	data, _, err := c.SDFInternalDecrypt_ECC(s, keyIndex, SGD_SM2_3, encData)
	assert.NoError(t, err)
	assert.Equal(t, inHashData, data)
}

func TestExternalECCSign(t *testing.T) {
	t.Skip()

	c, d, s := Connect(t)
	defer Release(t, c, d, s)

	var publicKey ECCrefPublicKey
	var privateKey ECCrefPrivateKey
	fmt.Println("===SDFGenerateKeyPair_ECC===")
	publicKey, privateKey, err := c.SDFGenerateKeyPair_ECC(s, SGD_SM2, 256)
	assert.NoError(t, err)

	fmt.Println("===SDFExternalSign_ECC===")
	inputData := sm3.Sm3Sum(plain)
	signData, err := c.SDFExternalSign_ECC(s, SGD_SM2_1, privateKey, inputData, uint(len(inputData)))
	if err != nil {
		fmt.Println("External Sign error: ", err)
	}
	fmt.Printf("signData R %x \n", []byte(signData.R))
	fmt.Printf("signData S %x \n", []byte(signData.S))

	fmt.Println("===SDFExternalVerify_ECC===")
	err = c.SDFExternalVerify_ECC(s, SGD_SM2_1, publicKey, inputData, uint(len(inputData)), signData)
	assert.NoError(t, err)
}

func TestExternalECCEnc(t *testing.T) {
	t.Skip()

	c, d, s := Connect(t)
	defer Release(t, c, d, s)

	var publicKey ECCrefPublicKey
	var privateKey ECCrefPrivateKey
	fmt.Println("===SDFGenerateKeyPair_ECC===")
	publicKey, privateKey, err := c.SDFGenerateKeyPair_ECC(s, SGD_SM2, 256)
	fmt.Printf("publicKey: %+v\n", publicKey)
	fmt.Printf("privateKey: %+v\n", privateKey)

	inputData := []byte{0xbc, 0xa3, 0xde, 0xa1, 0x2f, 0x89, 0xd7, 0x78, 0xe5, 0xb7, 0x0b, 0x86, 0x7d, 0x1e, 0x36, 0x0e, 0x93, 0x7d, 0x47, 0xcb, 0xbb, 0xac, 0x39, 0x06, 0x35, 0x81, 0xa4, 0xe1, 0x85, 0x76, 0x57, 0x31}
	fmt.Println("===SDFExternalEncrypt_ECC===")
	encData, err := c.SDFExternalEncrypt_ECC(s, SGD_SM2_3, publicKey, inputData, 32)
	assert.NoError(t, err)

	fmt.Println("===SDFExternalDecrypt_ECC===")
	decData, _, err := c.SDFExternalDecrypt_ECC(s, SGD_SM2_3, privateKey, encData)
	assert.NoError(t, err)

	assert.Equal(t, inputData, decData)
}

func TestEncryptFunc_SM4CBC(t *testing.T) {
	t.Skip()

	c, d, s := Connect(t)
	defer Release(t, c, d, s)

	keyHandle, err := c.SDFGetSymmKeyHandle(s, uint(1))
	assert.NoError(t, err)

	iv := make([]byte, 16)
	_, err = rand.Read(iv)
	assert.NoError(t, err)

	//copy it!! SDFEncrypt will change iv :-(
	iv2 := make([]byte, 16)
	copy(iv2, iv)

	plainWithPad := utils.PKCS5Padding(plain, 16)
	encData, encDataLength, err := c.SDFEncrypt(s, keyHandle, SGD_SMS4_CBC, iv, plainWithPad, uint(len(plainWithPad)))
	assert.NoError(t, err)

	decdata, decdataLength, err := c.SDFDecrypt(s, keyHandle, SGD_SMS4_CBC, iv2, encData, encDataLength)
	assert.NoError(t, err)

	plain2 := utils.PKCS5UnPadding(decdata[:decdataLength])
	assert.NoError(t, err)

	//fmt.Printf("%d, %X\n", len(plain), plain)
	//fmt.Printf("%d, %X\n", len(plainWithPad), plainWithPad)
	//fmt.Printf("%d, %X\n", encDataLength, encData)
	//fmt.Printf("%d, %X\n", decdataLength, decdata)

	assert.Equal(t, plainWithPad, decdata)
	assert.Equal(t, plain, plain2)
}

func TestEncryptFunc_SM4ECB(t *testing.T) {
	t.Skip()

	c, d, s := Connect(t)
	defer Release(t, c, d, s)

	keyHandle, err := c.SDFGetSymmKeyHandle(s, uint(1))
	assert.NoError(t, err)

	plainWithPad := utils.PKCS5Padding(plain, 16)
	encData, encDataLength, err := c.SDFEncrypt(s, keyHandle, SGD_SMS4_ECB, nil, plainWithPad, uint(len(plainWithPad)))
	assert.NoError(t, err)

	decdata, _, err := c.SDFDecrypt(s, keyHandle, SGD_SMS4_ECB, nil, encData, encDataLength)
	assert.NoError(t, err)
	assert.Equal(t, plainWithPad, decdata)

	plainUnpadding := utils.PKCS5UnPadding(decdata)
	assert.NoError(t, err)
	assert.Equal(t, plain, plainUnpadding)
}

func TestSM3(t *testing.T) {
	//t.Skip()
	c, d, s := Connect(t)
	defer Release(t, c, d, s)

	_, err := c.SDFHashInit(s, SGD_SM3, nil, 0)
	assert.NoError(t, err)

	err = c.SDFHashUpdate(s, plain, uint(len(plain)))
	assert.NoError(t, err)

	v, vl, err := c.SDFHashFinal(s)
	assert.NoError(t, err)
	assert.Equal(t, vl, uint(len(v)))
	assert.Equal(t, v, sm3.Sm3Sum(plain))
}

func TestHMAC(t *testing.T) {
	t.Skip()
	c, d, s := Connect(t)
	defer Release(t, c, d, s)

	k, err := c.SDFGetSymmKeyHandle(s, 1)
	assert.NoError(t, err)

	mac, macLen, err := c.SDFHMAC(s, k, SGD_SM3, plain, uint(len(plain)))
	assert.NoError(t, err)

	fmt.Println(macLen)
	fmt.Println(len(mac))
	assert.Equal(t, macLen, uint(len(mac)))

	log.Println(hex.EncodeToString(mac))
}
