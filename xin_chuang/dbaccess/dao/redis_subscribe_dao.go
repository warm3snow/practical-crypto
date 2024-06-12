/**
 * @Author: xueyanghan
 * @File: redis_subscribe_dao.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/30 12:01
 */

package dao

import "github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/model"

func (dao *DAO) AddRedisSubscribe(redisSubscribe *model.RedisSubscribe) error {
	return dao.db.Create(redisSubscribe).Error
}

func (dao *DAO) DeleteRedisSubscribe(id int) error {
	return dao.db.Delete(&model.RedisSubscribe{}, "id = ?", id).Error
}

func (dao *DAO) UpdateRedisSubscribe(redisSubscribe *model.RedisSubscribe) error {
	// note: gorm will update all fields, even if they are not changed
	return dao.db.Save(redisSubscribe).Error
}

func (dao *DAO) GetRedisSubscribe(id int) (*model.RedisSubscribe, error) {
	var redisSubscribe model.RedisSubscribe
	err := dao.db.Where("id = ?", id).
		First(&redisSubscribe).Error
	return &redisSubscribe, err
}

func (dao *DAO) UpdateRedisSubscribeStatus(id int, status int) error {
	return dao.db.Model(&model.RedisSubscribe{}).
		Where("id = ?", id).
		Update("message_status", status).Error
}

func (dao *DAO) GetRedisSubscribeByMessageId(messageId string) (*model.RedisSubscribe, error) {
	var redisSubscribe model.RedisSubscribe
	err := dao.db.Where("message_id = ?", messageId).
		First(&redisSubscribe).Error
	return &redisSubscribe, err
}

func (dao *DAO) UpdateRedisSubscribeStatusByMessageId(messageId string, status int) error {
	return dao.db.Model(&model.RedisSubscribe{}).
		Where("message_id = ?", messageId).
		Update("message_status", status).Error
}

func (dao *DAO) GetRedisSubscribeList(status []int) ([]*model.RedisSubscribe, error) {
	var redisSubscribeList = make([]*model.RedisSubscribe, 0)
	var err error
	if status == nil {
		err = dao.db.Model(&model.RedisSubscribe{}).Find(&redisSubscribeList).Error
	} else {
		err = dao.db.Model(&model.RedisSubscribe{}).
			Where("message_status IN (?)", status).
			Find(&redisSubscribeList).Error
	}
	return redisSubscribeList, err
}
