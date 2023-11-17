package pkcs11impl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/miekg/pkcs11"
)

//softhsm2-util --init-token --slot 0 --label test --pin 1234

var lib = "./libs/libsofthsm2.so"

func TestNew(t *testing.T) {
	_ = pkcs11.New("/usr/local/Cellar/softhsm/2.6.1/lib/softhsm/libsofthsm2.so")
}

func TestSHA1(t *testing.T) {
	p := pkcs11.New(lib)
	assert.NotNil(t, p)

	err := p.Initialize()
	assert.NoError(t, err)

	defer p.Destroy()
	defer p.Finalize()

	slots, err := p.GetSlotList(true)
	assert.NoError(t, err)
	assert.NotNil(t, slots) //must initialize softhsm locally. CMD: softhsm2-util --init-token --slot 0 --label test --pin 1234

	session, err := p.OpenSession(slots[0], pkcs11.CKF_SERIAL_SESSION|pkcs11.CKF_RW_SESSION)
	assert.NoError(t, err)
	defer p.CloseSession(session)

	err = p.Login(session, pkcs11.CKU_USER, "1234")
	assert.NoError(t, err)

	defer p.Logout(session)

	p.DigestInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_SHA_1, nil)})
	hash, err := p.Digest(session, []byte("this is a string"))
	assert.NoError(t, err)

	for _, d := range hash {
		fmt.Printf("%x", d)
	}
	fmt.Println()
}

func TestFindSlotLabel(t *testing.T) {
	ctx := pkcs11.New(lib)
	assert.NotNil(t, ctx)

	err := ctx.Initialize()
	assert.NoError(t, err)

	slots, err := ctx.GetSlotList(true)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(slots))

	for i, s := range slots {
		info, err := ctx.GetTokenInfo(s)
		if err != nil {
			continue
		}
		t.Logf("slot[%d], lable[%s]\n", i, info.Label)
	}
}
