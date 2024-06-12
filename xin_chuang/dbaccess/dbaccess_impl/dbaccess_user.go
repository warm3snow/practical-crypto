/**
 * @Author: xueyanghan
 * @File: dbaccess_service_redis_subscribe.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/5/11 12:05
 */

package dbaccess_impl

import (
	"github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/model"
	"github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/types"
)

func (d DBAccessServiceImpl) GetUserInfo(userName string) (userInfo *types.UserInfo, err error) {
	user, err := d.dao.GetUser(userName)
	if err != nil {
		return nil, err
	}
	userInfo = &types.UserInfo{
		UserName: user.UserName,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
		Role:     user.Role,
	}
	return userInfo, nil
}

func (d DBAccessServiceImpl) AddUser(user *types.UserInfo) (err error) {
	userModel := &model.User{
		UserName: user.UserName,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
		Role:     user.Role,
	}
	return d.dao.AddUser(userModel)
}

func (d DBAccessServiceImpl) DeleteUser(userName string) (err error) {
	return d.dao.DeleteUser(userName)
}

func (d DBAccessServiceImpl) UpdateUser(user *types.UserInfo) (err error) {
	userModel := &model.User{
		UserName: user.UserName,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
		Role:     user.Role,
	}
	return d.dao.UpdateUser(userModel)
}

func (d DBAccessServiceImpl) GetUsers() (users []*types.UserInfo, err error) {
	userModels, err := d.dao.GetUsers()
	if err != nil {
		return nil, err
	}
	users = make([]*types.UserInfo, 0)
	for _, userModel := range userModels {
		user := &types.UserInfo{
			UserName: userModel.UserName,
			Password: userModel.Password,
			Email:    userModel.Email,
			Phone:    userModel.Phone,
			Role:     userModel.Role,
		}
		users = append(users, user)
	}
	return users, nil
}
