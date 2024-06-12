/**
 * @Author: xueyanghan
 * @File: dbaccess_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/26 19:11
 */

package dbaccess

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warm3snow/practical-crypto/xin_chuang/config"
	"github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/dao"
	"github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/types"
)

func TestMain(m *testing.M) {
	mysqldsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123456", "localhost", 3306, "mira_testdb")
	cfg := config.DBConfig{
		Type: dao.DBTypeMysql,
		URL:  mysqldsn,
	}

	if err := InitDBAccessService(&cfg); err != nil {
		panic(err)
	}
	m.Run()
}

func TestRedisSubscribe(t *testing.T) {
	// test add
	err := DBAccessService.AddRedisSubscribe(&types.RedisSubscribe{
		StreamName:   "mystream",
		GroupName:    "mygroup",
		ConsumerName: "myconsumer",
		MessageId:    "0",
	})
	assert.NoError(t, err)

	// test get
	redisSubscribeList, err := DBAccessService.GetRedisSubscribeList(nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(redisSubscribeList))
	assert.Equal(t, 0, redisSubscribeList[0].MessageStatus)

	// update message status
	err = DBAccessService.UpdateRedisSubscribeStatusByMessageId("0", 1)
	assert.NoError(t, err)
	redisSubscribe, err := DBAccessService.GetRedisSubscribeByMessageId("0")
	assert.NoError(t, err)
	assert.Equal(t, 1, redisSubscribe.MessageStatus)

	err = DBAccessService.DeleteRedisSubscribe(redisSubscribeList[0].MessageId)
	assert.NoError(t, err)
}
