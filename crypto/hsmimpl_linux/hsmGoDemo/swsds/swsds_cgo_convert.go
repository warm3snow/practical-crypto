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
import "unsafe"

// --- c 语言转换go
func ConvertSgdUCharPtrToString(uChars *C.SGD_UCHAR, len int) string {
	if uChars == nil {
		return ""
	}
	bytes := ConvertSgdUCharPtrToBytes(uChars, len)
	return string(bytes)
}
func ConvertSgdUCharPtrToBytes(uChars *C.SGD_UCHAR, len int) []byte {
	var bytes []byte
	if uChars == nil {
		return bytes
	}
	uCharSlice := (*[MAX_INT]C.uchar)(unsafe.Pointer(uChars))[:len]
	for _, b := range uCharSlice {
		bytes = append(bytes, byte(b))
	}
	return bytes
}

func UcharArrToByteArr(buf []C.uchar) []byte {
	var ret []byte
	if buf == nil {
		return nil
	}
	for i := 0; i < len(buf); i++ {
		ret = append(ret, byte(buf[i]))
	}
	return ret
}
func ByteArrToUcharArr(buf []byte) []C.uchar {
	var ret []C.uchar
	if buf == nil {
		return nil
	}
	for i := 0; i < 32; i++ {
		if i < len(buf) {
			ret = append(ret, C.uchar(buf[i]))
		} else {
			ret = append(ret, C.uchar([]byte("0")[0]))
		}
	}
	return ret
}
func ByteArrToSgdUCHARArr32(buf []byte) []C.SGD_UCHAR {
	var ret []C.SGD_UCHAR
	if buf == nil {
		return nil
	}
	for i := 0; i < 32; i++ {
		if i < len(buf) {
			ret = append(ret, C.SGD_UCHAR(buf[i]))
		} else {
			ret = append(ret, C.SGD_UCHAR([]byte("0")[0]))
		}
	}
	return ret
}

func ByteArrToSgdUCHARArr(buf []byte) []C.SGD_UCHAR {
	var ret []C.SGD_UCHAR
	if buf == nil {
		return nil
	}
	for i := 0; i < len(buf); i++ {
		if i < len(buf) {
			ret = append(ret, C.SGD_UCHAR(buf[i]))
		} else {
			ret = append(ret, C.SGD_UCHAR([]byte("0")[0]))
		}
	}
	return ret
}

func SgdUCHARArrToByteArr(buf []C.SGD_UCHAR) []byte {
	var ret []byte
	if buf == nil {
		return nil
	}
	for i := 0; i < len(buf); i++ {
		ret = append(ret, byte(buf[i]))
	}
	return ret
}
