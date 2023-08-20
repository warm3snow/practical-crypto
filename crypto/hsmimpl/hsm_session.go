/**
 * @Author: xueyanghan
 * @File: hsm_session.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/20 14:22
 */

package hsmimpl

import (
	"github.com/pkg/errors"
	"github.com/warm3snow/practical-crypto/crypto/hsmimpl/base"
	"time"
)

// getSession get hsm session
func (h *hsmimpl) getSession() (base.SessionHandle, error) {
	var session base.SessionHandle
	select {
	case session = <-h.sessions:
		return session, nil
	default:
		var err error
		for i := 0; i < 3; i++ {
			session, err = h.ctx.SDFOpenSession(h.deviceHandle)
			if err == nil {
				return session, nil
			}
			time.Sleep(time.Millisecond * 100)
		}
		return nil, errors.WithMessage(err, "failed to create new session after 3 times attempt")
	}
}

// returnSession return hsm session
func (h *hsmimpl) returnSession(err error, session base.SessionHandle) {
	if err != nil {
		_ = h.ctx.SDFCloseSession(session)
	}
	select {
	case h.sessions <- session:
		return
	default:
		_ = h.ctx.SDFCloseSession(session)
		return
	}
}
