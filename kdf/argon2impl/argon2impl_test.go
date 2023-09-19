/**
 * @Author: xueyanghan
 * @File: argon2impl_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/19 15:26
 */

package pbkdf2impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	argon2impl *Argon2impl
)

func TestMain(m *testing.M) {
	argon2impl = New(-1, -1, -1, -1)
	m.Run()
}

func TestArgon2(t *testing.T) {
	_, err := argon2impl.DeriveKeyByPassword("123456")
	assert.NoError(t, err)

	dkStr1 := argon2impl.GetDeriveKeyStr()
	dkStr2 := argon2impl.GetDeriveKeyStr()
	assert.Equal(t, dkStr1, dkStr2)
	t.Log("DeriveKeyStr:", dkStr1)

	isOk, err := argon2impl.VerifyDeriveKeyStr(dkStr1, []byte("123456"))
	assert.NoError(t, err)
	assert.True(t, isOk)
}
