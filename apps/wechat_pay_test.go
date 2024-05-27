/**
 * @Author: xueyanghan
 * @File: wechat_pay_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/5/11 16:42
 */

package apps

import (
	"log"
	"testing"

	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

const (
	privateKeyPem = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCytm4ia0YUj3kd
ELI2+qmme4CA7X9FzkkjcyC3hx/QzUcSSk6CRSHhLZR4LbdWZ2cDZSJPc63eY4/W
nncVBt50YnX4/lvpFftpA7MZevXeXk79/BtI5zM36ZDD5MNuErue+6qjJyhrCrzO
19alXCM7Twa4xx98yy1mpA6MynZ3Fmq9JdxplsJ9Z1j55KEYTohRM7p/oL0hIOSZ
qi1KTgGZkauYoGOtCdmi1xjSGvpkdhL2B5dGmHkPbl2g03YNHxMkjIylIiQNpSB5
NV00YT9MMJfcZVb+wUk46qc9sQ0Uojt0DgHNJ3u+mG56N/kQ61HJtbhpPgs8qH5Y
bFKc2wetAgMBAAECggEABHVerAcnGW425E29VHfOcbPVEeJa8X/i/TprD4kUdx4X
7tHhwYg9/yMm+BST5VjxDfmg8IB/TvQG66DKccDZHegy70JRQrBktEr+Vd1mcuTB
bb6zY485DOUL0odLhNx8x8uZhzZ74KfE7BggZ8Gzs1AWZ9e35mIfQNdI8aQ/blsb
AwVJxMkm0jgkk9FN7vCThmgJYAmpsO2XmtKHzJWWRXfL6s9g+kg6HZNPEKbM0CYF
yix01cBG3j51LT9bUYZhYvvKZWb90dQ3iC+7YmCCg/XlqDJ203kEA5M7RgOjf/OP
s6zbnImm01KqrTmNcuM7oMiyMyyLrc+mtTPT9QhKNQKBgQD6EqKLMz/eMT2V8sgU
CdndVVdxudVRb9bxwcHv+NoR/cYPnPfb9v63LZcsBV+01+wqx2FpvpQq8lJENrmC
P6xHlSgYtmVZMq9SXH+y9T/O3/nkbO6t5XNAmxykt8gkona4RF5zIh73YXx44PaG
XAC+bCeYwxtR4E7vHe7ejwHCdwKBgQC28s25KpvYpyOH2UDqVj5rkt4KcbU6v6Y+
j4xjqTsfMuEoBcoSBRTvwjfW4qxTBcDf2l+Pq9rfVWz7jgezeAFy4bkZtxe6SVb7
TJSJ0Yt1C9sCBXF9bKChByQmVxuOyvA5jxvcvK9eef+ej1sFWTaXHt/btz39umZR
P2qjrLvL+wKBgQDMFds7uUadYhg9wNaSq/t2i8iCksPJ43lK8fMBklj+J9q1MUiC
5s+31Ogzz+rsnu8pdnBEqXkol3yWGQdKigZioRMJIAoUQq4cjErXVRmPDvFU/6XG
m6R0jGJS933KkCBNf9aJJcaZ46B+0jG/M2SUA5ZJMKmiJ2qj8QzZ62HnPwKBgF20
v+q4CFtQQMK0ZKuiwRYYg+KVm14cC+q/XAkwBThVtz3ouTfBoperwy4trFZZ/BCs
qYXNYK5D8y5l8Uvbi+Jr+4NQZLbmGcdd0jdFPUkWaXb/ksHLgfr4zWtV8qeRCrpw
srS9cJXTpzpv8w8qQuvkxISltvrfHsk+0kBDmW1BAoGBALoS+UlDDPlbSYIdP/VE
Bl3airb02yC1s0HMohue9FvQxfFqq/d7kERkgiggChsC0EvZOQaTO/KP9OUkLK79
MMrEJ5emwD4sVuAxFEX5BUpFOf5j/0TQv6drFgKfcFy7+x+Uf/1NLdVQf+wZXXv4
RonoLPvDQADoSeVVJowb/Tkf
-----END PRIVATE KEY-----
`
)

func TestWechatPay(t *testing.T) {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775B6A45ACD588826D15E583A95F5DD********" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	//mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	mchPrivateKey, err := utils.LoadPrivateKey(privateKeyPem)
	if err != nil {
		t.Errorf("load private key error:%s", err)
	}

	codeUrl, err := WechatPay(mchID, mchCertificateSerialNumber, mchAPIv3Key, mchPrivateKey)
	if err != nil {
		t.Error(err)
	}

	log.Println("codeUrl:", codeUrl)
}
