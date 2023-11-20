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
	//algo string, keySize int
	switch strings.ToUpper(algo) {
	case "SM4":
		session, err := h.getSession()
		if err != nil {
			return nil, err
		}
		defer h.returnSession(err, session)
		return base.SM4GenKey(h.ctx, session, 16)
	default:
		return nil, errors.New("Only support SM4 now")
	}
}

func (h *hsmimpl) Enc(algo, key string, plainText []byte, mode string) ([]byte, error) {
	session, err := h.getSession()
	if err != nil {
		return nil, err
	}
	defer h.returnSession(err, session)

	switch strings.ToUpper(algo) {
	case "SM4":
		// check keyId
		keyIndex, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}
		return base.SM4Encrypt(h.ctx, session, uint(keyIndex), plainText, mode)
	default:
		return nil, errors.New("Only support SM4 now")
	}
}

func (h *hsmimpl) Dec(algo, key string, cipherText []byte, mode string) ([]byte, error) {
	session, err := h.getSession()
	if err != nil {
		return nil, err
	}
	defer h.returnSession(err, session)

	switch strings.ToUpper(algo) {
	case "SM4":
		// check keyId
		keyIndex, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}
		return base.SM4Decrypt(h.ctx, session, uint(keyIndex), cipherText, mode)
	default:
		return nil, errors.New("Only support SM4 now")
	}
}

func (h *hsmimpl) Sign(algo, key string, plain []byte, option ...[]byte) ([]byte, error) {
	session, err := h.getSession()
	if err != nil {
		return nil, err
	}
	defer h.returnSession(err, session)

	switch strings.ToUpper(algo) {
	case "SM2":
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
	session, err := h.getSession()
	if err != nil {
		return false, err
	}
	defer h.returnSession(err, session)

	switch strings.ToUpper(algo) {
	case "SM2":
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
	session, err := h.getSession()
	if err != nil {
		return nil, err
	}
	defer h.returnSession(err, session)

	switch strings.ToUpper(algo) {
	case "SM3":
		return base.SM3Hash(h.ctx, session, origin)
	default:
		return nil, errors.New("Only support SM3 now")
	}
}

func (h hsmimpl) HMac(algo, key string, plain []byte) ([]byte, error) {
	session, err := h.getSession()
	if err != nil {
		return nil, err
	}
	defer h.returnSession(err, session)

	switch strings.ToUpper(algo) {
	case "SM3":
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
