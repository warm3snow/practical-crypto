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
)

// getSession get hsm session
func (h *hsmimpl) getSession() (base.SessionHandle, error) {
	var session base.SessionHandle
	select {
	case session = <-h.sessions:
		return session, nil
	default:
		if len(h.sessions) == 0 && len(h.sessions) < sessionCacheSize {
			var err error
			session, err = h.ctx.SDFOpenSession(h.deviceHandle)
			if err != nil {
				return nil, errors.WithMessage(err, "failed to create new session after 3 times attempt")
			}
			h.sessions <- session
		}
	}
	return <-h.sessions, nil
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
