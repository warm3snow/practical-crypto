/**
 * @Author: xueyanghan
 * @File: types.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/13 22:56
 */

package pre

import (
	"github.com/warm3snow/gmsm/sm2"
	"math/big"
)

type Capsule struct {
	E sm2.PublicKey
	V sm2.PublicKey
	S *big.Int
}

const (
	C1C2 = 1
)
