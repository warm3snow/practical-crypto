package dbaccess_impl

import (
	"github.com/warm3snow/practical-crypto/xin_chuang/config"
	"github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/dao"
)

type DBAccessServiceImpl struct {
	dao *dao.DAO
}

// InitDBAccessService returns a new database instance.
func InitDBAccessService(cfg *config.DBConfig) (*DBAccessServiceImpl, error) {
	daoObj, err := dao.New(cfg)
	if err != nil {
		return nil, err
	}
	return &DBAccessServiceImpl{daoObj}, nil
}
