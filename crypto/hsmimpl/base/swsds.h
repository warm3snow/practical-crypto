/*
* File: swsds.h
* Copyright (c) SWXA 2009
*
*/

#ifndef _SW_SDS_H_
#define _SW_SDS_H_ 1

#ifdef __cplusplus
	extern "C"{
#endif

/*RSA最大模长定义*/
#define SGD_RSA_MAX_BITS    4096

/*数据类型定义*/
typedef char				SGD_CHAR;
typedef char				SGD_INT8;
typedef short				SGD_INT16;
typedef int					SGD_INT32;
typedef long long			SGD_INT64;
typedef unsigned char		SGD_UCHAR;
typedef unsigned char		SGD_UINT8;
typedef unsigned short		SGD_UINT16;
typedef unsigned int		SGD_UINT32;
typedef unsigned long long	SGD_UINT64;
typedef unsigned int		SGD_RV;
typedef void*				SGD_OBJ;
typedef int					SGD_BOOL;
typedef void*				SGD_HANDLE;

/*设备信息*/
typedef struct DeviceInfo_st{
	unsigned char IssuerName[40];
	unsigned char DeviceName[16];
	unsigned char DeviceSerial[16];
	unsigned int  DeviceVersion;
	unsigned int  StandardVersion;
	unsigned int  AsymAlgAbility[2];
	unsigned int  SymAlgAbility;
	unsigned int  HashAlgAbility;
	unsigned int  BufferSize;
}DEVICEINFO;

/*设备运行信息--自定义扩展*/
typedef struct st_DeviceRunStatus{
	unsigned int onboot;		//服务是否开机自启动
	unsigned int service;		//当前服务状态，0-未启动，1-已启动，>1状态异常
	unsigned int concurrency;	//当前并发数
	unsigned int memtotal;		//内存大小
	unsigned int memfree;		//内存空闲
	unsigned int cpu;			//CPU占用率，不包含小数点部分
	unsigned int reserve1;
	unsigned int reserve2;
}DEVICE_RUN_STATUS;

/*RSA密钥*/
#define LiteRSAref_MAX_BITS    2048
#define LiteRSAref_MAX_LEN     ((LiteRSAref_MAX_BITS + 7) / 8)
#define LiteRSAref_MAX_PBITS   ((LiteRSAref_MAX_BITS + 1) / 2)
#define LiteRSAref_MAX_PLEN    ((LiteRSAref_MAX_PBITS + 7)/ 8)

typedef struct RSArefPublicKeyLite_st
{
	unsigned int  bits;
	unsigned char m[LiteRSAref_MAX_LEN];
	unsigned char e[LiteRSAref_MAX_LEN];
}RSArefPublicKeyLite;

typedef struct RSArefPrivateKeyLite_st
{
	unsigned int  bits;
	unsigned char m[LiteRSAref_MAX_LEN];
	unsigned char e[LiteRSAref_MAX_LEN];
	unsigned char d[LiteRSAref_MAX_LEN];
	unsigned char prime[2][LiteRSAref_MAX_PLEN];
	unsigned char pexp[2][LiteRSAref_MAX_PLEN];
	unsigned char coef[LiteRSAref_MAX_PLEN];
}RSArefPrivateKeyLite;

#define ExRSAref_MAX_BITS    4096
#define ExRSAref_MAX_LEN     ((ExRSAref_MAX_BITS + 7) / 8)
#define ExRSAref_MAX_PBITS   ((ExRSAref_MAX_BITS + 1) / 2)
#define ExRSAref_MAX_PLEN    ((ExRSAref_MAX_PBITS + 7)/ 8)

typedef struct RSArefPublicKeyEx_st
{
	unsigned int  bits;
	unsigned char m[ExRSAref_MAX_LEN];
	unsigned char e[ExRSAref_MAX_LEN];
} RSArefPublicKeyEx;

typedef struct RSArefPrivateKeyEx_st
{
	unsigned int  bits;
	unsigned char m[ExRSAref_MAX_LEN];
	unsigned char e[ExRSAref_MAX_LEN];
	unsigned char d[ExRSAref_MAX_LEN];
	unsigned char prime[2][ExRSAref_MAX_PLEN];
	unsigned char pexp[2][ExRSAref_MAX_PLEN];
	unsigned char coef[ExRSAref_MAX_PLEN];
} RSArefPrivateKeyEx;

#if defined(SGD_RSA_MAX_BITS) && (SGD_RSA_MAX_BITS > LiteRSAref_MAX_BITS)
#define RSAref_MAX_BITS    ExRSAref_MAX_BITS
#define RSAref_MAX_LEN     ExRSAref_MAX_LEN
#define RSAref_MAX_PBITS   ExRSAref_MAX_PBITS
#define RSAref_MAX_PLEN    ExRSAref_MAX_PLEN

typedef struct RSArefPublicKeyEx_st  RSArefPublicKey;
typedef struct RSArefPrivateKeyEx_st  RSArefPrivateKey;
#else
#define RSAref_MAX_BITS    LiteRSAref_MAX_BITS
#define RSAref_MAX_LEN     LiteRSAref_MAX_LEN
#define RSAref_MAX_PBITS   LiteRSAref_MAX_PBITS
#define RSAref_MAX_PLEN    LiteRSAref_MAX_PLEN

typedef struct RSArefPublicKeyLite_st  RSArefPublicKey;
typedef struct RSArefPrivateKeyLite_st  RSArefPrivateKey;
#endif

/*ECC密钥*/
#define ECCref_MAX_BITS			512 
#define ECCref_MAX_LEN			((ECCref_MAX_BITS+7) / 8)
#define ECCref_MAX_CIPHER_LEN	136

typedef struct ECCrefPublicKey_st
{
	unsigned int  bits;
	unsigned char x[ECCref_MAX_LEN]; 
	unsigned char y[ECCref_MAX_LEN]; 
} ECCrefPublicKey;

typedef struct ECCrefPrivateKey_st
{
    unsigned int  bits;
    unsigned char K[ECCref_MAX_LEN];
} ECCrefPrivateKey;

/*ECC 密文*/
typedef struct ECCCipher_st
{
	unsigned char x[ECCref_MAX_LEN]; 
	unsigned char y[ECCref_MAX_LEN]; 
  unsigned char M[32];
	unsigned int L; //C的有效长度
	unsigned char C[ECCref_MAX_CIPHER_LEN];
} ECCCipher;

/*ECC 签名*/
typedef struct ECCSignature_st
{
	unsigned char r[ECCref_MAX_LEN];	
	unsigned char s[ECCref_MAX_LEN];	
} ECCSignature;

#define SM9ref_MAX_BITS			256 
#define SM9ref_MAX_LEN			((SM9ref_MAX_BITS+7) / 8)


typedef struct SM9refSignMasterPrivateKey_st
{
	unsigned int  bits;
	unsigned char s[SM9ref_MAX_LEN];
} SM9refSignMasterPrivateKey;

typedef struct SM9refSignMasterPublicKey_st
{
	unsigned int  bits;
	unsigned char xa[SM9ref_MAX_LEN];
	unsigned char xb[SM9ref_MAX_LEN];
	unsigned char ya[SM9ref_MAX_LEN];
	unsigned char yb[SM9ref_MAX_LEN];
} SM9refSignMasterPublicKey;

typedef struct SM9refEncMasterPrivateKey_st
{
	unsigned int  bits;
	unsigned char s[SM9ref_MAX_LEN];
} SM9refEncMasterPrivateKey;

typedef struct SM9refEncMasterPublicKey_st
{
	unsigned int  bits;
	unsigned char x[SM9ref_MAX_LEN];
	unsigned char y[SM9ref_MAX_LEN];
} SM9refEncMasterPublicKey;

typedef struct SM9refSignUserPrivateKey_st
{
	unsigned int  bits;
	unsigned char x[SM9ref_MAX_LEN];
	unsigned char y[SM9ref_MAX_LEN];
} SM9refSignUserPrivateKey;

typedef struct SM9refEncUserPrivateKey_st
{
	unsigned int  bits;
	unsigned char xa[SM9ref_MAX_LEN];
	unsigned char xb[SM9ref_MAX_LEN];
	unsigned char ya[SM9ref_MAX_LEN];
	unsigned char yb[SM9ref_MAX_LEN];
} SM9refEncUserPrivateKey;

typedef struct SM9Signature_st
{
	unsigned char h[SM9ref_MAX_LEN];
	unsigned char x[SM9ref_MAX_LEN];
	unsigned char y[SM9ref_MAX_LEN];
} SM9Signature;

typedef struct SM9Cipher_st
{
	unsigned char x[SM9ref_MAX_LEN];
	unsigned char y[SM9ref_MAX_LEN];
	unsigned char h[SM9ref_MAX_LEN];
	unsigned int  L;
	unsigned char C[1024];
} SM9Cipher;

typedef struct SM9refKeyPackage_st
{
	unsigned char x[SM9ref_MAX_LEN];
	unsigned char y[SM9ref_MAX_LEN];
} SM9refKeyPackage;

//ECDSA
#define ECDSAref_MAX_BITS		640
#define ECDSAref_MAX_LEN			((ECDSAref_MAX_BITS+7) / 8)

//ECDSA密钥曲线参数
#define SGD_ECDSA_P			  0x00080001
#define SGD_ECDSA_K			  0x00080002
#define SGD_ECDSA_B			  0x00080003
#define SGD_ECDSA_BrainpoolR1 0x00080004
#define SGD_ECDSA_BrainpoolT1 0x00080005
#define SGD_ECDSA_WAPIP		  0x00080007
#define SGD_ECDSA_SecpK1	  0x00080008

	typedef struct ECDSArefPublicKey_st
	{
		unsigned int  bits;
		unsigned int  curvetype;
		unsigned char x[ECDSAref_MAX_LEN];
		unsigned char y[ECDSAref_MAX_LEN];
	} ECDSArefPublicKey;

	typedef struct ECDSArefPrivateKey_st
	{
		unsigned int  bits;
		unsigned int  curvetype;
		unsigned char D[ECDSAref_MAX_LEN];
	} ECDSArefPrivateKey;

	typedef struct ECDSASignature_st
	{
		unsigned char r[ECDSAref_MAX_LEN];
		unsigned char s[ECDSAref_MAX_LEN];
	} ECDSASignature;

/*EdDSA密钥*/
#define EdDSAref_MAX_BITS		640 
#define EdDSAref_MAX_LEN		((EdDSAref_MAX_BITS+7) / 8)
	typedef struct EdDSArefPublicKey_st
	{
		unsigned int  bits;
		unsigned int  curvetype;
		unsigned char A[EdDSAref_MAX_LEN];
	} EdDSArefPublicKey;
	typedef struct EdDSArefPrivateKey_st
	{
		unsigned int  bits;
		unsigned int  curvetype;
		unsigned char k[EdDSAref_MAX_LEN];
	} EdDSArefPrivateKey;
	typedef struct EdDSASignature_st
	{
		unsigned char r[EdDSAref_MAX_LEN];
		unsigned char s[EdDSAref_MAX_LEN];
	} EdDSASignature;

	
//DSA
#define LiteDSAref_MAX_BITS    4096
#define LiteDSAref_MAX_LEN     ((LiteDSAref_MAX_BITS + 7) / 8)
#define LiteDSAref_MAX_QLEN   32

	typedef struct DSArefPublicKeyLite_st
	{
		unsigned int  bits;
		unsigned char p[LiteDSAref_MAX_LEN];
		unsigned char q[LiteDSAref_MAX_QLEN];
		unsigned char g[LiteDSAref_MAX_LEN];
		unsigned char pub_key[LiteDSAref_MAX_LEN];
	}DSArefPublicKey;

	typedef struct DSArefPrivateKeyLite_st
	{
		unsigned int  bits;
		unsigned char p[LiteDSAref_MAX_LEN];
		unsigned char q[LiteDSAref_MAX_QLEN];
		unsigned char g[LiteDSAref_MAX_LEN];
		unsigned char pub_key[LiteDSAref_MAX_LEN];
		unsigned char priv_key[LiteDSAref_MAX_QLEN];
	}DSArefPrivateKey;

	typedef struct DSASignature_st
	{
		unsigned char r[LiteDSAref_MAX_QLEN];
		unsigned char s[LiteDSAref_MAX_QLEN];
	} DSASignature;

	
		/*常量定义*/
#define SGD_TRUE		0x00000001
#define SGD_FALSE		0x00000000

	/*对称密码算法标识*/
#define SGD_SM1_ECB		0x00000101
#define SGD_SM1_CBC		0x00000102
#define SGD_SM1_CFB		0x00000104
#define SGD_SM1_OFB		0x00000108
#define SGD_SM1_MAC		0x00000110
#define SGD_SM1_CTR		0x00000120

#define SGD_SMS4_ECB	0x00000401
#define SGD_SMS4_CBC	0x00000402
#define SGD_SMS4_CFB	0x00000404
#define SGD_SMS4_OFB	0x00000408
#define SGD_SMS4_MAC	0x00000410
#define SGD_SMS4_CTR	0x00000420
#define SGD_SMS4_GCM	0x00000440
#define SGD_SMS4_XTS	0x00000480
#define SGD_SMS4_CCM	0x000004A0
#define SGD_SMS4_CMAC	0x000004C0

	//ZUC算法
#define SGD_ZUC			0x00000800      
#define SGD_ZUC_EEA3    0x00000801      //ZUC祖冲之机密性算法 128-EEA3算法
#define SGD_ZUC_EIA3    0x00000802      //ZUC祖冲之完整性算法 128-EIA3算法

#define SGD_3DES_ECB	0x00001001
#define SGD_3DES_CBC	0x00001002
#define SGD_3DES_CFB	0x00001004
#define SGD_3DES_OFB	0x00001008
#define SGD_3DES_MAC	0x00001010
#define SGD_3DES_CTR	0x00001020
#define SGD_3DES_CMAC	0x000010C0


#define SGD_AES_ECB		0x00002001
#define SGD_AES_CBC		0x00002002
#define SGD_AES_CFB		0x00002004
#define SGD_AES_OFB		0x00002008
#define SGD_AES_MAC		0x00002010
#define SGD_AES_CTR		0x00002020
#define SGD_AES_GCM		0x00002040
#define SGD_AES_XTS		0x00002080
#define SGD_AES_CCM		0x000020A0
#define SGD_AES_CMAC	0x000020C0


#define SGD_SM7_ECB		0x00008001
#define SGD_SM7_CBC		0x00008002
#define SGD_SM7_CFB		0x00008004
#define SGD_SM7_OFB		0x00008008
#define SGD_SM7_MAC		0x00008010
#define SGD_SM7_CTR		0x00008020


	/*对称密码算法标识*/
#define SGD_RSA				0x00010000
#define SGD_RSA_SIGN		0x00010100
#define SGD_RSA_ENC			0x00010200

#define SGD_SM2 			0x00020100
#define SGD_SM2_1			0x00020200
#define SGD_SM2_2			0x00020400
#define SGD_SM2_3			0x00020800

#define SGD_DSA				0x00040000
#define SGD_DSA_SIGN		0x00040100

#define SGD_ECDSA			0x00080000
#define SGD_ECDSA_SIGN		0x00080100
#define SGD_ECDSA_ENC		0x00080200

#define SGD_SM9        	 	0x00100000   //SM9算法
#define SGD_SM9_1      	 	0x00100100   //SM9签名算法
#define SGD_SM9_2       	0x00100200   //SM9密钥交换算法
#define SGD_SM9_3       	0x00100400   //SM9密钥封装算法
#define SGD_SM9_4       	0x00100800   //SM9加密算法

//EdDSA
#define SGD_EdDSA			0x00200100
#define SGD_EdDSA_SIGN		0x00200200
#define SGD_EdDSA_ENC		0x00200400

#define SGD_SM3				0x00000001
#define SGD_SHA1			0x00000002
#define SGD_SHA256			0x00000004
#define SGD_SHA512			0x00000008
#define SGD_SHA384			0x00000010
#define SGD_SHA224			0x00000020
#define SGD_MD5				0x00000080

#define SGD_SHA3_256		0x00001004
#define SGD_SHA3_512		0x00001008
#define SGD_SHA3_384		0x00001010
#define SGD_SHA3_224		0x00001020

#define SGD_SHA3_KE128		0x00001040
#define SGD_SHA3_KE256		0x00001080
		

/*标准错误码定义*/
#define SDR_OK				0x0						/*成功*/
#define SDR_ERROR           0x1
#define SDR_BASE			0x01000000
#define SDR_UNKNOWERR		(SDR_BASE + 0x00000001)		/*未知错误*/
#define SDR_NOTSUPPORT		(SDR_BASE + 0x00000002)	    /*不支持的接口调用*/
#define SDR_COMMFAIL		(SDR_BASE + 0x00000003)		/*与设备通信错误*/
#define SDR_HARDFAIL		(SDR_BASE + 0x00000004)		/*硬件错误*/
#define SDR_OPENDEVICE		(SDR_BASE + 0x00000005)		/*打开设备错误*/
#define SDR_OPENSESSION		(SDR_BASE + 0x00000006)		/*打开会话句柄错误*/
#define SDR_PARDENY			(SDR_BASE + 0x00000007)		/*权限不满足*/
#define SDR_KEYNOTEXIST		(SDR_BASE + 0x00000008)		/*密钥不存在*/
#define SDR_ALGNOTSUPPORT	(SDR_BASE + 0x00000009)		/*不支持的算法*/
#define SDR_ALGMODNOTSUPPORT (SDR_BASE + 0x0000000A)    /*不支持的算法模式*/
#define SDR_PKOPERR			(SDR_BASE + 0x0000000B)		/*公钥运算错误*/
#define SDR_SKOPERR			(SDR_BASE + 0x0000000C)		/*私钥运算错误*/
#define SDR_SIGNERR			(SDR_BASE + 0x0000000D)		/*签名错误*/
#define SDR_VERIFYERR		(SDR_BASE + 0x0000000E)		/*验证错误*/
#define SDR_SYMOPERR		(SDR_BASE + 0x0000000F)		/*对称运算错误*/
#define SDR_STEPERR			(SDR_BASE + 0x00000010)		/*步骤错误*/
#define SDR_FILESIZEERR		(SDR_BASE + 0x00000011)		/*文件大小错误 | 数据长度错误*/
#define SDR_FILENOEXIST		(SDR_BASE + 0x00000012)		/*文件不存在*/
#define SDR_FILEOFSERR		(SDR_BASE + 0x00000013)		/*文件操作偏移量错误*/
#define SDR_KEYTYPEERR		(SDR_BASE + 0x00000014)		/*密钥类型错误*/
#define SDR_KEYERR			(SDR_BASE + 0x00000015)		/*密钥错误*/
//add 14/03/10    GM/T 0018-2012
#define SDR_ENCDATAERR		(SDR_BASE + 0x00000016)		/*ECC加密数据错误*/
#define SDR_RANDERR			(SDR_BASE + 0x00000017)		/*随机数产生失败*/
#define SDR_PRKRERR			(SDR_BASE + 0x00000018)		/*私钥使用权限获取失败，未获取 | 私钥访问控制权限释放失败，未获取*/
#define SDR_MACERR			(SDR_BASE + 0x00000019)		/*MAC运算失败*/
#define SDR_FILEEXISTS		(SDR_BASE + 0x0000001A)		/*指定文件已存在*/
#define SDR_FILEWERR		(SDR_BASE + 0x0000001B)		/*文件写入失败*/
#define SDR_NOBUFFER		(SDR_BASE + 0x0000001C)		/*存储空间不足*/
#define SDR_INARGERR		(SDR_BASE + 0x0000001D)		/*输入参数错误：1、数据长度错误*/
#define SDR_OUTARGERR		(SDR_BASE + 0x0000001E)		/*输出参数错误*/

/*扩展错误码*/
#define SWR_BASE				(SDR_BASE + 0x00010000)	/*自定义错误码基础值*/
#define SWR_INVALID_USER		(SWR_BASE + 0x00000001)	/*无效的用户名*/
#define SWR_INVALID_AUTHENCODE	(SWR_BASE + 0x00000002)	/*无效的授权码*/
#define SWR_PROTOCOL_VER_ERR	(SWR_BASE + 0x00000003)	/*不支持的协议版本*/
#define SWR_INVALID_PACKAGE		(SWR_BASE + 0x00000004)	/*错误的数据包格式*/
#define SWR_INVALID_PARAMETERS	(SWR_BASE + 0x00000005)	/*参数错误*/
#define SWR_FILE_ALREADY_EXIST	(SWR_BASE + 0x00000006)	/*已存在同名文件*/

#define SWR_SOCKET_ERR_MASK		0xFFFFFF00	/*用于检查是否是SOCKET错误*/
#define SWR_SOCKET_ERR_BASE		(SWR_BASE + 0x00000100)	/*用于检查是否是SOCKET错误*/
#define SWR_SOCKET_TIMEOUT		(SWR_BASE + 0x00000100)	/*超时错误*/
#define SWR_CONNECT_ERR			(SWR_BASE + 0x00000101)	/*连接服务器错误*/
#define SWR_SET_SOCKOPT_ERR		(SWR_BASE + 0x00000102)	/*设置Socket参数错误*/
#define SWR_SOCKET_SEND_ERR		(SWR_BASE + 0x00000104)	/*发送LOGINRequest错误*/
#define SWR_SOCKET_RECV_ERR		(SWR_BASE + 0x00000105)	/*接收LOGINRequest错误*/
#define SWR_SOCKET_RECV_0		(SWR_BASE + 0x00000106)	/*接收LOGINRequest为0*/
#define SWR_SOCKET_ERR			(SWR_BASE + 0x00000107)	/*socket错误*/

#define SWR_SEM_TIMEOUT			(SWR_BASE + 0x00000200)	/*超时错误*/
#define SWR_NO_AVAILABLE_HSM	(SWR_BASE + 0x00000201)	/*没有可用的加密机*/
#define SWR_NO_AVAILABLE_CSM	(SWR_BASE + 0x00000202)	/*加密机内没有可用的加密模块*/
#define SWR_ADD_HSM				(SWR_BASE + 0x00000203)	/*密码机负载数量已达最大值，不可再增加*/
#define SWR_DELETE_HSM			(SWR_BASE + 0x00000204)	/*密码机负载数量已达最小值，不可再减少*/

#define SWR_CONFIG_ERR			(SWR_BASE + 0x00000301)	/*配置信息错误*/
#define SWR_NULL_ERR			(SWR_BASE + 0x00000302)	/*空指针*/
#define SWR_TYPE_ERR			(SWR_BASE + 0x00000303)	/*不支持的输入类型*/
#define SWR_CALLOC_ERR			(SWR_BASE + 0x00000304)	/*申请内存空间calloc错误*/
#define SWR_MALLOC_ERR			(SWR_BASE + 0x00000305)	/*申请内存空间malloc错误*/

//openssl
#define SWR_LOAD_CA_ERR			(SWR_BASE + 0x00000401)	/*加载CA证书错误*/
#define SWR_PARSE_PFX_ERR		(SWR_BASE + 0x00000402)	/*解析个人证书错误*/
#define SWR_PATH_ERR			(SWR_BASE + 0x00000403)	/*路径错误*/
#define SWR_CHECK_PRIKEY_ERR	(SWR_BASE + 0x00000404)	/*验证pfx证书错误*/
#define SWR_SET_CiPHERLIST_ERR	(SWR_BASE + 0x00000405)	/*设置加密模式错误*/
#define SWR_CREATE_SEMAPHORE_ERR	(SWR_BASE + 0x00000406)	/*创建信号量错误*/
#define SWR_RELEASE_SEMAPHORE_ERR	(SWR_BASE + 0x00000407)	/*增加信号量计数错误*/
#define SWR_DELETE_SEMAPHORE_ERR	(SWR_BASE + 0x00000408)	/*删除信号量错误*/
#define SWR_CHANGE_SEMAPHORE_ERR	(SWR_BASE + 0x00000409)	/*改变信号量的值错误*/	

//server
#define SWR_PARAMETERS_ERROR		0x01000201			/*配置参数错误*/
#define SWR_FILENOEXIST				0x01000202			/*配置文件不存在*/
#define SWR_DATANOEXIST				0x01000203			/*配置参数不存在*/

#define SWR_PIN_LENGTH				0x01000302			//用户pin口令长度错误
#define SWR_CHECKPIN				0x01000303			//用户pin口令错误
#define SWR_KEY_OPENUSBKEY			0x01000304			//用户打开USBKEY错误
#define SWR_KEY_USERTYPE			0x01000306			//USBKEY用户类型错误
#define SWR_KEY_USERNOADDED			0x01000307			//用户未添加
#define SWR_KEY_USERMODIFYPASSWD	0x0100030B			//修改用户口令失败
#define SWR_KEY_USELOGOUT			0x0100030C			//用户注销失败
#define SWR_KEY_MAXMANAGENUM		0x0100030D			//管理员数目已经达到最大值
#define SWR_KEY_MANAEXISTED			0x0100030F			//管理员UKEY已添加
#define SWR_KEY_ADDMANGUSER			0x01000315			//增加管理员用户KEY失败
#define SWR_KEY_MAXOPERNUM			0x01000316			//操作员数目已经达到最大值
#define SWR_KEY_OPEREXISTED			0x01000318			//操作员UKEY已添加
#define SWR_KEY_ADDOPERUSER			0x01000319			//增加操作员用户KEY失败
#define SWR_KEY_USEDELETE			0x0100031A			//UKEY用户删除失败
#define SWR_KEY_MANAGEDELETE		0x0100031B			//管理员只剩一个不允许删除
#define SWR_KEY_MAXAUDITOR			0x01000320			//审计员数目已经达到最大值
#define SWR_KEY_AUDITEXISTED		0x01000322			//审计员UKEY已添加
#define SWR_KEY_ADDAUDITOR			0x01000323			//增加审计员用户KEY失败
#define SWR_KEY_PRIVILEGENODISSA	0x01000324			//权限不满足

#define SWR_CHECK_AUTHENCODE		 0x01000407			/*私钥访问控制码错误*/
#define SWR_KEYNOTEXIST				 0x01000408			/*密钥不存在*/
#define SWR_INDEX	                 0x01000409			/*索引超出范围*/
#define SWR_KEYLEN	                 0x0100040A			/*密钥长度错误;*/
#define SWR_AUTHENCODE	             0x0100040B			/*私钥访问控制码长度错误*/
#define SWR_FILENOEXISTEX				 0x01000412			/*文件不存在*/
#define SWR_FILEOFSERR		         0x01000413			/*文件操作偏移量错误，读写文件时，偏移量加长度超出文件大小*/
#define SWR_HANDLENULL		         0x01000414			/*输入的参数句柄或文件名为空*/
#define SWR_KEYERR					 0x01000415			/*密钥错误*/
#define SWR_INVALID_FILENAME_LEN	 0x01000416			/*无效的文件名或文件长度，文件名大于32字符或文件长度小于1*/
#define SWR_FILE_ALREADY_EXISTEX	     0x01000417			/*已存在同名文件*/
#define SWR_USERFILEERR			     0x01000418			/*用户文件错误*/
#define SWR_CREATEUSERFILE			 0x01000419			/*创建用户文件错误，*/
#define SWR_BACKUPCHECK				 0x01000420			/*备份文件错误*/
#define SWR_INVALID_COUNT		     0x01000424			/*密钥个数超出范围*/


#define SDR_MANAGEMENT_DENY			  0x00001001		/*管理权限不满足*/
#define SDR_OPERATION_DENY			  0x00001002		/*操作权限不满足*/
#define SDR_PARAMENT_ERR			  0x00001013		/*参数错误*/

#define SWR_BASEEX					  0x01000500
#define SWR_KEYBIITES	              0x01000501		 /*密钥长度错误*/
#define SWR_INVALID_PARAMETERSEX	  0x01000502		 /*无效的参数*/
#define SWR_GEN_KEY	                  0x01000503		 /*生成密钥错误*/
#define SWR_CURVE	                  0x01000504         /*曲线参数错误*/
#define SWR_LENGTH                    0x01000506		 /*数据长度错误*/
#define SWR_ALGNOTSUPPORT	          0x0100050A		 /*算法不支持*/
#define SWR_PKOPERR			          0x0100050B		 /*公钥运算错误*/
#define SWR_SKOPERR			          0x0100050C		 /*私钥运算错误*/
#define SWR_SIGNERR			          0x0100050D		 /*签名运算错误*/
#define SWR_VERIFYERR		          0x0100050E		 /*验签运算错误*/
#define SWR_HASH	                  0x01000520		 /*hash运算错误*/
#define SWR_HASH_INIT                 0x01000548		 /*分步哈希错误*/
#define SWR_HASH_Update               0x01000549		 /*分步哈希错误*/
#define SWR_HASH_Final                0x01000550		 /*分步哈希错误*/
#define SWR_DSA_GEN_PAR_ERR			  0x01000551		 /*生成DSA密钥参数错误*/
#define SWR_DSA_GEN_KEY_ERR			  0x01000552		 /*生成DSA密钥错误*/

#define SWR_IPFORMAT_INVALID		  0x01000701		 //IP 端口号初始化错误 网络地址格式错误
#define SWR_SERCER_NOEXIST			  0x01000704		 //服务根证不存在
#define SWR_SERPFX_NOEXIST			  0x01000705		 //服务个人证书不存在
#define SWR_SERPFXENC_NOEXIST		  0x01000706		 //服务加密个人证书不存在
#define SWR_SSLCONTEXT_ERR			  0x01000707		 //设置加密套件错误
#define SWR_FILE_ERR				  0x01000708		 //服务配置文件错误,SSL算法错误
#define SWR_PFX_ERR					  0x01000709		 //加载个人证书错误
#define SWR_ENCPFX_ERR				  0x0100070a		 //加载个人加密证书错误
#define SWR_SSL_CERTCHAIN_FAIL		  0x0100070b		 //SSL根证书验证失败
#define SWR_SSL_VERMODE_FAIL		  0x0100070c		 //SSL设置认证模式错误
#define NETERR_AUTHENCODE_INVALID	  0x01000715		 //无效的授权码
#define NETERR_BUFFERSIZE_INVALID	  0x01000716		 //无效的报文长度
#define SWR_PROTOCOL_VER_ERREX		  0x01000717		 //不支持的协议版本
#define NETERR_CONNECT_OVERSIZE		  0x01000718		 //连接数达到上限
#define NETERR_MEMORY_ERR			  0x01000719		 //内存错误
#define SWR_INVALID_PACKAGEEX		  0x0100071a		 //错误的数据包格式
#define SWR_PARDENY					  0x0100071b     	 //权限不满足
							   
#define SWR_NOTSUPPORT				  0x01000802		 //不支持
#define SWR_TIMEOUT					  0x01000803		 //超时错误
#define SWR_ALGUNSUPPORT			  0x01000809		 //不支持的算法




/*0018接口*/

//设备管理类接口
 SGD_RV SDF_OpenDevice(SGD_HANDLE *phDeviceHandle);
 SGD_RV SDF_CloseDevice(SGD_HANDLE hDeviceHandle);
 SGD_RV SDF_OpenSession(SGD_HANDLE hDeviceHandle, SGD_HANDLE *phSessionHandle);
 SGD_RV SDF_CloseSession(SGD_HANDLE hSessionHandle);
 SGD_RV SDF_GetDeviceInfo(SGD_HANDLE hSessionHandle, DEVICEINFO *pstDeviceInfo);
 SGD_RV SDF_GenerateRandom(SGD_HANDLE hSessionHandle, SGD_UINT32  uiLength, SGD_UCHAR *pucRandom);
 SGD_RV SDF_GetPrivateKeyAccessRight(SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex,SGD_UCHAR *pucPassword, SGD_UINT32  uiPwdLength);
 SGD_RV SDF_ReleasePrivateKeyAccessRight(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex);

//密钥管理类接口
 SGD_RV SDF_ExportSignPublicKey_RSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,RSArefPublicKey *pucPublicKey);
 SGD_RV SDF_ExportEncPublicKey_RSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,RSArefPublicKey *pucPublicKey);
 SGD_RV SDF_GenerateKeyPair_RSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyBits,RSArefPublicKey *pucPublicKey,RSArefPrivateKey *pucPrivateKey);
 SGD_RV SDF_GenerateKeyWithIPK_RSA (SGD_HANDLE hSessionHandle, SGD_UINT32 uiIPKIndex,SGD_UINT32 uiKeyBits,SGD_UCHAR *pucKey,SGD_UINT32 *puiKeyLength,SGD_HANDLE *phKeyHandle);
 SGD_RV SDF_GenerateKeyWithEPK_RSA (SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyBits,RSArefPublicKey *pucPublicKey,SGD_UCHAR *pucKey,SGD_UINT32 *puiKeyLength,SGD_HANDLE *phKeyHandle);
 SGD_RV SDF_ImportKeyWithISK_RSA (SGD_HANDLE hSessionHandle, SGD_UINT32 uiISKIndex,SGD_UCHAR *pucKey,SGD_UINT32 uiKeyLength,SGD_HANDLE *phKeyHandle);
 SGD_RV SDF_ExchangeDigitEnvelopeBaseOnRSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,RSArefPublicKey *pucPublicKey,SGD_UCHAR *pucDEInput,SGD_UINT32  uiDELength,SGD_UCHAR *pucDEOutput,SGD_UINT32  *puiDELength);

 SGD_RV SDF_ExportSignPublicKey_ECC(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,ECCrefPublicKey *pucPublicKey);
 SGD_RV SDF_ExportEncPublicKey_ECC(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,ECCrefPublicKey *pucPublicKey);
 SGD_RV SDF_GenerateKeyPair_ECC(SGD_HANDLE hSessionHandle, SGD_UINT32  uiAlgID,SGD_UINT32  uiKeyBits,ECCrefPublicKey *pucPublicKey,ECCrefPrivateKey *pucPrivateKey);
 SGD_RV SDF_GenerateKeyWithIPK_ECC (SGD_HANDLE hSessionHandle, SGD_UINT32 uiIPKIndex,SGD_UINT32 uiKeyBits,ECCCipher *pucKey,SGD_HANDLE *phKeyHandle);
 SGD_RV SDF_GenerateKeyWithEPK_ECC (SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyBits,SGD_UINT32  uiAlgID,ECCrefPublicKey *pucPublicKey,ECCCipher *pucKey,SGD_HANDLE *phKeyHandle);
 SGD_RV SDF_ImportKeyWithISK_ECC (SGD_HANDLE hSessionHandle,SGD_UINT32 uiISKIndex,ECCCipher *pucKey,SGD_HANDLE *phKeyHandle);
 SGD_RV SDF_GenerateAgreementDataWithECC (SGD_HANDLE hSessionHandle, SGD_UINT32 uiISKIndex,SGD_UINT32 uiKeyBits,SGD_UCHAR *pucSponsorID,SGD_UINT32 uiSponsorIDLength,ECCrefPublicKey  *pucSponsorPublicKey,ECCrefPublicKey  *pucSponsorTmpPublicKey,SGD_HANDLE *phAgreementHandle);
 SGD_RV SDF_GenerateKeyWithECC (SGD_HANDLE hSessionHandle, SGD_UCHAR *pucResponseID,SGD_UINT32 uiResponseIDLength,ECCrefPublicKey *pucResponsePublicKey,ECCrefPublicKey *pucResponseTmpPublicKey,SGD_HANDLE hAgreementHandle,SGD_HANDLE *phKeyHandle);
 SGD_RV SDF_GenerateAgreementDataAndKeyWithECC (SGD_HANDLE hSessionHandle, SGD_UINT32 uiISKIndex,SGD_UINT32 uiKeyBits,SGD_UCHAR *pucResponseID,SGD_UINT32 uiResponseIDLength,SGD_UCHAR *pucSponsorID,SGD_UINT32 uiSponsorIDLength,ECCrefPublicKey *pucSponsorPublicKey,ECCrefPublicKey *pucSponsorTmpPublicKey,ECCrefPublicKey  *pucResponsePublicKey,	ECCrefPublicKey  *pucResponseTmpPublicKey,SGD_HANDLE *phKeyHandle);
 SGD_RV SDF_ExchangeDigitEnvelopeBaseOnECC(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,SGD_UINT32  uiAlgID,ECCrefPublicKey *pucPublicKey,ECCCipher *pucEncDataIn,ECCCipher *pucEncDataOut);

 SGD_RV SDF_GenerateKeyWithKEK (SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyBits,SGD_UINT32  uiAlgID,SGD_UINT32 uiKEKIndex, SGD_UCHAR *pucKey, SGD_UINT32 *puiKeyLength, SGD_HANDLE *phKeyHandle);
 SGD_RV SDF_ImportKeyWithKEK (SGD_HANDLE hSessionHandle, SGD_UINT32  uiAlgID,SGD_UINT32 uiKEKIndex, SGD_UCHAR *pucKey, SGD_UINT32 uiKeyLength, SGD_HANDLE *phKeyHandle);
 SGD_RV SDF_ImportKey (SGD_HANDLE hSessionHandle, SGD_UCHAR *pucKey, SGD_UINT32 uiKeyLength,SGD_HANDLE *phKeyHandle);
 SGD_RV SDF_DestroyKey (SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle);

 SGD_RV SDF_ExternalPublicKeyOperation_RSA(SGD_HANDLE hSessionHandle, RSArefPublicKey *pucPublicKey,SGD_UCHAR *pucDataInput,SGD_UINT32  uiInputLength,SGD_UCHAR *pucDataOutput,SGD_UINT32  *puiOutputLength);
 SGD_RV SDF_ExternalPrivateKeyOperation_RSA(SGD_HANDLE hSessionHandle, RSArefPrivateKey *pucPrivateKey,SGD_UCHAR *pucDataInput,SGD_UINT32  uiInputLength,SGD_UCHAR *pucDataOutput,SGD_UINT32  *puiOutputLength);
 SGD_RV SDF_InternalPublicKeyOperation_RSA(SGD_HANDLE hSessionHandle,SGD_UINT32 uiKeyIndex,SGD_UCHAR *pucDataInput,SGD_UINT32  uiInputLength,SGD_UCHAR *pucDataOutput,SGD_UINT32  *puiOutputLength);
 SGD_RV SDF_InternalPrivateKeyOperation_RSA(SGD_HANDLE hSessionHandle,SGD_UINT32 uiKeyIndex,SGD_UCHAR *pucDataInput,SGD_UINT32  uiInputLength,SGD_UCHAR *pucDataOutput,SGD_UINT32  *puiOutputLength);

 SGD_RV SDF_ExternalSign_ECC(SGD_HANDLE hSessionHandle,SGD_UINT32 uiAlgID,ECCrefPrivateKey *pucPrivateKey,SGD_UCHAR *pucData,SGD_UINT32  uiDataLength,ECCSignature *pucSignature);
 SGD_RV SDF_ExternalVerify_ECC(SGD_HANDLE hSessionHandle,SGD_UINT32 uiAlgID,ECCrefPublicKey *pucPublicKey,SGD_UCHAR *pucDataInput,SGD_UINT32  uiInputLength,ECCSignature *pucSignature);
 SGD_RV SDF_InternalSign_ECC(SGD_HANDLE hSessionHandle,SGD_UINT32  uiISKIndex,SGD_UCHAR *pucData,SGD_UINT32  uiDataLength,ECCSignature *pucSignature);
 SGD_RV SDF_InternalVerify_ECC(SGD_HANDLE hSessionHandle,SGD_UINT32  uiISKIndex,SGD_UCHAR *pucData,SGD_UINT32  uiDataLength,ECCSignature *pucSignature);
 SGD_RV SDF_ExternalEncrypt_ECC(SGD_HANDLE hSessionHandle,SGD_UINT32 uiAlgID,ECCrefPublicKey *pucPublicKey,SGD_UCHAR *pucData,SGD_UINT32  uiDataLength,ECCCipher *pucEncData);
 SGD_RV SDF_ExternalDecrypt_ECC(SGD_HANDLE hSessionHandle,SGD_UINT32 uiAlgID,ECCrefPrivateKey *pucPrivateKey,ECCCipher *pucEncData,SGD_UCHAR *pucData,SGD_UINT32  *puiDataLength);

 SGD_RV SDF_Encrypt(SGD_HANDLE hSessionHandle,SGD_HANDLE hKeyHandle,SGD_UINT32 uiAlgID,SGD_UCHAR *pucIV,SGD_UCHAR *pucData,SGD_UINT32 uiDataLength,SGD_UCHAR *pucEncData,SGD_UINT32  *puiEncDataLength);
 SGD_RV SDF_Decrypt (SGD_HANDLE hSessionHandle,SGD_HANDLE hKeyHandle,SGD_UINT32 uiAlgID,SGD_UCHAR *pucIV,SGD_UCHAR *pucEncData,SGD_UINT32  uiEncDataLength,SGD_UCHAR *pucData,SGD_UINT32 *puiDataLength);
 SGD_RV SDF_CalculateMAC(SGD_HANDLE hSessionHandle,SGD_HANDLE hKeyHandle,SGD_UINT32 uiAlgID,SGD_UCHAR *pucIV,SGD_UCHAR *pucData,SGD_UINT32 uiDataLength,SGD_UCHAR *pucMAC,SGD_UINT32  *puiMACLength);

 SGD_RV SDF_HashInit(SGD_HANDLE hSessionHandle,SGD_UINT32 uiAlgID,ECCrefPublicKey *pucPublicKey,SGD_UCHAR *pucID,SGD_UINT32 uiIDLength);
 SGD_RV SDF_HashUpdate(SGD_HANDLE hSessionHandle,SGD_UCHAR *pucData,SGD_UINT32  uiDataLength);
 SGD_RV SDF_HashFinal(SGD_HANDLE hSessionHandle,SGD_UCHAR *pucHash,SGD_UINT32  *puiHashLength);

 SGD_RV SDF_CreateFile(SGD_HANDLE hSessionHandle,SGD_UCHAR *pucFileName,SGD_UINT32 uiNameLen,SGD_UINT32 uiFileSize);
 SGD_RV SDF_ReadFile(SGD_HANDLE hSessionHandle,SGD_UCHAR *pucFileName,SGD_UINT32 uiNameLen,SGD_UINT32 uiOffset,SGD_UINT32 *puiReadLength,SGD_UCHAR *pucBuffer);
 SGD_RV SDF_WriteFile(SGD_HANDLE hSessionHandle,SGD_UCHAR *pucFileName,SGD_UINT32 uiNameLen,SGD_UINT32 uiOffset,SGD_UINT32 uiWriteLength,SGD_UCHAR *pucBuffer);
 SGD_RV SDF_DeleteFile(SGD_HANDLE hSessionHandle,SGD_UCHAR *pucFileName,SGD_UINT32 uiNameLen);


/*扩展接口*/
 
 SGD_RV SDF_OpenDeviceWithPath(SGD_CHAR *pcCfgPath, SGD_HANDLE *phDeviceHandle);
 SGD_RV SDF_OpenDeviceWithPathAndName(char *pcCfgfile, void **phDeviceHandle);
 SGD_RV SDF_OpenDeviceWithParameter(SGD_HANDLE *phDeviceHandle, char **argv);
 SGD_RV SDF_OpenDeviceWithParameter_Ex(SGD_HANDLE *phDeviceHandle, int poolSize, char *sIP,int nPort, char*sPwd);
 
 SGD_RV SDF_GetKeyStatus(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyType, SGD_UINT32 *puiKeyStatus, SGD_UINT32 *puiKeyCount);
 SGD_RV SDF_GetVersion(unsigned int *puiVersion);
 SGD_RV SDF_AddHsm(SGD_HANDLE hDeviceHandle, SGD_CHAR *sIP, SGD_UINT32 nPort, SGD_CHAR *sPwd);
 SGD_RV SDF_DeleteHsm(SGD_HANDLE hDeviceHandle, SGD_CHAR *sIP);
 
 
 SGD_RV SDF_GetSymmKeyHandle(SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SGD_HANDLE *phKeyHandle);

 SGD_RV SDF_InternalPublicKeyOperation_RSA_Ex(SGD_HANDLE hSessionHandle,SGD_UINT32  uiKeyIndex,SGD_UINT32  uiKeyUsage,SGD_UCHAR *pucDataInput,SGD_UINT32  uiInputLength,SGD_UCHAR *pucDataOutput,SGD_UINT32  *puiOutputLength);
 SGD_RV SDF_InternalPrivateKeyOperation_RSA_Ex(SGD_HANDLE hSessionHandle,SGD_UINT32  uiKeyIndex,SGD_UINT32  uiKeyUsage,SGD_UCHAR *pucDataInput,SGD_UINT32  uiInputLength,SGD_UCHAR *pucDataOutput,SGD_UINT32  *puiOutputLength);

 SGD_RV SDF_InternalSign_ECC_Ex(SGD_HANDLE hSessionHandle,SGD_UINT32  uiISKIndex,SGD_UINT32 uiAlgID,SGD_UCHAR *pucData,SGD_UINT32  uiDataLength,ECCSignature *pucSignature);
 SGD_RV SDF_InternalVerify_ECC_Ex(SGD_HANDLE hSessionHandle,SGD_UINT32  uiISKIndex,SGD_UINT32 uiAlgID,SGD_UCHAR *pucData,SGD_UINT32  uiDataLength,ECCSignature *pucSignature);

 SGD_RV SDF_InternalEncrypt_ECC(SGD_HANDLE hSessionHandle,SGD_UINT32  uiISKIndex, SGD_UINT32 uiAlgID,SGD_UCHAR *pucData,SGD_UINT32  uiDataLength,ECCCipher *pucEncData);
 SGD_RV SDF_InternalDecrypt_ECC(SGD_HANDLE hSessionHandle,SGD_UINT32  uiISKIndex,SGD_UINT32 uiAlgID,ECCCipher *pucEncData,SGD_UCHAR *pucData,SGD_UINT32  *puiDataLength);

 SGD_RV SDF_Hash(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, ECCrefPublicKey *pucPublicKey, SGD_UCHAR *pucID, SGD_UINT32 uiIDLength, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, SGD_UCHAR *pucHash, SGD_UINT32  *puiHashLength);
 SGD_RV SDF_HMAC(SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle, SGD_UINT32 uiAlgID, SGD_UCHAR  *pucData, SGD_UINT32  uiDataLength, SGD_UCHAR  *pucHmac, SGD_UINT32 *puiHmacLen);


 SGD_RV SDF_GenerateSignMasterPrivateKey_SM9(SGD_HANDLE hSessionHandle,	SGD_UINT32  uiKeyBits,	SM9refSignMasterPublicKey *pucPublicKey, SM9refSignMasterPrivateKey *pucPrivateKey, SGD_UCHAR *pucPairG, SGD_UINT32 *puiPairGLen);

 SGD_RV SDF_GenerateEncMasterPrivateKey_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyBits, SM9refEncMasterPublicKey *pucPublicKey, SM9refEncMasterPrivateKey *pucPrivateKey, SGD_UCHAR *pucPairG, SGD_UINT32 *puiPairGLen);

 SGD_RV SDF_ImportSignMasterPrivateKey_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, SM9refSignMasterPublicKey *pucPublicKey, SM9refSignMasterPrivateKey *pucPrivateKey, SGD_UCHAR *pucPairG, SGD_UINT32 puiPairGLen);

 SGD_RV SDF_ImportEncMasterPrivateKey_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, SM9refEncMasterPublicKey *pucPublicKey, SM9refEncMasterPrivateKey *pucPrivateKey, SGD_UCHAR *pucPairG, SGD_UINT32 puiPairGLen);

 SGD_RV SDF_GenerateSignUserPrivateKey_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SGD_UCHAR hid, SGD_UCHAR *pucUserID, SGD_UINT32 uiUserIDLen, SM9refSignUserPrivateKey  *pUserPrivateKey);

 SGD_RV SDF_GenerateEncUserPrivateKey_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SGD_UCHAR hid, SGD_UCHAR *pucUserID, SGD_UINT32 uiUserIDLen, SM9refEncUserPrivateKey  *pUserPrivateKey);


 SGD_RV SDF_Sign_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SM9refSignUserPrivateKey  *pUserPrivateKey, SM9refSignMasterPublicKey *pMasterPublicKey, SGD_UCHAR *pucDataInput, SGD_UINT32 uiDataInputLen, SM9Signature  *pSignature);

 SGD_RV SDF_SignEx_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SM9refSignUserPrivateKey  *pUserPrivateKey, SM9refSignMasterPublicKey *pMasterPublicKey, SGD_UCHAR *pPairG, SGD_UINT32 uiPairGLen, SGD_UCHAR *pucDataInput, SGD_UINT32 uiDataInputLen, SM9Signature  *pSignature);

 SGD_RV SDF_Verify_SM9(SGD_HANDLE hSessionHandle, SGD_UCHAR hid, SGD_UCHAR *pucUserID, SGD_UINT32  uiUserIDLen, SM9refSignMasterPublicKey  *pMasterPublicKey, SGD_UCHAR *pucData, SGD_UINT32   uiDataInputLen, SM9Signature  *pSignature);

 SGD_RV SDF_VerifyEx_SM9(SGD_HANDLE hSessionHandle, SGD_UCHAR hid, SGD_UCHAR *pucUserID, SGD_UINT32  uiUserIDLen, SM9refSignMasterPublicKey  *pMasterPublicKey, SGD_UCHAR *pPairG, SGD_UINT32 uiPairGLen, SGD_UCHAR *pucData, SGD_UINT32   uiDataInputLen, SM9Signature  *pSignature);

 SGD_RV SDF_Encrypt_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32 flag, SGD_UCHAR hid, SGD_UCHAR *pucUserID, SGD_UINT32  uiUserIDLen, SM9refEncMasterPublicKey *pPubluicKey, SGD_UCHAR *pucData, SGD_UINT32   uiDataLength, SM9Cipher *pCipher);

 SGD_RV SDF_EncryptEx_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32 flag, SGD_UCHAR hid, SGD_UCHAR *pucUserID, SGD_UINT32  uiUserIDLen, SM9refEncMasterPublicKey *pPubluicKey, SGD_UCHAR *pPairG, SGD_UINT32  nPairGLen, SGD_UCHAR *pucData, SGD_UINT32   uiDataLength, SM9Cipher *pCipher);

 SGD_RV SDF_Decrypt_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32 flag, SGD_UCHAR *pucUserID, SGD_UINT32  uiUserIDLen, SGD_UINT32 uiKeyIndex, SM9refEncUserPrivateKey  *pUserPrivateKey, SM9Cipher * pCipher, SGD_UCHAR *pucPlainData, SGD_UINT32  *uiPlainDataLength);

 SGD_RV SDF_ExportSignMasterPublicKey_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, SM9refSignMasterPublicKey *pPublicKey);

 SGD_RV SDF_ExportEncMasterPublicKey_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, SM9refEncMasterPublicKey *pPublicKey);

 SGD_RV SDF_ExportSignMasterKeyPairG_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, SGD_UCHAR *pPairG, SGD_UINT32 *puiPairGLen);

 SGD_RV SDF_ExportEncMasterKeyPairG_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, SGD_UCHAR *pPairG, SGD_UINT32 *puiPairGLen);

 SGD_RV SDF_ImportUserSignPrivateKey_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SM9refSignUserPrivateKey  *pUserPrivateKey);

 SGD_RV SDF_ImportUserEncPrivateKey_SM9(SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SM9refEncUserPrivateKey  *pUserPrivateKey);

 SGD_RV  SDF_Encap_SM9(SGD_HANDLE hSessionHandle, SGD_UCHAR hid, SGD_UCHAR *pucUserID, SGD_UINT32  uiUserIDLen, SM9refEncMasterPublicKey  *pPublicKey, SGD_UINT32 uiKeyLen, SGD_UCHAR *pKey, SM9refKeyPackage *pKeyPackage);

 SGD_RV SDF_Decap_SM9(SGD_HANDLE hSessionHandle, SGD_UCHAR *pucUserID, SGD_UINT32  uiUserIDLen, SGD_UINT32 uiKeyIndex, SM9refEncUserPrivateKey  *pUserPrivateKey, SM9refKeyPackage *pKeyPackage, SGD_UINT32  uiKeyLen, SGD_UCHAR *pucKey);
 
 SGD_RV SDF_GenerateAgreementDataWithSM9(SGD_HANDLE hSessionHandle, SGD_UCHAR hid, SGD_UCHAR *pucResponseID, SGD_UINT32 uiResponseIDLength, SM9refEncMasterPublicKey  *pPublicKey, SM9refEncMasterPublicKey  *pucSponsorTmpPublicKey, void  **phAgreementHandle);

 SGD_RV SDF_GenerateAgreemetDataAndKeyWithSM9( SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyLen, SGD_UCHAR hid, SGD_UCHAR * pucResponseID, SGD_UINT32 uiResponseIDLen, SGD_UCHAR * pucSponsorID, SGD_UINT32 uiSponsorIDLen, SGD_UINT32 uiKeyIndex, SM9refEncUserPrivateKey  *pucResponsePrivateKey, SM9refEncMasterPublicKey *pucPublicKey, SM9refEncMasterPublicKey * pucSponsorTmpPublicKey, SM9refEncMasterPublicKey * pucResponseTmpPublicKey, SGD_UCHAR *pucHashSB, SGD_UINT32 *puiSBLen, SGD_UCHAR * pucHashS2, SGD_UINT32 *puiS2Len, void **phKeyHandle);

 SGD_RV  SDF_GenerateKeyWithSM9(SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyLen, SGD_UCHAR hid, SGD_UCHAR * pucSponsorID, SGD_UINT32 uiSponsorIDLen, SGD_UCHAR * pucResponseID, SGD_UINT32 uiResponseIDLen, SGD_UINT32 uiKeyIndex, SM9refEncUserPrivateKey   * pucSponsorPrivateKey, SM9refEncMasterPublicKey   *pucPublicKey, SM9refEncMasterPublicKey   * pucResponseTmpPublicKey, SGD_UCHAR *pucHashSB, SGD_UINT32 uiSBLen, SGD_UCHAR *pucHashSA, SGD_UINT32 *puiSALen, void *phAgreementHandle, void **phKeyHandle);

 SGD_RV  SDF_GenerateKeyVerifySM9(SGD_HANDLE hSessionHandle,SGD_UCHAR *pHashS2,SGD_UINT32  uiS2Len,SGD_UCHAR *pHashSA,SGD_UINT32 uiSALen);


 SGD_RV SDF_ExportPublicKey_DSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, DSArefPublicKey *pucPublicKey);
 SGD_RV SDF_InternalSign_DSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiISKIndex, SGD_UINT32 uiAlgID, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, SGD_UCHAR *pucSignature, SGD_UINT32 *uiSignatureDataLength);
 SGD_RV SDF_InternalVerify_DSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiISKIndex, SGD_UINT32 uiAlgID, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, SGD_UCHAR *pucSignature, SGD_UINT32 uiSignatureDataLength);
 SGD_RV SDF_GenerateKeyPair_DSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiAlgID, SGD_UINT32  uiKeyBits, DSArefPublicKey *pucPublicKey, DSArefPrivateKey *pucPrivateKey);
 SGD_RV SDF_ExternalSign_DSA(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, DSArefPrivateKey *pucPrivateKey, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, SGD_UCHAR *pucSignature, SGD_UINT32 *uiSignatureLength);
 SGD_RV SDF_ExternalVerify_DSA(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, DSArefPublicKey *pucPublicKey, SGD_UCHAR *pucDataInput, SGD_UINT32  uiInputLength, SGD_UCHAR *pucSignature, SGD_UINT32 uiSignatureLength);

 SGD_RV SDF_ExportSignPublicKey_ECDSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, ECDSArefPublicKey *pucPublicKey);
 SGD_RV SDF_ExportEncPublicKey_ECDSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, ECDSArefPublicKey *pucPublicKey);
 SGD_RV SDF_InternalSign_ECDSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiISKIndex, SGD_UINT32 uiAlgID, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, SGD_UCHAR *pucSignature, SGD_UINT32 *uiSignatureDataLength);
 SGD_RV SDF_InternalVerify_ECDSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiISKIndex, SGD_UINT32 uiAlgID, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, SGD_UCHAR *pucSignature, SGD_UINT32 uiSignatureDataLength);
 SGD_RV SDF_GenerateKeyPair_ECDSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiAlgID, SGD_UINT32  uiKeyBits, SGD_UINT32  uiCurveType, ECDSArefPublicKey *pucPublicKey, ECDSArefPrivateKey *pucPrivateKey);
 SGD_RV SDF_ExternalSign_ECDSA(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, ECDSArefPrivateKey *pucPrivateKey, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, SGD_UCHAR *pucSignature, SGD_UINT32 *uiSignatureLength);
 SGD_RV SDF_ExternalVerify_ECDSA(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, ECDSArefPublicKey *pucPublicKey, SGD_UCHAR *pucDataInput, SGD_UINT32  uiInputLength, SGD_UCHAR *pucSignature, SGD_UINT32 uiSignatureLength);
 SGD_RV SDF_GetKeyPair(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, SGD_UINT32 uiAlgID, ECDSArefPublicKey *pucPublicKey, ECDSArefPrivateKey *pucPrivateKey);
 SGD_RV SDF_InternalSign_ECDSA_Ex(SGD_HANDLE hSessionHandle, SGD_UINT32  uiISKIndex, SGD_UINT32 uiAlgID, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, ECDSASignature *pucSignature);
 SGD_RV SDF_InternalVerify_ECDSA_Ex(SGD_HANDLE hSessionHandle, SGD_UINT32  uiISKIndex, SGD_UINT32 uiAlgID, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, ECDSASignature *pucSignature);
 SGD_RV SDF_ExternalSign_ECDSA_Ex(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, ECDSArefPrivateKey *pucPrivateKey, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, ECDSASignature *pucSignature);
 SGD_RV SDF_ExternalVerify_ECDSA_Ex(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, ECDSArefPublicKey *pucPublicKey, SGD_UCHAR *pucDataInput, SGD_UINT32  uiInputLength, ECDSASignature *pucSignature);

 SGD_RV SDF_PBKDF2(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, SGD_UCHAR  *pucKey, SGD_UINT32  uiKeyLength, SGD_UCHAR  *pucSalt, SGD_UINT32  uiSaltLen, SGD_UINT32 uiCounter, SGD_UINT32 uiOutDataLen, SGD_UCHAR  *pucOutData);
 SGD_RV SDF_KeyAgreement_ECDH(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, SGD_UINT32 uiAlgID, ECDSArefPrivateKey *pucPrivateKey, ECDSArefPublicKey *pucPublicKey, SGD_UCHAR *pucShareKey, SGD_UINT32 *uiShareKeyLen);
 SGD_RV SDF_CMAC(SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle, SGD_UINT32 uiAlgID, SGD_UCHAR  *pucData, SGD_UINT32  uiDataLength, SGD_UCHAR  *pucCmac, SGD_UINT32 *puiCmacLen);

 SGD_RV SDF_Encrypt_CCM(SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle, SGD_UINT32 uiAlgID, SGD_UCHAR *pucIV, SGD_UINT32 uiIVLength, SGD_UCHAR *pucAAD, SGD_UINT32 uiAADLength, SGD_UCHAR *pucData, SGD_UINT32 uiDataLength, SGD_UCHAR *pucEncData, SGD_UINT32  *puiEncDataLength, SGD_UCHAR *pucTagData, SGD_UINT32  *puiTagDataLength);
 
 SGD_RV SDF_Decrypt_CCM(SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle, SGD_UINT32 uiAlgID, SGD_UCHAR *pucIV, SGD_UINT32 uiIVLength, SGD_UCHAR *pucAAD, SGD_UINT32 uiAADLength, SGD_UCHAR *pucTag, SGD_UINT32 uiTagLength, SGD_UCHAR *pucEncData, SGD_UINT32 puiEncDataLength, SGD_UCHAR *pucData, SGD_UINT32  *uiDataLength, SGD_UINT32  *puiResult);
 
 SGD_RV SDF_Encrypt_GCM(SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle, SGD_UINT32 uiAlgID, SGD_UCHAR *pucIV, SGD_UINT32 uiIVLength, SGD_UCHAR *pucAAD, SGD_UINT32 uiAADLength, SGD_UCHAR *pucData, SGD_UINT32 uiDataLength, SGD_UCHAR *pucEncData, SGD_UINT32  *puiEncDataLength, SGD_UCHAR *pucTagData, SGD_UINT32  *puiTagDataLength);
	
 SGD_RV SDF_Decrypt_GCM(SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle, SGD_UINT32 uiAlgID, SGD_UCHAR *pucIV, SGD_UINT32 uiIVLength, SGD_UCHAR *pucAAD, SGD_UINT32 uiAADLength, SGD_UCHAR *pucTag, SGD_UINT32 uiTagLength, SGD_UCHAR *pucEncData, SGD_UINT32 puiEncDataLength, SGD_UCHAR *pucData, SGD_UINT32  *uiDataLength, SGD_UINT32  *puiResult);				//输出，认证结果，1为认证通过，0为认证失败		

 SGD_RV SDF_Encrypt_XTS(SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle, SGD_HANDLE hKeyHandleSec, SGD_UINT32 uiAlgID, SGD_UCHAR *pucIV, SGD_UCHAR *pucData, SGD_UINT32 uiDataLength, SGD_UCHAR *pucEncData, SGD_UINT32  *puiEncDataLength);
 
 SGD_RV SDF_Decrypt_XTS(SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle, SGD_HANDLE hKeyHandleSec, SGD_UINT32 uiAlgID, SGD_UCHAR *pucIV, SGD_UCHAR *pucEncData, SGD_UINT32  uiEncDataLength, SGD_UCHAR *pucData, SGD_UINT32 *puiDataLength);

 SGD_RV SDF_GenerateKeyPair_EdDSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiAlgID, SGD_UINT32  uiKeyBits, SGD_UINT32  uiCurveType, EdDSArefPublicKey* pucPublicKey, EdDSArefPrivateKey* pucPrivateKey);
 SGD_RV SDF_ExportPublicKey_EdDSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, EdDSArefPublicKey* pucPublicKey);
 SGD_RV SDF_InternalSign_EdDSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, SGD_UINT32 uiAlgID, SGD_UCHAR* pucData, SGD_UINT32  uiDataLength, SGD_UCHAR* pucSignature, SGD_UINT32* uiSignatureDataLength);
 SGD_RV SDF_InternalVerify_EdDSA(SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex, SGD_UINT32 uiAlgID, SGD_UCHAR* pucData, SGD_UINT32  uiDataLength, SGD_UCHAR* pucSignature, SGD_UINT32 uiSignatureDataLength);
 SGD_RV SDF_ExternalSign_EdDSA(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, EdDSArefPrivateKey* pucPrivateKey, SGD_UCHAR* pucData, SGD_UINT32  uiDataLength, SGD_UCHAR* pucSignature, SGD_UINT32* uiSignatureLength);
 SGD_RV SDF_ExternalVerify_EdDSA(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, EdDSArefPublicKey* pucPublicKey, SGD_UCHAR* pucData, SGD_UINT32  uiDataLength, SGD_UCHAR* pucSignature, SGD_UINT32 uiSignatureLength);

 SGD_RV SDF_GetIV_ZUC(SGD_HANDLE hSessionHandle, SGD_UINT32  uiAlgID, SGD_UINT32  uiCount, SGD_UCHAR ucBearer, SGD_UCHAR ucDirection, SGD_UCHAR* pIv);
 SGD_RV SDF_Confidentiality_ZUC(SGD_HANDLE hSessionHandle, SGD_UCHAR* pucCKey, SGD_UCHAR *pIv, SGD_UCHAR* pucIBS, SGD_UINT32  uiLength, SGD_UCHAR* pucOBS);
 SGD_RV SDF_Integrity_ZUC(SGD_HANDLE hSessionHandle, SGD_UCHAR* pucIKey, SGD_UCHAR *pIv, SGD_UCHAR* pucM, SGD_UINT32 uiLength, SGD_UINT32* puiMAC);
 
 SGD_RV SDF_ExternalEncrypt_BigData_ECC(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, ECCrefPublicKey *pucPublicKey, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, ECCCipher *pucEncData);
 SGD_RV SDF_ExternalDecrypt_BigData_ECC(SGD_HANDLE hSessionHandle, SGD_UINT32 uiAlgID, ECCrefPrivateKey *pucPrivateKey, ECCCipher *pucEncData, SGD_UCHAR *pucData, SGD_UINT32  *puiDataLength);
 SGD_RV SDF_InternalEncrypt_BigData_ECC(SGD_HANDLE hSessionHandle, SGD_UINT32  uiISKIndex, SGD_UINT32 uiAlgID, SGD_UCHAR *pucData, SGD_UINT32  uiDataLength, ECCCipher *pucEncData);
 SGD_RV SDF_InternalDecrypt_BigData_ECC(SGD_HANDLE hSessionHandle, SGD_UINT32  uiISKIndex, SGD_UINT32 uiAlgID, ECCCipher *pucEncData, SGD_UCHAR *pucData, SGD_UINT32  *puiDataLength);

 #ifdef __cplusplus
}
#endif

#endif /*#ifndef _SW_SDS_H_*/