/**
 * @Author: xueyanghan
 * @File: dbaccess_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/26 19:11
 */

package dbaccess

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"

	oracle "github.com/godoes/gorm-oracle"

	"github.com/stretchr/testify/assert"
	"github.com/warm3snow/practical-crypto/xin_chuang/config"
	"github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/dao"
	"github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/types"
)

func initMysqlDBAccessService(t *testing.T) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123456", "localhost", 3306, "testdb")
	err := InitDBAccessService(&config.DBConfig{
		Type: dao.DBTypeMysql,
		URL:  url,
	})
	assert.NoError(t, err)
}

func initSqlite3DBAccessService(t *testing.T) {
	err := InitDBAccessService(&config.DBConfig{
		Type: dao.DBTypeSqlite3,
		URL:  "./testdb.db",
	})
	assert.NoError(t, err)
}

func initKingbaseDBAccessService(t *testing.T) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123456", "localhost", 54321, "testdb")
	err := InitDBAccessService(&config.DBConfig{
		Type: dao.DBTypeKingBaseMysql,
		URL:  url,
	})
	assert.NoError(t, err)
}

func initDM8DBAccessService(t *testing.T) {
	url := fmt.Sprintf("dm://%s:%s@%s:%d/%s?autoCommit=true&charset=UTF-8",
		"root", "123456", "localhost", 5236, "testdb")
	err := InitDBAccessService(&config.DBConfig{
		Type: dao.DBTypeDM8,
		URL:  url,
	})
	assert.NoError(t, err)
}

func initOracleDBAccessService(t *testing.T) {
	url := oracle.BuildUrl("localhost", 1521, "testdb",
		"SYSTEM", "123456", nil)
	err := InitDBAccessService(&config.DBConfig{
		Type: dao.DBTypeOracle,
		URL:  url,
	})
	assert.NoError(t, err)

}

func testUserInterface(t *testing.T) {
	userInfo1 := &types.UserInfo{
		UserName: "test",
		Password: hex.EncodeToString(md5.New().Sum([]byte("123456"))),
		Email:    "test@email.com",
		Phone:    "12345678901",
		Role:     0,
	}
	// test add
	err := DBAccessService.AddUser(userInfo1)
	assert.NoError(t, err)

	// test get
	userInfo2, err := DBAccessService.GetUserInfo("test")
	assert.NoError(t, err)
	assert.NotNil(t, userInfo2)
	assert.Equal(t, userInfo1.UserName, userInfo2.UserName)

	// test update
	err = DBAccessService.UpdateUser(&types.UserInfo{
		UserName: "test",
		Email:    "test1@email.com",
	})
	assert.NoError(t, err)
	userInfo3, err := DBAccessService.GetUserInfo("test")
	assert.NoError(t, err)
	assert.Equal(t, "test1@email.com", userInfo3.Email)

	// test delete
	err = DBAccessService.DeleteUser("test")
	assert.NoError(t, err)
	userInfo4, err := DBAccessService.GetUserInfo("test")
	assert.Error(t, err)
	assert.Nil(t, userInfo4)
}

func TestMysqlDBAccessService(t *testing.T) {
	initMysqlDBAccessService(t)
	testUserInterface(t)
}

func TestSqlite3DBAccessService(t *testing.T) {
	initSqlite3DBAccessService(t)
	testUserInterface(t)
}

func TestKingbaseDBAccessService(t *testing.T) {
	initKingbaseDBAccessService(t)
	testUserInterface(t)
}

func TestDM8DBAccessService(t *testing.T) {
	initDM8DBAccessService(t)
	testUserInterface(t)
}

func TestOracleDBAccessService(t *testing.T) {
	initOracleDBAccessService(t)
	testUserInterface(t)
}
