package swsds

const (
	//SGD_RSA_SIGN 加签方式标识
	SGD_RSA_SIGN CTypeKeyUsage = 0x00010100
	/*常量定义*/
	SGD_TRUE  CTypeAlgorithm = 0x00000001
	SGD_FALSE CTypeAlgorithm = 0x00000000

	/*对称密码算法标识*/
	SGD_SM1_ECB   CTypeAlgorithm = 0x00000101
	SGD_SM1_CBC   CTypeAlgorithm = 0x00000102
	SGD_SM1_CFB   CTypeAlgorithm = 0x00000104
	SGD_SM1_OFB   CTypeAlgorithm = 0x00000108
	SGD_SM1_MAC   CTypeAlgorithm = 0x00000110
	SGD_SM1_CTR   CTypeAlgorithm = 0x00000120
	SGD_SMS4_ECB  CTypeAlgorithm = 0x00000401
	SGD_SMS4_CBC  CTypeAlgorithm = 0x00000402
	SGD_SMS4_CFB  CTypeAlgorithm = 0x00000404
	SGD_SMS4_OFB  CTypeAlgorithm = 0x00000408
	SGD_SMS4_MAC  CTypeAlgorithm = 0x00000410
	SGD_SMS4_CTR  CTypeAlgorithm = 0x00000420
	SGD_SMS4_GCM  CTypeAlgorithm = 0x00000440
	SGD_SMS4_XTS  CTypeAlgorithm = 0x00000480
	SGD_SMS4_CCM  CTypeAlgorithm = 0x000004A0
	SGD_SMS4_CMAC CTypeAlgorithm = 0x000004C0

	MAX_INT int = 2147483647
)
