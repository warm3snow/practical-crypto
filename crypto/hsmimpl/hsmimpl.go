package hsmimpl

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/warm3snow/practical-crypto/crypto/hsmimpl/base"
	"os"
	"strconv"
	"strings"
)

const (
	sessionCacheSize = 10
)

type hsmimpl struct {
	ctx          *base.Ctx
	deviceHandle base.SessionHandle
	sessions     chan base.SessionHandle
}

func (h *hsmimpl) GenSymKey(algo string, keySize int) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (h *hsmimpl) Enc(algo string, key, plain []byte, mode string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (h *hsmimpl) Dec(algo string, key, plain []byte, mode string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (h *hsmimpl) Sign(algo, key string, plain []byte, option ...[]byte) ([]byte, error) {
	switch strings.ToUpper(algo) {
	case "SM2":
		session, err := h.getSession()
		if err != nil {
			return nil, err
		}
		defer h.returnSession(err, session)
		// check keyId
		keyIndex, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}
		// get keyPwd
		var keyPwd []byte
		if len(option) > 0 {
			keyPwd = option[0]
		}
		return base.SM2Sign(h.ctx, session, uint(keyIndex), keyPwd, plain)
	default:
		return nil, errors.New("Only support SM2 now")
	}
}

func (h *hsmimpl) Verify(algo, key string, plain, sig []byte) (bool, error) {
	switch strings.ToUpper(algo) {
	case "SM2":
		session, err := h.getSession()
		if err != nil {
			return false, err
		}
		defer h.returnSession(err, session)
		// check keyId
		keyIndex, err := strconv.Atoi(key)
		if err != nil {
			return false, err
		}
		return base.SM2Verify(h.ctx, session, uint(keyIndex), plain, sig)
	default:
		return false, errors.New("Only support SM2 now")
	}
}

func New(lib string) (*hsmimpl, error) {
	ctx := base.New(lib)
	if ctx == nil {
		libEnv := os.Getenv("HSM_LIB")
		ctx = base.New(libEnv)
		if ctx == nil {
			return nil, fmt.Errorf("[SDF] error: fail to initialize [%s]", libEnv)
		}
	}

	var err error
	var deviceHandle base.SessionHandle
	for i := 0; i < 3; i++ {
		deviceHandle, err = ctx.SDFOpenDevice()
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		return nil, fmt.Errorf("[SDF] error: fail to open device after 3 times [%v]", err)
	}
	sessions := make(chan base.SessionHandle, sessionCacheSize)
	handle := &hsmimpl{
		ctx:          ctx,
		deviceHandle: deviceHandle,
		sessions:     sessions,
	}
	return handle, nil
}

func (h hsmimpl) Hash(algo string, origin []byte) ([]byte, error) {
	switch strings.ToUpper(algo) {
	case "SM3":
		session, err := h.getSession()
		if err != nil {
			return nil, err
		}
		defer h.returnSession(err, session)
		return base.SM3Hash(h.ctx, session, origin)
	default:
		return nil, errors.New("Only support SM3 now")
	}
}

func (h hsmimpl) HMac(algo, key string, plain []byte) ([]byte, error) {
	switch strings.ToUpper(algo) {
	case "SM3":
		session, err := h.getSession()
		if err != nil {
			return nil, err
		}
		defer h.returnSession(err, session)
		//check keyId
		keyIndex, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}
		return base.SM3HMac(h.ctx, session, uint(keyIndex), plain)
	default:
		return nil, errors.New("Only support SM3 now")
	}
}
