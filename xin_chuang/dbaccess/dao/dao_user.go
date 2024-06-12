/**
 * @Author: xueyanghan
 * @File: dao_user.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/6/12 11:20
 */

package dao

import (
	"github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/model"
)

func (d *DAO) GetUser(userName string) (userInfo *model.User, err error) {
	result := d.db.Where("user_name = ?", userName).First(&userInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return userInfo, nil
}

func (d *DAO) AddUser(user *model.User) (err error) {
	result := d.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *DAO) DeleteUser(userName string) (err error) {
	result := d.db.Where("user_name = ?", userName).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *DAO) UpdateUser(user *model.User) (err error) {
	result := d.db.Model(&model.User{}).Where("user_name = ?", user.UserName).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *DAO) GetUsers() (users []*model.User, err error) {
	users = make([]*model.User, 0)
	result := d.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
