/**
 * @Author: xueyanghan
 * @File: dbaccess_service_redis_subscribe.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/5/11 12:05
 */

package dbaccess_impl

import "github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/types"

func (dbaccess DBAccessServiceImpl) AddRedisSubscribe(redisSubscribe *types.RedisSubscribe) error {
	//TODO implement me
	panic("implement me")
}

func (dbaccess DBAccessServiceImpl) DeleteRedisSubscribe(messageId string) error {
	//TODO implement me
	panic("implement me")
}

func (dbaccess DBAccessServiceImpl) UpdateRedisSubscribe(redisSubscribe *types.RedisSubscribe) error {
	//TODO implement me
	panic("implement me")
}

func (dbaccess DBAccessServiceImpl) GetRedisSubscribe(id int) (*types.RedisSubscribe, error) {
	//TODO implement me
	panic("implement me")
}

func (dbaccess DBAccessServiceImpl) GetRedisSubscribeList(status []int) ([]*types.RedisSubscribe, error) {
	//TODO implement me
	panic("implement me")
}

func (dbaccess DBAccessServiceImpl) UpdateRedisSubscribeStatus(id int, status int) error {
	//TODO implement me
	panic("implement me")
}

func (dbaccess DBAccessServiceImpl) GetRedisSubscribeByMessageId(messageId string) (*types.RedisSubscribe, error) {
	//TODO implement me
	panic("implement me")
}

func (dbaccess DBAccessServiceImpl) UpdateRedisSubscribeStatusByMessageId(messageId string, status int) error {
	//TODO implement me
	panic("implement me")
}
