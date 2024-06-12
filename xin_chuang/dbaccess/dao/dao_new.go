package dao

import (
	"database/sql"

	"github.com/warm3snow/practical-crypto/xin_chuang/config"
	"github.com/warm3snow/practical-crypto/xin_chuang/dbaccess/model"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DAO 数据库访问对象
type DAO struct {
	db  *gorm.DB
	cfg *config.DBConfig
}

// New returns a database instance by DB type.
func New(cfg *config.DBConfig) (*DAO, error) {
	var (
		db  *gorm.DB
		err error
	)
	switch cfg.Type {
	case DBTypeSqlite:
		db, err = gorm.Open(sqlite.Open(cfg.URL), &gorm.Config{})
	case DBTypeMysql, DBTypeKingBaseMysql:
		db, err = gorm.Open(mysql.Open(cfg.URL), &gorm.Config{})
	case DBTypeMemory:
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	case DBTypePostgres, DBTypeKingBasePgsql:
		db, err = gorm.Open(postgres.Open(cfg.URL), &gorm.Config{})
	default:
		return nil, errors.Errorf("db type not support, dbType = %s", cfg.Type)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open db, dbType = %s", cfg.Type)
	}

	// 自动建表
	err = db.AutoMigrate(&model.RedisSubscribe{})
	if err != nil {
		return nil, err
	}
	return &DAO{db: db, cfg: cfg}, nil
}

// BeginTx 打开数据库事务
func (dao *DAO) BeginTx() (*DAO, error) {
	tx := dao.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// 返回包含事务和相关对象的 SqlDB 结构体
	return &DAO{
		db: tx,
	}, nil
}

// Commit 提交数据库事务
func (dao *DAO) Commit() *gorm.DB {
	return dao.db.Commit()
}

// Rollback 回滚数据库事务
func (dao *DAO) Rollback() *gorm.DB {
	return dao.db.Rollback()
}

func (dao *DAO) Transaction(fc func(db *DAO) error, opts ...*sql.TxOptions) error {
	return dao.db.Transaction(func(tx *gorm.DB) error {
		err := fc(&DAO{db: tx})
		if err != nil {
			return err
		}
		return nil
	}, opts...)
}
