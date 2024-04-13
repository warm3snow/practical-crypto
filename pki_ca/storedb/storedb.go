package storedb

import (
	"github.com/warm3snow/crypto-service-backend/config"
	"github.com/warm3snow/crypto-service-backend/core/storedb/dao"
)

var (
	DBAccess DB
)

// DB -
type DB interface {

	// SetSubscribeBlockHeight sets the block height of the last block that has been subscribed.
	SetSubscribeBlockHeight(height int64) error

	// GetSubscribeBlockHeight gets the block height of the last block that has been subscribed.
	GetSubscribeBlockHeight() (int64, error)
}

// InitDBAccess returns a new database instance.
func InitDBAccess(cfg *config.DBConfig) (DB, error) {
	return dao.New(cfg)
}
