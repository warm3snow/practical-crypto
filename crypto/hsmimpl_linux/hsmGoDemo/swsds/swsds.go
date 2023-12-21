package swsds

/*
  #cgo CFLAGS: -I./
  #cgo linux LDFLAGS: -L./ -lswsds
  #if defined(__linux__) || defined(linux)
  	#include "swsds.h"
	#include <pthread.h>
  #endif
*/
import "C"

func SDF_OpenDevice() (handler CTypeSGDHandle, err error) {
	defer RecoverError(&err)

	var sgdHandle C.SGD_HANDLE
	result := C.SDF_OpenDevice(&sgdHandle)

	err = ConvertInt64toError(int64(result))
	handler = CTypeSGDHandle(sgdHandle)
	return
}

//SDF_CloseDevice 关闭设备
func SDF_CloseDevice(handler CTypeSGDHandle) (err error) {
	defer RecoverError(&err)

	result := C.SDF_CloseDevice(C.SGD_HANDLE(handler))
	err = ConvertInt64toError(int64(result))
	return
}

//SDF_OpenSession 打开session
func SDF_OpenSession(handler CTypeSGDHandle) (session CTypeSGDHandle, err error) {
	defer RecoverError(&err)

	var sgdSession C.SGD_HANDLE
	result := C.SDF_OpenSession(C.SGD_HANDLE(handler), &sgdSession)

	err = ConvertInt64toError(int64(result))
	session = CTypeSGDHandle(sgdSession)
	return
}

//SDF_CloseSession 关闭session
func SDF_CloseSession(session CTypeSGDHandle) (err error) {
	defer RecoverError(&err)

	result := C.SDF_CloseSession(C.SGD_HANDLE(session))
	err = ConvertInt64toError(int64(result))
	return
}

func SDF_GenerateRandom(session CTypeSGDHandle, len int) (data []byte, err error) {
	defer RecoverError(&err)

	charArr := make([]C.SGD_UCHAR, len)
	result := C.SDF_GenerateRandom(C.SGD_HANDLE(session), C.SGD_UINT32(len), &charArr[0])
	data = SgdUCHARArrToByteArr(charArr)
	err = ConvertInt64toError(int64(result))
	return
}

//SGD_RV SDF_GetSymmKeyHandle(SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SGD_HANDLE *phKeyHandle);
func SDF_GetSymmKeyHandle(session CTypeSGDHandle, index int) (hKeyHandle CTypeSGDHandle, err error) {
	defer RecoverError(&err)

	var handle C.SGD_HANDLE
	result := C.SDF_GetSymmKeyHandle(C.SGD_HANDLE(session), C.SGD_UINT32(index), &handle)
	hKeyHandle = CTypeSGDHandle(handle)
	err = ConvertInt64toError(int64(result))
	return

}

// SGD_RV SDF_ImportKey (SGD_HANDLE hSessionHandle, SGD_UCHAR *pucKey, SGD_UINT32 uiKeyLength,SGD_HANDLE *phKeyHandle);
func SDF_ImportKey(session CTypeSGDHandle, key []byte) (hKeyHandle CTypeSGDHandle, err error) {
	defer RecoverError(&err)

	pucKey := ByteArrToSgdUCHARArr(key)
	uiKeyLength := C.SGD_UINT32(len(key))
	var handle C.SGD_HANDLE
	result := C.SDF_ImportKey(C.SGD_HANDLE(session), &pucKey[0], uiKeyLength, &handle)
	hKeyHandle = CTypeSGDHandle(handle)
	err = ConvertInt64toError(int64(result))
	return

}

// SGD_RV SDF_DestroyKey (SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle);
func SDF_DestroyKey(session CTypeSGDHandle, hKeyHandle CTypeSGDHandle) (err error) {
	defer RecoverError(&err)

	result := C.SDF_DestroyKey(C.SGD_HANDLE(session), C.SGD_HANDLE(hKeyHandle))
	err = ConvertInt64toError(int64(result))
	return
}

// SGD_RV SDF_Encrypt(SGD_HANDLE hSessionHandle,SGD_HANDLE hKeyHandle,SGD_UINT32 uiAlgID,SGD_UCHAR *pucIV,SGD_UCHAR *pucData,SGD_UINT32 uiDataLength,SGD_UCHAR *pucEncData,SGD_UINT32  *puiEncDataLength);
func SDF_Encrypt(session CTypeSGDHandle, hKeyHandle CTypeSGDHandle, alg CTypeAlgorithm, pucIV []byte, pucData []byte) (pucEncData []byte, err error) {
	defer RecoverError(&err)

	pucIVs := ByteArrToSgdUCHARArr(pucIV)
	pucDatas := ByteArrToSgdUCHARArr(pucData)
	uiDataLength := C.SGD_UINT32(len(pucData))
	pucEncDatas := make([]C.SGD_UCHAR, len(pucData))
	var puiEncDataLength C.SGD_UINT32
	var result C.SGD_RV
	if pucIVs == nil {
		result = C.SDF_Encrypt(C.SGD_HANDLE(session), C.SGD_HANDLE(hKeyHandle), C.SGD_UINT32(alg), nil, &pucDatas[0], uiDataLength, &pucEncDatas[0], &puiEncDataLength)
	} else {
		result = C.SDF_Encrypt(C.SGD_HANDLE(session), C.SGD_HANDLE(hKeyHandle), C.SGD_UINT32(alg), &pucIVs[0], &pucDatas[0], uiDataLength, &pucEncDatas[0], &puiEncDataLength)
	}
	pucEncData = SgdUCHARArrToByteArr(pucEncDatas)
	err = ConvertInt64toError(int64(result))
	return
}

// SGD_RV SDF_Decrypt (SGD_HANDLE hSessionHandle,SGD_HANDLE hKeyHandle,SGD_UINT32 uiAlgID,SGD_UCHAR *pucIV,SGD_UCHAR *pucEncData,SGD_UINT32  uiEncDataLength,SGD_UCHAR *pucData,SGD_UINT32 *puiDataLength);
func SDF_Decrypt(session CTypeSGDHandle, hKeyHandle CTypeSGDHandle, alg CTypeAlgorithm, pucIV []byte, pucEncData []byte) (pucData []byte, err error) {
	defer RecoverError(&err)

	pucIVs := ByteArrToSgdUCHARArr(pucIV)
	pucEncDatas := ByteArrToSgdUCHARArr(pucEncData)
	uiEncDataLength := C.SGD_UINT32(len(pucEncData))
	pucDatas := make([]C.SGD_UCHAR, len(pucEncData))
	var puiDataLength C.SGD_UINT32
	var result C.SGD_RV
	if pucIVs == nil {
		result = C.SDF_Decrypt(C.SGD_HANDLE(session), C.SGD_HANDLE(hKeyHandle), C.SGD_UINT32(alg), nil, &pucEncDatas[0], uiEncDataLength, &pucDatas[0], &puiDataLength)
	} else {
		result = C.SDF_Decrypt(C.SGD_HANDLE(session), C.SGD_HANDLE(hKeyHandle), C.SGD_UINT32(alg), &pucIVs[0], &pucEncDatas[0], uiEncDataLength, &pucDatas[0], &puiDataLength)
	}
	pucData = SgdUCHARArrToByteArr(pucDatas)
	err = ConvertInt64toError(int64(result))
	return
}
