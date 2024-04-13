/**
 * @Author: xueyanghan
 * @File: types.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/13 22:56
 */

package pre

import "math/big"

type C struct {
	C2 []byte `json:"C2"`
	C1 struct {
		X *big.Int `json:"X"`
		Y *big.Int `json:"Y"`
	}
	C3 []byte `json:"C3"`
	C4 []byte `json:"C4"`
}

const (
	C2C1C3C4 = 1
)
