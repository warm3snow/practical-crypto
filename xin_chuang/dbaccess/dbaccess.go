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
	// AddRedisSubscribe adds redis subscribe.
	AddRedisSubscribe(redisSubscribe *types.RedisSubscribe) error

	// DeleteRedisSubscribe deletes redis subscribe of specified RedisSubscribeNo.
	DeleteRedisSubscribe(messageId string) error

	// UpdateRedisSubscribe updates redis subscribe.
	UpdateRedisSubscribe(redisSubscribe *types.RedisSubscribe) error

	// GetRedisSubscribe gets redis subscribe of specified RedisSubscribeNo.
	GetRedisSubscribe(id int) (*types.RedisSubscribe, error)

	// GetRedisSubscribeList gets redis subscribe list by status.
	GetRedisSubscribeList(status []int) ([]*types.RedisSubscribe, error)

	// UpdateRedisSubscribeStatus updates redis subscribe status.
	UpdateRedisSubscribeStatus(id int, status int) error

	// GetRedisSubscribeByMessageId gets redis subscribe of specified messageId.
	GetRedisSubscribeByMessageId(messageId string) (*types.RedisSubscribe, error)

	// UpdateRedisSubscribeStatusByMessageId updates redis subscribe of specified messageId.
	UpdateRedisSubscribeStatusByMessageId(messageId string, status int) error
}

// InitDBAccessService returns a new database instance.
func InitDBAccessService(cfg *config.DBConfig) (err error) {
	DBAccessService, err = dbaccess_impl.InitDBAccessService(cfg)
	if err != nil {
		return err
	}
	return nil
}
