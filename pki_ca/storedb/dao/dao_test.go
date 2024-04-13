/**
 * @Author: xueyanghan
 * @File: sqlite3_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/18 16:08
 */

package dao

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warm3snow/crypto-service-backend/config"
)

func getSubscribeBlockHeight(testdb *DAO, t *testing.T) {
	height, err := testdb.GetSubscribeBlockHeight()
	assert.NoError(t, err)
	t.Logf("initial block height: %v", height)
}

func setSubscribeBlockHeight(testdb *DAO, t *testing.T) {
	err := testdb.SetSubscribeBlockHeight(1)
	assert.NoError(t, err)
	height, err := testdb.GetSubscribeBlockHeight()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), height)

	err = testdb.SetSubscribeBlockHeight(100)
	assert.NoError(t, err)
	height, err = testdb.GetSubscribeBlockHeight()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), height)
}

func TestSqlite3(t *testing.T) {
	config.Conf = &config.Config{
		Common: config.CommonConfig{},
	}
	sqlite3dbName := "proxy.db"
	cfg := config.DBConfig{
		Type: "sqlite3",
		URL:  sqlite3dbName,
	}
	var err error
	testdb, err := New(&cfg)
	assert.NoError(t, err)
	defer func() {
		err = os.Remove(sqlite3dbName)
		assert.NoError(t, err)
	}()

	getSubscribeBlockHeight(testdb, t)
	setSubscribeBlockHeight(testdb, t)
}

func TestMysql(t *testing.T) {
	config.Conf = &config.Config{
		Common: config.CommonConfig{},
	}
	mysqldsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123456", "localhost", 3306, "proxy")
	cfg := config.DBConfig{
		Type: "mysql",
		URL:  mysqldsn,
	}
	var err error
	testdb, err := New(&cfg)
	assert.NoError(t, err)

	getSubscribeBlockHeight(testdb, t)
	setSubscribeBlockHeight(testdb, t)
}
