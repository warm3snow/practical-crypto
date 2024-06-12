/**
 * @Author: xueyanghan
 * @File: dbaccess_interface.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/5/11 11:58
 */

package dbaccess

import (
	"github.com/warm3snow/practical-crypto/xin_chuang/config"
	"github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/dbaccess_impl"
	"github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/types"
)

var (
	// DBAccessService - database access service
	DBAccessService DBAccessInterface
)

// DBAccessInterface -
type DBAccessInterface interface {
	// GetUserInfo - get user info by user name
	GetUserInfo(userName string) (userInfo *types.UserInfo, err error)
	// AddUser - add user
	AddUser(user *types.UserInfo) (err error)
	// DeleteUser - delete user
	DeleteUser(userName string) (err error)
	// UpdateUser - update user
	UpdateUser(user *types.UserInfo) (err error)
	// GetUsers - get all users
	GetUsers() (users []*types.UserInfo, err error)
}

// InitDBAccessService returns a new database instance.
func InitDBAccessService(cfg *config.DBConfig) (err error) {
	DBAccessService, err = dbaccess_impl.InitDBAccessService(cfg)
	if err != nil {
		return err
	}
	return nil
}
