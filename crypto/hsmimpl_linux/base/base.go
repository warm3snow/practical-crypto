package base

/*
  #cgo CFLAGS: -I./
  #cgo linux LDFLAGS: -L./ -lswsds
  #cgo darwin LDFLAGS: -L ./ -ldl
  //#if defined(__linux__) || defined(linux)
  #include "swsds.h"
  #include <pthread.h>
  #include <stdlib.h>
  #include <stdio.h>
  #include <string.h>
  #include <unistd.h>
  //#endif
*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

func ConvertToDeviceInfoGo(deviceInfo1 C.DEVICEINFO) (deviceInfo DeviceInfo) {
	deviceInfo = DeviceInfo{
		IssuerName:      strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&deviceInfo1.IssuerName[0]), 40)), " "),
		DeviceName:      strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&deviceInfo1.DeviceName[0]), 16)), " "),
		DeviceSerial:    strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&deviceInfo1.DeviceSerial[0]), 16)), " "),
		DeviceVersion:   uint(deviceInfo1.DeviceVersion),
		StandardVersion: uint(deviceInfo1.StandardVersion),
		SymAlgAbility:   uint(deviceInfo1.SymAlgAbility),
		HashAlgAbility:  uint(deviceInfo1.HashAlgAbility),
		BufferSize:      uint(deviceInfo1.BufferSize),
	}
	temp1 := C.GoBytes(unsafe.Pointer(&deviceInfo1.AsymAlgAbility[0]), 2)
	temp2 := C.GoBytes(unsafe.Pointer(&deviceInfo1.AsymAlgAbility[1]), 2)
	deviceInfo.AsymAlgAbility[0] = uint(temp1[0])
	deviceInfo.AsymAlgAbility[1] = uint(temp2[0])
	return deviceInfo
}

func ConvertToECCrefPublicKeyC(publicKey ECCrefPublicKey) (pucPublicKey C.ECCrefPublicKey) {

	pucPublicKey.bits = C.SGD_UINT32(publicKey.Bits)
	for i := 0; i < len(publicKey.X); i++ {
		pucPublicKey.x[i] = C.SGD_UCHAR(publicKey.X[i])
	}
	for i := 0; i < len(publicKey.Y); i++ {
		pucPublicKey.y[i] = C.SGD_UCHAR(publicKey.Y[i])
	}
	return pucPublicKey
}

func ConvertToECCrefPublicKeyGo(pucPublicKey C.ECCrefPublicKey) (publicKey ECCrefPublicKey) {
	publicKey = ECCrefPublicKey{
		Bits: uint(pucPublicKey.bits),
		X:    strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPublicKey.x[0]), 64)), " "),
		Y:    strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPublicKey.y[0]), 64)), " "),
	}
	return publicKey
}

func ConvertToECCrefPrivateKeyC(privateKey ECCrefPrivateKey) (pucPrivateKey C.ECCrefPrivateKey) {
	pucPrivateKey.bits = C.SGD_UINT32(privateKey.Bits)
	for i := 0; i < len(privateKey.K); i++ {
		pucPrivateKey.K[i] = C.SGD_UCHAR(privateKey.K[i])
	}
	return pucPrivateKey
}

func ConvertToECCrefPrivateKeyGo(pucPrivateKey C.ECCrefPrivateKey) (privateKey ECCrefPrivateKey) {
	privateKey = ECCrefPrivateKey{
		Bits: uint(pucPrivateKey.bits),
		K:    strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPrivateKey.K[0]), 64)), " "),
	}
	return privateKey
}

func ConvertToECCCipherC(encData ECCCipher) (pucEncData C.ECCCipher) {
	for i := 0; i < len(encData.X); i++ {
		pucEncData.x[i] = C.SGD_UCHAR(encData.X[i])
	}
	for i := 0; i < len(encData.Y); i++ {
		pucEncData.y[i] = C.SGD_UCHAR(encData.Y[i])
	}
	for i := 0; i < len(encData.M); i++ {
		pucEncData.M[i] = C.SGD_UCHAR(encData.M[i])
	}
	pucEncData.L = C.SGD_UINT32(encData.L)
	for i := 0; i < len(encData.C); i++ {
		pucEncData.C[i] = C.SGD_UCHAR(encData.C[i])
	}
	return pucEncData
}
func ConvertToECCCipherGo(pucKey C.ECCCipher) (key ECCCipher) {
	key = ECCCipher{
		X: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucKey.x[0]), 64)), " "),
		Y: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucKey.y[0]), 64)), " "),
		M: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucKey.M[0]), 32)), " "),
		L: uint(pucKey.L),
		C: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucKey.C[0]), C.int(pucKey.L))), " "),
	}
	return key
}

func ConvertToECCSignatureC(signature ECCSignature) (pSignature C.ECCSignature) {
	for i := 0; i < len(signature.R); i++ {
		pSignature.r[i] = C.SGD_UCHAR(signature.R[i])
	}
	for i := 0; i < len(signature.S); i++ {
		pSignature.s[i] = C.SGD_UCHAR(signature.S[i])
	}
	return pSignature
}

func ConvertToECCSignatureGo(pucSignature C.ECCSignature) (signature ECCSignature) {
	signature = ECCSignature{
		R: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucSignature.r[0]), 64)), " "),
		S: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucSignature.s[0]), 64)), " "),
	}

	return signature
}

func New(libPath string) *Ctx {
	c := new(Ctx)
	//mod := C.CString(libPath)
	//defer C.free(unsafe.Pointer(mod))
	//c.libHandle = C.NewLib(mod)
	//if c.libHandle == nil {
	//	return nil
	//}
	return c
}

func ToError(e C.SGD_RV) error {
	if e == C.SDR_OK {
		return nil
	}
	return fmt.Errorf("sdf: 0x%X", uint(e))
}

func (c *Ctx) Destroy() {
	//if c == nil || c.libHandle == nil {
	//	return
	//}
	//C.DestroyLib(c.libHandle)
	//c.libHandle = nil
}

type Ctx struct {
	libHandle *C.struct_LibHandle
}

type SessionHandle C.SGD_HANDLE

var stubData = []byte{0}

func CMessage(data []byte) (dataPtr C.SGD_UCHAR_PRT) {
	l := len(data)
	if l == 0 {
		return nil
	}
	dataPtr = C.SGD_UCHAR_PRT(unsafe.Pointer(&data[0]))
	return dataPtr
}

// SDFOpenDevice 1.打开设备
func (c *Ctx) SDFOpenDevice() (deviceHandle SessionHandle, err error) {
	var err1 C.SGD_RV
	var dH C.SGD_HANDLE
	err1 = C.SDFOpenDevice(&dH)
	err = ToError(err1)
	deviceHandle = SessionHandle(dH)
	return deviceHandle, err
}

// SDFCloseDevice 2.关闭设备
func (c *Ctx) SDFCloseDevice(deviceHandle SessionHandle) (err error) {
	var err1 = C.SDFCloseDevice(C.SGD_HANDLE(deviceHandle))
	return ToError(err1)
}

// SDFOpenSession 3.创建会话
func (c *Ctx) SDFOpenSession(deviceHandle SessionHandle) (SessionHandle, error) {
	var s C.SGD_HANDLE
	var err1 = C.SDFOpenSession(C.SGD_HANDLE(deviceHandle), &s)
	return SessionHandle(s), ToError(err1)
}

// SDFCloseSession 4.关闭会话
func (c *Ctx) SDFCloseSession(sessionHandle SessionHandle) (err error) {
	var err1 = C.SDFCloseSession(C.SGD_HANDLE(sessionHandle))
	return ToError(err1)
}

// SDFGetDeviceInfo 5.获取设备信息
func (c *Ctx) SDFGetDeviceInfo(sessionHandle SessionHandle) (deviceInfo DeviceInfo, err error) {
	var deviceInfo1 C.DEVICEINFO
	var err1 = C.SDFGetDeviceInfo(C.SGD_HANDLE(sessionHandle), &deviceInfo1)
	deviceInfo = ConvertToDeviceInfoGo(deviceInfo1)
	err = ToError(err1)
	return deviceInfo, err
}

// SDFGenerateRandom 6.产生随机数
func (c *Ctx) SDFGenerateRandom(sessionHandle SessionHandle, length uint) (randomData []byte, err error) {
	var random C.SGD_UCHAR_PRT
	var err1 = C.SDFGenerateRandom(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(length), &random)
	err = ToError(err1)
	randomData = C.GoBytes(unsafe.Pointer(random), C.int(length))
	C.free(unsafe.Pointer(random))
	return randomData, err
}

// SDFGetPrivateKeyAccessRight 7.获取私钥使用权限
func (c *Ctx) SDFGetPrivateKeyAccessRight(sessionHandle SessionHandle, keyIndex uint, password []byte, pwdLength uint) (err error) {
	var err1 = C.SDFGetPrivateKeyAccessRight(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(keyIndex), CMessage(password), C.SGD_UINT32(pwdLength))
	err = ToError(err1)
	return err
}

// SDFReleasePrivateKeyAccessRight 8.释放私钥使用权限
func (c *Ctx) SDFReleasePrivateKeyAccessRight(sessionHandle SessionHandle, keyIndex uint) (err error) {
	var err1 = C.SDFReleasePrivateKeyAccessRight(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(keyIndex))
	err = ToError(err1)
	return err
}

// SDFExportSignPublicKey_ECC 16.导出 ECC签名公钥
func (c *Ctx) SDFExportSignPublicKey_ECC(sessionHandle SessionHandle, uiKeyIndex uint) (publicKey ECCrefPublicKey, err error) {
	var pucPublicKey C.ECCrefPublicKey
	var err1 = C.SDFExportSignPublicKey_ECC(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiKeyIndex), &pucPublicKey)
	publicKey = ConvertToECCrefPublicKeyGo(pucPublicKey)
	err = ToError(err1)
	return publicKey, err
}

// SDFExportEncPublicKey_ECC 17.导出 ECC加密公钥
func (c *Ctx) SDFExportEncPublicKey_ECC(sessionHandle SessionHandle, uiKeyIndex uint) (publicKey ECCrefPublicKey, err error) {
	var err1 C.SGD_RV
	var pucPublicKey C.ECCrefPublicKey
	err1 = C.SDFExportEncPublicKey_ECC(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiKeyIndex), &pucPublicKey)
	publicKey = ConvertToECCrefPublicKeyGo(pucPublicKey)
	err = ToError(err1)
	return publicKey, err
}

// SDFGenerateKeyPair_ECC 18.产生 ECC非对称密钥对并输出
func (c *Ctx) SDFGenerateKeyPair_ECC(sessionHandle SessionHandle, uiAlgID uint, uiKeyBits uint) (publicKey ECCrefPublicKey, privateKey ECCrefPrivateKey, err error) {
	var err1 C.SGD_RV
	var pucPublicKey C.ECCrefPublicKey
	var pucPrivateKey C.ECCrefPrivateKey
	err1 = C.SDFGenerateKeyPair_ECC(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiAlgID), C.SGD_UINT32(uiKeyBits), &pucPublicKey, &pucPrivateKey)
	publicKey = ConvertToECCrefPublicKeyGo(pucPublicKey)
	privateKey = ConvertToECCrefPrivateKeyGo(pucPrivateKey)
	err = ToError(err1)
	return publicKey, privateKey, err
}

// SDFImportKey 28.导入明文会话密钥
func (c *Ctx) SDFImportKey(sessionHandle SessionHandle, pucKey []byte, uiKeyLength uint) (keyHandle SessionHandle, err error) {
	var err1 C.SGD_RV
	var phKeyHandle C.SGD_HANDLE
	err1 = C.SDFImportKey(C.SGD_HANDLE(sessionHandle), CMessage(pucKey), C.SGD_UINT32(uiKeyLength), &phKeyHandle)
	keyHandle = SessionHandle(phKeyHandle)
	err = ToError(err1)
	return keyHandle, err
}

// SDFDestroyKey 29.销毁会话密钥
func (c *Ctx) SDFDestroyKey(sessionHandle SessionHandle, hKeyHandle SessionHandle) (err error) {
	var err1 = C.SDFDestroyKey(C.SGD_HANDLE(sessionHandle), C.SGD_HANDLE(hKeyHandle))
	err = ToError(err1)
	return err
}

// SDFExternalSign_ECC 34. 外部密钥ECC签名
func (c *Ctx) SDFExternalSign_ECC(sessionHandle SessionHandle, uiAlgID uint, privateKey ECCrefPrivateKey, pucData []byte, uiDataLength uint) (signature ECCSignature, err error) {
	pucPrivateKey := ConvertToECCrefPrivateKeyC(privateKey)
	var pucSignature C.ECCSignature
	var err1 = C.SDFExternalSign_ECC(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiAlgID), &pucPrivateKey, CMessage(pucData), C.SGD_UINT32(uiDataLength), &pucSignature)
	signature = ConvertToECCSignatureGo(pucSignature)
	err = ToError(err1)
	return signature, err
}

// SDFExternalVerify_ECC 35.外部密钥 ECC验证
func (c *Ctx) SDFExternalVerify_ECC(sessionHandle SessionHandle, uiAlgID uint, publicKey ECCrefPublicKey, inputData []byte, uiInputLength uint, signature ECCSignature) (err error) {
	var err1 C.SGD_RV
	pucPublicKey := ConvertToECCrefPublicKeyC(publicKey)
	pucSignature := ConvertToECCSignatureC(signature)
	err1 = C.SDFExternalVerify_ECC(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiAlgID), &pucPublicKey, CMessage(inputData), C.SGD_UINT32(uiInputLength), &pucSignature)
	err = ToError(err1)
	return err
}

// SDFInternalSign_ECC 36.内部密钥 ECC签名
func (c *Ctx) SDFInternalSign_ECC(sessionHandle SessionHandle, uiISKIndex uint, pucData []byte, uiDataLength uint) (signature ECCSignature, err error) {
	var err1 C.SGD_RV
	var pucSignature C.ECCSignature
	err1 = C.SDFInternalSign_ECC(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiISKIndex), CMessage(pucData), C.SGD_UINT32(uiDataLength), &pucSignature)
	signature = ConvertToECCSignatureGo(pucSignature)
	err = ToError(err1)
	return signature, err
}

// SDFInternalVerify_ECC 37.内部密钥 ECC验证
func (c *Ctx) SDFInternalVerify_ECC(sessionHandle SessionHandle, uiISKIndex uint, pucData []byte, uiDataLength uint, signature ECCSignature) (err error) {
	var err1 C.SGD_RV
	var pucSignature C.ECCSignature
	for i := 0; i < len(signature.R); i++ {
		pucSignature.r[i] = C.SGD_UCHAR(signature.R[i])
	}
	for i := 0; i < len(signature.S); i++ {
		pucSignature.s[i] = C.SGD_UCHAR(signature.S[i])
	}
	err1 = C.SDFInternalVerify_ECC(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiISKIndex), CMessage(pucData), C.SGD_UINT32(uiDataLength), &pucSignature)
	err = ToError(err1)
	return err
}

// SDFExternalEncrypt_ECC 38.外部密钥 ECC加密
func (c *Ctx) SDFExternalEncrypt_ECC(sessionHandle SessionHandle, uiAlgID uint, publicKey ECCrefPublicKey, data []byte, dataLength uint) (encData ECCCipher, err error) {
	var err1 C.SGD_RV
	pucPublicKey := ConvertToECCrefPublicKeyC(publicKey)
	var pucEncData C.ECCCipher
	err1 = C.SDFExternalEncrypt_ECC(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiAlgID), &pucPublicKey, CMessage(data), C.SGD_UINT32(dataLength), &pucEncData)
	encData = ConvertToECCCipherGo(pucEncData)
	err = ToError(err1)
	return encData, err
}

// SDFExternalDecrypt_ECC 39.外部密钥 ECC解密
func (c *Ctx) SDFExternalDecrypt_ECC(sessionHandle SessionHandle, uiAlgID uint, privateKey ECCrefPrivateKey, encData ECCCipher) (data []byte, dataLength uint, err error) {
	var err1 C.SGD_RV
	pucPrivateKey := ConvertToECCrefPrivateKeyC(privateKey)
	pucEncData := ConvertToECCCipherC(encData)
	var pucData C.SGD_UCHAR_PRT
	var puiDataLength C.SGD_UINT32
	err1 = C.SDFExternalDecrypt_ECC(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiAlgID), &pucPrivateKey, &pucEncData, &pucData, &puiDataLength)
	data = C.GoBytes(unsafe.Pointer(pucData), C.int(puiDataLength))
	dataLength = uint(puiDataLength)
	C.free(unsafe.Pointer(pucData))
	err = ToError(err1)
	return data, dataLength, err
}

// SDFEncrypt 40.对称加密
func (c *Ctx) SDFEncrypt(sessionHandle SessionHandle, keyHandle SessionHandle, algID uint, iv []byte, data []byte, dataLength uint) (encData []byte, encDataLength uint, err error) {
	var pucEncData C.SGD_UCHAR_PRT
	var puiEncDataLength C.SGD_UINT32
	var err1 = C.SDFEncrypt(C.SGD_HANDLE(sessionHandle), C.SGD_HANDLE(keyHandle), C.SGD_UINT32(algID), CMessage(iv), CMessage(data), C.SGD_UINT32(dataLength), &pucEncData, &puiEncDataLength)
	encData = C.GoBytes(unsafe.Pointer(pucEncData), C.int(puiEncDataLength))
	err = ToError(err1)
	C.free(unsafe.Pointer(pucEncData))
	return encData, uint(puiEncDataLength), err
}

// SDFDecrypt 41.对称解密
func (c *Ctx) SDFDecrypt(sessionHandle SessionHandle, hKeyHandle SessionHandle, uiAlgID uint, iv []byte, encData []byte, encDataLength uint) (data []byte, dataLength uint, err error) {
	var pucData C.SGD_UCHAR_PRT
	var puiDataLength C.SGD_UINT32
	var err1 = C.SDFDecrypt(C.SGD_HANDLE(sessionHandle), C.SGD_HANDLE(hKeyHandle), C.SGD_UINT32(uiAlgID), CMessage(iv), CMessage(encData), C.SGD_UINT32(encDataLength), &pucData, &puiDataLength)
	data = C.GoBytes(unsafe.Pointer(pucData), C.int(puiDataLength))
	dataLength = uint(puiDataLength)
	C.free(unsafe.Pointer(pucData))
	err = ToError(err1)
	return data, dataLength, err
}

// SDFGetSymmKeyHandle 50.
func (c *Ctx) SDFGetSymmKeyHandle(sessionHandle SessionHandle, uiKeyIndex uint) (keyHandle SessionHandle, err error) {
	var phKeyHandle C.SGD_HANDLE
	var err1 = C.SDFGetSymmKeyHandle(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiKeyIndex), &phKeyHandle)
	keyHandle = SessionHandle(phKeyHandle)
	err = ToError(err1)
	return keyHandle, err
}

// SDFInternalEncrypt_ECC 51. ECC方式的加密
func (c *Ctx) SDFInternalEncrypt_ECC(sessionHandle SessionHandle, uiISKIndex uint, uiAlgID uint, pucData []byte, uiDataLength uint) (encData ECCCipher, err error) {
	var pucEncData C.ECCCipher
	var err1 = C.SDFInternalEncrypt_ECC(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiISKIndex), C.SGD_UINT32(uiAlgID), CMessage(pucData), C.SGD_UINT32(uiDataLength), &pucEncData)
	encData = ConvertToECCCipherGo(pucEncData)
	err = ToError(err1)
	return encData, err
}

// SDFInternalDecrypt_ECC 52. ECC方式的解密
func (c *Ctx) SDFInternalDecrypt_ECC(sessionHandle SessionHandle, uiISKIndex uint, uiAlgID uint, encData ECCCipher) (data []byte, dataLength uint, err error) {
	var err1 C.SGD_RV
	var pucEncData C.ECCCipher
	for i := 0; i < len(encData.X); i++ {
		pucEncData.x[i] = C.SGD_UCHAR(encData.X[i])
	}
	for i := 0; i < len(encData.Y); i++ {
		pucEncData.y[i] = C.SGD_UCHAR(encData.Y[i])
	}
	for i := 0; i < len(encData.M); i++ {
		pucEncData.M[i] = C.SGD_UCHAR(encData.M[i])
	}
	pucEncData.L = C.SGD_UINT32(encData.L)
	for i := 0; i < len(encData.C); i++ {
		pucEncData.C[i] = C.SGD_UCHAR(encData.C[i])
	}
	var pucData C.SGD_UCHAR_PRT
	var puiDataLength C.SGD_UINT32
	err1 = C.SDFInternalDecrypt_ECC(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiISKIndex), C.SGD_UINT32(uiAlgID), &pucEncData, &pucData, &puiDataLength)
	data = C.GoBytes(unsafe.Pointer(pucData), C.int(puiDataLength))
	C.free(unsafe.Pointer(pucData))
	dataLength = uint(puiDataLength)
	err = ToError(err1)
	return data, dataLength, err
}

// SDFHashInit 43
func (c *Ctx) SDFHashInit(sessionHandle SessionHandle, uiAlgID uint, pucID []byte, uiIDLength uint) (publicKey ECCrefPublicKey, err error) {
	var pucPublicKey C.ECCrefPublicKey
	var err1 = C.SDFHashInit(C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiAlgID), nil, nil, C.SGD_UINT32(0))
	publicKey = ConvertToECCrefPublicKeyGo(pucPublicKey)
	return publicKey, ToError(err1)
}

// SDFHashUpdate 44
func (c *Ctx) SDFHashUpdate(sessionHandle SessionHandle, pucData []byte, uiDataLength uint) (err error) {
	var err1 = C.SDFHashUpdate(C.SGD_HANDLE(sessionHandle), CMessage(pucData), C.SGD_UINT32(uiDataLength))
	return ToError(err1)
}

// SDFHashFinal 45
func (c *Ctx) SDFHashFinal(sessionHandle SessionHandle) (hash []byte, hashLength uint, err error) {
	var pucData C.SGD_UCHAR_PRT
	var pubHashLength C.SGD_UINT32
	var err1 = C.SDFHashFinal(C.SGD_HANDLE(sessionHandle), &pucData, &pubHashLength)
	hash = C.GoBytes(unsafe.Pointer(pucData), C.int(pubHashLength))
	hashLength = uint(pubHashLength)
	C.free(unsafe.Pointer(pucData))
	return hash, hashLength, ToError(err1)
}

func (c *Ctx) SDFCalculateMAC(sessionHandle SessionHandle, hKeyHandle SessionHandle, uiAlgID uint, iv []byte, data []byte, dataLength uint) (mac []byte, macLength uint, err error) {
	var pucMAC C.SGD_UCHAR_PRT
	var puiMACLength C.SGD_UINT32
	err1 := C.SDFCalculateMAC(C.SGD_HANDLE(sessionHandle), C.SGD_HANDLE(hKeyHandle), C.SGD_UINT32(uiAlgID), CMessage(iv), CMessage(data), C.SGD_UINT32(dataLength), &pucMAC, &puiMACLength)
	mac = C.GoBytes(unsafe.Pointer(pucMAC), C.int(puiMACLength))
	macLength = uint(puiMACLength)
	C.free(unsafe.Pointer(pucMAC))
	return mac, macLength, ToError(err1)
}

// HMAC SDFHMAC used to calculate HMAC-SM3
func (c *Ctx) SDFHMAC(sessionHandle SessionHandle, hKeyHandle SessionHandle, uiAlgID uint, data []byte, dataLength uint) (mac []byte, macLength uint, err error) {
	var pucMAC C.SGD_UCHAR_PRT
	var puiMACLength C.SGD_UINT32
	err1 := C.SDFHMAC(C.SGD_HANDLE(sessionHandle), C.SGD_HANDLE(hKeyHandle), C.SGD_UINT32(uiAlgID), CMessage(data), C.SGD_UINT32(dataLength), &pucMAC, &puiMACLength)
	mac = C.GoBytes(unsafe.Pointer(pucMAC), C.int(puiMACLength))
	macLength = uint(puiMACLength)
	C.free(unsafe.Pointer(pucMAC))
	return mac, macLength, ToError(err1)
}
