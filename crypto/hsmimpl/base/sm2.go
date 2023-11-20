/**
 * @Author: xueyanghan
 * @File: sm2.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/11/17 14:34
 */

package base

import (
	"encoding/asn1"
	"github.com/warm3snow/gmsm/sm2"
	"github.com/warm3snow/gmsm/sm3"
	"math/big"
)

const (
	CRYPTO_DEFAULT_UID = "1234567812345678"
)

type eccSignature struct {
	R, S *big.Int
}

func SM2Sign(c *Ctx, s SessionHandle, keyIndex uint, keyPwd, origin []byte) (sign []byte, err error) {
	err = c.SDFGetPrivateKeyAccessRight(s, keyIndex+10000, keyPwd, uint(len(keyPwd)))
	if err != nil {
		return nil, err
	}

	//defer func() {
	//	err = c.SDFReleasePrivateKeyAccessRight(s, keyIndex+10000)
	//	if err != nil {
	//		log.Fatalf("failed to SDFReleasePrivateKeyAccessRight, err = %v", err)
	//	}
	//}()

	// must use sm3 to hash origin
	pub, err := ExportECDSAPublicKey(c, s, keyIndex)
	if err != nil {
		return nil, err
	}
	originHash, err := calDataHash(pub, origin, CRYPTO_DEFAULT_UID)
	if err != nil {
		return nil, err
	}

	signature, err := c.SDFInternalSign_ECC(s, keyIndex, originHash, uint(len(originHash)))
	if err != nil {
		return nil, err
	}

	eccR := big.NewInt(0).SetBytes([]byte(signature.R))
	eccS := big.NewInt(0).SetBytes([]byte(signature.S))

	return asn1.Marshal(eccSignature{eccR, eccS})
}

func SM2Verify(c *Ctx, s SessionHandle, keyIndex uint, origin, sign []byte) (bool, error) {
	var signature eccSignature
	_, err := asn1.Unmarshal(sign, &signature)
	if err != nil {
		return false, err
	}

	// must use sm3 to hash origin
	pub, err := ExportECDSAPublicKey(c, s, keyIndex)
	if err != nil {
		return false, err
	}

	//originHash, err := calDataHash(pub, origin, CRYPTO_DEFAULT_UID)
	//if err != nil {
	//	return false, err
	//}
	//var eccSinagure ECCSignature
	//eccSinagure.R = string(signature.R.Bytes())
	//eccSinagure.S = string(signature.S.Bytes())

	//err = c.SDFInternalVerify_ECC(s, keyIndex, originHash, uint(len(originHash)), eccSinagure)
	//if err != nil {
	//	return false, err
	//}
	//return true, nil

	return sm2.Sm2Verify(pub, origin, []byte(CRYPTO_DEFAULT_UID), signature.R, signature.S), nil
}

// ExportECDSAPublicKey export a ecc publickey
func ExportECDSAPublicKey(c *Ctx, s SessionHandle, keyIndex uint) (*sm2.PublicKey, error) {
	pub, err := c.SDFExportSignPublicKey_ECC(s, keyIndex)
	if err != nil {
		return nil, err
	}

	x, y := big.NewInt(0), big.NewInt(0)
	x.SetBytes([]byte(pub.X))
	y.SetBytes([]byte(pub.Y))

	sm2PubKey := &sm2.PublicKey{
		Curve: sm2.P256Sm2(),
		X:     x,
		Y:     y,
	}
	return sm2PubKey, err
}

func calDataHash(pkSM2 *sm2.PublicKey, origin []byte, uid string) ([]byte, error) {
	if len(uid) == 0 {
		uid = CRYPTO_DEFAULT_UID
	}
	za, err := sm2.ZA(pkSM2, []byte(uid))
	if err != nil {
		return nil, err
	}
	e := sm3.New()
	e.Write(za)
	e.Write(origin)
	return e.Sum(nil)[:32], nil
}
