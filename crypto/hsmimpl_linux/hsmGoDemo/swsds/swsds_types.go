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

// --------- 加签机方法封装  -----------
type CTypeSGDHandle C.SGD_HANDLE
type CTypeKeyUsage C.SGD_UINT32
type CTypeAlgorithm C.SGD_UINT32
type CTypeSGDRV C.SGD_RV

// DeviceInfo 设备信息
type DeviceInfo struct {
	IssuerName      string
	DeviceName      string
	DeviceSerial    string
	DeviceVersion   int
	StandardVersion int
	AsymAlgAbility  []int
	SymAlgAbility   int
	HashAlgAbility  int
	BufferSize      int
}
