/**
 * @Author: xueyanghan
 * @File: wechat_pay.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/5/11 16:34
 */

package apps

import (
	"context"
	"crypto/rsa"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
)

func WechatPay(mchID, mchCertificateSerialNumber, mchAPIv3Key string, mchPrivateKey *rsa.PrivateKey) (string, error) {
	ctx := context.Background()

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
		return "", err
	}
	// 以 Native 支付为例
	svc := native.NativeApiService{Client: client}
	// 发送请求
	resp, result, err := svc.Prepay(ctx,
		native.PrepayRequest{
			Appid:       core.String("wxd678efh567hg6787"),
			Mchid:       core.String("1900009191"),
			Description: core.String("Image形象店-深圳腾大-QQ公仔"),
			OutTradeNo:  core.String("1217752501201407033233368018"),
			Attach:      core.String("自定义数据说明"),
			NotifyUrl:   core.String("https://www.weixin.qq.com/wxpay/pay.php"),
			Amount: &native.Amount{
				Total: core.Int64(100),
			},
		},
	)
	if err != nil {
		log.Fatalf("prepay err:%s", err)
		return "", err
	}

	// 使用微信扫描 resp.code_url 对应的二维码，即可体验Native支付
	log.Printf("status=%d resp=%v", result.Response.StatusCode, resp)

	return *resp.CodeUrl, nil
}
