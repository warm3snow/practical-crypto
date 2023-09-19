/**
 * @Author: xueyanghan
 * @File: kdf.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/19 14:15
 */

package kdf

type KDF interface {
	DeriveKeyByPassword(password string) (deriveKey []byte, err error)
	VerifyDeriveKeyStr(kdfKeyStr string, password []byte) (isOk bool, err error)
	GetDeriveKeyStr() string
}
