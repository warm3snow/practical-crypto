package swsds

import (
	"bytes"
	"log"
)

func Test_Encrypt_Decrypt() {
	log.Println("Test_Encrypt_Decrypt")

	var err error
	handler, err := SDF_OpenDevice()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(handler)

	session, err := SDF_OpenSession(handler)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(session)

	d, err := SDF_GenerateRandom(session, 16)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(d)

	//keyHandle, err := SDF_ImportKey(session, d)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	keyHandle, err := SDF_GetSymmKeyHandle(session, 1)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(keyHandle)

	// 加密
	inData, err := SDF_GenerateRandom(session, 32)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(inData)
	encData, err := SDF_Encrypt(session, keyHandle, SGD_SMS4_ECB, nil, inData)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(encData)

	// 解密
	decData, err := SDF_Decrypt(session, keyHandle, SGD_SMS4_ECB, nil, encData)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(decData)

	// 对比结果
	if bytes.Equal(inData, decData) {
		log.Println("运算结果：加密、解密及结果比较均正确。")
	} else {
		log.Println("运算结果：解密结果错误。")
	}

	err = SDF_DestroyKey(session, keyHandle)
	if err != nil {
		log.Println(err)
		return
	}

	err = SDF_CloseSession(session)
	if err != nil {
		log.Println(err)
		return
	}

	err = SDF_CloseDevice(handler)
	if err != nil {
		log.Println(err)
		return
	}
}
