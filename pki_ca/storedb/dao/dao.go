package dao

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/warm3snow/crypto-service-backend/config"
	"github.com/warm3snow/crypto-service-backend/core/storedb/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DAO struct {
	db  *gorm.DB
	cfg *config.DBConfig
}

// SetSubscribeBlockHeight sets the height of the subscribe block.
func (s DAO) SetSubscribeBlockHeight(height int64) error {
	subscribeBlock := model.SubscribeBlock{
		Height: height,
	}
	result := s.db.Model(&model.SubscribeBlock{}).
		Where("Fchain_name = ?", model.ChainMakerChainName).
		Updates(&subscribeBlock)
	return result.Error
}

// GetSubscribeBlockHeight gets the height of the subscribe block.
func (s DAO) GetSubscribeBlockHeight() (int64, error) {
	var subscribeBlock model.SubscribeBlock
	result := s.db.Model(&model.SubscribeBlock{}).
		Where("Fchain_name = ?", model.ChainMakerChainName).
		First(&subscribeBlock)

	return subscribeBlock.Height, result.Error
}

// New returns a database instance by DB type.
func New(cfg *config.DBConfig) (*DAO, error) {
	var (
		db  *gorm.DB
		err error
	)
	switch cfg.Type {
	case "sqlite3":
		db, err = gorm.Open(sqlite.Open(cfg.URL), &gorm.Config{})
	case "mysql":
		db, err = gorm.Open(mysql.Open(cfg.URL), &gorm.Config{})
	default:
		return nil, fmt.Errorf("db type not support, dbType = %s", cfg.Type)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open db, dbType = %s", cfg.Type)
	}

	// 自动建表
	err = db.AutoMigrate(&model.SubscribeBlock{})
	if err != nil {
		return nil, err
	}

	// 设置链订阅初始高度
	result := db.First(&model.SubscribeBlock{}, "Fchain_name = ?", model.ChainMakerChainName)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result := db.Create(&model.SubscribeBlock{
			ID:        0,
			ChainName: model.ChainMakerChainName,
		})
		if result.Error != nil {
			return nil, err
		}
	} else if result.Error != nil {
		return nil, err
	}

	return &DAO{db: db, cfg: cfg}, nil
}
