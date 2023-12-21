package swsds

import (
	"fmt"
	"strings"
)

func init() {
	signErrorCodeMapInit()
}

func ConvertInt64toError(result int64) error {
	if result == 0 {
		return nil
	}

	errMsg, ok := SignErrorCodeMap[result]
	if !ok {
		return fmt.Errorf("[code:0x%X,msg:%s]", result, "未知错误码")
	}

	return fmt.Errorf("[code:0x%X,msg:%s]", result, errMsg)
}

func RecoverError(errPtr *error) {
	if err := recover(); err != nil {
		var err2 error

		if err3, ok := err.(error); !ok {
			err2 = err3
		} else {
			err2 = fmt.Errorf("Unknown error, %v ", err)
		}

		// 包装栈信息
		if errPtr == nil {
			errPtr = &err2
			return
		}

		// 合并两个error，通过字符串拼接
		var errStrings []string
		errStrings = append(errStrings, (*errPtr).Error())
		errStrings = append(errStrings, err2.Error())
		err2 = fmt.Errorf(strings.Join(errStrings, "\n"))
		errPtr = &err2
	}
}

const (
	/*标准错误码定义*/
	SDR_OK               int64 = 0x0 /*成功*/
	SDR_BASE             int64 = 0x01000000
	SDR_UNKNOWERR        int64 = (SDR_BASE + 0x00000001) /*未知错误*/
	SDR_NOTSUPPORT       int64 = (SDR_BASE + 0x00000002) /*不支持*/
	SDR_COMMFAIL         int64 = (SDR_BASE + 0x00000003) /*通信错误*/
	SDR_HARDFAIL         int64 = (SDR_BASE + 0x00000004) /*硬件错误*/
	SDR_OPENDEVICE       int64 = (SDR_BASE + 0x00000005) /*打开设备错误*/
	SDR_OPENSESSION      int64 = (SDR_BASE + 0x00000006) /*打开会话句柄错误*/
	SDR_PARDENY          int64 = (SDR_BASE + 0x00000007) /*权限不满足*/
	SDR_KEYNOTEXIST      int64 = (SDR_BASE + 0x00000008) /*密钥不存在*/
	SDR_ALGNOTSUPPORT    int64 = (SDR_BASE + 0x00000009) /*不支持的算法*/
	SDR_ALGMODNOTSUPPORT int64 = (SDR_BASE + 0x0000000A) /*不支持的算法模式*/
	SDR_PKOPERR          int64 = (SDR_BASE + 0x0000000B) /*公钥运算错误*/
	SDR_SKOPERR          int64 = (SDR_BASE + 0x0000000C) /*私钥运算错误*/
	SDR_SIGNERR          int64 = (SDR_BASE + 0x0000000D) /*签名错误*/
	SDR_VERIFYERR        int64 = (SDR_BASE + 0x0000000E) /*验证错误*/
	SDR_SYMOPERR         int64 = (SDR_BASE + 0x0000000F) /*对称运算错误*/
	SDR_STEPERR          int64 = (SDR_BASE + 0x00000010) /*步骤错误*/
	SDR_FILESIZEERR      int64 = (SDR_BASE + 0x00000011) /*文件大小错误*/
	SDR_FILENOEXIST      int64 = (SDR_BASE + 0x00000012) /*文件不存在*/
	SDR_FILEOFSERR       int64 = (SDR_BASE + 0x00000013) /*文件操作偏移量错误*/
	SDR_KEYTYPEERR       int64 = (SDR_BASE + 0x00000014) /*密钥类型错误*/
	SDR_KEYERR           int64 = (SDR_BASE + 0x00000015) /*密钥错误*/

	SDR_ENCDATAERR int64 = (SDR_BASE + 0x00000016) /*ECC加密数据错误*/
	SDR_RANDERR    int64 = (SDR_BASE + 0x00000017) /*随机数产生失败*/
	SDR_PRKRERR    int64 = (SDR_BASE + 0x00000018) /*私钥使用权限获取失败，未获取 | 私钥访问控制权限释放失败，未获取*/
	SDR_MACERR     int64 = (SDR_BASE + 0x00000019) /*MAC运算失败*/
	SDR_FILEEXISTS int64 = (SDR_BASE + 0x0000001A) /*指定文件已存在*/
	SDR_FILEWERR   int64 = (SDR_BASE + 0x0000001B) /*文件写入失败*/
	SDR_NOBUFFER   int64 = (SDR_BASE + 0x0000001C) /*存储空间不足*/
	SDR_INARGERR   int64 = (SDR_BASE + 0x0000001D) /*输入参数错误：1、数据长度错误*/
	SDR_OUTARGERR  int64 = (SDR_BASE + 0x0000001E) /*输出参数错误*/

	/*扩展错误码*/
	SWR_BASE               int64 = (SDR_BASE + 0x00010000) /*自定义错误码基础值*/
	SWR_INVALID_USER       int64 = (SWR_BASE + 0x00000001) /*无效的用户名*/
	SWR_INVALID_AUTHENCODE int64 = (SWR_BASE + 0x00000002) /*无效的授权码*/
	SWR_PROTOCOL_VER_ERR   int64 = (SWR_BASE + 0x00000003) /*不支持的协议版本*/
	SWR_INVALID_COMMAND    int64 = (SWR_BASE + 0x00000004) /*错误的命令字*/
	SWR_INVALID_PACKAGE    int64 = (SWR_BASE + 0x00000005) /*错误的数据包格式*/
	SWR_FILE_ALREADY_EXIST int64 = (SWR_BASE + 0x00000006) /*已存在同名文件*/

	SWR_SOCKET_TIMEOUT  int64 = (SWR_BASE + 0x00000100) /*超时错误*/
	SWR_CONNECT_ERR     int64 = (SWR_BASE + 0x00000101) /*连接服务器错误*/
	SWR_SET_SOCKOPT_ERR int64 = (SWR_BASE + 0x00000102) /*设置Socket参数错误*/
	SWR_SOCKET_SEND_ERR int64 = (SWR_BASE + 0x00000104) /*发送LOGINRequest错误*/
	SWR_SOCKET_RECV_ERR int64 = (SWR_BASE + 0x00000105) /*发送LOGINRequest错误*/
	SWR_SOCKET_RECV_0   int64 = (SWR_BASE + 0x00000106) /*发送LOGINRequest错误*/

	SWR_NO_AVAILABLE_HSM int64 = (SWR_BASE + 0x00000201) /*没有可用的加密机*/
	SWR_NO_AVAILABLE_CSM int64 = (SWR_BASE + 0x00000202) /*加密机内没有可用的加密模块*/
	SWR_CONFIG_ERR       int64 = (SWR_BASE + 0x00000301) /*配置文件错误*/

	SWR_CARD_BASE        int64 = (SDR_BASE + 0x00020000)      /*密码卡错误码*/
	SDR_BUFFER_TOO_SMALL int64 = (SWR_CARD_BASE + 0x00000101) /*接收参数的缓存区太小*/
	SDR_DATA_PAD         int64 = (SWR_CARD_BASE + 0x00000102) /*数据没有按正确格式填充，或解密得到的脱密数据不符合填充格式*/
	SDR_DATA_SIZE        int64 = (SWR_CARD_BASE + 0x00000103) /*明文或密文长度不符合相应的算法要求*/
	SDR_CRYPTO_NOT_INIT  int64 = (SWR_CARD_BASE + 0x00000104) /*步骤错误*/

	SWR_MANAGEMENT_DENY   int64 = (SWR_CARD_BASE + 0x00001001) //管理权限不满足
	SWR_OPERATION_DENY    int64 = (SWR_CARD_BASE + 0x00001002) //操作权限不满足
	SWR_DEVICE_STATUS_ERR int64 = (SWR_CARD_BASE + 0x00001003) //当前设备状态不满足现有操作

	SWR_LOGIN_ERR    int64 = (SWR_CARD_BASE + 0x00001011) //登录失败
	SWR_USERID_ERR   int64 = (SWR_CARD_BASE + 0x00001012) //用户ID数目/号码错误
	SWR_PARAMENT_ERR int64 = (SWR_CARD_BASE + 0x00001013) //参数错误
	SWR_KEYTYPEERR   int64 = (SWR_CARD_BASE + 0x00000020) //密钥类型错误
)

var SignErrorCodeMap = make(map[int64]string)

func signErrorCodeMapInit() {
	SignErrorCodeMap[SDR_UNKNOWERR] = "未知错误"
	SignErrorCodeMap[SDR_NOTSUPPORT] = "不支持"
	SignErrorCodeMap[SDR_COMMFAIL] = "通信错误"
	SignErrorCodeMap[SDR_HARDFAIL] = "硬件错误"
	SignErrorCodeMap[SDR_OPENDEVICE] = "打开设备错误"
	SignErrorCodeMap[SDR_OPENSESSION] = "打开会话句柄错误"
	SignErrorCodeMap[SDR_PARDENY] = "权限不满足"
	SignErrorCodeMap[SDR_KEYNOTEXIST] = "密钥不存在"
	SignErrorCodeMap[SDR_ALGNOTSUPPORT] = "不支持的算法"
	SignErrorCodeMap[SDR_ALGMODNOTSUPPORT] = "不支持的算法模式"
	SignErrorCodeMap[SDR_PKOPERR] = "公钥运算错误"
	SignErrorCodeMap[SDR_SKOPERR] = "私钥运算错误"
	SignErrorCodeMap[SDR_SIGNERR] = "签名错误"
	SignErrorCodeMap[SDR_VERIFYERR] = "验证错误"
	SignErrorCodeMap[SDR_SYMOPERR] = "对称运算错误"
	SignErrorCodeMap[SDR_STEPERR] = "步骤错误"
	SignErrorCodeMap[SDR_FILESIZEERR] = "文件大小错误"
	SignErrorCodeMap[SDR_FILENOEXIST] = "文件不存在"
	SignErrorCodeMap[SDR_FILEOFSERR] = "文件操作偏移量错误"
	SignErrorCodeMap[SDR_KEYTYPEERR] = "密钥类型错误"
	SignErrorCodeMap[SDR_KEYERR] = "密钥错误"

	SignErrorCodeMap[SDR_ENCDATAERR] = "ECC加密数据错误"
	SignErrorCodeMap[SDR_RANDERR] = "随机数产生失败"
	SignErrorCodeMap[SDR_PRKRERR] = "私钥使用权限获取失败，未获取 | 私钥访问控制权限释放失败，未获取"
	SignErrorCodeMap[SDR_MACERR] = "MAC运算失败"
	SignErrorCodeMap[SDR_FILEEXISTS] = "指定文件已存在"
	SignErrorCodeMap[SDR_FILEWERR] = "文件写入失败"
	SignErrorCodeMap[SDR_NOBUFFER] = "存储空间不足"
	SignErrorCodeMap[SDR_INARGERR] = "输入参数错误：1、数据长度错误"
	SignErrorCodeMap[SDR_OUTARGERR] = "输出参数错误"

	/*扩展错误码*/
	SignErrorCodeMap[SWR_BASE] = "自定义错误码基础值"
	SignErrorCodeMap[SWR_INVALID_USER] = "无效的用户名"
	SignErrorCodeMap[SWR_INVALID_AUTHENCODE] = "无效的授权码"
	SignErrorCodeMap[SWR_PROTOCOL_VER_ERR] = "不支持的协议版本"
	SignErrorCodeMap[SWR_INVALID_COMMAND] = "错误的命令字"
	SignErrorCodeMap[SWR_INVALID_PACKAGE] = "错误的数据包格式"
	SignErrorCodeMap[SWR_FILE_ALREADY_EXIST] = "已存在同名文件"

	SignErrorCodeMap[SWR_SOCKET_TIMEOUT] = "超时错误"
	SignErrorCodeMap[SWR_CONNECT_ERR] = "连接服务器错误"
	SignErrorCodeMap[SWR_SET_SOCKOPT_ERR] = "设置Socket参数错误"
	SignErrorCodeMap[SWR_SOCKET_SEND_ERR] = "发送LOGINRequest错误"
	SignErrorCodeMap[SWR_SOCKET_RECV_ERR] = "发送LOGINRequest错误"
	SignErrorCodeMap[SWR_SOCKET_RECV_0] = "发送LOGINRequest错误"

	SignErrorCodeMap[SWR_NO_AVAILABLE_HSM] = "没有可用的加密机"
	SignErrorCodeMap[SWR_NO_AVAILABLE_CSM] = "加密机内没有可用的加密模块"
	SignErrorCodeMap[SWR_CONFIG_ERR] = "配置文件错误"

	SignErrorCodeMap[SWR_CARD_BASE] = "密码卡错误码"
	SignErrorCodeMap[SDR_BUFFER_TOO_SMALL] = "接收参数的缓存区太小"
	SignErrorCodeMap[SDR_DATA_PAD] = "数据没有按正确格式填充，或解密得到的脱密数据不符合填充格式"
	SignErrorCodeMap[SDR_DATA_SIZE] = "明文或密文长度不符合相应的算法要求"
	SignErrorCodeMap[SDR_CRYPTO_NOT_INIT] = "步骤错误"

	SignErrorCodeMap[SWR_MANAGEMENT_DENY] = "管理权限不满足"
	SignErrorCodeMap[SWR_OPERATION_DENY] = "操作权限不满足"
	SignErrorCodeMap[SWR_DEVICE_STATUS_ERR] = "当前设备状态不满足现有操作"

	SignErrorCodeMap[SWR_LOGIN_ERR] = "登录失败"
	SignErrorCodeMap[SWR_USERID_ERR] = "用户ID数目/号码错误"
	SignErrorCodeMap[SWR_PARAMENT_ERR] = "参数错误"
	SignErrorCodeMap[SWR_KEYTYPEERR] = "密钥类型错误"
}
