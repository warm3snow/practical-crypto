package hsmimpl

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"github.com/warm3snow/gmsm/sm3"
	"log"
	"testing"
)

func TestHsmSM3(t *testing.T) {
	t.Skip()

	origin := []byte("hello")

	c, err := New("./lib/libswsds.dylib")
	assert.NoError(t, err)

	v, err := c.Hash("SM3", origin)
	assert.NoError(t, err)

	assert.Equal(t, v, sm3.Sm3Sum(origin))
}

func TestHsmSM3HMac(t *testing.T) {
	origin := []byte("hello")

	c, err := New("./lib/libswsds.dylib")
	assert.NoError(t, err)

	v, err := c.HMac("SM3", "1", origin)
	assert.NoError(t, err)

	log.Println(hex.EncodeToString(v))
}
