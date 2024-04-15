/**
 * @Author: xueyanghan
 * @File: utils.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/15 18:31
 */

package pre

import "C"
import (
	"crypto/elliptic"
	"github.com/warm3snow/gmsm/sm2"
	"github.com/warm3snow/gmsm/sm3"
	"math/big"
)

var (
	Curve = sm2.P256Sm2()
	N     = sm2.P256Sm2().Params().N
)

type CurvePoint = sm2.PublicKey

func CurvePointToBytes(point *CurvePoint) []byte {
	return elliptic.MarshalCompressed(Curve, point.X, point.Y)
}

func Hash(x []byte) []byte {
	return sm3.Sm3Sum(x)
}

func HashToCurve(h []byte) *big.Int {
	hashInt := new(big.Int).SetBytes(h)
	return hashInt.Mod(hashInt, N)
}
